package main

import (
	"compose-test/auth-app/app"
	"compose-test/auth-app/handler"
	"compose-test/auth-app/models"
	"compose-test/auth-app/ui"
	"context"
)

func main() {

	mainApp := app.NewApp()
	conn := models.DBInit()
	defer conn.Close(context.Background())

Menu:
	for {

		choose := ui.ShowMenu(mainApp.Scanner)
		switch choose {
		case "l":
			ui.ViewClear()
			//handler.UserLogin(conn)
		case "r":
			ui.ViewClear()
			handler.UserRegister(mainApp, conn)
		case "q":
			break Menu
		default:
			ui.ViewClear()
			continue Menu
		}
	}

}
