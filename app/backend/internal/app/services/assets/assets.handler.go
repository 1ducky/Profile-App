package assets

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"profile-template/internal/assets"
	"profile-template/internal/auth"
	"profile-template/internal/contextapp"
	"profile-template/internal/transport"
)

type Handler struct {
	AssetApp AssetApp
}

func NewHandler(assetsService AssetApp) Handler {
	return Handler{
		AssetApp: assetsService,
	}
}

func (h *Handler) ServeAssets(w http.ResponseWriter, r *http.Request) error {
	fmt.Println(r.PathValue(KEY))
	reader, err := h.AssetApp.ServePublicAssets(r.Context(), r.PathValue(KEY))
	if err != nil {
		return err
	}
	defer reader.Close()

	data, err := io.ReadAll(reader)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return err
	}

	w.Header().Set("Content-Type", http.DetectContentType(data))
	w.Write(data)

	return nil
}

func (h *Handler) UploadAvatar(w http.ResponseWriter, r *http.Request) error {
	user, ok := contextapp.GetAuth(r.Context())
	if !ok {
		return auth.ErrUnauthorized
	}
	r.Body = http.MaxBytesReader(w, r.Body, 2*assets.MB)
	fmt.Println("Limit Reader")
	file, header, err := r.FormFile("image")
	if err != nil {
		return assets.ErrAssetNotFound
	}
	defer file.Close()
	fmt.Println("Get Reader")

	// file adalah multipart.File, sudah memenuhi io.Reader
	// header berisi metadata: header.Filename, header.Size, header.Header
	fmt.Println("metadata", header.Filename, header.Size)

	// 3. Pakai sebagai io.Reader
	res, err := h.AssetApp.UploadPublicAssets(r.Context(), UploadPublicAssetsReq{UserID: user.ID, Reader: file})
	if err != nil {
		return err
	}
	return json.NewEncoder(w).Encode(transport.Response[AssetsUploadRes]{Status: transport.StatusSuccess, Data: res})

}
