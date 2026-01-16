package main

import (
	"compose-test/auth-app/app"
	"compose-test/auth-app/handler"
	"compose-test/auth-app/models"
	"compose-test/auth-app/ui"
	"context"

	"github.com/jackc/pgx/v5"
)

func main() {

	mainApp := app.NewApp()
	conn := models.DBInit()
	defer conn.Close(context.Background())

	MainLoop(mainApp, conn)

}

func MainLoop(mainApp *app.App, conn *pgx.Conn) {

	mainApp.AppState = app.StateMenu

	for mainApp.AppState != app.StateExit {
		switch mainApp.AppState {
		case app.StateMenu:

			choose := ui.ShowMenu(mainApp.Scanner)
			switch choose {
			case "l":
				ui.ViewClear()
				if handler.UserLogin(mainApp, conn) {
					mainApp.AppState = app.StatePanel
				}
			case "r":
				ui.ViewClear()
				if handler.UserRegister(mainApp, conn) {
					mainApp.AppState = app.StatePanel
				}
			case "q":
				mainApp.AppState = app.StateExit
			}

		case app.StatePanel:
			ui.PanelLoop(mainApp)
		}
	}
}
