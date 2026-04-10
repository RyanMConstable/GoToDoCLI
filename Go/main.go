package main

func main() {
	data := Data{}

	data.file = Setup()
	data.flags, data.config = SetFlags()
	data.tasks = ParseFile(data.file)
	UnmarshalPostgres(data.config)

	DecisionTree(&data)

	ShowTasks(data, []string{"Name", "Date Created", "Date Completed", "Due Date"})
}
