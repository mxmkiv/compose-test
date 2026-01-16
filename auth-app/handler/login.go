package handler

import (
	"compose-test/auth-app/app"
	"compose-test/auth-app/models"
	"compose-test/auth-app/ui"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

func UserLogin(a *app.App, conn *pgx.Conn) bool {

	fmt.Print("login: ")
	login := ui.GetInput(a.Scanner)
	fmt.Print("password: ")
	password := ui.GetInput(a.Scanner)

	exists := CheckUserExists(conn, login)
	if exists == false {
		ui.ViewClear()
		fmt.Println("Пользователя с таким именем не существует")
		return false
	}

	validPass := CheckPassword(conn, password, login)
	if validPass == false {
		ui.ViewClear()
		fmt.Println("Неправильный пароль")
		return false
	}

	a.CurrentUser = models.NewUser(login, "")

	/*
		необходимо удаление поля passwordHash
		присваивать не имя одного поля тк CurrentUser = nil
		присваивание структуры пользователя
	*/

	return true

}

/*
	добавить повторный ввод пароля 3 попытки
	объединить функции индентицикации и аутентицикации в одну функцию проверки

*/

func CheckUserExists(conn *pgx.Conn, login string) bool {

	// переписать через EXISTS, возвращать не просто тру фолс а результат запроса

	queryString := `SELECT login FROM users WHERE login=$1`

	var existsFlag string
	err := conn.QueryRow(context.Background(), queryString, login).Scan(&existsFlag)
	if err != nil {
		if err == pgx.ErrNoRows {
			return false
		}
		log.Printf("database: query error %s\n", err)
	}

	return true

}

func CheckPassword(conn *pgx.Conn, password string, login string) bool {

	queryString := `SELECT password_hash FROM users WHERE login=$1`

	var passwordHash []byte
	err := conn.QueryRow(context.Background(), queryString, login).Scan(&passwordHash)
	if err != nil {
		log.Printf("database: password get error %s\n", err)
		return false
	}

	return bcrypt.CompareHashAndPassword(passwordHash, []byte(password)) == nil

}
