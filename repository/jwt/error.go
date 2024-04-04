package jwt

import "errors"

var ErrTokenParse = errors.New("error parsing jwt token")
var ErrTokenInvalid = errors.New("invalid jwt token")
var ErrTokenNotAcceptable = errors.New("jwt token not acceptable")
var ErrTokenExpired = errors.New("token expired")
