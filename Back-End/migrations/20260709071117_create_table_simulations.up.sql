CREATE TABLE simulations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    user_id UUID NOT NULL,
    template_id UUID,

    name VARCHAR(150) NOT NULL,

    account_size NUMERIC(20,8) NOT NULL,
    current_equity NUMERIC(20,8) NOT NULL,

    currency CHAR(3) NOT NULL DEFAULT 'USD',

    status simulation_status NOT NULL DEFAULT 'ACTIVE',

    started_at DATE NOT NULL DEFAULT CURRENT_DATE,
    passed_at TIMESTAMPTZ,
    failed_at TIMESTAMPTZ,

    notes TEXT,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,

    CONSTRAINT fk_simulations_user_id
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_simulations_template_id
        FOREIGN KEY (template_id)
        REFERENCES challenge_templates(id)
        ON DELETE SET NULL,

    CONSTRAINT chk_simulations_account_size
        CHECK (account_size > 0),

    CONSTRAINT chk_simulations_current_equity
        CHECK (current_equity >= 0),

    CONSTRAINT chk_simulations_currency
        CHECK (currency ~ '^[A-Z]{3}$')
);

CREATE INDEX idx_simulations_user_id
    ON simulations (user_id);

CREATE INDEX idx_simulations_status
    ON simulations (status);

CREATE INDEX idx_simulations_deleted_at
    ON simulations (deleted_at);

CREATE INDEX idx_simulations_user_status
    ON simulations (user_id, status)
    WHERE deleted_at IS NULL;

CREATE TRIGGER trg_simulations_set_updated_at
BEFORE UPDATE ON simulations
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();