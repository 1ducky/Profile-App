package registry

import (
	articleApp "profile-template/internal/app/services/article"
	assetsApp "profile-template/internal/app/services/assets"
	authApp "profile-template/internal/app/services/auth"
	profileApp "profile-template/internal/app/services/profile"
)

type RegisterApp struct {
	authService    authApp.AuthApp
	profileService profileApp.ProfileApp
	assetsService  assetsApp.AssetApp
	articleService articleApp.ArticleApp
}

func NewRegisterApp(ucs RegisterUC, uows RegisterUoW) RegisterApp {
	return RegisterApp{
		authService:    *authApp.NewAuthApp(ucs.auth, ucs.profile, uows.RegisterCase),
		profileService: *profileApp.NewProfileApp(ucs.profile, uows.ProfileCase),
		assetsService:  *assetsApp.NewAssetsApp(ucs.assets),
		articleService: *articleApp.NewArticleApp(ucs.Article, ucs.profile, uows.ArticleCase),
	}
}
