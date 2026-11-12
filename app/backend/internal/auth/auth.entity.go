package auth

import (
	"profile-template/internal/db"
	"time"
)

type TokenEntity struct {
	ID       string
	Email    string
	UserID   string
	Role     Role
	Version  int
	RevokeAt time.Time
}

type CredentialEntity struct {
	ID       string
	Email    string
	Password string
	Version  int
	Role     Role
}

type SessionEntity struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	RevokeAt     time.Time `json:"revoke_at"`
	Version      int       `json:"version"`
}

const CREDENTIAL_TABLE db.Table = "credential"
const SESSION_TABLE db.Table = "session"

const (
	CREDENTIAL_ID       db.Collom = "id"       // primary
	CREDENTIAL_USER_ID  db.Collom = "user_id"  //unique
	CREDENTIAL_EMAIL    db.Collom = "email"    //uniquer
	CREDENTIAL_PASSWORD db.Collom = "password" // hashed
	CREDENTIAL_VERSION  db.Collom = "version"  // int
	CREDENTIAL_ROLE     db.Collom = "role"     // enum
)

const (
	SESSION_ID            db.Collom = "id"
	SESSION_USER_ID       db.Collom = "user_id"
	SESSION_ACCESS_TOKEN  db.Collom = "access_token"
	SESSION_REFRESH_TOKEN db.Collom = "refresh_token"
	SESSION_REVOKE_AT     db.Collom = "revoke_at"
	SESSION_VERSION       db.Collom = "version"
)
