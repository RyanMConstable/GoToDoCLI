package main

func main() {
	column_names := []string{"Name", "Date Created", "Date Completed", "Due Date"}
	data := Data{}

	data.file = Setup()
	data.flags, data.config = SetFlags()
	data.tasks = ParseFile(data.file)

	//Currently under testing
	psqldata := Data{}
	psqldata.file = data.file
	psqldata.flags = data.flags
	psqldata.config = data.config

	psqldata.tasks = UnmarshalPostgres(data.config)
	ShowTasks(psqldata, column_names)
	//End testing environment

	DecisionTree(&data)

	ShowTasks(data, column_names)
}
