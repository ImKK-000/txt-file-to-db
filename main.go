package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var filename string
	flag.StringVar(&filename, "filename", "input_file", "set filename")
	flag.Parse()

	file, errorOpenFile := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	defer file.Close()
	if errorOpenFile != nil {
		log.Fatal(errorOpenFile)
	}

	fileInfo, errorGetFileInfo := file.Stat()
	if errorGetFileInfo != nil {
		log.Fatal(errorGetFileInfo)
	}

	buffer := make([]byte, fileInfo.Size())
	_, errorReadData := file.Read(buffer)
	if errorReadData != nil {
		log.Fatal(errorReadData)
	}

	fmt.Print(string(buffer))
}
