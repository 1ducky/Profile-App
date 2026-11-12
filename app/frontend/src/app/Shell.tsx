"use client";

import { SidebarProvider } from "../components/dashboard/Sidebar/SidebarContext";




export function AppShell({ children }: { children: React.ReactNode }) {
    return (
        <SidebarProvider>
            {children}
        </SidebarProvider>
    );
}