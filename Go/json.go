package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func ParseFile(file string) Tasks {
	jsonFile, err := os.Open(file)

	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	var tasks Tasks
	err = json.Unmarshal(byteValue, &tasks)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
	}

	return tasks
}

func WriteJson(tasks Tasks, file string) error {
	data, err := json.Marshal(tasks)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return os.WriteFile(file, data, 0644)
}
