CREATE TABLE strategies (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    user_id UUID NOT NULL,

    name VARCHAR(100) NOT NULL,
    description TEXT,

    color_hex CHAR(7) NOT NULL DEFAULT '#6366F1',

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,

    CONSTRAINT fk_strategies_user_id
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE,

    CONSTRAINT chk_strategies_color_hex
        CHECK (color_hex ~ '^#[0-9A-Fa-f]{6}$')
);

CREATE INDEX idx_strategies_user_id
    ON strategies (user_id)
    WHERE deleted_at IS NULL;

CREATE UNIQUE INDEX idx_strategies_user_name
    ON strategies (user_id, name)
    WHERE deleted_at IS NULL;

CREATE TRIGGER trg_strategies_set_updated_at
BEFORE UPDATE ON strategies
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();