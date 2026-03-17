package main

import (
	"flag"
	"fmt"
	"os"
	"log"

	tea "charm.land/bubbletea/v2"
)


func main() {
	file := IntakeParamaters()
	fmt.Println("Opening ToDo List At:", file)

	ParseFile(file)
}

func IntakeParamaters() string {
	file := flag.String("file", "todo.json", "Storage location for todo information")

	flag.Parse()

	fmt.Println("File Path:", *file)

	positionalArgs := flag.Args()
	fmt.Println("Positional Args:", positionalArgs)

	fmt.Println("END IntakeParameters\n")
	return *file
}

func ParseFile(path string) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
}
