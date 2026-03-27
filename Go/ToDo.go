package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Tasks struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	Name          string `json:"name"`
	Completed     bool   `json:"completed"`
	DateCreated   string `json:"datecreated"`
	DateCompleted string `json:"datecompleted"`
	DueDate       string `json:"duedate"`
}

type Flags struct {
	add       bool
	duedate   bool
	completed bool
}

type FlagValues struct {
	name    string
	duedate string
}

func main() {
	file := Setup()

	//Function to read the file and create a slice of slices to hold the task information
	tasks := ParseFile(file)

	//SET FLAG STRUCTS
	flags, flagvalues := SetFlags()

	//MODIFY JSON STRUCTS
	if flags.add {
		AddTask(&tasks, flags, flagvalues)
		WriteJson(tasks, file)
	}

	if flags.completed {
		CompleteTask(&tasks, flags, flagvalues)
		WriteJson(tasks, file)
	}

	fields := []string{"Name", "Completed", "Date Created", "Date Completed", "Due Date"}

	ShowTasks(tasks, fields)
}

func SetFlags() (Flags, FlagValues) {
	var flags Flags
	var flagvalues FlagValues

	_ = *flag.Bool("help", false, "Show Commands")
	flagAdd := flag.Bool("add", false, "Add a task")
	flagName := flag.String("name", "", "Name of the task")
	flagDuedate := flag.String("duedate", "", "Date the task is due")
	flagCompleted := flag.String("completed", "", "If the task is completed or not")

	flag.Parse()

	flags.add = *flagAdd
	flagvalues.name = *flagName
	flagvalues.duedate = *flagDuedate

	flagSet := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) {
		flagSet[f.Name] = true
	})

	if flagSet["help"] {
		fmt.Println("Help")
		os.Exit(1)
	}

	if flagSet["add"] {
		if flagSet["name"] == false {
			fmt.Println("--name is required in conjunction with the --add flag")
			os.Exit(1)
		}
		if flagSet["duedate"] {
			flags.duedate = true
		}
		if flagSet["completed"] {
			fmt.Println("Cannot use completed and add in the same command")
			os.Exit(1)
		}
	} else {
		flags.add = false
	}

	if flagSet["completed"] {
		flags.completed = true
		flagvalues.name = *flagCompleted
	}

	return flags, flagvalues
}

func Setup() string {
	//CREATING DIRECTORY
	homeDirectory, _ := os.UserHomeDir()
	defaultWorkingDirectory := homeDirectory + "/.todo"
	defaultJsonFile := homeDirectory + "/.todo/todo.json"

	if _, err := os.Stat(defaultWorkingDirectory); os.IsNotExist(err) {
		fmt.Println("Creating directory!")
		os.Mkdir(defaultWorkingDirectory, 0o744)
	}
	//END CREATE DIRECTORY

	//CREATE FILE IF NOT EXISTING
	if _, err := os.Stat(defaultJsonFile); err != nil {
		data := []byte(`{"tasks": []}`)
		err = os.WriteFile(defaultJsonFile, data, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
	//END FILE CREATION

	return defaultJsonFile
}

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

func AddTask(tasks *Tasks, flags Flags, flagvalues FlagValues) {
	var newTask Task

	if flags.duedate {
		newTask.DueDate = flagvalues.duedate
	}
	newTask.Name = flagvalues.name

	tasks.Tasks = append(tasks.Tasks, newTask)
}

func CompleteTask(tasks *Tasks, flags Flags, flagvalues FlagValues) {
	found := false
	for i, value := range tasks.Tasks {
		if flagvalues.name == value.Name {
			tasks.Tasks[i].Completed = true
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Task name does not exist")
	}
}

func WriteJson(tasks Tasks, file string) error {
	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}
	return os.WriteFile(file, data, 0644)
}

func ShowTasks(tasks Tasks, columnNames []string) {
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

	//PRINT OUT HEADER TABLE
	fmt.Println(strings.Repeat("=", lineLength))
	fmt.Println(fmt.Sprintf("|%v%v%v%v%v", lineValues["Name"], lineValues["Due Date"], lineValues["Date Created"], lineValues["Date Completed"], lineValues["Completed"]))
	fmt.Println(strings.Repeat("=", lineLength))
	//END HEADER TABLE

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
