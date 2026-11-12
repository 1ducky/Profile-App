import {
    useRef,
    type ReactNode,
} from "react";
import { useDragSlider } from "../../hooks/useDrag";


/* ------------------------------------------------------------------ */
/* Tipe & data                                                         */
/* ------------------------------------------------------------------ */

type Stat = { value: string; label: string };

type MiniCard = {
    icon: ReactNode;
    title: string;
    description: string;
    href: string;
};

export type ProductSlide = {
    id: string;
    name: string;
    logoSrc: string;
    logoAlt?: string;
    tagline: string;
    description: string;
    href: string;
    imageSrc: string;
    imageAlt: string;
    stats: Stat[];
    cards: MiniCard[];
};

const iconProps = {
    width: 22,
    height: 22,
    viewBox: "0 0 24 24",
    fill: "none",
    stroke: "currentColor",
    strokeWidth: 1.6,
    strokeLinecap: "round" as const,
    strokeLinejoin: "round" as const,
    "aria-hidden": true,
};

const IconChip = () => (
    <svg {...iconProps}>
        <rect x="6" y="6" width="12" height="12" rx="2" />
        <path d="M9 2v4M15 2v4M9 18v4M15 18v4M2 9h4M2 15h4M18 9h4M18 15h4" />
    </svg>
);
const IconUsers = () => (
    <svg {...iconProps}>
        <circle cx="9" cy="8" r="3.5" />
        <path d="M2.5 20c0-3.6 2.9-6 6.5-6s6.5 2.4 6.5 6" />
        <path d="M16 4.6a3.5 3.5 0 0 1 0 6.8M18.5 14.4c1.8.8 3 2.6 3 5.6" />
    </svg>
);
const IconSpark = () => (
    <svg {...iconProps}>
        <path d="M12 3v4M12 17v4M3 12h4M17 12h4M6 6l2.5 2.5M15.5 15.5L18 18M18 6l-2.5 2.5M8.5 15.5L6 18" />
    </svg>
);

/** Ganti dengan data asli. Tiap item = satu "lembar" slider. */
export const DEFAULT_SLIDES: ProductSlide[] = [
    {
        id: "algebra",
        name: "Algebra Edu",
        logoSrc: "/assets/favicon.png",
        logoAlt: "Logo Algebra",
        tagline: "Digitalisasi Pendidikan",
        description:
            "Platform pembelajaran dan manajemen pendidikan untuk sekolah, guru, dan siswa.",
        href: "#",
        imageSrc: "/assets/preview/web-prev-1.png",
        imageAlt: "Tampilan dasbor Algebra di laptop dan tablet",
        stats: [
            { value: "+10K", label: "Pengguna aktif" },
            { value: "+500", label: "Sekolah mitra" },
            { value: "+20", label: "Provinsi" },
        ],
        cards: [
            {
                icon: <IconChip />,
                title: "Proyek Teknologi Lainnya",
                description:
                    "Aplikasi, sistem dan solusi digital untuk berbagai kebutuhan.",
                href: "#",
            },
            {
                icon: <IconUsers />,
                title: "Kolaborasi & Kemitraan",
                description: "Membangun ekosistem yang lebih kuat.",
                href: "#",
            },
            {
                icon: <IconSpark />,
                title: "Inovasi Berkelanjutan",
                description: "Mengembangkan solusi yang relevan dan berdampak jangka panjang.",
                href: "#",
            },
        ],
    },
    // CONTOH slide kedua: ganti/hapus sesuai kebutuhan.
    {
        id: "produk-2",
        name: "Nama Produk",
        logoSrc: "/assets/favicon.png",
        tagline: "Kategori Produk",
        description: "Deskripsi singkat produk atau karya berikutnya.",
        href: "#",
        imageSrc: "/assets/preview/web-prev-2.png",
        imageAlt: "Tampilan produk",
        stats: [
            { value: "+0", label: "Metrik 1" },
            { value: "+0", label: "Metrik 2" },
            { value: "+0", label: "Metrik 3" },
        ],
        cards: [
            { icon: <IconChip />, title: "Judul Kartu 1", description: "Deskripsi singkat.", href: "#" },
            { icon: <IconUsers />, title: "Judul Kartu 2", description: "Deskripsi singkat.", href: "#" },
            { icon: <IconSpark />, title: "Judul Kartu 3", description: "Deskripsi singkat.", href: "#" },
        ],
    },
];

/* ------------------------------------------------------------------ */
/* Ikon panah                                                          */
/* ------------------------------------------------------------------ */

const Arrow = ({ dir = "right" }: { dir?: "left" | "right" }) => (
    <svg
        width="14"
        height="14"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        strokeWidth="2"
        strokeLinecap="round"
        strokeLinejoin="round"
        aria-hidden
        className={dir === "left" ? "rotate-180" : undefined}
    >
        <path d="M5 12h14M13 6l6 6-6 6" />
    </svg>
);

/* ------------------------------------------------------------------ */
/* Satu lembar slide                                                   */
/* ------------------------------------------------------------------ */

