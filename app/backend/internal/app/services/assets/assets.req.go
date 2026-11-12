package assets

import "io"

type UploadPublicAssetsReq struct {
	UserID string
	Reader io.Reader
}
