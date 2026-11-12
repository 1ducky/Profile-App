package uow

import (
	"context"
	"database/sql"
	"log"
	"profile-template/internal/db"
)

type UnitOfWork[R any] interface {
	Do(ctx context.Context, fn func(ctx context.Context, r R) error) error
}

type UoW[R any] struct {
	db    *sql.DB
	build func(db.DBTX) R
}

func NewUoW[R any](db *sql.DB, build func(db.DBTX) R) UnitOfWork[R] {
	return &UoW[R]{
		db:    db,
		build: build,
	}
}

func (u *UoW[R]) Do(ctx context.Context, fn func(ctx context.Context, r R) error) error {
	// transaction start
	log.Print("Start Transaction")
	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		err := tx.Rollback()
		if err != nil {
			log.Println("Rollback Transaction Error: ", err.Error())
		}
	}()
	if err := fn(ctx, u.build(tx)); err != nil {
		log.Print("Rollback Transaction Error: ", err.Error())
		return err
	}
	return tx.Commit()

}
