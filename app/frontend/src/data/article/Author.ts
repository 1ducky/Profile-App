import type { Author } from "../../types/article/entity";

// ============================================================
// A. MOCK: AUTHOR
// ============================================================

export const mockAuthors: Record<string, Author> = {
    dahlanIskan: {
        name: "Dahlan Iskan & Redaksi",
        initials: "DI",
        url: "/author/dahlan-iskan",
        isColumnist: true,
    },
    mayaSaraswati: {
        name: "Maya Saraswati",
        initials: "MS",
        url: "/author/maya-saraswati",
    },
    azrulAnanda: {
        name: "Azrul Ananda",
        initials: "AA",
        url: "/author/azrul-ananda",
        isColumnist: true,
    },
    redaksi: {
        name: "Redaksi Disway",
        initials: "RD",
    },
    biroBeijing: {
        name: "Biro Beijing Disway Media",
        initials: "BJ",
    },
};








