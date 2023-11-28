package service

import (
	"context"
	"crypto/rsa"
	"errors"
	"github.com/SawitProRecruitment/UserService/domain"
	"github.com/SawitProRecruitment/UserService/domain/entity"
	mock_repository "github.com/SawitProRecruitment/UserService/repository/mock"
	mock2 "github.com/SawitProRecruitment/UserService/service/mock"
	"github.com/SawitProRecruitment/UserService/util"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

// create unit test for Register function
// Path: service/user_service_test.go
func TestUserService_Register(t *testing.T) {

	mockRepo := mock_repository.NewMockRepository()
	mockRepo.On("Register", mock.Anything, mock.Anything).Return(int32(123), nil)
	mockRepo.On("GetById", mock.Anything, mock.Anything).Return(&entity.User{}, nil)

	mockCredentialValidation := mock2.NewMockCredentialValidationService()
	mockCredentialValidation.On("GenerateHashPassword", mock.Anything).Return("Passw0rd!", nil)

	userService := NewUserService(mockRepo, mockCredentialValidation)

	t.Run("Valid input", func(t *testing.T) {
		// Create a valid input
		validInput := &domain.CreateUserRequest{
			PhoneNumber: "+62123456789",
			FullName:    "John Doe",
			Password:    "Passw0rd!",
		}

		// Validate the input
		_, err := userService.Register(context.Background(), validInput)

		// Assert that there are no errors and the error is nil
		assert.Nil(t, err)
	})
}

func TestUserService_Login(t *testing.T) {
	mockRepo := mock_repository.NewMockRepository()
	mockRepo.On("GetByPhone", mock.Anything, mock.Anything).Return(&entity.User{
		Password: "Passw0rd!",
	}, nil)
	mockRepo.On("Update", mock.Anything, mock.Anything).Return(nil)

	mockCredentialValidation := mock2.NewMockCredentialValidationService()
	mockCredentialValidation.On("Hash", mock.Anything, mock.Anything).Return("Passw0rd!")
	mockCredentialValidation.On("GeneratePrivateKey", mock.Anything).Return(rsa.PrivateKey{}, nil)
	mockCredentialValidation.On("GenerateJWTToken", mock.Anything, mock.Anything, mock.Anything).Return("token", nil)

	userService := NewUserService(mockRepo, mockCredentialValidation)

	t.Run("Valid input", func(t *testing.T) {
		// Create a valid input
		validInput := &domain.LoginRequest{
			PhoneNumber: "+62123456789",
			Password:    "Passw0rd!",
		}

		// Validate the input
		_, err := userService.Login(context.Background(), validInput)

		// Assert that there are no errors and the error is nil
		assert.Nil(t, err)
	})

}

func TestUserService_GetUser(t *testing.T) {

	mockRepo := mock_repository.NewMockRepository()
	mockRepo.On("GetById", mock.Anything, mock.Anything).Return(&entity.User{
		FullName:    "John Doe",
		PhoneNumber: "+62123456789",
	}, nil)

	userService := NewUserService(mockRepo, nil)

	t.Run("Valid input", func(t *testing.T) {
		// Validate the input
		_, err := userService.GetUser(context.Background(), jwt.Claims(&util.Claims{
			UserId: 123,
		}))

		// Assert that there are no errors and the error is nil
		assert.Nil(t, err)
	})
}

func TestUserService_Update(t *testing.T) {

	mockRepo := mock_repository.NewMockRepository()
	mockRepo.On("GetById", mock.Anything, mock.Anything).Return(&entity.User{
		FullName:    "John Doe",
		PhoneNumber: "+62123456789",
	}, nil)
	mockRepo.On("GetById", mock.Anything, mock.Anything).Return(&entity.User{})
	mockRepo.On("GetByPhoneOther", mock.Anything, mock.Anything, mock.Anything).Return(&entity.User{
		Id: 0,
	}, nil)
	mockRepo.On("Update", mock.Anything, mock.Anything).Return(nil)

	userService := NewUserService(mockRepo, nil)

	t.Run("Valid input", func(t *testing.T) {
		// Create a valid input
		validInput := &domain.UpdateUserRequest{
			FullName:    "John Doe",
			PhoneNumber: "+62123456789",
		}

		// Validate the input
		_, err := userService.Update(context.Background(), validInput, jwt.Claims(&util.Claims{
			UserId: 123,
		}))

		// Assert that there are no errors and the error is nil
		assert.Nil(t, err)
	})
}

func TestUserService_ValidateToken(t *testing.T) {

	t.Run("Valid input", func(t *testing.T) {

		mockRepo := mock_repository.NewMockRepository()
		mockCredentialValidation := mock2.NewMockCredentialValidationService()
		mockCredentialValidation.On("GeneratePublicKey", mock.Anything).Return(rsa.PublicKey{}, nil)
		mockCredentialValidation.On("ValidateJWTToken", mock.Anything, mock.Anything).Return(jwt.Claims(&util.Claims{
			UserId: 123,
		}), nil)

		userService := NewUserService(mockRepo, mockCredentialValidation)
		// Validate the input
		_, err := userService.ValidateToken("token")

		// Assert that there are no errors and the error is nil
		assert.Nil(t, err)
	})

	t.Run("Invalid input", func(t *testing.T) {
		mockRepo := mock_repository.NewMockRepository()
		mockCredentialValidation := mock2.NewMockCredentialValidationService()
		mockCredentialValidation.On("GeneratePublicKey", mock.Anything).Return(rsa.PublicKey{}, nil)
		mockCredentialValidation.On("ValidateJWTToken", mock.Anything, mock.Anything).Return(jwt.Claims(&util.Claims{
			UserId: 123,
		}), nil)

		userService := NewUserService(mockRepo, mockCredentialValidation)

		// Validate the input
		_, err := userService.ValidateToken("")

		// Assert that there are no errors and the error is nil
		assert.NotNil(t, err)
	})

	t.Run("Invalid input", func(t *testing.T) {

		mockRepo := mock_repository.NewMockRepository()
		mockCredentialValidation := mock2.NewMockCredentialValidationService()
		mockCredentialValidation.On("GeneratePublicKey", mock.Anything).Return(rsa.PublicKey{}, nil)
		mockCredentialValidation.On("ValidateJWTToken", mock.Anything, mock.Anything).Return(jwt.Claims(&util.Claims{
			UserId: 123,
		}), errors.New("Invalid credential"))

		userService := NewUserService(mockRepo, mockCredentialValidation)

		// Validate the input
		_, err := userService.ValidateToken("jwtToken")

		// Assert that there are no errors and the error is nil
		assert.NotNil(t, err)
	})
}

func TestUserService_ValidateUserRequest(t *testing.T) {

	mockRepo := mock_repository.NewMockRepository()
	mockCrendentialValidation := mock2.NewMockCredentialValidationService()
	userService := NewUserService(mockRepo, mockCrendentialValidation)

	t.Run("Valid input", func(t *testing.T) {
		// Create a valid input
		validInput := &domain.CreateUserRequest{
			PhoneNumber: "+62123456789",
			FullName:    "John Doe",
			Password:    "Passw0rd!",
		}

		// Mock any UserService method calls if needed

		// Validate the input
		errors, err := userService.ValidateUserRequest(validInput)

		// Assert that there are no errors and the error is nil
		assert.NoError(t, err)
		assert.Empty(t, errors)
	})

	t.Run("Invalid phone number", func(t *testing.T) {
		// Create an input with an invalid phone number
		invalidPhoneNumberInput := &domain.CreateUserRequest{
			PhoneNumber: "123",
			FullName:    "John Doe",
			Password:    "Passw0rd!",
		}

		// Mock any UserService method calls if needed

		// Validate the input
		errors, err := userService.ValidateUserRequest(invalidPhoneNumberInput)

		// Assert that there is an error and the error message is as expected
		assert.Error(t, err)
		assert.Contains(t, errors, domain.CustomError{
			Field:      "phone_number",
			ErrMessage: "phone number must be between 10 and 13 characters",
		})
		assert.Contains(t, errors, domain.CustomError{
			Field:      "phone_number",
			ErrMessage: "phone number must start with +62",
		})
	})

	t.Run("Invalid Password", func(t *testing.T) {
		// Create an input with an invalid phone number
		invalidPasswordInput := &domain.CreateUserRequest{
			PhoneNumber: "123",
			FullName:    "John Doe",
			Password:    "Password",
		}

		// Mock any UserService method calls if needed

		// Validate the input
		errors, err := userService.ValidateUserRequest(invalidPasswordInput)

		// Assert that there is an error and the error message is as expected
		assert.Error(t, err)
		assert.Contains(t, errors, domain.CustomError{
			Field:      "password",
			ErrMessage: "password must contain at least one number and one uppercase letter and one special character",
		})
	})

	t.Run("Invalid Length Password", func(t *testing.T) {
		// Create an input with an invalid phone number
		invalidPasswordInput := &domain.CreateUserRequest{
			PhoneNumber: "123",
			FullName:    "John Doe",
			Password:    "123",
		}

		// Mock any UserService method calls if needed

		// Validate the input
		errors, err := userService.ValidateUserRequest(invalidPasswordInput)

		// Assert that there is an error and the error message is as expected
		assert.Error(t, err)
		assert.Contains(t, errors, domain.CustomError{
			Field:      "password",
			ErrMessage: "password must be between 6 and 64 characters",
		})
	})

}
