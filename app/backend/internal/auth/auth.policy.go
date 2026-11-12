package auth

import "time"

type Role string

const (
	Admin Role = "ADMIN"
	User  Role = "USER" // default
)

func ParseRole(role string) (Role, error) {
	switch Role(role) {
	case Admin, User:
		return Role(role), nil
	default:
		return "", ErrRoleInvalid
	}
}

func (r Role) String() string {
	return string(r)
}

func (r Role) Compare(diff Role) bool {
	return r == diff
}

const PasswordLength int = 8
const (
	AccessTokenDuration  time.Duration = 5 * time.Minute
	RefreshTokenDuration time.Duration = 60 * 24 * 30 * time.Minute
)
