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
	c := d.config

	pool, err := pgxpool.New(ctx, fmt.Sprintf("postgresql://%v:%v@%v:%v/%v", c.DB_USER, c.DB_PASSWORD, c.DB_HOST, c.DB_PORT, c.DB_NAME))

	if err != nil {
		return err
	}
	defer pool.Close()

	_, err = pool.Exec(ctx, `
    INSERT INTO todo (
        name,
        completed,
        datecreated,
        datecompleted,
        duedate,
        inprogress,
        stale,
	parent_id
    ) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
`,
		t.Name,
		t.Completed,
		t.DateCreated,
		t.DateCompleted,
		t.DueDate,
		t.InProgress,
		t.Stale,
		t.Parent_ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func CompleteTaskInPostgres(n string, d Data) error {
	c := d.config

	pool, err := pgxpool.New(ctx, fmt.Sprintf("postgresql://%v:%v@%v:%v/%v", c.DB_USER, c.DB_PASSWORD, c.DB_HOST, c.DB_PORT, c.DB_NAME))

	if err != nil {
		return err
	}
	defer pool.Close()

	var id string

	err = pool.QueryRow(ctx, `SELECT id FROM todo WHERE name = $1`, n).Scan(&id)
	if err != nil {
		return err
	}

	//Now we need to update that row as complete
	_, err = pool.Exec(ctx, `UPDATE todo SET completed = true, datecompleted = NOW() WHERE id = $1`, id)

	return nil
}

func UpdateOneColumn(t Task, d Data) error {
	return nil
}
