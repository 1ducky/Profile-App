import { useRef } from "react";
import type { Article, ArticleStats } from "../../types/article/entity";
import { useDragSlider } from "../../hooks/useDrag";
import { ArticleThumbnail } from "./Thubnail";

export default function GalerySection({ galery }: { galery: Article[] }) {
    const drag = useRef({ startX: 0, startT: 0, dx: 0, active: false, moved: false });
    const count = galery.length;
    const { viewportRef, index, go, dragging, onKeyDown, trackProps } =
        useDragSlider(count);
    return (
        <div className="">
            {count > 1 && (
                <div className="my-3 flex items-center justify-between">
                    <div className="flex items-center gap-2">
                        {galery.map((s, i) => (
                            <button
                                key={s.id}
                                type="button"
                                onClick={() => go(i)}
                                aria-label={`Ke slide ${i + 1}`}
                                aria-current={i === index}
                                className={`h-2 rounded-full transition-all focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary ${i === index ? "w-6 bg-primary" : "w-2 bg-border hover:bg-muted"
                                    }`}
                            />
                        ))}
                    </div>
                </div>
            )}
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
                    {galery.map((item, i) => (
                        <div
                            key={item.id}
                            role="group"
                            aria-roledescription="slide"
                            aria-label={`${i + 1} dari ${count}`}
                            {...(i !== index ? ({ inert: "" } as object) : {})}
                            className="relative basis-full shrink-0 grow-0 ">
                            <ArticleThumbnail
                                image={item.image}
                                priority={true}
                                stats={item.stats}
                                // 3. Pastikan gambar memenuhi container (absolute inset-0 atau w-full h-full object-cover)
                                className="absolute inset-0 w-full h-full object-cover z-0"
                            />
                        </div>
                    ))}
                </div>

            </div>
        </div>
    )
}