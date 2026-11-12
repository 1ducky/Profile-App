import type { ProfileResponse } from "./profile";

export interface ArticleResponse {
  id: string;
  slug: string;
  url: string;
  title: string;
  excerpt: string;
  category: string;
  badges: string[];
  tags: string[];
  cover: string; //prefix with asset:
  author: ProfileResponse;
  published_at: string | null;
  updated_at: string | null;
  content?: string;
}

export interface ArticleListResponse {
  articles: ArticleResponse[];
  meta: {
    total: number;
    page: number;
    page_size: number;
  };
}

export interface ArticleField {
  title: string,
  excerpt: string,
  category: string,
  badges: string[],
  tags: string[],
  cover: string,
  content: string,
  url: string
}