package main

func main() {
	data := Data{}

	data.file = Setup()
	data.tasks = ParseFile(data.file)
	data.flags = SetFlags()

	DecisionTree(&data)

	ShowTasks(data, []string{"Name", "Date Created", "Date Completed", "Due Date"})
}
