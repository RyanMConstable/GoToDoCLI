package main

func main() {
	data := Data{}

	file := Setup()

	data.tasks = ParseFile(file)

	data.flags = SetFlags()

	DecisionTree(data.flags, &data.tasks, file)

	fields := []string{"Name", "Date Created", "Date Completed", "Due Date"}

	ShowTasks(data.tasks, fields, data.flags)
}
