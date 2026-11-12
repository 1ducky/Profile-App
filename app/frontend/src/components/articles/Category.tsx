import { ARTICLES_DEFAULTS, type ArticleCategoryFilter, type CategoryFilterProps } from "../../types/articles";


export function CategoryFilter({
    categories,
    active,
    onSelect,
    allLabel = ARTICLES_DEFAULTS.allCategoryLabel,
}: CategoryFilterProps) {
    const options: { value: ArticleCategoryFilter; label: string }[] = [
        { value: "semua", label: allLabel },
        ...categories.map((c) => ({ value: c.slug, label: c.label })),
    ];

    return (
        <div role="group" aria-label="Filter kategori" className="mt-4">
            <ul className="m-0 flex list-none snap-x gap-3 overflow-x-auto px-5 py-1 [scrollbar-width:none] [&::-webkit-scrollbar]:hidden lg:flex-wrap lg:overflow-visible lg:px-0">
                {options.map((opt) => {
                    const isActive = opt.value === active;
                    return (
                        <li key={opt.value} className="shrink-0 snap-start">
                            <button
                                type="button"
                                aria-pressed={isActive}
                                onClick={() => onSelect(opt.value)}
                                className={
                                    `min-h-11 rounded-pill border px-[22px] text-base transition-colors focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary motion-reduce:transition-none
                  ${isActive
                                        ? "border-ink bg-primary text-white dark:border-night-chip dark:bg-night-chip"
                                        : "border-line text-ink hover:border-muted dark:border-night-line dark:text-night-text dark:hover:border-night-muted"}`
                                }
                            >
                                {opt.label}
                            </button>
                        </li>
                    );
                })}
            </ul>
        </div>
    );
}