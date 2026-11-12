package profile

import (
	"profile-template/internal/db"
	"time"
)

type Profile struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Tag       string    `json:"tag"`
	Bio       string    `json:"bio"`
	AvatarURL string    `json:"avatar_url"`
	CoverURL  string    `json:"cover_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

const TABLE db.Table = "profiles"

const (
	ID         db.Collom = "id"       //primary key
	USERNAME   db.Collom = "username" //displayname
	TAG        db.Collom = "tag"      //unique //indexed
	BIO        db.Collom = "bio"      //
	AVATAR_URL db.Collom = "avatar_url"
	COVER_URL  db.Collom = "cover_url"
	CREATED_AT db.Collom = "created_at"
	UPDATED_AT db.Collom = "updated_at"
)
