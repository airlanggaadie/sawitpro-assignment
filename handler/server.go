package handler

import (
	"github.com/SawitProRecruitment/UserService/usecase"
)

type Server struct {
	usecase usecase.UsecaseInterface
}

type NewServerOptions struct {
	Usecase usecase.UsecaseInterface
}

func NewServer(opts NewServerOptions) *Server {
	return &Server{
		usecase: opts.Usecase,
	}
}
