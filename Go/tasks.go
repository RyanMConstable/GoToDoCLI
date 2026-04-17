package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

func AddTask(d *Data) error {
	var newTask Task

	t := &d.tasks
	f := d.flags

	duedate := time.Now().Add(72 * time.Hour)

	if f.duedate != "" {
		duedate, _ = time.Parse("2006-01-02", f.duedate)
	}

	currentTime := time.Now()

	newTask.Name = f.add
	newTask.DueDate = &duedate
	newTask.DateCreated = &currentTime
	newTask.InProgress = false
	newTask.Stale = false

	t.Tasks = append(t.Tasks, newTask)

	if d.flags.db {
		err := AddTaskToPostgres(newTask, *d)
		if err != nil {
			return err
		}
	}

	return nil
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
			t := time.Now()
			tasks.Tasks[i].DateCompleted = &t
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
			if tasks.Tasks[i].InProgress {
				tasks.Tasks[i].InProgress = false
			} else {
				tasks.Tasks[i].InProgress = true
			}
			break
		}
	}
}

func CheckStale(t *Tasks) {
	for i, value := range t.Tasks {
		if value.Stale || value.Completed {
			continue
		}

		dateCreated := value.DateCreated

		dueDate := value.DueDate
		if time.Since(*dateCreated) >= 72*time.Hour {
			t.Tasks[i].Stale = true
		} else if time.Now().After(*dueDate) {
			t.Tasks[i].Stale = true
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
