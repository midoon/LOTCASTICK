CREATE TABLE simulation_rules (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    simulation_id UUID NOT NULL,

    drawdown_type drawdown_type_enum NOT NULL,

    max_drawdown_pct NUMERIC(6,4),
    max_drawdown_amount NUMERIC(20,8),

    daily_drawdown_pct NUMERIC(6,4),
    daily_drawdown_amount NUMERIC(20,8),

    trailing_drawdown_pct NUMERIC(6,4),
    trailing_drawdown_amount NUMERIC(20,8),
    trailing_high_water_mark NUMERIC(20,8),

    profit_target_pct NUMERIC(6,4) NOT NULL,
    profit_target_amount NUMERIC(20,8) NOT NULL,

    min_trading_days SMALLINT NOT NULL DEFAULT 0,

    consistency_rule_enabled BOOLEAN NOT NULL DEFAULT FALSE,
    consistency_threshold_pct NUMERIC(6,4),

    daily_reset_timezone VARCHAR(64) NOT NULL DEFAULT 'UTC',
    daily_reset_time TIME NOT NULL DEFAULT '00:00:00',

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_simulation_rules_simulation_id
        FOREIGN KEY (simulation_id)
        REFERENCES simulations(id)
        ON DELETE CASCADE,

    CONSTRAINT uq_simulation_rules_simulation_id
        UNIQUE (simulation_id),

    CONSTRAINT chk_simulation_rules_profit_target_pct
        CHECK (profit_target_pct > 0),

    CONSTRAINT chk_simulation_rules_min_trading_days
        CHECK (min_trading_days >= 0),

    CONSTRAINT chk_simulation_rules_consistency
        CHECK (
            CASE
                WHEN consistency_rule_enabled = TRUE
                    THEN consistency_threshold_pct IS NOT NULL
                     AND consistency_threshold_pct > 0
                ELSE TRUE
            END
        )
);

CREATE UNIQUE INDEX idx_sim_rules_simulation_id
    ON simulation_rules (simulation_id);

CREATE TRIGGER trg_simulation_rules_set_updated_at
BEFORE UPDATE ON simulation_rules
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();