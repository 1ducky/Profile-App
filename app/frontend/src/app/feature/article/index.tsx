
import type { ArticleField, ArticleListResponse, ArticleResponse } from "../../../contract/article";
import { api } from "../../../lib/api/client";
import { buildArticleQuery, type ArticleListQuery } from "./query";


export async function GetListArticle(filter: ArticleListQuery): Promise<ArticleListResponse> {
    return (await api.get<ArticleListResponse>(`/articles${buildArticleQuery(filter)}`)).data
}
export async function GetBySlug(slug: string): Promise<ArticleResponse> {
    return (await api.get<ArticleResponse>(`/articles/${slug}`)).data
}
export async function GetByID(id: string): Promise<ArticleResponse> {
    return (await api.get<ArticleResponse>(`/article/${id}`)).data
}
export async function PublishArticle(req: ArticleField): Promise<ArticleResponse> {
    return (await api.post<ArticleResponse>(`/articles`, req)).data
}
export async function UpdateArticles(req: ArticleField, articleId: string): Promise<ArticleResponse> {
    return (await api.patch<ArticleResponse>(`/articles/${articleId}`, req)).data
}
export async function DeleteArticles(articleId: string): Promise<{message:string}> {
    return (await api.delete<{message:string}>(`/articles/${articleId}`)).data
}