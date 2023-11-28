package handler

import (
	"bytes"
	"encoding/json"
	"github.com/SawitProRecruitment/UserService/domain"
	error_handler "github.com/SawitProRecruitment/UserService/error"
	mock_service "github.com/SawitProRecruitment/UserService/service/mock"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegister(t *testing.T) {
	testCases := []struct {
		Name           string
		HttpMethod     string
		RequestBody    map[string]interface{}
		ExpectedStatus int
		Mock           func() *mock_service.MockUserService
		ExpectedBody   map[string]interface{}
	}{
		{
			Name:           "Valid registration",
			HttpMethod:     http.MethodPost,
			RequestBody:    map[string]interface{}{"full_name": "some_value", "phone_number": "some_value"},
			ExpectedStatus: http.StatusCreated,
			Mock: func() *mock_service.MockUserService {
				mockService := mock_service.NewMockUserService()
				mockService.On("ValidateUserRequest", mock.Anything).Return([]domain.CustomError{}, nil)
				mockService.On("Register", mock.Anything, mock.Anything).Return(&domain.UserResponse{
					FullName:    "jonathan",
					PhoneNumber: "08123456789",
				}, nil)
				return mockService
			},
			ExpectedBody: map[string]interface{}{
				"full_name":    "jonathan",
				"phone_number": "08123456789",
			},
		},
		{
			Name:           "Invalid Registration",
			HttpMethod:     http.MethodPost,
			RequestBody:    map[string]interface{}{"full_name": "some_value", "phone_number": "some_value"},
			ExpectedStatus: http.StatusBadRequest,
			Mock: func() *mock_service.MockUserService {
				mockService := mock_service.NewMockUserService()
				mockService.On("ValidateUserRequest", mock.Anything).Return([]domain.CustomError{
					{
						Field:      "full_name",
						ErrMessage: "Less than 3 characters",
					},
				}, nil)
				return mockService
			},
			ExpectedBody: map[string]interface{}{
				"full_name":    "jonathan",
				"phone_number": "08123456789",
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {

			mockService := tc.Mock()

			server := Server{
				Service: mockService,
			}

			e := echo.New()

			requestBytes, _ := json.Marshal(tc.RequestBody)
			req := httptest.NewRequest(tc.HttpMethod, "/register", bytes.NewBuffer(requestBytes))
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()

			ctx := e.NewContext(req, rec)
			err := server.Register(ctx, domain.CreateUserRequest{})

			assert.NoError(t, err)
			assert.Equal(t, tc.ExpectedStatus, rec.Code)

			mockService.AssertExpectations(t)
		})
	}
}

func TestLogin(t *testing.T) {

	testCases := []struct {
		Name           string
		HttpMethod     string
		RequestBody    map[string]interface{}
		ExpectedStatus int
		Mock           func() *mock_service.MockUserService
		ExpectedBody   map[string]interface{}
	}{
		{
			Name:           "Valid Login",
			HttpMethod:     http.MethodPut,
			RequestBody:    map[string]interface{}{"phone_number": "some_value"},
			ExpectedStatus: http.StatusOK,
			Mock: func() *mock_service.MockUserService {
				mockService := mock_service.NewMockUserService()
				mockService.On("Login", mock.Anything, mock.Anything).Return(domain.LoginResponse{}, nil)
				return mockService
			},
			ExpectedBody: map[string]interface{}{
				"full_name":    "jonathan",
				"phone_number": "08123456789",
			},
		},
		{
			Name:           "Invalid Login",
			HttpMethod:     http.MethodPut,
			RequestBody:    map[string]interface{}{"phone_number": "some_value"},
			ExpectedStatus: http.StatusBadRequest,
			Mock: func() *mock_service.MockUserService {
				mockService := mock_service.NewMockUserService()
				mockService.On("Login", mock.Anything, mock.Anything).Return(&domain.LoginResponse{}, &error_handler.CustomError{
					Code:    400,
					Message: "Invalid Password",
				})
				return mockService
			},
			ExpectedBody: map[string]interface{}{
				"full_name":    "jonathan",
				"phone_number": "08123456789",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {

			mockService := tc.Mock()

			server := Server{
				Service: mockService,
			}

			e := echo.New()

			requestBytes, _ := json.Marshal(tc.RequestBody)
			req := httptest.NewRequest(tc.HttpMethod, "/login", bytes.NewBuffer(requestBytes))
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()

			ctx := e.NewContext(req, rec)
			server.Login(ctx, domain.LoginRequest{})
			assert.Equal(t, tc.ExpectedStatus, rec.Code)
		})
	}
}

func TestGetUser(t *testing.T) {
	testCases := []struct {
		Name           string
		HttpMethod     string
		RequestBody    map[string]interface{}
		ExpectedStatus int
		Mock           func() *mock_service.MockUserService
		ExpectedBody   map[string]interface{}
	}{
		{
			Name:           "Valid Get User",
			HttpMethod:     http.MethodPut,
			RequestBody:    map[string]interface{}{"phone_number": "some_value"},
			ExpectedStatus: http.StatusOK,
			Mock: func() *mock_service.MockUserService {
				mockService := mock_service.NewMockUserService()
				mockService.On("ValidateToken", mock.Anything).Return(jwt.RegisteredClaims{}, nil)
				mockService.On("GetUser", mock.Anything, mock.Anything).Return(domain.UserResponse{}, nil)
				return mockService
			},
			ExpectedBody: map[string]interface{}{
				"full_name":    "jonathan",
				"phone_number": "08123456789",
			},
		},
		{
			Name:           "Get User 403 Error",
			HttpMethod:     http.MethodGet,
			RequestBody:    map[string]interface{}{"phone_number": "some_value"},
			ExpectedStatus: http.StatusForbidden,
			Mock: func() *mock_service.MockUserService {
				mockService := mock_service.NewMockUserService()
				mockService.On("ValidateToken", mock.Anything).Return(jwt.RegisteredClaims{}, &error_handler.CustomError{
					Code:    403,
					Message: "Forbidden",
				})
				mockService.On("GetUser", mock.Anything, mock.Anything).Return(domain.UserResponse{}, nil)
				return mockService
			},
			ExpectedBody: map[string]interface{}{
				"full_name":    "jonathan",
				"phone_number": "08123456789",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {

			mockService := tc.Mock()

			server := Server{
				Service: mockService,
			}

			e := echo.New()

			requestBytes, _ := json.Marshal(tc.RequestBody)
			req := httptest.NewRequest(tc.HttpMethod, "/user", bytes.NewBuffer(requestBytes))
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()

			ctx := e.NewContext(req, rec)
			server.GetUser(ctx)

			assert.Equal(t, tc.ExpectedStatus, rec.Code)
		})
	}
}

func TestUpdate(t *testing.T) {
	testCases := []struct {
		Name           string
		HttpMethod     string
		RequestBody    map[string]interface{}
		ExpectedStatus int
		Mock           func() *mock_service.MockUserService
		ExpectedBody   map[string]interface{}
	}{
		{
			Name:           "Valid Update User",
			HttpMethod:     http.MethodPatch,
			RequestBody:    map[string]interface{}{"phone_number": "some_value"},
			ExpectedStatus: http.StatusOK,
			Mock: func() *mock_service.MockUserService {
				mockService := mock_service.NewMockUserService()
				mockService.On("ValidateToken", mock.Anything).Return(jwt.RegisteredClaims{}, nil)
				mockService.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(nil)
				return mockService
			},
			ExpectedBody: map[string]interface{}{
				"full_name":    "jonathan",
				"phone_number": "08123456789",
			},
		},
		{
			Name:           "403 Error",
			HttpMethod:     http.MethodPatch,
			RequestBody:    map[string]interface{}{"phone_number": "some_value"},
			ExpectedStatus: http.StatusForbidden,
			Mock: func() *mock_service.MockUserService {
				mockService := mock_service.NewMockUserService()
				mockService.On("ValidateToken", mock.Anything).Return(jwt.RegisteredClaims{}, &error_handler.CustomError{
					Code: 403,
				})
				mockService.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(nil)
				return mockService
			},
			ExpectedBody: map[string]interface{}{
				"full_name":    "jonathan",
				"phone_number": "08123456789",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {

			mockService := tc.Mock()

			server := Server{
				Service: mockService,
			}

			e := echo.New()

			requestBytes, _ := json.Marshal(tc.RequestBody)
			req := httptest.NewRequest(tc.HttpMethod, "/update", bytes.NewBuffer(requestBytes))
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()

			ctx := e.NewContext(req, rec)
			server.Update(ctx, domain.UpdateUserRequest{})

			assert.Equal(t, tc.ExpectedStatus, rec.Code)
		})
	}
}
