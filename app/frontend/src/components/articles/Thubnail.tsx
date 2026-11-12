import type { ArticleImage } from "../../types/articles";


interface Props {
    image?: ArticleImage;
    /** Featured dimuat langsung, grid lazy (design.md §8) */
    priority?: boolean;
    className?: string;
}

export function ArticleThumbnail({ image, priority = false, className }: Props) {
    return (
        <div
            className={
                `aspect-[4/3] overflow-hidden bg-line dark:bg-night-line ${className}`
            }
        >
            {image && (

                <img
                    src={image.src}
                    alt={image.alt}
                    width={image.width}
                    height={image.height}
                    loading={priority ? "eager" : "lazy"}
                    decoding="async"
                    draggable={false}
                    className="block size-full object-cover"
                />
            )}
        </div>
    );
}