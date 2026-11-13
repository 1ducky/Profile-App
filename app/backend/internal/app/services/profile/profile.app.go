package profile

import (
	"context"
	"fmt"
	"path/filepath"
	assetApp "profile-template/internal/app/services/assets"
	"profile-template/internal/app/uow"
	"profile-template/internal/assets"
	"profile-template/internal/contextapp"
	"profile-template/internal/profile"
	"strings"
)

type TXUsecase struct {
	Profile profile.UsecaseContract
	Asset   assets.UsecaseContract
}

type ProfileApp struct {
	uow  uow.UnitOfWork[TXUsecase]
	prof profile.UsecaseContract
}

func NewProfileApp(prof profile.UsecaseContract, uow uow.UnitOfWork[TXUsecase]) *ProfileApp {
	return &ProfileApp{prof: prof, uow: uow}
}

func (p *ProfileApp) GetOwnProfile(ctx context.Context) (ProfileResource, error) {
	identity, ok := contextapp.GetAuth(ctx)
	if !ok {
		return ProfileResource{}, profile.ErrNotFound
	}
	pro, err := p.prof.GetUserByID(ctx, identity.ID)
	if err != nil {
		return ProfileResource{}, err
	}
	return ProfileResource{Username: pro.Username, Tag: pro.Tag, Avatar: pro.AvatarURL, Bio: pro.Bio}, nil
}

func (p *ProfileApp) GetProfileByID(ctx context.Context, id string) (ProfileResource, error) {
	pro, err := p.prof.GetUserByID(ctx, id)
	if err != nil {
		return ProfileResource{}, err
	}
	return ProfileResource{Username: pro.Username, Tag: pro.Tag, Avatar: pro.AvatarURL, Bio: pro.Bio}, nil
}

func (p *ProfileApp) UpdateProfile(ctx context.Context, req UpdateProfileRequest) error {
	user, ok := contextapp.GetAuth(ctx)
	if !ok {
		return profile.ErrNotFound
	}
	err := p.prof.UpdateProfile(ctx, profile.UpdateRequest{ID: user.ID, Username: req.Username, Tag: req.Tag, Bio: req.Bio})
	if err != nil {
		return err
	}
	return nil
}

func (p *ProfileApp) UploadAvatar(ctx context.Context, userID string, filekey string) (profileAvatarResource, error) {
	cleanFilekey := p.clearToValidFileKey(filekey)
	if userID == "" || cleanFilekey == "" {
		return profileAvatarResource{}, profile.ErrNotFound
	}
	fmt.Println("app: " + cleanFilekey)
	var AssetInfo assets.AssetMetaData
	err := p.uow.Do(ctx, func(ctx context.Context, r TXUsecase) error {
		err := r.Profile.UploadAvatar(ctx, userID, cleanFilekey)
		if err != nil {
			return err
		}
		_ = r.Asset.ReleaseByParentID(ctx, []string{userID})
		fmt.Println("Release Assets")
		err = r.Asset.LinkingAssetByFileKeys(ctx, []string{filekey}, userID)
		return err
	})

	AssetInfo.FileKey = filepath.ToSlash(AssetInfo.FileKey)
	fmt.Print("after Commit")
	fmt.Print(AssetInfo)

	return profileAvatarResource{Metadata: AssetInfo, Avatar: AssetInfo.FileKey}, err
}

func (a *ProfileApp) clearToValidFileKey(filekey string) string {
	return filepath.FromSlash(strings.TrimPrefix(filekey, assetApp.PREFIX))
}
