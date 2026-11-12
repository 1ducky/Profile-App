import { useState } from "react";
import { Sidebar } from "../../components/dashboard/Article/Sidebar";
import { IconButton } from "../../components/global/primitive/button";
import { Bell, Menu, Search } from "lucide-react";
import type { ProfileResponse } from "../../contract/profile";

export function ArticleManagementLayout({ children,Profile }: { children: React.ReactNode,Profile: ProfileResponse  }) {
    const [drawer, setDrawer] = useState(false);

    return (
        <div className="min-h-screen bg-[#F8FAFC] font-sans text-[#1F2937]">
            {/* Sidebar desktop */}
            <aside className="fixed inset-y-0 left-0 hidden w-64 border-r border-gray-200 lg:block">
                <Sidebar />
            </aside>

            {/* Sidebar mobile (drawer) */}
            {drawer && (
                <div className="fixed inset-0 z-40 lg:hidden">
                    <div className="absolute inset-0 bg-[#1F2937]/60" onClick={() => setDrawer(false)} aria-hidden="true" />
                    <aside className="relative h-full w-72 max-w-[85%] shadow-xl">
                        <Sidebar onNavigate={() => setDrawer(false)} />
                    </aside>
                </div>
            )}

            <div className="lg:pl-64">
                {/* Topbar sticky */}
                <header className="sticky top-0 z-30 flex h-16 items-center gap-3 border-b border-gray-200 bg-white/90 px-4 backdrop-blur sm:px-8">
                    <IconButton label="Buka menu" className="lg:hidden" onClick={() => setDrawer(true)}>
                        <Menu className="h-5 w-5" />
                    </IconButton>

                    <div className="relative hidden max-w-md flex-1 sm:block">
                        <Search className="pointer-events-none absolute left-4 top-1/2 h-4 w-4 -translate-y-1/2 text-[#6B7280]" aria-hidden="true" />
                        <input
                            type="search"
                            aria-label="Cari di seluruh panel"
                            placeholder="Cari konten, media, pesan"
                            className={`${"w-full rounded-xl border border-gray-300 bg-white px-4 py-2.5 text-sm text-[#1F2937] placeholder:text-[#6B7280]/70 focus-visible:border-[#0B63CE] focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-[#0B63CE]/30"} rounded-full bg-[#F8FAFC] pl-11`}
                        />
                    </div>

                    <div className="ml-auto flex items-center gap-1">
                        <IconButton label="Notifikasi">
                            <Bell className="h-5 w-5" />
                        </IconButton>
                        <span
            className="ml-2 flex h-9 w-9 shrink-0 items-center justify-center overflow-hidden rounded-full bg-[#1AA6B7] text-sm font-bold text-white"
            aria-label={`Akun: ${Profile.username}#${Profile.tag}`}
            title={`${Profile.username}#${Profile.tag}${Profile.bio ? ` — ${Profile.bio}` : ""}`}
        >
            {Profile.avatar != "" ? (
                <img
                    src={Profile.avatar}
                    alt={Profile.username}
                    className="h-full w-full object-cover"
                />
            ) : (
                Profile.username.charAt(0).toUpperCase()
            )}
        </span>
                    </div>
                </header>

                <main className="mx-auto w-full max-w-[1200px] px-4 py-8 sm:px-8 sm:py-10">
                    {children}
                </main>
            </div>
        </div>
    );
}