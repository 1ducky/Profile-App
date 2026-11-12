package assets

import (
	"profile-template/internal/db"
	"time"
)

type AssetMetaData struct {
	ID        string        `json:"id"`
	ParentID  string        `json:"parent_id"`
	Filename  string        `json:"filename"`
	FileKey   string        `json:"file_key"`
	Mime      Mime          `json:"mime"`
	Size      int64         `json:"size"`
	AuthorID  string        `json:"author_id"`
	Status    AssetStatus   `json:"status"`
	Category  AssetCategory `json:"category"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

const TABLE db.Table = "assets"

const (
	ID         db.Collom = "id"
	PARENT_ID  db.Collom = "parent_id"
	FILENAME   db.Collom = "filename"
	FILE_KEY   db.Collom = "file_key"
	MIME       db.Collom = "mime"
	SIZE       db.Collom = "size"
	AUTHOR_ID  db.Collom = "author_id"
	STATUS     db.Collom = "status"
	CATEGORY   db.Collom = "category"
	CREATED_AT db.Collom = "created_at"
	UPDATED_AT db.Collom = "updated_at"
)
