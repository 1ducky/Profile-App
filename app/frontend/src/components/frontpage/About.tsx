import type { ReactNode } from "react";

type Value = { icon: ReactNode; title: string; description: string };

const svg = {
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

const VALUES: Value[] = [
    {
        icon: (
            <svg {...svg}>
                <path d="M8 8l-4 4 4 4M16 8l4 4-4 4M13.5 5l-3 14" />
            </svg>
        ),
        title: "IT Expert & Product Builder",
        description:
            "Berpengalaman membangun berbagai solusi teknologi dan produk digital.",
    },
    {
        icon: (
            <svg {...svg}>
                <path d="M12 21s-7-4.4-7-10.2A4.3 4.3 0 0 1 9.3 6.5c1.2 0 2.2.5 2.7 1.4.5-.9 1.5-1.4 2.7-1.4A4.3 4.3 0 0 1 19 10.8C19 16.6 12 21 12 21z" />
            </svg>
        ),
        title: "Pembelajar Sepanjang Hayat",
        description:
            "Selalu terbuka untuk belajar hal baru dari siapa saja dan di mana saja.",
    },
    {
        icon: (
            <svg {...svg}>
                <circle cx="9" cy="8" r="3.5" />
                <path d="M2.5 20c0-3.6 2.9-6 6.5-6s6.5 2.4 6.5 6M17 11l2 2 3-3.5" />
            </svg>
        ),
        title: "Aktivis Sosial",
        description:
            "Terlibat langsung di masyarakat, terutama di bidang pendidikan dan pemberdayaan.",
    },
];

export default function AboutSection({ href = "#" }: { href?: string }) {
    return (
        <section id="tentang" className="bg-bg py-(--section-space)">
            <div className="mx-auto grid max-w-container gap-12 px-(--container-pad) lg:grid-cols-12 lg:items-center">
                <div className="lg:col-span-6">
                    <p className="mb-3 text-(length:--fs-eyebrow) font-medium uppercase tracking-eyebrow text-muted">
                        Siapa Saya
                    </p>
                    <h2 className="text-(length:--fs-h2) font-semibold tracking-tight text-ink">
                        Tentang Saya
                    </h2>
                    <p className="mt-4 max-w-[48ch] text-base leading-relaxed text-muted">
                        Seorang teknolog, pembelajar, dan pegiat sosial yang percaya bahwa
                        kolaborasi adalah kunci untuk menciptakan dampak yang lebih besar.
                    </p>
                    <a
                        href={href}
                        className="mt-8 inline-flex items-center justify-center gap-2 rounded-full bg-primary px-7 py-3.5 text-sm font-semibold text-white shadow-cta transition-colors hover:bg-primary-hover focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary max-sm:w-full"
                    >
                        Selengkapnya tentang saya
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
                        >
                            <path d="M5 12h14M13 6l6 6-6 6" />
                        </svg>
                    </a>
                </div>

                <ul className="space-y-7 lg:col-span-6 lg:pl-8">
                    {VALUES.map((v) => (
                        <li key={v.title} className="flex items-start gap-4">
                            <span className="grid h-11 w-11 shrink-0 place-items-center rounded-full border border-border bg-white text-ink">
                                {v.icon}
                            </span>
                            <div>
                                <h3 className="text-sm font-semibold text-ink">{v.title}</h3>
                                <p className="mt-1 max-w-[46ch] text-sm leading-relaxed text-muted">
                                    {v.description}
                                </p>
                            </div>
                        </li>
                    ))}
                </ul>
            </div>
        </section>
    );
}