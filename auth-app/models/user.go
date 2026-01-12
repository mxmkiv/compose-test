package models

import (
	"compose-test/auth-app/auth"
	"fmt"
)

type User struct {
	Login        string
	PasswordHash []byte
}

func (u User) Whoami() {
	fmt.Printf("now u login as %s\n", u.Login)
}

func NewUser(login, password string) *User {

	return &User{
		Login:        login,
		PasswordHash: auth.GetHash(password),
	}

}
