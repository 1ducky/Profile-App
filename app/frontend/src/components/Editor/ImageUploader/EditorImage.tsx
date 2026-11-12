import { success } from "astro:schema";
import { useEffect, useId, useRef, useState, type DragEvent } from "react";
import { ResolveUrlAsset } from "../../../utils/resolve-url";

/**
 * UploadImage — island Astro (client:load / client:visible).
 *
 * Kontrak backend:
 *  POST multipart/form-data { file, title, alt }
 *  2xx -> { url: string }
 *  4xx/5xx -> { error?: string }
 */

export type UploadImageValue = {
    url: string;
    title: string;
    alt: string;
    source: "upload" | "external";
};

export type UploadImageProps = {
    /** URL backend penerima file */
    endpoint?: string;
    /** Nama field file pada multipart */
    fieldName?: string;
    /** Batas ukuran file (MB) */
    maxSizeMB?: number;
    /** Prefix name untuk hidden input (dipakai di <form> biasa) */
    name?: string;
    onChange?: (value: UploadImageValue) => void;
    /** Dipanggil saat hasil sebelumnya dibatalkan (Ganti gambar / ganti sumber) */
    onReset?: () => void;
    /** Tanpa kartu pembungkus (untuk dipakai di dalam modal) */
    bare?: boolean;
};

type Mode = "file" | "url";
type Status = "idle" | "uploading" | "done" | "error";
type ExtState = "idle" | "loading" | "ok" | "bad";
type ErrorResponse = {
    status: "error";
    code: string;
    message: string;
    request_id: string;
};

type SuccessResponse<T> = {
    status: "success";
    data: T;
};

type UploadData = {
    filekey: string;      // wajib kalau sukses
    metadata?: unknown;
};

type UploadResponse = SuccessResponse<UploadData> | ErrorResponse;

const ACCEPT = ["image/jpeg", "image/png", "image/webp", "image/avif", "image/gif"] as const;

const cn = (...c: Array<string | false | null | undefined>) => c.filter(Boolean).join(" ");

const btnBase =
    "rounded-pill px-5 py-2.5 text-sm font-medium transition-colors focus-visible:outline-3 focus-visible:outline-offset-2 focus-visible:outline-primary/50 disabled:cursor-not-allowed disabled:opacity-60";
const btnPrimary = cn(btnBase, "border border-primary bg-primary font-semibold text-white hover:enabled:bg-primary/90");
const btnGhost = cn(btnBase, "border border-text-muted/30 bg-white text-text hover:enabled:bg-bg");
const inputCls =
    "w-full rounded-xl border border-text-muted/30 bg-white px-3 py-2.5 text-text placeholder:text-text-muted/70 focus-visible:outline-3 focus-visible:outline-offset-1 focus-visible:outline-primary/50 disabled:bg-bg disabled:text-text-muted aria-[invalid=true]:border-danger";

