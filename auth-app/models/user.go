package models

import (
	"fmt"
)

type Role string

const (
	RoleAdmin = "admin"
	RoleUser  = "user"
)

type User struct {
	Login     string
	Role      Role
	isBlocked bool
}

func (u User) Whoami() {
	fmt.Printf("| now u login as %s|\n", u.Login)
}

func NewUser(login string) *User {

	return &User{
		Login:     login,
		Role:      RoleUser,
		isBlocked: false,
	}

}
