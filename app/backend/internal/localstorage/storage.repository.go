package localstorage

import (
	"context"
	"io"
	"os"
	"path"
	"profile-template/internal/assets"
	"profile-template/internal/config"
)

type LocalStorage struct {
	conf config.StorageConfig
}

func NewLocalStorage(conf config.StorageConfig) (assets.RepositoryContract, error) {
	// Initial Folder
	dirs := []string{
		conf.StorageRoot,
		path.Join(conf.StorageRoot, conf.StoragePathTemp),
		path.Join(conf.StorageRoot, conf.StoragePathUpload),
		path.Join(conf.StorageRoot, conf.StoragePathDocument),
		path.Join(conf.StorageRoot, conf.StoragePathAvatar),
		path.Join(conf.StorageRoot, conf.StoragePathVideo),
	}
	_AttemtDir := 10

	var attempDir []string

	for _AttemtDir > 0 {
		for _, dir := range dirs {
			if err := os.MkdirAll(dir, os.ModePerm); err != nil {
				attempDir = append(attempDir, dir)
			}
		}
		dirs = attempDir
		attempDir = []string{}
		_AttemtDir--
	}

	if len(attempDir) > 0 {
		return nil, assets.ErrStorageUnavailable
	}

	return &LocalStorage{conf: conf}, nil
}

func (l *LocalStorage) Write(ctx context.Context, reader io.Reader, destination string) (assets.StoreResult, error) {
	finalDestination, err := l.resolveFilePath(destination)
	if err != nil {
		return assets.StoreResult{}, err
	}
	dst, err := os.Create(path.Join(finalDestination))
	if err != nil {
		return assets.StoreResult{}, assets.ErrAssetFailedCreate
	}
	defer dst.Close()

	size, err := io.Copy(dst, reader)
	if err != nil || size == 0 {
		dst.Close()
		if delErr := l.Delete(ctx, destination); delErr != nil {
			return assets.StoreResult{}, delErr
		}
		return assets.StoreResult{}, assets.ErrAssetFailedCreate
	}

	return assets.StoreResult{FileKey: finalDestination, Size: size}, nil
}

func (l *LocalStorage) Delete(ctx context.Context, destination string) error {
	finalDestination, err := l.resolveFilePath(destination)
	if err != nil {
		return assets.ErrAssetInvalidPath
	}
	err = os.Remove(finalDestination)
	if err != nil {
		if os.IsNotExist(err) {
			return assets.ErrAssetNotFound
		}
		return err
	}
	return nil
}

func (l *LocalStorage) Read(ctx context.Context, destination string) (io.ReadCloser, error) {
	finalDestination, err := l.resolveFilePath(destination)
	if err != nil {
		return nil, err
	}
	info, err := os.Stat(finalDestination)
	if err != nil && os.IsNotExist(err) {
		return nil, assets.ErrAssetNotFound
	}
	if err == nil && info.IsDir() {
		return nil, assets.ErrAssetInvalidPath
	}
	f, err := os.Open(finalDestination)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, assets.ErrAssetNotFound
		}
		return nil, err
	}
	return f, nil
}

func (l *LocalStorage) Move(ctx context.Context, oldDestination, newDestination string) (assets.StoreResult, error) {
	finalOldDestination, err := l.resolveFilePath(oldDestination)
	if err != nil {
		return assets.StoreResult{}, err
	}
	finalNewDestination, err := l.resolveFilePath(newDestination)
	if err != nil {
		return assets.StoreResult{}, err
	}
	err = os.Rename(finalOldDestination, finalNewDestination)
	if err != nil {
		if os.IsNotExist(err) {
			return assets.StoreResult{}, assets.ErrAssetNotFound
		}
		return assets.StoreResult{}, err
	}
	return assets.StoreResult{FileKey: finalNewDestination}, nil
}
