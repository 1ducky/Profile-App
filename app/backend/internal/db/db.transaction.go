package db

import (
	"context"
	"database/sql"
)

type TransactionManager struct {
	DB *sql.DB
}

func NewDBTransaction(Db *sql.DB) *TransactionManager {
	return &TransactionManager{DB: Db}
}

func (t *TransactionManager) Do(ctx context.Context, fn func(*sql.Tx) error) error {
	tx, err := t.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		}
	}()

	err = fn(tx)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
