CREATE OR REPLACE FUNCTION LogSLOEntryChanges()
  RETURNS TRIGGER
  LANGUAGE PLPGSQL
  AS
$$
BEGIN
IF (TG_OP = 'DELETE') THEN
	INSERT INTO slo_entry_log(id,slo_entry_id,name,changed_ts)
	VALUES(nextval('slo_entry_log_id'),OLD.id,OLD.name,TG_OP,now());
ELSE
    INSERT INTO slo_entry_log(id,slo_entry_id,name,changed_ts)
	VALUES(nextval('slo_entry_log_id'),NEW.id,NEW.name,TG_OP,now());
END IF;
RETURN NULL;
END;
$$

CREATE OR REPLACE FUNCTION msre.GetSLOEntry(id int)
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

CREATE OR REPLACE FUNCTION msre.GetSLOEntryByName(customerId int,name varchar(40))
  RETURNS SET OF slo_entry
  LANGUAGE PLPGSQL
  AS
$$
BEGIN
    SELECT *
    FROM slo_entry e
    WHERE e.customer_id = customerId AND e.name = name
END;
$$

CREATE OR REPLACE FUNCTION msre.GetSLOEntryBySegment(segments int, remainder int)
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

CREATE OR REPLACE FUNCTION msre.GetSLOEntryLog(low timestamp, high timestamp)
  RETURNS SET OF slo_entry
  LANGUAGE PLPGSQL
  AS
$$
BEGIN
    SELECT *
    FROM slo_entry_log l
    WHERE l.change_ts > low AND l.change_ts <= high
END;
$$