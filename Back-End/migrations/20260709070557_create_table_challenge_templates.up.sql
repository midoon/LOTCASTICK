CREATE TABLE challenge_templates (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    name VARCHAR(150) NOT NULL,
    firm_name VARCHAR(100) NOT NULL,

    account_size NUMERIC(20,8) NOT NULL,

    currency CHAR(3) NOT NULL DEFAULT 'USD',

    drawdown_type drawdown_type_enum NOT NULL,

    max_drawdown_pct NUMERIC(6,4),
    daily_drawdown_pct NUMERIC(6,4),
    trailing_drawdown_pct NUMERIC(6,4),

    profit_target_pct NUMERIC(6,4) NOT NULL,

    min_trading_days SMALLINT NOT NULL DEFAULT 0,

    consistency_rule_enabled BOOLEAN NOT NULL DEFAULT FALSE,
    consistency_threshold_pct NUMERIC(6,4),

    is_active BOOLEAN NOT NULL DEFAULT TRUE,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT uq_challenge_templates_name
        UNIQUE (name),

    CONSTRAINT chk_challenge_templates_account_size
        CHECK (account_size > 0),

    CONSTRAINT chk_challenge_templates_profit_target
        CHECK (profit_target_pct > 0),

    CONSTRAINT chk_challenge_templates_currency
        CHECK (currency ~ '^[A-Z]{3}$')
);


CREATE INDEX idx_challenge_templates_firm_name
    ON challenge_templates (firm_name);

CREATE INDEX idx_challenge_templates_is_active
    ON challenge_templates (is_active);