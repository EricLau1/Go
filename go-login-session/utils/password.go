package utils

import (
	"golang.org/x/crypto/bcrypt"	
)

func Hash(password string) ([]byte, error) {

	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

}

func VerifyHash(hash, password []byte) (bool, error) {

	err := bcrypt.CompareHashAndPassword(hash, password)

	if err != nil {

		return false, err

	}

	return true, nil

}