// ============================================================
// D. MOCK: ARTICLE — per contentType
// ============================================================

import type { Article } from "../../types/article/entity";
import { mockAuthors } from "./Author";
import { mockImages } from "./Image";
import { mockStats } from "./Stats";

/** 1) ARTICLE — teks biasa (feature/berita) */
export const mockArticleNews: Article = {
    id: "1001",
    slug: "disway-business-adventures-beijing-tianjin-2",
    url: "https://disway.id/read/1001/disway-business-adventures-beijing-tianjin-2",
    contentType: "article",
    title: "Disway Business Adventures: Beijing–Tianjin #2",
    excerpt:
        "Perjalanan bisnis Disway Media menyusuri pelabuhan Tianjin, melihat langsung bagaimana China membangun konektivitas logistik kelas dunia.",
    body: "<p>...</p>",
    category: "Business",
    subCategory: "Adventure",
    badges: ["Eksklusif", "Laporan Khusus"],
    tags: ["#DiswayBusinessAdventures", "#China", "#Tianjin"],
    image: {
        src: "/assets/demo/sea.png",
        alt: "Harga Beras Naik 5 Persen di Pasar Tradisional",
        caption: "Harga Beras Naik 5 Persen di Pasar Tradisional",
        aspectRatio: "16/9",
    },
    author: mockAuthors.dahlanIskan,
    publishedAt: "2026-10-28T14:30:00+07:00",
    updatedAt: "2026-10-28T16:10:00+07:00",
    stats: mockStats.viral,
    related: [],
};

/** 2) GALLERY — foto jurnalistik */
export const mockArticleGallery: Article = {
    id: "1002",
    slug: "potret-pelabuhan-tianjin-foto-jurnalistik",
    contentType: "gallery",
    title: "Potret Pelabuhan Tianjin: Foto Jurnalistik",
    excerpt: "24 foto pilihan dari kunjungan Disway Media ke salah satu pelabuhan tersibuk di dunia.",
    category: "Foto",
    subCategory: "Jurnalistik",
    tags: ["#Tianjin", "#Pelabuhan", "#PhotoEssay"],
    image: mockImages.galleryEpaper,
    author: mockAuthors.biroBeijing,
    publishedAt: "2026-10-28T09:00:00+07:00",
    stats: mockStats.medium,
};

/** 3) VIDEO — readvideo / podcast */
export const mockArticleVideo: Article = {
    id: "1003",
    slug: "podcast-disway-ep-45-masa-depan-ai",
    contentType: "video",
    title: "Podcast Disway Ep. 45: Masa Depan AI & Media",
    excerpt: "Diskusi santai bareng praktisi AI tentang bagaimana redaksi berubah.",
    category: "Podcast",
    badges: ["Video"],
    tags: ["#PodcastDisway", "#AI"],
    image: mockImages.podcastStudio,
    author: mockAuthors.redaksi,
    publishedAt: "2026-10-27T20:00:00+07:00",
    stats: mockStats.medium,
    related: [],
};

/** 4) EPAPER — edisi e-paper */
export const mockArticleEpaper: Article = {
    id: "1004",
    slug: "epaper-disway-28-oktober-2026",
    contentType: "epaper",
    title: "E-Paper Disway — 28 Oktober 2026",
    excerpt: "Baca edisi cetak hari ini dalam format digital.",
    category: "E-Paper",
    image: mockImages.galleryEpaper,
    author: mockAuthors.redaksi,
    publishedAt: "2026-10-28T05:00:00+07:00",
    stats: mockStats.fresh,
};

/** 5) ARTICLE — kolumnis tetap (Dahlan Iskan) */
export const mockArticleColumn: Article = {
    id: "1005",
    slug: "catatan-dahlan-iskandar-melihat-china-dari-dekat",
    contentType: "article",
    title: "Melihat China dari Dekat",
    excerpt: "Kolom Dahlan Iskan: ada yang berubah cepat, ada yang tetap sama.",
    category: "Kolom",
    subCategory: "Dahlan Iskan",
    badges: ["Kolom"],
    tags: ["#DahlanIskan", "#China"],
    image: mockImages.beijingHero,
    author: mockAuthors.dahlanIskan,
    publishedAt: "2026-10-28T06:00:00+07:00",
    stats: mockStats.medium,
};

