CREATE TABLE assets (
    id VARCHAR(64) PRIMARY KEY,
    parent_id VARCHAR(64),
    filename VARCHAR(255) NOT NULL,
    file_key VARCHAR(255) NOT NULL,
    mime VARCHAR(255) NOT NULL,
    size BIGINT NOT NULL DEFAULT 0,
    author_id VARCHAR(40) NOT NULL,
    status ENUM('temp', 'pending', 'active', 'orphan', 'deleted', 'failed') NOT NULL,
    category ENUM('avatars', 'document', 'post', 'temp', 'video') NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
