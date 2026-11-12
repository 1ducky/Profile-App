package session_test

import (
	"context"
	"profile-template/internal/auth"
	"profile-template/internal/auth/session"
	"profile-template/internal/config"
	"profile-template/internal/db"
	"profile-template/internal/pipeline"
	"testing"
	"time"
)

func CreatemySQL(t *testing.T) db.DBTX {
	conf := config.LoadMySQLDatabaseConfig()
	mySQL, err := db.NewMySQLDatabase(&conf)
	if err != nil {
		t.Fatal("Error: ", err)
	}
	return mySQL

}

func CreateRepository(mysql db.DBTX) auth.SessionRepositoryContract {
	return session.NewMySQLRepository(mysql)
}

func CleanUpTable(t *testing.T, mysql db.DBTX, ctx context.Context) {
	t.Log("Trunecate Testing Table")
	query := `TRUNCATE TABLE session`
	_, err := mysql.ExecContext(ctx, query)
	if err != nil {
		t.Error("Error: ", err)
	}
	t.Log("Trunecate Testing Table Done")
}

func TestSessionUsecase(t *testing.T) {
	mysql := CreatemySQL(t)
	repo := CreateRepository(mysql)
	ctx := context.Background()
	CleanUpTable(t, mysql, ctx)

	t.Run("Create Session Successfully", func(t *testing.T) {
		err := repo.Create(ctx, auth.CreateSessionPayload{
			SessionID:    "VALID_SESSION_ID_1",
			UserID:       "VALID_USER_ID_1",
			AccessToken:  "VALID_ACCESS_TOKEN_1",
			RefreshToken: "VALID_REFRESH_TOKEN_1",
			RevokeAt:     time.Now(),
		})
		if err != nil {
			t.Error("Error: ", err)
		}
	})
	t.Run("Dup Session Invalid", func(t *testing.T) {
		err := repo.Create(ctx, auth.CreateSessionPayload{
			SessionID:    "VALID_SESSION_ID_1",
			UserID:       "VALID_USER_ID_1",
			AccessToken:  "VALID_ACCESS_TOKEN_1",
			RefreshToken: "VALID_REFRESH_TOKEN_1",
			RevokeAt:     time.Now(),
		})
		if err == nil {
			t.Error("Error: ShouldDuplicateSession")
		}
		t.Log(err)
	})
	t.Run("Rotate Session Valid", func(t *testing.T) {
		err := repo.Rotate(ctx, auth.UpdateSessionPayload{
			UserID:          "VALID_USER_ID_1",
			AccessToken:     "VALID_ACCESS_TOKEN_2",
			RefreshToken:    "VALID_REFRESH_TOKEN_2",
			OldrefreshToken: "VALID_REFRESH_TOKEN_1",
			RevokeAt:        time.Now(),
		})
		if err != nil {
			t.Error("Error: ", err)
		}
		t.Log(err)
	})
	t.Run("Rotate Session Invalid (Old Refresh Token Mismatch)", func(t *testing.T) {
		err := repo.Rotate(ctx, auth.UpdateSessionPayload{
			UserID:          "VALID_USER_ID_1",
			AccessToken:     "VALID_ACCESS_TOKEN_2",
			RefreshToken:    "VALID_REFRESH_TOKEN_2",
			OldrefreshToken: "VALID_REFRESH_TOKEN_1_WRONG",
			RevokeAt:        time.Now(),
		})
		if err == nil {
			t.Error("Error: Should not update session")
		}
		t.Log(err)
	})

	t.Run("Get Version By Access Token", func(t *testing.T) {
		ver, err := repo.GetVersion(ctx, "VALID_ACCESS_TOKEN_2")
		if err != nil {
			t.Error("Error: ", err)
		}
		if ver != 2 {
			t.Error("Error: Invalid Version")
		}
		t.Log(ver)
	})
	t.Run("Get Version By Access Token Invalid", func(t *testing.T) {
		ver, err := repo.GetVersion(ctx, "INVALID_ACCESS_TOKEN_2")
		if err == nil {
			t.Error("Error: notget")
		}
		t.Log(ver, err)
	})
	t.Run("Get By Access Token", func(t *testing.T) {
		ver, err := repo.GetByAccessToken(ctx, "VALID_ACCESS_TOKEN_2")
		if err != nil {
			t.Error("Error: ", err)
		}
		t.Log(ver)
	})
	t.Run("Get By Access Token Invalid", func(t *testing.T) {
		ver, err := repo.GetByAccessToken(ctx, "INVALID_ACCESS_TOKEN_2")
		if err == nil {
			t.Error("Error: notget")
		}
		t.Log(ver, err)
	})
	t.Run("Get By Refresh Token", func(t *testing.T) {
		ver, err := repo.GetByRefreshToken(ctx, "VALID_REFRESH_TOKEN_2")
		if err != nil {
			t.Error("Error: ", err)
		}
		t.Log(ver)
	})
	t.Run("Get By Refresh Token Invalid", func(t *testing.T) {
		ver, err := repo.GetByRefreshToken(ctx, "INVALID_REFRESH_TOKEN_2")
		if err == nil {
			t.Error("Error: notget")
		}
		t.Log(ver, err)
	})

	t.Run("Delete Session Valid", func(t *testing.T) {
		err := repo.Delete(ctx, auth.DeleteSessionPayload{
			SessionID:    "VALID_SESSION_ID_1",
			UserID:       "VALID_USER_ID_1",
			RefreshToken: "VALID_REFRESH_TOKEN_1",
		})
		if err != nil {
			t.Error("Error: ", err)
		}
		t.Log(err)
	})
	t.Run("Delete Session Invalid", func(t *testing.T) {
		err := repo.Delete(ctx, auth.DeleteSessionPayload{
			SessionID:    "VALID_SESSION_ID_1",
			UserID:       "VALID_USER_ID_1",
			RefreshToken: "VALID_REFRESH_TOKEN_1",
		})
		if err != nil {
			t.Error("Error: ", err)
		}
		t.Log(err)
	})
	t.Run("Delete Session Invalid (Old Refresh Token Mismatch)", func(t *testing.T) {
		err := repo.Delete(ctx, auth.DeleteSessionPayload{
			SessionID:    "VALID_SESSION_ID_1",
			UserID:       "VALID_USER_ID_1",
			RefreshToken: "VALID_REFRESH_TOKEN_1_WRONG",
		})
		if err != nil {
			t.Error("Error: ", err)
		}
		t.Log(err)
	})
	t.Run("Delete Session Invalid (Old Refresh Token Mismatch)", func(t *testing.T) {
		err := repo.Delete(ctx, auth.DeleteSessionPayload{
			SessionID:    "VALID_SESSION_ID_1",
			UserID:       "VALID_USER_ID_1",
			RefreshToken: "VALID_REFRESH_TOKEN_1_WRONG",
		})
		if err != nil {
			t.Error("Error: ", err)
		}
		t.Log(err)
	})
	t.Run("Delete Session Invalid (Old Refresh Token Mismatch)", func(t *testing.T) {
		err := repo.Delete(ctx, auth.DeleteSessionPayload{
			SessionID:    "VALID_SESSION_ID_1",
			UserID:       "VALID_USER_ID_1",
			RefreshToken: "VALID_REFRESH_TOKEN_1_WRONG",
		})
		if err != nil {
			t.Error("Error: ", err)
		}
		t.Log(err)
	})
	t.Run("Delete Session Invalid (Old Refresh Token Mismatch)", func(t *testing.T) {
		err := repo.Delete(ctx, auth.DeleteSessionPayload{
			SessionID:    "VALID_SESSION_ID_1",
			UserID:       "VALID_USER_ID_1",
			RefreshToken: "VALID_REFRESH_TOKEN_1_WRONG",
		})
		if err != nil {
			t.Error("Error: ", err)
		}
		t.Log(err)
	})
	t.Run("Get Version By Access Token", func(t *testing.T) {
		ver, err := repo.GetVersion(ctx, "VALID_ACCESS_TOKEN_2")
		if err != nil {
			t.Error("Error: ", err)
		}
		if ver != 2 {
			t.Error("Error: Invalid Version")
		}
		t.Log(ver)
	})
}

