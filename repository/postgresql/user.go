package postgresql

import (
	"context"

	"github.com/SawitProRecruitment/UserService/model"
	"github.com/google/uuid"
)

// GetUserById implements repository.PostgresqlRepositoryInterface.
func (p *postgresqlRepository) GetUserById(ctx context.Context, id uuid.UUID) (model.User, error) {

	return model.User{}, nil
}

// GetUserByPhonenumber implements repository.PostgresqlRepositoryInterface.
func (p *postgresqlRepository) GetUserByPhonenumber(ctx context.Context, phonenumber string) (model.User, error) {
	return model.User{}, nil
}

// GetUserPasswordById implements repository.PostgresqlRepositoryInterface.
func (p *postgresqlRepository) GetUserPasswordById(ctx context.Context, id uuid.UUID) (model.UserPassword, error) {
	return model.UserPassword{}, nil
}

// CheckPhonenumberExists implements repository.PostgresqlRepositoryInterface.
func (p *postgresqlRepository) CheckPhonenumberExists(ctx context.Context, phonenumber string) (bool, error) {
	return false, nil
}

// CountLoginSession implements repository.PostgresqlRepositoryInterface.
func (p *postgresqlRepository) CountLoginSession(ctx context.Context, userId uuid.UUID) error {
	return nil
}

// InsertNewUser implements repository.PostgresqlRepositoryInterface.
func (p *postgresqlRepository) InsertNewUser(ctx context.Context, newUser model.User, userAuth model.UserPassword) (uuid.UUID, error) {
	return uuid.UUID{}, nil
}

// UpdateUser implements repository.PostgresqlRepositoryInterface.
func (p *postgresqlRepository) UpdateUser(ctx context.Context, userId uuid.UUID, fullname string, phonenumber string) (uuid.UUID, error) {
	return uuid.UUID{}, nil
}
