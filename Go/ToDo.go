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
	defaultWorkingDirectory := homeDirectory + "/.todo"
	defaultJsonFile := defaultWorkingDirectory + "/todo.json"

	if _, err := os.Stat(defaultWorkingDirectory); os.IsNotExist(err) {
		fmt.Println("Creating directory!")
		os.Mkdir(defaultWorkingDirectory, 0o744)
	}
	//END CREATE DIRECTORY

	//CREATE FILE IF NOT EXISTING
	if _, err := os.Stat(defaultJsonFile); err != nil {
		fmt.Println("File doesn't exist!")
		f, err := os.Create(defaultJsonFile)
		defer f.Close()
		if err != nil {
			fmt.Println(err)
		}

		data := []byte(`{"tasks": []}`)
		err = os.WriteFile(defaultJsonFile, data, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
	}
	//END FILE CREATION

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
