package util

import (
	"crypto/rsa"
	"fmt"
	error_handler "github.com/SawitProRecruitment/UserService/error"
	jwt "github.com/golang-jwt/jwt/v4"
	"time"
)

type Claims struct {
	UserId      int32  `json:"user_id"`
	PhoneNumber string `json:"phone_number"`
	jwt.RegisteredClaims
}

func GenerateJWTToken(userId int32, phoneNumber string, privateKey *rsa.PrivateKey) (string, error) {
	claims := Claims{
		UserId:      userId,
		PhoneNumber: phoneNumber,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	result, err := token.SignedString(privateKey)
	if err != nil {
		return "", error_handler.NewCustomError(500, err.Error())
	}
	return result, nil
}

func ValidateJWTToken(tokenString string, publicKey *rsa.PublicKey) (jwt.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})

	if err != nil {
		return nil, error_handler.NewCustomError(500, err.Error())
	}

	if !token.Valid {
		return nil, error_handler.NewCustomError(403, fmt.Sprintf("invalid token"))
	}

	return token.Claims, nil
}
