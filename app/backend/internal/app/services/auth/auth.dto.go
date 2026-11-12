package auth

import (
	"profile-template/internal/auth"
)

type LoginResource struct {
	token   auth.TokenResponse
	Profile LoginSession
}
type LoginSession struct {
	SessionResource
	ProfileResource
}

type SessionResource struct {
	UserID  string `json:"id"`
	Email   string `json:"email"`
	Role    string `json:"role"`
	Version int    `json:"version"`
}

type ProfileResource struct {
	Username string `json:"username"`
	Tag      string `json:"tag"`
	Avatar   string `json:"avatar"`
	Bio      string `json:"bio"`
}
