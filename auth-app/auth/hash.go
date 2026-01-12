package auth

import "golang.org/x/crypto/bcrypt"

func GetHash(password string) []byte {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return hash

}
