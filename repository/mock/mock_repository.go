package mock_repository

import (
	"context"
	"github.com/SawitProRecruitment/UserService/domain/entity"
	error_handler "github.com/SawitProRecruitment/UserService/error"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func NewMockRepository() *MockRepository {
	return &MockRepository{}
}

// Update mocks the Update method of RepositoryInterface
func (m *MockRepository) Update(ctx context.Context, input *entity.User) error {
	args := m.Called(ctx, input)
	if args.Get(0) != nil {
		return args.Get(0).(error)
	}
	return nil
}

// Register mocks the Register method of RepositoryInterface
func (m *MockRepository) Register(ctx context.Context, input *entity.User) (int32, error) {
	args := m.Called(ctx, input)
	// Check if the first return value is non-nil and of type int32
	if id, ok := args.Get(0).(int32); ok {
		// Check if the second return value is non-nil
		if args.Get(1) != nil {
			return id, args.Get(1).(*error_handler.CustomError)
		}
		return id, nil
	}
	return 0, nil
}

// GetById mocks the GetById method of RepositoryInterface
func (m *MockRepository) GetById(ctx context.Context, id int32) (*entity.User, error) {
	args := m.Called(ctx, id)
	// Check if the first return value is non-nil and of type int32
	if id, ok := args.Get(0).(*entity.User); ok {
		// Check if the second return value is non-nil
		if args.Get(1) != nil {
			return id, args.Get(1).(*error_handler.CustomError)
		}
		return id, nil
	}
	return nil, nil
}

// GetByPhone mocks the GetByPhone method of RepositoryInterface
func (m *MockRepository) GetByPhone(ctx context.Context, phone string) (*entity.User, error) {
	args := m.Called(ctx, phone)
	// Check if the first return value is non-nil and of type int32
	if id, ok := args.Get(0).(*entity.User); ok {
		// Check if the second return value is non-nil
		if args.Get(1) != nil {
			return id, args.Get(1).(*error_handler.CustomError)
		}
		return id, nil
	}
	return nil, nil
}

// GetByPhoneOther mocks the GetByPhoneOther method of RepositoryInterface
func (m *MockRepository) GetByPhoneOther(ctx context.Context, phone string, id int32) (*entity.User, error) {
	args := m.Called(ctx, phone, id)
	// Check if the first return value is non-nil and of type int32
	if id, ok := args.Get(0).(*entity.User); ok {
		// Check if the second return value is non-nil
		if args.Get(1) != nil {
			return id, args.Get(1).(*error_handler.CustomError)
		}
		return id, nil
	}
	return nil, nil
}
