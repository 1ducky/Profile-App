

import { useMemo, useState } from "react";
import type { Article, ArticleImage } from "../../types/article/entity";
import { InsertImageButton } from "./EditorImage";
import EditorMarkdown from "./EditorMarkDown";
// import type { Article, ArticleImage } from "@/types/article";
// import MarkdownEditor from "@/components/MarkdownEditor";   // <- komponen editor markdown kamu
// import ImagePickerButton from "@/components/ImagePickerButton"; // <- komponen upload/pilih gambar kamu

/**
 * Field yang diisi dari form.
 * Dibuat backend: id, slug, contentType, publishedAt, updatedAt, author, stats, related.
 */
export type ArticleFormValues = Pick<
    Article,
    "title" | "excerpt" | "category" | "subCategory" | "tags" | "badges"
> & {
    mode: "write" | "external";
    /** Mode external: tujuan redirect (mis. https://youtube.com/...). Mode write: dikosongkan. */
    url?: string;
    /** Mode write: isi markdown. */
    Content?: string;
    image?: ArticleImage;
};

type Props = {
    initialValues?: Partial<ArticleFormValues>;
    onSaveDraft?: (values: ArticleFormValues) => Promise<void> | void;
    onPublish?: (values: ArticleFormValues) => Promise<void> | void;
};

const ASPECT_OPTIONS = ["16/9", "4/3", "1/1"];

const inputCls =
    "w-full rounded-[10px] border border-[#D1D5DB] bg-white px-3.5 py-2.5 text-[15px] text-[#1F2937] placeholder:text-[#9CA3AF] outline-none transition focus:border-[#0B63CE] focus:ring-2 focus:ring-[#0B63CE]/20";

function Field({
    label,
    hint,
    error,
    htmlFor,
    children,
}: {
    label: string;
    hint?: string;
    error?: string;
    htmlFor?: string;
    children: React.ReactNode;
}) {
    return (
        <div className="flex flex-col gap-1.5">
            <label htmlFor={htmlFor} className="text-sm font-semibold text-[#1F2937]">
                {label}
            </label>
            {children}
            {error ? (
                <p className="text-sm text-[#B91C1C]">{error}</p>
            ) : hint ? (
                <p className="text-sm text-[#6B7280]">{hint}</p>
            ) : null}
        </div>
    );
}

/** Input tag: Enter / koma menambah, klik × menghapus. */
function TagInput({
    id,
    value,
    onChange,
    placeholder,
}: {
    id: string;
    value: string[];
    onChange: (v: string[]) => void;
    placeholder?: string;
}) {
    const [draft, setDraft] = useState("");

    const commit = () => {
        const t = draft.trim();
        if (t && !value.includes(t)) onChange([...value, t]);
        setDraft("");
    };

    return (
        <div className="flex flex-wrap items-center gap-2 rounded-[10px] border border-[#D1D5DB] bg-white px-2.5 py-2 focus-within:border-[#0B63CE] focus-within:ring-2 focus-within:ring-[#0B63CE]/20">
            {value.map((t) => (
                <span
                    key={t}
                    className="inline-flex items-center gap-1 rounded-full bg-[#F5EFE3] px-2.5 py-1 text-sm text-[#1F2937]"
                >
                    {t}
                    <button
                        type="button"
                        aria-label={`Hapus ${t}`}
                        onClick={() => onChange(value.filter((x) => x !== t))}
                        className="text-[#6B7280] hover:text-[#1F2937]"
                    >
                        ×
                    </button>
                </span>
            ))}
            <input
                id={id}
                value={draft}
                placeholder={value.length ? "" : placeholder}
                onChange={(e) => setDraft(e.target.value)}
                onKeyDown={(e) => {
                    if (e.key === "Enter" || e.key === ",") {
                        e.preventDefault();
                        commit();
                    } else if (e.key === "Backspace" && !draft && value.length) {
                        onChange(value.slice(0, -1));
                    }
                }}
                onBlur={commit}
                className="min-w-[120px] flex-1 bg-transparent px-1 py-1 text-[15px] text-[#1F2937] outline-none placeholder:text-[#9CA3AF]"
            />
        </div>
    );
}

function isValidUrl(v: string) {
    try {
        const u = new URL(v);
        return u.protocol === "http:" || u.protocol === "https:";
    } catch {
        return false;
    }
}

