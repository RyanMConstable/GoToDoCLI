package main

import (
	"flag"
	"fmt"

	"github.com/BurntSushi/toml"
)

func SetFlags(s SetupFiles) (Flags, Config) {
	var flags Flags
	var c Config

	SetBoolFlags(&flags.all, "A", "all", false, "Show all tasks")
	SetStringFlags(&flags.add, "a", "add", "", "Add a task")
	SetStringFlags(&flags.duedate, "d", "duedate", "", "Date the task is due")
	SetStringFlags(&flags.completed, "c", "complete", "", "Mark a task complete")
	SetStringFlags(&flags.remove, "r", "remove", "", "Remove task by name")
	SetStringFlags(&flags.inprogress, "p", "in-progress", "", "Mark a task as in progress")
	SetStringFlags(&flags.config, "C", "config", "", "Set the path location for a config file")

	flag.Usage = func() {
		fmt.Println("Usage of todo:")
		fmt.Println("  -add,         -a  Add a task")
		fmt.Println("  -all,         -A  Show all tasks")
		fmt.Println("  -name,        -n  Name of the task")
		fmt.Println("  -complete,    -c  Mark a task complete")
		fmt.Println("  -duedate,     -d  Date the task is due")
		fmt.Println("  -remove,      -r  Remove a task by name")
		fmt.Println("  -in-progress, -p  Mark a task in progress")
		fmt.Println("  -config,      -C  Add a configuration file")
	}

	flag.Parse()

	if flags.config != "" {
		_, err := toml.DecodeFile(flags.config, &c)
		if err == nil {
			flags.db = true
		}
	}

	return flags, c
}

func SetStringFlags(flagvalue *string, shortFlag string, longFlag string, defaultValue string, description string) {
	flag.StringVar(flagvalue, shortFlag, defaultValue, description)
	flag.StringVar(flagvalue, longFlag, defaultValue, description)
}

func SetBoolFlags(flagvalue *bool, shortFlag string, longFlag string, defaultValue bool, description string) {
	flag.BoolVar(flagvalue, shortFlag, defaultValue, description)
	flag.BoolVar(flagvalue, longFlag, defaultValue, description)
}

func DecisionTree(d *Data) {
	f := d.flags
	file := d.setup.dataFile
	t := &d.tasks

	CheckStale(t)

	if f.add != "" {
		AddTask(t, f)
	}

	if f.completed != "" {
		CompleteTask(t, f)
	}

	if f.remove != "" {
		RemoveTask(t, f)
	}

	if f.inprogress != "" {
		MarkTaskInProgress(t, f)
	}

	if f.db {
		fmt.Println("Writing to database")
	} else {
		WriteJson(*t, file)
	}
}
