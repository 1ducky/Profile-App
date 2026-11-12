CREATE TABLE session (
    id            VARCHAR(64) PRIMARY KEY,
    user_id       VARCHAR(40) NOT NULL,
    access_token  TEXT NOT NULL,
    refresh_token TEXT NOT NULL,
    revoke_at     DATETIME NOT NULL,
    version       INT NOT NULL DEFAULT 1
);

-- CREATE INDEX idx_session_user_id ON session(user_id);
-- CREATE UNIQUE INDEX idx_session_access_token ON session(access_token(255));
-- CREATE UNIQUE INDEX idx_session_refresh_token ON session(refresh_token(255));