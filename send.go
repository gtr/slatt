package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
)

// BUFFERSIZE is 131,072 bytes of buffer data
const (
	BUFFERSIZE = 131072
	PORT       = "9999"
)

var (
	currDir  string
	currFile string
)

// Main function for send.go
func send(directory, fileName string) {
	if fileName != "" {
		path := directory + "/" + fileName
		if beginServer(path) {
			fmt.Println("Sent")
		} else {
			fmt.Println("Failed to send")
		}
	} else {
		fmt.Println("Please specify filename")
	}
}

// Begin server for tcp connection
func beginServer(path string) bool {
	ip := getThisIP()
	listener, err := net.Listen("tcp", ip+PORT)
	handleErr(err, "Could not begin server")
	log.Println("Server started")

	// Wait for connection
	for {
		connection, err := listener.Accept()
		handleErr(err, "Could not connect to server")
		log.Println("Server connected")

		fsInfo, err := os.Stat(path)
		handleErr(err, "")

		// Handling files and directories
		mode := fsInfo.Mode()
		if mode.IsDir() {
			log.Println("dir")
			handleErr(err, "Could not open directory")
			newZip := zipDir(fsInfo.Name())
			handleFile(newZip, connection)
		} else if mode.IsRegular() {
			log.Println("file")
			currFile, err := os.Open(path)
			handleErr(err, "Could not open file")
			log.Println(currFile.Name())
			handleFile(currFile, connection)
		}
	}
}

// Sending a directory
func handleDirectory(directory *os.File, connection net.Conn) {

	defer directory.Close()
	fi, err := directory.Readdir(-1)
	handleErr(err, "")
	directoryInfo(directory.Name(), connection)
	for _, fi := range fi {
		if fi.Mode().IsRegular() {
			currFile, err := os.Open(fi.Name())
			fmt.Println(fi.Name(), fi.Size(), "bytes")
			handleErr(err, "")
			handleFile(currFile, connection)
		} else {
			currDir, err := os.Open(fi.Name() + "/")
			handleErr(err, "")
			handleDirectory(currDir, connection)
		}
	}
}

func directoryInfo(name string, connection net.Conn) {
	dirName := fixString(name, 64)
	dirSize := fixString("", 10)
	connection.Write([]byte(dirSize))
	connection.Write([]byte(dirName))
}

// Handle one file
func handleFile(file *os.File, connection net.Conn) {
	fileInfo, err := file.Stat()
	handleErr(err, "")
	fileSize := strconv.FormatInt(fileInfo.Size(), 10)
	fileName := fileInfo.Name()

	log.Println("Beginning transfer")
	log.Println("FileName:\t\t", fileName)
	log.Println("FileSize:\t\t", fileSize, " bytes")

	fileSize = fixString(fileSize, 10)
	fileName = fixString(fileName, 64)
	connection.Write([]byte(fileSize))
	connection.Write([]byte(fileName))

	transfer(file, connection)
}

// Sending file in chunks
func transfer(file *os.File, connection net.Conn) {
	currBuffer := make([]byte, BUFFERSIZE)
	for {
		_, err := file.Read(currBuffer)
		if err == io.EOF {
			return
		}
		handleErr(err, "")
		connection.Write(currBuffer)
	}
}
