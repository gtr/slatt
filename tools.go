package main

import (
	"log"
	"net"
	"time"

	"github.com/fatih/color"
)

func termLog(text string) {
	start := time.Now()
	now := start.Format("2008-03-09 15:04:05")
	color.Cyan(now)
}

func fixString(myString string, length int) string {
	for len(myString) < length {
		myString += ":"
	}
	return myString
}

func handleErr(err error, output string) {
	if err != nil {
		log.Fatal(err, output)
	}
}

func getIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip := ipnet.IP.String()
				ip += ":"
				return ip
			}
		}
	}
	return ""
}
