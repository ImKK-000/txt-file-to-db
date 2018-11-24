package main

import (
	"bufio"
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

	fileReader := bufio.NewReader(file)
	var line, err = "", error(nil)
	for err == nil {
		fmt.Print(line)
		line, err = fileReader.ReadString('\n')
	}
}
