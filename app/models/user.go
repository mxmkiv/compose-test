package models

import (
	"compose-test/app/auth"
	"fmt"
)

type User struct {
	login        string
	passwordHash []byte
}

func (u User) Whoami() {
	fmt.Printf("now u login as %s, ur hash %s\n", u.login, u.passwordHash)
}

func NewUser(login, password string) *User {

	return &User{
		login:        login,
		passwordHash: auth.GetHash(password),
	}

}
