package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/handler"
	"github.com/SawitProRecruitment/UserService/repository/postgresql"
	"github.com/SawitProRecruitment/UserService/usecase"

	"github.com/labstack/echo/v4"

	_ "github.com/lib/pq"
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

func initPostgresql() (*sql.DB, error) {
	dbDsn := os.Getenv("DATABASE_URL")
	if dbDsn == "" {
		return nil, fmt.Errorf("please configure DATABASE_URL environment variable")
	}

	db, err := sql.Open("postgres", dbDsn)
	if err != nil {
		return nil, fmt.Errorf("[main][initPostgresql] open database error: %v", err)
	}

	return db, nil
}
