import SocialLinks from "./const";

type HeroProps = {
    /** Foto latar (akan ditampilkan grayscale) */
    backgroundSrc: string;
    /** Foto subjek dengan latar transparan (PNG/WebP cutout), tetap berwarna */
    portraitSrc: string;
    portraitAlt?: string;
    primaryHref?: string;
    secondaryHref?: string;
};

export default function Hero({
    backgroundSrc,
    portraitSrc,
    portraitAlt = "Puguh Pamungkas tersenyum, mengenakan kemeja gelap",
    primaryHref = "#produk",
    secondaryHref = "#tentang",
}: HeroProps) {
    return (
        <section
            id="beranda"
            className="relative overflow-hidden bg-bg"
        >
            <div className="mx-auto grid max-w-container items-center px-(--container-pad) lg:min-h-160 lg:grid-cols-12">
                {/* Teks */}
                <div className="relative z-10 order-2 py-12 lg:order-1 lg:col-span-6 lg:py-24">
                    <p className="mb-5 text-(length:--fs-eyebrow) font-medium uppercase tracking-eyebrow text-muted">
                        Teknologi untuk Pendidikan &amp; Masyarakat
                    </p>

                    <h1 className="max-w-[16ch] text-4xl font-bold leading-[1.1] tracking-tight text-ink sm:text-5xl lg:text-[3.25rem]">
                        Membangun Solusi Teknologi, Menguatkan Manusia.
                    </h1>

                    <p className="mt-6 max-w-[48ch] text-base leading-relaxed text-muted">
                        Saya percaya bahwa teknologi akan bermakna lebih besar ketika bisa
                        diakses oleh lebih banyak orang dan memberi dampak nyata bagi
                        masyarakat.
                    </p>

                    <div className="mt-8 flex flex-col gap-4 sm:flex-row sm:items-center">
                        <a
                            href={primaryHref}
                            className="inline-flex items-center justify-center rounded-full bg-primary px-7 py-3.5 text-sm font-semibold text-white shadow-cta transition-colors hover:bg-primary-hover  focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary"
                        >
                            Lihat Karya &amp; Dampak
                        </a>

                        <a
                            href={secondaryHref}
                            className="group inline-flex items-center justify-center gap-3 rounded-full px-2 py-2 text-sm font-medium text-ink  focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary"
                        >
                            Tentang Saya
                            <span
                                aria-hidden
                                className="grid h-8 w-8 place-items-center rounded-full border border-border transition-colors group-hover:border-primary group-hover:text-primary"
                            >
                                <svg
                                    width="14"
                                    height="14"
                                    viewBox="0 0 24 24"
                                    fill="none"
                                    stroke="currentColor"
                                    strokeWidth="2"
                                    strokeLinecap="round"
                                    strokeLinejoin="round"
                                >
                                    <path d="M5 12h14M13 6l6 6-6 6" />
                                </svg>
                            </span>
                        </a>
                    </div>
                    <SocialLinks />
                </div>

                {/* Foto */}
                <div className="relative order-1 -mx-(--container-pad) h-95 sm:h-115 lg:order-2 lg:col-span-6 lg:mx-0 lg:h-full lg:min-h-160">
                    {/* Latar grayscale, memudar ke kiri */}
                    <div
                        aria-hidden
                        className="absolute inset-0 lg:-right-[30vw] lg:left-[-6rem]
        [-webkit-mask-image:linear-gradient(to_bottom,black_70%,transparent)]
        [mask-image:linear-gradient(to_bottom,black_70%,transparent)]
        lg:[-webkit-mask-image:linear-gradient(to_right,transparent,black_45%)]
        lg:[mask-image:linear-gradient(to_right,transparent,black_45%)]"
                    >
                        <img
                            src={backgroundSrc}
                            alt=""
                            className="h-full w-full object-cover grayscale"
                            fetchPriority="high"
                        />
                    </div>

                    {/* Bentuk organik teal */}
                    <svg
                        aria-hidden
                        viewBox="0 0 400 400"
                        className="absolute bottom-[40%] left-[38%] w-[55%] max-w-110 text-[color-mix(in_srgb,var(--color-accent-teal)_40%,transparent)]"
                    >
                        <path
                            fill="currentColor"
                            d="M318 62c47 38 78 102 62 160-16 58-79 110-143 126-64 16-130-4-171-52C25 248 12 174 46 118 80 62 154 24 218 22c34-1 75 8 100 40z"
                        />
                    </svg>
                    <svg
                        aria-hidden
                        viewBox="0 0 400 400"
                        className="absolute bottom-[5%] left-[24%] w-[55%] max-w-110 text-[color-mix(in_srgb,var(--color-accent-teal)_40%,transparent)]"
                    >
                        <path
                            fill="currentColor"
                            d="M318 62c47 38 78 102 62 160-16 58-79 110-143 126-64 16-130-4-171-52C25 248 12 174 46 118 80 62 154 24 218 22c34-1 75 8 100 40z"
                        />
                    </svg>
                    <svg
                        aria-hidden
                        viewBox="0 0 400 400"
                        className="absolute bottom-0 left-[48%] w-[55%] max-w-110 text-[color-mix(in_srgb,var(--color-accent-teal)_40%,transparent)]"
                    >
                        <path
                            fill="currentColor"
                            d="M318 62c47 38 78 102 62 160-16 58-79 110-143 126-64 16-130-4-171-52C25 248 12 174 46 118 80 62 154 24 218 22c34-1 75 8 100 40z"
                        />
                    </svg>

                    {/* Subjek, sedikit offset ke kanan */}
                    <img
                        src={portraitSrc}
                        alt={portraitAlt}
                        className="absolute -bottom-30 h-[120%] right-[10%] lg:left-[30%] w-auto max-w-none object-contain object-center"
                        fetchPriority="high"
                    />
                </div>
            </div>
        </section>
    );
}