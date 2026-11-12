"use client";

import {
    createContext, useCallback, useContext, useEffect, useMemo, useState,
    type ReactNode,
} from "react";

type SidebarContextValue = {
    /** Desktop: sidebar diciutkan jadi mode ikon (72px) */
    collapsed: boolean;
    /** Mobile/tablet: drawer terbuka */
    mobileOpen: boolean;
    toggleCollapsed: () => void;
    openMobile: () => void;
    closeMobile: () => void;
};

const SidebarContext = createContext<SidebarContextValue | null>(null);

export function SidebarProvider({ children }: { children: ReactNode }) {
    const [collapsed, setCollapsed] = useState(true);
    const [mobileOpen, setMobileOpen] = useState(false);

    const toggleCollapsed = useCallback(() => setCollapsed((v) => !v), []);
    const openMobile = useCallback(() => setMobileOpen(false), []);
    const closeMobile = useCallback(() => setMobileOpen(false), []);

    // Tutup drawer dengan tombol Escape
    useEffect(() => {
        if (!mobileOpen) return;
        const onKey = (e: KeyboardEvent) => e.key === "Escape" && setMobileOpen(false);
        window.addEventListener("keydown", onKey);
        return () => window.removeEventListener("keydown", onKey);
    }, [mobileOpen]);

    const value = useMemo(
        () => ({ collapsed, mobileOpen, toggleCollapsed, openMobile, closeMobile }),
        [collapsed, mobileOpen, toggleCollapsed, openMobile, closeMobile]
    );

    return <SidebarContext.Provider value={value}>{children}</SidebarContext.Provider>;
}

export function useSidebar() {
    const ctx = useContext(SidebarContext);
    if (!ctx) throw new Error("useSidebar harus dipakai di dalam <SidebarProvider>");
    return ctx;
}