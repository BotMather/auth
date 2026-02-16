package auth

import "errors"

var (
	ErrUserAlreadyExists       = errors.New("user already exists")
	ErrInvalidRefreshToken     = errors.New("invalid refresh token")
	ErrPhoneNumberNotConfirmed = errors.New("phone number not confirmed")
	ErrInvalidPassword         = errors.New("Invalid password")
)