function Slide({ slide }: { slide: ProductSlide }) {
    return (
        <div className="space-y-5">
            {/* Kartu unggulan */}
            <article className="grid items-center gap-8 rounded-card border border-border bg-white p-6 shadow-card md:grid-cols-2 md:p-10">
                <div>
                    <div className="flex items-center gap-3">
                        <img
                            src={slide.logoSrc}
                            alt={slide.logoAlt ?? ""}
                            className="h-10 w-10 object-contain"
                            draggable={false}
                        />
                        <h3 className="text-3xl font-semibold tracking-tight text-ink">
                            {slide.name}
                        </h3>
                    </div>
                    <p className="mt-5 text-sm font-semibold text-ink">{slide.tagline}</p>
                    <p className="mt-2 max-w-[44ch] text-sm leading-relaxed text-muted">
                        {slide.description}
                    </p>
                    <a
                        href={slide.href}
                        draggable={false}
                        className="mt-6 inline-flex items-center gap-2 rounded-full border border-border bg-white px-5 py-2.5 text-sm font-medium text-primary transition-colors hover:border-primaryfocus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary"
                    >
                        Pelajari lebih lanjut <Arrow />
                    </a>
                </div>

                <div>
                    <img
                        src={slide.imageSrc}
                        alt={slide.imageAlt}
                        loading="lazy"
                        draggable={false}
                        className="mx-auto w-full max-w-115 object-contain"
                    />
                    <dl className="mt-6 grid grid-cols-3 gap-2 text-center">
                        {slide.stats.map((s) => (
                            <div key={s.label}>
                                <dt className="sr-only">{s.label}</dt>
                                <dd className="text-xl font-semibold text-ink">{s.value}</dd>
                                <p aria-hidden className="mt-0.5 text-xs text-muted">
                                    {s.label}
                                </p>
                            </div>
                        ))}
                    </dl>
                </div>
            </article>

            {/* Tiga kartu kecil */}
            <div className="grid gap-5 md:grid-cols-3">
                {slide.cards.map((c) => (
                    <a
                        key={c.title}
                        href={c.href}
                        draggable={false}
                        className="group block rounded-card border border-border bg-white p-5 shadow-card transition-colors hover:border-primaryfocus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary"
                    >
                        <div className="flex items-start justify-between text-ink">
                            {c.icon}
                            <span className="grid h-7 w-7 place-items-center rounded-full border border-border text-primary transition-colors group-hover:border-primary">
                                <Arrow />
                            </span>
                        </div>
                        <h4 className="mt-4 text-sm font-semibold text-ink">{c.title}</h4>
                        <p className="mt-1 text-xs leading-relaxed text-muted">
                            {c.description}
                        </p>
                    </a>
                ))}
            </div>
        </div>
    );
}

/* ------------------------------------------------------------------ */
/* Slider (overflow-hidden, track flex, geser dengan pointer)          */
/* ------------------------------------------------------------------ */

function Slider({ slides }: { slides: ProductSlide[] }) {
    const drag = useRef({ startX: 0, startT: 0, dx: 0, active: false, moved: false });
    const count = slides.length;
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
                className="-mx-3 overflow-hidden"
                tabIndex={0}
                onKeyDown={onKeyDown}
            >
                {/* Track: children memanjang ke samping */}
                <div
                    className={`flex cursor-grab touch-pan-y select-none active:cursor-grabbing ${dragging
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
                    {slides.map((slide, i) => (
                        <div
                            key={slide.id}
                            role="group"
                            aria-roledescription="slide"
                            aria-label={`${i + 1} dari ${count}`}
                            // slide yang tidak aktif tidak bisa difokus
                            {...(i !== index ? ({ inert: "" } as object) : {})}
                            className="basis-full shrink-0 grow-0 p-3"
                        >
                            <Slide slide={slide} />
                        </div>
                    ))}
                </div>
            </div>

            {/* Kontrol */}
            {count > 1 && (
                <div className="mt-6 flex items-center justify-between">
                    <div className="flex items-center gap-2">
                        {slides.map((s, i) => (
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
                            <Arrow dir="left" />
                        </button>
                        <button
                            type="button"
                            onClick={() => go(index + 1)}
                            disabled={index === count - 1}
                            aria-label="Slide berikutnya"
                            className="grid h-9 w-9 place-items-center rounded-full border border-border bg-white text-ink transition-colors hover:border-primary disabled:opacity-40 disabled:hover:border-border focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary"
                        >
                            <Arrow />
                        </button>
                    </div>
                </div>
            )}

            <p className="sr-only" aria-live="polite">
                Slide {index + 1} dari {count}
            </p>
        </div>
    );
}

/* ------------------------------------------------------------------ */
/* Section                                                             */
/* ------------------------------------------------------------------ */

export default function ProductSection({
    slides = DEFAULT_SLIDES,
    allHref = "#",
}: {
    slides?: ProductSlide[];
    allHref?: string;
}) {
    return (
        <section
            id="produk"
            className="bg-bg py-(--section-space)"
        >
            <div className="mx-auto max-w-container px-(--container-pad)">
                <header className="mb-10 flex flex-col gap-4 md:flex-row md:items-end md:justify-between">
                    <div>
                        <p className="mb-3 text-(--fs-eyebrow) font-medium uppercase tracking-eyebrow text-muted">
                            Fokus Utama
                        </p>
                        <h2 className="text-(--fs-h2) font-semibold tracking-tight text-ink">
                            Produk &amp; Karya
                        </h2>
                        <p className="mt-3 max-w-[56ch] text-sm leading-relaxed text-muted">
                            Menghadirkan solusi teknologi yang berdampak, mulai dari
                            digitalisasi pendidikan hingga berbagai inisiatif inovasi lainnya.
                        </p>
                    </div>
                    <a
                        href={allHref}
                        className="inline-flex items-center gap-2 text-sm font-medium text-primary hover:underline focus-visible:outline-2 focus-visible:outline-offset-4 focus-visible:outline-primary"
                    >
                        Lihat semua <Arrow />
                    </a>
                </header>

                <Slider slides={slides} />
            </div>
        </section>
    );
}