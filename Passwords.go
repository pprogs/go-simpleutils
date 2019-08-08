package simpleutils

import (
	"golang.org/x/crypto/bcrypt"
)

//GeneratePasswordHash generates hash for a given password using bcrypt with default cost
func GeneratePasswordHash(pass []byte) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return hash, err
}

//GeneratePasswordHashWithCost generates hash for a given password using bcrypt with specified cost
func GeneratePasswordHashWithCost(pass []byte, cost int) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword(pass, cost)
	if err != nil {
		panic(err)
	}
	return hash, err
}

//CheckPasswordsMatch check if plainPass match provided hash using bcrypt
func CheckPasswordsMatch(plainPass string, hash []byte) bool {
	pass := []byte(plainPass)
	err := bcrypt.CompareHashAndPassword(hash, pass)
	return err == nil
}
