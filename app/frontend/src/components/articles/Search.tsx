import { useId } from "react";
import { ARTICLES_DEFAULTS, type ArticleSearchBarProps } from "../../types/articles";

/**
 * Controlled & presentational. Debounce ditangani di hook (useArticles).
 */
export function ArticleSearchBar({
    value,
    placeholder = ARTICLES_DEFAULTS.searchPlaceholder,
    onChange,
    onClear,
}: ArticleSearchBarProps) {
    const inputId = useId();

    return (
        <form
            role="search"
            onSubmit={(e) => e.preventDefault()}
            className="mx-5 mt-5 flex min-h-14 items-center gap-3 rounded-card border border-line bg-white px-4 text-muted transition focus-within:border-primary focus-within:ring-2 focus-within:ring-primary/25 motion-reduce:transition-none dark:border-night-line dark:bg-night-surface dark:text-night-muted lg:mx-0"
        >
            <label htmlFor={inputId} className="sr-only">
                Cari artikel
            </label>
            <svg
                viewBox="0 0 24 24"
                width="22"
                height="22"
                aria-hidden="true"
                fill="none"
                stroke="currentColor"
                strokeWidth="2"
                strokeLinecap="round"
                className="shrink-0"
            >
                <circle cx="11" cy="11" r="7" />
                <path d="m20 20-3.5-3.5" />
            </svg>
            <input
                id={inputId}
                type="search"
                value={value}
                placeholder={placeholder}
                autoComplete="off"
                onChange={(e) => onChange(e.target.value)}
                className="min-w-0 flex-1 appearance-none border-0 bg-transparent text-[1.0625rem] text-ink outline-none placeholder:text-muted dark:text-night-text dark:placeholder:text-night-muted [&::-webkit-search-cancel-button]:hidden"
            />
            {value && onClear && (
                <button
                    type="button"
                    onClick={onClear}
                    aria-label="Hapus pencarian"
                    className="grid size-8 place-items-center rounded-pill text-muted hover:text-ink focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary dark:text-night-muted dark:hover:text-night-text"
                >
                    <svg
                        viewBox="0 0 24 24"
                        width="18"
                        height="18"
                        aria-hidden="true"
                        fill="none"
                        stroke="currentColor"
                        strokeWidth="2"
                        strokeLinecap="round"
                    >
                        <path d="M6 6l12 12M18 6 6 18" />
                    </svg>
                </button>
            )}
        </form>
    );
}