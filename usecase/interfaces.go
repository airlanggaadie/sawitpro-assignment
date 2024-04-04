package usecase

import (
	"context"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/google/uuid"
)

type UsecaseInterface interface {
	Register(ctx context.Context, request generated.RegisterRequest) (generated.RegisterResponse, error)
	Login(ctx context.Context, request generated.LoginRequest) (generated.LoginResponse, error)
	GetProfile(ctx context.Context, userId uuid.UUID) (generated.ProfileResponse, error)
	UpdateProfile(ctx context.Context, userId uuid.UUID, request generated.UpdateProfileRequest) (generated.UpdateProfileResponse, error)
}
