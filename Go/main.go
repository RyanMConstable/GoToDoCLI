package main

import "fmt"

func main() {
	data := Data{}

	data.file = Setup()
	data.tasks = ParseFile(data.file)
	data.flags, data.config = SetFlags()
	fmt.Println(data.config)

	DecisionTree(&data)

	ShowTasks(data, []string{"Name", "Date Created", "Date Completed", "Due Date"})
}
