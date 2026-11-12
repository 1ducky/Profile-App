import { FileText, ImagePlus, Inbox, LayoutDashboard, Settings, UserRound } from "lucide-react";
import type { LucideIcon } from "lucide-react/types";

export interface NavItem {
    label: string
    icon: LucideIcon
    href: string
    active?: boolean
}

const NAV_ITEMS: NavItem[] = [
    { label: "Ringkasan", icon: LayoutDashboard, href: "/admin" },
    { label: "Konten", icon: FileText, href: "/admin/konten", active: true },
    { label: "Media", icon: ImagePlus, href: "/admin/media" },
    { label: "Halaman Tentang", icon: UserRound, href: "/admin/tentang" },
    { label: "Pesan Masuk", icon: Inbox, href: "/admin/pesan" },
    { label: "Pengaturan", icon: Settings, href: "/admin/pengaturan" },
];

export function Sidebar({ onNavigate, navItem = NAV_ITEMS }: { onNavigate?: () => void, navItem?: NavItem[] }) {
    return (
        <nav aria-label="Navigasi admin" className="flex h-full flex-col gap-6 bg-white px-4 py-6">
            <a href="/admin" className="flex items-center gap-3 px-2">
                <span className="flex h-10 w-10 items-center justify-center rounded-full bg-[#0B63CE] text-sm font-bold text-white">
                    PP
                </span>
                <span className="leading-tight">
                    <span className="block text-sm font-bold text-[#1F2937]">Puguh Pamungkas</span>
                    <span className="block text-xs text-[#6B7280]">Panel konten</span>
                </span>
            </a>

            <ul className="flex flex-1 flex-col gap-1">
                {NAV_ITEMS.map(({ label, icon: Icon, href, active, }) => (
                    <li key={label}>
                        <a
                            href={href}
                            onClick={onNavigate}
                            aria-current={active ? "page" : undefined}
                            className={`flex items-center gap-3 rounded-full px-4 py-2.5 text-sm font-medium transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-[#0B63CE] ${active
                                ? "bg-[#0B63CE]/10 font-semibold text-[#0B63CE]"
                                : "text-[#6B7280] hover:bg-gray-100 hover:text-[#1F2937]"
                                }`}
                        >
                            <Icon className="h-5 w-5 shrink-0" aria-hidden="true" />
                            <span className="flex-1">{label}</span>

                        </a>
                    </li>
                ))}
            </ul>

        </nav>
    );
}