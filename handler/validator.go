package handler

import (
	"errors"
	"regexp"
	"strings"

	"github.com/SawitProRecruitment/UserService/generated"
)

func (s Server) validateLoginRequest(request generated.LoginRequest) error {
	errMessage := []string{}

	errMessagePhonenumber := s.validatePhonenumber(request.Phonenumber)
	errMessage = append(errMessage, errMessagePhonenumber...)

	errMessagePassword, err := s.validatePassword(request.Password)
	if err != nil {
		return err
	}

	errMessage = append(errMessage, errMessagePassword...)
	if len(errMessage) > 0 {
		return errors.New(strings.Join(errMessage, "|"))
	}

	return nil
}

func (s Server) validateRegisterRequest(request generated.RegisterRequest) error {
	errMessage := []string{}

	errMessageFullname := s.validateFullname(request.Fullname)
	errMessage = append(errMessage, errMessageFullname...)

	errMessagePhonenumber := s.validatePhonenumber(request.Phonenumber)
	errMessage = append(errMessage, errMessagePhonenumber...)

	errMessagePassword, err := s.validatePassword(request.Password)
	if err != nil {
		return err
	}

	errMessage = append(errMessage, errMessagePassword...)
	if len(errMessage) > 0 {
		return errors.New(strings.Join(errMessage, "|"))
	}

	return nil
}

func (s Server) validateUpdateRequest(request generated.UpdateProfileRequest) error {
	errMessage := []string{}

	errMessageFullname := s.validateFullname(request.Fullname)
	errMessage = append(errMessage, errMessageFullname...)

	errMessagePhonenumber := s.validatePhonenumber(request.Phonenumber)
	errMessage = append(errMessage, errMessagePhonenumber...)

	if len(errMessage) > 0 {
		return errors.New(strings.Join(errMessage, "|"))
	}

	return nil
}

func (s Server) validatePhonenumber(phonenumber string) []string {
	errMessage := []string{}
	prefix := "+62"
	if !strings.HasPrefix(phonenumber, prefix) {
		errMessage = append(errMessage, "phonenumber must begin with "+prefix)
	}

	if len(phonenumber) < 12 && len(phonenumber) > 15 {
		errMessage = append(errMessage, "phonenumber must be between 10 and 13 digits excluding "+prefix)
	}

	return errMessage
}

func (s Server) validatePassword(password string) ([]string, error) {
	errMessage := []string{}
	if len(password) < 6 && len(password) > 64 {
		errMessage = append(errMessage, "password must be between 6 and 64 characters")
	}

	regex, err := regexp.Compile(`[A-Z]+`)
	if err != nil {
		return nil, err
	}

	regexValidation := regex.MatchString(password)
	if !regexValidation {
		errMessage = append(errMessage, "password must contain at least 1 capital characters")
	}

	regex, err = regexp.Compile(`\d`)
	if err != nil {
		return nil, err
	}

	regexValidation = regex.MatchString(password)
	if !regexValidation {
		errMessage = append(errMessage, "password must contain at least 1 number")
	}
	regex, err = regexp.Compile(`[\W_]`)
	if err != nil {
		return nil, err
	}

	regexValidation = regex.MatchString(password)
	if !regexValidation {
		errMessage = append(errMessage, "password must contain at least 1 special characters")
	}

	return errMessage, nil
}

func (s Server) validateFullname(fullname string) []string {
	errMessage := []string{}
	if len(fullname) < 6 && len(fullname) > 60 {
		errMessage = append(errMessage, "fullname must be between 6 and 60 characters")
	}

	return errMessage
}
