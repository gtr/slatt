package main

import (
	"fmt"
	"log"
	"net"
)

func receiveFile() {
	// localhost:27001
	connection, err := net.Dial("tcp", "localhost:27001")
	if err != nil {
		log.Fatal(err)
	}
	if download(connection) {
		fmt.Println("Downloaded")
	} else {
		fmt.Println("Unable to download")
	}
}

func download(connection net.Conn) bool {
	return true
}
