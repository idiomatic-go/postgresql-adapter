GRANT SELECT, INSERT, UPDATE, DELETE, TRIGGER
ON slo_entry
TO slo_servicing;

GRANT SELECT, INSERT, UPDATE, DELETE
ON slo_entry_log
TO slo_servicing;

GRANT USAGE, SELECT, UPDATE
    ON slo_entry_id
    TO slo_servicing;

GRANT USAGE, SELECT, UPDATE
    ON slo_entry_log_id
    TO slo_servicing;

GRANT slo_servicing TO markb