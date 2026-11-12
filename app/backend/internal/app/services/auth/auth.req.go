package auth

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type UpdateProfileRequest struct {
	Username string `json:"username"`
	Tag      string `json:"tag"`
	Bio      string `json:"bio"`
}
