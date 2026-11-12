import type { ArticlesPageHeaderProps } from "../../types/articles";


export function ArticlesPageHeader({
    breadcrumb,
    title,
    subtitle,
}: ArticlesPageHeaderProps) {
    return (
        <header className="px-5 pt-6 lg:px-0">
            <nav aria-label="Breadcrumb">
                <ol className="flex flex-wrap gap-2 text-sm text-muted dark:text-night-muted">
                    {breadcrumb.map((item, i) => {
                        const isLast = i === breadcrumb.length - 1;
                        return (
                            <li
                                key={`${item.label}-${i}`}
                                className="flex gap-2 [&:not(:first-child)]:before:content-['/'] [&:not(:first-child)]:before:opacity-60"
                            >
                                {item.href && !isLast ? (
                                    <a
                                        href={item.href}
                                        className="hover:text-primary hover:underline"
                                    >
                                        {item.label}
                                    </a>
                                ) : (
                                    <span aria-current={isLast ? "page" : undefined}>
                                        {item.label}
                                    </span>
                                )}
                            </li>
                        );
                    })}
                </ol>
            </nav>
            <h1 className="mb-2 mt-3 text-[1.75rem] font-bold leading-[1.15] tracking-tight sm:text-4xl lg:text-[2.5rem]">
                {title}
            </h1>
            <p className="max-w-[60ch] text-[1.0625rem] leading-relaxed text-muted dark:text-night-muted">
                {subtitle}
            </p>
        </header>
    );
}