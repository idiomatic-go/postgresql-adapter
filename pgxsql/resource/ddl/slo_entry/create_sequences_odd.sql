/* Defaults to bigint data type */
DROP SEQUENCE IF EXISTS slo_entry_id

CREATE SEQUENCE slo_entry_id
    START 1
INCREMENT 2
MINVALUE 1;

DROP SEQUENCE IF EXISTS slo_entry_log_id

CREATE SEQUENCE slo_entry_log_id
    START 1
INCREMENT 2
MINVALUE 1;


/*
 DROP SEQUENCE slo_entry_id
 */