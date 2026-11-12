package registry

import (
	"profile-template/internal/article"
	"profile-template/internal/assets"
	"profile-template/internal/auth"
	"profile-template/internal/config"
	"profile-template/internal/profile"
)

type RegisterUC struct {
	auth    auth.UsecaseContract
	profile profile.UsecaseContract
	assets  assets.UsecaseContract
	Article article.ArticleUsecaseContract
}

func NewRegisterUC(repos RegisterRepo, conf config.AppConfig) RegisterUC {
	return buildRegisterUC(repos, conf)
}

func buildRegisterUC(repos RegisterRepo, conf config.AppConfig) RegisterUC {
	return RegisterUC{
		auth:    auth.NewAuthService(repos.Cred, repos.Token, repos.Session),
		profile: profile.NewUsecase(repos.Profile),
		assets:  assets.NewUsecase(conf.StorageConfig, repos.BucketPersistent, repos.Assets),
		Article: article.NewUsecase(repos.Article),
	}
}
