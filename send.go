package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

const BUFFERSIZE = 1024

func beginServer(path string) bool {
	listener, err := net.Listen("tcp", "localhost:27001")
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
		termLog("Server connected")
		transfer(connection, path)
	}
}

func transfer(conn net.Conn, path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("could not open file", err)
	}
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Transferring -> ", fileInfo.Name())
	log.Print("File Size    -> ", fileInfo.Size(), "bytes")
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
