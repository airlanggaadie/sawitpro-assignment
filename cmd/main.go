package main

import (
	"log"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/handler"
	"github.com/SawitProRecruitment/UserService/repository/postgresql"
	"github.com/SawitProRecruitment/UserService/usecase"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	var server generated.ServerInterface = newServer()

	generated.RegisterHandlers(e, server)
	e.Logger.Fatal(e.Start(":1323"))
}

func newServer() *handler.Server {
	db, err := initPostgresql()
	if err != nil {
		log.Fatalln(err)
	}

	postgresql := postgresql.NewPostgresqlRepository(db)

	usecase := usecase.NewUsecase(postgresql)

	opts := handler.NewServerOptions{
		Usecase: usecase,
	}

	return handler.NewServer(opts)
}
