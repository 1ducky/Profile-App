package middleware

import (
	"log"
	"net/http"
	"profile-template/internal/auth"
	"profile-template/internal/contextapp"
	"profile-template/internal/transport"
)

func NewAuthMiddleware(authUC auth.UsecaseContract) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie(transport.CookieSessionName)
			var user auth.TokenEntity
			if err == nil {

				user, err = authUC.Verify(r.Context(), cookie.Value)
				if err != nil {

				}
			}

			log.Print(user)

			next.ServeHTTP(w, r.WithContext(contextapp.WithAuth(r.Context(), contextapp.SessionContext{ID: user.UserID, Email: user.Email, Version: user.Version, Role: user.Role})))
		})
	}
}
