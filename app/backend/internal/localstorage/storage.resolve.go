package localstorage

import (
	"path/filepath"
	"profile-template/internal/assets"
	"strings"
)

// Resolve Path Option end Return valid path
func (l *LocalStorage) resolveFilePath(destination string) (string, error) {
	validDestination := map[string]string{
		l.conf.StoragePathAvatar:   l.conf.StoragePathAvatar,
		l.conf.StoragePathDocument: l.conf.StoragePathDocument,
		l.conf.StoragePathVideo:    l.conf.StoragePathVideo,
		l.conf.StoragePathTemp:     l.conf.StoragePathTemp,
		l.conf.StoragePathUpload:   l.conf.StoragePathUpload,
	}

	if destination == "" {
		return "", assets.ErrAssetInvalidPath
	}
	part := strings.Split(filepath.ToSlash(destination), "/")
	if len(part) != 2 || validDestination[part[0]] == "" {

		return "", assets.ErrAssetInvalidPath
	}

	root, err := filepath.Abs(l.conf.StorageRoot)
	if err != nil {
		return "", err
	}

	fullPath, err := filepath.Abs(filepath.Join(root, destination))
	if err != nil {
		return "", err
	}

	if fullPath == root {
		return "", assets.ErrAssetInvalidPath
	}

	rel, err := filepath.Rel(root, fullPath)
	if err != nil {
		return "", err
	}

	if rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
		return "", assets.ErrAssetInvalidPath
	}

	return fullPath, nil
}
