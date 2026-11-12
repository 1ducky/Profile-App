package registry

import (
	"database/sql"
	articleApp "profile-template/internal/app/services/article"
	authApp "profile-template/internal/app/services/auth"
	profileApp "profile-template/internal/app/services/profile"
	"profile-template/internal/app/uow"
	"profile-template/internal/article"
	"profile-template/internal/assets"
	"profile-template/internal/auth"
	"profile-template/internal/auth/credential"
	"profile-template/internal/auth/jwt"
	"profile-template/internal/auth/session"
	"profile-template/internal/config"
	"profile-template/internal/db"
	"profile-template/internal/localstorage"
	"profile-template/internal/profile"
)

type RegisterUoW struct {
	RegisterCase uow.UnitOfWork[authApp.TXUsecase]
	ProfileCase  uow.UnitOfWork[profileApp.TXUsecase]
	ArticleCase  uow.UnitOfWork[articleApp.TXUsecase]
}

func NewUoWCase(persistent *sql.DB, conf config.AppConfig) RegisterUoW {
	return RegisterUoW{
		RegisterCase: buildRegisterCase(persistent, conf),
		ProfileCase:  buildProfileCase(persistent, conf),
		ArticleCase:  buildArticleCase(persistent, conf),
	}
}

func buildRegisterCase(persistent *sql.DB, conf config.AppConfig) uow.UnitOfWork[authApp.TXUsecase] {
	return uow.NewUoW(persistent, func(tx db.DBTX) authApp.TXUsecase {
		return authApp.TXUsecase{
			Auth:    auth.NewAuthService(credential.NewMySQLRepository(tx), jwt.NewJWT(conf.AuthConfig), session.NewMySQLRepository(tx)),
			Profile: profile.NewUsecase(profile.NewMySQLRepository(tx)),
		}
	})
}
func buildProfileCase(persistent *sql.DB, conf config.AppConfig) uow.UnitOfWork[profileApp.TXUsecase] {
	return uow.NewUoW(persistent, func(tx db.DBTX) profileApp.TXUsecase {
		bucket, err := localstorage.NewLocalStorage(conf.StorageConfig)
		if err != nil {
			return profileApp.TXUsecase{}
		}
		return profileApp.TXUsecase{
			Profile: profile.NewUsecase(profile.NewMySQLRepository(tx)),
			Asset:   assets.NewUsecase(conf.StorageConfig, bucket, assets.NewMysqlRepo(tx)),
		}
	})
}

func buildArticleCase(persistent *sql.DB, conf config.AppConfig) uow.UnitOfWork[articleApp.TXUsecase] {
	return uow.NewUoW(persistent, func(tx db.DBTX) articleApp.TXUsecase {
		return articleApp.TXUsecase{
			Article: article.NewUsecase(article.NewMySQLRepository(tx)),
			Asset:   assets.NewUsecase(conf.StorageConfig, nil, assets.NewMysqlRepo(tx)),
		}
	})
}
