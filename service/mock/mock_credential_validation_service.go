package mock

import (
	"crypto/rsa"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/mock"
)

type MockCredentialValidationService struct {
	mock.Mock
}

func NewMockCredentialValidationService() *MockCredentialValidationService {
	return &MockCredentialValidationService{}
}

func (m *MockCredentialValidationService) GenerateHashPassword(password string) (string, string, error) {
	args := m.Called(password)
	if hashPassword, ok := args.Get(0).(string); ok {
		if salt, ok := args.Get(1).(string); ok {
			if args.Get(2) != nil {
				return hashPassword, salt, args.Get(2).(error)
			}
			return hashPassword, salt, nil
		}
	}
	return "", "", nil
}

func (m *MockCredentialValidationService) Hash(password, salt string) string {
	args := m.Called(password, salt)
	if hashPassword, ok := args.Get(0).(string); ok {
		return hashPassword
	}
	return ""
}

func (m *MockCredentialValidationService) GeneratePrivateKey(privateKey string) (*rsa.PrivateKey, error) {
	args := m.Called(privateKey)
	if privateKey, ok := args.Get(0).(*rsa.PrivateKey); ok {
		if args.Get(1) != nil {
			return privateKey, args.Get(1).(error)
		}
		return privateKey, nil
	}
	return nil, nil
}

func (m *MockCredentialValidationService) GeneratePublicKey(publicKey string) (*rsa.PublicKey, error) {
	args := m.Called(publicKey)
	if publicKey, ok := args.Get(0).(*rsa.PublicKey); ok {
		if args.Get(1) != nil {
			return publicKey, args.Get(1).(error)
		}
		return publicKey, nil
	}
	return nil, nil
}

func (m *MockCredentialValidationService) GenerateJWTToken(id int32, phoneNumber string, privateKey *rsa.PrivateKey) (string, error) {
	args := m.Called(id, phoneNumber, privateKey)
	if token, ok := args.Get(0).(string); ok {
		if args.Get(1) != nil {
			return token, args.Get(1).(error)
		}
		return token, nil
	}
	return "", nil
}

func (m *MockCredentialValidationService) ValidateJWTToken(token string, publicKey *rsa.PublicKey) (jwt.Claims, error) {
	args := m.Called(token, publicKey)
	if claims, ok := args.Get(0).(jwt.Claims); ok {
		if args.Get(1) != nil {
			return claims, args.Get(1).(error)
		}
		return claims, nil
	}
	return nil, nil
}
