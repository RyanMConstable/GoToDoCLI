package main

func main() {
	file := Setup()

	tasks := ParseFile(file)

	flags := SetFlags()

	DecisionTree(flags, &tasks, file)

	fields := []string{"Name", "Completed", "Date Created", "Date Completed", "Due Date"}

	ShowTasks(tasks, fields, flags)
}
