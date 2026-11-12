import type { ArticleCardProps } from "../../types/articles";
import { BookmarkButton } from "./BookMark";
import { ArticleThumbnail } from "./Thubnail";

export function ArticleList({
    article,
    isBookmarked,
    onToggleBookmark,
    onOpen,
    ctaLabel = "Baca artikel",
}: ArticleCardProps) {
    return (
        <article className="group relative flex flex-row items-center justify-between gap-4 overflow-hidden bg-white p-4 shadow-card transition-all hover:bg-gray-50/50 dark:border-night-line dark:bg-night-surface dark:shadow-none dark:hover:bg-night-muted/10">

            {/* Bagian Kiri: Kategori, Judul, Eksplanasi, Waktu/CTA */}
            <div className="flex flex-1 flex-col gap-1.5">
                {/* Kategori & Waktu (Opsional jika ingin mirip referensi) */}


                {/* Judul Artikel */}
                <h3 className="text-base font-bold leading-snug tracking-tigh text-blue-400 md:text-lg">
                    <a
                        href={article.href}
                        onClick={() => onOpen?.(article)}
                        className="outline-none after:absolute after:inset-0 after:z-10 after:content-[''] focus-visible:after:ring-2 focus-visible:after:ring-primary"
                    >
                        {article.title}
                    </a>
                </h3>
                <div className="flex items-center gap-2 text-xs font-semibold uppercase tracking-wider text-primary">
                    <span>{article.category.label}</span>
                    {article.readingTimeMinutes && (
                        <>
                            <span className="text-muted">•</span>
                            <span className="text-muted font-normal normal-case">
                                {article.readingTimeMinutes} Menit
                            </span>
                        </>
                    )}
                </div>

                {/* Eksplanasi Singkat (Opsional, bisa dihapus jika ingin super minimalis seperti gambar) */}
                {/* {article.excerpt && (
                    <p className="line-clamp-2 text-sm leading-normal text-muted dark:text-night-muted">
                        {article.excerpt}
                    </p>
                )} */}
            </div>

            {/* Bagian Kanan: Thumbnail Gambar & Tombol Bookmark */}
            <div className="relative flex-shrink-0">
                <div className="relative h-20 w-28 overflow-hidden rounded-md sm:h-24 sm:w-36 md:h-28 md:w-44">
                    <ArticleThumbnail
                        image={article.thumbnail}
                        className="size-full object-cover transition-transform duration-300 group-hover:scale-105"
                    />
                </div>

                {/* Tombol Bookmark di atas Thumbnail (Opsional) */}
                {/* {onToggleBookmark && (
                    <div className="absolute right-2 top-2 z-20">
                        <BookmarkButton
                            articleId={article.id}
                            isBookmarked={isBookmarked}
                            onToggle={onToggleBookmark}
                        />
                    </div>
                )} */}
            </div>

        </article>
    );
}