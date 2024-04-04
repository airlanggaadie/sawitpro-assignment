// This file contains types that are used in the usecase and repository layer.
package model

import (
	"fmt"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/repository/utils"
	"github.com/google/uuid"
)

func NewUser(fullname, phonenumber string) (User, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return User{}, fmt.Errorf("[model][newuser] generate id error: %v", err)
	}

	return User{
		Id:          id,
		Fullname:    fullname,
		Phonenumber: phonenumber,
	}, nil
}

type User struct {
	Id          uuid.UUID
	Fullname    string
	Phonenumber string
}

func (u *User) ProfileResponse() generated.ProfileResponse {
	return generated.ProfileResponse{
		Fullname:    u.Fullname,
		Phonenumber: u.Phonenumber,
	}
}

func NewPassword(id uuid.UUID, password string) (UserPassword, error) {
	var userPassword UserPassword
	userPassword.Id = id
	if err := userPassword.setPassword(password); err != nil {
		return UserPassword{}, fmt.Errorf("[model][newpassword] set password error: %v", err)
	}

	return userPassword, nil
}

type UserPassword struct {
	Id       uuid.UUID
	password string
}

func (u *UserPassword) setPassword(password string) error {
	hashPassword, err := utils.HashPassword(password)
	if err != nil {
		return fmt.Errorf("[model][newpassword][setpassword] hash password error: %v", err)
	}

	u.password = hashPassword
	return nil
}

func (u *UserPassword) GetPassword() string {
	return u.password
}
