package utils

import (
	"math/rand"

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

//GenerateVerificationTocken for Registration
func GenerateVerificationTocken() string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, 15)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

//GenerateTemporaryPassword if Password is forgot
func GenerateTemporaryPassword() string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, 8)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
