package main

func main() {
	column_names := []string{"Name", "Date Created", "Date Completed", "Due Date"}
	data := Data{}

	data.file, _ = Setup()
	data.flags, data.config = SetFlags()

	if data.flags.db {
		data.tasks = UnmarshalPostgres(data.config)
	} else {
		data.tasks = ParseFile(data.file)
	}

	DecisionTree(&data)

	ShowTasks(data, column_names)
}
