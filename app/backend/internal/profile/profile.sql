CREATE TABLE profiles (
    id VARCHAR(64) PRIMARY KEY,
    username VARCHAR(64) NOT NULL,
    tag VARCHAR(64) NOT NULL UNIQUE,
    bio TEXT,
    avatar_url TEXT,
    cover_url TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Index untuk mempercepat pencarian berdasarkan tag (karena ditandai sebagai //indexed di entity)
CREATE INDEX idx_profiles_tag ON profiles(tag);
