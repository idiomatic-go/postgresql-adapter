/* SLO entry - tables */
GRANT SELECT, INSERT, UPDATE, DELETE, TRIGGER
ON msre.slo_entry
TO slo_servicing;

GRANT SELECT, INSERT, UPDATE, DELETE
ON msre.slo_entry_log
TO slo_servicing;

/* SLO entry - sequences */
GRANT USAGE, SELECT, UPDATE
    ON msre.slo_entry_id
    TO slo_servicing;

GRANT USAGE, SELECT, UPDATE
    ON msre.slo_entry_log_id
    TO slo_servicing;

/* SLO entry - role */
GRANT slo_servicing TO markb


