package handler

import (
	"compose-test/auth-app/app"
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
	userData := models.NewUser(login, password)
	err := DBQueryRegister(userData, conn)
	if err != nil {
		fmt.Println("/////////// ошибка добавления пользователя ///////////")
		return false
	}

	a.CurrentUser = userData
	return true
}

func DBQueryRegister(userData *models.User, conn *pgx.Conn) error {

	queryString := `INSERT INTO users (login, password_hash, created_at)
	VALUES ($1, $2, $3)`

	_, err := conn.Exec(context.Background(), queryString,
		userData.Login,
		userData.PasswordHash,
		time.Now(),
	)

	return err

}
