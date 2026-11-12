import {
    createContext,
    useCallback,
    useContext,
    useEffect,
    useRef,
    useState,
    type ReactNode,
} from "react";
import { CircleHelp, Check, X, Timer } from "lucide-react";

/**
 * Cara pakai singkat:
 *
 *   const confirm = useAlert();
 *   const ok = await confirm({ title: "Hapus data ini?", timer: 10 });
 *   if (ok) { ... }          // true = Ya, false = Tidak / timer habis / ditutup
 *
 * Di Astro, pasang SEKALI di layout:
 *   <AlertProvider client:load />
 * Island lain (React terpisah) tetap bisa memanggil lewat `useAlert()`
 * atau lewat `alert()` langsung — keduanya memakai provider yang sama.
 */

/* ---------------------------------- Types --------------------------------- */

export type AlertOptions = {
    /** Judul / pertanyaan yang ditampilkan */
    title: string;
    /** Opsional: batas waktu dalam detik. Jika habis, hasilnya `false` (Tidak). */
    timer?: number;
};

export type AlertFn = (options: AlertOptions | string) => Promise<boolean>;

type AlertItem = AlertOptions & {
    id: number;
    resolve: (value: boolean) => void;
};

/* ---------- Singleton: menjembatani island Astro yang terpisah ---------- */

// Disimpan di globalThis agar tetap satu instance walau modul ini ter-bundle
// ganda (HMR / chunk berbeda antar island).
type Registry = { open: AlertFn | null; waiters: Array<() => void> };
const g = globalThis as typeof globalThis & { __alert?: Registry };
const registry: Registry = (g.__alert ??= { open: null, waiters: [] });

const WAIT_MS = 3000;

export const alert: AlertFn = (options) => {
    if (registry.open) return registry.open(options);

    // Provider mungkin belum selesai hydrate: tunggu sebentar sebelum menyerah.
    return new Promise<boolean>((resolve, reject) => {
        const run = () => {
            clearTimeout(timer);
            registry.open!(options).then(resolve, reject);
        };
        const timer = setTimeout(() => {
            registry.waiters = registry.waiters.filter((w) => w !== run);
            reject(
                new Error(
                    "AlertProvider belum terpasang. Pastikan <AlertProvider client:load /> ada di layout halaman ini."
                )
            );
        }, WAIT_MS);
        registry.waiters.push(run);
    });
};

/* ---------- Context (untuk komponen di dalam pohon yang sama) ---------- */

const AlertContext = createContext<AlertFn | null>(null);

export function useAlert(): AlertFn {
    const ctx = useContext(AlertContext);
    // Di luar pohon provider (island lain) otomatis memakai singleton.
    return ctx ?? alert;
}

/* -------------------------------- Provider -------------------------------- */

export function AlertProvider({ children }: { children?: ReactNode }) {
    // Antrean: jika alert dipanggil saat yang lain masih terbuka, ditampilkan bergiliran.
    const [queue, setQueue] = useState<AlertItem[]>([]);
    const idRef = useRef(0);

    const show = useCallback<AlertFn>((options) => {
        const opts = typeof options === "string" ? { title: options } : options;
        return new Promise<boolean>((resolve) => {
            const item: AlertItem = { ...opts, id: ++idRef.current, resolve };
            setQueue((q) => [...q, item]);
        });
    }, []);

    // Daftarkan ke singleton
    useEffect(() => {
        registry.open = show;
        registry.waiters.splice(0).forEach((run) => run());
        return () => {
            if (registry.open === show) registry.open = null;
        };
    }, [show]);

    const current = queue[0];

    const close = useCallback(
        (result: boolean) => {
            if (!current) return;
            current.resolve(result);
            setQueue((q) => q.slice(1));
        },
        [current]
    );

    return (
        <AlertContext.Provider value={show}>
            {children}
            {current && (
                <AlertDialog
                    key={current.id}
                    title={current.title}
                    timer={current.timer}
                    onClose={close}
                />
            )}
        </AlertContext.Provider>
    );
}

export default AlertProvider;

/* --------------------------------- Dialog --------------------------------- */

