CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY, 
    uuid UUID DEFAULT uuid_generate_v4(),
    email VARCHAR NOT NULL,
    username VARCHAR(20) NOT NULL,
    timezone VARCHAR,
    followers BIGINT DEFAULT 0,
    following BIGINT DEFAULT 0,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    CONSTRAINT unique_users_username UNIQUE (username),
    CONSTRAINT unique_users_email UNIQUE (email),
    CONSTRAINT unique_users_uuid UNIQUE (uuid)
);
CREATE INDEX IF NOT EXISTS idx_users_deleted_at ON users (deleted_at) WHERE deleted_at IS NOT NULL;
CREATE INDEX IF NOT EXISTS idx_users_created_at ON users (created_at) WHERE created_at IS NOT NULL;
CREATE UNIQUE INDEX IF NOT EXISTS idx_users_username ON users (username);
CREATE UNIQUE INDEX IF NOT EXISTS idx_users_uuid ON users (uuid);
CREATE UNIQUE INDEX IF NOT EXISTS idx_users_email ON users (email) WHERE email IS NOT NULL;


CREATE TABLE IF NOT EXISTS followers (
    followed_id BIGINT NOT NULL,
    following_id BIGINT NOT NULL,
    created_at timestamp,
    CONSTRAINT fk_followed_user_id FOREIGN KEY (followed_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_following_user_id FOREIGN KEY (following_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT unique_followed_following_id UNIQUE (followed_id,following_id)
);
CREATE INDEX IF NOT EXISTS idx_followers_created_at ON followers (created_at) WHERE created_at IS NOT NULL;
CREATE INDEX IF NOT EXISTS idx_followers_followed_id ON followers (followed_id);
