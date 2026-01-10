package handler

import (
	"bufio"
	"compose-test/app/models"
	"compose-test/app/ui"
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

func UserRegister(scanner *bufio.Scanner, conn *pgx.Conn) {

	fmt.Print("login: ")
	login := ui.GetInput(scanner)
	fmt.Print("password: ")
	password := ui.GetInput(scanner)
	userData := models.NewUser(login, password)
	DBQueryRegister(userData, conn)
}

func DBQueryRegister(userData *models.User, conn *pgx.Conn) {

	queryString := `INSERT INTO users (login, password_hash, created_at)
	VALUES ($1, $2, $3)`

	_, err := conn.Exec(context.Background(), queryString,
		userData.Login,
		userData.PasswordHash,
		time.Now(),
	)
	if err != nil {
		fmt.Println("/////////// ошибка добавления пользователя ///////////")
		fmt.Println(err)
	}

}
