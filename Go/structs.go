package main

type Data struct {
	tasks  Tasks
	flags  Flags
	config Config
}

type Tasks struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	Name          string `json:"name"`
	Completed     bool   `json:"completed"`
	DateCreated   string `json:"datecreated"`
	DateCompleted string `json:"datecompleted"`
	DueDate       string `json:"duedate"`
	InProgress    bool   `json:"inprogress"`
	Stale         bool   `json:"stale"`
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
	db_user     string
	db_password string
	db_host     string
	db_port     string
	db_name     string
}
