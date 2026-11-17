package registry

import (
	"profile-template/internal/article"
	"profile-template/internal/article/category"
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
	ArticleCategory  article.CategoryRepositoryContract
	BucketPersistent assets.RepositoryContract
}

func NewRegisterRepo(persistent RegisterPersistent, conf config.AppConfig) RegisterRepo {
	return RegisterRepo{
		Cred:             credential.NewPGRepository(persistent.DB),
		Token:            jwt.NewJWT(conf.AuthConfig),
		Session:          session.NewPGSQLRepository(persistent.DB),
		Profile:          profile.NewPostgresRepository(persistent.DB),
		Assets:           assets.NewPostgresRepo(persistent.DB),
		Article:          article.NewPGSQLRepository(persistent.DB),
		BucketPersistent: persistent.Bucket,
		ArticleCategory:  category.NewPQSQLRepository(persistent.DB),
	}
}
