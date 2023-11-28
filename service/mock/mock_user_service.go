package mock

import (
	"context"
	"github.com/SawitProRecruitment/UserService/domain"
	"github.com/SawitProRecruitment/UserService/domain/entity"
	error_handler "github.com/SawitProRecruitment/UserService/error"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) Update(ctx context.Context, input *domain.UpdateUserRequest, claims jwt.Claims) (response *domain.UserResponse, err *error_handler.CustomError) {
	args := m.Called(ctx, input, claims)
	if response, ok := args.Get(0).(*domain.UserResponse); ok {
		if args.Get(1) != nil {
			return response, args.Get(1).(*error_handler.CustomError)
		}
		return response, nil
	}
	return nil, nil
}

func (m *MockUserService) Register(ctx context.Context, input *domain.CreateUserRequest) (response *domain.UserResponse, err error) {
	args := m.Called(ctx, input)
	if id, ok := args.Get(0).(*domain.UserResponse); ok {
		if args.Get(1) != nil {
			return id, args.Get(1).(error)
		}
		return id, nil
	}
	return nil, nil
}

func (m *MockUserService) GetById(ctx context.Context, id int32) (*entity.User, error) {
	args := m.Called(ctx, id)
	if id, ok := args.Get(0).(*entity.User); ok {
		if args.Get(1) != nil {
			return id, args.Get(1).(error)
		}
		return id, nil
	}
	return nil, nil
}

func (m *MockUserService) GetByPhone(ctx context.Context, phone string) (*entity.User, error) {
	args := m.Called(ctx, phone)
	if id, ok := args.Get(0).(*entity.User); ok {
		if args.Get(1) != nil {
			return id, args.Get(1).(error)
		}
		return id, nil
	}
	return nil, nil
}

func (m *MockUserService) GetByPhoneOther(ctx context.Context, phone string, id int32) (*entity.User, error) {
	args := m.Called(ctx, phone, id)
	if id, ok := args.Get(0).(*entity.User); ok {
		if args.Get(1) != nil {
			return id, args.Get(1).(error)
		}
		return id, nil
	}
	return nil, nil
}

func (m *MockUserService) ValidateUserRequest(input *domain.CreateUserRequest) (res []domain.CustomError, err error) {
	args := m.Called(input)
	if res, ok := args.Get(0).([]domain.CustomError); ok {
		if args.Get(1) != nil {
			return res, args.Get(1).(error)
		}
		return res, nil
	}
	return nil, nil
}

func (m *MockUserService) ValidateToken(jwtToken string) (claims jwt.Claims, err *error_handler.CustomError) {
	args := m.Called(jwtToken)
	if claims, ok := args.Get(0).(jwt.Claims); ok {
		if args.Get(1) != nil {
			return claims, args.Get(1).(*error_handler.CustomError)
		}
		return claims, nil
	}
	return nil, nil
}

func (m *MockUserService) Login(ctx context.Context, input *domain.LoginRequest) (response *domain.LoginResponse, err *error_handler.CustomError) {
	args := m.Called(ctx, input)
	if response, ok := args.Get(0).(*domain.LoginResponse); ok {
		if args.Get(1) != nil {
			return response, args.Get(1).(*error_handler.CustomError)
		}
		return response, nil
	}
	return nil, nil
}

func (m *MockUserService) GetUser(ctx context.Context, claims jwt.Claims) (response *domain.UserResponse, err *error_handler.CustomError) {
	args := m.Called(ctx, claims)
	if response, ok := args.Get(0).(*domain.UserResponse); ok {
		if args.Get(1) != nil {
			return response, args.Get(1).(*error_handler.CustomError)
		}
		return response, nil
	}
	return nil, nil
}

func NewMockUserService() *MockUserService {
	return &MockUserService{}
}

