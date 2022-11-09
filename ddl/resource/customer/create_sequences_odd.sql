/* Defaults to bigint data type */
DROP SEQUENCE IF EXISTS msre.customer_id

CREATE SEQUENCE msre.customer_id
    START 1
INCREMENT 2
MINVALUE 1;

DROP SEQUENCE IF EXISTS msre.customer_log_id

CREATE SEQUENCE msre.customer_log_id
    START 1
INCREMENT 2
MINVALUE 1;

DROP SEQUENCE IF EXISTS msre.customer_locality_id

CREATE SEQUENCE msre.customer_locality_id
    START 1
INCREMENT 2
MINVALUE 1;

DROP SEQUENCE IF EXISTS msre.customer_metric_id

CREATE SEQUENCE msre.customer_metric_id
    START 1
INCREMENT 2
MINVALUE 1;


/*
 DROP SEQUENCE customerid
 */