/**
 * Type Contract — Halaman "Artikel & Wawasan"
 * Website Profil Puguh Pamungkas
 *
 * Struktur folder yang disarankan:
 *
 * src/features/articles/
 * ├── types/
 * │   └── articles.contract.ts        ← file ini
 * ├── api/
 * │   └── articles.service.ts         (fetch list, detail, bookmark)
 * ├── hooks/
 * │   ├── useArticles.ts              (query, filter, search, debounce)
 * │   └── useBookmarks.ts             (persist bookmark)
 * ├── components/
 * │   ├── ArticlesPageHeader.tsx      (breadcrumb, judul, subjudul)
 * │   ├── ArticleSearchBar.tsx
 * │   ├── CategoryFilter.tsx          (chip: Semua, Teknologi, ...)
 * │   ├── BookmarkButton.tsx
 * │   ├── FeaturedArticleCard.tsx     (kartu besar)
 * │   ├── ArticleCard.tsx             (kartu grid 2 kolom)
 * │   ├── ArticleGrid.tsx
 * │   └── ArticlesEmptyState.tsx
 * └── pages/
 *     └── ArticlesPage.tsx
 */

/* ============================================================
 * 1. PRIMITIVES
 * ============================================================ */

/** ISO 8601, mis. "2026-10-27T08:00:00Z" */
export type ISODateString = string;

/** URL-safe identifier, mis. "transformasi-digital" */
export type Slug = string;

/** Kategori sesuai chip filter; "semua" hanya dipakai di UI, bukan data artikel. */
export type ArticleCategorySlug =
    | "teknologi"
    | "pendidikan"
    | "sosial"
    | "perspektif";

export type ArticleCategoryFilter = "semua" | ArticleCategorySlug;

/* ============================================================
 * 2. DOMAIN MODELS
 * ============================================================ */

export interface ArticleCategory {
    slug: ArticleCategorySlug;
    label: string; // "Teknologi"
}

export interface ArticleImage {
    src: string;
    /** Wajib & deskriptif (WCAG, lihat design.md §8) */
    alt: string;
    width?: number;
    height?: number;
    /** LQIP / blur placeholder opsional */
    blurDataUrl?: string;
}

export interface Article {
    id: string;
    slug: Slug;
    title: string;
    /** Ringkasan singkat untuk kartu */
    excerpt: string;
    category: ArticleCategory;
    thumbnail: ArticleImage;
    publishedAt: ISODateString;
    /** Estimasi menit baca, mis. 5 → "5 MENIT BACA" */
    readingTimeMinutes: number;
    /** Tandai artikel unggulan (maksimal 1 tampil di hero list) */
    isFeatured: boolean;
    /** URL tujuan "Baca selengkapnya / Baca artikel" */
    href: string;
}

/** Status bookmark per pengguna (dipisah agar Article tetap murni) */
export interface ArticleBookmarkState {
    articleId: string;
    isBookmarked: boolean;
}

/* ============================================================
 * 3. API CONTRACT
 * ============================================================ */

export interface ListArticlesQuery {
    /** Pencarian teks (judul + ringkasan) */
    search?: string;
    category?: ArticleCategoryFilter;
    page?: number; // default 1
    pageSize?: number; // default 9
    sort?: "newest" | "oldest";
}

export interface PaginationMeta {
    page: number;
    pageSize: number;
    totalItems: number;
    totalPages: number;
    hasNextPage: boolean;
}

export interface ListArticlesResponse {
    /** Artikel unggulan; null bila tidak ada / sedang filter aktif */
    featured: Article | null;
    items: Article[];
    categories: ArticleCategory[];
    meta: PaginationMeta;
}

export interface ToggleBookmarkRequest {
    articleId: string;
}

export interface ToggleBookmarkResponse extends ArticleBookmarkState { }

export interface ApiError {
    code: "NETWORK" | "NOT_FOUND" | "VALIDATION" | "UNKNOWN";
    message: string;
}

