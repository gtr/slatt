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

var receivedBytes int64

func receiveFile() {
	ip := getIP()
	connection, err := net.Dial("tcp", ip+PORT)
	handleErr(err, "")
	log.Println("Connected")
	if download(connection) {
		log.Println("Downloaded")
	} else {
		log.Println("Unable to download")
	}
}

func download(connection net.Conn) bool {
	bufferFileName := make([]byte, 64)
	bufferFileSize := make([]byte, 10)

	for {
		connection.Read(bufferFileSize)
		connection.Read(bufferFileName)
		fileSize, _ := strconv.ParseInt(strings.Trim(string(bufferFileSize), ":"), 10, 64)
		fileName := strings.Trim(string(bufferFileName), ":")

		log.Println("Beginning transfer")
		log.Println("Downloading " + fileName)
		log.Println("FileName:\t\t", fileName)
		log.Println("FileSize:\t\t", fileSize, " bytes")

		newFile, err := os.Create(fileName)
		handleErr(err, "")
		initialTime := time.Now()

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

		totalTime := time.Since(initialTime)
		log.Print("Time elapsed:\t", totalTime)
		return true
	}
}
