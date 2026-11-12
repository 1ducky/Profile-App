package auth_test

import (
	"profile-template/internal/auth"
	"profile-template/internal/auth/credential"
	"profile-template/internal/auth/jwt"
	"profile-template/internal/auth/session"
	"profile-template/internal/config"
	"profile-template/internal/db"
	"testing"
)

func CleanUpTableCredential(t *testing.T, dbconn db.DBTX) {
	query := `TRUNCATE TABLE credential`
	_, err := dbconn.ExecContext(t.Context(), query)
	if err != nil {
		t.Fatal("Error: ", err)
	}
}
func CleanUpTableSession(t *testing.T, dbconn db.DBTX) {
	query := `TRUNCATE TABLE session`
	_, err := dbconn.ExecContext(t.Context(), query)
	if err != nil {
		t.Fatal("Error: ", err)
	}
}

func NewMySQLDatabase(t *testing.T, conf config.DBConfig) db.DBTX {
	conn, err := db.NewMySQLDatabase(&conf)
	if err != nil {
		t.Fatal("Error: ", err)
	}
	return conn
}

func NewMysqlCredentialRepo(t *testing.T, db db.DBTX) auth.CredentialRepositoryContract {
	return credential.NewMySQLRepository(db)
}
func NewMysqlSessionRepo(t *testing.T, db db.DBTX) auth.SessionRepositoryContract {
	return session.NewMySQLRepository(db)
}

func TestMain(t *testing.T) {
	conf, err := config.LoadConfig()
	if err != nil {
		t.Fatal("Error: ", err)
	}
	mysql := NewMySQLDatabase(t, conf.DBConfig)
	CleanUpTableCredential(t, mysql)
	CleanUpTableSession(t, mysql)
	credentialRepo := NewMysqlCredentialRepo(t, mysql)
	sessionRepo := NewMysqlSessionRepo(t, mysql)
	JWT := jwt.NewJWT(conf.AuthConfig)

	AuthUsecase := auth.NewAuthService(credentialRepo, JWT, sessionRepo)
	if AuthUsecase == nil {
		t.Fatal("Error: Usecase should not nil")
	}

	var tokens []auth.TokenResponse

	t.Run("VALID_REGISTRATION", func(t *testing.T) {
		err := AuthUsecase.Register(t.Context(), auth.RegisterPayload{
			Email:        "valide@mail.com",
			Password:     "password",
			CredentialID: "VALID_CREDENTIAL_RECORD",
			UserID:       "VALID_USER_ID",
		})
		if err != nil {
			t.Fatal("Error: ", err)
		}
	})
	t.Run("INVALID_EMAIL_REGISTRATION", func(t *testing.T) {
		err := AuthUsecase.Register(t.Context(), auth.RegisterPayload{
			Email:        "[EMAIL_ADDRESS]",
			Password:     "password",
			CredentialID: "VALID_CREDENTIAL_RECORD",
			UserID:       "VALID_USER_ID",
		})
		if err == nil {
			t.Fatal("Error: Should not register with invalid email")
		}
	})
	t.Run("INVALID_PW_REGISTRATION", func(t *testing.T) {
		err := AuthUsecase.Register(t.Context(), auth.RegisterPayload{
			Email:        "[EMAIL_ADDRESS]",
			Password:     "passwd",
			CredentialID: "VALID_CREDENTIAL_RECORD",
			UserID:       "VALID_USER_ID",
		})
		if err == nil {
			t.Fatal("Error: Should not register with invalid password")
		}
	})
	t.Run("DUPLICATE_REGISTRATION", func(t *testing.T) {
		err := AuthUsecase.Register(t.Context(), auth.RegisterPayload{
			Email:        "valide@mail.com",
			Password:     "password",
			CredentialID: "VALID_CREDENTIAL_RECORD",
			UserID:       "VALID_USER_ID",
		})
		if err == nil {
			t.Fatal("Error: Should not register with duplicate email")
		}
	})
	t.Run("VALID_LOGIN", func(t *testing.T) {
		tokenResponse, err := AuthUsecase.Login(t.Context(), auth.LoginPayload{
			Email:    "valide@mail.com",
			Password: "password",
		})
		if err != nil {
			t.Fatal("Error: ", err)
		}
		if tokenResponse.AccessToken == "" || tokenResponse.RefreshToken == "" {
			t.Fatal("Error: Token should not empty")
		}
		tokenEntity, _ := JWT.VerifyToken(t.Context(), tokenResponse.RefreshToken)
		t.Logf("tokenEntity: %v", tokenEntity)

		tokens = append(tokens, tokenResponse)
	})
	t.Run("INVALID_CREDENTIAL_LOGIN", func(t *testing.T) {
		_, err := AuthUsecase.Login(t.Context(), auth.LoginPayload{
			Email:    "ivalid@mail.com",
			Password: "password",
		})
		if err == nil {
			t.Fatal("Error: Should not login with invalid email")
		}
	})

	t.Run("VALID_REFRESH_TOKEN", func(t *testing.T) {
		tokenResponse, err := AuthUsecase.Refresh(t.Context(), tokens[0].RefreshToken)
		if err != nil {
			t.Log("Session", tokenResponse.AccessToken)
			t.Log("Token", tokenResponse.RefreshToken)
			t.Fatal("Error: ", err)
		}
		if tokenResponse.AccessToken == "" || tokenResponse.RefreshToken == "" {
			t.Fatal("Error: Token should not empty")
		}

		tokens = append(tokens, tokenResponse)
	})
	t.Run("INVALID_OLD_REFRESH_TOKEN", func(t *testing.T) {
		_, err := AuthUsecase.Refresh(t.Context(), tokens[0].RefreshToken)
		if err == nil {
			t.Fatal("Error: Should not refresh with invalid token")
		}
	})

	t.Run("VALID_LOG_OUT", func(t *testing.T) {
		err := AuthUsecase.Logout(t.Context(), tokens[1].RefreshToken)
		if err != nil {
			t.Fatal("Error: ", err)
		}
	})
	t.Run("LOG_OUT_TWICE", func(t *testing.T) {
		err := AuthUsecase.Logout(t.Context(), tokens[1].RefreshToken)
		if err != nil {
			t.Fatal("Error: ", err)
		}
	})
	t.Run("INVALID_REFRESH_TOKEN", func(t *testing.T) {
		_, err := AuthUsecase.Refresh(t.Context(), tokens[0].RefreshToken)
		if err == nil {
			t.Fatal("Error: Should not refresh with invalid token")
		}
	})

}
