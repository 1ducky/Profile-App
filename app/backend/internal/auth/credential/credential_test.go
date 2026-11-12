package credential_test

import (
	"profile-template/internal/auth"
	"profile-template/internal/auth/credential"
	"profile-template/internal/config"
	"profile-template/internal/db"
	"testing"
)

func CreateRepository(t *testing.T, mysql db.DBTX) auth.CredentialRepositoryContract {
	return credential.NewMySQLRepository(mysql)
}

func CreateConn(t *testing.T, cfg config.DBConfig) db.DBTX {
	mySQL, err := db.NewMySQLDatabase(&cfg)
	if err != nil {
		t.Fatal("Error: ", err)
	}
	return mySQL
}

func CleanUpTable(t *testing.T, mysql db.DBTX) {
	query := `TRUNCATE TABLE credential`
	_, err := mysql.ExecContext(t.Context(), query)
	if err != nil {
		t.Fatal("Error: ", err)
	}
}

func TestMain(t *testing.T) {
	cfg := config.LoadMySQLDatabaseConfig()
	conn := CreateConn(t, cfg)
	repo := CreateRepository(t, conn)
	if repo == nil {
		t.Fatal("Error: repo should not nil")
	}

	CleanUpTable(t, conn)

	t.Run("Registration Successful", func(t *testing.T) {
		err := repo.Registration(t.Context(), auth.RegisterPayload{
			Email:        "[EMAIL_ADDRESS]",
			Password:     "password",
			CredentialID: "VALID_CREDENTIAL_RECORD",
			UserID:       "VALID_USER_ID",
		})
		if err != nil {
			t.Error("Error: ", err)
		}
	})

	t.Run("Registration Duplicated Email", func(t *testing.T) {
		err := repo.Registration(t.Context(), auth.RegisterPayload{
			Email:        "[EMAIL_ADDRESS]",
			Password:     "password",
			UserID:       "VALID_USER_ID",
			CredentialID: "INVALID_CREDENTIAL_DUPLICATE",
		})
		if err == nil {
			t.Error("Error: Should not register with same email")
		}
	})
	t.Run("Registration Duplicated UserID", func(t *testing.T) {
		err := repo.Registration(t.Context(), auth.RegisterPayload{
			Email:        "[EMAIL_ADDRESS]1",
			Password:     "password",
			UserID:       "VALID_USER_ID",
			CredentialID: "INVALID_CREDENTIAL_DUPLICATE",
		})
		if err == nil {
			t.Error("Error: Should not register with same UserID")
		}
	})

	t.Run("VALID_AUTHENTICATION", func(t *testing.T) {
		_, err := repo.Authentication(t.Context(), "[EMAIL_ADDRESS]")
		if err != nil {
			t.Error("Error: ", err)
		}

	})
	t.Run("INVALID_AUTHENTICATION", func(t *testing.T) {
		_, err := repo.Authentication(t.Context(), "[EMAIL_ADDRESS]qweqweqwe")
		if err == nil {
			t.Error("Error: Should not authenticate with invalid email")
		}

	})
	t.Run("VALID_GETROLE", func(t *testing.T) {
		_, err := repo.GetRoleByUserId(t.Context(), "VALID_USER_ID")
		if err != nil {
			t.Error("Error: ", err)
		}

	})
	t.Run("INVALID_GETROLE", func(t *testing.T) {
		_, err := repo.GetRoleByUserId(t.Context(), "INVALID_USER_ID")
		if err == nil {
			t.Error("Error: Should not get role with invalid UserID")
		}

	})

}
