import type { ArticleCardProps } from "../../types/articles";
import { BookmarkButton } from "./BookMark";
import { ArticleThumbnail } from "./Thubnail";


export function ArticleCard({
    article,
    isBookmarked,
    onToggleBookmark,
    onOpen,
    ctaLabel = "Baca artikel",
}: ArticleCardProps) {
    return (
        <article className="group relative flex h-full flex-col overflow-hidden rounded-card border border-line bg-white shadow-card dark:border-night-line dark:bg-night-surface dark:shadow-none">
            <div className="relative">
                <ArticleThumbnail image={article.thumbnail} />
                <div className="absolute right-3 top-3 z-20">
                    <BookmarkButton
                        articleId={article.id}
                        isBookmarked={isBookmarked}
                        onToggle={onToggleBookmark}
                    />
                </div>
            </div>

            <div className="flex flex-1 flex-col gap-1.5 p-3 md:p-4">
                <p className="text-sm text-muted dark:text-night-muted">
                    {article.category.label}
                </p>
                <h3 className="text-lg font-semibold leading-snug tracking-tight md:text-xl">
                    <a
                        href={article.href}
                        onClick={() => onOpen?.(article)}
                        className="outline-none after:absolute after:inset-0 after:z-10 after:rounded-card after:content-[''] focus-visible:after:ring-2 focus-visible:after:ring-primary focus-visible:after:ring-offset-2"
                    >
                        {article.title}
                    </a>
                </h3>
                <p className="leading-normal text-muted dark:text-night-muted">
                    {article.excerpt}
                </p>
                <span
                    aria-hidden="true"
                    className="mt-auto inline-flex items-center gap-2 pt-2 font-medium group-hover:text-primary"
                >
                    {ctaLabel} →
                </span>
            </div>
        </article>
    );
}