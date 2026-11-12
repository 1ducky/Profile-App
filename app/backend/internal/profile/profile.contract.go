package profile

import "context"

type UsecaseContract interface {
	CreateUser(ctx context.Context, req CreateRequest) error
	UpdateProfile(ctx context.Context, req UpdateRequest) error
	GetUserByID(ctx context.Context, id string) (Profile, error)
	GetUsersByIDs(ctx context.Context, id []string) ([]Profile, error)
	DeleteUser(ctx context.Context, id string) error

	UploadAvatar(ctx context.Context, id string, path string) error
	UploadCover(ctx context.Context, id string) error
}

type RepoContract interface {
	CreateUser(ctx context.Context, req CreateData) error
	UpdateProfile(ctx context.Context, req UpdateRequest) error
	GetUserByID(ctx context.Context, id string) (Profile, error)
	GetUsersByIDs(ctx context.Context, id []string) ([]Profile, error)
	DeleteUser(ctx context.Context, id string) error
	SetProfilePhoto(ctx context.Context, id string, path string) error
}
