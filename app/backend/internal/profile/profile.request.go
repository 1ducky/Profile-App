package profile

type CreateRequest struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Tag      string `json:"tag"`
}

type UpdateRequest struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Tag      string `json:"tag"`
	Bio      string `json:"bio"`
}

type CreateData struct {
	ID       string
	Username string
	Tag      string
}
