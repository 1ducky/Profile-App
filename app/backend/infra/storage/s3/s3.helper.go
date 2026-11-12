package s3

import (
	"path"
	"profile-template/internal/assets"
	"strings"
)

func resolveObjectKey(key string) (string, error) {
	if key == "" {
		return "", assets.ErrAssetInvalidPath
	}

	// Normalize Windows separator to S3 separator.
	key = strings.ReplaceAll(key, "\\", "/")

	// Reject absolute paths.
	if strings.HasPrefix(key, "/") {
		return "", assets.ErrAssetInvalidPath
	}

	// Reject traversal segments.
	for _, segment := range strings.Split(key, "/") {
		if segment == "" || segment == "." || segment == ".." {
			return "", assets.ErrAssetInvalidPath
		}
	}

	return path.Clean(key), nil
}
