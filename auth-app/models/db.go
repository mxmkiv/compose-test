package models

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func DBInit() *pgx.Conn {

	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL_DOCKER"))
	if err != nil {
		log.Fatal("database connect error ", err)
	}

	return conn
}