/** 6) ARTICLE — olahraga, headline viral */
export const mockArticleSports: Article = {
    id: "1006",
    slug: "harry-kane-ballon-dor-2026",
    contentType: "article",
    title: "Harry Kane Akhirnya Raih Ballon d'Or 2026",
    excerpt: "Setelah bertahun-tahun dianggap 'selalu hampir', Kane menutup musim terbaiknya dengan trofi paling bergengsi.",
    category: "Olahraga",
    subCategory: "Sepak Bola",
    badges: ["Breaking", "Headline"],
    tags: ["#BallonDor2026", "#HarryKane"],
    image: mockImages.ballonDor,
    author: mockAuthors.mayaSaraswati,
    publishedAt: "2026-10-28T03:15:00+07:00",
    stats: { views: 987_600, comments: 2_341, shares: 8_920 },
    related: [],
};

export const mockArticleRelate: Article = {
    id: "1006",
    slug: "relating",
    contentType: "article",
    title: "Harry Kane Akhirnya Raih Ballon d'Or 2026",
    excerpt: "Setelah bertahun-tahun dianggap 'selalu hampir', Kane menutup musim terbaiknya dengan trofi paling bergengsi.",
    category: "Olahraga",
    subCategory: "Sepak Bola",
    badges: ["Breaking", "Headline"],
    tags: ["#BallonDor2026", "#HarryKane"],
    image: mockImages.ballonDor,
    author: mockAuthors.mayaSaraswati,
    publishedAt: "2026-10-28T03:15:00+07:00",
    stats: { views: 987_600, comments: 2_341, shares: 8_920 },
    related: [mockArticleNews, mockArticleVideo],
};

/** 7) ARTICLE — minimal (tanpa image, excerpt, stats) */
export const mockArticleMinimal: Article = {
    id: "1007",
    slug: "pengumuman-maintenance-sistem",
    contentType: "article",
    title: "Pengumuman: Maintenance Sistem 29 Oktober 2026",
    category: "Pengumuman",
    author: mockAuthors.redaksi,
    publishedAt: "2026-10-28T18:00:00+07:00",
};

/** 8) ARTICLE — tanpa body (hanya untuk card stream) */
export const mockArticleCardOnly: Article = {
    id: "1008",
    slug: "harga-beras-naik-5-persen",
    contentType: "article",
    title: "Harga Beras Naik 5 Persen di Pasar Tradisional",
    excerpt: "Kenaikan dipicu musim panen yang mundur dan biaya distribusi.",
    category: "Ekonomi",
    badges: ["Terkini"],
    image: {
        src: "/assets/demo/beras.png",
        alt: "Harga Beras Naik 5 Persen di Pasar Tradisional",
        caption: "Harga Beras Naik 5 Persen di Pasar Tradisional",
        aspectRatio: "16/9",
    },
    author: mockAuthors.mayaSaraswati,
    publishedAt: "2026-10-28T11:45:00+07:00",
    stats: mockStats.fresh,
};
// ============================================================
// E. MOCK: KOLEKSI (untuk stream / feed)
// ============================================================

export const mockArticles: Article[] = [
    mockArticleSports,        // headline viral
    mockArticleNews,          // feature
    mockArticleColumn,        // kolom
    mockArticleCardOnly,      // card-only
    mockArticleGallery,       // gallery
    mockArticleVideo,         // video/podcast
    mockArticleEpaper,        // epaper
    mockArticleMinimal,       // minimal
    mockArticleRelate,       // minimal
];

/** Filter per contentType — untuk testing render komponen. */
export const mockByType = {
    article: mockArticles.filter((a) => a.contentType === "article"),
    gallery: mockArticles.filter((a) => a.contentType === "gallery"),
    video: mockArticles.filter((a) => a.contentType === "video"),
    epaper: mockArticles.filter((a) => a.contentType === "epaper"),
};

/** Untuk resolve `relatedIds` → Article. */
export const mockArticleMap: Record<string, Article> = Object.fromEntries(
    mockArticles.map((a) => [a.id, a])
);

// export function getRelated(article: Article): Article[] {
//     return (article.relatedIds ?? [])
//         .map((id) => mockArticleMap[id])
//         .filter(Boolean);
// }