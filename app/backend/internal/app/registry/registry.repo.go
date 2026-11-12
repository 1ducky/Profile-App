package registry

import (
	"profile-template/internal/article"
	"profile-template/internal/assets"
	"profile-template/internal/auth"
	"profile-template/internal/auth/credential"
	"profile-template/internal/auth/jwt"
	"profile-template/internal/auth/session"
	"profile-template/internal/config"
	"profile-template/internal/profile"
)

type RegisterRepo struct {
	Cred             auth.CredentialRepositoryContract
	Token            auth.TokenContract
	Session          auth.SessionRepositoryContract
	Profile          profile.RepoContract
	Assets           assets.AssetRepository
	Article          article.ArticlaeRepositoryContract
	BucketPersistent assets.RepositoryContract
}

func NewRegisterRepo(persistent RegisterPersistent, conf config.AppConfig) RegisterRepo {
	return RegisterRepo{
		Cred:             credential.NewMySQLRepository(persistent.DB),
		Token:            jwt.NewJWT(conf.AuthConfig),
		Session:          session.NewMySQLRepository(persistent.DB),
		Profile:          profile.NewMySQLRepository(persistent.DB),
		Assets:           assets.NewMysqlRepo(persistent.DB),
		Article:          article.NewMySQLRepository(persistent.DB),
		BucketPersistent: persistent.Bucket,
	}
}
