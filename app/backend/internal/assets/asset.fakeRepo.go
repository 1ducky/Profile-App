package assets

import (
	"context"
	"errors"
	"slices"
	"sync"
	"time"
)

// FakeRepo is an in-memory implementation of AssetRepository for testing.
type FakeRepo struct {
	mu     sync.RWMutex
	assets map[string]AssetMetaData
}

// NewFakeRepo creates a new instance of FakeRepo.
func NewFakeRepo() AssetRepository {
	return &FakeRepo{
		assets: make(map[string]AssetMetaData),
	}
}
func (f *FakeRepo) Record(ctx context.Context, payload RecordPayload) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	asset := AssetMetaData{
		ID:        payload.ID,
		Filename:  payload.Filename,
		FileKey:   payload.FileKey,
		Mime:      payload.Mime,
		AuthorID:  payload.AuthorID,
		Status:    payload.Status,
		Category:  payload.Category,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	f.assets[payload.ID] = asset
	return nil
}
func (f *FakeRepo) MarkAsDeleted(ctx context.Context, payload DeletePayload) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	for _, id := range payload.IDs {
		asset, ok := f.assets[id]
		if !ok {
			return errors.New("asset not found")
		}
		asset.Status = AssetStatusDeleted
		asset.UpdatedAt = time.Now()
		f.assets[id] = asset
	}
	return nil
}
func (f *FakeRepo) Update(ctx context.Context, payload UpdatePayload) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	asset, ok := f.assets[payload.Id]
	if !ok {
		return errors.New("asset not found")
	}
	if payload.Size != 0 {
		asset.Size = payload.Size
	}
	if payload.Status != "" {
		asset.Status = payload.Status
	}
	asset.UpdatedAt = time.Now()
	f.assets[payload.Id] = asset
	return nil
}
func (f *FakeRepo) GetByIDs(ctx context.Context, ids []string) ([]AssetMetaData, error) {
	f.mu.RLock()
	defer f.mu.RUnlock()
	var result []AssetMetaData
	for _, id := range ids {
		if asset, ok := f.assets[id]; ok {
			result = append(result, asset)
		}
	}
	return result, nil
}
func (f *FakeRepo) GetByFileKeys(ctx context.Context, fileKeys []string) ([]AssetMetaData, error) {
	f.mu.RLock()
	defer f.mu.RUnlock()
	var result []AssetMetaData
	for _, fileKey := range fileKeys {
		for _, asset := range f.assets {
			if asset.FileKey == fileKey {
				result = append(result, asset)
			}
		}
	}
	return result, nil
}

func (f *FakeRepo) ReadByParentID(ctx context.Context, parentIDs []string) ([]AssetMetaData, error) {
	f.mu.RLock()
	defer f.mu.RUnlock()
	var result []AssetMetaData
	for _, id := range parentIDs {
		for _, asset := range f.assets {
			if asset.ParentID == id {
				result = append(result, asset)
			}
		}
	}
	return result, nil
}
func (f *FakeRepo) LinkingAsset(ctx context.Context, assetIDs []string, parentID string) error {
	if parentID == "" {
		return ErrAssetNotFound
	}
	f.mu.Lock()
	defer f.mu.Unlock()
	for _, id := range assetIDs {
		if asset, ok := f.assets[id]; ok {
			asset.ParentID = parentID
			asset.UpdatedAt = time.Now()
			f.assets[id] = asset

		}
	}
	return nil
}

func (f *FakeRepo) LinkingAssetByFileKeys(ctx context.Context, fileKeys []string, parentID string) error {
	if parentID == "" {
		return ErrAssetNotFound
	}
	f.mu.Lock()
	defer f.mu.Unlock()
	for _, fileKey := range fileKeys {
		for _, asset := range f.assets {
			if asset.FileKey == fileKey {
				asset.ParentID = parentID
				asset.UpdatedAt = time.Now()
				f.assets[asset.ID] = asset
			}
		}
	}
	return nil
}

func (f *FakeRepo) ReleaseByParentID(ctx context.Context, parentIDs []string) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	var assetIDs []string
	for _, asset := range f.assets {
		if slices.Contains(parentIDs, asset.ParentID) {
			asset.ParentID = ""
			asset.UpdatedAt = time.Now()
			f.assets[asset.ID] = asset
			assetIDs = append(assetIDs, asset.ID)
		}
	}
	if len(assetIDs) == 0 {
		return ErrAssetNotFound
	}
	return nil
}
func (f *FakeRepo) ReleaseByFileKeys(ctx context.Context, fileKeys []string) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	var assetIDs []string
	for _, asset := range f.assets {
		if slices.Contains(fileKeys, asset.FileKey) {
			asset.ParentID = ""
			asset.UpdatedAt = time.Now()
			f.assets[asset.ID] = asset
			assetIDs = append(assetIDs, asset.ID)
		}
	}
	if len(assetIDs) == 0 {
		return ErrAssetNotFound
	}
	return nil
}
