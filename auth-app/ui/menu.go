package ui

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func ShowMenu(scanner *bufio.Scanner) string {

	fmt.Println("l - login | r - register | q - quit")
	choose := GetInput(scanner)
	return choose
}

func GetInput(scanner *bufio.Scanner) string {

	/*

		добавить удаление лишних символов (пробелов)

	*/

	scanner.Scan()
	choose := scanner.Text()
	return choose

}

func ViewClear() {

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()

}
