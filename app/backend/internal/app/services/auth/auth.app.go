package auth

import (
	"context"
	"profile-template/internal/app/uow"
	"profile-template/internal/auth"
	"profile-template/internal/profile"
	"strings"

	"github.com/google/uuid"
)

type TXUsecase struct {
	Auth    auth.UsecaseContract
	Profile profile.UsecaseContract
}

type AuthApp struct {
	uow     uow.UnitOfWork[TXUsecase]
	auth    auth.UsecaseContract
	profile profile.UsecaseContract
}

func NewAuthApp(authUsecase auth.UsecaseContract, profileUsecase profile.UsecaseContract, uow uow.UnitOfWork[TXUsecase]) *AuthApp {
	return &AuthApp{
		auth:    authUsecase,
		profile: profileUsecase,
		uow:     uow,
	}
}

// session auth
func (a *AuthApp) Login(ctx context.Context, req LoginRequest) (LoginResource, error) {
	token, err := a.auth.Login(ctx, auth.LoginPayload{Email: req.Email, Password: req.Password})
	if err != nil {
		return LoginResource{}, err
	}
	user, err := a.auth.Verify(ctx, token.AccessToken)
	if err != nil {
		return LoginResource{}, err
	}
	pro, err := a.profile.GetUserByID(ctx, user.UserID)
	if err != nil {
		return LoginResource{}, err
	}
	return LoginResource{token: token, Profile: LoginSession{SessionResource{UserID: user.UserID, Role: string(user.Role), Email: user.Email, Version: user.Version}, ProfileResource{Username: pro.Username, Tag: pro.Tag, Avatar: pro.AvatarURL, Bio: pro.Bio}}}, nil
}
func (a *AuthApp) Register(ctx context.Context, req RegisterRequest) error {
	if req.Email == "" || req.Password == "" || req.Username == "" {
		return auth.ErrInvalidCredentials
	}
	err := a.uow.Do(ctx, func(ctx context.Context, r TXUsecase) error {
		userID := uuid.NewString()
		err := r.Auth.Register(ctx, auth.RegisterPayload{UserID: userID, Email: req.Email, Password: req.Password})
		if err != nil {
			return err
		}
		err = r.Profile.CreateUser(ctx, profile.CreateRequest{ID: userID, Username: req.Username, Tag: strings.Split(req.Username, "")[0] + userID[0:5]})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (a *AuthApp) UpdateProfile(ctx context.Context, token string, req UpdateProfileRequest) error {
	user, err := a.auth.Verify(ctx, token)
	if err != nil {
		return err
	}
	err = a.profile.UpdateProfile(ctx, profile.UpdateRequest{ID: user.UserID, Username: req.Username, Tag: req.Tag, Bio: req.Bio})
	if err != nil {
		return err
	}
	return nil
}

func (a *AuthApp) GetProfile(ctx context.Context, token string) (ProfileResource, error) {

	user, err := a.auth.Verify(ctx, token)
	if err != nil {
		return ProfileResource{}, err
	}
	res, err := a.profile.GetUserByID(ctx, user.UserID)
	if err != nil {
		return ProfileResource{}, err
	}

	return ProfileResource{Username: res.Username, Tag: res.Tag, Avatar: res.AvatarURL, Bio: res.Bio}, nil
}

func (a *AuthApp) GetSession(ctx context.Context, token string) (SessionResource, error) {

	user, err := a.auth.Verify(ctx, token)
	if err != nil {
		return SessionResource{}, err
	}
	res, err := a.profile.GetUserByID(ctx, user.UserID)
	if err != nil {
		return SessionResource{}, err
	}

	return SessionResource{UserID: res.ID, Role: string(user.Role), Email: user.Email, Version: user.Version}, nil
}
