package handler

import (
	"compose-test/auth-app/app"
	"compose-test/auth-app/auth"
	"compose-test/auth-app/models"
	"compose-test/auth-app/ui"
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

func UserRegister(a *app.App, conn *pgx.Conn) bool {

	fmt.Print("login: ")
	login := ui.GetInput(a.Scanner)
	fmt.Print("password: ")
	password := ui.GetInput(a.Scanner)
	userData := models.NewUser(login)
	err := DBQueryRegister(userData, conn, password)
	if err != nil {
		fmt.Println("/////////// ошибка добавления пользователя ///////////")
		return false
	}

	a.CurrentUser = userData
	return true
}

func DBQueryRegister(userData *models.User, conn *pgx.Conn, password string) error {

	queryString := `INSERT INTO users (login, password_hash, created_at)
	VALUES ($1, $2, $3)`

	_, err := conn.Exec(context.Background(), queryString,
		userData.Login,
		auth.GetHash(password),
		time.Now(),
	)

	return err

}
