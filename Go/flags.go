package main

import (
	"flag"
	"fmt"
)

func SetFlags(d *Data) Flags {
	var f Flags

	SetBoolFlags(&f.all, "A", "all", false, "Show all tasks")
	SetStringFlags(&f.add, "a", "add", "", "Add a task")
	SetStringFlags(&f.duedate, "d", "duedate", "", "Date the task is due")
	SetStringFlags(&f.completed, "c", "complete", "", "Mark a task complete")
	SetStringFlags(&f.remove, "r", "remove", "", "Remove task by name")
	SetStringFlags(&f.inprogress, "p", "in-progress", "", "Mark a task as in progress")
	SetStringFlags(&f.config, "C", "config", "", "Set the path location for a config file")

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

	return f
}

func SetStringFlags(flagvalue *string, shortFlag string, longFlag string, defaultValue string, description string) {
	flag.StringVar(flagvalue, shortFlag, defaultValue, description)
	flag.StringVar(flagvalue, longFlag, defaultValue, description)
}

func SetBoolFlags(flagvalue *bool, shortFlag string, longFlag string, defaultValue bool, description string) {
	flag.BoolVar(flagvalue, shortFlag, defaultValue, description)
	flag.BoolVar(flagvalue, longFlag, defaultValue, description)
}

func DecisionTree(d *Data) error {
	f := d.flags
	file := d.setup.dataFile
	t := &d.tasks

	CheckStale(t)

	if f.add != "" {
		err := AddTask(d)
		if err != nil {
			return err
		}
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

	if !f.db {
		err := WriteJson(*t, file)
		if err != nil {
			return err
		}
	}

	return nil
}
