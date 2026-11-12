import { ArrowUpRight, CalendarDays, ImageOff } from "lucide-react";
import { useState } from "react";
import type { ArticleResponse } from "../../contract/article";
import { ArticleThumbnail } from "./Thubnail";
import { ResolveUrlAsset } from "../../utils/resolve-url";
import { timeAgo } from "../../utils/time";

// Sesuaikan dengan definisi ProfileResponse di project kamu.
// Di sini hanya field minimal yang dipakai komponen.


interface ArticleMiniCardProps {
    article: ArticleResponse;
    /** Ubah "asset:xxx" menjadi URL gambar yang bisa dimuat. */
    resolveAsset?: (path: string) => string;
    className?: string;
}



export function ArticleMiniCard({
    article,

    className = "",
}: ArticleMiniCardProps) {


    return (
        <article
            className={`group relative flex flex-col overflow-hidden rounded-2xl border border-gray-200/70 bg-white shadow-[0_4px_16px_rgba(31,41,55,0.06)] transition-shadow focus-within:ring-2 focus-within:ring-[#0B63CE] focus-within:ring-offset-2 hover:shadow-[0_8px_24px_rgba(31,41,55,0.10)] motion-reduce:transition-none ${className}`}
        >
            {/* Thumbnail */}
            <div className="relative aspect-[16/10] w-full overflow-hidden bg-[#F5EFE3]">
                <ArticleThumbnail image={{ alt: article.title, src: ResolveUrlAsset(article.cover) }} />

                {article.category && (
                    <span className="absolute left-3 top-3 rounded-full bg-white/95 px-3 py-1 text-xs font-medium text-[#0B63CE] shadow-sm">
                        {article.category}
                    </span>
                )}
            </div>

            {/* Body */}
            <div className="flex flex-1 flex-col gap-2 p-4">
                {article.published_at && (
                    <time
                        dateTime={article.published_at ?? undefined}
                        className="flex items-center gap-1.5 text-xs text-[#6B7280]"
                    >
                        <CalendarDays className="h-3.5 w-3.5" aria-hidden="true" />
                        {timeAgo(article.published_at || "")}
                    </time>
                )}

                <h3 className="line-clamp-2 text-base font-semibold leading-snug tracking-tight text-[#1F2937]">
                    {/* Stretched link: seluruh kartu bisa diklik */}
                    <a
                        href={article.url != "" ? article.url : "/articles/" + article.slug}
                        className="outline-none before:absolute before:inset-0 before:content-['']"
                    >
                        {article.title}
                    </a>
                </h3>

                <p className="line-clamp-2 text-sm leading-relaxed text-[#6B7280]">
                    {article.excerpt}
                </p>

                {(article.badges.length > 0 || article.tags.length > 0) && (
                    <ul className="mt-1 flex flex-wrap gap-1.5">
                        {article.badges.slice(0, 2).map((badge) => (
                            <li
                                key={`badge-${badge}`}
                                className="rounded-full bg-[#16A34A]/10 px-2.5 py-0.5 text-xs font-medium text-[#15803D]"
                            >
                                {badge}
                            </li>
                        ))}
                        {article.tags.slice(0, 2).map((tag) => (
                            <li
                                key={`tag-${tag}`}
                                className="rounded-full bg-[#F8FAFC] px-2.5 py-0.5 text-xs text-[#6B7280] ring-1 ring-inset ring-gray-200"
                            >
                                #{tag}
                            </li>
                        ))}
                    </ul>
                )}

                <span className="mt-auto inline-flex items-center gap-1 pt-3 text-sm font-medium text-[#0B63CE]">
                    Baca artikel
                    <ArrowUpRight
                        className="h-4 w-4 transition-transform group-hover:-translate-y-0.5 group-hover:translate-x-0.5 motion-reduce:transition-none"
                        aria-hidden="true"
                    />
                </span>
            </div>
        </article>
    );
}

export default ArticleMiniCard;