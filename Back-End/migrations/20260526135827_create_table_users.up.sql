CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    email VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    display_name VARCHAR(100) NOT NULL,

    timezone VARCHAR(64) NOT NULL DEFAULT 'UTC',

    default_currency CHAR(3) NOT NULL DEFAULT 'USD'
        CHECK (default_currency ~ '^[A-Z]{3}$'),

    login_attempts SMALLINT NOT NULL DEFAULT 0,
    locked_until TIMESTAMPTZ,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,

    CONSTRAINT chk_users_login_attempts
        CHECK (login_attempts >= 0)
);

CREATE UNIQUE INDEX idx_users_email
    ON users (email)
    WHERE deleted_at IS NULL;

CREATE INDEX idx_users_deleted_at
    ON users (deleted_at);

CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_users_set_updated_at
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();