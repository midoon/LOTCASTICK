CREATE TABLE trades (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    simulation_id UUID NOT NULL,
    user_id UUID NOT NULL,
    strategy_id UUID,

    symbol VARCHAR(20) NOT NULL,
    direction trade_direction NOT NULL,

    entry_price NUMERIC(20,8) NOT NULL,
    exit_price NUMERIC(20,8) NOT NULL,
    stop_loss NUMERIC(20,8),
    take_profit NUMERIC(20,8),

    lot_size NUMERIC(12,4) NOT NULL,

    pnl NUMERIC(20,8) NOT NULL,
    pnl_pct NUMERIC(10,6) NOT NULL,

    risk_amount NUMERIC(20,8),
    rr_ratio NUMERIC(10,4),

    session trade_session,

    entry_time TIMESTAMPTZ NOT NULL,
    exit_time TIMESTAMPTZ NOT NULL,
    trade_date DATE NOT NULL,

    notes TEXT,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,

    CONSTRAINT fk_trades_simulation_id
        FOREIGN KEY (simulation_id)
        REFERENCES simulations(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_trades_user_id
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_trades_strategy_id
        FOREIGN KEY (strategy_id)
        REFERENCES strategies(id)
        ON DELETE SET NULL,

    CONSTRAINT chk_trades_lot_size
        CHECK (lot_size > 0),

    CONSTRAINT chk_trades_entry_time
        CHECK (entry_time < exit_time),

    CONSTRAINT chk_trades_entry_price
        CHECK (entry_price > 0),

    CONSTRAINT chk_trades_exit_price
        CHECK (exit_price > 0)
);

CREATE INDEX idx_trades_simulation_id
    ON trades (simulation_id)
    WHERE deleted_at IS NULL;

CREATE INDEX idx_trades_user_id
    ON trades (user_id)
    WHERE deleted_at IS NULL;

CREATE INDEX idx_trades_trade_date
    ON trades (simulation_id, trade_date)
    WHERE deleted_at IS NULL;

CREATE INDEX idx_trades_symbol
    ON trades (simulation_id, symbol)
    WHERE deleted_at IS NULL;

CREATE INDEX idx_trades_strategy_id
    ON trades (strategy_id)
    WHERE deleted_at IS NULL;

CREATE INDEX idx_trades_exit_time
    ON trades (simulation_id, exit_time DESC)
    WHERE deleted_at IS NULL;

CREATE TRIGGER trg_trades_set_updated_at
BEFORE UPDATE ON trades
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();