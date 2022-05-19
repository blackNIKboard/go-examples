package pack

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

var ConnString = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
var Connection *pgx.Conn
var flag bool

func ConnectDB(s string) {
	var err error

	if !flag {
		logrus.Info("init connection")

		Connection, err = pgx.Connect(context.Background(), s)
		if err != nil {
			panic(err)
		}

		flag = true
	} else {
		logrus.Info("ready connection")
	}
}
