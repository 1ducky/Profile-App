import { Edit, Monitor, PanelLeftClose, PanelLeftOpen, Users, X } from "lucide-react";
import { useSidebar } from "./SidebarContext";

const NAV_ITEMS = [
    { key: "dashboard", label: "Dashboard", href: "/dashboard", icon: <Monitor size={20} /> },
    { key: "users", label: "Users", href: "/dashboard/users", icon: <Users size={20} /> },
    { key: "settings", label: "Editor", href: "/dashboard/editor", icon: <Edit size={20} /> },
];

function NavList({ compact, onNavigate }: { compact: boolean; onNavigate?: () => void }) {
    return (
        <nav aria-label="Menu utama" className="flex flex-col gap-1 px-3">
            {NAV_ITEMS.map(({ key, label, href, icon: Icon }) => {
                return (
                    <a

                        href={href}
                        key={key}
                        onClick={onNavigate}
                        title={compact ? label : undefined}
                        className={[
                            "group relative flex items-center gap-3 rounded-full py-3 text-sm font-semibold transition-colors text-white/80 hover:bg-white/10 hover:text-white",
                            "focus-visible:outline-none focus-visible:ring-[3px] focus-visible:ring-primary/30",
                            compact ? "justify-center px-3" : "px-4",
                        ].join(" ")}
                    >
                        {Icon}
                        {!compact && <span className="truncate font-sans">{label}</span>}
                    </a>
                );
            })}
        </nav>
    );
}

export default function Sidebar() {
    const { collapsed, mobileOpen, toggleCollapsed, closeMobile } = useSidebar();

    return (
        <>
            {/* Desktop */}
            <aside
                className={[
                    "fixed inset-y-0 left-0 z-30 hidden flex-col bg-primary transition-[width] duration-200 lg:flex mt-16",
                    collapsed ? "w-18" : "w-64",
                ].join(" ")}
            >

                <div className="flex-1 overflow-y-auto">
                    <NavList compact={collapsed} />
                </div>
                <div className="p-3">
                    <button
                        type="button"
                        onClick={toggleCollapsed}
                        aria-label={collapsed ? "Perluas sidebar" : "Ciutkan sidebar"}
                        className={[
                            "flex w-full items-center gap-3 rounded-full py-3 text-sm font-semibold text-white/75 transition-colors hover:bg-white/10 hover:text-white",
                            "focus-visible:outline-none focus-visible:ring-[3px] focus-visible:ring-primary/30",
                            collapsed ? "justify-center px-3" : "px-4",
                        ].join(" ")}
                    >
                        {collapsed ? <PanelLeftOpen size={20} /> : <PanelLeftClose size={20} />}
                        {!collapsed && <span className="font-sans">Ciutkan</span>}
                    </button>
                </div>
            </aside>

            {/* Mobile / tablet drawer */}
            <div className={["fixed inset-0 z-40 lg:hidden", mobileOpen ? "" : "pointer-events-none"].join(" ")}>
                <div
                    onClick={closeMobile}
                    aria-hidden
                    className={[
                        "absolute inset-0 bg-foreground/50 transition-opacity duration-200",
                        mobileOpen ? "opacity-100" : "opacity-0",
                    ].join(" ")}
                />
                <aside
                    role="dialog"
                    aria-modal="true"
                    aria-label="Menu navigasi"
                    className={[
                        "absolute inset-y-0 left-0 flex w-[280px] max-w-[85vw] flex-col bg-primary shadow-lift transition-transform duration-200",
                        mobileOpen ? "translate-x-0" : "-translate-x-full",
                    ].join(" ")}
                >
                    <div className="flex items-center justify-between pr-3">
                        <button
                            type="button"
                            onClick={closeMobile}
                            aria-label="Tutup menu"
                            className="grid h-10 w-10 place-items-center rounded-full text-white/75 hover:bg-white/10 hover:text-white focus-visible:outline-none focus-visible:ring-[3px] focus-visible:ring-primary/30"
                        >
                            <X size={20} />
                        </button>
                    </div>
                    <div className="flex-1 overflow-y-auto">
                        <NavList compact={false} onNavigate={closeMobile} />
                    </div>
                </aside>
            </div>
        </>
    );
}

