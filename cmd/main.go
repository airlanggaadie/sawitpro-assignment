package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/handler"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/SawitProRecruitment/UserService/repository/jwt"
	"github.com/SawitProRecruitment/UserService/repository/postgresql"
	"github.com/SawitProRecruitment/UserService/usecase"

	"github.com/labstack/echo/v4"

	_ "github.com/lib/pq"
)

func main() {
	e := echo.New()
	e.Debug = true

	jwtAuthenticator, err := initJwt()
	if err != nil {
		log.Fatal(err)
	}

	var server generated.ServerInterface = newServer(jwtAuthenticator)

	generated.RegisterHandlers(e, server)
	e.Logger.Fatal(e.Start(":1323"))
}

func newServer(jwtAuthenticator repository.JWTRepositoryInterface) *handler.Server {
	db, err := initPostgresql()
	if err != nil {
		log.Fatalln(err)
	}

	postgresql := postgresql.NewPostgresqlRepository(db)

	usecase := usecase.NewUsecase(postgresql, jwtAuthenticator)

	opts := handler.NewServerOptions{
		Usecase: usecase,
		Jwt:     jwtAuthenticator,
	}

	return handler.NewServer(opts)
}

func initJwt() (repository.JWTRepositoryInterface, error) {
	secret, ok := os.LookupEnv("JWT_SECRET")
	if !ok {
		return nil, fmt.Errorf("please configure JWT_SECRET environment variable")
	}

	issuer, ok := os.LookupEnv("JWT_ISSUER")
	if !ok {
		return nil, fmt.Errorf("please configure JWT_ISSUER environment variable")
	}

	return jwt.NewJwtRepository(secret, issuer), nil
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
