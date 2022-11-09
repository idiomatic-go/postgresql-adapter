DROP INDEX IF EXISTS idx_slo_entry_name

CREATE UNIQUE INDEX idx_slo_entry_name
    ON slo_entry(name);