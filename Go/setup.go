package main

import (
	"fmt"
	"log"
	"os"
)

func Setup() string {
	//CREATING DIRECTORY
	homeDirectory, _ := os.UserHomeDir()
	defaultWorkingDirectory := homeDirectory + "/.todo"
	defaultJsonFile := homeDirectory + "/.todo/todo.json"

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

	return defaultJsonFile
}
