package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
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
	add       bool
	duedate   bool
	completed bool
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
		WriteJson(tasks, file)
	}

	if flags.completed {
		CompleteTask(&tasks, flags, flagvalues)
		WriteJson(tasks, file)
	}

	fields := []string{"Name", "Completed", "Date Created", "Date Completed", "Due Date"}

	ShowTasks(tasks, fields)
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
	} else {
		newTask.DueDate = "----"
	}
	newTask.Name = flagvalues.name

	//TIME
	currentTime := time.Now().Format("2006-01-02")
	newTask.DateCreated = currentTime

	tasks.Tasks = append(tasks.Tasks, newTask)
}

func CompleteTask(tasks *Tasks, flags Flags, flagvalues FlagValues) {
	found := false
	for i, value := range tasks.Tasks {
		if flagvalues.name == value.Name {
			tasks.Tasks[i].Completed = true
			tasks.Tasks[i].DateCompleted = time.Now().Format("2006-01-02")
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Task name does not exist")
	}
}

func WriteJson(tasks Tasks, file string) error {
	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}
	return os.WriteFile(file, data, 0644)
}
