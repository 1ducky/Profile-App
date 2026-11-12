CREATE TABLE IF NOT EXISTS `article` (
    `id`           VARCHAR(36)   NOT NULL,
    `slug`         VARCHAR(255)  NOT NULL,
    `url`          VARCHAR(500)  NOT NULL,
    `title`        VARCHAR(255)  NOT NULL,
    `excerpt`      TEXT          NULL,
    `category`     VARCHAR(100)  NOT NULL,
    `badges`       JSON          NULL,
    `tags`         JSON          NULL,
    `cover`        VARCHAR(500)  NULL,
    `author_id`    VARCHAR(36)   NOT NULL,
    `content`      LONGTEXT      NOT NULL,
    `published_at` DATETIME      NULL,
    `updated_at`   DATETIME      NULL,
    `created_at`   DATETIME      NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`),
    UNIQUE KEY `uq_article_slug` (`slug`),
    KEY `idx_article_category` (`category`),
    KEY `idx_article_author_id` (`author_id`),
    KEY `idx_article_published_at` (`published_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;