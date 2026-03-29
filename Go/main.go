package main

import ()

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
	remove    bool
}

type FlagValues struct {
	name    string
	duedate string
}

func main() {
	file := Setup()

	tasks := ParseFile(file)

	flags, flagvalues := SetFlags()

	if flags.add {
		AddTask(&tasks, flags, flagvalues)
		WriteJson(tasks, file)
	}

	if flags.completed {
		CompleteTask(&tasks, flags, flagvalues)
		WriteJson(tasks, file)
	}

	if flags.remove {
		RemoveTask(&tasks, flags, flagvalues)
	}

	fields := []string{"Name", "Completed", "Date Created", "Date Completed", "Due Date"}

	ShowTasks(tasks, fields)
}
