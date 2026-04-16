package main

import (
	"fmt"
	"os"
)

func main() {
	column_names := []string{"Name", "Date Created", "Date Completed", "Due Date"}
	data := Data{}

	Setup(&data)
	data.flags = SetFlags(&data)

	err := LoadConfig(&data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = IsDatabaseConfigured(&data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if data.flags.db {
		data.tasks = UnmarshalPostgres(data.config)
	} else {
		data.tasks = ParseFile(data.setup.dataFile)
	}

	DecisionTree(&data)

	ShowTasks(data, column_names)
}
