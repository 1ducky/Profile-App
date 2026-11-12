

import { useState } from "react";
import {
    CalendarDays,
    Clock,
    ImageOff,
    Link2,
    Loader2,
    Pencil,
    Tag,
    Trash2,
    User,
} from "lucide-react";
import { useAlert } from "../../../hooks/useAllert";
import type { ArticleResponse } from "../../../contract/article";
import { timeAgo } from "../../../utils/time";
import { ArticleThumbnail } from "../../Content/Thubnail";
import { ResolveUrlAsset } from "../../../utils/resolve-url";


/* ---------------------------------- Props --------------------------------- */

export interface ArticleListItemProps {
    article: ArticleResponse;
    /** Path dasar untuk edit. Tombol Update menuju `${root}/${article.id}` */
    root: string;
    /** Dipanggil setelah user menekan "Ya" pada konfirmasi hapus */
    onDelete: (id: string) => void | Promise<void>;
}

/* --------------------------------- Helpers -------------------------------- */

const joinPath = (root: string, id: string) =>
    `${root.replace(/\/+$/, "")}/${id}`;

/* -------------------------------- Component ------------------------------- */

export function ArticleListItem({ article, root, onDelete }: ArticleListItemProps) {
    const confirm = useAlert();
    const [deleting, setDeleting] = useState(false);


    const published = timeAgo(article.published_at || "");
    const updated = timeAgo(article.updated_at || "");
    const isDraft = !article.published_at;

    const handleDelete = async () => {
        if (deleting) return;
        const ok = await confirm({ title: `Hapus artikel "${article.title}"?` });
        if (!ok) return;

        try {
            setDeleting(true);
            await onDelete(article.id);
        } finally {
            setDeleting(false);
        }
    };

    return (
        <article
            aria-busy={deleting}
            className={`group flex flex-col gap-4 rounded-[16px] border border-slate-200 bg-white p-4 shadow-[0_4px_16px_rgba(31,41,55,0.06)] transition motion-reduce:transition-none hover:border-slate-300 sm:flex-row sm:items-stretch sm:gap-5 ${deleting ? "pointer-events-none opacity-60" : ""
                }`}
        >
            {/* Cover */}
            <div className="relative aspect-video w-full shrink-0 overflow-hidden rounded-xl bg-[#F5EFE3] sm:aspect-auto sm:h-auto sm:w-48 lg:w-56">
                {article.cover ? (
                    // eslint-disable-next-line @next/next/no-img-element
                    <ArticleThumbnail image={{ alt: article.title, src: ResolveUrlAsset(article.cover) }} className="absolute inset-0 h-full w-full object-cover" />
                ) : (
                    <div className="absolute inset-0 flex items-center justify-center text-[#6B7280]">
                        <ImageOff className="h-7 w-7" aria-hidden="true" />
                    </div>
                )}
            </div>

            {/* Konten */}
            <div className="flex min-w-0 flex-1 flex-col gap-3">
                <div className="flex flex-wrap items-center gap-2">
                    <span className="text-xs font-semibold uppercase tracking-[0.12em] text-[#0B63CE]">
                        {article.category}
                    </span>

                    <span
                        className={`rounded-full px-2.5 py-0.5 text-xs font-medium ${isDraft
                            ? "bg-[#E6CFA8]/50 text-[#1F2937]"
                            : "bg-[#16A34A]/10 text-[#16A34A]"
                            }`}
                    >
                        {isDraft ? "Draft" : "Terbit"}
                    </span>

                    {article.badges.map((badge) => (
                        <span
                            key={badge}
                            className="rounded-full bg-[#1AA6B7]/10 px-2.5 py-0.5 text-xs font-medium text-[#1AA6B7]"
                        >
                            {badge}
                        </span>
                    ))}
                </div>

                <div className="min-w-0">
                    <h3 className="line-clamp-2 text-lg font-semibold leading-snug tracking-tight text-[#1F2937]">
                        {article.title}
                    </h3>
                    <p className="mt-1 line-clamp-2 text-sm leading-relaxed text-[#6B7280]">
                        {article.excerpt}
                    </p>
                </div>

                {/* Metadata */}
                <dl className="flex flex-wrap items-center gap-x-4 gap-y-1.5 text-xs text-[#6B7280]">
                    <div className="flex items-center gap-1.5">
                        <dt className="sr-only">Penulis</dt>
                        <User className="h-3.5 w-3.5" aria-hidden="true" />
                        <dd>{article.author?.username ?? "—"}</dd>
                    </div>

                    <div className="flex items-center gap-1.5">
                        <dt className="sr-only">Tanggal terbit</dt>
                        <CalendarDays className="h-3.5 w-3.5" aria-hidden="true" />
                        <dd>{published ?? "Belum terbit"}</dd>
                    </div>

                    {updated && (
                        <div className="flex items-center gap-1.5">
                            <dt className="sr-only">Terakhir diperbarui</dt>
                            <Clock className="h-3.5 w-3.5" aria-hidden="true" />
                            <dd>Diperbarui {updated}</dd>
                        </div>
                    )}

                    <div className="flex min-w-0 items-center gap-1.5">
                        <dt className="sr-only">URL</dt>
                        <Link2 className="h-3.5 w-3.5 shrink-0" aria-hidden="true" />
                        <dd className="truncate">{article.url || `/${article.slug}`}</dd>
                    </div>
                </dl>

                {article.tags.length > 0 && (
                    <ul className="flex flex-wrap items-center gap-1.5" aria-label="Tag">
                        <Tag className="h-3.5 w-3.5 text-[#6B7280]" aria-hidden="true" />
                        {article.tags.map((tag) => (
                            <li
                                key={tag}
                                className="rounded-md bg-slate-100 px-2 py-0.5 text-xs text-[#6B7280]"
                            >
                                #{tag}
                            </li>
                        ))}
                    </ul>
                )}
            </div>

            {/* Aksi */}
            <div className="flex shrink-0 gap-2 sm:flex-col sm:justify-start">
                <a
                    href={joinPath(root, article.id)}
                    className="inline-flex flex-1 items-center justify-center gap-2 rounded-full bg-[#0B63CE] px-4 py-2 text-sm font-medium text-white transition hover:bg-[#0953ad] focus:outline-none focus-visible:ring-2 focus-visible:ring-[#0B63CE] focus-visible:ring-offset-2 sm:flex-none"
                >
                    <Pencil className="h-4 w-4" aria-hidden="true" />
                    Update
                </a>

                <button
                    type="button"
                    onClick={handleDelete}
                    disabled={deleting}
                    className="inline-flex flex-1 items-center justify-center gap-2 rounded-full border border-slate-300 bg-white px-4 py-2 text-sm font-medium text-[#1F2937] transition hover:border-red-300 hover:bg-red-50 hover:text-red-600 focus:outline-none focus-visible:ring-2 focus-visible:ring-red-500 focus-visible:ring-offset-2 disabled:cursor-not-allowed sm:flex-none"
                >
                    {deleting ? (
                        <Loader2 className="h-4 w-4 animate-spin motion-reduce:animate-none" aria-hidden="true" />
                    ) : (
                        <Trash2 className="h-4 w-4" aria-hidden="true" />
                    )}
                    {deleting ? "Menghapus…" : "Hapus"}
                </button>
            </div>
        </article>
    );
}