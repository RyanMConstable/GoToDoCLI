package main

func main() {
	column_names := []string{"Name", "Date Created", "Date Completed", "Due Date"}
	data := Data{}

	Setup(&data)
	data.flags, data.config = SetFlags(data.setup)

	//err := LoadConfig(&data)

	if data.flags.db {
		data.tasks = UnmarshalPostgres(data.config)
	} else {
		data.tasks = ParseFile(data.setup.dataFile)
	}

	DecisionTree(&data)

	ShowTasks(data, column_names)
}
