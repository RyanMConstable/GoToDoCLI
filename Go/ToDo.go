package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	file := Setup()
	ParseFile(file)
}

func Setup() string {
	//CREATING DIRECTORY
	homeDirectory, _ := os.UserHomeDir()
	defaultWorkingDirectory := homeDirectory + "/.todo"
	defaultJsonFile := defaultWorkingDirectory + "/todo.json"

	if _, err := os.Stat(defaultWorkingDirectory); os.IsNotExist(err) {
		fmt.Println("Creating directory!")
		os.Mkdir(defaultWorkingDirectory, 0o744)
	}
	//END CREATE DIRECTORY

	//CREATE FILE IF NOT EXISTING
	if _, err := os.Stat(defaultJsonFile); err != nil {
		data := []byte(`{"tasks": []}`)
		err = os.WriteFile(defaultJsonFile, data, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
	//END FILE CREATION

	file := flag.String("file", "todo.json", "Storage location for todo information")

	flag.Parse()

	positionalArgs := flag.Args()
	if len(positionalArgs) != 0 {
		fmt.Println("Args given")
	}

	return *file
}

func ParseFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
}
