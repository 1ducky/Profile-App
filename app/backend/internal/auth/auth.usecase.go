package auth

import (
	"context"
	"profile-template/internal/db/mapper"
	"profile-template/internal/utils"
	"strconv"
	"strings"
	"time"
)

type Service struct {
	CredentialRepository CredentialRepositoryContract
	Token                TokenContract
	SessionRepo          SessionRepositoryContract
}

func NewAuthService(CredentialRepo CredentialRepositoryContract, Token TokenContract, Session SessionRepositoryContract) UsecaseContract {
	return &Service{
		CredentialRepository: CredentialRepo,
		Token:                Token,
		SessionRepo:          Session,
	}
}

func (s *Service) Login(ctx context.Context, payload LoginPayload) (TokenResponse, error) {
	if payload.Email == "" || payload.Password == "" || !utils.IsEmailValid(payload.Email) || len(strings.Split(payload.Password, "")) < PasswordLength {
		return TokenResponse{}, ErrInvalidCredentials
	}
	credentials, err := s.CredentialRepository.Authentication(ctx, payload.Email)
	if err != nil {
		translate := mapper.MapMySQLError(err)
		if translate.Error() == mapper.ErrNoRows.Error() {
			return TokenResponse{}, ErrInvalidCredentials
		}
		return TokenResponse{}, translate
	}
	// Compare hash password
	if !utils.CheckHashString(payload.Password, credentials.Password) {
		return TokenResponse{}, ErrInvalidCredentials
	}

	accessTokenRevokedAt := time.Now().Add(AccessTokenDuration)
	refreshTokenRevokedAt := time.Now().Add(RefreshTokenDuration)
	SessionID := GeneratedUUIDSession(credentials.UserID)

	token, err := s.Token.CreateToken(ctx, TokenEntity{ID: SessionID, Email: credentials.Email, UserID: credentials.UserID, Role: credentials.Role, Version: 1, RevokeAt: accessTokenRevokedAt})
	if err != nil {
		return TokenResponse{}, err
	}
	refreshToken, err := s.Token.CreateToken(ctx, TokenEntity{ID: SessionID, Email: credentials.Email, UserID: credentials.UserID, Role: credentials.Role, Version: 1, RevokeAt: refreshTokenRevokedAt})
	if err != nil {
		return TokenResponse{}, err
	}
	err = s.SessionRepo.Create(ctx, CreateSessionPayload{
		SessionID:    SessionID,
		AccessToken:  token,
		RefreshToken: refreshToken,
		RevokeAt:     accessTokenRevokedAt,
		UserID:       credentials.UserID,
	})
	if err != nil {
		return TokenResponse{}, err
	}
	return TokenResponse{
		AccessToken:  token,
		RefreshToken: refreshToken,
	}, nil

}

func (s *Service) Register(ctx context.Context, payload RegisterPayload) error {
	// hash password

	switch {
	case len(strings.Split(payload.Password, "")) < PasswordLength:
		return ErrPasswordShort
	}

	password, err := utils.HashString(payload.Password)
	if err != nil {
		return err
	}

	// check email valid
	if !utils.IsEmailValid(payload.Email) {
		return ErrInvalidEmail
	}

	payload.Password = password
	payload.CredentialID = GeneratedUUIDCredential(payload.UserID)
	err = s.CredentialRepository.Registration(ctx, payload)
	if err != nil {
		translate := mapper.MapMySQLError(err)
		if translate.Error() == mapper.ErrDuplicate.Error() {
			return ErrEmailAlreadyExists
		}
		return translate
	}
	return nil

}

func (s *Service) Refresh(ctx context.Context, oldrefreshToken string) (TokenResponse, error) {
	token, err := s.Token.VerifyToken(ctx, oldrefreshToken)
	if err != nil {
		return TokenResponse{}, err
	}
	// Revoke Version
	session, err := s.SessionRepo.GetByRefreshToken(ctx, oldrefreshToken)
	if err != nil {
		translate := mapper.MapMySQLError(err)
		if translate.Error() == mapper.ErrNoRows.Error() {
			return TokenResponse{}, ErrInvalidCredentials
		}
		return TokenResponse{}, translate
	}
	if session.Version != token.Version {
		return TokenResponse{AccessToken: strconv.Itoa(session.Version), RefreshToken: strconv.Itoa(token.Version)}, ErrTokenExpired
	}

	if !session.RevokeAt.IsZero() && session.RevokeAt.Before(time.Now()) {
		return TokenResponse{}, ErrTokenExpired
	}
	nextVersion := session.Version + 1
	accessTokenRevokedAt := time.Now().Add(AccessTokenDuration)
	refreshTokenRevokedAt := time.Now().Add(RefreshTokenDuration)
	role, err := s.CredentialRepository.GetRoleByUserId(ctx, session.UserID)
	if err != nil {
		translate := mapper.MapMySQLError(err)
		if translate.Error() == mapper.ErrNoRows.Error() {
			return TokenResponse{}, ErrInvalidCredentials
		}
		return TokenResponse{}, translate
	}

	newToken, err := s.Token.CreateToken(ctx, TokenEntity{ID: session.ID, Email: token.Email, Role: role, Version: nextVersion, UserID: session.ID, RevokeAt: accessTokenRevokedAt})
	if err != nil {
		return TokenResponse{}, err
	}
	newRefreshToken, err := s.Token.CreateToken(ctx, TokenEntity{ID: session.ID, Email: token.Email, Role: role, Version: nextVersion, UserID: session.ID, RevokeAt: refreshTokenRevokedAt})
	if err != nil {
		return TokenResponse{}, err
	}
	err = s.SessionRepo.Rotate(ctx, UpdateSessionPayload{
		UserID:          session.UserID,
		OldrefreshToken: oldrefreshToken,
		AccessToken:     newToken,
		RefreshToken:    newRefreshToken,
		RevokeAt:        accessTokenRevokedAt,
	})
	if err != nil {
		translate := mapper.MapMySQLError(err)
		if translate.Error() == mapper.ErrNoRows.Error() {
			return TokenResponse{}, ErrInvalidCredentials
		}
		return TokenResponse{}, translate
	}
	return TokenResponse{AccessToken: newToken, RefreshToken: newRefreshToken}, nil
}

func (s *Service) Logout(ctx context.Context, oldrefreshToken string) error {
	// Revoke Version
	user, err := s.Token.VerifyToken(ctx, oldrefreshToken)
	if err != nil {
		return err
	}
	err = s.SessionRepo.Delete(ctx, DeleteSessionPayload{
		UserID:       user.ID,
		RefreshToken: oldrefreshToken,
	})

	if err != nil {
		translate := mapper.MapMySQLError(err)
		if translate.Error() == mapper.ErrNoRows.Error() {
			return ErrInvalidCredentials
		}
		return translate
	}
	return nil
}

func (s *Service) Verify(ctx context.Context, token string) (TokenEntity, error) {
	tokenMd, err := s.Token.VerifyToken(ctx, token)
	if err != nil {
		return TokenEntity{}, ErrTokenExpired
	}
	// Call repository to check version
	version, err := s.SessionRepo.GetVersion(ctx, token)
	if err != nil {
		translate := mapper.MapMySQLError(err)
		if translate.Error() == mapper.ErrNoRows.Error() {
			return TokenEntity{}, ErrInvalidCredentials
		}
		return TokenEntity{}, translate
	}

	if tokenMd.Version != version {
		return TokenEntity{}, ErrTokenExpired
	}
	return tokenMd, nil
}
