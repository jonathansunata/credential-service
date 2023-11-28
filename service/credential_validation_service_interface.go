package service

import (
	"crypto/rsa"
	"github.com/golang-jwt/jwt/v4"
)

type CredentialValidationServiceInterface interface {
	GenerateHashPassword(password string) (string, string, error)
	Hash(password, salt string) string
	GeneratePrivateKey(string) (*rsa.PrivateKey, error)
	GeneratePublicKey(string) (*rsa.PublicKey, error)
	GenerateJWTToken(int32, string, *rsa.PrivateKey) (string, error)
	ValidateJWTToken(string, *rsa.PublicKey) (jwt.Claims, error)
}
