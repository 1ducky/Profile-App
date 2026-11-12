package article

import (
	"profile-template/internal/article"
)

type ArticleResponse struct {
	article.ArticleResponse
	ProfileResource `json:"author"`
}

type ProfileResource struct {
	Username string `json:"username"`
	Tag      string `json:"tag"`
	Avatar   string `json:"avatar"`
	Bio      string `json:"bio"`
}
type ListArticleResponses struct {
	Articles []ArticleResponse `json:"articles"`
	Meta     article.ListMeta  `json:"meta"`
}
