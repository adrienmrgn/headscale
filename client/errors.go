package client

import (
	"errors"
)

var (
	// ErrUserNotFound : error returned when a headscale user is not found
	ErrUserNotFound = errors.New("Headscale : User not found")
	ErrUnauthorized = errors.New("Headscale : Unauthorized")
)
