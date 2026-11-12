import { useMemo, useState } from "react";
import { ARTICLE_CATEGORIES, type Article, type ArticleCategoryFilter } from "../../types/articles";
import { ArticlesPageLayout } from "./Layout";
import { ArticlesPageHeader } from "./Header";
import { ArticleSearchBar } from "./Search";
import { CategoryFilter } from "./Category";
import { FeaturedArticleCard } from "./Featured";
import { ArticleGrid } from "./CardGrid";


/** Preview statis: state lokal & data mock, belum terhubung service. */
export default function ArticlesDemo({ Articles }: { Articles: Article[] }) {
    const [search, setSearch] = useState("");
    const [article, setArticle] = useState<Article[]>(Articles)
    const [category, setCategory] = useState<ArticleCategoryFilter>("semua");
    const [bookmarked, setBookmarked] = useState<ReadonlySet<string>>(new Set());

    const featured = useMemo(() => article.filter((a) => a.isFeatured), [article]);
    const filtered = useMemo(() => {
        return article.filter((a) => {
            const matchesSearch = a.title.toLowerCase().includes(search.toLowerCase()) && !a.isFeatured;
            const matchesCategory = category === "semua" || a.category.slug === category;
            return matchesSearch && matchesCategory;
        });
    }, [article, search, category]);


    const toggleBookmark = (id: string) =>
        setBookmarked((prev) => {
            const next = new Set(prev);
            next.has(id) ? next.delete(id) : next.add(id);
            return next;
        });

    return (
        <ArticlesPageLayout
            header={
                <ArticlesPageHeader
                    breadcrumb={[
                        { label: "Beranda", href: "/" },
                        { label: "Artikel & Wawasan" },
                    ]}
                    title="Artikel & Wawasan"
                    subtitle="Gagasan, pengalaman, dan perkembangan terbaru."
                />
            }
            toolbar={
                <>
                    <ArticleSearchBar
                        value={search}
                        onChange={setSearch}
                        onClear={() => setSearch("")}
                    />
                    <CategoryFilter
                        categories={[...ARTICLE_CATEGORIES]}
                        active={category}
                        onSelect={setCategory}
                    />
                </>
            }
            featured={featured && (
                <FeaturedArticleCard
                    articles={filtered}
                    isBookmarked={false}
                    onToggleBookmark={toggleBookmark}
                />)}
        >
            <ArticleGrid
                articles={filtered}
                bookmarkedIds={bookmarked}
                onToggleBookmark={toggleBookmark}
            />
        </ArticlesPageLayout>
    );
}