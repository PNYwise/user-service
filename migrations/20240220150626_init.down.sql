DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS followers;

DROP INDEX IF EXISTS idx_users_deleted_at;
DROP INDEX IF EXISTS idx_users_created_at;

DROP INDEX IF EXISTS idx_users_username;
DROP INDEX IF EXISTS idx_users_email;

DROP INDEX IF EXISTS idx_followers_created_at;
DROP INDEX IF EXISTS idx_followers_followed_id;