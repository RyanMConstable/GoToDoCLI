package main

import (
	"fmt"
	"os"
)

const (
	VERSION = "BETA:0.0.0"
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

	err = DecisionTree(&data, column_names)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
