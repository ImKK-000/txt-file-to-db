package main

import (
	"bytes"
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var filename string
	flag.StringVar(&filename, "filename", "input_file", "set filename")
	flag.Parse()

	rawBytes, errorReadFile := ioutil.ReadFile(filename)
	if errorReadFile != nil {
		log.Fatal(errorReadFile)
	}

	buffer := bytes.Split(rawBytes, []byte{13, 10})
	dbConnection, _ := sql.Open("mysql", "root:@tcp(:9999)/txt_db")
	defer dbConnection.Close()
	for index, line := range buffer {
		if len(line) > 0 {
			statementInsert, _ := dbConnection.Prepare("INSERT INTO txt_table(data) VALUES(?)")
			statementInsert.Exec(string(line))
			defer statementInsert.Close()
			fmt.Printf("%d: %s\n", index, string(line))
		}
	}
}
