package util

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
)

func Hash(password, salt string) string {
	hash := sha256.New()
	hash.Write([]byte(password + salt))
	hashedPassword := hash.Sum(nil)
	return base64.StdEncoding.EncodeToString(hashedPassword)
}

func GenerateSalt(length int) (string, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	return base64.StdEncoding.EncodeToString(salt), err
}
