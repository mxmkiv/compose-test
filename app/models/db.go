package models

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func DBInit() *pgx.Conn {

	_ = godotenv.Load("../.env")

	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal("database connect error ", err)
	}

	return conn
}
