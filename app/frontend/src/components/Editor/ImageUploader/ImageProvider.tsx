import {
    createContext,
    useCallback,
    useContext,
    useEffect,
    useRef,
    useState,
    type ReactNode,
} from "react";
import UploadImage, { type UploadImageProps, type UploadImageValue } from "./EditorImage";

/**
 * Cara pakai singkat:
 *
 *   const { open } = useImageUploader();
 *   const img = await open({ title: "Foto sampul artikel" });
 *   if (img) setCover(img.url);   // null = pengguna menutup tanpa memilih
 *
 * Di Astro, pasang SEKALI di layout:
 *   <ImageUploaderProvider client:load />
 * Island lain (React terpisah) tetap bisa memanggil lewat hook yang sama
 * atau lewat `imageUploader.open()` — keduanya memakai provider yang sama.
 */

export type OpenOptions = Partial<Pick<UploadImageProps, "endpoint" | "fieldName" | "maxSizeMB">> & {
    /** Judul dialog */
    title?: string;
};

export type OpenImageUploader = (options?: OpenOptions) => Promise<UploadImageValue | null>;

/* ---------- Singleton: menjembatani island Astro yang terpisah ---------- */

// Disimpan di globalThis agar tetap satu instance walau modul ini ter-bundle
// ganda (HMR / chunk berbeda antar island).
type Registry = { open: OpenImageUploader | null; waiters: Array<() => void> };
const g = globalThis as typeof globalThis & { __imageUploader?: Registry };
const registry: Registry = (g.__imageUploader ??= { open: null, waiters: [] });

const WAIT_MS = 3000;

export const imageUploader: { open: OpenImageUploader } = {
    open: (options) => {
        if (registry.open) return registry.open(options);

        // Provider mungkin belum selesai hydrate: tunggu sebentar sebelum menyerah.
        return new Promise((resolve, reject) => {
            const run = () => {
                clearTimeout(timer);
                registry.open!(options).then(resolve, reject);
            };
            const timer = setTimeout(() => {
                registry.waiters = registry.waiters.filter((w) => w !== run);
                reject(
                    new Error(
                        "ImageUploaderProvider belum terpasang. Pastikan <ImageUploaderProvider client:load /> ada di layout halaman ini.",
                    ),
                );
            }, WAIT_MS);
            registry.waiters.push(run);
        });
    },
};

/* ---------- Context (untuk komponen di dalam pohon yang sama) ---------- */

const ImageUploaderContext = createContext<OpenImageUploader | null>(null);

export function useImageUploader(): { open: OpenImageUploader } {
    const ctx = useContext(ImageUploaderContext);
    // Di luar pohon provider (island lain) otomatis memakai singleton.
    return { open: ctx ?? imageUploader.open };
}

/* ---------- Provider ---------- */

type ProviderProps = Pick<UploadImageProps, "endpoint" | "fieldName" | "maxSizeMB"> & {
    children?: ReactNode;
};

const btn =
    "rounded-pill px-5 py-2.5 text-sm font-medium transition-colors focus-visible:outline-3 focus-visible:outline-offset-2 focus-visible:outline-primary/50 disabled:cursor-not-allowed disabled:opacity-60";

export function ImageUploaderProvider({ children, endpoint, fieldName, maxSizeMB }: ProviderProps) {
    const dialogRef = useRef<HTMLDialogElement>(null);
    const resolverRef = useRef<((v: UploadImageValue | null) => void) | null>(null);

    const [isOpen, setIsOpen] = useState(false);
    const [session, setSession] = useState(0); // ganti key agar UploadImage mulai bersih
    const [options, setOptions] = useState<OpenOptions>({});
    const [value, setValue] = useState<UploadImageValue | null>(null);

    const settle = useCallback((v: UploadImageValue | null) => {
        resolverRef.current?.(v);
        resolverRef.current = null;
    }, []);

    const open = useCallback<OpenImageUploader>(
        (opts = {}) => {
            settle(null); // jika masih ada pemanggilan sebelumnya, anggap dibatalkan
            return new Promise((resolve) => {
                resolverRef.current = resolve;
                setOptions(opts);
                setValue(null);
                setSession((s) => s + 1);
                setIsOpen(true);
            });
        },
        [settle],
    );

    // Daftarkan ke singleton
    useEffect(() => {
        registry.open = open;
        registry.waiters.splice(0).forEach((run) => run());
        return () => {
            if (registry.open === open) registry.open = null;
            settle(null);
        };
    }, [open, settle]);

    // Sinkronkan state dengan <dialog> native (focus trap, Esc, backdrop)
    useEffect(() => {
        const d = dialogRef.current;
        if (!d) return;
        if (isOpen && !d.open) d.showModal();
        if (!isOpen && d.open) d.close();
    }, [isOpen]);

    function confirm() {
        if (!value) return;
        settle(value);
        setIsOpen(false);
    }

    const heading = options.title ?? "Unggah gambar";

    return (
        <ImageUploaderContext.Provider value={open}>
            {children}

            <dialog
                ref={dialogRef}
                aria-labelledby="image-uploader-title"
                onClose={() => {
                    setIsOpen(false);
                    settle(null); // Esc / tombol tutup
                }}
                onClick={(e) => {
                    // klik di area backdrop menutup dialog
                    if (e.target === e.currentTarget) e.currentTarget.close();
                }}
                className="m-auto w-[calc(100%-2rem)] max-w-xl rounded-card bg-white p-0 text-text shadow-card backdrop:bg-text/50 backdrop:backdrop-blur-sm"
            >
                {isOpen && (
                    <div className="grid gap-4 p-5 font-sans">
                        <div className="flex items-start justify-between gap-4">
                            <h2 id="image-uploader-title" className="text-lg font-semibold tracking-tight">
                                {heading}
                            </h2>
                            <button
                                type="button"
                                aria-label="Tutup"
                                onClick={() => dialogRef.current?.close()}
                                className="grid size-8 place-items-center rounded-pill text-text-muted transition-colors hover:bg-bg hover:text-text focus-visible:outline-3 focus-visible:outline-primary/50"
                            >
                                <svg viewBox="0 0 20 20" className="size-4" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" aria-hidden="true">
                                    <path d="M5 5l10 10M15 5L5 15" />
                                </svg>
                            </button>
                        </div>

                        <UploadImage
                            key={session}
                            bare
                            endpoint={options.endpoint ?? endpoint}
                            fieldName={options.fieldName ?? fieldName}
                            maxSizeMB={options.maxSizeMB ?? maxSizeMB}
                            onChange={setValue}
                            onReset={() => setValue(null)}
                        />

                        <div className="flex flex-wrap justify-end gap-2.5 border-t border-text-muted/20 pt-4 max-md:[&>button]:basis-full">
                            <button
                                type="button"
                                onClick={() => dialogRef.current?.close()}
                                className={`${btn} border border-text-muted/30 bg-white text-text hover:bg-bg`}
                            >
                                Batal
                            </button>
                            <button
                                type="button"
                                disabled={!value}
                                onClick={confirm}
                                className={`${btn} border border-primary bg-primary font-semibold text-white hover:enabled:bg-primary/90`}
                            >
                                Pakai gambar ini
                            </button>
                        </div>
                    </div>
                )}
            </dialog>
        </ImageUploaderContext.Provider>
    );
}

export default ImageUploaderProvider;