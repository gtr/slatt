package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
)

// BUFFERSIZE is 1024 bytes of buffer data
const (
	BUFFERSIZE = 1024
	PORT       = "9999"
)

func beginServer(path string) bool {
	ip := getIP()
	listener, err := net.Listen("tcp", ip+PORT)
	if err != nil {
		log.Fatal("could not begin server", err)
		return false
	}
	log.Println("Server started")
	for {
		// Wait for connection
		connection, err := listener.Accept()
		if err != nil {
			log.Fatal("could not connect to server: ", err)
		}
		log.Println("Server connected")
		transfer(connection, path)
	}
}

func transfer(connection net.Conn, path string) {
	file, err := os.Open(path)
	handleErr(err, "Could not open file")
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

	// Sending in chunks
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

func sendFile(directory, fileName string) {
	if fileName != "" {
		path := directory + "/" + fileName
		if beginServer(path) {
			fmt.Println("sent")
		} else {
			fmt.Println("failed to send")
		}
	} else {
		fmt.Println("Please specify filename")
	}
}
