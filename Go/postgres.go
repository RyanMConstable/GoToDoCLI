package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool
var ctx = context.Background()

func UnmarshalPostgres(c Config) {
	_, err := pgxpool.New(ctx, fmt.Sprintf("postgresql://%v:%v@%v:%v/%v", c.DB_USER, c.DB_PASSWORD, c.DB_HOST, c.DB_PORT, c.DB_NAME))

	if err != nil {
		fmt.Println(err)
		return
	}
}
