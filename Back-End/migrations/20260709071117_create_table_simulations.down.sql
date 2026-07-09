DROP TRIGGER IF EXISTS trg_simulations_set_updated_at
ON simulations;

DROP INDEX IF EXISTS idx_simulations_user_status;

DROP INDEX IF EXISTS idx_simulations_deleted_at;

DROP INDEX IF EXISTS idx_simulations_status;

DROP INDEX IF EXISTS idx_simulations_user_id;

DROP TABLE IF EXISTS simulations;