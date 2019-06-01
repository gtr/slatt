package main

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fatih/color"
)

var (
	charPool  = "abcdefghijklmnopqrstuvwxyz0123456789"
	filenames []string
	currChar  byte
	ind       int
)

// Log with colors
func termLog(text string) {
	start := time.Now()
	now := start.Format("2008-03-09 15:04:05")
	color.Cyan(now)
}

// Fix string to certain length
func fixString(myString string, length int) string {
	for len(myString) < length {
		myString += ":"
	}
	return myString
}

// Error handling can be tedius
func handleErr(err error, output string) {
	if err != nil {
		log.Fatal(err, output)
		return
	}
}

// Get local IP of current machine
func getThisIP() string {
	connection, err := net.Dial("udp", "8.8.8.8:80")
	handleErr(err, "")
	defer connection.Close()

	localAddr := connection.LocalAddr().(*net.UDPAddr)
	ip := localAddr.IP

	return ip.String() + ":"
}

// Get IP of machine sending the file
func getIP() string {
	ip, err := net.ResolveTCPAddr("tcp", "192.168.1.100")
	handleErr(err, "")
	return ip.IP.String() + ":"
	// return "192.168.1.100:"
}

// Wrapper to zipping files
func zipDir(path string) *os.File {
	if string(path[len(path)-1]) == "/" {
		path = path[0 : len(path)-1]
	}

	newZip, err := os.Create(path + ".zip")
	handleErr(err, "")

	path += "/"
	writer := zip.NewWriter(newZip)
	includeFiles(writer, path, "")
	err = writer.Close()
	handleErr(err, "")

	return newZip
}

// Recursively add files to the zipped directory
func includeFiles(writer *zip.Writer, path, zipPath string) {
	files, err := ioutil.ReadDir(path)
	handleErr(err, "")

	for _, file := range files {
		if file.IsDir() {
			currDir := file.Name() + "/"
			newBase := path + currDir
			includeFiles(writer, newBase, currDir)
		} else {
			chunk, err := ioutil.ReadFile(path + file.Name())
			handleErr(err, "Cannot read file")
			fileInZip, err := writer.Create(zipPath + file.Name())
			handleErr(err, "")
			_, err = fileInZip.Write(chunk)
			handleErr(err, "")
		}
	}
}

// Unzip .zip file
func unzip(zipped, output string) {
	reader, err := zip.OpenReader(zipped)
	handleErr(err, "Unable to open zip")

	for _, f := range reader.File {
		rc, err := f.Open()
		handleErr(err, "")
		defer rc.Close()

		fpath := filepath.Join(output, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, f.Mode())

		} else {
			var fdir string
			if lastIndex := strings.LastIndex(fpath, string(os.PathSeparator)); lastIndex > -1 {
				fdir = fpath[:lastIndex]
			}

			err = os.MkdirAll(fdir, f.Mode())
			handleErr(err, "")
			f, err := os.OpenFile(
				fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			handleErr(err, "")

			_, err = io.Copy(f, rc)
			handleErr(err, "")
		}
	}
}
