import { useRef, useState } from "react";
import type { ArticleField } from "../../contract/article";
import { InsertImageButton } from "./EditorImage";
import { ResolveUrlAsset } from "../../utils/resolve-url";
import { TagInput, type TagInputHandle } from "./EditorTagInput";
import type { MDXEditorMethods } from "@mdxeditor/editor";
import EditorMarkdown from "./EditorMarkDown";

const Sanitize = (req: ArticleField, mode: "write" | "external"): ArticleField => ({
    title: req.title.trim(),
    excerpt: req.excerpt.trim(),
    category: req.category.trim(),
    tags: req.tags.length ? req.tags : [],
    badges: req.badges.length ? req.badges : [],
    url: mode === "external" ? req.url.trim() : "",
    content: mode === "write" ? req.content : "",
    cover: req.cover.trim()
});

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
                <p className="text-sm text-text-muted">{hint}</p>
            ) : null}
        </div>
    );
}



export function EditorArticle({ onPublish, defaultValue }: { onPublish: (req: ArticleField) => void, defaultValue?: ArticleField }) {
    const tagsRef = useRef<TagInputHandle>(null);
    const badgesRef = useRef<TagInputHandle>(null);
    const mdEditorRef = useRef<MDXEditorMethods>(null)
    const [mode, setMode] = useState<"write" | "external">("write")
    const [cover, setCover] = useState<string>(defaultValue?.cover || "")

    const [isSubmmiting, setIsSubmmiting] = useState<Boolean>(false)

    function onSubmit(e: React.FormEvent<HTMLFormElement>) {
        e.preventDefault()
        setIsSubmmiting(true)
        const formData = new FormData(e.currentTarget)
        var req: ArticleField | null = null
        req = Sanitize({
            badges: badgesRef.current?.getValue() || [],
            tags: tagsRef.current?.getValue() || [],
            content: mdEditorRef.current?.getMarkdown() || "",
            category: String(formData.get("category") ?? ""),
            excerpt: String(formData.get("excerpt") ?? ""),
            title: String(formData.get("title") ?? ""),
            url: String(formData.get("url") ?? ""),
            cover: cover,

        }, mode)
        try {
            console.log(req)
            onPublish(req)

        } catch (e) {

        } finally {
            setIsSubmmiting(false)
        }

    }

    return (
        <div className="min-h-screen bg-bg px-4 py-8 font-sans text-[#1F2937] sm:px-6 sm:py-12">
            <form onSubmit={(e) => onSubmit(e)} className="mx-auto flex w-full max-w-215 flex-col gap-8">
                <header className="flex flex-col gap-1.5">
                    <h1 className="text-[28px] font-bold leading-tight tracking-tight sm:text-[32px]">
                        {defaultValue ? "Perbarui Article" : "Article Baru"}
                    </h1>
                    <p className="text-[15px] text-text-muted">
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
                                className={`rounded-2xl border p-4 text-left transition focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary/40 ${active
                                    ? "border-primary bg-white shadow-card"
                                    : "border-bg bg-transparent hover:bg-white"
                                    }`}
                            >
                                <span className="flex items-center gap-2.5">
                                    <span
                                        className={`grid h-4 w-4 place-items-center rounded-full border ${active ? "border-primary" : "border-[#9CA3AF]"
                                            }`}
                                    >
                                        {active && <span className="h-2 w-2 rounded-full bg-primary" />}
                                    </span>
                                    <span className="text-[15px] font-semibold">{o.t}</span>
                                </span>
                                <span className="mt-1.5 block pl-6.5 text-sm text-text-muted">{o.d}</span>
                            </button>
                        );
                    })}
                </div>

                {/* Konten utama */}
                <section className="flex flex-col gap-5 rounded-2xl border border-border bg-white p-5 shadow-card sm:p-7">
                    <Field label="Judul" htmlFor="title" error={undefined}>
                        <input
                            id="title"
                            name="title"
                            defaultValue={defaultValue?.title || ""}
                            placeholder="Judul yang jelas dan singkat"
                            className="w-full rounded-[10px] border border-[#D1D5DB] bg-white px-3.5 py-2.5 text-[15px] text-[#1F2937] placeholder:text-[#9CA3AF] outline-none transition focus:border-primary focus:ring-2 focus:ring-primary/20"
                        />
                    </Field>

                    {mode === "external" && (
                        <Field
                            label="Link tujuan"
                            htmlFor="url"
                            error={undefined}
                            hint="Pembaca akan diarahkan ke alamat ini saat mengklik artikel."
                        >
                            <input
                                id="url"
                                name="url"
                                type="url"
                                inputMode="url"
                                defaultValue={defaultValue?.url || ""}
                                placeholder="https://youtube.com/watch?v=..."
                                className="w-full rounded-[10px] border border-[#D1D5DB] bg-white px-3.5 py-2.5 text-[15px] text-[#1F2937] placeholder:text-[#9CA3AF] outline-none transition focus:border-primary focus:ring-2 focus:ring-primary/20"
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
                            name="excerpt"
                            defaultValue={defaultValue?.excerpt || ""}
                            rows={3}
                            className={"w-full rounded-[10px] border border-[#D1D5DB] bg-white px-3.5 py-2.5 text-[15px] text-[#1F2937] placeholder:text-[#9CA3AF] outline-none transition focus:border-primary focus:ring-2 focus:ring-primary/20" + `resize-y leading-relaxed`}
                        />
                    </Field>

                    {/* Gambar */}
                    <fieldset className="flex flex-col gap-4 border-t border-border pt-5">
                        <legend className="sr-only">Gambar sampul</legend>
                        <div className="flex items-center justify-between gap-3">
                            <span className="text-sm font-semibold">Gambar sampul</span>

                            <InsertImageButton onSuccess={(image) => {
                                setCover(image.src);
                            }} />
                        </div>

                        {cover ? (
                            <div className="flex flex-col gap-4 sm:flex-row">
                                <div
                                    className="w-full overflow-hidden rounded-xl bg-section-warm sm:w-60 sm:shrink-0 aspect-video"
                                >
                                    {/* eslint-disable-next-line @next/next/no-img-element */}
                                    <img
                                        src={ResolveUrlAsset(cover)}
                                        alt={"Pratinjau gambar sampul"}
                                        className="h-full w-full object-cover"
                                    />
                                </div>
                                <div className="flex flex-1 flex-col gap-4">
                                    <button
                                        type="button"
                                        onClick={() => {
                                            setCover("");
                                        }}
                                        className="rounded-full px-4 py-2.5 text-sm font-medium text-[#B91C1C] hover:bg-[#FEF2F2]"
                                    >
                                        Hapus gambar
                                    </button>
                                </div>
                            </div>

                        ) : (
                            <p className="text-sm text-text-muted">Belum ada gambar dipilih.</p>
                        )}
                    </fieldset>
                </section>
                <section className="flex flex-col gap-5 rounded-2xl border border-border bg-white p-5 shadow-card sm:p-7">
                    <h2 className="text-lg font-semibold">Kategori dan label</h2>
                    <div className="grid grid-cols-1 gap-5 sm:grid-cols-2">
                        <Field label="Kategori" htmlFor="category" error={undefined}>
                            <input
                                id="category"
                                name="category"
                                defaultValue={defaultValue?.category || ""}
                                placeholder="mis. Teknologi"
                                className={"w-full rounded-[10px] border border-[#D1D5DB] bg-white px-3.5 py-2.5 text-[15px] text-[#1F2937] placeholder:text-[#9CA3AF] outline-none transition focus:border-primary focus:ring-2 focus:ring-primary/20"}
                            />
                        </Field>
                    </div>
                    <Field label="Tag" htmlFor="tags" hint="Tekan Enter atau koma untuk menambah.">
                        <TagInput id="tags" ref={tagsRef} placeholder="#Pendidikan" defaultValue={defaultValue?.tags || []} />
                    </Field>
                    <Field label="Badge" htmlFor="badges" hint="Penanda singkat di kartu, mis. Opini atau Video.">
                        <TagInput id="badges" ref={badgesRef} placeholder="Opini" defaultValue={defaultValue?.badges || []} />
                    </Field>
                </section>
                {mode == "write" && (
                    <section className="flex flex-col gap-3 rounded-2xl border border-border bg-white p-5 shadow-card sm:p-7">
                        <h2 className="text-lg font-semibold">Isi artikel</h2>
                        <EditorMarkdown
                            markdown={defaultValue?.content || ""}
                            ref={mdEditorRef}
                        />


                    </section>
                )}
                <div className="flex flex-col-reverse gap-3 sm:flex-row sm:justify-end">

                    <button
                        className="rounded-full bg-[#0B63CE] px-6 py-3 text-[15px] font-semibold text-white transition hover:bg-[#0A55B0] focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-[#0B63CE]/40 focus-visible:ring-offset-2 disabled:opacity-60"
                    >
                        {isSubmmiting ? "Menerbitkan…" : "Terbitkan"}
                    </button>
                </div>
            </form>
        </div>
    )
}