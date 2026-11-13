package profile

type ProfileResource struct {
	Username string `json:"username"`
	Tag      string `json:"tag"`
	Avatar   string `json:"avatar"`
	Bio      string `json:"bio"`
}

type profileAvatarResource struct {
	Avatar string `json:"avatar"`
}
