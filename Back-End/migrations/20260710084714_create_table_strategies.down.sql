DROP TRIGGER IF EXISTS trg_strategies_set_updated_at
ON strategies;

DROP INDEX IF EXISTS idx_strategies_user_name;

DROP INDEX IF EXISTS idx_strategies_user_id;

DROP TABLE IF EXISTS strategies;