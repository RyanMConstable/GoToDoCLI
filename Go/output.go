package main

import (
	"fmt"
	"strings"
)

func ShowTasks(tasks Tasks, columnNames []string) {
	if len(tasks.Tasks) <= 0 {
		return
	}

	longestValues := map[string]int{}

	for _, columnName := range columnNames {
		longestValues[columnName] = len(columnName) + 2
		for _, value := range tasks.Tasks {
			switch columnName {
			case "Name":
				if longestValues[columnName] < len(value.Name)+2 {
					longestValues[columnName] = len(value.Name) + 2
				}
			case "Date Created":
				if longestValues[columnName] < len(value.DateCreated)+2 {
					longestValues[columnName] = len(value.DateCreated) + 2
				}
			case "Date Completed":
				if longestValues[columnName] < len(value.DateCompleted)+2 {
					longestValues[columnName] = len(value.DateCompleted) + 2
				}
			case "Due Date":
				if longestValues[columnName] < len(value.DueDate)+2 {
					longestValues[columnName] = len(value.DueDate) + 2
				}
			}
		}
	}

	lineLength := 1

	lineValues := map[string]string{}

	for k, value := range longestValues {
		lineLength += (value) + 1

		value -= len(k)
		leftGap, rightGap := value/2, value/2
		if value%2 != 0 {
			rightGap = (value / 2) + 1
		}
		lineValues[k] = fmt.Sprintf("%v%v%v|", strings.Repeat(" ", leftGap), k, strings.Repeat(" ", rightGap))
	}

	fmt.Println(strings.Repeat("=", lineLength))
	fmt.Println(fmt.Sprintf("|%v%v%v%v%v", lineValues["Name"], lineValues["Due Date"], lineValues["Date Created"], lineValues["Date Completed"], lineValues["Completed"]))
	fmt.Println(strings.Repeat("=", lineLength))

	for _, task := range tasks.Tasks {
		taskLine := ""
		leftGap, rightGap := 0, 0
		numSideGap := 0

		totalSpace := longestValues["Name"] - len(task.Name)
		numSideGap = totalSpace / 2
		leftGap = numSideGap
		rightGap = numSideGap
		if totalSpace%2 != 0 {
			rightGap = numSideGap + 1
		}
		taskLine += fmt.Sprintf("|%v%v%v|", strings.Repeat(" ", leftGap), task.Name, strings.Repeat(" ", rightGap))

		totalSpace = longestValues["Due Date"] - len(task.DueDate)
		numSideGap = totalSpace / 2
		leftGap = numSideGap
		rightGap = numSideGap
		if totalSpace%2 != 0 {
			rightGap = numSideGap + 1
		}
		taskLine += fmt.Sprintf("%v%v%v|", strings.Repeat(" ", leftGap), task.DueDate, strings.Repeat(" ", rightGap))

		totalSpace = longestValues["Date Created"] - len(task.DateCreated)
		numSideGap = totalSpace / 2
		leftGap = numSideGap
		rightGap = numSideGap
		if totalSpace%2 != 0 {
			rightGap = numSideGap + 1
		}
		taskLine += fmt.Sprintf("%v%v%v|", strings.Repeat(" ", leftGap), task.DateCreated, strings.Repeat(" ", rightGap))

		totalSpace = longestValues["Date Completed"] - len(task.DateCompleted)
		numSideGap = totalSpace / 2
		leftGap = numSideGap
		rightGap = numSideGap
		if totalSpace%2 != 0 {
			rightGap = numSideGap + 1
		}
		taskLine += fmt.Sprintf("%v%v%v|", strings.Repeat(" ", leftGap), task.DateCompleted, strings.Repeat(" ", leftGap))

		totalSpace = longestValues["Completed"] - len(fmt.Sprintf("%v", task.Completed))
		numSideGap = totalSpace / 2
		leftGap = numSideGap
		rightGap = numSideGap
		if totalSpace%2 != 0 {
			rightGap = numSideGap + 1
		}
		taskLine += fmt.Sprintf("%v%v%v|", strings.Repeat(" ", leftGap), task.Completed, strings.Repeat(" ", rightGap))

		fmt.Println(taskLine)
	}

	fmt.Println(strings.Repeat("=", lineLength))
}
