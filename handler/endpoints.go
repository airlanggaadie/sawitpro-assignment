package handler

import (
	"errors"
	"net/http"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// This is just a health check endpoint to get service information.
// (GET /healthz)
func (s *Server) Healthz(ctx echo.Context) error {
	response := generated.HealthCheckResponse{Status: "OK"}
	return ctx.JSON(http.StatusOK, response)
}

// This endpoint will check the database whether the combination exists.
// (POST /login)
func (s *Server) Login(ctx echo.Context) error {
	var bodyRequest generated.LoginRequest
	err := ctx.Bind(&bodyRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{
			Message: "invalid payload",
		})
	}

	response, err := s.usecase.Login(ctx.Request().Context(), bodyRequest)
	if err != nil {
		if errors.Is(err, model.ErrAuthentication) {
			return ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{
				Message: "wrong phone number / password",
			})
		}

		return err
	}

	return ctx.JSON(http.StatusOK, response)
}

// This endpoint will check the database and giving user information.
// (GET /profile)
func (s *Server) Profile(ctx echo.Context) error {
	userId, ok := ctx.Get("userId").(string)
	if userId == "" || !ok {
		return ctx.NoContent(http.StatusForbidden)
	}

	userUUID, err := uuid.Parse(userId)
	if err != nil {
		return ctx.NoContent(http.StatusForbidden)
	}

	profile, err := s.usecase.GetProfile(ctx.Request().Context(), userUUID)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, profile)
}

// This endpoint should store the newly created user in the database. The password should be hashed and salted in the database. Successful request should return the ID of the user.
// (POST /register)
func (s *Server) Register(ctx echo.Context) error {
	var bodyRequest generated.RegisterRequest
	err := ctx.Bind(&bodyRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{
			Message: "invalid payload",
		})
	}

	response, err := s.usecase.Register(ctx.Request().Context(), bodyRequest)
	if err != nil {
		if errors.Is(err, model.ErrDuplicateData) {
			return ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{
				Message: "phone number already registered",
			})
		}

		return err
	}

	return ctx.JSON(http.StatusOK, response)
}

// This endpoint will update user information.
// (PUT /profile)
func (s *Server) UpdateProfile(ctx echo.Context) error {
	userId, ok := ctx.Get("userId").(string)
	if userId == "" || !ok {
		return ctx.NoContent(http.StatusForbidden)
	}

	userUUID, err := uuid.Parse(userId)
	if err != nil {
		return ctx.NoContent(http.StatusForbidden)
	}

	var bodyRequest generated.UpdateProfileRequest
	err = ctx.Bind(&bodyRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{
			Message: "invalid payload",
		})
	}

	response, err := s.usecase.UpdateProfile(ctx.Request().Context(), userUUID, bodyRequest)
	if err != nil {
		if errors.Is(err, model.ErrDuplicateData) {
			return ctx.NoContent(http.StatusConflict)
		}

		return err
	}

	return ctx.JSON(http.StatusOK, response)
}
