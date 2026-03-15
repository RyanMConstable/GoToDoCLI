package main

import (
	"flag"
	"fmt"
	"encoding/json"
	"log"
)

func main() {
	file := IntakeParamaters()
	fmt.Println("Opening ToDo List At:", file)

}

func IntakeParamaters() string {
	file := flag.String("file", "~/.todo", "Storage location for todo information")

	flag.Parse()

	fmt.Println("File Path:", *file)

	positionalArgs := flag.Args()
	fmt.Println("Positional Args:", positionalArgs)

	fmt.Println("END IntakeParameters\n")
	return *file
}
