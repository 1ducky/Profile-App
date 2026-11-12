package registry

import (
	"profile-template/internal/assets"
	"profile-template/internal/db"
)

type RegisterPersistent struct {
	DB     db.DBTX
	Bucket assets.RepositoryContract
}

func NewRegisterPersistent(DB db.DBTX, Bucket assets.RepositoryContract) RegisterPersistent {
	return RegisterPersistent{
		DB:     DB,
		Bucket: Bucket,
	}
}
