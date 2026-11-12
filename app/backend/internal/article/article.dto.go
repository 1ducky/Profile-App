package article

import "time"

// reposiotry

// CreateRequest adalah data yang dibutuhkan untuk membuat artikel.
type CreateRequest struct {
	ID        string
	Slug      string
	URL       string
	Title     string
	Excerpt   string
	Category  string
	Badges    []string
	Tags      []string
	Cover     string
	AuthorID  string
	Content   string
	PublishAt *time.Time // nil = draft / belum dipublikasi
}

// UpdateRequest untuk partial update.
// Field bertipe pointer: nil berarti tidak diubah.
type UpdateRequest struct {
	ID        string // wajib
	Slug      *string
	URL       *string
	Title     *string
	Excerpt   *string
	Category  *string
	Badges    *[]string
	Tags      *[]string
	Cover     *string
	AuthorID  *string
	Content   *string
	PublishAt *time.Time
	UpdateAt  time.Time // diisi service, biasanya time.Now()
}

type DeleteRequest struct {
	ID       string
	AuthorID *string
}

// SortField adalah kolom yang boleh dipakai untuk sorting (whitelist).
type SortField string

const (
	SortByPublishedAt SortField = "published_at"
	SortByUpdatedAt   SortField = "updated_at"
	SortByCreatedAt   SortField = "created_at"
	SortByTitle       SortField = "title"
)

// SortOrder arah sorting.
type SortOrder string

const (
	OrderAsc  SortOrder = "ASC"
	OrderDesc SortOrder = "DESC"
)
