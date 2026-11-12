import React, { useState } from "react";
import type { ArticleImage, ArticleStats } from "../../types/article/entity";
import { Eye, MessageCircle } from "lucide-react";

interface Props {
    image?: ArticleImage
    stats?: ArticleStats;
    priority?: boolean;
    className?: string;
}

export function ArticleThumbnail({ image, stats, priority = false, className = "" }: Props) {
    const [isError, setIsError] = useState(false);

    // Anggap gambar tidak ada jika data image kosong atau gagal dimuat
    const hasValidImage = image && !isError;

    return (
        <div
            className={`
                relative w-full aspect-[4/3] 
                overflow-hidden bg-gray-200 dark:bg-gray-800 
                ${className}
            `}
        >
            {hasValidImage ? (
                <img
                    src={image.src}
                    alt={image.alt}
                    loading={priority ? "eager" : "lazy"}
                    decoding="async"
                    draggable={false}
                    onError={() => setIsError(true)} // Fallback jika gambar rusak
                    className="absolute inset-0 block size-full object-cover"
                />
            ) : (
                // Placeholder jika gambar tidak ada / gagal load
                <div className="absolute inset-0 flex items-center justify-center text-gray-400 dark:text-gray-600">
                    {/* Ikon SVG sederhana sebagai placeholder (opsional) */}
                    <svg
                        className="w-8 h-8 opacity-50"
                        fill="none"
                        viewBox="0 0 24 24"
                        stroke="currentColor"
                    >
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
                    </svg>
                </div>
            )}
            {stats && (

                <div className="stat flex gap-2 absolute bottom-0 left-0 right-0 bg-surface p-2">
                    <span className="flex items-center gap-1 text-primary text-xs">
                        <Eye size={12} />
                        <span>{stats.views}</span>
                    </span>
                    <span className="flex items-center gap-1 text-primary text-xs">
                        <MessageCircle size={12} />
                        <span>{stats.comments}</span>
                    </span>
                </div>
            )}
        </div>
    );
}