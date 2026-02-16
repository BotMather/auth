package auth

import "errors"

var (
	ErrUserAlreadyExists       = errors.New("user already exists")
	ErrInvalidRefreshToken     = errors.New("invalid refresh token")
	ErrPhoneNumberNotConfirmed = errors.New("phone number not confirmed")
	ErrInvalidPassword         = errors.New("Invalid password")
	ErrInvalidCredentions      = errors.New("Invalid credentions")
	ErrRateLimit               = errors.New("It takes 2 minutes to resend the SMS.")
)
