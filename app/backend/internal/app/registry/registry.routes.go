package registry

import (
	"net/http"
	"profile-template/internal/app/err"
	articleApp "profile-template/internal/app/services/article"
	assetsApp "profile-template/internal/app/services/assets"
	authApp "profile-template/internal/app/services/auth"
	profileApp "profile-template/internal/app/services/profile"
	"profile-template/internal/transport"
	"profile-template/internal/transport/middleware"
	"strings"
)

func NewRegisterRoutes(handlers RegisterHandler, uc RegisterUC, adapter *transport.Adapter) http.Handler {
	mux := http.NewServeMux()
	routeAuth(mux, handlers.AuthHandler, adapter)
	routeProfile(mux, handlers.ProfileHandler, adapter)
	routeAssets(mux, handlers.AssetsHandler, adapter)
	routeArticle(mux, handlers.ArticleHandler, adapter)

	return middleware.Chain(
		mux,
		middleware.CORSMiddleware,
		middleware.TrackerMiddleware(),
		middleware.NewAuthMiddleware(uc.auth),
	)
}

func routeAuth(mux *http.ServeMux, handler authApp.Handler, adapter *transport.Adapter) {
	mux.HandleFunc("POST /auth/register", adapter.Adapt(handler.Register, err.Auth))
	mux.HandleFunc("POST /auth/login", adapter.Adapt(handler.Login, err.Auth))
	mux.HandleFunc("GET /auth/session", adapter.Adapt(handler.GetSession, err.Auth))
}

func routeProfile(mux *http.ServeMux, handler profileApp.Handler, adapter *transport.Adapter) {
	mux.HandleFunc("GET /profile", adapter.Adapt(handler.GetOwnProfile, err.Profile))
	mux.HandleFunc("PATCH /profile", adapter.Adapt(handler.UpdateProfile, err.Profile))
	mux.HandleFunc("POST /profile/avatar", adapter.Adapt(handler.UploadAvatar, err.Profile))
	mux.HandleFunc(strings.ReplaceAll("GET /profile/{*}", "*", profileApp.DynamicIDPath), adapter.Adapt(handler.GetProfileByID, err.Profile))
}
func routeAssets(mux *http.ServeMux, handler assetsApp.Handler, adapter *transport.Adapter) {
	mux.HandleFunc(strings.ReplaceAll("GET /assets/{*...}", "*", assetsApp.KEY), adapter.Adapt(handler.ServeAssets, err.Assets))
	mux.HandleFunc("POST /assets/upload", adapter.Adapt(handler.UploadAvatar, err.Assets))
}

func routeArticle(mux *http.ServeMux, handler articleApp.Handler, adapter *transport.Adapter) {
	mux.HandleFunc("POST /articles", adapter.Adapt(handler.Create, err.Article))
	mux.HandleFunc("GET /articles", adapter.Adapt(handler.List, err.Article))
	mux.HandleFunc("GET /articles/{slug}", adapter.Adapt(handler.GetBySlug, err.Article))
	mux.HandleFunc("GET /article/{id}", adapter.Adapt(handler.GetByID, err.Article))
	mux.HandleFunc("PATCH /articles/{id}", adapter.Adapt(handler.Update, err.Article))
	mux.HandleFunc("DELETE /articles/{id}", adapter.Adapt(handler.Delete, err.Article))
}
