package handler

import (
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/SawitProRecruitment/UserService/usecase"
)

type Server struct {
	usecase usecase.UsecaseInterface
	jwt     repository.JWTRepositoryInterface
}

type NewServerOptions struct {
	Usecase usecase.UsecaseInterface
	Jwt     repository.JWTRepositoryInterface
}

func NewServer(opts NewServerOptions) *Server {
	return &Server{
		usecase: opts.Usecase,
		jwt:     opts.Jwt,
	}
}
