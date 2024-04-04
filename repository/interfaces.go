// This file contains the interfaces for the repository layer.
// The repository layer is responsible for interacting with the database.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package repository

import (
	"context"
	"time"

	"github.com/SawitProRecruitment/UserService/model"
	"github.com/google/uuid"
)

type PostgresqlRepositoryInterface interface {
	// GetUserById
	GetUserById(ctx context.Context, id uuid.UUID) (model.User, error)

	// GetUserByPhoneNumber
	GetUserByPhonenumber(ctx context.Context, phonenumber string) (model.User, error)

	// GetUserPasswordById
	GetUserPasswordById(ctx context.Context, id uuid.UUID) (model.UserPassword, error)

	// CheckPhonenumberExists
	CheckPhonenumberExists(ctx context.Context, phonenumber string) (bool, error)

	// CountLoginSession
	CountLoginSession(ctx context.Context, userId uuid.UUID) error

	// InsertNewUser
	InsertNewUser(ctx context.Context, newUser model.User, userAuth model.UserPassword) (uuid.UUID, error)

	// UpdateUser
	UpdateUser(ctx context.Context, userId uuid.UUID, fullname, phonenumber string) (uuid.UUID, error)
}

type JWTRepositoryInterface interface {
	Generate(userId uuid.UUID, additionalClaims map[string]string, expiry time.Duration) (string, error)
	Verify(jwt string) (uuid.UUID, error)
}
