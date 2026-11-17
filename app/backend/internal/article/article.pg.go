package article

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"profile-template/internal/db"
	"strconv"
	"strings"
	"time"

	"github.com/lib/pq"
)

type PGSQLRepository struct {
	db db.DBTX
}

// NewPGSQLRepository membuat Repository berbasis PostgreSQL.
func NewPGSQLRepository(db db.DBTX) ArticlaeRepositoryContract {
	return &PGSQLRepository{db: db}
}

// ph mengembalikan placeholder Postgres ke-n, mis. "$3".
func ph(n int) string { return "$" + strconv.Itoa(n) }

// ---------- Create ----------

func (r *PGSQLRepository) Create(ctx context.Context, req CreateRequest) (*ArticleEntity, error) {
	if req.ID == "" {
		return nil, ErrInvalidID
	}

	const q = `INSERT INTO ` + tableName + `
		(id, slug, url, title, excerpt, category, badges, tags, cover, author_id, content, published_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`

	now := time.Now()
	_, err := r.db.ExecContext(ctx, q,
		req.ID,
		req.Slug,
		req.URL,
		req.Title,
		db.NullString(req.Excerpt),
		req.Category,
		stringListPg(req.Badges),
		stringListPg(req.Tags),
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

func (r *PGSQLRepository) Update(ctx context.Context, req UpdateRequest) error {
	if req.ID == "" {
		return ErrInvalidID
	}

	sets := make([]string, 0, 11)
	args := make([]any, 0, 12)

	add := func(col string, val any) {
		args = append(args, val)
		sets = append(sets, col+" = "+ph(len(args)))
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
		add("badges", stringListPg(*req.Badges))
	}
	if req.Tags != nil {
		add("tags", stringListPg(*req.Tags))
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
	q := "UPDATE " + tableName + " SET " + strings.Join(sets, ", ") +
		" WHERE id = " + ph(len(args)-1) + " AND author_id = " + ph(len(args))

	res, err := r.db.ExecContext(ctx, q, args...)
	if err != nil {
		if db.IsDuplicateEntry(err) {
			return ErrDuplicateSlug
		}
		return fmt.Errorf("article: update: %w", err)
	}

	// Postgres menghitung baris yang cocok WHERE, jadi 0 berarti memang tidak ada.
	affected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("article: update rows affected: %w", err)
	}
	if affected == 0 {
		return ErrNotFound
	}
	return nil
}

// ---------- Delete ----------

func (r *PGSQLRepository) Delete(ctx context.Context, req DeleteRequest) error {
	if req.ID == "" {
		return ErrInvalidID
	}

	res, err := r.db.ExecContext(ctx, "DELETE FROM "+tableName+" WHERE id = $1 AND author_id = $2", req.ID, req.AuthorID)
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

func (r *PGSQLRepository) FindByID(ctx context.Context, id string) (*ArticleEntity, error) {
	if id == "" {
		return nil, ErrInvalidID
	}
	return r.findOne(ctx, "id = $1", id)
}

func (r *PGSQLRepository) FindBySlug(ctx context.Context, slug string) (*ArticleEntity, error) {
	return r.findOne(ctx, "slug = $1", slug)
}

func (r *PGSQLRepository) findOne(ctx context.Context, where string, arg any) (*ArticleEntity, error) {
	q := "SELECT " + colsWithoutContent + ", content FROM " + tableName + " WHERE " + where + " LIMIT 1"

	e, err := pgScanArticle(r.db.QueryRowContext(ctx, q, arg))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("article: find: %w", err)
	}
	return e, nil
}

// ---------- FindAll ----------

func (r *PGSQLRepository) FindAll(ctx context.Context, req FindAllRequest) ([]ArticleEntity, int64, error) {
	req.Normalize()

	where, args, err := buildFilterPGQuery(req)
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
	n := len(args)
	q := "SELECT " + colsWithoutContent + ", " + contentCol +
		" FROM " + tableName + where +
		fmt.Sprintf(" ORDER BY %s %s, id ASC LIMIT %s OFFSET %s",
			req.SortBy, req.Order, ph(n+1), ph(n+2))

	listArgs := append(append([]any{}, args...), req.PageSize, req.Offset())

	rows, err := r.db.QueryContext(ctx, q, listArgs...)
	if err != nil {
		return nil, 0, fmt.Errorf("article: find all: %w", err)
	}
	defer rows.Close()

	items := make([]ArticleEntity, 0, req.PageSize)
	for rows.Next() {
		e, err := pgScanArticle(rows)
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

// buildFilterPGQuery membuat klausa WHERE (diawali spasi, atau string kosong) beserta argumennya.
func buildFilterPGQuery(req FindAllRequest) (string, []any, error) {
	conds := make([]string, 0, 8)
	args := make([]any, 0, 8)

	// bind menambahkan argumen dan mengembalikan placeholder-nya.
	bind := func(v any) string {
		args = append(args, v)
		return ph(len(args))
	}

	if s := strings.TrimSpace(req.Search); s != "" {
		// Satu parameter dipakai ulang untuk kedua kolom.
		p := bind("%" + db.EscapeLike(s) + "%")
		conds = append(conds, "(title ILIKE "+p+" OR excerpt ILIKE "+p+")")
	}
	if req.Category != "" {
		conds = append(conds, "category = "+bind(req.Category))
	}
	if req.AuthorID != "" {
		conds = append(conds, "author_id = "+bind(req.AuthorID))
	}
	if len(req.Tags) > 0 {
		// Setara JSON_OVERLAPS: true jika salah satu elemen ada di kolom JSONB.
		conds = append(conds, "jsonb_exists_any(tags, "+bind(pq.Array(req.Tags))+")")
	}
	if len(req.Badges) > 0 {
		conds = append(conds, "jsonb_exists_any(badges, "+bind(pq.Array(req.Badges))+")")
	}
	if req.OnlyPublish {
		conds = append(conds, "published_at IS NOT NULL AND published_at <= "+bind(time.Now()))
	}
	if req.PublishFrom != nil {
		conds = append(conds, "published_at >= "+bind(*req.PublishFrom))
	}
	if req.PublishTo != nil {
		conds = append(conds, "published_at <= "+bind(*req.PublishTo))
	}

	if len(conds) == 0 {
		return "", args, nil
	}
	return " WHERE " + strings.Join(conds, " AND "), args, nil
}

// ---------- ExistsBySlug ----------

func (r *PGSQLRepository) ExistsBySlug(ctx context.Context, slug string, excludeID string) (bool, error) {
	q := "SELECT EXISTS(SELECT 1 FROM " + tableName + " WHERE slug = $1"
	args := []any{slug}
	if excludeID != "" {
		q += " AND id <> $2"
		args = append(args, excludeID)
	}
	q += ")"

	var exists bool
	if err := r.db.QueryRowContext(ctx, q, args...).Scan(&exists); err != nil {
		return false, fmt.Errorf("article: exists by slug: %w", err)
	}
	return exists, nil
}

// ===================== Scan helpers =====================

type pgRowScanner interface {
	Scan(dest ...any) error
}

// pgScanArticle membaca satu baris dengan urutan kolom:
// id, slug, url, title, excerpt, category, badges, tags, cover, author_id,
// published_at, updated_at, content
func pgScanArticle(s pgRowScanner) (*ArticleEntity, error) {
	var (
		e         ArticleEntity
		excerpt   sql.NullString
		cover     sql.NullString
		badges    stringListPg
		tags      stringListPg
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

// stringListPg menyimpan []string ke kolom JSONB dan membacanya kembali.
type stringListPg []string

// Value mengimplementasikan driver.Valuer.
func (l stringListPg) Value() (driver.Value, error) {
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
func (l *stringListPg) Scan(src any) error {
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
		return fmt.Errorf("article: cannot scan %T into stringListPg", src)
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
