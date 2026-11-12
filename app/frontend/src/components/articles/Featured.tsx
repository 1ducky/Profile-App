import { useRef } from "react";
import { useDragSlider } from "../../hooks/useDrag";
import type { Article, FeaturedArticleCardProps } from "../../types/articles";
import { BookmarkButton } from "./BookMark";
import { ArticleThumbnail } from "./Thubnail";


export function FeaturedArticleCard({
    articles,
    isBookmarked,
    onToggleBookmark,
    onOpen,
    ctaLabel = "Baca selengkapnya",
}: FeaturedArticleCardProps) {
    const drag = useRef({ startX: 0, startT: 0, dx: 0, active: false, moved: false });
    const count = articles.length;
    const { viewportRef, index, go, dragging, onKeyDown, trackProps } =
        useDragSlider(count);
    return (
        <div
            role="region"
            aria-roledescription="carousel"
            aria-label="Produk dan karya"
        >
            {/* Parent: overflow hidden */}
            <div
                ref={viewportRef}
                className="-mx-3 overflow-hidden relative rounded-card"
                tabIndex={0}
                onKeyDown={onKeyDown}
            >
                {/* Badge Label "HEADLINE" di pojok kiri atas */}
                <div className="absolute top-4 left-4 z-10 bg-black/80 px-3 py-1 text-xs font-bold tracking-wider text-white uppercase backdrop-blur-sm rounded">
                    Headline
                </div>
                <div
                    className={`flex rela cursor-grab touch-pan-y select-none active:cursor-grabbing ${dragging
                        ? "transition-none"
                        : "transition-transform duration-500 ease-soft motion-reduce:transition-none"
                        }`}
                    style={trackProps.style}
                    onPointerDown={trackProps.onPointerDown}
                    onPointerMove={trackProps.onPointerMove}
                    onPointerUp={trackProps.onPointerUp}
                    onPointerCancel={trackProps.onPointerCancel}
                    onClickCapture={trackProps.onClickCapture}
                >
                    {articles.map((article, i) => (
                        <div
                            key={article.id}
                            role="group"
                            aria-roledescription="slide"
                            aria-label={`${i + 1} dari ${count}`}
                            // slide yang tidak aktif tidak bisa difokus
                            {...(i !== index ? ({ inert: "" } as object) : {})}
                            className="basis-full shrink-0 grow-0 p-3"
                        >

                            <ArticleHeadline key={article.id} article={article} />
                        </div>
                    ))}
                </div>
            </div>

            {/* Kontrol */}
            {count > 1 && (
                <div className="mt-6 flex items-center justify-between">
                    <div className="flex items-center gap-2">
                        {articles.map((s, i) => (
                            <button
                                key={s.id}
                                type="button"
                                onClick={() => go(i)}
                                aria-label={`Ke slide ${i + 1}`}
                                aria-current={i === index}
                                className={`h-2 rounded-full transition-all focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary ${i === index ? "w-6 bg-primary" : "w-2 bg-border hover:bg-muted"
                                    }`}
                            />
                        ))}
                    </div>

                    <div className="flex items-center gap-2">
                        <button
                            type="button"
                            onClick={() => go(index - 1)}
                            disabled={index === 0}
                            aria-label="Slide sebelumnya"
                            className="grid h-9 w-9 place-items-center rounded-full border border-border bg-white text-ink transition-colors hover:border-primary disabled:opacity-40 disabled:hover:border-border focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary"
                        >
                            {/* <Arrow dir="left" /> */}
                            sebelumnya
                        </button>
                        <button
                            type="button"
                            onClick={() => go(index + 1)}
                            disabled={index === count - 1}
                            aria-label="Slide berikutnya"
                            className="grid h-9 w-9 place-items-center rounded-full border border-border bg-white text-ink transition-colors hover:border-primary disabled:opacity-40 disabled:hover:border-border focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary"
                        >
                            selanjutnya
                            {/* <Arrow /> */}
                        </button>
                    </div>
                </div>
            )}
        </div>
    );
}

export default function ArticleHeadline({ article }: { article: Article }) {
    return (
        <article className="group relative flex flex-col overflow-hidden rounded-xl border border-line bg-white shadow-card dark:border-night-line dark:bg-night-surface dark:shadow-none lg:flex-row">
            <div className="relative w-full overflow-hidden lg:flex-[1.2]">
                {/* Thumbnail / Gambar Artikel */}
                <ArticleThumbnail
                    image={article.thumbnail}
                    priority
                    className="aspect-[16/10] w-full object-cover transition-transform duration-500 group-hover:scale-105 lg:aspect-auto lg:h-full"
                />

                {/* Gradient Overlay agar teks judul kontras dan terbaca dengan jelas di atas gambar */}
                <div className="absolute inset-0 bg-gradient-to-t from-black/80 via-black/40 to-black/20 pointer-events-none" />



                {/* Konten Judul & Waktu di bagian bawah gambar */}
                <div className="absolute bottom-0 left-0 w-full p-5 lg:p-8 z-10 flex flex-col gap-2">
                    <h1 className="text-2xl sm:text-3xl lg:text-4xl font-bold text-white leading-snug text-balance">
                        {article.title}
                    </h1>

                    {article.readingTimeMinutes && (
                        <span className="text-sm font-medium text-gray-300">
                            {article.readingTimeMinutes} menit baca
                        </span>
                    )}
                </div>
            </div>
        </article>
    );
}