func TestRaceCondition(t *testing.T) {
	mysql := CreatemySQL(t)
	repo := CreateRepository(mysql)
	CleanUpTable(t, mysql, t.Context())

	jobs := pipeline.ProduceJob(t.Context(), []string{
		"VALID_ACCESS_TOKEN_1",
		"VALID_ACCESS_TOKEN_1",
		"VALID_ACCESS_TOKEN_1",
		"VALID_ACCESS_TOKEN_1",
		"VALID_ACCESS_TOKEN_1",
		"VALID_ACCESS_TOKEN_1",
		"VALID_ACCESS_TOKEN_1",
		"VALID_ACCESS_TOKEN_1",
		"VALID_ACCESS_TOKEN_1",
		"VALID_ACCESS_TOKEN_1",
	}, 10)
	pool := pipeline.NewWorkerPool[string, error](4)

	t.Run("Init Session", func(t *testing.T) {
		err := repo.Create(t.Context(), auth.CreateSessionPayload{
			SessionID:    "VALID_SESSION_ID_1",
			UserID:       "VALID_USER_ID_1",
			AccessToken:  "VALID_ACCESS_TOKEN_1",
			RefreshToken: "VALID_REFRESH_TOKEN_1",
			RevokeAt:     time.Now(),
		})
		if err != nil {
			t.Error("Error: ", err)
		}
	})

	workerReport := pool.Run(t.Context(), jobs, func(ctx context.Context, s string) error {
		var raceErr error
		t.Run("Rotate", func(t *testing.T) {
			err := repo.Rotate(t.Context(), auth.UpdateSessionPayload{
				AccessToken:     "NEW_ACCESS_TOKEN",
				RefreshToken:    "NEW_REFRESH_TOKEN",
				OldrefreshToken: "VALID_REFRESH_TOKEN_1",
				UserID:          "VALID_USER_ID_1",
				RevokeAt:        time.Now(),
			})
			if err != nil {
				raceErr = err
			}
		})
		return raceErr
	})

	res := pipeline.Collect(t.Context(), workerReport, 10)

	var errCount int
	var successCount int

	for _, v := range res {
		if v != nil {
			t.Log(v)
			errCount++
		} else {
			successCount++
		}
	}

	if successCount > 1 || successCount != 1 {
		t.Error("Error: ", errCount)
	}
	t.Log("errCount:", errCount)
	t.Log("SuccessCount:", successCount)

}