function AlertDialog({
    title,
    timer,
    onClose,
}: {
    title: string;
    timer?: number;
    onClose: (result: boolean) => void;
}) {
    const hasTimer = typeof timer === "number" && timer > 0;
    const totalMs = hasTimer ? timer! * 1000 : 0;
    const [remainingMs, setRemainingMs] = useState(totalMs);
    const noRef = useRef<HTMLButtonElement>(null);
    const closedRef = useRef(false);

    const finish = useCallback(
        (result: boolean) => {
            if (closedRef.current) return; // cegah resolve ganda
            closedRef.current = true;
            onClose(result);
        },
        [onClose]
    );

    // Timer hitung mundur
    useEffect(() => {
        if (!hasTimer) return;
        const deadline = Date.now() + totalMs;
        const interval = setInterval(() => {
            const left = deadline - Date.now();
            if (left <= 0) {
                clearInterval(interval);
                setRemainingMs(0);
                finish(false);
            } else {
                setRemainingMs(left);
            }
        }, 100);
        return () => clearInterval(interval);
    }, [hasTimer, totalMs, finish]);

    // Fokus awal ke "Tidak" (pilihan paling aman) + Escape = Tidak + kunci scroll
    useEffect(() => {
        noRef.current?.focus();
        const onKey = (e: KeyboardEvent) => {
            if (e.key === "Escape") finish(false);
        };
        window.addEventListener("keydown", onKey);
        const prevOverflow = document.body.style.overflow;
        document.body.style.overflow = "hidden";
        return () => {
            window.removeEventListener("keydown", onKey);
            document.body.style.overflow = prevOverflow;
        };
    }, [finish]);

    const progress = hasTimer ? (remainingMs / totalMs) * 100 : 0;
    const secondsLeft = Math.ceil(remainingMs / 1000);

    return (
        <div
            className="fixed inset-0 z-[100] flex items-center justify-center bg-[#1F2937]/50 p-4 backdrop-blur-sm animate-[fadeIn_150ms_ease-out] motion-reduce:animate-none"
            onMouseDown={(e) => {
                if (e.target === e.currentTarget) finish(false);
            }}
        >
            <style>{`
        @keyframes fadeIn { from { opacity: 0 } to { opacity: 1 } }
        @keyframes popIn { from { opacity: 0; transform: translateY(8px) scale(.98) } to { opacity: 1; transform: none } }
      `}</style>

            <div
                role="alertdialog"
                aria-modal="true"
                aria-labelledby="alert-title"
                className="relative w-full max-w-md overflow-hidden rounded-[16px] border border-slate-200 bg-white shadow-[0_4px_16px_rgba(31,41,55,0.06),0_24px_48px_rgba(31,41,55,0.18)] animate-[popIn_200ms_ease-out] motion-reduce:animate-none"
            >
                <div className="flex items-start gap-4 p-6">
                    <span className="flex h-11 w-11 shrink-0 items-center justify-center rounded-full bg-[#0B63CE]/10 text-[#0B63CE]">
                        <CircleHelp className="h-6 w-6" aria-hidden="true" />
                    </span>

                    <div className="min-w-0 flex-1 pt-0.5">
                        <h2
                            id="alert-title"
                            className="text-lg font-semibold leading-snug tracking-tight text-[#1F2937]"
                        >
                            {title}
                        </h2>

                        {hasTimer && (
                            <p className="mt-2 inline-flex items-center gap-1.5 text-sm text-[#6B7280]">
                                <Timer className="h-4 w-4" aria-hidden="true" />
                                Otomatis menutup dalam{" "}
                                <span className="font-semibold tabular-nums text-[#1F2937]">
                                    {secondsLeft} dtk
                                </span>
                            </p>
                        )}
                    </div>
                </div>

                <div className="flex flex-col-reverse gap-2 px-6 pb-6 sm:flex-row sm:justify-end">
                    <button
                        ref={noRef}
                        type="button"
                        onClick={() => finish(false)}
                        className="inline-flex items-center justify-center gap-2 rounded-full border border-slate-300 bg-white px-5 py-2.5 text-sm font-medium text-[#1F2937] transition hover:bg-slate-50 focus:outline-none focus-visible:ring-2 focus-visible:ring-[#0B63CE] focus-visible:ring-offset-2"
                    >
                        <X className="h-4 w-4" aria-hidden="true" />
                        Tidak
                    </button>
                    <button
                        type="button"
                        onClick={() => finish(true)}
                        className="inline-flex items-center justify-center gap-2 rounded-full bg-[#0B63CE] px-5 py-2.5 text-sm font-medium text-white transition hover:bg-[#0953ad] focus:outline-none focus-visible:ring-2 focus-visible:ring-[#0B63CE] focus-visible:ring-offset-2"
                    >
                        <Check className="h-4 w-4" aria-hidden="true" />
                        Ya
                    </button>
                </div>

                {hasTimer && (
                    <div className="h-1 w-full bg-slate-100" aria-hidden="true">
                        <div
                            className="h-full bg-[#1AA6B7]"
                            style={{ width: `${progress}%`, transition: "width 100ms linear" }}
                        />
                    </div>
                )}
            </div>
        </div>
    );
}