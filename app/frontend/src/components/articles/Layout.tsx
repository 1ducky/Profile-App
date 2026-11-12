import type { ReactNode } from "react";

interface ArticlesPageLayoutProps {
    header?: ReactNode;
    /** Search bar + filter kategori */
    toolbar?: ReactNode;
    /** Kartu unggulan (opsional) */
    featured?: ReactNode;
    /** Grid artikel / empty state / error */
    children: ReactNode;
}

export function ArticlesPageLayout({
    header,
    toolbar,
    featured,
    children,
}: ArticlesPageLayoutProps) {
    return (
        <main className="mx-auto min-h-screen max-w-content px-10 bg-bg pb-14 font-sans text-ink dark:bg-night-bg dark:text-night-text md:pb-20">
            {header}
            <section aria-label="Cari dan filter artikel">{toolbar}</section>
            <ArticleLayout>

                {featured && (
                    <section
                        aria-label="Artikel unggulan"
                        className="mt-5 px-5 lg:px-0"
                    >
                        {featured}
                    </section>
                )}
                <section aria-label="Daftar artikel" className="mt-10 px-5 lg:px-0 py-10">
                    {children}
                </section>
            </ArticleLayout>
        </main>
    );
}
function ArticleLayout({ children, Sidebar }: { children: ReactNode, Sidebar?: ReactNode }) {
    return (
        <div className="mx-auto max-w-8xl p-4 sm:p-6">
            {/* 
              Container utama:
              - Mobile (default): flex column-reverse (Sisi kiri/kecil otomatis di atas, sisi kanan di bawah)
              - Desktop (lg): Berubah menjadi CSS Grid 2 kolom (Sisi kiri di kiri, sisi kanan di kanan)
            */}
            <div className="flex flex-col-reverse lg:grid lg:grid-cols-[320px_1fr] lg:gap-8">

                {/* ========================================= */}
                <div>
                    {/* SISI KIRI (Sidebar / Kolom Kecil: Catatan, E-Paper, dll) */}
                    <div className="flex flex-col gap-6 mt-6 lg:mt-0">

                        {/* Contoh Widget 1: Catatan Dahlan Iskan */}
                        <div className="rounded-card border border-line bg-white shadow-card dark:border-night-line dark:bg-night-surface dark:shadow-none overflow-hidden">
                            <div className="bg-red-700 px-4 py-3 text-white font-bold text-sm tracking-wide">
                                CATATAN DAHLAN ISKAN
                            </div>
                            <ul className="divide-y divide-line dark:divide-night-line p-2">
                                <li className="p-2 text-sm hover:bg-gray-50 dark:hover:bg-night-muted/10 cursor-pointer">1. IKN Tiongkok</li>
                                <li className="p-2 text-sm hover:bg-gray-50 dark:hover:bg-night-muted/10 cursor-pointer">2. Gas Asap</li>
                                <li className="p-2 text-sm hover:bg-gray-50 dark:hover:bg-night-muted/10 cursor-pointer">3. Perang Hujan</li>
                            </ul>
                        </div>

                        {/* Contoh Widget 2: Disway E-Paper */}
                        <div className="rounded-card border border-line bg-white p-4 shadow-card dark:border-night-line dark:bg-night-surface dark:shadow-none">
                            <h3 className="font-bold text-base mb-3">DISWAY E-PAPER</h3>
                            <div className="aspect-[3/4] bg-gray-200 rounded-md overflow-hidden">
                                {/* Thumbnail E-paper */}
                                <img src="/assets/preview/mobile-prev-1.png" alt="E-Paper" className="size-full object-cover" />
                            </div>
                        </div>

                    </div>
                </div>

                {/* ========================================= */}
                <div>
                    {/* SISI KANAN (Konten Utama: Headline, Kolom Berita, dll) */}
                    <div className="flex flex-col gap-6">
                        {children}

                    </div>
                </div>

            </div>
        </div>
    );
}