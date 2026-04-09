package main

func main() {
	data := Data{}

	data.file = Setup()

	data.tasks = ParseFile(data.file)

	data.flags = SetFlags()

	DecisionTree(data.flags, &data.tasks, data.file)

	fields := []string{"Name", "Date Created", "Date Completed", "Due Date"}

	ShowTasks(data.tasks, fields, data.flags)
}
