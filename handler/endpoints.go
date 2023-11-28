package handler

import (
	"fmt"
	"github.com/SawitProRecruitment/UserService/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) Register(ctx echo.Context, params domain.CreateUserRequest) error {
	errList, err := s.Service.ValidateUserRequest(&params)
	if len(errList) > 0 {
		return ctx.JSON(http.StatusBadRequest, errList)
	}

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	res, err := s.Service.Register(ctx.Request().Context(), &params)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
	}
	return ctx.JSON(http.StatusCreated, res)
}

func (s *Server) Login(ctx echo.Context, params domain.LoginRequest) error {
	res, err := s.Service.Login(ctx.Request().Context(), &params)
	if err != nil && err.Code == 400 {
		return ctx.JSON(http.StatusBadRequest, "Invalid Password")

	}
	if err != nil {
		fmt.Errorf("Error when login : %v, err: %v", params.PhoneNumber, err)
		return ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
	}
	return ctx.JSON(http.StatusOK, res)
}

func (s *Server) GetUser(ctx echo.Context) error {
	jwtToken := ctx.Request().Header.Get("Authorization")
	claims, err := s.Service.ValidateToken(jwtToken)
	if err != nil && err.Code == 403 {
		return ctx.JSON(http.StatusForbidden, err)
	}

	res, err := s.Service.GetUser(ctx.Request().Context(), claims)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
	}
	return ctx.JSON(http.StatusOK, res)
}

func (s *Server) Update(ctx echo.Context, params domain.UpdateUserRequest) error {
	jwtToken := ctx.Request().Header.Get("Authorization")
	claims, err := s.Service.ValidateToken(jwtToken)
	if err != nil && err.Code == 403 {
		return ctx.JSON(http.StatusForbidden, err)
	}

	res, err := s.Service.Update(ctx.Request().Context(), &params, claims)
	if err != nil && err.Code == 409 {
		return ctx.JSON(http.StatusConflict, err)
	}
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
	}
	return ctx.JSON(http.StatusOK, res)
}
