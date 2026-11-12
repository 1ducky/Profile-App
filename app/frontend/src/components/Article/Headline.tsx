import { useRef } from "react";
import type { Article } from "../../types/article/entity";
import { useDragSlider } from "../../hooks/useDrag";
import { ArticleThumbnail } from "./Thubnail";

export default function HeadlineSection({ article }: { article: Article[] }) {
    const drag = useRef({ startX: 0, startT: 0, dx: 0, active: false, moved: false });
    const count = article.length;
    const { viewportRef, index, go, dragging, onKeyDown, trackProps } =
        useDragSlider(count);
    return (
        <section className="w-full max-w-[1200px] mx-auto bg-white pb-10">

            {/* 1. Label Headline */}


            {/* 2. Hero Article (Gambar Besar) */}
            <div className="relative w-full aspect-3/2 md:aspect-2/1 overflow-hidden group cursor-pointer mb-4">
                <div className="absolute top-0 left-0 flex z-2">
                    <div className="bg-[#1a1a1a] text-white px-5 py-1.5 text-[18px] font-bold tracking-wide uppercase">
                        Headline
                    </div>
                </div>
                <div
                    ref={viewportRef}
                    className="-mx-3 overflow-hidden h-full"
                    tabIndex={0}
                    onKeyDown={onKeyDown}
                >
                    {/* Track: children memanjang ke samping */}
                    <div
                        className={`flex cursor-grab touch-pan-y select-none active:cursor-grabbing h-full ${dragging
                            ? "transition-none"
                            : "transition-transform duration-500 ease-soft motion-reduce:transition-none"
                            }`}
                        style={trackProps.style}
                        onPointerDown={trackProps.onPointerDown}
                        onPointerMove={trackProps.onPointerMove}
                        onPointerUp={trackProps.onPointerUp}
                        onPointerCancel={trackProps.onPointerCancel}
                        onClickCapture={trackProps.onClickCapture}
                    >
                        {article.map((article, i) => (
                            <article
                                key={article.id}
                                role="group"
                                aria-roledescription="slide"
                                aria-label={`${i + 1} dari ${count}`}
                                {...(i !== index ? ({ inert: "" } as object) : {})}
                                // 1. Tambahkan relative & h-full. Hapus p-3 jika ingin gambar mentok ujung.
                                // Jika ingin ada jarak antar slide, bungkus konten di dalam div dengan padding, 
                                // tapi biarkan article-nya full.
                                className="relative basis-full shrink-0 grow-0 h-full overflow-hidden"
                            >
                                {/* 2. Bungkus gambar dan gradient dalam satu container relative/absolute */}


                                {/* Gradient Overlay */}
                                <div className="absolute inset-0 z-10 bg-gradient-to-t from-black/90 via-black/40 to-transparent pointer-events-none"></div>

                                {/* Thumbnail Gambar */}
                                <ArticleThumbnail
                                    image={article.image}
                                    priority={true}
                                    stats={article.stats}
                                    // 3. Pastikan gambar memenuhi container (absolute inset-0 atau w-full h-full object-cover)
                                    className="absolute inset-0 w-full h-full object-cover z-0"
                                />

                                {/* Konten Teks (Judul & Tanggal) */}
                                <div className="absolute bottom-0 left-0 p-6 md:p-10 w-full flex flex-col justify-end text-white z-20">
                                    <h2 className="text-2xl md:text-4xl lg:text-5xl font-extrabold leading-tight tracking-tight drop-shadow-lg mb-3">
                                        {article.title}
                                    </h2>
                                    <span className="text-sm font-medium text-gray-300">
                                        {article.publishedAt}
                                    </span>
                                </div>
                            </article>
                        ))}
                    </div>
                </div>

            </div>

            {/* 3. Stream Articles (Grid 5 Kolom di Desktop) */}
            <div className="lg:grid hidden grid-cols-2 md:grid-cols-3 lg:grid-cols-5 gap-4">
                {article.map((article) => (
                    <a
                        key={article.id}
                        className="flex flex-col group cursor-pointer"
                        href={article.url ? article.url : `/articles/${article.slug}`}
                    >
                        {/* Thumbnail Gambar */}
                        <ArticleThumbnail image={article.image}
                            priority={false}
                            className="w-full h-full object-cover "
                        />

                        {/* Judul */}
                        <h3 className="text-sm font-bold leading-snug text-gray-900 line-clamp-4 group-hover:text-blue-600 transition-colors">
                            {article.title}
                        </h3>
                    </a>
                ))}
            </div>

        </section>
    );
}