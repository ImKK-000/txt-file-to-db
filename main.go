package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	var filename string
	flag.StringVar(&filename, "filename", "input_file", "set filename")
	flag.Parse()

	rawBytes, errorReadFile := ioutil.ReadFile(filename)
	if errorReadFile != nil {
		log.Fatal(errorReadFile)
	}

	buffer := bytes.NewBuffer(rawBytes)
	line, err := "", error(nil)
	for err == nil {
		fmt.Print(line)
		line, err = buffer.ReadString('\n')
	}
}
