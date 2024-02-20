CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY, 
    uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email VARCHAR NOT NULL,
    username VARCHAR(20) NOT NULL,
    timezone VARCHAR,
    follower BIGINT DEFAULT 0;
    following BIGINT DEFAULT 0;
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    CONSTRAINT unique_users_username UNIQUE (username)
);
CREATE INDEX IF NOT EXISTS idx_users_deleted_at ON users (deleted_at) WHERE deleted_at IS NOT NULL;
CREATE INDEX IF NOT EXISTS idx_users_created_at ON users (created_at) WHERE created_at IS NOT NULL;


CREATE TABLE IF NOT EXISTS followers (
    followed_id BIGINT NOT NULL,
    following_id BIGINT NOT NULL,
    created_at timestamp,
    CONSTRAINT fk_followed_user_id FOREIGN KEY (follower_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_following_user_id FOREIGN KEY (following_id) REFERENCES users(id) ON DELETE CASCADE
);
CREATE INDEX IF NOT EXISTS idx_created_at ON followers (created_at) WHERE created_at IS NOT NULL;
