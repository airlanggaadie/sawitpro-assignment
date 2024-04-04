// This file contains error list that are used in the handler, usecase, and repository layer.
package model

import "errors"

var (
	ErrAuthentication error = errors.New("invalid authentication")
	ErrDuplicateData  error = errors.New("duplicate data")
)
