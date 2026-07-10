DROP TRIGGER IF EXISTS trg_trades_set_updated_at
ON trades;

DROP INDEX IF EXISTS idx_trades_exit_time;

DROP INDEX IF EXISTS idx_trades_strategy_id;

DROP INDEX IF EXISTS idx_trades_symbol;

DROP INDEX IF EXISTS idx_trades_trade_date;

DROP INDEX IF EXISTS idx_trades_user_id;

DROP INDEX IF EXISTS idx_trades_simulation_id;

DROP TABLE IF EXISTS trades;