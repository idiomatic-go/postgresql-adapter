/* Defaults to bigint data type */
DROP SEQUENCE IF EXISTS msre.slo_entry_id

CREATE SEQUENCE msre.slo_entry_id
    START 2
INCREMENT 2
MINVALUE 2;

DROP SEQUENCE IF EXISTS msre.slo_entry_log_id

CREATE SEQUENCE msre.slo_entry_log_id
    START 2
INCREMENT 2
MINVALUE 2;


/*
 DROP SEQUENCE slo_entry_id
 */