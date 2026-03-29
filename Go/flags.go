package main

import (
	"flag"
	"fmt"
	"os"
)

func SetFlags() (Flags, FlagValues) {
	var flags Flags
	var flagvalues FlagValues

	var flagAdd bool
	var flagName string
	var flagDuedate string
	var flagCompleted string
	var flagRemove string

	flag.BoolVar(&flagAdd, "add", false, "Add a task")
	flag.BoolVar(&flagAdd, "a", false, "Short version of add")

	flag.StringVar(&flagName, "name", "", "Name of the task")
	flag.StringVar(&flagName, "n", "", "Short version of name")

	flag.StringVar(&flagDuedate, "duedate", "", "Date the task is due")
	flag.StringVar(&flagDuedate, "d", "", "Short version of duedate")

	flag.StringVar(&flagCompleted, "complete", "", "Mark a task complete")
	flag.StringVar(&flagCompleted, "c", "", "Short version of complete")

	flag.StringVar(&flagRemove, "remove", "", "Remove task by name")
	flag.StringVar(&flagRemove, "r", "", "Short version of remove")

	flag.Parse()

	flags.add = flagAdd
	flagvalues.name = flagName
	flagvalues.duedate = flagDuedate

	flagSet := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) {
		flagSet[f.Name] = true
	})

	if flagSet["add"] || flagSet["a"] {
		if flagSet["name"] == false && flagSet["n"] == false {
			fmt.Println("--name is required in conjunction with the --add flag")
			os.Exit(1)
		}
		if flagSet["duedate"] || flagSet["d"] {
			flags.duedate = true
		}
		if flagSet["complete"] || flagSet["c"] {
			fmt.Println("Cannot use completed and add in the same command")
			os.Exit(1)
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
