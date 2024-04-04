package jwt

import (
	"fmt"
	"time"

	"github.com/SawitProRecruitment/UserService/repository"
	jsonwebtoken "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func NewJwtRepository(secret string, issuer string) repository.JWTRepositoryInterface {
	return &jwtRepository{
		secret: []byte(secret),
		issuer: issuer,
	}
}

type jwtRepository struct {
	secret []byte
	issuer string
}

// Generate implements service.JwtRepository.
func (d jwtRepository) Generate(userId uuid.UUID, additionalClaims map[string]string, expiry time.Duration) (string, error) {
	token := jsonwebtoken.New(jsonwebtoken.SigningMethodHS256)
	claims, ok := token.Claims.(jsonwebtoken.MapClaims)
	if !ok {
		return "", fmt.Errorf("unable to cast claims")
	}

	claims["nbf"] = time.Now().Unix()
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(expiry).Unix()
	claims["id"] = userId.String()
	claims["iss"] = d.issuer
	claims["sub"] = d.issuer

	// this code is used to handle the value of jwt claims that can be dynamic value
	if len(additionalClaims) > 0 {
		for additonalClaimKey, additonalClaimValue := range additionalClaims {
			claims[additonalClaimKey] = additonalClaimValue
		}
	}

	s, err := token.SignedString(d.secret)
	if err != nil {
		return "", err
	}

	return s, nil
}

// Verify implements service.JwtRepository.
func (d jwtRepository) Verify(jwt string) (uuid.UUID, error) {
	token, err := jsonwebtoken.Parse(jwt, func(t *jsonwebtoken.Token) (interface{}, error) {
		_, ok := t.Method.(*jsonwebtoken.SigningMethodHMAC)
		if !ok {
			return nil, ErrTokenParse
		}
		return d.secret, nil
	})
	if err != nil {
		if token != nil && !token.Valid {
			if ve, ok := err.(*jsonwebtoken.ValidationError); ok {
				if ve.Errors&(jsonwebtoken.ValidationErrorExpired) != 0 {
					return uuid.Nil, ErrTokenExpired
				}

				if ve.Errors&(jsonwebtoken.ValidationErrorMalformed|jsonwebtoken.ValidationErrorSignatureInvalid|jsonwebtoken.ValidationErrorUnverifiable|jsonwebtoken.ValidationErrorNotValidYet) != 0 {
					return uuid.Nil, ErrTokenInvalid
				}

				if ve.Errors&(jsonwebtoken.ValidationErrorAudience|jsonwebtoken.ValidationErrorIssuedAt|jsonwebtoken.ValidationErrorClaimsInvalid|jsonwebtoken.ValidationErrorIssuer) != 0 {
					return uuid.Nil, ErrTokenNotAcceptable
				}
			}
		}

		return uuid.Nil, err
	}

	claims, ok := token.Claims.(jsonwebtoken.MapClaims)
	if !ok || !token.Valid {
		return uuid.Nil, ErrTokenInvalid
	}

	if claims["iss"] != d.issuer {
		return uuid.Nil, ErrTokenNotAcceptable
	}

	id, ok := claims["id"].(string)
	if !ok {
		return uuid.Nil, ErrTokenInvalid
	}

	userId, err := uuid.Parse(id)
	if err != nil {
		return uuid.Nil, ErrTokenInvalid
	}

	return userId, nil
}
