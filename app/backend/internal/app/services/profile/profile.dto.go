package profile

import "profile-template/internal/assets"

type ProfileResource struct {
	Username string `json:"username"`
	Tag      string `json:"tag"`
	Avatar   string `json:"avatar"`
	Bio      string `json:"bio"`
}

type profileAvatarResource struct {
	Avatar   string               `json:"avatar"`
	Metadata assets.AssetMetaData `json:"metadata"`
}
