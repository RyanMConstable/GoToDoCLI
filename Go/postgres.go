package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool
var ctx = context.Background()

func UnmarshalPostgres(c Config) Tasks {
	pool, err := pgxpool.New(ctx, fmt.Sprintf("postgresql://%v:%v@%v:%v/%v", c.DB_USER, c.DB_PASSWORD, c.DB_HOST, c.DB_PORT, c.DB_NAME))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rows, err := pool.Query(ctx, `SELECT * FROM todo`)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	tasks, err := pgx.CollectRows(rows, pgx.RowToStructByName[Task])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return Tasks{Tasks: tasks}
}

func AddTaskToPostgres(t Task, d Data) error {
	return nil
}
