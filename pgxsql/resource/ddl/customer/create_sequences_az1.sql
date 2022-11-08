/* Defaults to bigint data type */
DROP SEQUENCE IF EXISTS customer_id

CREATE SEQUENCE customer_id
    START 2
INCREMENT 2
MINVALUE 2;

DROP SEQUENCE IF EXISTS customer_log_id

CREATE SEQUENCE customer_log_id
    START 2
INCREMENT 2
MINVALUE 2;

DROP SEQUENCE IF EXISTS customer_locality_id

CREATE SEQUENCE customer_locality_id
    START 2
INCREMENT 2
MINVALUE 2;



/*
 DROP SEQUENCE slo_entry_id
 */