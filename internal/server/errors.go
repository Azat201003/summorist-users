package server

import (
	"errors"
)

var (
	ErrNotPermitted = errors.New("Not permitted")
	ErrBadRequest = errors.New("Bad request")
)

