/* Defaults to bigint data type */
DROP SEQUENCE IF EXISTS customer_id

CREATE SEQUENCE customer_id
    START 1
INCREMENT 2
MINVALUE 1;

DROP SEQUENCE IF EXISTS customer_log_id

CREATE SEQUENCE customer_log_id
    START 1
INCREMENT 2
MINVALUE 1;

DROP SEQUENCE IF EXISTS customer_locality_id

CREATE SEQUENCE customer_locality_id
    START 1
INCREMENT 2
MINVALUE 1;

DROP SEQUENCE IF EXISTS customer_metric_id

CREATE SEQUENCE customer_metric_id
    START 1
INCREMENT 2
MINVALUE 1;


/*
 DROP SEQUENCE customerid
 */