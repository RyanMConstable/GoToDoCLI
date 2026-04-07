package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool
var ctx = context.Background()

func unmarshalPostgres(c Config) {
	_, err := pgxpool.New(ctx, fmt.Sprintf("postgresql://%v:%v@%v:%v/%v", c.db_user, c.db_password, c.db_host, c.db_port, c.db_name))

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Connection successful")
}
