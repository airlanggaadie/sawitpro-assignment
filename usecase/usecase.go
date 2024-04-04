package usecase

import (
	"github.com/SawitProRecruitment/UserService/repository"
)

type usecase struct {
	repository repository.PostgresqlRepositoryInterface
}

func NewUsecase(repository repository.PostgresqlRepositoryInterface) UsecaseInterface {
	return &usecase{
		repository: repository,
	}
}
