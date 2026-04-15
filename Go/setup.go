package main

import (
	"fmt"
	"log"
	"os"
)

func Setup(d *Data) {
	homeDirectory, _ := os.UserHomeDir()
	defaultWorkingDirectory := homeDirectory + "/.todo"
	defaultJsonFile := homeDirectory + "/.todo/todo.json"
	defaultConfigFile := homeDirectory + "/.todo/config.yaml"

	if _, err := os.Stat(defaultWorkingDirectory); os.IsNotExist(err) {
		fmt.Println("Creating directory!")
		os.Mkdir(defaultWorkingDirectory, 0o744)
	}

	if _, err := os.Stat(defaultJsonFile); err != nil {
		data := []byte(`{"tasks": []}`)
		err = os.WriteFile(defaultJsonFile, data, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}

	if _, err := os.Stat(defaultConfigFile); err != nil {
		d.setup.confFile = ""
	}

	d.setup.dataFile = defaultJsonFile
}
