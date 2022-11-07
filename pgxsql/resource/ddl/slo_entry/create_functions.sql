CREATE OR REPLACE FUNCTION log_slo_entry_changes()
  RETURNS TRIGGER
  LANGUAGE PLPGSQL
  AS
$$
BEGIN
IF (TG_OP = 'DELETE') THEN
	INSERT INTO slo_entry_log(id,slo_entry_id,name,changed_on)
	VALUES(nextval('slo_entry_log_id'),OLD.id,OLD.name,TG_OP,now());
ELSE
    INSERT INTO slo_entry_log(id,slo_entry_id,name,changed_on)
	VALUES(nextval('slo_entry_log_id'),NEW.id,NEW.name,TG_OP,now());
END IF;
RETURN NULL;
END;
$$