package app

import (
	"bufio"
	"compose-test/auth-app/models"
	"os"
)

type App struct {
	Scanner     *bufio.Scanner
	CurrentUser *models.User
}

func NewApp() *App {

	return &App{
		Scanner: bufio.NewScanner(os.Stdin),
		// current user nil
	}
}
