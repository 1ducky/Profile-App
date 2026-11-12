package assets

import (
	"context"
	"errors"
	"io"
	"log"
	"profile-template/internal/config"
	"profile-template/internal/pipeline"
	"profile-template/internal/reader"
	"profile-template/internal/utils"

	"github.com/google/uuid"
)

type Usecase struct {
	StorageRepo RepositoryContract
	AssetRepo   AssetRepository
	conf        config.StorageConfig
}

func NewUsecase(config config.StorageConfig, StorageRepo RepositoryContract, AssetRepo AssetRepository) UsecaseContract {
	return &Usecase{StorageRepo: StorageRepo, AssetRepo: AssetRepo, conf: config}
}

func (u *Usecase) Upload(ctx context.Context, r io.Reader, upload UploadPolicy) (AssetMetaData, error) {
	var entity AssetMetaData

	// Policy
	policy := PolicyAsset[upload.Category]
	maxSize := min(policy.MaxSize, upload.MaxSize)
	mime, newR, err := reader.DetectMime(r)
	if err != nil {
		return AssetMetaData{}, err
	}

	// Reader handle
	limitedR := u.limitReader(newR, maxSize)
	typeBuffer, ok := isAllowedExt(Mime(mime))
	if !ok || !isAllowedMimeByPolicy(policy, typeBuffer.Mime) {
		return AssetMetaData{}, ErrAssetInvalidMime
	}

	// Resolve path
	entity.ID = uuid.NewString()
	entity.AuthorID = upload.UserID
	entity.Mime = typeBuffer.Mime
	entity.Category = upload.Category
	entity.Filename = upload.UserID + "_" + uuid.NewString() + string(typeBuffer.Ext)
	entity.Size = 0
	entity.Status = AssetStatusPending
	// entity.CreatedAt = upload.UserID
	// entity.UpdatedAt = upload.UserID
	var fullpath string
	fullpath, err = u.resolvePath(entity.Category, entity.Filename)
	if err != nil {
		return AssetMetaData{}, err
	}
	entity.FileKey = fullpath
	// Record
	err = u.AssetRepo.Record(ctx, RecordPayload{ID: entity.ID, Status: entity.Status, FileKey: entity.FileKey, Filename: entity.Filename, Mime: entity.Mime, AuthorID: entity.AuthorID, Category: entity.Category})
	if err != nil {
		return AssetMetaData{}, err
	}
	res, err := u.StorageRepo.Write(ctx, limitedR, fullpath)
	if err != nil {
		return AssetMetaData{}, err
	}
	// log.Printf("res size %d , max size %d, user:%s", res.Size, maxSize, upload.UserID)
	if res.Size > maxSize {
		return AssetMetaData{}, ErrAssetTooLarge
	}
	entity.Size = res.Size
	err = u.AssetRepo.Update(ctx, UpdatePayload{
		Id:        entity.ID,
		Size:      entity.Size,
		Status:    AssetStatusActive,
		OldStatus: entity.Status,
		AuthorID:  upload.UserID,
	})
	if err != nil {
		return AssetMetaData{}, err
	}
	entity.Status = AssetStatusActive
	return entity, nil
}

