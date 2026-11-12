import { useEffect, useState } from "react";

type NavItem = { label: string; href: string };

const NAV_ITEMS: NavItem[] = [
    { label: "Beranda", href: "#beranda" },
    { label: "Tentang", href: "#tentang" },
    { label: "Produk", href: "#produk" },
    { label: "Dampak Sosial", href: "#dampak-sosial" },
    { label: "Artikel", href: "#artikel" },
    { label: "Kontak", href: "#kontak" },
];

type NavbarProps = {
    brandName?: string;
    logoSrc?: string;
    ctaLabel?: string;
    ctaHref?: string;
};

export default function Navbar({
    brandName = "Cak Oyim",
    logoSrc,
    ctaLabel = "Hubungi Saya",
    ctaHref = "#kontak",
}: NavbarProps) {
    const [open, setOpen] = useState(false);
    const [active, setActive] = useState("#beranda");
    const [scrolled, setScrolled] = useState(false);

    // Bayangan halus setelah scroll
    useEffect(() => {
        const onScroll = () => setScrolled(window.scrollY > 8);
        onScroll();
        window.addEventListener("scroll", onScroll, { passive: true });
        return () => window.removeEventListener("scroll", onScroll);
    }, []);

    // Tandai menu aktif sesuai section yang terlihat
    useEffect(() => {
        const sections = NAV_ITEMS.map((i) =>
            document.querySelector(i.href)
        ).filter(Boolean) as Element[];
        if (!sections.length) return;

        const observer = new IntersectionObserver(
            (entries) => {
                entries.forEach((e) => {
                    if (e.isIntersecting) setActive(`#${e.target.id}`);
                });
            },
            { rootMargin: "-40% 0px -55% 0px" }
        );
        sections.forEach((s) => observer.observe(s));
        return () => observer.disconnect();
    }, []);

    // Tutup menu mobile dengan Escape
    useEffect(() => {
        if (!open) return;
        const onKey = (e: KeyboardEvent) => e.key === "Escape" && setOpen(false);
        window.addEventListener("keydown", onKey);
        return () => window.removeEventListener("keydown", onKey);
    }, [open]);

    return (
        <header
            className={`sticky top-0 z-50 border-b bg-white/90 backdrop-blur transition-shadow ${scrolled
                ? "border-border shadow-card"
                : "border-transparent"
                }`}
        >
            <nav
                aria-label="Navigasi utama"
                className="mx-auto flex h-(--navbar-height) max-w-container items-center justify-between px-(--container-pad)"
            >
                {/* Logo */}
                <a
                    href="#beranda"
                    className="flex items-center gap-2 rounded-md text-[15px] font-semibold tracking-tight text-ink  focus-visible:outline-2 focus-visible:outline-offset-4 focus-visible:outline-primary"
                >
                    {logoSrc ? (
                        <img src={logoSrc} alt="" className="h-6 w-6" />
                    ) : (
                        <span
                            aria-hidden
                            className="grid h-6 w-6 place-items-center rounded-md bg-primary text-xs font-bold text-white"
                        >
                            P
                        </span>
                    )}
                    {brandName}
                </a>

                {/* Menu desktop */}
                <ul className="hidden items-center gap-8 lg:flex">
                    {NAV_ITEMS.map((item) => {
                        const isActive = active === item.href;
                        return (
                            <li key={item.href}>
                                <a
                                    href={item.href}
                                    aria-current={isActive ? "page" : undefined}
                                    className={`rounded text-sm transition-colors  focus-visible:outline-2 focus-visible:outline-offset-4 focus-visible:outline-primary ${isActive
                                        ? "font-semibold text-primary"
                                        : "text-ink hover:text-primary"
                                        }`}
                                >
                                    {item.label}
                                </a>
                            </li>
                        );
                    })}
                </ul>

                {/* CTA desktop */}
                <a
                    href={ctaHref}
                    className="hidden rounded-full bg-navy px-5 py-2.5 text-sm font-semibold text-white transition-colors hover:bg-primarye focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary lg:inline-block"
                >
                    {ctaLabel}
                </a>

                {/* Hamburger */}
                <button
                    type="button"
                    onClick={() => setOpen((v) => !v)}
                    aria-expanded={open}
                    aria-controls="menu-mobile"
                    aria-label={open ? "Tutup menu" : "Buka menu"}
                    className="grid h-10 w-10 place-items-center rounded-md text-ink focus-visible:outline-2 focus-visible:outline-primary lg:hidden"
                >
                    <svg
                        width="22"
                        height="22"
                        viewBox="0 0 24 24"
                        fill="none"
                        stroke="currentColor"
                        strokeWidth="2"
                        strokeLinecap="round"
                        aria-hidden
                    >
                        {open ? (
                            <path d="M6 6l12 12M18 6L6 18" />
                        ) : (
                            <path d="M4 7h16M4 12h16M4 17h16" />
                        )}
                    </svg>
                </button>
            </nav>

            {/* Menu mobile */}
            <div
                id="menu-mobile"
                hidden={!open}
                className="border-t border-border bg-white lg:hidden"
            >
                <ul className="mx-auto max-w-container px-(--container-pad) py-3">
                    {NAV_ITEMS.map((item) => (
                        <li key={item.href}>
                            <a
                                href={item.href}
                                onClick={() => setOpen(false)}
                                aria-current={active === item.href ? "page" : undefined}
                                className={`block rounded-md px-2 py-3 text-base ${active === item.href
                                    ? "font-semibold text-primary"
                                    : "text-ink"
                                    }`}
                            >
                                {item.label}
                            </a>
                        </li>
                    ))}
                    <li className="pt-2 pb-3">
                        <a
                            href={ctaHref}
                            onClick={() => setOpen(false)}
                            className="block w-full rounded-full bg-navy px-(--container-pad) py-3 text-center text-sm font-semibold text-white"
                        >
                            {ctaLabel}
                        </a>
                    </li>
                </ul>
            </div>
        </header>
    );
}