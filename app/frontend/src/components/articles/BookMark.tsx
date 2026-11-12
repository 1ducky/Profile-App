import type { BookmarkButtonProps } from "../../types/articles";


export function BookmarkButton({
    articleId,
    isBookmarked,
    onToggle,
    labelSaved = "Hapus dari simpanan",
    labelUnsaved = "Simpan artikel",
    disabled,
}: BookmarkButtonProps) {
    return (
        <button
            type="button"
            aria-pressed={isBookmarked}
            aria-label={isBookmarked ? labelSaved : labelUnsaved}
            disabled={disabled}
            onClick={() => onToggle(articleId)}
            className={
                `grid size-10 place-items-center rounded-pill text-white backdrop-blur-sm transition-colors focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-white focus-visible:ring-4 focus-visible:ring-primary disabled:opacity-50 motion-reduce:transition-none md:size-11"
                ${isBookmarked ? "bg-primary" : "bg-black/55 hover:bg-black/70"}`
            }
        >
            <svg
                viewBox="0 0 24 24"
                width="22"
                height="22"
                aria-hidden="true"
                fill={isBookmarked ? "currentColor" : "none"}
                stroke="currentColor"
                strokeWidth="2"
                strokeLinejoin="round"
            >
                <path d="M6 3h12a1 1 0 0 1 1 1v17l-7-4.5L5 21V4a1 1 0 0 1 1-1Z" />
            </svg>
        </button>
    );
}