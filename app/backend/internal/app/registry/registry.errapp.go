package registry

import (
	"profile-template/internal/app/err"
	"profile-template/internal/article"
	"profile-template/internal/assets"
	"profile-template/internal/auth"
	"profile-template/internal/profile"
	"slices"
)

type RegisterErr struct {
	AuthKey    err.ErrDomain
	AssetKey   err.ErrDomain
	ProfileKey err.ErrDomain
	ArticleKey err.ErrDomain
}

func NewErrRegistry(keys RegisterErr) *err.ErrApp {
	errApp := err.NewErrApp()
	// Auth
	authErr := []err.ErrEntry{
		{Err: auth.ErrUserNotFound, Code: auth.CodeUserNotFound, StatusCode: 404, Message: "User not found"},
		{Err: auth.ErrInvalidCredentials, Code: auth.CodeInvalidCredentials, StatusCode: 401, Message: "Invalid credentials"},
		{Err: auth.ErrPasswordShort, Code: auth.CodePasswordShort, StatusCode: 400, Message: "Password is too short"},
		{Err: auth.ErrEmailAlreadyExists, Code: auth.CodeEmailAlreadyExists, StatusCode: 400, Message: "Email already exists"},
		{Err: auth.ErrInvalidEmail, Code: auth.CodeInvalidEmail, StatusCode: 400, Message: "Invalid email"},
		{Err: auth.ErrUnauthorized, Code: auth.CodeUnauthorized, StatusCode: 401, Message: "Unauthorized"},
		{Err: auth.ErrTokenInvalid, Code: auth.CodeTokenInvalid, StatusCode: 401, Message: "Invalid token"},
		{Err: auth.ErrTokenExpired, Code: auth.CodeTokenExpired, StatusCode: 401, Message: "Token expired"},
		{Err: auth.ErrRoleInvalid, Code: auth.CodeRoleInvalid, StatusCode: 403, Message: "Invalid role"},
	}
	errApp.Register(keys.AuthKey, authErr)

	errApp.Register(keys.AssetKey, slices.Concat([]err.ErrEntry{
		{Err: assets.ErrAssetNotFound, Code: assets.CodeAssetNotFound, StatusCode: 404, Message: "Asset not found"},
		{Err: assets.ErrAssetAlreadyExists, Code: assets.CodeAssetAlreadyExists, StatusCode: 400, Message: "Asset already exists"},
		{Err: assets.ErrAssetInvalidMime, Code: assets.CodeAssetInvalidMime, StatusCode: 400, Message: "Invalid mime"},
		{Err: assets.ErrAssetFailedCreate, Code: assets.CodeAssetFailedCreate, StatusCode: 500, Message: "Asset failed create"},
		{Err: assets.ErrInvalidAssetInput, Code: assets.CodeInvalidInput, StatusCode: 400, Message: "Invalid asset input"},
		{Err: assets.ErrAssetTooLarge, Code: assets.CodeAssetTooLarge, StatusCode: 413, Message: "Asset too large"},
		{Err: assets.ErrUnsupportedFormat, Code: assets.CodeUnsupportedFormat, StatusCode: 415, Message: "Unsupported format"},
		{Err: assets.ErrStorageUnavailable, Code: assets.CodeStorageUnavailable, StatusCode: 503, Message: "Storage unavailable"},
		{Err: assets.ErrPermissionDenied, Code: assets.CodePermissionDenied, StatusCode: 403, Message: "Permission denied"},
		{Err: assets.ErrUnauthorized, Code: assets.CodeUnauthorized, StatusCode: 401, Message: "Unauthorized"},
		{Err: assets.ErrAssetFailedUpdate, Code: assets.CodeAssetFailedUpdate, StatusCode: 500, Message: "Asset failed update"},
		{Err: assets.ErrAssetFailedDelete, Code: assets.CodeAssetFailedDelete, StatusCode: 500, Message: "Asset failed delete"},
		{Err: assets.ErrAssetInvalidPath, Code: assets.CodeAssetInvalidPath, StatusCode: 400, Message: "Asset invalid path"},
	}, authErr))
	errApp.Register(keys.ProfileKey, slices.Concat([]err.ErrEntry{
		{Err: profile.ErrNotFound, Code: profile.CodeUserNotFound, StatusCode: 404, Message: "User not found"},
		{Err: profile.ErrTagAlreadyAssigned, Code: profile.CodeTagAlreadyAssigned, StatusCode: 400, Message: "Tag already assigned"},
		{Err: profile.ErrInvalidUsername, Code: profile.CodeInvalidUsername, StatusCode: 400, Message: "Invalid username"},
		{Err: profile.ErrInvalidInput, Code: profile.CodeInvalidInput, StatusCode: 400, Message: "Invalid input"},
		{Err: profile.ErrInvalidAvatar, Code: profile.CodeInvalidAvatar, StatusCode: 400, Message: "Invalid avatar"},
		{Err: auth.ErrUnauthorized, Code: auth.CodeUnauthorized, StatusCode: 401, Message: "Unauthorized"},
		{Err: auth.ErrTokenInvalid, Code: auth.CodeTokenInvalid, StatusCode: 401, Message: "Invalid token"},
		{Err: auth.ErrTokenExpired, Code: auth.CodeTokenExpired, StatusCode: 401, Message: "Token expired"},
	}, authErr))
	errApp.Register(keys.ArticleKey, slices.Concat([]err.ErrEntry{
		{Err: article.ErrNotFound, Code: article.CodeNotFound, StatusCode: 404, Message: "Article not found"},
		{Err: article.ErrInvalidID, Code: article.CodeInvalidID, StatusCode: 400, Message: "Article ID is required"},
		{Err: article.ErrDuplicateSlug, Code: article.CodeDuplicate, StatusCode: 400, Message: "Article slug already exists"},
		{Err: article.ErrNothingToUpdate, Code: article.CodeNothingToUpdate, StatusCode: 400, Message: "Nothing to update"},
		{Err: article.ErrUnauthenticated, Code: article.CodeUnauthenticated, StatusCode: 401, Message: "Author ID is required"},
		{Err: article.ErrForbidden, Code: article.CodeForbidden, StatusCode: 403, Message: "Not the owner of this article"},
		{Err: article.ErrInvalidInput, Code: article.CodeInvalidInput, StatusCode: 400, Message: "Invalid input"},
		{Err: article.ErrBadRequest, Code: article.CodeBadRequest, StatusCode: 400, Message: "Bad Request"},
	}, authErr))
	return errApp
}
