import type { ReactNode } from "react";

export type SocialLink = { label: string; href: string; icon: ReactNode };

const p = {
    width: 18,
    height: 18,
    viewBox: "0 0 24 24",
    fill: "none",
    stroke: "currentColor",
    strokeWidth: 1.8,
    strokeLinecap: "round" as const,
    strokeLinejoin: "round" as const,
    "aria-hidden": true,
};

/** Ganti href dengan akun asli. */
export const DEFAULT_SOCIALS: SocialLink[] = [
    {
        label: "LinkedIn",
        href: "https://linkedin.com/in/username",
        icon: (
            <svg {...p}>
                <path d="M16 8a6 6 0 0 1 6 6v7h-4v-7a2 2 0 0 0-2-2 2 2 0 0 0-2 2v7h-4v-7a6 6 0 0 1 6-6z" />
                <rect x="2" y="9" width="4" height="12" />
                <circle cx="4" cy="4" r="2" />
            </svg>
        ),
    },
    {
        label: "Instagram",
        href: "https://instagram.com/username",
        icon: (
            <svg {...p}>
                <rect x="2" y="2" width="20" height="20" rx="5" />
                <path d="M16 11.37A4 4 0 1 1 12.63 8 4 4 0 0 1 16 11.37z" />
                <path d="M17.5 6.5h.01" />
            </svg>
        ),
    },
    {
        label: "GitHub",
        href: "https://github.com/username",
        icon: (
            <svg {...p}>
                <path d="M15 22v-4a4.8 4.8 0 0 0-1-3.5c3 0 6-2 6-5.5.08-1.25-.27-2.48-1-3.5.28-1.15.28-2.35 0-3.5 0 0-1 0-3 1.5-2.64-.5-5.36-.5-8 0C6 2 5 2 5 2c-.3 1.15-.3 2.35 0 3.5A5.403 5.403 0 0 0 4 9c0 3.5 3 5.5 6 5.5-.39.49-.68 1.05-.85 1.65-.17.6-.22 1.23-.15 1.85v4" />
                <path d="M9 18c-4.51 2-5-2-7-2" />
            </svg>
        ),
    },
    {
        label: "YouTube",
        href: "https://youtube.com/@username",
        icon: (
            <svg {...p}>
                <path d="M2.5 17a24.12 24.12 0 0 1 0-10 2 2 0 0 1 1.4-1.4 49.56 49.56 0 0 1 16.2 0A2 2 0 0 1 21.5 7a24.12 24.12 0 0 1 0 10 2 2 0 0 1-1.4 1.4 49.55 49.55 0 0 1-16.2 0A2 2 0 0 1 2.5 17" />
                <path d="m10 15 5-3-5-3z" />
            </svg>
        ),
    },
];

export default function SocialLinks({
    links = DEFAULT_SOCIALS,
    tone = "light",
}: {
    links?: SocialLink[];
    /** light: untuk latar terang. dark: untuk latar gelap (footer). */
    tone?: "light" | "dark";
}) {
    const style =
        tone === "dark"
            ? "border-white/25 text-white hover:border-white hover:bg-white/10 focus-visible:outline-white"
            : "border-border bg-white text-ink hover:border-primary hover:text-primary focus-visible:outline-primary";

    return (
        <ul className="flex flex-wrap items-center gap-3">
            {links.map((l) => (
                <li key={l.label}>
                    <a
                        href={l.href}
                        target="_blank"
                        rel="noopener noreferrer"
                        aria-label={l.label}
                        className={`grid h-10 w-10 place-items-center rounded-full border transition-colors focus-visible:outline-2 focus-visible:outline-offset-2 ${style}`}
                    >
                        {l.icon}
                    </a>
                </li>
            ))}
        </ul>
    );
}