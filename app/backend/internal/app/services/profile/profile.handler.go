package profile

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"profile-template/internal/assets"
	"profile-template/internal/auth"
	"profile-template/internal/contextapp"
	"profile-template/internal/profile"
	"profile-template/internal/transport"
)

type Handler struct {
	app ProfileApp
}

func NewHandler(app ProfileApp) Handler {
	return Handler{app: app}
}

func (h *Handler) GetOwnProfile(w http.ResponseWriter, r *http.Request) error {
	log.Print("GetOwnProfile")
	res, err := h.app.GetOwnProfile(r.Context())
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)

	return json.NewEncoder(w).Encode(transport.Response[ProfileResource]{Status: transport.StatusSuccess, Data: res})
}

func (h *Handler) GetProfileByID(w http.ResponseWriter, r *http.Request) error {
	id := r.PathValue(DynamicIDPath)
	if id == "" {
		return profile.ErrNotFound
	}
	res, err := h.app.GetProfileByID(r.Context(), id)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)

	return json.NewEncoder(w).Encode(transport.Response[ProfileResource]{Status: transport.StatusSuccess, Data: res})
}

func (h *Handler) UpdateProfile(w http.ResponseWriter, r *http.Request) error {
	var payload UpdateProfileRequest
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		return err
	}

	err = h.app.UpdateProfile(r.Context(), payload)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)

	return json.NewEncoder(w).Encode(transport.Response[map[string]string]{Status: transport.StatusSuccess, Data: map[string]string{"message": "Profile updated successfully"}})
}

func (h *Handler) UploadAvatar(w http.ResponseWriter, r *http.Request) error {
	// 1. Batasi ukuran body (misal 10 MB) supaya tidak membebani memori
	user, ok := contextapp.GetAuth(r.Context())
	if !ok {
		return auth.ErrUnauthorized
	}
	r.Body = http.MaxBytesReader(w, r.Body, 3*assets.MB)

	// 2. Ambil file dari form field bernama "file"
	file, header, err := r.FormFile("avatar")
	if err != nil {
		http.Error(w, "file tidak ditemukan: "+err.Error(), http.StatusBadRequest)
		return profile.ErrInvalidAvatar
	}
	defer file.Close()

	// file adalah multipart.File, sudah memenuhi io.Reader
	// header berisi metadata: header.Filename, header.Size, header.Header
	fmt.Println(header.Filename, header.Size)

	// 3. Pakai sebagai io.Reader
	res, err := h.app.UploadAvatar(r.Context(), UploadAvatarRequest{id: user.ID, reader: file})
	if err != nil {
		return err
	}
	return json.NewEncoder(w).Encode(transport.Response[profileAvatarResource]{Status: transport.StatusSuccess, Data: res})

}
