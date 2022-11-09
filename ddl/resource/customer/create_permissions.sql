/* Customer - tables */
GRANT SELECT, INSERT, UPDATE, DELETE, TRIGGER
    ON customer
    TO customer_servicing;

GRANT SELECT, INSERT, UPDATE, DELETE
    ON customer_log
    TO customer_servicing;

/* Customer - sequences */
GRANT USAGE, SELECT, UPDATE
    ON customer_id
    TO customer_servicing;

GRANT USAGE, SELECT, UPDATE
    ON customer_log_id
    TO customer_servicing;

GRANT USAGE, SELECT, UPDATE
    ON customer_locality_id
    TO customer_servicing;

GRANT USAGE, SELECT, UPDATE
    ON customer_metric_id
    TO customer_servicing;

/* Customer - role */
GRANT customer_servicing TO markb