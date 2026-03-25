package main

import (
	"encoding/json"
	"flag"
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

type Flags struct {
	add     bool
	duedate bool
}

type FlagValues struct {
	name    string
	duedate string
}

func main() {
	file := Setup()

	//Function to read the file and create a slice of slices to hold the task information
	tasks := ParseFile(file)

	//SET FLAG STRUCTS
	flags, flagvalues := SetFlags()

	//MODIFY JSON STRUCTS
	if flags.add {
		AddTask(&tasks, flags, flagvalues)
	}
}

func SetFlags() (Flags, FlagValues) {
	var flags Flags
	var flagvalues FlagValues

	flagAdd := flag.Bool("add", false, "Add a task")
	flagName := flag.String("name", "", "Name of the task")
	flagDuedate := flag.String("duedate", "", "Date the task is due")

	flag.Parse()

	flags.add = *flagAdd
	flagvalues.name = *flagName
	flagvalues.duedate = *flagDuedate

	flagSet := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) {
		flagSet[f.Name] = true
	})

	if flagSet["add"] {
		if flagSet["name"] == false {
			fmt.Println("--name is required in conjunction with the --add flag")
			os.Exit(1)
		}
		if flagSet["duedate"] {
			flags.duedate = true
		}
	} else {
		flags.add = false
	}

	return flags, flagvalues
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

func AddTask(tasks *Tasks, flags Flags, flagvalues FlagValues) {
	var newTask Task

	if flags.duedate {
		newTask.DueDate = flagvalues.duedate
	}
	newTask.Name = flagvalues.name

	tasks.Tasks = append(tasks.Tasks, newTask)
}
