package main

import (
	"flag"
	"fmt"
)

func SetFlags() Flags {
	var flags Flags

	SetBoolFlags(&flags.all, "A", "all", false, "Show all tasks")

	flag.StringVar(&flags.add, "add", "", "Add a task")
	flag.StringVar(&flags.add, "a", "", "Short version of add")

	flag.StringVar(&flags.duedate, "duedate", "", "Date the task is due")
	flag.StringVar(&flags.duedate, "d", "", "Short version of duedate")

	flag.StringVar(&flags.completed, "complete", "", "Mark a task complete")
	flag.StringVar(&flags.completed, "c", "", "Short version of complete")

	flag.StringVar(&flags.remove, "remove", "", "Remove task by name")
	flag.StringVar(&flags.remove, "r", "", "Short version of remove")

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

	flagSet := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) {
		flagSet[f.Name] = true
	})

	return flags
}

func SetStringFlags(flagvalue *string, shortFlag string, longFlag string, defaultValue string, description string) {
	flag.StringVar(flagvalue, shortFlag, defaultValue, description)
	flag.StringVar(flagvalue, longFlag, defaultValue, description)
}

func SetBoolFlags(flagvalue *bool, shortFlag string, longFlag string, defaultValue bool, description string) {
	flag.BoolVar(flagvalue, shortFlag, defaultValue, description)
	flag.BoolVar(flagvalue, longFlag, defaultValue, description)
}
