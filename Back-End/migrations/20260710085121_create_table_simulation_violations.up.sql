CREATE TABLE simulation_violations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    simulation_id UUID NOT NULL,
    trade_id UUID,

    violation_type VARCHAR(50) NOT NULL,

    equity_at_violation NUMERIC(20,8) NOT NULL,
    rule_floor_at_violation NUMERIC(20,8) NOT NULL,

    occurred_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_simulation_violations_simulation_id
        FOREIGN KEY (simulation_id)
        REFERENCES simulations(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_simulation_violations_trade_id
        FOREIGN KEY (trade_id)
        REFERENCES trades(id)
        ON DELETE SET NULL
);

CREATE INDEX idx_violations_simulation_id
    ON simulation_violations (simulation_id);

CREATE INDEX idx_violations_occurred_at
    ON simulation_violations (occurred_at);