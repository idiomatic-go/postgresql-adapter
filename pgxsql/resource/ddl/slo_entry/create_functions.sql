CREATE OR REPLACE FUNCTION LogSLOEntryChanges()
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

CREATE OR REPLACE FUNCTION GetSLOEntry(id int)
  RETURNS SET OF slo_entry
  LANGUAGE PLPGSQL
  AS
$$
BEGIN
    SELECT *
    FROM slo_entry e
    WHERE e.id = id
END;
$$

CREATE OR REPLACE FUNCTION GetSLOEntryByName(name varchar(40))
  RETURNS SET OF slo_entry
  LANGUAGE PLPGSQL
  AS
$$
BEGIN
    SELECT *
    FROM slo_entry e
    WHERE e.name = name
END;
$$

CREATE OR REPLACE FUNCTION GetSLOEntryBySegment(segments int, remainder int)
  RETURNS SET OF slo_entry
  LANGUAGE PLPGSQL
  AS
$$
BEGIN
    SELECT *
    FROM slo_entry e
    WHERE MOD(e.id,segments) = remainder
END;
$$