DROP TRIGGER IF EXISTS trg_simulation_rules_set_updated_at
ON simulation_rules;

DROP INDEX IF EXISTS idx_sim_rules_simulation_id;

DROP TABLE IF EXISTS simulation_rules;