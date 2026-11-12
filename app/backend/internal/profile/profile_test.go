package profile_test

import (
	"context"
	"profile-template/internal/config"
	"profile-template/internal/db"
	"profile-template/internal/db/mapper"
	"profile-template/internal/profile"
	"testing"
)

func TestRepo(t *testing.T) {
	db, err := db.NewMySQLDatabase(&config.DBConfig{
		Host:     "localhost",
		Port:     "3306",
		User:     "root",
		Password: "",
		DbName:   "sgintern-test",
	})
	if err != nil {
		t.Fatal(err)
	}
	repo := profile.NewMySQLRepository(db)
	t.Run("CreateUser", func(t *testing.T) {
		err = repo.CreateUser(context.Background(), profile.CreateData{
			ID:       "Test",
			Username: "Test",
			Tag:      "Test",
		})
		if err != nil {
			transalteErr := mapper.MapMySQLError(err)
			if transalteErr.Error() == mapper.ErrDuplicate.Error() {
				t.Log("Data Sudah ada")
				return
			}
			t.Log(err)
			t.Fatal("Error ", err)

		}
	})

	t.Run("GetUser", func(t *testing.T) {
		res, err := repo.GetUserByID(context.Background(), "Test")
		if err != nil {
			t.Log(err)
			t.Fatal("Error ", err)

		}
		if res.ID == "" {
			t.Log(err)
			t.Fatal("Error ", err)
		}
		t.Log(res)
	})

}
