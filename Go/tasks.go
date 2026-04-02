package main

import (
	"fmt"
	"time"
)

func AddTask(tasks *Tasks, flags Flags) {
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

func CompleteTask(tasks *Tasks, flags Flags) {
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

func RemoveTask(tasks *Tasks, flags Flags) {
	for i, value := range tasks.Tasks {
		if value.Name == flagvalues.name {
			fmt.Println("Removing!")
		}
		fmt.Println(i, value)
	}
}
