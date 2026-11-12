package assets

import (
	"io"
	"path/filepath"
)

func (l *Usecase) resolvePath(dir AssetCategory, fileName string) (string, error) {
	switch dir {
	case CategoryPost:
		return filepath.Join(l.conf.StoragePathUpload, fileName), nil
	case CategoryTemp:
		return filepath.Join(l.conf.StoragePathTemp, fileName), nil
	case CategoryProfile:
		return filepath.Join(l.conf.StoragePathAvatar, fileName), nil
	case CategoryDocument:
		return filepath.Join(l.conf.StoragePathDocument, fileName), nil
	case CategoryVideo:
		return filepath.Join(l.conf.StoragePathVideo, fileName), nil
	default:
		return "", ErrStorageUnavailable
	}
}

func (l *Usecase) limitReader(reader io.Reader, maxSize int64) io.Reader {
	if maxSize+1 == 0 {
		return reader
	}
	return io.LimitReader(reader, maxSize+1)
}