export default function UploadImage({
    endpoint = "/api/upload",
    fieldName = "image",
    maxSizeMB = 5,
    name = "image",
    onChange,
    onReset,
    bare = false,
}: UploadImageProps) {
    const uid = useId();
    const inputRef = useRef<HTMLInputElement>(null);
    const xhrRef = useRef<XMLHttpRequest | null>(null);

    const [mode, setMode] = useState<Mode>("file");
    const [file, setFile] = useState<File | null>(null);
    const [preview, setPreview] = useState("");
    const [status, setStatus] = useState<Status>("idle");
    const [progress, setProgress] = useState(0);
    const [message, setMessage] = useState("");
    const [dragging, setDragging] = useState(false);

    const [title, setTitle] = useState("");
    const [alt, setAlt] = useState("");
    const [altError, setAltError] = useState("");

    const [extUrl, setExtUrl] = useState("");
    const [extState, setExtState] = useState<ExtState>("idle");
    const [finalUrl, setFinalUrl] = useState("");

    // Lepas object URL lama & batalkan request saat unmount
    useEffect(() => {
        return () => {
            if (preview.startsWith("blob:")) URL.revokeObjectURL(preview);
        };
    }, [preview]);
    useEffect(() => () => xhrRef.current?.abort(), []);

    const uploading = status === "uploading";

    function reset() {
        onReset?.();
        xhrRef.current?.abort();
        setFile(null);
        setPreview("");
        setStatus("idle");
        setProgress(0);
        setMessage("");
        setFinalUrl("");
        if (inputRef.current) inputRef.current.value = "";
    }

    function pickFile(f?: File | null) {
        if (!f) return;
        if (!(ACCEPT as readonly string[]).includes(f.type)) {
            setStatus("error");
            setMessage("Format tidak didukung. Gunakan JPG, PNG, WebP, AVIF, atau GIF.");
            return;
        }
        if (f.size > maxSizeMB * 1024 * 1024) {
            setStatus("error");
            setMessage(`Ukuran file ${(f.size / 1048576).toFixed(1)} MB melebihi batas ${maxSizeMB} MB.`);
            return;
        }
        setFile(f);
        setPreview(URL.createObjectURL(f));
        setStatus("idle");
        setMessage("");
        setProgress(0);
        setFinalUrl("");
    }

    function validateAlt(): boolean {
        if (!alt.trim()) {
            setAltError("Teks alternatif wajib diisi agar gambar bisa dibaca pembaca layar.");
            return false;
        }
        setAltError("");
        return true;
    }

    function DemoUpload() {
        if (!file || !validateAlt()) return;

        const url = URL.createObjectURL(file)

        setFinalUrl(url);
        console.log(url)
        setStatus("done");
        setMessage("");
        onChange?.({ url, title: title.trim(), alt: alt.trim(), source: "upload" });
    }
    function upload() {
        if (!file || !validateAlt()) return;

        const body = new FormData();
        body.append(fieldName, file);

        const xhr = new XMLHttpRequest();
        xhrRef.current = xhr;
        xhr.open("POST", endpoint);
        xhr.responseType = "json";

        xhr.upload.onprogress = (e) => {
            if (e.lengthComputable) setProgress(Math.round((e.loaded / e.total) * 100));
        };
        xhr.onload = () => {
            const res = xhr.response as UploadResponse;
            console.log(res)
            if (xhr.status >= 200 && xhr.status < 300 && res?.status == "success") {
                setFinalUrl(res?.data?.filekey);
                setStatus("done");
                setMessage("");
                onChange?.({ url: res?.data?.filekey, title: title.trim(), alt: alt.trim(), source: "upload" });
            } else {
                setStatus("error");
                setMessage(
                    res?.status == "error" ?
                        (xhr.status === 413
                            ? "File terlalu besar untuk diterima server."
                            : `Unggah gagal (kode ${xhr.status}). Coba lagi sebentar lagi. RequestID: ${res.request_id} ERRCODE:${res.code}`) : "Kesalahan Internal."
                );
                console.log(res?.status == "error" && res)
            }
        };
        xhr.onerror = () => {
            setStatus("error");
            setMessage("Tidak bisa terhubung ke server. Periksa koneksi internet lalu coba lagi.");
        };

        setStatus("uploading");
        setProgress(0);
        setMessage("");
        xhr.send(body);
    }

    function cancelUpload() {
        xhrRef.current?.abort();
        setStatus("idle");
        setProgress(0);
    }

    function checkExternal(value: string) {
        if (status === "done") onReset?.();
        setExtUrl(value);
        setFinalUrl("");
        setStatus("idle");
        const v = value.trim();
        if (!v) {
            setExtState("idle");
            setMessage("");
            return;
        }
        try {
            const u = new URL(v);
            if (u.protocol !== "http:" && u.protocol !== "https:") throw new Error("protocol");
            setExtState("loading");
            setMessage("");
        } catch {
            setExtState("bad");
            setMessage("Masukkan alamat lengkap yang diawali http:// atau https://");
        }
    }

    function useExternal() {
        if (extState !== "ok" || !validateAlt()) return;
        const url = extUrl.trim();
        setFinalUrl(url);
        setStatus("done");
        onChange?.({ url, title: title.trim(), alt: alt.trim(), source: "external" });
    }

    function onDrop(e: DragEvent<HTMLDivElement>) {
        e.preventDefault();
        setDragging(false);
        pickFile(e.dataTransfer.files?.[0]);
    }

    const shownPreview =
        mode === "file" ? preview : extState !== "bad" && extUrl.trim() ? extUrl.trim() : "";
    const showMeta = (mode === "file" && !!file) || (mode === "url" && extState === "ok");

    return (
        <div className={cn("grid gap-4 font-sans text-text", !bare && "max-w-lg rounded-card border border-text-muted/20 bg-white p-5 shadow-card")}>
            {/* Pilih sumber */}
            <div role="tablist" aria-label="Sumber gambar" className="inline-flex w-fit rounded-pill border border-text-muted/20 bg-bg p-1">
                {(
                    [
                        ["file", "Unggah file"],
                        ["url", "URL eksternal"],
                    ] as const
                ).map(([key, label]) => (
                    <button
                        key={key}
                        role="tab"
                        type="button"
                        aria-selected={mode === key}
                        disabled={uploading}
                        onClick={() => {
                            if (status === "done") onReset?.();
                            setMode(key);
                            setMessage("");
                            setStatus("idle");
                            setFinalUrl("");
                        }}
                        className={cn(
                            "rounded-pill px-4 py-2 text-sm transition-colors disabled:cursor-not-allowed disabled:opacity-50",
                            mode === key ? "bg-white font-semibold text-text shadow-sm" : "text-text-muted hover:text-text",
                        )}
                    >
                        {label}
                    </button>
                ))}
            </div>

            {/* Dropzone / input URL */}
            {mode === "file" && !file && (
                <div
                    onDragOver={(e) => {
                        e.preventDefault();
                        setDragging(true);
                    }}
                    onDragLeave={() => setDragging(false)}
                    onDrop={onDrop}
                    className={cn(
                        "grid justify-items-center gap-2 rounded-card border-2 border-dashed px-4 py-8 text-center transition-colors",
                        dragging ? "border-primary bg-primary/10" : "border-text-muted/30 bg-bg",
                    )}
                >
                    <p className="font-semibold">Tarik gambar ke sini</p>
                    <p className="text-sm text-text-muted">JPG, PNG, WebP, AVIF, atau GIF. Maksimal {maxSizeMB} MB.</p>
                    <button type="button" className={cn(btnPrimary, "mt-1")} onClick={() => inputRef.current?.click()}>
                        Pilih dari perangkat
                    </button>
                </div>
            )}

            {mode === "url" && (
                <div className="grid gap-1.5">
                    <label htmlFor={`${uid}-url`} className="text-sm font-semibold">
                        Alamat gambar
                    </label>
                    <input
                        id={`${uid}-url`}
                        type="url"
                        inputMode="url"
                        placeholder="https://contoh.com/foto.jpg"
                        value={extUrl}
                        onChange={(e) => checkExternal(e.target.value)}
                        className={inputCls}
                    />
                </div>
            )}

            <input
                ref={inputRef}
                type="file"
                accept={ACCEPT.join(",")}
                hidden
                onChange={(e) => pickFile(e.target.files?.[0])}
            />

            {/* Preview */}
            {shownPreview && (
                <figure className="relative m-0 aspect-16/10 overflow-hidden rounded-card border border-text-muted/20 bg-bg">
                    <img
                        src={ResolveUrlAsset(shownPreview)}
                        alt={alt.trim() || "Pratinjau gambar"}
                        onLoad={() => {
                            if (mode === "url") {
                                setExtState("ok");
                                setMessage("");
                            }
                        }}
                        onError={() => {
                            if (mode === "url") {
                                setExtState("bad");
                                setMessage("Gambar tidak bisa dimuat dari alamat ini. Periksa kembali URL-nya.");
                            }
                        }}
                        className={cn("h-full w-full object-cover transition-[filter] duration-300 motion-reduce:transition-none", uploading && "brightness-90 grayscale")}
                    />

                    {uploading && (
                        <div role="status" aria-live="polite" className="absolute inset-0 flex flex-col items-center justify-center gap-2.5 bg-primary/35 font-semibold text-white">
                            <span aria-hidden="true" className="size-9 animate-spin rounded-full border-4 border-white/35 border-t-white motion-reduce:[animation-duration:2.4s]" />
                            <span>Mengunggah… {progress}%</span>
                        </div>
                    )}

                    {status === "done" && (
                        <span className="absolute left-3 top-3 rounded-pill bg-secondary px-3 py-1 text-xs font-semibold text-white">Tersimpan</span>
                    )}
                </figure>
            )}

            {uploading && (
                <div
                    role="progressbar"
                    aria-label="Progres unggah"
                    aria-valuemin={0}
                    aria-valuemax={100}
                    aria-valuenow={progress}
                    className="h-1.5 overflow-hidden rounded-pill bg-text-muted/20"
                >
                    <span className="block h-full rounded-pill bg-primary transition-[width] duration-200 motion-reduce:transition-none" style={{ width: `${progress}%` }} />
                </div>
            )}

            {/* Title & alt */}
            {showMeta && (
                <div className="grid gap-3 sm:grid-cols-2">
                    <div className="grid gap-1.5">
                        <label htmlFor={`${uid}-title`} className="text-sm font-semibold">
                            Judul
                        </label>
                        <input
                            id={`${uid}-title`}
                            type="text"
                            value={title}
                            disabled={uploading || status === "done"}
                            onChange={(e) => setTitle(e.target.value)}
                            placeholder="Mis. Pelatihan literasi digital"
                            className={inputCls}
                        />
                    </div>
                    <div className="grid gap-1.5">
                        <label htmlFor={`${uid}-alt`} className="text-sm font-semibold">
                            Teks alternatif (alt)
                        </label>
                        <input
                            id={`${uid}-alt`}
                            type="text"
                            value={alt}
                            disabled={uploading || status === "done"}
                            aria-invalid={!!altError}
                            aria-describedby={`${uid}-alt-help`}
                            onChange={(e) => {
                                setAlt(e.target.value);
                                if (altError) setAltError("");
                            }}
                            placeholder="Jelaskan isi gambar secara singkat"
                            className={inputCls}
                        />
                        <small id={`${uid}-alt-help`} className={cn("text-[0.8125rem]", altError ? "text-danger" : "text-text-muted")}>
                            {altError || "Contoh: Siswa belajar memakai tablet di ruang kelas."}
                        </small>
                    </div>
                </div>
            )}

            {/* Pesan gagal */}
            {message && (
                <div role="alert" className="rounded-xl border border-danger/30 bg-danger/10 px-3.5 py-3 text-sm text-danger">
                    {message}
                </div>
            )}

            {/* Hasil */}
            {status === "done" && finalUrl && (
                <div className="grid gap-1 rounded-xl bg-section-warm px-3.5 py-3">
                    <span className="text-sm text-text-muted">URL gambar</span>
                    <code className="break-all text-[0.8125rem]">{finalUrl}</code>
                </div>
            )}

            {/* Aksi */}
            <div className="flex flex-wrap gap-2.5 max-md:[&>button]:basis-full">
                {mode === "file" && file && status !== "done" && (
                    <button type="button" className={btnPrimary} disabled={uploading} onClick={upload}>
                        {uploading ? "Mengunggah…" : status === "error" ? "Coba unggah lagi" : "Unggah gambar"}
                    </button>
                )}
                {mode === "url" && extState === "ok" && status !== "done" && (
                    <button type="button" className={btnPrimary} onClick={useExternal}>
                        Pakai URL ini
                    </button>
                )}
                {uploading ? (
                    <button type="button" className={btnGhost} onClick={cancelUpload}>
                        Batalkan
                    </button>
                ) : (
                    (file || status === "done") && (
                        <button type="button" className={btnGhost} onClick={reset}>
                            {status === "done" ? "Ganti gambar" : "Hapus pilihan"}
                        </button>
                    )
                )}
            </div>

            {/* Untuk <form> biasa */}
            {status === "done" && (
                <>
                    <input type="hidden" name={`${name}[url]`} value={finalUrl} />
                    <input type="hidden" name={`${name}[title]`} value={title} />
                    <input type="hidden" name={`${name}[alt]`} value={alt} />
                </>
            )}
        </div>
    );
}