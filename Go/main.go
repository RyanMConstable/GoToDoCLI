package main

func main() {
	data := Data{}

	data.file, data.tasks, data.flags = Setup(), ParseFile(data.file), SetFlags()

	DecisionTree(&data)

	ShowTasks(data, []string{"Name", "Date Created", "Date Completed", "Due Date"})
}
