DROP INDEX IF EXISTS idx_customer_org

CREATE UNIQUE INDEX idx_customer_org
    ON customer(org_id);


DROP INDEX IF EXISTS idx_customer_log_id

CREATE INDEX idx_customer_log_id
    ON customer_log(customer_id);