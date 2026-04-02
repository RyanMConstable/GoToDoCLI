package main

import (
	"fmt"
	"strings"
)

const (
	reset = "\033[0m"

	bold = "\033[1m"

	green = "\033[32m"

	Cornsilk   = "\033[38;2;254;250;224m" // #FEFAE0 - headers
	OliveLeaf  = "\033[38;2;96;108;56m"   // #606C38 - incomplete tasks
	SunlitClay = "\033[38;2;221;161;94m"  // #DDA15E - complete tasks
	Copperwood = "\033[38;2;188;108;37m"  // #BC6C25 - completed col values

	SageGreen   = "\033[38;2;175;213;170m" // ##AFD5AA - For border
	BlackForest = "\033[38;2;40;54;24m"    // #283618  - dark accents
	Mauve       = "\033[38;2;109;104;117m"
)

func ShowTasks(tasks Tasks, columnNames []string, flags Flags) {
	tasksShown := 0

	longestValues := map[string]int{}

	for _, columnName := range columnNames {
		longestValues[columnName] = len(columnName) + 2
		for _, value := range tasks.Tasks {
			switch columnName {
			case "Name":
				if value.Completed == false || flags.all == true {
					if longestValues[columnName] < len(value.Name)+2 {
						longestValues[columnName] = len(value.Name) + 2
					}
				}

			case "Date Created":
				if value.Completed == false || flags.all == true {
					if longestValues[columnName] < len(value.DateCreated)+2 {
						longestValues[columnName] = len(value.DateCreated) + 2
					}
				}
			case "Date Completed":
				if value.Completed == false || flags.all == true {
					if longestValues[columnName] < len(value.DateCompleted)+2 {
						longestValues[columnName] = len(value.DateCompleted) + 2
					}
				}
			case "Due Date":
				if value.Completed == false || flags.all == true {
					if longestValues[columnName] < len(value.DueDate)+2 {
						longestValues[columnName] = len(value.DueDate) + 2
					}
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
		lineValues[k] = fmt.Sprintf(Copperwood+bold+"%v%v%v"+reset+Mauve+"|"+reset, strings.Repeat(" ", leftGap), k, strings.Repeat(" ", rightGap))
	}

	fmt.Println(strings.Repeat(Mauve+"=", lineLength))
	fmt.Println(fmt.Sprintf("|%v%v%v%v%v", lineValues["Name"], lineValues["Due Date"], lineValues["Date Created"], lineValues["Date Completed"], lineValues["Completed"]))
	fmt.Println(strings.Repeat(Mauve+"=", lineLength))

	for _, task := range tasks.Tasks {
		if flags.all == false && task.Completed == true {
			continue
		}
		tasksShown += 1
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
		if task.Completed {
			taskLine += fmt.Sprintf(Mauve+"|"+SunlitClay+"%v%v%v"+reset+Mauve+"|"+reset, strings.Repeat(" ", leftGap), task.Name, strings.Repeat(" ", rightGap))
		} else {
			taskLine += fmt.Sprintf(Mauve+"|"+OliveLeaf+"%v%v%v"+reset+Mauve+"|"+reset, strings.Repeat(" ", leftGap), task.Name, strings.Repeat(" ", rightGap))
		}

		totalSpace = longestValues["Due Date"] - len(task.DueDate)
		numSideGap = totalSpace / 2
		leftGap = numSideGap
		rightGap = numSideGap
		if totalSpace%2 != 0 {
			rightGap = numSideGap + 1
		}
		if task.Completed {
			taskLine += fmt.Sprintf(SunlitClay+"%v%v%v"+reset+Mauve+"|"+reset, strings.Repeat(" ", leftGap), task.DueDate, strings.Repeat(" ", rightGap))
		} else {
			taskLine += fmt.Sprintf(OliveLeaf+"%v%v%v"+reset+Mauve+"|"+reset, strings.Repeat(" ", leftGap), task.DueDate, strings.Repeat(" ", rightGap))
		}

		totalSpace = longestValues["Date Created"] - len(task.DateCreated)
		numSideGap = totalSpace / 2
		leftGap = numSideGap
		rightGap = numSideGap
		if totalSpace%2 != 0 {
			rightGap = numSideGap + 1
		}
		if task.Completed {
			taskLine += fmt.Sprintf(SunlitClay+"%v%v%v"+reset+Mauve+"|"+reset, strings.Repeat(" ", leftGap), task.DateCreated, strings.Repeat(" ", rightGap))
		} else {
			taskLine += fmt.Sprintf(OliveLeaf+"%v%v%v"+reset+Mauve+"|"+reset, strings.Repeat(" ", leftGap), task.DateCreated, strings.Repeat(" ", rightGap))
		}

		totalSpace = longestValues["Date Completed"] - len(task.DateCompleted)
		numSideGap = totalSpace / 2
		leftGap = numSideGap
		rightGap = numSideGap
		if totalSpace%2 != 0 {
			rightGap = numSideGap + 1
		}
		if task.Completed {
			taskLine += fmt.Sprintf(SunlitClay+"%v%v%v"+reset+Mauve+"|"+reset, strings.Repeat(" ", leftGap), task.DateCompleted, strings.Repeat(" ", leftGap))
		} else {
			taskLine += fmt.Sprintf(OliveLeaf+"%v%v%v"+reset+Mauve+"|"+reset, strings.Repeat(" ", leftGap), task.DateCompleted, strings.Repeat(" ", leftGap))
		}

		totalSpace = longestValues["Completed"] - len(fmt.Sprintf("%v", task.Completed))
		numSideGap = totalSpace / 2
		leftGap = numSideGap
		rightGap = numSideGap
		if totalSpace%2 != 0 {
			rightGap = numSideGap + 1
		}
		if task.Completed {
			taskLine += fmt.Sprintf(SunlitClay+"%v%v%v"+reset+Mauve+"|"+reset, strings.Repeat(" ", leftGap), task.Completed, strings.Repeat(" ", rightGap))
		} else {
			taskLine += fmt.Sprintf(OliveLeaf+"%v%v%v"+reset+Mauve+"|"+reset, strings.Repeat(" ", leftGap), task.Completed, strings.Repeat(" ", rightGap))
		}

		fmt.Println(taskLine)
	}

	if tasksShown == 0 {
		fmt.Println(fmt.Sprintf(Mauve+"|"+reset+"%v"+Mauve+"|"+reset, strings.Repeat(" ", lineLength-2)))
	}

	fmt.Println(strings.Repeat(Mauve+"="+reset, lineLength))
}
