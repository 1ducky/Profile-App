package article

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type usecase struct {
	repo ArticlaeRepositoryContract
}

func NewUsecase(repo ArticlaeRepositoryContract) ArticleUsecaseContract {
	return &usecase{repo: repo}
}

// ---------- Create ----------

func (uc *usecase) Create(ctx context.Context, in CreateInput) (*ArticleEntity, error) {
	if in.AuthorID == "" {
		return nil, ErrUnauthenticated
	}
	title := strings.TrimSpace(in.Title)
	if title == "" {
		return nil, errors.New("title cannot be empty")
	}
	slug := slugify(in.Title)

	if strings.TrimSpace(in.Content) == "" && in.URL == "" {
		return nil, errors.New("content cannot be empty without URL")
	}
	if strings.TrimSpace(in.Category) == "" {
		return nil, errors.New("title cannot be empty without category")
	}

	taken, err := uc.repo.ExistsBySlug(ctx, slug, "")
	if err != nil {
		return nil, err
	}
	if taken {
		return nil, ErrDuplicateSlug
	}

	// Jika terjadi race, repository tetap mengembalikan ErrDuplicateSlug.
	return uc.repo.Create(ctx, CreateRequest{
		ID:        uuid.NewString(),
		Slug:      slug,
		URL:       in.URL,
		Title:     title,
		Excerpt:   strings.TrimSpace(in.Excerpt),
		Category:  strings.TrimSpace(in.Category),
		Badges:    in.Badges,
		Tags:      in.Tags,
		Cover:     in.Cover,
		AuthorID:  in.AuthorID,
		Content:   in.Content,
		PublishAt: in.PublishAt,
	})
}

// ---------- Update ----------

func (uc *usecase) Update(ctx context.Context, in UpdateInput) error {
	if in.AuthorID == "" {
		return ErrUnauthenticated
	}
	fmt.Println("usecase:")
	fmt.Println("url:", *in.URL)
	fmt.Println("content:", *in.Content)

	if in.ID == "" {
		return ErrInvalidID
	}
	slug := slugify(*in.Title)

	req := UpdateRequest{
		ID:        in.ID,
		Excerpt:   in.Excerpt,
		URL:       in.URL,
		Slug:      &slug,
		Cover:     in.Cover,
		Badges:    in.Badges,
		Tags:      in.Tags,
		PublishAt: in.PublishAt,
		UpdateAt:  time.Now(),
		Category:  in.Category,
		AuthorID:  &in.AuthorID,
	}

	if in.Title != nil {
		t := strings.TrimSpace(*in.Title)
		if t == "" {
			return ErrBadRequest
		}
		req.Title = &t
	}
	if in.Category != nil {
		if strings.TrimSpace(*in.Category) == "" {
			return ErrBadRequest
		}
		req.Category = in.Category
	}
	if in.Content != nil {
		if strings.TrimSpace(*in.Content) == "" && in.URL == nil {
			return ErrBadRequest
		}
		req.Content = in.Content
	}

	// if in.Slug != nil {
	// 	// s := slugify(*in.Slug)
	// 	// if s == "" {
	// 	// 	return invalid("slug cannot be empty")
	// 	// }
	// 	if slug != "" {
	// 		taken, err := uc.repo.ExistsBySlug(ctx, slug, in.ID)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		if taken {
	// 			return ErrDuplicateSlug
	// 		}
	// 		// u := articleURL(s)
	// 		// req.Slug = &s
	// 		// req.URL = &u // URL harus ikut berubah saat slug berubah
	// 	}
	// }
	fmt.Println("usecase req:", *req.URL)
	return uc.repo.Update(ctx, req)
}

// ---------- Delete ----------

func (uc *usecase) Delete(ctx context.Context, authorID string, id string) error {
	if authorID == "" {
		return ErrUnauthenticated
	}
	if id == "" {
		return ErrInvalidID
	}
	// if _, err := uc.loadOwned(ctx, id, authorID); err != nil {
	// 	return err
	// }
	return uc.repo.Delete(ctx, DeleteRequest{ID: id, AuthorID: &authorID})
}

// ---------- GetBySlug ----------

func (uc *usecase) GetBySlug(ctx context.Context, slug string, viewerID string) (*ArticleEntity, error) {
	a, err := uc.repo.FindBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}

	// Draft / belum waktunya terbit hanya terlihat oleh pemiliknya.
	// Selain itu pura-pura tidak ada agar keberadaan draft tidak bocor.
	if !uc.isPublished(a) && (viewerID == "" || viewerID != a.AuthorID) {
		return nil, ErrNotFound
	}
	return a, nil
}

// ---------- GetBySlug ----------

func (uc *usecase) GetByID(ctx context.Context, ID string, viewerID string) (*ArticleEntity, error) {
	a, err := uc.repo.FindByID(ctx, ID)
	if err != nil {
		return nil, err
	}
	if !uc.isPublished(a) && (viewerID == "" || viewerID != a.AuthorID) {
		return nil, ErrNotFound
	}
	return a, nil
}

// ---------- List ----------

func (uc *usecase) List(ctx context.Context, in ListInput) (*ListOutput, error) {
	// Melihat daftar milik sendiri = boleh melihat draft. Selain itu hanya yang terbit.
	ownList := in.AuthorID != "" && in.AuthorID == in.ViewerID

	req := FindAllRequest{
		Search:      in.Search,
		Category:    in.Category,
		AuthorID:    in.AuthorID,
		Tags:        in.Tags,
		OnlyPublish: !ownList,
		SortBy:      SortByPublishedAt,
		Order:       OrderDesc,
		Page:        in.Page,
		PageSize:    in.PageSize,
	}
	req.Normalize()

	items, total, err := uc.repo.FindAll(ctx, req)
	if err != nil {
		return nil, err
	}
	return &ListOutput{Items: items, Total: total, Page: req.Page, PageSize: req.PageSize}, nil
}
