package main

import (
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	receivedBytes  int64
	bufferFileSize byte
	bufferFilename byte
	bufferFileType byte
)

// Main function for receive.go
func receive() {
	ip := getIP()
	println(ip + PORT)
	connection, err := net.Dial("tcp", ip+PORT)
	handleErr(err, "")
	log.Println("Connected")
	// connection.RemoteAddr()
	if beginClient(connection) {
		log.Println("Downloaded")
	} else {
		log.Println("Unable to download")
	}
}

// Catch a single file transfer
func catchTransfer(fileName string, fileSize int64, connection net.Conn) {
	receivedBytes = 0
	newFile, err := os.Create(fileName)
	handleErr(err, "")

	// Writing to the new file in chunks
	for receivedBytes < fileSize {
		// Less than one chunk left
		if (fileSize - receivedBytes) < BUFFERSIZE {
			lastChunk := fileSize - receivedBytes
			io.CopyN(newFile, connection, lastChunk)
			receivedBytes += lastChunk
			break
		}
		io.CopyN(newFile, connection, BUFFERSIZE)
		receivedBytes += BUFFERSIZE
	}
}

// Begin server for tcp connection
func beginClient(connection net.Conn) bool {
	bufferFileName := make([]byte, 64)
	bufferFileSize := make([]byte, 10)

	for {
		connection.Read(bufferFileSize)
		connection.Read(bufferFileName)
		if string(bufferFileType) == "d" {

		}
		fileSize, _ := strconv.ParseInt(strings.Trim(string(bufferFileSize), ":"), 10, 64)
		fileName := strings.Trim(string(bufferFileName), ":")

		log.Println("Beginning transfer")
		log.Println("Downloading " + fileName)
		log.Println("FileName:\t\t", fileName)
		log.Println("FileSize:\t\t", fileSize, "bytes")

		initialTime := time.Now()

		catchTransfer(fileName, fileSize, connection)

		totalTime := time.Since(initialTime)
		log.Print("Total time:\t\t ", totalTime)
		return true
	}
}
