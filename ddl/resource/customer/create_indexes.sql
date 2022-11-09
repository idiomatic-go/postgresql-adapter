DROP INDEX IF EXISTS msre.idx_customer_org

CREATE UNIQUE INDEX msre.idx_customer_org
    ON customer(org_id);


