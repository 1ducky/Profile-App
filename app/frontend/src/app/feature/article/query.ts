export interface ArticleListQuery {
    page?: number;
    pageSize?: number;
    authorId?: string;
    search?: string;
    category?: string;
    tags?: string[];
}

export function buildArticleQuery(input: ArticleListQuery = {}): string {
    const params = new URLSearchParams();

    const set = (key: string, value: string | number | undefined | null) => {
        if (value === undefined || value === null) return;
        const str = String(value).trim();
        if (str === "") return;
        params.set(key, str);
    };

    set("q", input.search);
    set("category", input.category);
    set("author_id", input.authorId);

    // tag berulang: ?tag=a&tag=b (cocok dengan q["tag"] di Go)
    for (const tag of input.tags ?? []) {
        const t = tag.trim();
        if (t) params.append("tag", t);
    }

    set("page", input.page);
    set("page_size", input.pageSize);

    const qs = params.toString();
    return qs ? `?${qs}` : "";
}