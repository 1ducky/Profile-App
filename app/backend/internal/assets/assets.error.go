package assets

import (
	"errors"
	"profile-template/internal/app/err"
)

// Sentinel errors. Usecase & repository layer harus me-wrap error asli
var (
	ErrAssetNotFound      = errors.New("asset not found")
	ErrAssetAlreadyExists = errors.New("asset already exists")
	ErrAssetInvalidMime   = errors.New("Unsupport Mime")
	ErrAssetFailedCreate  = errors.New("asset failed create")
	ErrInvalidAssetInput  = errors.New("invalid asset input")
	ErrAssetTooLarge      = errors.New("asset size exceeds limit")
	ErrUnsupportedFormat  = errors.New("unsupported asset format")
	ErrStorageUnavailable = errors.New("storage backend unavailable")
	ErrPermissionDenied   = errors.New("permission denied")
	ErrUnauthorized       = errors.New("unauthorized")
	ErrAssetFailedUpdate  = errors.New("asset failed update")
	ErrAssetFailedDelete  = errors.New("asset failed delete")
	ErrAssetInvalidPath   = errors.New("asset invalid path")
)

// Error Code Contract Frontend
const (
	CodeAssetNotFound      err.ErrCode = "ASSET_NOT_FOUND"
	CodeAssetAlreadyExists err.ErrCode = "ASSET_ALREADY_EXISTS"
	CodeInvalidInput       err.ErrCode = "INVALID_INPUT"
	CodeAssetTooLarge      err.ErrCode = "ASSET_TOO_LARGE"
	CodeUnsupportedFormat  err.ErrCode = "UNSUPPORTED_FORMAT"
	CodeStorageUnavailable err.ErrCode = "STORAGE_UNAVAILABLE"
	CodePermissionDenied   err.ErrCode = "PERMISSION_DENIED"
	CodeUnauthorized       err.ErrCode = "UNAUTHORIZED"
	CodeAssetInvalidMime   err.ErrCode = "ASSET_INVALID_MIME"
	CodeInternal           err.ErrCode = "INTERNAL_ERROR"
	CodeAssetFailedDelete  err.ErrCode = "ASSET_FAILED_DELETE"
	CodeAssetFailedUpdate  err.ErrCode = "ASSET_FAILED_UPDATE"
	CodeAssetFailedCreate  err.ErrCode = "ASSET_FAILED_CREATE"
	CodeAssetInvalidPath   err.ErrCode = "ASSET_INVALID_PATH"
)
