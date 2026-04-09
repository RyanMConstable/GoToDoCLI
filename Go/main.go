package main

func main() {
	data := Data{}

	file := Setup()

	tasks := ParseFile(file)

	data.flags = SetFlags()

	DecisionTree(data.flags, &tasks, file)

	fields := []string{"Name", "Date Created", "Date Completed", "Due Date"}

	ShowTasks(tasks, fields, data.flags)
}
