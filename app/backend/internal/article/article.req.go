package article

import "time"

// Usecase Input
// ===================== Input / Output =====================

type CreateInput struct {
	AuthorID  string // wajib, dari identitas login, bukan dari body request
	Title     string
	URL       string //opsional
	Excerpt   string
	Category  string
	Badges    []string
	Tags      []string
	Cover     string //filekey
	Content   string
	PublishAt *time.Time // nil = draft
}

type UpdateInput struct {
	AuthorID  string // wajib, harus pemilik
	ID        string
	Slug      *string
	URL       *string
	Title     *string
	Excerpt   *string
	Category  *string
	Badges    *[]string
	Tags      *[]string
	Cover     *string
	Content   *string
	PublishAt *time.Time
	// AuthorID artikel sengaja tidak bisa diubah di sini (transfer kepemilikan = usecase terpisah).
}

type ListInput struct {
	ViewerID string // opsional, identitas pemanggil
	AuthorID string // opsional, filter by author
	Search   string
	Category string
	Tags     []string
	Page     int
	PageSize int
}

type ListOutput struct {
	Items    []ArticleEntity
	Total    int64
	Page     int
	PageSize int
}

// ===================== Request body (JSON) =====================

type CreateArticleBody struct {
	Title       string     `json:"title"`
	Excerpt     string     `json:"excerpt"`
	Category    string     `json:"category"`
	Badges      []string   `json:"badges"`
	Tags        []string   `json:"tags"`
	Cover       string     `json:"cover"`
	Content     string     `json:"content"`
	URL         string     `json:"url"`
	PublishedAt *time.Time `json:"published_at"` // RFC3339, null/absen = draft
}

// updateArticleBody memakai pointer: field yang tidak dikirim = tidak diubah.
type UpdateArticleBody struct {
	Title    *string   `json:"title"`
	URL      *string   `json:"url"`
	Excerpt  *string   `json:"excerpt"`
	Category *string   `json:"category"`
	Badges   *[]string `json:"badges"`
	Tags     *[]string `json:"tags"`
	Cover    *string   `json:"cover"`
	Content  *string   `json:"content"`
}

// ===================== Response (JSON) =====================

type ArticleResponse struct {
	ID          string     `json:"id"`
	Slug        string     `json:"slug"`
	URL         string     `json:"url"`
	Title       string     `json:"title"`
	Excerpt     string     `json:"excerpt"`
	Category    string     `json:"category"`
	Badges      []string   `json:"badges"`
	Tags        []string   `json:"tags"`
	Cover       string     `json:"cover"`
	AuthorID    string     `json:"author_id"`
	PublishedAt *time.Time `json:"published_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	Content     string     `json:"content,omitempty"` // kosong di list
}
