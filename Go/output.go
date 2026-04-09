package main

import (
	"fmt"
	"strings"
)

const (
	reset = "\033[0m"

	bold = "\033[1m"

	green = "\033[32m"

	OliveLeaf  = "\033[38;2;96;108;56m"  // #606C38 - incomplete tasks
	SunlitClay = "\033[38;2;221;161;94m" // #DDA15E - complete tasks
	Copperwood = "\033[38;2;188;108;37m" // #BC6C25 - completed col values

	Mauve    = "\033[38;2;109;104;117m" // #6d6875 - Borders for table
	MintMist = "\033[38;2;203;243;240m" // #CBF3F0 - For in progress

	Sunlit = "\033[38;2;255;60;56m"
)

func ShowTasks(d Data, columnNames []string) {
	tasks := d.tasks
	flags := d.flags

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
	fmt.Println(fmt.Sprintf("|%v%v%v%v", lineValues["Name"], lineValues["Due Date"], lineValues["Date Created"], lineValues["Date Completed"]))
	fmt.Println(strings.Repeat(Mauve+"=", lineLength))

	for _, task := range tasks.Tasks {
		if flags.all == false && task.Completed == true {
			continue
		}
		tasksShown += 1
		taskLine := Mauve + "|" + reset

		taskLine += CalculateTaskLine(SunlitClay, OliveLeaf, Mauve, longestValues["Name"]-len(fmt.Sprintf("%v", task.Name)), task.Completed, task.Name, MintMist, task.InProgress, Sunlit, task.Stale)
		taskLine += CalculateTaskLine(SunlitClay, OliveLeaf, Mauve, longestValues["Due Date"]-len(fmt.Sprintf("%v", task.DueDate)), task.Completed, task.DueDate, MintMist, task.InProgress, Sunlit, task.Stale)
		taskLine += CalculateTaskLine(SunlitClay, OliveLeaf, Mauve, longestValues["Date Created"]-len(fmt.Sprintf("%v", task.DateCreated)), task.Completed, task.DateCreated, MintMist, task.InProgress, Sunlit, task.Stale)
		taskLine += CalculateTaskLine(SunlitClay, OliveLeaf, Mauve, longestValues["Date Completed"]-len(fmt.Sprintf("%v", task.DateCompleted)), task.Completed, task.DateCompleted, MintMist, task.InProgress, Sunlit, task.Stale)

		fmt.Println(taskLine)
	}

	if tasksShown == 0 {
		fmt.Println(fmt.Sprintf(Mauve+"|"+reset+"%v"+Mauve+"|"+reset, strings.Repeat(" ", lineLength-2)))
	}

	fmt.Println(strings.Repeat(Mauve+"="+reset, lineLength))
}

func CalculateTaskLine(completedColor string, wordColor string, borderColor string, totalSpace int, completed bool, value string, inprogressColor string, inprogress bool, staleColor string, stale bool) string {
	numSideGap := totalSpace / 2
	leftGap := numSideGap
	rightGap := numSideGap
	if totalSpace%2 != 0 {
		rightGap = numSideGap + 1
	}

	if completed {
		return fmt.Sprintf(completedColor+"%v%v%v"+reset+borderColor+"|"+reset, strings.Repeat(" ", leftGap), value, strings.Repeat(" ", rightGap))
	} else if inprogress {
		return fmt.Sprintf(inprogressColor+"%v%v%v"+reset+borderColor+"|"+reset, strings.Repeat(" ", leftGap), value, strings.Repeat(" ", rightGap))
	} else if stale {
		return fmt.Sprintf(staleColor+"%v%v%v"+reset+borderColor+"|"+reset, strings.Repeat(" ", leftGap), value, strings.Repeat(" ", rightGap))
	}

	return fmt.Sprintf(wordColor+"%v%v%v"+reset+borderColor+"|"+reset, strings.Repeat(" ", leftGap), value, strings.Repeat(" ", rightGap))
}
