import { useDragSlider } from "../../hooks/useDrag";



type Photo = { src: string; alt: string };

export type ImpactSlide = {
    id: string;
    story: { title: string; description: string };
    /** [foto besar, foto kecil atas, foto kecil bawah] */
    photos: [Photo, Photo, Photo];
};

/** Ganti dengan dokumentasi kegiatan asli. */
export const DEFAULT_IMPACT_SLIDES: ImpactSlide[] = [
    {
        id: "literasi-digital",
        story: {
            title: "Pendidikan & Literasi Digital",
            description:
                "Membantu akses pendidikan dan literasi digital di berbagai daerah.",
        },
        photos: [
            { src: "/assets/preview/mobile-prev-1.png", alt: "Puguh berbincang dengan anak-anak di desa saat kegiatan belajar" },
            { src: "/assets/preview/mobile-prev-2.png", alt: "Pelatihan literasi digital bersama warga" },
            { src: "/assets/preview/mobile-prev-3.png", alt: "Warga berkumpul dalam kegiatan komunitas" },
        ],
    },
    // CONTOH: tambahkan lembar lain sesuai kegiatan
    {
        id: "kegiatan-2",
        story: {
            title: "Judul Kegiatan",
            description: "Deskripsi singkat kegiatan dan dampak yang dicapai.",
        },
        photos: [
            { src: "/assets/preview/web-prev-1.png", alt: "Deskripsi foto kegiatan" },
            { src: "/assets/preview/web-prev-2.png", alt: "Deskripsi foto kegiatan" },
            { src: "/assets/preview/web-prev-3.png", alt: "Deskripsi foto kegiatan" },
        ],
    },
];

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

const IconBook = () => (
    <svg
        width="22"
        height="22"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        strokeWidth="1.6"
        strokeLinecap="round"
        strokeLinejoin="round"
        aria-hidden
    >
        <path d="M4 5.5A2.5 2.5 0 0 1 6.5 3H20v16H6.5A2.5 2.5 0 0 0 4 21.5z" />
        <path d="M4 19V5.5M9 8h6" />
    </svg>
);

const photoClass =
    "h-full w-full rounded-card object-cover";

export default function SocialImpactSection({
    slides = DEFAULT_IMPACT_SLIDES,
    allHref = "#",
}: {
    slides?: ImpactSlide[];
    allHref?: string;
}) {
    const count = slides.length;
    const { viewportRef, index, go, dragging, onKeyDown, trackProps } =
        useDragSlider(count);
    const story = slides[index].story;

    return (
        <section id="dampak-sosial" className="bg-warm py-(--section-space)">
            <div className="mx-auto grid max-w-container items-center gap-10 px-(--container-pad) lg:grid-cols-12 lg:gap-12">
                {/* Teks */}
                <div className="lg:col-span-5">
                    <p className="mb-3 text-[length(var(--fs-eyebrow))] font-medium uppercase tracking-eyebrow text-muted">
                        Pengabdian
                    </p>
                    <h2 className="text-(length:--fs-h2) font-semibold tracking-tight text-ink">
                        Dampak Sosial
                    </h2>
                    <p className="mt-3 max-w-[44ch] text-sm leading-relaxed text-muted">
                        Turun langsung ke masyarakat, mendengarkan, belajar, dan
                        berkontribusi bersama untuk perubahan yang lebih baik.
                    </p>
                    <a
                        href={allHref}
                        className="mt-5 inline-flex items-center gap-2 text-sm font-medium text-primary hover:underline focus-visible:outline-2 focus-visible:outline-offset-4 focus-visible:outline-primary"
                    >
                        Lihat semua <Arrow />
                    </a>

                    {/* Cerita singkat: mengikuti slide aktif */}
                    <div
                        className="mt-10 flex items-start gap-4"
                        aria-live="polite"
                    >
                        <span className="grid h-11 w-11 shrink-0 place-items-center rounded-xl bg-white text-primary shadow-card">
                            <IconBook />
                        </span>
                        <div>
                            <h3 className="text-sm font-semibold text-ink">{story.title}</h3>
                            <p className="mt-1 max-w-[36ch] text-xs leading-relaxed text-muted">
                                {story.description}
                            </p>
                        </div>
                    </div>
                </div>

                {/* Galeri carousel */}
                <div
                    className="lg:col-span-7"
                    role="region"
                    aria-roledescription="carousel"
                    aria-label="Galeri dampak sosial"
                >
                    <div
                        ref={viewportRef}
                        className="overflow-hidden rounded-card"
                        tabIndex={0}
                        onKeyDown={onKeyDown}
                    >
                        <div
                            className={`flex cursor-grab touch-pan-y select-none active:cursor-grabbing ${dragging
                                ? "transition-none"
                                : "transition-transform duration-500 ease-soft motion-reduce:transition-none"
                                }`}
                            {...trackProps}
                        >
                            {slides.map((s, i) => (
                                <div
                                    key={s.id}
                                    role="group"
                                    aria-roledescription="slide"
                                    aria-label={`${i + 1} dari ${count}`}
                                    {...(i !== index ? ({ inert: "" } as object) : {})}
                                    className="basis-full shrink-0 grow-0 p-5"
                                >
                                    <div className="grid h-[260px] grid-cols-3 grid-rows-2 gap-3 sm:h-85 lg:h-95">
                                        <img
                                            src={s.photos[0].src}
                                            alt={s.photos[0].alt}
                                            loading="lazy"
                                            draggable={false}
                                            className={`${photoClass} col-span-2 row-span-2`}
                                        />
                                        <img
                                            src={s.photos[1].src}
                                            alt={s.photos[1].alt}
                                            loading="lazy"
                                            draggable={false}
                                            className={photoClass}
                                        />
                                        <img
                                            src={s.photos[2].src}
                                            alt={s.photos[2].alt}
                                            loading="lazy"
                                            draggable={false}
                                            className={photoClass}
                                        />
                                    </div>
                                </div>
                            ))}
                        </div>
                    </div>

                    {count > 1 && (
                        <div className="mt-5 flex items-center justify-between">
                            <div className="flex items-center gap-2">
                                {slides.map((s, i) => (
                                    <button
                                        key={s.id}
                                        type="button"
                                        onClick={() => go(i)}
                                        aria-label={`Ke slide ${i + 1}`}
                                        aria-current={i === index}
                                        className={`h-2 rounded-full transition-all focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary ${i === index ? "w-6 bg-primary" : "w-2 bg-muted/40 hover:bg-muted"
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
                </div>
            </div>
        </section>
    );
}