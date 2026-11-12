package profile

import (
	"context"
	"profile-template/internal/db/mapper"
	"profile-template/internal/utils"
)

type Usecase struct {
	repo RepoContract
}

func NewUsecase(repo RepoContract) UsecaseContract {
	return &Usecase{repo: repo}
}

func (uc *Usecase) CreateUser(ctx context.Context, req CreateRequest) error {
	var CreateInput CreateData
	if req.ID == "" {
		return ErrInvalidUsername
	}
	CreateInput.ID = req.ID
	if req.Username == "" {
		return ErrInvalidUsername
	}
	CreateInput.Username = req.Username
	CreateInput.Tag = CreateInput.Username[0:4] + req.ID[0:4]
	err := uc.repo.CreateUser(ctx, CreateInput)
	if err != nil {
		transalte := mapper.MapMySQLError(err)
		if transalte == mapper.ErrDuplicate {
			return transalte
		}
	}
	return nil
}

func (uc *Usecase) GetUserByID(ctx context.Context, id string) (Profile, error) {
	res, err := uc.repo.GetUserByID(ctx, id)
	if err != nil {
		translate := mapper.MapMySQLError(err)
		if translate.Error() == mapper.ErrNoRows.Error() {
			return Profile{}, ErrNotFound
		}
		return Profile{}, translate
	}
	return res, nil
}
func (uc *Usecase) GetUsersByIDs(ctx context.Context, ids []string) ([]Profile, error) {
	uniqueIDs := utils.Unique(ids)
	res, err := uc.repo.GetUsersByIDs(ctx, uniqueIDs)
	if err != nil {
		translate := mapper.MapMySQLError(err)
		if translate.Error() == mapper.ErrNoRows.Error() {
			return []Profile{}, ErrNotFound
		}
		return []Profile{}, translate
	}
	return res, nil
}
func (uc *Usecase) UpdateProfile(ctx context.Context, req UpdateRequest) error {
	if req.Tag == "" || req.Username == "" || req.ID == "" {
		return ErrInvalidInput
	}
	if len(req.Tag) < 4 || len(req.Username) < 4 {
		return ErrInvalidInput
	}
	res := uc.repo.UpdateProfile(ctx, req)
	if res != nil {
		transalte := mapper.MapMySQLError(res)
		if transalte.Error() == mapper.ErrNoRows.Error() {
			return ErrNotFound
		}
		if transalte.Error() == mapper.ErrDuplicate.Error() {
			return ErrTagAlreadyAssigned
		}
		return transalte
	}
	return nil
}

func (uc *Usecase) DeleteUser(ctx context.Context, id string) error {
	return nil
}

func (uc *Usecase) UploadAvatar(ctx context.Context, id string, path string) error {
	if id == "" {
		return ErrNotFound
	}
	if path == "" {
		return ErrInvalidAvatar
	}
	return uc.repo.SetProfilePhoto(ctx, id, path)
}

func (uc *Usecase) UploadCover(ctx context.Context, id string) error {
	return nil
}
