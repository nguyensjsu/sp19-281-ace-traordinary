package utils

import (
	"golang.org/x/crypto/bcrypt"
)

//Tutorial https://hackernoon.com/how-to-store-passwords-example-in-go-62712b1d2212
type Hash struct{}

//EncodePassword encodes the given password before inserting to DB
func (c *Hash) EncodePassword(password string) (string, error) {
	saltedBytes := []byte(password)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	hash := string(hashedBytes[:])
	return hash, nil
}

//DecodePassword Decodes the password for validation
func DecodePassword(password string) string {
	decpassword := ""
	return decpassword
}
