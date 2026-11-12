package assets

import (
	"context"
)

type Repo struct {
	assets []AssetMetaData
}

func NewRepository() AssetRepository {
	return &Repo{}
}

func (r *Repo) Record(ctx context.Context, payload RecordPayload) error {
	return nil
}

func (r *Repo) MarkAsDeleted(ctx context.Context, payload DeletePayload) error {
	return nil
}

func (r *Repo) Update(ctx context.Context, payload UpdatePayload) error {
	return nil
}

func (r *Repo) GetByIDs(ctx context.Context, id []string) ([]AssetMetaData, error) {
	return []AssetMetaData{}, nil
}
func (r *Repo) GetByFileKeys(ctx context.Context, fileKeys []string) ([]AssetMetaData, error) {
	return []AssetMetaData{}, nil
}

func (r *Repo) ReadByParentID(ctx context.Context, parentID []string) ([]AssetMetaData, error) {
	return []AssetMetaData{}, nil
}

func (r *Repo) LinkingAsset(ctx context.Context, assetID []string, parentID string) error {
	return nil
}
func (r *Repo) LinkingAssetByFileKeys(ctx context.Context, fileKeys []string, parentID string) error {
	return nil
}

func (r *Repo) ReleaseByParentID(ctx context.Context, parentIDs []string) error {
	return nil
}
func (r *Repo) ReleaseByFileKeys(ctx context.Context, parentIDs []string) error {
	return nil
}
