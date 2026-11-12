
/** Tipe konten — menentukan komponen render yang tepat. */
export type ContentType =
    | "article"   // artikel teks biasa
    | "gallery"   // foto jurnalistik (readgallery)
    | "video"     // readvideo / podcast
    | "epaper";   // edisi e-paper

/** Badge overlay di atas gambar (hero & stream). */

// ============================================================
// 2. SUB-TIPE
// ============================================================

/** Penulis / kolumnis. */
export interface Author {
    /** Nama tampil, mis. "Dahlan Iskan & Redaksi", "Maya Saraswati". */
    name: string;
    /** Inisial 1–2 huruf untuk avatar fallback, mis. "DI", "MS". */
    initials?: string;
    /** URL profil penulis (opsional). */
    url?: string;
    /** True jika penulis adalah kolumnis tetap (mis. Dahlan Iskan, Azrul Ananda). */
    isColumnist?: boolean;
}

/** Metadata gambar. */
export interface ArticleImage {
    /** URL gambar (absolut atau path /public). */
    src: string;
    /** Alt text wajib untuk aksesibilitas & SEO. */
    alt: string;
    /** Caption foto, mis. "Foto: Biro Beijing Disway Media". */
    caption?: string;
    /** Rasio aspek yang diharapkan, mis. "16/9", "4/3", "1/1". */
    aspectRatio?: string;
}

/** Statistik engagement. */
export interface ArticleStats {
    views?: number;
    comments?: number;
    shares?: number;
}

// ============================================================
// 3. ARTICLE (KANONIK)
// ============================================================

export interface Article {
    // --- Identitas ---
    /** ID numerik dari URL /read/{id}/{slug}. Disimpan sebagai string agar fleksibel. */
    id: string;
    /** Slug URL, mis. "disway-business-adventures-beijing-tianjin-2". */
    slug: string;
    /** URL absolut halaman artikel. */
    url?: string;
    /** Tipe konten — default "article". */
    contentType: ContentType;

    // --- Konten ---
    title: string;
    /** Ringkasan 1–3 kalimat. Dipakai di stream card (line-clamp-2/3). */
    excerpt?: string;
    /** Isi lengkap (opsional — hanya jika halaman butuh render penuh). */
    body?: string;

    // --- Taksonomi ---
    /** Rubrik utama. */
    category: string;
    subCategory?: string;
    badges?: string[];
    /** Tag/hashtag, mis. ["#BallonDor2026", "#HarryKane"]. */
    tags?: string[];

    // --- Media ---
    image?: ArticleImage;

    // --- Penulis ---
    author: Author;

    // --- Waktu ---
    /** ISO 8601, mis. "2026-10-28T14:30:00+07:00". */
    publishedAt: string;
    /** ISO 8601, opsional jika pernah diupdate. */
    updatedAt?: string;

    // --- Engagement ---
    stats?: ArticleStats;

    related?: Article[];
    Content?: string; //Markdown
}