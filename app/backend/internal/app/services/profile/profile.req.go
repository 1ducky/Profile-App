package profile

import "io"

type UpdateProfileRequest struct {
	Username string `json:"username"`
	Tag      string `json:"tag"`
	Bio      string `json:"bio"`
}

type UploadAvatarRequest struct {
	id     string
	reader io.Reader
}