func (u *Usecase) Delete(ctx context.Context, payload DeletePayload) error {
	meta, err := u.AssetRepo.GetByIDs(ctx, payload.IDs)
	if err != nil {
		return err
	}
	validMeta := utils.MapField(meta, func(amd AssetMetaData) (AssetMetaData, bool) {
		if amd.Status == AssetStatusDeleted || amd.ParentID != "" {
			return amd, false
		}
		return amd, true
	})
	if len(validMeta) == 0 {
		return ErrAssetNotFound
	}
	workerCount := pipeline.CalculateWorkerCount(len(validMeta), pipeline.WorkerPoolConfig{
		MaxWorkers:    8,
		MinWorkers:    2,
		JobsPerWorker: 5,
	})
	pool := pipeline.NewWorkerPool[AssetMetaData, DeleteReaport](workerCount)
	jobs := pipeline.ProduceJob(ctx, validMeta, workerCount*2)
	reports := pool.Run(ctx, jobs, func(ctx context.Context, amd AssetMetaData) DeleteReaport {

		if amd.FileKey == "" {
			return DeleteReaport{id: amd.ID, err: ErrAssetInvalidPath}
		}
		delErr := u.StorageRepo.Delete(ctx, amd.FileKey)
		if delErr != nil {
			return DeleteReaport{id: amd.ID, err: delErr}
		}
		return DeleteReaport{id: amd.ID, err: nil}
	})
	res := pipeline.Collect(ctx, reports, len(meta))

	ids := utils.MapField(res, func(dr DeleteReaport) (string, bool) {
		if dr.err != nil && !errors.Is(dr.err, ErrAssetNotFound) {
			return "", false
		}
		return dr.id, true
	}) //get deleted ids successfull
	err = u.AssetRepo.MarkAsDeleted(ctx, DeletePayload{IDs: ids}) //mark as deleted
	if err != nil {
		return err
	}
	return nil
}
func (u *Usecase) Read(ctx context.Context, url string) (io.ReadCloser, error) {
	if url == "" {
		log.Print(url)
		return nil, ErrAssetNotFound
	}
	res, err := u.StorageRepo.Read(ctx, url)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *Usecase) LinkingAsset(ctx context.Context, assetIDs []string, parentID string) error {
	md, err := u.AssetRepo.GetByIDs(ctx, assetIDs)
	if err != nil {
		return err
	}
	validIds := utils.MapField(md, func(amd AssetMetaData) (string, bool) {
		if amd.Status == AssetStatusDeleted || amd.ParentID != "" {
			return "", false
		}
		return amd.ID, true
	})
	if len(validIds) == 0 {
		return ErrAssetNotFound
	}
	err = u.AssetRepo.LinkingAsset(ctx, validIds, parentID)
	if err != nil {
		return err
	}
	return nil
}
func (u *Usecase) LinkingAssetByFileKeys(ctx context.Context, fileKeys []string, parentID string) error {
	md, err := u.AssetRepo.GetByFileKeys(ctx, fileKeys)
	if err != nil {
		return err
	}
	validIds := utils.MapField(md, func(amd AssetMetaData) (string, bool) {
		if amd.Status == AssetStatusDeleted || amd.ParentID != "" {
			return "", false
		}
		return amd.ID, true
	})
	if len(validIds) == 0 {
		return ErrAssetNotFound
	}
	err = u.AssetRepo.LinkingAsset(ctx, validIds, parentID)
	if err != nil {
		return err
	}
	return nil
}

func (u *Usecase) ReadByParentID(ctx context.Context, parentIDs []string) ([]AssetMetaData, error) {
	data, err := u.AssetRepo.ReadByParentID(ctx, parentIDs)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, ErrAssetNotFound
	}
	return data, nil
}

func (u *Usecase) ReleaseByParentID(ctx context.Context, parentID []string) error {
	validParentID := utils.MapField(parentID, func(s string) (string, bool) {
		if s != "" {
			return s, true
		}
		return "", false
	})
	if len(validParentID) == 0 {
		return ErrAssetNotFound
	}
	err := u.AssetRepo.ReleaseByParentID(ctx, validParentID)
	if err != nil {
		return err
	}
	return nil
}
func (u *Usecase) ReleaseByFileKeys(ctx context.Context, fileKeys []string) error {
	validFileKeys := utils.MapField(fileKeys, func(s string) (string, bool) {
		if s != "" {
			return s, true
		}
		return "", false
	})
	if len(validFileKeys) == 0 {
		return ErrAssetNotFound
	}
	err := u.AssetRepo.ReleaseByFileKeys(ctx, validFileKeys)
	if err != nil {
		return err
	}
	return nil
}
