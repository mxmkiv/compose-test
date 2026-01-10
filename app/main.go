package main

import (
	"bufio"
	"compose-test/app/handler"
	"compose-test/app/models"
	"compose-test/app/ui"
	"context"
	"os"
)

type App struct {
	scanner     *bufio.Scanner
	currentUser *models.User
}

func NewApp() *App {

	return &App{
		scanner: bufio.NewScanner(os.Stdin),
		// current user nil
	}

}

func (a *App) Run() {

	conn := models.DBInit()
	defer conn.Close(context.Background())

Menu:
	for {

		choose := ui.ShowMenu(a.scanner)
		switch choose {
		case "l":
			ui.ViewClear()
			//handler.UserLogin(conn)
		case "r":
			ui.ViewClear()
			handler.UserRegister(a.scanner, conn)
		case "q":
			break Menu
		default:
			ui.ViewClear()
			continue Menu
		}
	}
}

func main() {

	mainApp := NewApp()
	mainApp.Run()

}
