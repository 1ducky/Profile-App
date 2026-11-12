package jwt_test

import (
	"profile-template/internal/auth"
	"profile-template/internal/auth/jwt"
	"profile-template/internal/config"
	"testing"
	"time"
)

func CreateJWT(t *testing.T, cfg config.AuthConfig) auth.TokenContract {
	return jwt.NewJWT(cfg)
}

type TestToken struct {
	Token       string
	ExpectedErr error
}

func TestCreateToken(t *testing.T) {
	validCfg := config.AuthConfig{
		SecretKeyJWT: "secretkey1",
		Issuer:       "issuer1",
	}
	SecondvalidCfg := config.AuthConfig{
		SecretKeyJWT: "secretkey2",
		Issuer:       "issuer2",
	}
	jwtService := CreateJWT(t, validCfg)
	if jwtService == nil {
		t.Fatalf("should not nil")
	}
	secondjwtService := CreateJWT(t, SecondvalidCfg)
	if secondjwtService == nil {
		t.Fatalf("should not nil")
	}

	var PrimaryToken []string
	var SecondToken []string

	t.Run("VALID_TOKEN_CREATE", func(t *testing.T) {
		pToken, err := jwtService.CreateToken(t.Context(), auth.TokenEntity{
			ID:       "1",
			UserID:   "1",
			Role:     "ADMIN",
			Version:  0,
			RevokeAt: time.Now().Add(time.Minute * 5),
		})
		if err != nil {
			t.Fatalf("should not err: %v", err)
		}
		if pToken == "" {
			t.Fatalf("should not empty")
		}
		PrimaryToken = append(PrimaryToken, pToken)
	})
	t.Run("VALID_SECOND_TOKEN_CREATE", func(t *testing.T) {
		sToken, err := secondjwtService.CreateToken(t.Context(), auth.TokenEntity{
			ID:       "2",
			UserID:   "2",
			Role:     "USER",
			Version:  0,
			RevokeAt: time.Now().Add(time.Minute * 5),
		})
		if err != nil {
			t.Fatalf("should not err")
		}
		if sToken == "" {
			t.Fatalf("should not empty")
		}
		SecondToken = append(SecondToken, sToken)
	})

	t.Run("VALID_VERIFY_TOKEN", func(t *testing.T) {
		_, pErr := jwtService.VerifyToken(t.Context(), PrimaryToken[0])
		if pErr != nil {
			t.Fatalf("should not err: %v, token: %s", pErr, PrimaryToken[0])
		}

	})
	t.Run("VALID_SECOND_VERIFY_TOKEN", func(t *testing.T) {
		_, sErr := secondjwtService.VerifyToken(t.Context(), SecondToken[0])
		if sErr != nil {
			t.Fatalf("should not err: %v, token: %s", sErr, SecondToken[0])
		}
	})

	t.Run("INVALID_VERIFY_TOKEN", func(t *testing.T) {
		_, pErr := jwtService.VerifyToken(t.Context(), SecondToken[0])
		if pErr == nil {
			t.Fatalf("should err")
		}
	})
	t.Run("INVALID_SECOND_VERIFY_TOKEN", func(t *testing.T) {
		_, sErr := secondjwtService.VerifyToken(t.Context(), PrimaryToken[0])
		if sErr == nil {
			t.Fatalf("should err")
		}
	})

}
