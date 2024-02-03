package constant

import "errors"

// error message.
const (
	MsgHeaderTokenNotFound            = "header `authorization` not found"
	MsgHeaderRefreshTokenNotFound     = "header `refresh-token` not found"
	MsgHeaderTokenUnauthorized        = "unauthorized token"
	MsgHeaderRefreshTokenUnauthorized = "unauthorized refresh token"
	MsgHeaderTokenInvalid             = "invalid token"
	MsgHeaderRefreshTokenInvalid      = "invalid refresh token"
)

// error interface.
var (
	// 400.
	ErrFailedParseRequest       = errors.New("failed to parse request")
	ErrFileNotFound             = errors.New("file not found")
	ErrValidationFailed         = errors.New("validation failed")
	ErrPasswordResetLinkExpired = errors.New("your password reset link has expired")

	ErrPasswordIncorrect       = errors.New("password incorrect")
	ErrPasswordNotMatch        = errors.New("password not match")
	ErrConfirmPasswordNotMatch = errors.New("confirm password not match")

	// 401.
	ErrMissingHeaderData     = errors.New("missing header data")
	ErrUnauthorizedTokenData = errors.New("unauthorized token data")
	ErrInvalidTokenData      = errors.New("invalid token data")

	// 404.
	ErrAccountNotFound = errors.New("account not found")
	ErrDataNotFound    = errors.New("data not found")

	// 500.
	ErrUnknownSource          = errors.New("unknown error")
	ErrAccountNotHavePassword = errors.New("account does not have password")
)
