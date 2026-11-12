package article

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"profile-template/internal/db"
	"strings"
	"time"
)

type mysqlRepository struct {
	db db.DBTX
}

// NewMySQLRepository membuat Repository berbasis MySQL.
func NewMySQLRepository(db db.DBTX) ArticlaeRepositoryContract {
	return &mysqlRepository{db: db}
}

const tableName = "article"

// Kolom yang di-select untuk entity (urutan harus sama dengan scanArticle).
const (
	colsWithoutContent = "id, slug, url, title, excerpt, category, badges, tags, cover, author_id, published_at, updated_at"
)

// ---------- Create ----------

func (r *mysqlRepository) Create(ctx context.Context, req CreateRequest) (*ArticleEntity, error) {
	if req.ID == "" {
		return nil, ErrInvalidID
	}

	const q = `INSERT INTO ` + tableName + `
		(id, slug, url, title, excerpt, category, badges, tags, cover, author_id, content, published_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	now := time.Now()
	_, err := r.db.ExecContext(ctx, q,
		req.ID,
		req.Slug,
		req.URL,
		req.Title,
		db.NullString(req.Excerpt),
		req.Category,
		stringList(req.Badges),
		stringList(req.Tags),
		db.NullString(req.Cover),
		req.AuthorID,
		req.Content,
		db.NullTimePtr(req.PublishAt),
		now,
	)
	if err != nil {
		if db.IsDuplicateEntry(err) {
			return nil, ErrDuplicateSlug
		}
		return nil, fmt.Errorf("article: create: %w", err)
	}

	return r.FindByID(ctx, req.ID)
}

// ---------- Update ----------

func (r *mysqlRepository) Update(ctx context.Context, req UpdateRequest) error {
	if req.ID == "" {
		return ErrInvalidID
	}

	sets := make([]string, 0, 11)
	args := make([]any, 0, 12)

	add := func(col string, val any) {
		sets = append(sets, col+" = ?")
		args = append(args, val)
	}

	if req.Slug != nil {
		add("slug", *req.Slug)
	}
	if req.URL != nil {
		add("url", *req.URL)
	}
	if req.Title != nil {
		add("title", *req.Title)
	}
	if req.Excerpt != nil {
		add("excerpt", db.NullString(*req.Excerpt))
	}
	if req.Category != nil {
		add("category", *req.Category)
	}
	if req.Badges != nil {
		add("badges", stringList(*req.Badges))
	}
	if req.Tags != nil {
		add("tags", stringList(*req.Tags))
	}
	if req.Cover != nil {
		add("cover", db.NullString(*req.Cover))
	}
	if req.Content != nil {
		add("content", *req.Content)
	}
	if req.PublishAt != nil {
		add("published_at", *req.PublishAt)
	}

	if len(sets) == 0 {
		return ErrNothingToUpdate
	}

	updatedAt := req.UpdateAt
	if updatedAt.IsZero() {
		updatedAt = time.Now()
	}
	add("updated_at", updatedAt)

	args = append(args, req.ID, req.AuthorID)
	q := "UPDATE " + tableName + " SET " + strings.Join(sets, ", ") + " WHERE id = ? AND author_id = ?"

	res, err := r.db.ExecContext(ctx, q, args...)
	if err != nil {
		if db.IsDuplicateEntry(err) {
			return ErrDuplicateSlug
		}
		return fmt.Errorf("article: update: %w", err)
	}

	// MySQL mengembalikan 0 affected rows jika nilai tidak berubah,
	// jadi pastikan dulu baris memang ada sebelum menyatakan not found.
	affected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("article: update rows affected: %w", err)
	}
	if affected == 0 {
		exists, err := r.existsByID(ctx, req.ID)
		if err != nil {
			return err
		}
		if !exists {
			return ErrNotFound
		}
	}
	// fmt.Println("QUERY:", q)
	// fmt.Println("ARGS:", args)
	// fmt.Println("AFFECTED:", affected)
	return nil
}

// ---------- Delete ----------

func (r *mysqlRepository) Delete(ctx context.Context, req DeleteRequest) error {
	if req.ID == "" {
		return ErrInvalidID
	}

	res, err := r.db.ExecContext(ctx, "DELETE FROM "+tableName+" WHERE id = ? AND author_id = ?", req.ID, req.AuthorID)
	if err != nil {
		return fmt.Errorf("article: delete: %w", err)
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("article: delete rows affected: %w", err)
	}
	if affected == 0 {
		return ErrNotFound
	}
	return nil
}

// ---------- FindByID / FindBySlug ----------

func (r *mysqlRepository) FindByID(ctx context.Context, id string) (*ArticleEntity, error) {
	if id == "" {
		return nil, ErrInvalidID
	}
	return r.findOne(ctx, "id = ?", id)
}

func (r *mysqlRepository) FindBySlug(ctx context.Context, slug string) (*ArticleEntity, error) {
	return r.findOne(ctx, "slug = ?", slug)
}

func (r *mysqlRepository) findOne(ctx context.Context, where string, arg any) (*ArticleEntity, error) {
	q := "SELECT " + colsWithoutContent + ", content FROM " + tableName + " WHERE " + where + " LIMIT 1"

	e, err := scanArticle(r.db.QueryRowContext(ctx, q, arg))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("article: find: %w", err)
	}
	return e, nil
}

// ---------- FindAll ----------

func (r *mysqlRepository) FindAll(ctx context.Context, req FindAllRequest) ([]ArticleEntity, int64, error) {
	req.Normalize()

	where, args, err := buildFilter(req)
	if err != nil {
		return nil, 0, err
	}

	// Total baris (untuk pagination)
	var total int64
	countQ := "SELECT COUNT(*) FROM " + tableName + where
	if err := r.db.QueryRowContext(ctx, countQ, args...).Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("article: count: %w", err)
	}
	if total == 0 {
		return []ArticleEntity{}, 0, nil
	}

	contentCol := "'' AS content"
	if req.IncludeContent {
		contentCol = "content"
	}

	// SortBy dan Order sudah divalidasi whitelist di Normalize(),
	// jadi aman disisipkan langsung. id dipakai sebagai tie-breaker.
	q := "SELECT " + colsWithoutContent + ", " + contentCol +
		" FROM " + tableName + where +
		fmt.Sprintf(" ORDER BY %s %s, id ASC LIMIT ? OFFSET ?", req.SortBy, req.Order)

	listArgs := append(append([]any{}, args...), req.PageSize, req.Offset())

	rows, err := r.db.QueryContext(ctx, q, listArgs...)
	if err != nil {
		return nil, 0, fmt.Errorf("article: find all: %w", err)
	}
	defer rows.Close()

	items := make([]ArticleEntity, 0, req.PageSize)
	for rows.Next() {
		e, err := scanArticle(rows)
		if err != nil {
			return nil, 0, fmt.Errorf("article: scan: %w", err)
		}
		items = append(items, *e)
	}
	if err := rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("article: rows: %w", err)
	}

	return items, total, nil
}

// buildFilter membuat klausa WHERE (diawali spasi, atau string kosong) beserta argumennya.
func buildFilter(req FindAllRequest) (string, []any, error) {
	conds := make([]string, 0, 8)
	args := make([]any, 0, 8)

	if s := strings.TrimSpace(req.Search); s != "" {
		like := "%" + db.EscapeLike(s) + "%"
		conds = append(conds, "(title LIKE ? OR excerpt LIKE ?)")
		args = append(args, like, like)
	}
	if req.Category != "" {
		conds = append(conds, "category = ?")
		args = append(args, req.Category)
	}
	if req.AuthorID != "" {
		conds = append(conds, "author_id = ?")
		args = append(args, req.AuthorID)
	}
	if len(req.Tags) > 0 {
		b, err := json.Marshal(req.Tags)
		if err != nil {
			return "", nil, fmt.Errorf("article: marshal tags filter: %w", err)
		}
		// Butuh MySQL >= 8.0.17
		conds = append(conds, "JSON_OVERLAPS(tags, CAST(? AS JSON))")
		args = append(args, string(b))
	}
	if len(req.Badges) > 0 {
		b, err := json.Marshal(req.Badges)
		if err != nil {
			return "", nil, fmt.Errorf("article: marshal badges filter: %w", err)
		}
		conds = append(conds, "JSON_OVERLAPS(badges, CAST(? AS JSON))")
		args = append(args, string(b))
	}
	if req.OnlyPublish {
		conds = append(conds, "published_at IS NOT NULL AND published_at <= ?")
		args = append(args, time.Now())
	}
	if req.PublishFrom != nil {
		conds = append(conds, "published_at >= ?")
		args = append(args, *req.PublishFrom)
	}
	if req.PublishTo != nil {
		conds = append(conds, "published_at <= ?")
		args = append(args, *req.PublishTo)
	}

	if len(conds) == 0 {
		return "", args, nil
	}
	return " WHERE " + strings.Join(conds, " AND "), args, nil
}

// ---------- ExistsBySlug ----------

func (r *mysqlRepository) ExistsBySlug(ctx context.Context, slug string, excludeID string) (bool, error) {
	q := "SELECT EXISTS(SELECT 1 FROM " + tableName + " WHERE slug = ?"
	args := []any{slug}
	if excludeID != "" {
		q += " AND id <> ?"
		args = append(args, excludeID)
	}
	q += ")"

	var exists bool
	if err := r.db.QueryRowContext(ctx, q, args...).Scan(&exists); err != nil {
		return false, fmt.Errorf("article: exists by slug: %w", err)
	}
	return exists, nil
}

func (r *mysqlRepository) existsByID(ctx context.Context, id string) (bool, error) {
	var exists bool
	err := r.db.QueryRowContext(ctx,
		"SELECT EXISTS(SELECT 1 FROM "+tableName+" WHERE id = ?)", id).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("article: exists by id: %w", err)
	}
	return exists, nil
}

// ===================== Scan helpers =====================

type rowScanner interface {
	Scan(dest ...any) error
}

// scanArticle membaca satu baris dengan urutan kolom:
// id, slug, url, title, excerpt, category, badges, tags, cover, author_id,
// published_at, updated_at, content
func scanArticle(s rowScanner) (*ArticleEntity, error) {
	var (
		e         ArticleEntity
		excerpt   sql.NullString
		cover     sql.NullString
		badges    stringList
		tags      stringList
		publishAt sql.NullTime
		updateAt  sql.NullTime
	)

	err := s.Scan(
		&e.ID,
		&e.Slug,
		&e.URL,
		&e.Title,
		&excerpt,
		&e.Category,
		&badges,
		&tags,
		&cover,
		&e.AuthorID,
		&publishAt,
		&updateAt,
		&e.Content,
	)
	if err != nil {
		return nil, err
	}

	e.Excerpt = excerpt.String
	e.Cover = cover.String
	e.Badges = []string(badges)
	e.Tags = []string(tags)
	if publishAt.Valid {
		e.PublishAt = publishAt.Time
	}
	if updateAt.Valid {
		e.UpdateAt = updateAt.Time
	}
	return &e, nil
}

// stringList menyimpan []string ke kolom JSON dan membacanya kembali.
type stringList []string

// Value mengimplementasikan driver.Valuer.
func (l stringList) Value() (driver.Value, error) {
	if l == nil {
		return "[]", nil
	}
	b, err := json.Marshal([]string(l))
	if err != nil {
		return nil, err
	}
	return string(b), nil
}

// Scan mengimplementasikan sql.Scanner.
func (l *stringList) Scan(src any) error {
	var raw []byte
	switch v := src.(type) {
	case nil:
		*l = []string{}
		return nil
	case []byte:
		raw = v
	case string:
		raw = []byte(v)
	default:
		return fmt.Errorf("article: cannot scan %T into stringList", src)
	}
	if len(raw) == 0 {
		*l = []string{}
		return nil
	}
	var out []string
	if err := json.Unmarshal(raw, &out); err != nil {
		return err
	}
	if out == nil {
		out = []string{}
	}
	*l = out
	return nil
}
