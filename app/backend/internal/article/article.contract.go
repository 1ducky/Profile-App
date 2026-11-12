package article

import "context"

// Repository adalah kontrak akses data untuk tabel `article`.
type ArticlaeRepositoryContract interface {
	// Create menyimpan artikel baru dan mengembalikan entity hasil simpan.
	Create(ctx context.Context, req CreateRequest) (*ArticleEntity, error)

	// Update mengubah artikel berdasarkan ID (partial update, hanya field non-nil).
	Update(ctx context.Context, req UpdateRequest) error

	// Delete menghapus artikel berdasarkan ID.
	Delete(ctx context.Context, req DeleteRequest) error

	// FindByID mengambil satu artikel berdasarkan ID.
	FindByID(ctx context.Context, id string) (*ArticleEntity, error)

	// FindBySlug mengambil satu artikel berdasarkan slug.
	FindBySlug(ctx context.Context, slug string) (*ArticleEntity, error)

	// FindAll mengambil daftar artikel dengan filter, sorting, dan pagination.
	// Mengembalikan data dan total seluruh baris yang cocok (untuk pagination).
	FindAll(ctx context.Context, req FindAllRequest) ([]ArticleEntity, int64, error)

	// ExistsBySlug mengecek apakah slug sudah dipakai.
	// ExcludeID opsional, berguna saat update agar artikel sendiri tidak dihitung.
	ExistsBySlug(ctx context.Context, slug string, excludeID string) (bool, error)
}

type ArticleUsecaseContract interface {
	// Create membuat artikel milik AuthorID. AuthorID wajib.
	Create(ctx context.Context, in CreateInput) (*ArticleEntity, error)

	// Update mengubah artikel. AuthorID wajib dan harus pemilik artikel.
	Update(ctx context.Context, in UpdateInput) error

	// Delete menghapus artikel. AuthorID wajib dan harus pemilik artikel.
	Delete(ctx context.Context, authorID string, id string) error

	// GetBySlug membaca artikel. Publik hanya melihat yang sudah terbit,
	// pemilik (ViewerID == AuthorID) juga bisa melihat draft miliknya.
	GetBySlug(ctx context.Context, slug string, viewerID string) (*ArticleEntity, error)
	GetByID(ctx context.Context, ID string, viewerID string) (*ArticleEntity, error)

	// List mengembalikan daftar artikel dengan pagination.
	List(ctx context.Context, in ListInput) (*ListOutput, error)
}

// Repository adalah contract CRUD untuk category.
type CategoryRepositoryContract interface {
	Create(ctx context.Context, e *CategoryEntity) error
	GetByID(ctx context.Context, id string) (*CategoryEntity, error)
	GetAll(ctx context.Context) ([]CategoryEntity, error)
	Update(ctx context.Context, e *CategoryEntity) error
	Delete(ctx context.Context, id string) error
}
