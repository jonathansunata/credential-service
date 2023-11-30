package generated

import (
	"fmt"
	"github.com/SawitProRecruitment/UserService/domain"
	"github.com/SawitProRecruitment/UserService/handler"
	"github.com/labstack/echo/v4"
)

func RegisterHandlers(e *echo.Echo, h *handler.Server) {
	e.POST("/register", func(ctx echo.Context) error {
		var params domain.CreateUserRequest
		if err := ctx.Bind(&params); err != nil {
			return err
		}
		return h.Register(ctx, params)
	})

	e.PUT("/login", func(ctx echo.Context) error {
		var params domain.LoginRequest
		fmt.Printf("Login user")
		if err := ctx.Bind(&params); err != nil {
			fmt.Printf("Error when login : %v \n", err)
			return err
		}
		return h.Login(ctx, params)
	})

	e.GET("/user", func(ctx echo.Context) error {
		return h.GetUser(ctx)
	})

	e.PATCH("/update", func(ctx echo.Context) error {
		var params domain.UpdateUserRequest
		if err := ctx.Bind(&params); err != nil {
			return err
		}
		return h.Update(ctx, params)
	})
}
