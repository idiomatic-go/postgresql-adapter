DROP INDEX IF EXISTS msre.idx_slo_entry_name

CREATE UNIQUE INDEX msre.idx_slo_entry_name
    ON msre.slo_entry(name);