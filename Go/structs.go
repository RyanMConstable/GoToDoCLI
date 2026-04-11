package main

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
	Name          string `json:"name"          db:"name"`
	Completed     bool   `json:"completed"     db:"completed"`
	DateCreated   string `json:"datecreated"   db:"datecreated"`
	DateCompleted string `json:"datecompleted" db:"datecompleted"`
	DueDate       string `json:"duedate"       db:"duedate"`
	InProgress    bool   `json:"inprogress"    db:"inprogress"`
	Stale         bool   `json:"stale"         db:"stale"`
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
