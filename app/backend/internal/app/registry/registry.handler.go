package registry

import (
	articleApp "profile-template/internal/app/services/article"
	assetsApp "profile-template/internal/app/services/assets"
	authApp "profile-template/internal/app/services/auth"
	profileApp "profile-template/internal/app/services/profile"
)

type RegisterHandler struct {
	AuthHandler    authApp.Handler
	ProfileHandler profileApp.Handler
	AssetsHandler  assetsApp.Handler
	ArticleHandler articleApp.Handler
}

func NewRegisterHandler(apps RegisterApp) RegisterHandler {
	return RegisterHandler{
		AuthHandler:    authApp.NewHandler(apps.authService),
		ProfileHandler: profileApp.NewHandler(apps.profileService),
		AssetsHandler:  assetsApp.NewHandler(apps.assetsService),
		ArticleHandler: articleApp.NewHandler(apps.articleService),
	}
}
