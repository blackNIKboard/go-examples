package main

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/jackc/pgx/v4"

	"testing-database/pack"
)

func main() {
	conn, err := pgx.Connect(context.Background(), pack.ConnString)
	if err != nil {
		panic(err)
	}

	if err1 := CreateTable(context.Background(), conn); err1 != nil {
		panic(err1)
	}
}

func CreateTable(ctx context.Context, conn *pgx.Conn) error {
	exec, err := conn.Exec(ctx, "CREATE TABLE IF NOT EXISTS public.test(column_1 int);")
	if err != nil {
		return err
	}

	spew.Dump(exec)

	return nil
}