export default function EditorForm({
    initialValues,
    onSaveDraft,
    onPublish,
}: Props) {
    const [mode, setMode] = useState<ArticleFormValues["mode"]>(
        initialValues?.mode ?? "write"
    );
    const [title, setTitle] = useState(initialValues?.title ?? "");
    const [excerpt, setExcerpt] = useState(initialValues?.excerpt ?? "");
    const [category, setCategory] = useState(initialValues?.category ?? "");
    const [subCategory, setSubCategory] = useState(initialValues?.subCategory ?? "");
    const [tags, setTags] = useState<string[]>(initialValues?.tags ?? []);
    const [badges, setBadges] = useState<string[]>(initialValues?.badges ?? []);
    const [externalUrl, setExternalUrl] = useState(initialValues?.url ?? "");
    const [markdown, setMarkdown] = useState(initialValues?.Content ?? "");

    // Gambar: src diisi oleh komponen upload kamu (lewat setImageSrc)
    const [imageSrc, setImageSrc] = useState(initialValues?.image?.src ?? "");
    const [imageAlt, setImageAlt] = useState(initialValues?.image?.alt ?? "");
    const [imageCaption, setImageCaption] = useState(initialValues?.image?.caption ?? "");
    const [aspectRatio, setAspectRatio] = useState(initialValues?.image?.aspectRatio ?? "16/9");

    const [submitting, setSubmitting] = useState<null | "draft" | "publish">(null);
    const [touched, setTouched] = useState(false);

    const errors = useMemo(() => {
        const e: Record<string, string> = {};
        if (!title.trim()) e.title = "Judul wajib diisi.";
        if (!category.trim()) e.category = "Kategori wajib diisi.";
        if (imageSrc && !imageAlt.trim()) e.imageAlt = "Teks alternatif wajib diisi jika ada gambar.";
        if (mode === "external") {
            if (!externalUrl.trim()) e.url = "Link tujuan wajib diisi.";
            else if (!isValidUrl(externalUrl.trim())) e.url = "Gunakan link lengkap, mis. https://youtube.com/watch?v=...";
        } else if (!markdown.trim()) {
            e.content = "Isi artikel belum ditulis.";
        }
        return e;
    }, [title, category, imageSrc, imageAlt, mode, externalUrl, markdown]);

    const buildValues = (): ArticleFormValues => ({
        mode,
        title: title.trim(),
        excerpt: excerpt.trim() || undefined,
        category: category.trim(),
        subCategory: subCategory.trim() || undefined,
        tags: tags.length ? tags : undefined,
        badges: badges.length ? badges : undefined,
        url: mode === "external" ? externalUrl.trim() : undefined,
        Content: mode === "write" ? markdown : undefined,
        image: imageSrc
            ? {
                src: imageSrc,
                alt: imageAlt.trim(),
                caption: imageCaption.trim() || undefined,
                aspectRatio,
            }
            : undefined,
    });

    const submit = async (kind: "draft" | "publish") => {
        setTouched(true);
        // Draft hanya butuh judul; publish butuh semua valid.
        if (kind === "draft" ? errors.title : Object.keys(errors).length) return;
        setSubmitting(kind);
        try {
            const values = buildValues();
            console.log(values)
            if (kind === "draft") await onSaveDraft?.(values);
            else await onPublish?.(values);
        } finally {
            setSubmitting(null);
        }
    };

    const err = (k: string) => (touched ? errors[k] : undefined);

    return (
        <div className="min-h-screen bg-[#F8FAFC] px-4 py-8 font-sans text-[#1F2937] sm:px-6 sm:py-12">
            <form
                noValidate
                onSubmit={(e) => e.preventDefault()}
                className="mx-auto flex w-full max-w-[860px] flex-col gap-8"
            >
                <header className="flex flex-col gap-1.5">
                    <h1 className="text-[28px] font-bold leading-tight tracking-tight sm:text-[32px]">
                        Artikel baru
                    </h1>
                    <p className="text-[15px] text-[#6B7280]">
                        Tulis artikel sendiri, atau arahkan pembaca ke tautan di situs lain.
                    </p>
                </header>

                {/* Pilihan mode */}
                <div
                    role="radiogroup"
                    aria-label="Jenis artikel"
                    className="grid grid-cols-1 gap-3 sm:grid-cols-2"
                >
                    {(
                        [
                            { v: "write", t: "Tulis artikel", d: "Isi ditulis di sini dan tampil di situs." },
                            { v: "external", t: "Link eksternal", d: "Pembaca diarahkan ke situs lain, mis. YouTube." },
                        ] as const
                    ).map((o) => {
                        const active = mode === o.v;
                        return (
                            <button
                                key={o.v}
                                type="button"
                                role="radio"
                                aria-checked={active}
                                onClick={() => setMode(o.v)}
                                className={`rounded-2xl border p-4 text-left transition focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-[#0B63CE]/40 ${active
                                    ? "border-[#0B63CE] bg-white shadow-[0_4px_16px_rgba(31,41,55,0.06)]"
                                    : "border-[#E5E7EB] bg-transparent hover:bg-white"
                                    }`}
                            >
                                <span className="flex items-center gap-2.5">
                                    <span
                                        className={`grid h-4 w-4 place-items-center rounded-full border ${active ? "border-[#0B63CE]" : "border-[#9CA3AF]"
                                            }`}
                                    >
                                        {active && <span className="h-2 w-2 rounded-full bg-[#0B63CE]" />}
                                    </span>
                                    <span className="text-[15px] font-semibold">{o.t}</span>
                                </span>
                                <span className="mt-1.5 block pl-[26px] text-sm text-[#6B7280]">{o.d}</span>
                            </button>
                        );
                    })}
                </div>

                {/* Konten utama */}
                <section className="flex flex-col gap-5 rounded-2xl border border-[#E5E7EB] bg-white p-5 shadow-[0_4px_16px_rgba(31,41,55,0.06)] sm:p-7">
                    <Field label="Judul" htmlFor="title" error={err("title")}>
                        <input
                            id="title"
                            value={title}
                            onChange={(e) => setTitle(e.target.value)}
                            placeholder="Judul yang jelas dan singkat"
                            className={inputCls}
                        />
                    </Field>

                    {mode === "external" && (
                        <Field
                            label="Link tujuan"
                            htmlFor="url"
                            error={err("url")}
                            hint="Pembaca akan diarahkan ke alamat ini saat mengklik artikel."
                        >
                            <input
                                id="url"
                                type="url"
                                inputMode="url"
                                value={externalUrl}
                                onChange={(e) => setExternalUrl(e.target.value)}
                                placeholder="https://youtube.com/watch?v=..."
                                className={inputCls}
                            />
                        </Field>
                    )}

                    <Field
                        label="Ringkasan"
                        htmlFor="excerpt"
                        hint="1–3 kalimat. Tampil di kartu artikel."
                    >
                        <textarea
                            id="excerpt"
                            rows={3}
                            value={excerpt}
                            onChange={(e) => setExcerpt(e.target.value)}
                            className={`${inputCls} resize-y leading-relaxed`}
                        />
                    </Field>

                    {/* Gambar */}
                    <fieldset className="flex flex-col gap-4 border-t border-[#E5E7EB] pt-5">
                        <legend className="sr-only">Gambar sampul</legend>
                        <div className="flex items-center justify-between gap-3">
                            <span className="text-sm font-semibold">Gambar sampul</span>

                            <InsertImageButton onSuccess={(image) => {
                                setImageSrc(image.src);
                                setImageAlt(image.alt);
                                setImageCaption(image.caption || "")
                            }} />
                        </div>

                        {imageSrc ? (
                            <div className="flex flex-col gap-4 sm:flex-row">
                                <div
                                    className="w-full overflow-hidden rounded-xl bg-[#F5EFE3] sm:w-60 sm:shrink-0"
                                    style={{ aspectRatio }}
                                >
                                    {/* eslint-disable-next-line @next/next/no-img-element */}
                                    <img
                                        src={imageSrc}
                                        alt={imageAlt || "Pratinjau gambar sampul"}
                                        className="h-full w-full object-cover"
                                    />
                                </div>
                                <div className="flex flex-1 flex-col gap-4">
                                    <Field label="Teks alternatif" htmlFor="alt" error={err("imageAlt")} hint="Jelaskan isi gambar untuk pembaca layar dan SEO.">
                                        <input
                                            id="alt"
                                            value={imageAlt}
                                            onChange={(e) => setImageAlt(e.target.value)}
                                            className={inputCls}
                                        />
                                    </Field>
                                    <Field label="Keterangan foto" htmlFor="caption">
                                        <input
                                            id="caption"
                                            value={imageCaption}
                                            onChange={(e) => setImageCaption(e.target.value)}
                                            placeholder="Foto: nama fotografer / sumber"
                                            className={inputCls}
                                        />
                                    </Field>
                                    <div className="flex items-end gap-3">
                                        <Field label="Rasio gambar" htmlFor="ratio">
                                            <select
                                                id="ratio"
                                                value={aspectRatio}
                                                onChange={(e) => setAspectRatio(e.target.value)}
                                                className={inputCls}
                                            >
                                                {ASPECT_OPTIONS.map((r) => (
                                                    <option key={r} value={r}>
                                                        {r.replace("/", ":")}
                                                    </option>
                                                ))}
                                            </select>
                                        </Field>
                                        <button
                                            type="button"
                                            onClick={() => {
                                                setImageSrc("");
                                                setImageAlt("");
                                                setImageCaption("");
                                            }}
                                            className="rounded-full px-4 py-2.5 text-sm font-medium text-[#B91C1C] hover:bg-[#FEF2F2]"
                                        >
                                            Hapus gambar
                                        </button>
                                    </div>
                                </div>
                            </div>
                        ) : (
                            <p className="text-sm text-[#6B7280]">Belum ada gambar dipilih.</p>
                        )}
                    </fieldset>
                </section>

                {/* Metadata */}
                <section className="flex flex-col gap-5 rounded-2xl border border-[#E5E7EB] bg-white p-5 shadow-[0_4px_16px_rgba(31,41,55,0.06)] sm:p-7">
                    <h2 className="text-lg font-semibold">Kategori dan label</h2>
                    <div className="grid grid-cols-1 gap-5 sm:grid-cols-2">
                        <Field label="Kategori" htmlFor="category" error={err("category")}>
                            <input
                                id="category"
                                value={category}
                                onChange={(e) => setCategory(e.target.value)}
                                placeholder="mis. Teknologi"
                                className={inputCls}
                            />
                        </Field>
                        <Field label="Subkategori" htmlFor="sub">
                            <input
                                id="sub"
                                value={subCategory}
                                onChange={(e) => setSubCategory(e.target.value)}
                                className={inputCls}
                            />
                        </Field>
                    </div>
                    <Field label="Tag" htmlFor="tags" hint="Tekan Enter atau koma untuk menambah.">
                        <TagInput id="tags" value={tags} onChange={setTags} placeholder="#Pendidikan" />
                    </Field>
                    <Field label="Badge" htmlFor="badges" hint="Penanda singkat di kartu, mis. Opini atau Video.">
                        <TagInput id="badges" value={badges} onChange={setBadges} placeholder="Opini" />
                    </Field>
                </section>

                {/* Isi artikel (hanya mode tulis) */}
                {/* {mode === "write" && (
                    <section className="flex flex-col gap-3 rounded-2xl border border-[#E5E7EB] bg-white p-5 shadow-[0_4px_16px_rgba(31,41,55,0.06)] sm:p-7">
                        <h2 className="text-lg font-semibold">Isi artikel</h2>
                        <EditorMarkdown
                            markdown={initialValues?.Content ?? ""}
                            onChange={(md) => setMarkdown(md)}
                        />

                        {err("content") && <p className="text-sm text-[#B91C1C]">{err("content")}</p>}
                    </section>
                )} */}

                {/* Aksi */}
                <div className="flex flex-col-reverse gap-3 sm:flex-row sm:justify-end">
                    <button
                        type="button"
                        disabled={!!submitting}
                        onClick={() => submit("draft")}
                        className="rounded-full border border-[#D1D5DB] bg-white px-6 py-3 text-[15px] font-semibold text-[#1F2937] transition hover:bg-[#F3F4F6] focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-[#0B63CE]/40 disabled:opacity-60"
                    >
                        {submitting === "draft" ? "Menyimpan…" : "Simpan draf"}
                    </button>
                    <button
                        type="button"
                        disabled={!!submitting}
                        onClick={() => submit("publish")}
                        className="rounded-full bg-[#0B63CE] px-6 py-3 text-[15px] font-semibold text-white transition hover:bg-[#0A55B0] focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-[#0B63CE]/40 focus-visible:ring-offset-2 disabled:opacity-60"
                    >
                        {submitting === "publish" ? "Menerbitkan…" : "Terbitkan"}
                    </button>
                </div>
            </form>
        </div>
    );
}