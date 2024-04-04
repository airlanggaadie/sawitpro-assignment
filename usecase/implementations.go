package usecase

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/model"
	"github.com/SawitProRecruitment/UserService/repository/utils"
	"github.com/google/uuid"
)

// GetProfile implements UsecaseInterface.
func (u usecase) GetProfile(ctx context.Context, userId uuid.UUID) (generated.ProfileResponse, error) {
	user, err := u.repository.GetUserById(ctx, userId)
	if err != nil {
		return generated.ProfileResponse{}, fmt.Errorf("[usecase][getprofile] get user error: %v", err)
	}

	return user.ProfileResponse(), nil
}

// Login implements UsecaseInterface.
func (u usecase) Login(ctx context.Context, request generated.LoginRequest) (generated.LoginResponse, error) {
	user, err := u.repository.GetUserByPhonenumber(ctx, request.Phonenumber)
	if err != nil {
		return generated.LoginResponse{}, fmt.Errorf("[usecase][login] get user error: %v", err)
	}

	userPassword, err := u.repository.GetUserPasswordById(ctx, user.Id)
	if err != nil {
		return generated.LoginResponse{}, fmt.Errorf("[usecase][login] get user password error: %v", err)
	}

	if !utils.VerifyPassword(request.Password, userPassword.GetPassword()) {
		return generated.LoginResponse{}, model.ErrAuthentication
	}

	// TODO: generate token when verified
	token, err := u.jwt.Generate(user.Id, map[string]string{}, 30*time.Minute)
	if err != nil {
		return generated.LoginResponse{}, fmt.Errorf("[usecase][login] generate token error: %v", err)
	}

	go u.countLoginSession(user.Id)

	return generated.LoginResponse{
		Id:    user.Id,
		Token: token,
	}, nil
}

func (u usecase) countLoginSession(userId uuid.UUID) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := u.repository.CountLoginSession(ctx, userId)
	if err != nil {
		log.Printf("[usecase][countloginsession] error: %v", err)
	}
}

// Register implements UsecaseInterface.
func (u usecase) Register(ctx context.Context, request generated.RegisterRequest) (generated.RegisterResponse, error) {
	exists, err := u.repository.CheckPhonenumberExists(ctx, request.Phonenumber)
	if err != nil {
		return generated.RegisterResponse{}, fmt.Errorf("[usecase][register] check phonenumber error: %v", err)
	}

	if exists {
		return generated.RegisterResponse{}, model.ErrDuplicateData
	}

	user, err := model.NewUser(request.Fullname, request.Phonenumber)
	if err != nil {
		return generated.RegisterResponse{}, fmt.Errorf("[usecase][register] new user error: %v", err)
	}

	userPassword, err := model.NewPassword(user.Id, request.Password)
	if err != nil {
		return generated.RegisterResponse{}, fmt.Errorf("[usecase][register] new user passsword error: %v", err)
	}

	id, err := u.repository.InsertNewUser(ctx, user, userPassword)
	if err != nil {
		return generated.RegisterResponse{}, fmt.Errorf("[usecase][register] insert user error: %v", err)
	}

	return generated.RegisterResponse{Id: id}, nil
}

// UpdateProfile implements UsecaseInterface.
func (u usecase) UpdateProfile(ctx context.Context, userId uuid.UUID, request generated.UpdateProfileRequest) (generated.UpdateProfileResponse, error) {
	// TODO: get userData information by id
	userData, err := u.repository.GetUserById(ctx, userId)
	if err != nil {
		return generated.UpdateProfileResponse{}, fmt.Errorf("[usecase][updateprofile] get user error: %v", err)
	}

	if userData.Phonenumber != request.Phonenumber {
		exists, err := u.repository.CheckPhonenumberExists(ctx, request.Phonenumber)
		if err != nil {
			return generated.UpdateProfileResponse{}, fmt.Errorf("[usecase][updateprofile] check phonenumber error: %v", err)
		}

		if exists {
			return generated.UpdateProfileResponse{}, model.ErrDuplicateData
		}
	}

	id, err := u.repository.UpdateUser(ctx, userData.Id, request.Fullname, request.Phonenumber)
	if err != nil {
		return generated.UpdateProfileResponse{}, fmt.Errorf("[usecase][updateprofile] update user error: %v", err)
	}

	return generated.UpdateProfileResponse{Id: id}, nil
}
