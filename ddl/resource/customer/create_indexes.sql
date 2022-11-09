DROP INDEX IF EXISTS idx_customer_org

CREATE UNIQUE INDEX idx_customer_org
    ON customer(org_id);


