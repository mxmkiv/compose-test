package models

import (
	"compose-test/app/auth"
	"fmt"
)

type User struct {
	Login        string
	PasswordHash []byte
}

func (u User) Whoami() {
	fmt.Printf("now u login as %s, ur hash %s\n", u.Login, u.PasswordHash)
}

func NewUser(login, password string) *User {

	return &User{
		Login:        login,
		PasswordHash: auth.GetHash(password),
	}

}
