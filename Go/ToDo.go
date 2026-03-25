package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Tasks struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	Name          string `json:"name"`
	Completed     bool   `json:"completed"`
	DateCreated   string `json:"datecreated"`
	DateCompleted string `json:"datecompleted"`
	DueDate       string `json:"duedate"`
}

func main() {
	file := Setup()

	//Function to read the file and create a slice of slices to hold the task information
	tasks := ParseFile(file)
	fmt.Println(tasks)
}

func Setup() string {
	//CREATING DIRECTORY
	defaultWorkingDirectory := ".todo"
	defaultJsonFile := ".todo/todo.json"

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

func ParseFile(file string) Tasks {
	jsonFile, err := os.Open(file)

	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	var tasks Tasks
	err = json.Unmarshal(byteValue, &tasks)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
	}

	return tasks
}
