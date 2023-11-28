package main

import (
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/handler"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/SawitProRecruitment/UserService/service"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	server := newServer()

	generated.RegisterHandlers(e, server)
	e.Logger.Fatal(e.Start(":1323"))
}

func newServer() *handler.Server {
	var repo repository.RepositoryInterface = repository.NewRepository(repository.NewRepositoryOptions{
		ConnectionUrl: os.Getenv("DATABASE_URL"),
	})

	//"postgres://postgres:nokia3500@localhost:5433/postgres?sslmode=disable",

	var credentialValidation service.CredentialValidationServiceInterface = service.NewCredentialValidationService()

	var service service.ServiceInterface = service.NewUserService(repo, credentialValidation)
	opts := handler.NewServerOptions{
		Service: service,
	}
	return handler.NewServer(opts)
}
