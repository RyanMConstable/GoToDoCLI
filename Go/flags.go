package main

import (
	"flag"
	"fmt"
)

func SetFlags() (Flags, FlagValues) {
	var flags Flags
	var flagvalues FlagValues

	var flagAdd string
	var flagName string
	var flagDuedate string
	var flagCompleted string
	var flagRemove string

	flag.BoolVar(&flags.all, "all", false, "Show all tasks")
	flag.BoolVar(&flags.all, "A", false, "Show all tasks")

	flag.StringVar(&flagAdd, "add", "", "Add a task")
	flag.StringVar(&flagAdd, "a", "", "Short version of add")

	flag.StringVar(&flagName, "n", "", "Short version of name")
	flag.StringVar(&flagName, "name", "", "Name of the task")

	flag.StringVar(&flagDuedate, "duedate", "", "Date the task is due")
	flag.StringVar(&flagDuedate, "d", "", "Short version of duedate")

	flag.StringVar(&flagCompleted, "complete", "", "Mark a task complete")
	flag.StringVar(&flagCompleted, "c", "", "Short version of complete")

	flag.StringVar(&flagRemove, "remove", "", "Remove task by name")
	flag.StringVar(&flagRemove, "r", "", "Short version of remove")

	flag.Usage = func() {
		fmt.Println("Usage of todo:")
		fmt.Println("  -add, -a          Add a task")
		fmt.Println("  -all, -A          Show all tasks")
		fmt.Println("  -name, -n         Name of the task")
		fmt.Println("  -complete, -c     Mark a task complete")
		fmt.Println("  -duedate, -d      Date the task is due")
		fmt.Println("  -remove, -r       Remove a task by name")
	}

	flag.Parse()

	flagvalues.name = flagName
	flagvalues.duedate = flagDuedate

	flagSet := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) {
		flagSet[f.Name] = true
	})

	if flagSet["add"] || flagSet["a"] {
		flags.add = true
		flagvalues.name = flagAdd
		if flagSet["duedate"] || flagSet["d"] {
			flags.duedate = true
		}
	} else {
		flags.add = false
	}

	if flagSet["complete"] || flagSet["c"] {
		flags.completed = true
		flagvalues.name = flagCompleted
	}

	if flagSet["remove"] || flagSet["r"] {
		flags.remove = true
		flagvalues.name = flagRemove
	}

	return flags, flagvalues
}
