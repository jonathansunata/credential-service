package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/SawitProRecruitment/UserService/domain"
	"github.com/SawitProRecruitment/UserService/domain/entity"
	error_handler "github.com/SawitProRecruitment/UserService/error"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/SawitProRecruitment/UserService/util"
	"github.com/golang-jwt/jwt/v4"
)

type UserService struct {
	repository           repository.RepositoryInterface
	credentialValidation CredentialValidationServiceInterface
}

func NewUserService(r repository.RepositoryInterface, c CredentialValidationServiceInterface) *UserService {
	return &UserService{
		repository:           r,
		credentialValidation: c,
	}
}

func (s *UserService) ValidateUserRequest(input *domain.CreateUserRequest) (res []domain.CustomError, err error) {
	errorField := make([]domain.CustomError, 0)

	if input.PhoneNumber != "" {
		if len(input.PhoneNumber) < 10 || len(input.PhoneNumber) > 13 {
			errorField = append(errorField, domain.CustomError{
				Field:      "phone_number",
				ErrMessage: "phone number must be between 10 and 13 characters",
			})
		}

		if !util.StartsWithCountryCode(input.PhoneNumber, "+62") {
			errorField = append(errorField, domain.CustomError{
				Field:      "phone_number",
				ErrMessage: "phone number must start with +62",
			})
		}
	}

	if input.FullName != "" {
		if len(input.FullName) < 3 || len(input.FullName) > 60 {
			errorField = append(errorField, domain.CustomError{
				Field:      "full_name",
				ErrMessage: "full name must be between 3 and 60 characters",
			})
		}
	}

	if input.Password != "" {
		if len(input.Password) < 6 || len(input.Password) > 64 {
			errorField = append(errorField, domain.CustomError{
				Field:      "password",
				ErrMessage: "password must be between 6 and 64 characters",
			})
		}

		if !util.HasCharacters(input.Password) {
			errorField = append(errorField, domain.CustomError{
				Field:      "password",
				ErrMessage: "password must contain at least one number and one uppercase letter and one special character",
			})
		}
	}

	if len(errorField) > 0 {
		err = errors.New("invalid input")
		return errorField, err
	}
	return
}

func (s *UserService) Register(ctx context.Context, input *domain.CreateUserRequest) (res *domain.UserResponse, err error) {
	password, salt, err := s.credentialValidation.GenerateHashPassword(input.Password)

	fmt.Println("password", password)
	fmt.Println("salt", salt)
	user := &entity.User{
		FullName:    input.FullName,
		PhoneNumber: input.PhoneNumber,
		Password:    password,
		Salt:        salt,
	}

	id, err := s.repository.Register(ctx, user)
	if err != nil {
		return
	}

	userRes, err := s.repository.GetById(ctx, id)
	if err != nil {
		return
	}
	//convert User to UserResponse
	res = &domain.UserResponse{
		FullName:    userRes.FullName,
		PhoneNumber: userRes.PhoneNumber,
	}

	return
}

func (s *UserService) Login(ctx context.Context, input *domain.LoginRequest) (*domain.LoginResponse, *error_handler.CustomError) {
	fmt.Printf("Login with password %v \n", input.Password)
	user, err := s.repository.GetByPhone(ctx, input.PhoneNumber)
	if err != nil {
		fmt.Errorf("error when get user by phone : %v, err: %v \n", input.PhoneNumber, err)
		return nil, error_handler.NewCustomError(400, err.Error())
	}

	fmt.Printf("success get user id : %v \n", user.Id)

	password := s.credentialValidation.Hash(input.Password, user.Salt)
	if password != user.Password {
		fmt.Errorf("Error when hash the password with salt : %v, err: %v", user.Salt, err)
		return nil, error_handler.NewCustomError(400, err.Error())
	}

	fmt.Printf("success validate the password : %v \n", user.Id)

	privateKey, err := s.credentialValidation.GeneratePrivateKey("/app/key/private_key.pem")
	if err != nil {
		fmt.Errorf("error when generate private key : %v, err: %v", user.FullName, err)
		return nil, error_handler.NewCustomError(500, err.Error())
	}

	fmt.Printf("success genertae private key for user id : %v \n", user.Id)

	token, err := s.credentialValidation.GenerateJWTToken(user.Id, user.PhoneNumber, privateKey)
	if err != nil {
		fmt.Errorf("error when Generate JWT Token : %v, err: %v", user.FullName, err)
		return nil, error_handler.NewCustomError(500, err.Error())
	}

	fmt.Printf("success genertae private key for user id : %v \n", user.Id)

	if user.CountLogin == nil {
		user.CountLogin = new(int32)
		*user.CountLogin = 0
	}

	*user.CountLogin = *user.CountLogin + 1

	err = s.repository.Update(ctx, user)
	if err != nil {
		fmt.Errorf("error when update user : %v, err: %v \n", user.FullName, err)
		return nil, error_handler.NewCustomError(500, err.Error())
	}

	fmt.Printf("success update for user id : %v \n", user.Id)

	res := &domain.LoginResponse{
		AccessToken: token,
		Id:          user.Id,
	}

	return res, nil
}

func (s *UserService) GetUser(ctx context.Context, claims jwt.Claims) (*domain.UserResponse, *error_handler.CustomError) {
	userId := claims.(*util.Claims).UserId
	userData, err := s.repository.GetById(ctx, userId)
	if err != nil {
		return nil, error_handler.NewCustomError(500, err.Error())
	}

	res := &domain.UserResponse{
		FullName:    userData.FullName,
		PhoneNumber: userData.PhoneNumber,
	}

	return res, nil
}

func (s *UserService) ValidateToken(jwtToken string) (jwt.Claims, *error_handler.CustomError) {
	if jwtToken == "" {
		return nil, error_handler.NewCustomError(403, "token not found")
	}

	publicKey, err := s.credentialValidation.GeneratePublicKey("/app/key/public_key.pem")
	if err != nil {
		return nil, error_handler.NewCustomError(500, err.Error())
	}

	claims, err := s.credentialValidation.ValidateJWTToken(jwtToken, publicKey)
	if err != nil {
		return nil, error_handler.NewCustomError(403, err.Error())
	}

	return claims, nil
}

func (s *UserService) Update(ctx context.Context, input *domain.UpdateUserRequest, claims jwt.Claims) (*domain.UserResponse, *error_handler.CustomError) {

	userId := claims.(*util.Claims).UserId
	userData, err := s.repository.GetById(ctx, userId)
	if err != nil {
		return nil, error_handler.NewCustomError(500, err.Error())
	}

	userByPhone, err := s.repository.GetByPhoneOther(ctx, input.PhoneNumber, userId)
	if userByPhone != nil && userByPhone.Id != 0 {
		return nil, error_handler.NewCustomError(409, "phone number already exists")
	}

	userData.Merge(input)

	err = s.repository.Update(ctx, userData)
	if err != nil {
		return nil, error_handler.NewCustomError(500, err.Error())
	}

	res := &domain.UserResponse{
		FullName:    userData.FullName,
		PhoneNumber: userData.PhoneNumber,
	}

	return res, nil
}
