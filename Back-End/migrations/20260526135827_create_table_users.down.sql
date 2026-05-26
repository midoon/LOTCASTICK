DROP TRIGGER IF EXISTS trg_users_set_updated_at ON users;

DROP FUNCTION IF EXISTS set_updated_at;

DROP INDEX IF EXISTS idx_users_deleted_at;

DROP INDEX IF EXISTS idx_users_email;

DROP TABLE IF EXISTS users;

DROP EXTENSION IF EXISTS pgcrypto;