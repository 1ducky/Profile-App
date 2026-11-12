CREATE TABLE article_category (
    id   VARCHAR(36)  NOT NULL,
    slug VARCHAR(255) NOT NULL,
    PRIMARY KEY (id),
    UNIQUE KEY uq_article_category_slug (slug)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;