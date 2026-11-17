-- Postgres schema (converted from MySQL dump)

-- ---------- ENUM types ----------
CREATE TYPE asset_status   AS ENUM ('temp','pending','active','orphan','deleted','failed');
CREATE TYPE asset_category AS ENUM ('avatars','document','post','temp','video');
CREATE TYPE credential_role AS ENUM ('ADMIN','USER');

-- ---------- Trigger helper for "ON UPDATE CURRENT_TIMESTAMP" ----------
CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = CURRENT_TIMESTAMP;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- ---------- article ----------
CREATE TABLE article (
  id           varchar(36)  NOT NULL,
  slug         varchar(255) NOT NULL,
  url          varchar(500) DEFAULT NULL,
  title        varchar(255) NOT NULL,
  excerpt      text,
  category     varchar(100) NOT NULL,
  badges       jsonb DEFAULT NULL,
  tags         jsonb DEFAULT NULL,
  cover        varchar(500) DEFAULT NULL,
  author_id    varchar(36)  NOT NULL,
  content      text,
  published_at timestamp DEFAULT NULL,
  updated_at   timestamp DEFAULT NULL,
  created_at   timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  CONSTRAINT uq_article_slug UNIQUE (slug)
);
CREATE INDEX idx_article_category     ON article (category);
CREATE INDEX idx_article_author_id    ON article (author_id);
CREATE INDEX idx_article_published_at ON article (published_at);

-- ---------- article_category ----------
CREATE TABLE article_category (
  id   varchar(36)  NOT NULL,
  slug varchar(255) NOT NULL,
  PRIMARY KEY (id),
  CONSTRAINT uq_article_category_slug UNIQUE (slug)
);

-- ---------- assets ----------
CREATE TABLE assets (
  id         varchar(64)  NOT NULL,
  parent_id  varchar(64)  DEFAULT NULL,
  filename   varchar(255) NOT NULL,
  file_key   varchar(255) NOT NULL,
  mime       varchar(255) NOT NULL,
  size       bigint       NOT NULL DEFAULT 0,
  author_id  varchar(40)  NOT NULL,
  status     asset_status   NOT NULL,
  category   asset_category NOT NULL,
  created_at timestamp DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);
CREATE TRIGGER trg_assets_updated_at
  BEFORE UPDATE ON assets
  FOR EACH ROW EXECUTE FUNCTION set_updated_at();

-- ---------- credential ----------
CREATE TABLE credential (
  id       varchar(64)  NOT NULL,
  user_id  varchar(40)  NOT NULL,
  email    varchar(255) NOT NULL,
  password varchar(255) NOT NULL,
  role     credential_role NOT NULL DEFAULT 'USER',
  PRIMARY KEY (id),
  CONSTRAINT credential_user_id_key UNIQUE (user_id),
  CONSTRAINT credential_email_key   UNIQUE (email)
);

-- ---------- profiles ----------
CREATE TABLE profiles (
  id         varchar(64) NOT NULL,
  username   varchar(64) NOT NULL,
  tag        varchar(64) NOT NULL,
  bio        text,
  avatar_url text,
  cover_url  text,
  created_at timestamp DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  CONSTRAINT profiles_tag_key UNIQUE (tag)
);
CREATE TRIGGER trg_profiles_updated_at
  BEFORE UPDATE ON profiles
  FOR EACH ROW EXECUTE FUNCTION set_updated_at();

-- ---------- session ----------
CREATE TABLE session (
  id            varchar(64) NOT NULL,
  user_id       varchar(40) NOT NULL,
  access_token  text        NOT NULL,
  refresh_token text        NOT NULL,
  revoke_at     timestamp   NOT NULL,
  version       integer     NOT NULL DEFAULT 1,
  PRIMARY KEY (id)
);