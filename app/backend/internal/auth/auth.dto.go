package auth

import "time"

type AuthEntity struct {
	ID       string
	Username string
	Role     Role
	Version  int
}
type AuthLogin struct {
	UserID   string
	Email    string
	Password string
	Username string
	Role     Role
}
type RegisterPayload struct {
	CredentialID string
	UserID       string
	Email        string
	Password     string
}
type LoginPayload struct {
	Email    string
	Password string
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// Session
type CreateSessionPayload struct {
	SessionID    string
	AccessToken  string
	RefreshToken string
	UserID       string
	RevokeAt     time.Time
}
type UpdateSessionPayload struct {
	AccessToken     string
	OldrefreshToken string
	RefreshToken    string
	UserID          string
	RevokeAt        time.Time
}

type DeleteSessionPayload struct {
	SessionID    string
	UserID       string
	RefreshToken string
}
