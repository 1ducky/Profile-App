package assets

import "profile-template/internal/assets"

type AssetsUploadRes struct {
	FileKey  string               `json:"filekey"`
	Metadata assets.AssetMetaData `json:"metadata"`
}