/** Service interface — implementasi bebas (REST, CMS, mock). */
export interface ArticlesService {
    list(query: ListArticlesQuery): Promise<ListArticlesResponse>;
    toggleBookmark(req: ToggleBookmarkRequest): Promise<ToggleBookmarkResponse>;
    getBookmarkedIds(): Promise<string[]>;
}

/* ============================================================
 * 4. UI STATE
 * ============================================================ */

export type AsyncStatus = "idle" | "loading" | "success" | "error" | "empty";

export interface ArticlesViewState {
    status: AsyncStatus;
    query: Required<Pick<ListArticlesQuery, "search" | "category" | "page">>;
    data: ListArticlesResponse | null;
    bookmarkedIds: ReadonlySet<string>;
    error: ApiError | null;
}

/* ============================================================
 * 5. COMPONENT PROPS
 * ============================================================ */

export interface BreadcrumbItem {
    label: string;
    href?: string; // item terakhir tanpa href = halaman aktif
}

export interface ArticlesPageHeaderProps {
    breadcrumb: BreadcrumbItem[];
    title: string; // "Artikel & Wawasan"
    subtitle: string; // "Gagasan, pengalaman, dan perkembangan terbaru."
}

export interface ArticleSearchBarProps {
    value: string;
    placeholder?: string; // default "Cari artikel..."
    onChange: (value: string) => void;
    /** Debounce dalam ms; default 300 */
    debounceMs?: number;
    onClear?: () => void;
}

export interface CategoryFilterProps {
    categories: ArticleCategory[];
    active: ArticleCategoryFilter;
    onSelect: (category: ArticleCategoryFilter) => void;
    /** Label chip "Semua"; default "Semua" */
    allLabel?: string;
}

export interface BookmarkButtonProps {
    articleId: string;
    isBookmarked: boolean;
    onToggle: (articleId: string) => void;
    /** Untuk screen reader: "Simpan artikel" / "Hapus dari simpanan" */
    labelSaved?: string;
    labelUnsaved?: string;
    disabled?: boolean;
}

interface ArticleCardBaseProps {
    article: Article;
    isBookmarked: boolean;
    onToggleBookmark: (articleId: string) => void;
    onOpen?: (article: Article) => void;
}

export interface FeaturedArticleCardProps {
    articles: Article[];
    isBookmarked: boolean;
    onToggleBookmark: (articleId: string) => void;
    onOpen?: (article: Article) => void;
    /** Teks tautan; default "Baca selengkapnya" */
    ctaLabel?: string;
}

export interface ArticleCardProps extends ArticleCardBaseProps {
    /** Teks tautan; default "Baca artikel" */
    ctaLabel?: string;
}

export interface ArticleGridProps {
    articles: Article[];
    bookmarkedIds: ReadonlySet<string>;
    onToggleBookmark: (articleId: string) => void;
    /** Mobile 2 kolom (sesuai desain); desktop 3 kolom */
    columns?: { mobile: 1 | 2; tablet: 2 | 3; desktop: 3 | 4 };
}

export interface ArticlesEmptyStateProps {
    search: string;
    category: ArticleCategoryFilter;
    onReset: () => void;
}

export interface ArticlesPageProps {
    initialData?: ListArticlesResponse;
    initialQuery?: Partial<ListArticlesQuery>;
    service?: ArticlesService; // injectable untuk testing
}

/* ============================================================
 * 6. UI CONSTANTS (single source of truth)
 * ============================================================ */

export const ARTICLE_CATEGORIES: readonly ArticleCategory[] = [
    { slug: "teknologi", label: "Teknologi" },
    { slug: "pendidikan", label: "Pendidikan" },
    { slug: "sosial", label: "Sosial" },
    { slug: "perspektif", label: "Perspektif" },
] as const;

export const ARTICLES_DEFAULTS = {
    pageSize: 9,
    searchDebounceMs: 300,
    allCategoryLabel: "Semua",
    searchPlaceholder: "Cari artikel...",
} as const;