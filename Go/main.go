package main

func main() {
	file := Setup()

	tasks := ParseFile(file)

	flags := SetFlags()

	if flags.add != "" {
		AddTask(&tasks, flags)
		WriteJson(tasks, file)
	}

	if flags.completed != "" {
		CompleteTask(&tasks, flags)
		WriteJson(tasks, file)
	}

	if flags.remove != "" {
		RemoveTask(&tasks, flags)
	}

	if flags.inprogress != "" {
		MarkTaskInProgress(&tasks, flags)
	}

	fields := []string{"Name", "Completed", "Date Created", "Date Completed", "Due Date"}

	ShowTasks(tasks, fields, flags)
}
