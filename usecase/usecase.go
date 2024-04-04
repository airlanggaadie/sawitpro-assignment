package usecase

import (
	"github.com/SawitProRecruitment/UserService/repository"
)

type usecase struct {
	repository repository.PostgresqlRepositoryInterface
	jwt        repository.JWTRepositoryInterface
}

func NewUsecase(repository repository.PostgresqlRepositoryInterface, jwt repository.JWTRepositoryInterface) UsecaseInterface {
	return &usecase{
		repository: repository,
		jwt:        jwt,
	}
}
