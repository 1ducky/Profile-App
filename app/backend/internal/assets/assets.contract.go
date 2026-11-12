package assets

import (
	"context"
	"io"
)

type RepositoryContract interface {
	Write(ctx context.Context, reader io.Reader, path string) (StoreResult, error)
	Delete(ctx context.Context, path string) error
	Read(ctx context.Context, path string) (io.ReadCloser, error)
}

type UsecaseContract interface {
	Upload(ctx context.Context, reader io.Reader, upload UploadPolicy) (AssetMetaData, error)
	Delete(ctx context.Context, payload DeletePayload) error
	Read(ctx context.Context, url string) (io.ReadCloser, error)
	LinkingAsset(ctx context.Context, assetID []string, parentID string) error
	LinkingAssetByFileKeys(ctx context.Context, fileKeys []string, parentID string) error
	ReadByParentID(ctx context.Context, parentID []string) ([]AssetMetaData, error)
	ReleaseByParentID(ctx context.Context, parentIDs []string) error
	ReleaseByFileKeys(ctx context.Context, fileKeys []string) error
}

type AssetRepository interface {
	Record(ctx context.Context, payload RecordPayload) error
	MarkAsDeleted(ctx context.Context, payload DeletePayload) error
	Update(ctx context.Context, payload UpdatePayload) error
	GetByIDs(ctx context.Context, id []string) ([]AssetMetaData, error)
	GetByFileKeys(ctx context.Context, fileKeys []string) ([]AssetMetaData, error)
	ReadByParentID(ctx context.Context, parentID []string) ([]AssetMetaData, error)
	LinkingAsset(ctx context.Context, assetID []string, parentID string) error
	LinkingAssetByFileKeys(ctx context.Context, fileKeys []string, parentID string) error
	ReleaseByParentID(ctx context.Context, parentIDs []string) error
	ReleaseByFileKeys(ctx context.Context, FileKey []string) error
}
