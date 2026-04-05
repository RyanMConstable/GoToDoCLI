package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

func AddTask(tasks *Tasks, flags Flags) {
	var newTask Task
	var duedate string

	if flags.duedate != "" {
		duedate = flags.duedate
	} else {
		duedate = "----"
	}

	//TIME
	currentTime := time.Now().Format("2006-01-02")

	newTask.Name = flags.add
	newTask.DueDate = duedate
	newTask.DateCreated = currentTime
	newTask.InProgress = false

	tasks.Tasks = append(tasks.Tasks, newTask)
}

func CompleteTask(tasks *Tasks, flags Flags) {
	name, err := SearchForTaskName(*tasks, flags.completed)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i, value := range tasks.Tasks {
		if name == value.Name {
			tasks.Tasks[i].Completed = true
			tasks.Tasks[i].DateCompleted = time.Now().Format("2006-01-02")
			break
		}
	}

}

func RemoveTask(tasks *Tasks, flags Flags) {
	for i, value := range tasks.Tasks {
		if value.Name == flags.remove {
			fmt.Println("Removing!")
		}
		fmt.Println(i, value)
	}
}

func MarkTaskInProgress(tasks *Tasks, flags Flags) {
	name, err := SearchForTaskName(*tasks, flags.inprogress)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i, value := range tasks.Tasks {
		if value.Name == name {
			tasks.Tasks[i].InProgress = true
			break
		}
	}
}

func CheckStale(t *Tasks) {
	for i, value := range t.Tasks {
		if value.Stale || value.Completed {
			continue
		}

		dateCreated, err := time.Parse(time.DateOnly, value.DateCreated)
		if err != nil {
			fmt.Println("ERROR")
		}

		dueDate, err := time.Parse(time.DateOnly, value.DueDate)
		if err != nil {
			if time.Since(dateCreated) >= 72*time.Hour {
				t.Tasks[i].Stale = true
			}
		} else {
			if time.Now().After(dueDate) {
				t.Tasks[i].Stale = true
			}
		}
	}
}

func SearchForTaskName(t Tasks, search string) (string, error) {
	tasksFound := []string{}

	for _, v := range t.Tasks {
		if strings.HasPrefix(v.Name, search) && v.Completed == false {
			tasksFound = append(tasksFound, v.Name)
		}

	}

	if len(tasksFound) == 0 {
		return "", errors.New("No tasks found")
	} else if len(tasksFound) > 1 {
		return "", errors.New("Too many tasks found")
	}

	return tasksFound[0], nil
}
