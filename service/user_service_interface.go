package service

import (
	"context"
	"github.com/SawitProRecruitment/UserService/domain"
	error_handler "github.com/SawitProRecruitment/UserService/error"
	"github.com/golang-jwt/jwt/v4"
)

type ServiceInterface interface {
	ValidateUserRequest(input *domain.CreateUserRequest) (res []domain.CustomError, err error)
	ValidateToken(jwtToken string) (claims jwt.Claims, err *error_handler.CustomError)
	Register(ctx context.Context, input *domain.CreateUserRequest) (response *domain.UserResponse, err error)
	Login(ctx context.Context, input *domain.LoginRequest) (response *domain.LoginResponse, err *error_handler.CustomError)
	GetUser(ctx context.Context, claims jwt.Claims) (response *domain.UserResponse, err *error_handler.CustomError)
	Update(ctx context.Context, input *domain.UpdateUserRequest, claims jwt.Claims) (response *domain.UserResponse, err *error_handler.CustomError)
}
