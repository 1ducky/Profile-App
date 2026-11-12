package article

import (
	"errors"
	"profile-template/internal/app/err"
)

var (
	ErrNotFound        = errors.New("article: not found")
	ErrInvalidID       = errors.New("article: id is required")
	ErrDuplicateSlug   = errors.New("article: slug already exists")
	ErrNothingToUpdate = errors.New("article: nothing to update")
	ErrUnauthenticated = errors.New("article: author id is required")
	ErrForbidden       = errors.New("article: not the owner of this article")
	ErrInvalidInput    = errors.New("article: invalid input")
	ErrBadRequest      = errors.New("article: Bad Request")
)

const (
	CodeNotFound        err.ErrCode = "ARTICLE_NOT_FOUND"
	CodeInvalidID       err.ErrCode = "ARTICLE_INVALID_ID"
	CodeDuplicate       err.ErrCode = "ARTICLE_DUPLICATE"
	CodeNothingToUpdate err.ErrCode = "ARTICLE_NOTHING_TO_UPDATE"
	CodeUnauthenticated err.ErrCode = "ARTICLE_UNAUTHENTICATED"
	CodeForbidden       err.ErrCode = "ARTICLE_FORBIDDEN"
	CodeInvalidInput    err.ErrCode = "ARTICLE_INVALID_INPUT"
	CodeBadRequest      err.ErrCode = "ARTICLE_BAD_REQUEST"
)
