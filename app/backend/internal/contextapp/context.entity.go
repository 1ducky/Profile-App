package contextapp

import "profile-template/internal/auth"

type SessionContext struct {
	ID      string
	Email   string
	Version int
	Role    auth.Role
}
