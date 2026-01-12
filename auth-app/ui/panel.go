package ui

import "compose-test/auth-app/app"

func PanelInit(a *app.App) {

	ViewClear()
	a.CurrentUser.Whoami()

}
