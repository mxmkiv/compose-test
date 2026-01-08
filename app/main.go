package main

import (
	"bufio"
	"compose-test/app/handler"
	"compose-test/app/models"
	"compose-test/app/ui"
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

	choose := ui.ShowMenu(a.scanner)
	switch choose {
	case "l":
		ui.ViewClear()
		handler.Login()
	case "r":
		ui.ViewClear()
		handler.Register()
	case "q":
		return
	}

}

func main() {

	mainApp := NewApp()
	mainApp.Run()

}
