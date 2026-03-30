package main

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

	ShowTasks(tasks, fields, flags)
}
