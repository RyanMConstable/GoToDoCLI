package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	file := IntakeParamaters()
	fmt.Println("Opening ToDo List At:", file)

	ParseFile(file)
}

func IntakeParamaters() string {
	//CREATING DIRECTORY
	homeDirectory, _ := os.UserHomeDir()
	defaultWorkingDirectory = homeDirectory + "/.todo"

	if _, err := os.Stat(defaultWorkingDirectory); os.IsNotExist(err) {
		fmt.Println("Creating directory!")
		os.Mkdir(defaultWorkingDirectory, 0o644)
	}
	//END CREATE DIRECTORY

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
}
