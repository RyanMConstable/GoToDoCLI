package main

import (
	"time"
)

type Data struct {
	tasks  Tasks
	flags  Flags
	config Config

	file string
}

type Tasks struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	ID            int        `json:"id"            db:"id"`
	Name          string     `json:"name"          db:"name"`
	Completed     bool       `json:"completed"     db:"completed"`
	DateCreated   *time.Time `json:"datecreated"   db:"datecreated"`
	DateCompleted time.Time  `json:"datecompleted" db:"datecompleted"`
	DueDate       *time.Time `json:"duedate"       db:"duedate"`
	InProgress    bool       `json:"inprogress"    db:"inprogress"`
	Stale         bool       `json:"stale"         db:"stale"`
	Parent_ID     int        `json:"parent_id"     db:"parent_id"`
}

type Flags struct {
	add        string
	duedate    string
	completed  string
	remove     string
	all        bool
	inprogress string
	config     string
}

type Config struct {
	DB_USER     string `toml:"DB_USER"`
	DB_PASSWORD string `toml:"DB_PASSWORD"`
	DB_HOST     string `toml:"DB_HOST"`
	DB_PORT     string `toml:"DB_PORT"`
	DB_NAME     string `toml:"DB_NAME"`
}
