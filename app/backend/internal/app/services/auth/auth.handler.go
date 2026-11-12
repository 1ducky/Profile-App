package auth

import (
	"encoding/json"
	"net/http"
	"profile-template/internal/auth"
	"profile-template/internal/contextapp"
	"profile-template/internal/transport"
)

type Handler struct {
	AuthApp AuthApp
}

func NewHandler(authService AuthApp) Handler {
	return Handler{
		AuthApp: authService,
	}
}

func (h Handler) Login(w http.ResponseWriter, r *http.Request) error {
	var payload LoginRequest

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		return err
	}

	res, err := h.AuthApp.Login(r.Context(), payload)
	if err != nil {
		return err
	}

	cookie := http.Cookie{
		Name:     transport.CookieSessionName,
		Value:    res.token.AccessToken,           // The unencrypted data
		Path:     transport.CookieSessionPath,     // Accessible across the entire site
		MaxAge:   transport.CookieSessionMaxAge,   // Cookie expires in 1 hour (in seconds)
		HttpOnly: transport.CookieSessionHTTPOnly, // Prevents JavaScript access
		Secure:   transport.CookieSessionSecure,   // Recommended: Only sent over HTTPS
		SameSite: http.SameSiteLaxMode,            // Recommended: Protects against CSRF
	}

	// Send the cookie to the client in the response header
	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)

	return json.NewEncoder(w).Encode(transport.Response[LoginSession]{Status: transport.StatusSuccess, Data: res.Profile})
}

func (h Handler) Register(w http.ResponseWriter, r *http.Request) error {
	var payload RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		return err
	}

	err = h.AuthApp.Register(r.Context(), payload)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)

	return json.NewEncoder(w).Encode(transport.Response[map[string]string]{Status: transport.StatusSuccess, Data: map[string]string{"message": "User created successfully"}})
}

func (h Handler) GetSession(w http.ResponseWriter, r *http.Request) error {
	user, ok := contextapp.GetAuth(r.Context())
	if !ok {
		return auth.ErrTokenInvalid
	}
	if user.ID == "" {
		return auth.ErrUnauthorized
	}

	return json.NewEncoder(w).Encode(transport.Response[SessionResource]{Status: transport.StatusSuccess, Data: SessionResource{UserID: user.ID, Email: user.Email, Role: string(user.Role), Version: user.Version}})
}

func (h Handler) UpdateProfile(w http.ResponseWriter, r *http.Request) error {
	var payload UpdateProfileRequest

	cookie, err := r.Cookie(transport.CookieSessionName)
	if err != nil {
		return auth.ErrTokenInvalid
	}

	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		return err
	}

	err = h.AuthApp.UpdateProfile(r.Context(), cookie.Value, payload)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)

	return json.NewEncoder(w).Encode(transport.Response[map[string]string]{Status: transport.StatusSuccess, Data: map[string]string{"message": "Profile updated successfully"}})
}

func (h Handler) GetProfile(w http.ResponseWriter, r *http.Request) error {

	cookie, err := r.Cookie(transport.CookieSessionName)
	if err != nil {
		return auth.ErrTokenInvalid
	}

	res, err := h.AuthApp.GetProfile(r.Context(), cookie.Value)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)

	return json.NewEncoder(w).Encode(transport.Response[ProfileResource]{Status: transport.StatusSuccess, Data: res})
}
