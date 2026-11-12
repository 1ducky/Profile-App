package article

import (
	"profile-template/internal/db"
	"time"
)

type ArticleEntity struct {
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
	PublishAt time.Time
	UpdateAt  time.Time
	Content   string
}

const TABLE db.Table = "article"

const (
	ID          db.Collom = "id"
	Slug        db.Collom = "slug"
	URL         db.Collom = "url"
	Title       db.Collom = "title"
	Excerpt     db.Collom = "excerpt"
	Category    db.Collom = "category"
	Badges      db.Collom = "badges"
	Tags        db.Collom = "tags"
	Cover       db.Collom = "cover"
	AuthorID    db.Collom = "author_id"
	PublishedAt db.Collom = "published_at"
	UpdatedAt   db.Collom = "updated_at"
	CreatedAt   db.Collom = "created_at"
	Content     db.Collom = "content"
)

type CategoryEntity struct {
	ID   string `json:"id"`
	Slug string `json:"slug"`
}

const CATEGORY_TABLE db.Table = "article_category"
const (
	CATEGORY_ID   db.Collom = "id"
	CATEGORY_Slug db.Collom = "slug"
)
