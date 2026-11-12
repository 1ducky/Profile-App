package article

import (
	"context"
	"log"
	"path/filepath"
	assetsApp "profile-template/internal/app/services/assets"
	"profile-template/internal/app/uow"
	"profile-template/internal/article"
	"profile-template/internal/assets"
	"profile-template/internal/profile"
	"profile-template/internal/utils"
	"strings"
	"time"
)

type TXUsecase struct {
	Article article.ArticleUsecaseContract
	Asset   assets.UsecaseContract
}

type ArticleApp struct {
	uow     uow.UnitOfWork[TXUsecase]
	artc    article.ArticleUsecaseContract
	profile profile.UsecaseContract
}

func NewArticleApp(artc article.ArticleUsecaseContract, profle profile.UsecaseContract, uow uow.UnitOfWork[TXUsecase]) *ArticleApp {
	return &ArticleApp{artc: artc, profile: profle, uow: uow}
}

func (a *ArticleApp) CreateArticle(ctx context.Context, req article.CreateInput) (ArticleResponse, error) {
	if req.AuthorID == "" {
		return ArticleResponse{}, article.ErrUnauthenticated
	}
	filekeys := article.ExtractFileKeys(req.Content)
	filekeys = append(filekeys, req.Cover)

	filekeys = utils.MapField(filekeys, func(key string) (string, bool) {
		if key == "" {
			return "", false
		}
		return a.clearToValidFileKey(key), true
	})
	log.Println("filekeys:", filekeys)
	var articleRes article.ArticleEntity
	err := a.uow.Do(ctx, func(ctx context.Context, r TXUsecase) error {
		cratedAt := time.Now()
		res, err := r.Article.Create(ctx, article.CreateInput{
			AuthorID:  req.AuthorID,
			Title:     req.Title,
			URL:       req.URL,
			Content:   req.Content,
			Cover:     req.Cover,
			Excerpt:   req.Excerpt,
			Category:  req.Category,
			Badges:    req.Badges,
			Tags:      req.Tags,
			PublishAt: &cratedAt,
		})
		if err != nil {
			return err
		}
		articleRes = *res
		if len(filekeys) == 0 {
			return nil
		}
		err = r.Asset.LinkingAssetByFileKeys(ctx, filekeys, res.ID)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return ArticleResponse{}, err
	}
	author, err := a.profile.GetUserByID(ctx, articleRes.AuthorID)
	if err != nil {
		log.Println("error get author:", err)
		log.Println("error get author:", articleRes.AuthorID)
		return ArticleResponse{}, err
	}

	return ArticleResponse{
		ArticleResponse: article.ToResponse(&articleRes),
		ProfileResource: ProfileResource{
			Username: author.Username,
			Tag:      author.Tag,
			Avatar:   author.AvatarURL,
			Bio:      author.Bio,
		},
	}, err

}

func (a *ArticleApp) UpdateArticle(ctx context.Context, req article.UpdateInput) error {
	if req.AuthorID == "" {
		return article.ErrUnauthenticated
	}
	newFileKeys := article.ExtractFileKeys(*req.Content)
	newFileKeys = append(newFileKeys, *req.Cover)
	newFileKeys = utils.MapField(newFileKeys, func(key string) (string, bool) {
		if key == "" {
			return "", false
		}
		return a.clearToValidFileKey(key), true
	})
	log.Print(newFileKeys)
	return a.uow.Do(ctx, func(ctx context.Context, r TXUsecase) error {
		lastArticle, err := r.Article.GetByID(ctx, req.ID, req.AuthorID)
		if err != nil {
			return article.ErrNotFound
		}

		err = r.Article.Update(ctx, req)
		if err != nil {
			return err
		}
		err = r.Asset.ReleaseByParentID(ctx, []string{lastArticle.ID})
		if err != nil && len(newFileKeys) > 0 {
			return err
		}
		if len(newFileKeys) == 0 {
			return nil
		}
		err = r.Asset.LinkingAssetByFileKeys(ctx, newFileKeys, lastArticle.ID)
		if err != nil {
			return err
		}
		return nil
	})

}

func (a *ArticleApp) DeleteArticle(ctx context.Context, req article.DeleteRequest) error {
	if req.AuthorID == nil {
		return article.ErrUnauthenticated
	}
	return a.uow.Do(ctx, func(ctx context.Context, r TXUsecase) error {
		err := r.Article.Delete(ctx, *req.AuthorID, req.ID)
		if err != nil {
			return err
		}
		err = r.Asset.ReleaseByParentID(ctx, []string{req.ID})
		if err != nil {
			return err
		}
		return nil
	})
}

func (a *ArticleApp) GetBySlugArticle(ctx context.Context, userID string, slug string) (ArticleResponse, error) {
	res, err := a.artc.GetBySlug(ctx, slug, userID)
	if err != nil {
		return ArticleResponse{}, err
	}
	author, err := a.profile.GetUserByID(ctx, res.AuthorID)
	if err != nil {
		log.Println("error get author:", err)
		return ArticleResponse{}, err
	}

	return ArticleResponse{
		ArticleResponse: article.ToResponse(res),
		ProfileResource: ProfileResource{
			Username: author.Username,
			Tag:      author.Tag,
			Avatar:   author.AvatarURL,
			Bio:      author.Bio,
		},
	}, err
}
func (a *ArticleApp) GetByIdArticle(ctx context.Context, userID string, id string) (ArticleResponse, error) {
	res, err := a.artc.GetByID(ctx, id, userID)
	if err != nil {
		return ArticleResponse{}, err
	}
	author, err := a.profile.GetUserByID(ctx, res.AuthorID)
	if err != nil {
		log.Println("error get author:", err)
		return ArticleResponse{}, err
	}

	return ArticleResponse{
		ArticleResponse: article.ToResponse(res),
		ProfileResource: ProfileResource{
			Username: author.Username,
			Tag:      author.Tag,
			Avatar:   author.AvatarURL,
			Bio:      author.Bio,
		},
	}, err
}

func (a *ArticleApp) GetListArticle(ctx context.Context, req article.ListInput) (ListArticleResponses, error) {
	res, err := a.artc.List(ctx, req)
	if err == article.ErrNotFound {
		return ListArticleResponses{
			Articles: []ArticleResponse{},
			Meta: article.ListMeta{
				Total:    0,
				Page:     req.Page,
				PageSize: req.PageSize,
			},
		}, err
	} else if err != nil {
		return ListArticleResponses{}, err
	}
	responses := ListArticleResponses{
		Articles: make([]ArticleResponse, len(res.Items)),
	}

	authorIDs := make([]string, len(res.Items))
	for i := range res.Items {
		responses.Articles[i].ArticleResponse = article.ToResponse(&res.Items[i])
		authorIDs[i] = res.Items[i].AuthorID
	}
	authorIDs = utils.Unique(authorIDs)
	authors, err := a.profile.GetUsersByIDs(ctx, authorIDs)
	if err != nil {
		return ListArticleResponses{}, err
	}
	authorMap := make(map[string]ProfileResource)
	for _, author := range authors {
		authorMap[author.ID] = ProfileResource{
			Username: author.Username,
			Tag:      author.Tag,
			Avatar:   author.AvatarURL,
			Bio:      author.Bio,
		}
	}
	for i, item := range responses.Articles {
		if author, ok := authorMap[item.ArticleResponse.AuthorID]; ok {
			responses.Articles[i].ProfileResource = author
		}
	}

	Metadata := article.ListMeta{
		Total:    res.Total,
		Page:     res.Page,
		PageSize: res.PageSize,
	}
	return ListArticleResponses{
		Articles: responses.Articles,
		Meta:     Metadata,
	}, nil
}

func (a *ArticleApp) clearToValidFileKey(filekey string) string {
	return filepath.FromSlash(strings.TrimPrefix(filekey, assetsApp.PREFIX))
}
