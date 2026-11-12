package category

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"profile-template/internal/article"
	"profile-template/internal/db"
)

// ErrNotFound dikembalikan saat data tidak ditemukan.
var ErrNotFound = errors.New("category: not found")

type repository struct {
	db db.DBTX
}

// NewRepository membuat implementasi Repository berbasis database/sql.
func NewRepository(db db.DBTX) article.CategoryRepositoryContract {
	return &repository{db: db}
}

// Create menyimpan category baru.
func (r *repository) Create(ctx context.Context, e *article.CategoryEntity) error {
	query := fmt.Sprintf(
		"INSERT INTO %s (%s, %s) VALUES (?, ?)",
		article.CATEGORY_TABLE, article.CATEGORY_ID, article.CATEGORY_Slug,
	)
	_, err := r.db.ExecContext(ctx, query, e.ID, e.Slug)
	return err
}

// GetByID mengambil satu category berdasarkan id.
func (r *repository) GetByID(ctx context.Context, id string) (*article.CategoryEntity, error) {
	query := fmt.Sprintf(
		"SELECT %s, %s FROM %s WHERE %s = ?",
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
func (r *repository) GetAll(ctx context.Context) ([]article.CategoryEntity, error) {
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
func (r *repository) Update(ctx context.Context, e *article.CategoryEntity) error {
	query := fmt.Sprintf(
		"UPDATE %s SET %s = ? WHERE %s = ?",
		article.CATEGORY_TABLE, article.CATEGORY_Slug, article.CATEGORY_ID,
	)
	res, err := r.db.ExecContext(ctx, query, e.Slug, e.ID)
	if err != nil {
		return err
	}
	return checkAffected(res)
}

// Delete menghapus category berdasarkan id.
func (r *repository) Delete(ctx context.Context, id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s = ?", article.CATEGORY_TABLE, article.CATEGORY_ID)
	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return checkAffected(res)
}

func checkAffected(res sql.Result) error {
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return ErrNotFound
	}
	return nil
}
