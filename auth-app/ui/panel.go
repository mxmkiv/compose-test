package ui

import (
	"compose-test/auth-app/app"
	"fmt"
)

func PanelInit(a *app.App) {

	ViewClear()
	a.CurrentUser.Whoami()
	fmt.Println("q - logout | ----")

}

func PanelLoop(mainApp *app.App) {

	for {

		PanelInit(mainApp)
		choose := GetInput(mainApp.Scanner)
		switch choose {
		case "q":
			ViewClear()
			mainApp.AppState = app.StateMenu
			mainApp.CurrentUser = nil
			return
		default:
			continue
		}

	}
}
