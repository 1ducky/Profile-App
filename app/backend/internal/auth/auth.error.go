package auth

import (
	"errors"
	"profile-template/internal/app/err"
)

// Sentinel errors. Usecase & repository layer harus me-wrap error asli
var (
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid password")
	ErrPasswordShort      = errors.New("password is too short")
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrInvalidEmail       = errors.New("email is invalid")
	ErrUnauthorized       = errors.New("unauthorized")
	ErrTokenInvalid       = errors.New("invalid token")
	ErrTokenExpired       = errors.New("token expired")
	ErrRoleInvalid        = errors.New("invalid role")
)

// Error Code Contract Frontend
const (
	CodeUserNotFound       err.ErrCode = "USER_NOT_FOUND"
	CodeInvalidCredentials err.ErrCode = "INVALID_CREDENTIALS"
	CodePasswordShort      err.ErrCode = "PASSWORD_SHORT"
	CodeEmailAlreadyExists err.ErrCode = "EMAIL_ALREADY_EXISTS"
	CodeInvalidEmail       err.ErrCode = "INVALID_EMAIL"
	CodeUnauthorized       err.ErrCode = "UNAUTHORIZED"
	CodeTokenInvalid       err.ErrCode = "TOKEN_INVALID"
	CodeTokenExpired       err.ErrCode = "TOKEN_EXPIRED"
	CodeRoleInvalid        err.ErrCode = "ROLE_INVALID"
	CodeInternal           err.ErrCode = "INTERNAL_ERROR"
)
