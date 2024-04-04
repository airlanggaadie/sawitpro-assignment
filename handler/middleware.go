package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var (
	ErrNoAuthHeader      = errors.New("authorization header is missing")
	ErrInvalidAuthHeader = errors.New("authorization header is malformed")
)

// // GetJWTFromRequest extracts a JWS string from an Authorization: Bearer <jws> header
func (s Server) getJWTFromRequest(req *http.Request) (string, error) {
	authHdr := req.Header.Get("Authorization")
	// Check for the Authorization header.
	if authHdr == "" {
		return "", ErrNoAuthHeader
	}
	// We expect a header value of the form "Bearer <token>", with 1 space after
	// Bearer, per spec.
	prefix := "Bearer "
	if !strings.HasPrefix(authHdr, prefix) {
		return "", ErrInvalidAuthHeader
	}
	return strings.TrimPrefix(authHdr, prefix), nil
}

func (s Server) authenticate(ctx echo.Context) (uuid.UUID, error) {
	jwt, err := s.getJWTFromRequest(ctx.Request())
	if err != nil {
		return uuid.Nil, err
	}

	userId, err := s.jwt.Verify(jwt)
	if err != nil {
		return uuid.Nil, err
	}

	return userId, nil
}
