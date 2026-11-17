package category

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"profile-template/internal/article"
	"profile-template/internal/db"
)

type PGSQLRepository struct {
	db db.DBTX
}

// NewPGSQLRepository membuat implementasi Repository berbasis PostgreSQL.
func NewPQSQLRepository(db db.DBTX) article.CategoryRepositoryContract {
	return &PGSQLRepository{db: db}
}

// Create menyimpan category baru.
func (r *PGSQLRepository) Create(ctx context.Context, e *article.CategoryEntity) error {
	query := fmt.Sprintf(
		"INSERT INTO %s (%s, %s) VALUES ($1, $2)",
		article.CATEGORY_TABLE, article.CATEGORY_ID, article.CATEGORY_Slug,
	)
	_, err := r.db.ExecContext(ctx, query, e.ID, e.Slug)
	return err
}

// GetByID mengambil satu category berdasarkan id.
func (r *PGSQLRepository) GetByID(ctx context.Context, id string) (*article.CategoryEntity, error) {
	query := fmt.Sprintf(
		"SELECT %s, %s FROM %s WHERE %s = $1",
		article.CATEGORY_ID, article.CATEGORY_Slug, article.CATEGORY_TABLE, article.CATEGORY_ID,
	)

	var e article.CategoryEntity
	err := r.db.QueryRowContext(ctx, query, id).Scan(&e.ID, &e.Slug)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &e, nil
}

// GetAll mengambil semua category.
func (r *PGSQLRepository) GetAll(ctx context.Context) ([]article.CategoryEntity, error) {
	query := fmt.Sprintf("SELECT %s, %s FROM %s", article.CATEGORY_ID, article.CATEGORY_Slug, article.CATEGORY_TABLE)

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := make([]article.CategoryEntity, 0)
	for rows.Next() {
		var e article.CategoryEntity
		if err := rows.Scan(&e.ID, &e.Slug); err != nil {
			return nil, err
		}
		list = append(list, e)
	}
	return list, rows.Err()
}

// Update mengubah slug category berdasarkan id.
func (r *PGSQLRepository) Update(ctx context.Context, e *article.CategoryEntity) error {
	query := fmt.Sprintf(
		"UPDATE %s SET %s = $1 WHERE %s = $2",
		article.CATEGORY_TABLE, article.CATEGORY_Slug, article.CATEGORY_ID,
	)
	res, err := r.db.ExecContext(ctx, query, e.Slug, e.ID)
	if err != nil {
		return err
	}
	return checkAffected(res)
}

// Delete menghapus category berdasarkan id.
func (r *PGSQLRepository) Delete(ctx context.Context, id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s = $1", article.CATEGORY_TABLE, article.CATEGORY_ID)
	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return checkAffected(res)
}
