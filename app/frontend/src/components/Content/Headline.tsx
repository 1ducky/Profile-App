import { useRef } from "react";
import { useDragSlider } from "../../hooks/useDrag";
import { ArticleThumbnail } from "./Thubnail";
import type { ArticleResponse } from "../../contract/article";
import { ResolveUrlAsset } from "../../utils/resolve-url";
import { timeAgo } from "../../utils/time";
import ArticleMiniCard from "./MiniCard";
import { Calendar } from "lucide-react";
import { Hash, Search, X } from "lucide-react";
import { useState } from "react";
import type { NavItem } from "../../layouts/Article/Header.astro";

export default function HeadlineSection({ article }: { article: ArticleResponse[] }) {
    const drag = useRef({ startX: 0, startT: 0, dx: 0, active: false, moved: false });
    const count = article.length;
    const { viewportRef, index, go, dragging, onKeyDown, trackProps } =
        useDragSlider(count);
    return (
        <section className="w-full max-w-[1200px] mx-auto bg-white pb-10">

            {/* 1. Label Headline */}


            {/* 2. Hero Article (Gambar Besar) */}
            <div className="relative w-full aspect-3/2 md:aspect-2/1 overflow-hidden group cursor-pointer mb-4 md:rounded-2xl">
                <div className="absolute top-0 left-0 flex z-2">
                    <div className="bg-[#1a1a1a] text-white px-5 py-1.5 text-[18px] font-bold tracking-wide uppercase">
                        Headline
                    </div>
                </div>
                <div
                    ref={viewportRef}
                    className="-mx-3 overflow-hidden h-full"
                    tabIndex={0}
                    onKeyDown={onKeyDown}
                >
                    {/* Track: children memanjang ke samping */}
                    <div
                        className={`flex cursor-grab touch-pan-y select-none active:cursor-grabbing h-full ${dragging
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
                        {article.map((article, i) => (
                            <article
                                key={article.id}
                                role="group"
                                aria-roledescription="slide"
                                aria-label={`${i + 1} dari ${count}`}
                                {...(i !== index ? ({ inert: "" } as object) : {})}
                                // 1. Tambahkan relative & h-full. Hapus p-3 jika ingin gambar mentok ujung.
                                // Jika ingin ada jarak antar slide, bungkus konten di dalam div dengan padding, 
                                // tapi biarkan article-nya full.
                                className="relative basis-full shrink-0 grow-0 h-full overflow-hidden"
                            >
                                {/* 2. Bungkus gambar dan gradient dalam satu container relative/absolute */}


                                {/* Gradient Overlay */}
                                <div className="absolute inset-0 z-10 bg-gradient-to-t from-black/90 via-black/40 to-transparent pointer-events-none"></div>

                                {/* Thumbnail Gambar */}
                                <ArticleThumbnail
                                    image={{ alt: article.title, src: ResolveUrlAsset(article.cover), caption: "Article Headline" }}
                                    priority={true}
                                    // 3. Pastikan gambar memenuhi container (absolute inset-0 atau w-full h-full object-cover)
                                    className="absolute inset-0 w-full h-full object-cover z-0"
                                />

                                {/* Konten Teks (Judul & Tanggal) */}
                                <div className="absolute bottom-0 left-0 p-6 md:p-10 w-full flex flex-col justify-end text-white z-20">
                                    <h2 className="text-2xl md:text-4xl lg:text-5xl font-extrabold leading-tight tracking-tight drop-shadow-lg ">
                                        {article.title}
                                    </h2>
                                    <span className="text-md font-medium text-gray-300">
                                        {article.category.toUpperCase()}
                                    </span>
                                    <span className="text-sm font-medium text-text-muted flex flex-row flex-wrap gap-3">
                                        {article.badges.map((badge, i) => (
                                            <p className="px-2 py-1 rounded-2xl bg-surface">#{badge}</p>
                                        ))}

                                    </span>
                                    <span className="text-sm font-medium text-gray-300 flex flex-row flex-nowrap items-center gap-2">
                                        <Calendar /> {timeAgo(article.published_at || "")} yang lalu
                                    </span>
                                </div>
                            </article>
                        ))}
                    </div>
                </div>

            </div>

            {/* 3. Stream Articles (Grid 5 Kolom di Desktop) */}
            <div className="lg:grid hidden grid-cols-2 md:grid-cols-3 lg:grid-cols-5 gap-4 px-2">
                {article.map((article) => (

                    <ArticleMiniCard article={article} className="h-full" />
                ))}
            </div>

        </section>
    );
}



export interface LinkItem {
    label: string;
    path: string;
}

interface ArticleHeaderProps {
    title?: string;
    description?: string;
    categories: NavItem[];
    tags: NavItem[];
    /** Path yang sedang aktif (kategori atau tag) untuk penanda. */
    activePath?: string;
    /** Nilai awal kolom pencarian, mis. dari query string. */
    initialQuery?: string;
    /** Tujuan form pencarian (method GET, parameter `q`). */
    searchAction?: string;
    /** Opsional: tangani pencarian di client tanpa reload halaman. */
    onSearch?: (query: string) => void;
    className?: string;
}

export function ArticleHeaderNew({
    title = "Artikel & Insights",
    description = "Tulisan, opini, dan cerita kegiatan terbaru seputar teknologi dan dampaknya bagi masyarakat.",
    categories,
    tags,
    activePath,
    initialQuery = "",
    searchAction = "/articles",
    onSearch,
    className = "",
}: ArticleHeaderProps) {
    const [query, setQuery] = useState(initialQuery);

    const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        if (onSearch) {
            e.preventDefault();
            onSearch(query.trim());
        }
    };

    return (
        <header className={`w-full ${className}`}>
            <div className="mx-auto flex max-w-[1200px] flex-col gap-8 px-4 py-14 md:py-20">
                {/* Judul + pencarian */}
                <div className="flex flex-col gap-6 md:flex-row md:items-end md:justify-between">
                    <div className="max-w-xl">
                        <p className="text-xs font-semibold uppercase tracking-[0.14em] text-primary bg-white p-3">
                            Harian Cak Oyim
                        </p>
                        <h1 className="mt-2 text-3xl font-bold tracking-tight text-white md:text-4xl">
                            {title}
                        </h1>
                        <p className="mt-3 leading-relaxed text-white">{description}</p>
                    </div>
                    <div className="flex flex-col flex-nowrap gap-5">


                        <form
                            role="search"
                            action={searchAction}
                            method="get"
                            onSubmit={handleSubmit}
                            className="w-full "
                        >
                            <label htmlFor="article-search" className="sr-only">
                                Cari artikel
                            </label>
                            <div className="relative">
                                <Search
                                    className="pointer-events-none absolute left-4 top-1/2 h-4 w-4 -translate-y-1/2 text-[#6B7280]"
                                    aria-hidden="true"
                                />
                                <input
                                    id="article-search"
                                    name="q"
                                    type="search"
                                    value={query}
                                    onChange={(e) => setQuery(e.target.value)}
                                    placeholder="Cari judul atau topik…"
                                    autoComplete="off"
                                    className="h-12 w-full rounded-full border border-gray-200 bg-white pl-11 pr-24 text-sm text-[#1F2937] shadow-[0_4px_16px_rgba(31,41,55,0.06)] outline-none transition-colors placeholder:text-[#6B7280] focus:border-[#0B63CE] focus:ring-2 focus:ring-[#0B63CE]/30 [&::-webkit-search-cancel-button]:hidden"
                                />
                                {query && (
                                    <button
                                        type="button"
                                        onClick={() => setQuery("")}
                                        aria-label="Hapus pencarian"
                                        className="absolute right-[4.5rem] top-1/2 flex h-7 w-7 -translate-y-1/2 items-center justify-center rounded-full text-[#6B7280] hover:bg-gray-100 focus-visible:outline-2 focus-visible:outline-[#0B63CE]"
                                    >
                                        <X className="h-4 w-4" aria-hidden="true" />
                                    </button>
                                )}
                                <button
                                    type="submit"
                                    className="absolute right-1.5 top-1/2 h-9 -translate-y-1/2 rounded-full bg-[#0B63CE] px-4 text-sm font-medium text-white transition-colors hover:bg-[#0A56B3] focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-[#0B63CE]"
                                >
                                    Cari
                                </button>
                            </div>
                        </form>
                        {tags.length > 0 && (
                            <nav
                                aria-label="Tag artikel"
                                className="flex flex-wrap items-center gap-x-2 gap-y-2"
                            >
                                <span className="mr-1 flex items-center gap-1 text-sm text-white">
                                    <Hash className="h-4 w-4" aria-hidden="true" />
                                    Tag
                                </span>
                                {tags.map((tag) => {
                                    const active = activePath === tag.path;
                                    return (
                                        <a
                                            key={tag.path}
                                            href={"/articles?tag=" + tag.path}
                                            aria-current={active ? "page" : undefined}
                                            className={`rounded-full px-3 py-1 text-xs font-medium transition-colors focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-[#0B63CE] ${active
                                                ? "bg-white text-primary"
                                                : "bg-primary  text-white hover:bg-primary-hover"
                                                }`}
                                        >
                                            #{tag.label}
                                        </a>
                                    );
                                })}
                            </nav>
                        )}
                    </div>
                </div>

                {/* Kategori */}
                {categories.length > 0 && (
                    <nav aria-label="Kategori artikel">
                        <ul className="-mx-4 flex gap-2 overflow-x-auto px-4 pb-1 md:mx-0 md:flex-wrap md:overflow-visible md:px-0 md:pb-0">
                            {categories.map((cat) => {
                                const active = activePath === cat.path;
                                return (
                                    <li key={cat.path} className="shrink-0">
                                        <a
                                            href={"/articles?category=" + cat.path}
                                            aria-current={active ? "page" : undefined}
                                            className={`inline-flex h-9 items-center rounded-full px-4 text-sm font-medium transition-colors focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-[#0B63CE] ${active
                                                ? "bg-white text-primary"
                                                : "border border-gray-200 bg-primary text-white hover:border-[#0B63CE] hover:text-[#0B63CE]"
                                                }`}
                                        >
                                            {cat.label}
                                        </a>
                                    </li>
                                );
                            })}
                        </ul>
                    </nav>
                )}

                {/* Tag */}

            </div>
        </header>
    );
}
