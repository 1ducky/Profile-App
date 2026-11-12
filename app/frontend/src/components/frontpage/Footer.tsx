import SocialLinks from "./const";

const NAV = [
    { label: "Beranda", href: "#beranda" },
    { label: "Tentang", href: "#tentang" },
    { label: "Produk", href: "#produk" },
    { label: "Dampak Sosial", href: "#dampak-sosial" },
    { label: "Artikel", href: "#artikel" },
    { label: "Kontak", href: "#kontak" },
];

export default function Footer({
    brandName = "Puguh Pamungkas",
    email = "halo@contoh.com",
}: {
    brandName?: string;
    email?: string;
}) {
    return (
        <footer className="bg-primary text-white">
            <div className="mx-auto max-w-container px-(--container-pad) py-14">
                <div className="grid gap-10 md:grid-cols-12">
                    <div className="md:col-span-5">
                        <a
                            href="#beranda"
                            className="inline-flex items-center gap-2 text-base font-semibold tracking-tight focus-visible:outline-2 focus-visible:outline-offset-4 focus-visible:outline-white"
                        >
                            <span
                                aria-hidden
                                className="grid h-7 w-7 place-items-center rounded-md bg-primary text-xs font-bold"
                            >
                                P
                            </span>
                            {brandName}
                        </a>
                        <p className="mt-4 max-w-[34ch] text-sm italic leading-relaxed text-white/75">
                            “Teknologi untuk manusia, bukan sebaliknya.”
                        </p>
                    </div>

                    <nav aria-label="Navigasi footer" className="md:col-span-4">
                        <h2 className="text-xs font-semibold uppercase tracking-eyebrow text-white/60">
                            Navigasi
                        </h2>
                        <ul className="mt-4 grid grid-cols-2 gap-x-6 gap-y-3">
                            {NAV.map((n) => (
                                <li key={n.href}>
                                    <a
                                        href={n.href}
                                        className="text-sm text-white/85 hover:text-white hover:underline focus-visible:outline-2 focus-visible:outline-offset-4 focus-visible:outline-white"
                                    >
                                        {n.label}
                                    </a>
                                </li>
                            ))}
                        </ul>
                    </nav>

                    <div className="md:col-span-3">
                        <h2 className="text-xs font-semibold uppercase tracking-eyebrow text-white/60">
                            Terhubung
                        </h2>
                        <a
                            href={`mailto:${email}`}
                            className="mt-4 inline-block text-sm text-white/85 hover:text-white hover:underline"
                        >
                            {email}
                        </a>
                        <div className="mt-5">
                            <SocialLinks tone="dark" />
                        </div>
                    </div>
                </div>

                <div className="mt-12 flex flex-col gap-3 border-t border-white/15 pt-6 text-xs text-white/60 sm:flex-row sm:items-center sm:justify-between">
                    <p>
                        © {new Date().getFullYear()} {brandName}. Hak cipta dilindungi.
                    </p>
                    <a
                        href="#beranda"
                        className="hover:text-white hover:underline focus-visible:outline-2 focus-visible:outline-offset-4 focus-visible:outline-white"
                    >
                        Kembali ke atas ↑
                    </a>
                </div>
            </div>
        </footer>
    );
}