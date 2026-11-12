package auth

import (
	"context"
	"profile-template/internal/auth"
)

type AuthAppInterface interface {
	Login(ctx context.Context, req auth.LoginPayload) (auth.TokenResponse, error)
	Register(ctx context.Context, req auth.RegisterPayload) error
}
