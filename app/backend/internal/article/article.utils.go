package article

import "time"

// FindAllRequest adalah filter + pagination untuk daftar artikel.
type FindAllRequest struct {
	// Filter
	Search      string     // cari di title/excerpt
	Category    string     // exact match
	AuthorID    string     // exact match
	Tags        []string   // artikel yang memiliki salah satu tag
	Badges      []string   // artikel yang memiliki salah satu badge
	OnlyPublish bool       // true = hanya yang published_at <= now
	PublishFrom *time.Time // rentang published_at (awal)
	PublishTo   *time.Time // rentang published_at (akhir)

	// Sorting
	SortBy SortField // default: published_at
	Order  SortOrder // default: DESC

	// Pagination
	Page     int // mulai dari 1
	PageSize int // default 10, maksimal 100

	// IncludeContent false = kolom content tidak di-select (lebih ringan untuk list).
	IncludeContent bool
}

// Normalize mengisi nilai default dan membatasi rentang nilai.
func (r *FindAllRequest) Normalize() {
	if r.Page < 1 {
		r.Page = 1
	}
	if r.PageSize < 1 {
		r.PageSize = 10
	}
	if r.PageSize > 100 {
		r.PageSize = 100
	}
	switch r.SortBy {
	case SortByPublishedAt, SortByUpdatedAt, SortByCreatedAt, SortByTitle:
	default:
		r.SortBy = SortByPublishedAt
	}
	if r.Order != OrderAsc {
		r.Order = OrderDesc
	}
}

// Offset menghitung offset untuk query LIMIT/OFFSET.
func (r FindAllRequest) Offset() int {
	return (r.Page - 1) * r.PageSize
}
