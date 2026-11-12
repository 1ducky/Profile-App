export type CareerItem = {
    period: string;
    role: string;
    organization: string;
    description: string;
    /** Tandai posisi saat ini */
    current?: boolean;
};

/** CONTOH data. Ganti dengan riwayat karir asli (urut dari terbaru). */
export const DEFAULT_CAREER: CareerItem[] = [
    {
        period: "2022 — Sekarang",
        role: "Jabatan / Peran Saat Ini",
        organization: "Nama Organisasi",
        description:
            "Ringkasan tanggung jawab utama dan pencapaian paling berdampak pada peran ini.",
        current: true,
    },
    {
        period: "2019 — 2022",
        role: "Jabatan / Peran Sebelumnya",
        organization: "Nama Organisasi",
        description: "Ringkasan tanggung jawab dan hasil yang dicapai.",
    },
    {
        period: "2016 — 2019",
        role: "Jabatan / Peran Sebelumnya",
        organization: "Nama Organisasi",
        description: "Ringkasan tanggung jawab dan hasil yang dicapai.",
    },
    {
        period: "2012 — 2016",
        role: "Awal Karir / Pendidikan",
        organization: "Nama Institusi",
        description: "Titik awal perjalanan di dunia teknologi dan pendidikan.",
    },
];

export default function CareerSection({
    items = DEFAULT_CAREER,
}: {
    items?: CareerItem[];
}) {
    return (
        <section id="karir" className="bg-white py-[var(--section-space)]">
            <div className="mx-auto grid max-w-container gap-10 px-[var(--container-pad)] lg:grid-cols-12 lg:gap-16">
                <div className="lg:col-span-4">
                    <div className="lg:sticky lg:top-28">
                        <p className="mb-3 text-[length:var(--fs-eyebrow)] font-medium uppercase tracking-[var(--tracking-eyebrow)] text-muted">
                            Pengalaman
                        </p>
                        <h2 className="text-[length:var(--fs-h2)] font-semibold tracking-tight text-ink">
                            Perjalanan Karir
                        </h2>
                        <p className="mt-3 max-w-[36ch] text-sm leading-relaxed text-muted">
                            Langkah demi langkah membangun teknologi yang bermanfaat bagi
                            banyak orang.
                        </p>
                    </div>
                </div>

                <ol className="relative ml-2 border-l border-border lg:col-span-8">
                    {items.map((item) => (
                        <li
                            key={`${item.period}-${item.role}`}
                            className="relative pb-10 pl-8 last:pb-0"
                        >
                            <span
                                aria-hidden
                                className={`absolute -left-[9px] top-1 h-[17px] w-[17px] rounded-full border-4 border-white ring-1 ${item.current
                                        ? "bg-primary ring-primary"
                                        : "bg-muted/50 ring-border"
                                    }`}
                            />
                            <p
                                className={`text-xs font-semibold uppercase tracking-[var(--tracking-eyebrow)] ${item.current ? "text-primary" : "text-muted"
                                    }`}
                            >
                                {item.period}
                            </p>
                            <h3 className="mt-2 text-lg font-semibold tracking-tight text-ink">
                                {item.role}
                            </h3>
                            <p className="text-sm font-medium text-secondary">
                                {item.organization}
                            </p>
                            <p className="mt-2 max-w-[56ch] text-sm leading-relaxed text-muted">
                                {item.description}
                            </p>
                        </li>
                    ))}
                </ol>
            </div>
        </section>
    );
}