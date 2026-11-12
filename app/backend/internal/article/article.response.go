package article

type ListResponse struct {
	Articles []ArticleResponse `json:"articles"`
	Meta     ListMeta          `json:"meta"`
}

type ListMeta struct {
	Total    int64 `json:"total"`
	Page     int   `json:"page"`
	PageSize int   `json:"page_size"`
}

func ToResponse(a *ArticleEntity) ArticleResponse {
	return ArticleResponse{
		ID:          a.ID,
		Slug:        a.Slug,
		URL:         a.URL,
		Title:       a.Title,
		Excerpt:     a.Excerpt,
		Category:    a.Category,
		Badges:      nonNil(a.Badges),
		Tags:        nonNil(a.Tags),
		Cover:       a.Cover,
		AuthorID:    a.AuthorID,
		PublishedAt: timePtr(a.PublishAt),
		UpdatedAt:   timePtr(a.UpdateAt),
		Content:     a.Content,
	}
}
