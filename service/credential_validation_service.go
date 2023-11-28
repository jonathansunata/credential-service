package service

import (
	"crypto/rsa"
	"fmt"
	"github.com/SawitProRecruitment/UserService/util"
	"github.com/golang-jwt/jwt/v4"
)

type CredentialValidationService struct {
}

func NewCredentialValidationService() *CredentialValidationService {
	return &CredentialValidationService{}
}

func (e *CredentialValidationService) GenerateHashPassword(password string) (string, string, error) {
	salt, err := util.GenerateSalt(16)
	if err != nil {
		fmt.Println("Error generating salt:", err)
		return "", "", err
	}
	hashPassword := util.Hash(password, salt)
	if err != nil {
		fmt.Println("Error generating hash:", err)
		return "", "", err
	}

	return hashPassword, salt, nil
}

func (e *CredentialValidationService) Hash(password, salt string) string {
	return util.Hash(password, salt)
}

func (e *CredentialValidationService) GeneratePrivateKey(privateKey string) (*rsa.PrivateKey, error) {
	return util.GeneratePrivateKey(privateKey)
}

func (e *CredentialValidationService) GeneratePublicKey(publicKey string) (*rsa.PublicKey, error) {
	return util.GeneratePublicKey(publicKey)
}

func (e *CredentialValidationService) GenerateJWTToken(id int32, phoneNumber string, privateKey *rsa.PrivateKey) (string, error) {
	return util.GenerateJWTToken(id, phoneNumber, privateKey)
}

func (e *CredentialValidationService) ValidateJWTToken(token string, publicKey *rsa.PublicKey) (jwt.Claims, error) {
	return util.ValidateJWTToken(token, publicKey)
}
