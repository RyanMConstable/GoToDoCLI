package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file := IntakeParamaters()
	fmt.Println("Opening ToDo List At:", file)

	//TMP
	homeDir, _ := os.UserHomeDir()
	fmt.Println("User home directory:", homeDir)
	//TMP

	ParseFile(file)
}

func IntakeParamaters() string {
	file := flag.String("file", "todo.json", "Storage location for todo information")

	flag.Parse()

	fmt.Println("File Path:", *file)

	positionalArgs := flag.Args()
	fmt.Println("Positional Args:", positionalArgs)

	fmt.Println("END IntakeParameters")
	return *file
}

func ParseFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buf := make([]byte, 1024)

	for {
		data, err := file.Read(buf)

		if err != nil && err != io.EOF {
			panic(err)
		}
		if data == 0 {
			break
		}
		fmt.Printf("%s\n", string(buf[:data]))
	}
}
