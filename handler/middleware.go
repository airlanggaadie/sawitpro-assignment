package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v4"
	middleware "github.com/oapi-codegen/echo-middleware"
)

type contextKey string

const userIdKey contextKey = "userId"

var (
	ErrNoAuthHeader      = errors.New("authorization header is missing")
	ErrInvalidAuthHeader = errors.New("authorization header is malformed")
)

func CreateMiddleware(v repository.JWTRepositoryInterface) ([]echo.MiddlewareFunc, error) {
	spec, err := generated.GetSwagger()
	if err != nil {
		return nil, fmt.Errorf("loading spec: %w", err)
	}

	validator := middleware.OapiRequestValidatorWithOptions(spec,
		&middleware.Options{
			Options: openapi3filter.Options{
				AuthenticationFunc: NewAuthenticator(v),
			},
		})

	return []echo.MiddlewareFunc{validator}, nil
}

func NewAuthenticator(v repository.JWTRepositoryInterface) openapi3filter.AuthenticationFunc {
	return func(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
		return Authenticate(v, ctx, input)
	}
}

func Authenticate(v repository.JWTRepositoryInterface, ctx context.Context, input *openapi3filter.AuthenticationInput) error {
	// Our security scheme is named BearerAuth, ensure this is the case
	if input.SecuritySchemeName != "bearerAuth" {
		return fmt.Errorf("security scheme %s != 'bearerAuth'", input.SecuritySchemeName)
	}

	// Now, we need to get the JWS from the request, to match the request expectations
	// against request contents.
	jws, err := GetJWSFromRequest(input.RequestValidationInput.Request)
	if err != nil {
		return fmt.Errorf("getting jws: %w", err)
	}

	// if the JWS is valid, we have a JWT, which will contain a bunch of claims.
	token, err := v.Verify(jws)
	if err != nil {
		return fmt.Errorf("validating JWS: %w", err)
	}

	// Set the property on the echo context so the handler is able to
	// access the claims data we generate in here.
	echoCtx := middleware.GetEchoContext(ctx)
	authdCtx := context.WithValue(echoCtx.Request().Context(), userIdKey, token.String())
	requestWithAuth := echoCtx.Request().WithContext(authdCtx)
	echoCtx.SetRequest(requestWithAuth)

	return nil
}

// GetJWSFromRequest extracts a JWS string from an Authorization: Bearer <jws> header
func GetJWSFromRequest(req *http.Request) (string, error) {
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
