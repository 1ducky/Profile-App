package db

import (
	"context"
	"database/sql"
	"fmt"
	"profile-template/internal/config"
	"time"

	_ "github.com/go-sql-driver/mysql" //initalize
)

func NewMySQLDatabase(config *config.DBConfig) (*sql.DB, error) {
	var dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.User, config.Password, config.Host, config.Port, config.DbName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Hour)
	db.SetConnMaxIdleTime(15 * time.Minute)

	return db, nil
}
