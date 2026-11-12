import type { ArticleListQuery } from "./query";

export function parseArticleQuery(params: URLSearchParams): ArticleListQuery {
    const num = (key: string) => {
        const v = params.get(key);
        if (v === null || v === "") return undefined;
        const n = Number(v);
        return Number.isInteger(n) ? n : undefined;
    };

    return {
        search: params.get("q") ?? undefined,
        category: params.get("category") ?? undefined,
        authorId: params.get("author_id") ?? undefined,
        tags: params.getAll("tag"),
        page: num("page"),
        pageSize: num("page_size"),
    };
}