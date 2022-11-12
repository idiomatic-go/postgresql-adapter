CREATE OR REPLACE FUNCTION msre.LogCustomerChanges()
  RETURNS TRIGGER
  LANGUAGE PLPGSQL
  AS
$$
BEGIN
IF (TG_OP = 'DELETE') THEN
	INSERT INTO customer_log(id,customer_id,org_id,changed_ts)
	VALUES(nextval('customer_log_id'),OLD.id,OLD.org_id,TG_OP,now());
ELSE
    INSERT INTO customer_log(id,customer_id,org_id,changed_ts)
	VALUES(nextval('customer_log_id'),NEW.id,NEW.org_id,TG_OP,now());
END IF;
RETURN NULL;
END;
$$

CREATE OR REPLACE FUNCTION msre.GetCustomer(id int)
  RETURNS SET OF customer
  LANGUAGE PLPGSQL
  AS
$$
BEGIN
    SELECT *
    FROM customer c
    WHERE c.id = id
END;
$$

CREATE OR REPLACE FUNCTION msre.GetCustomerByOrg(id int,org_id varchar(40))
  RETURNS SET OF customer
  LANGUAGE PLPGSQL
  AS
$$
BEGIN
    SELECT *
    FROM customer c
    WHERE c.id = id AND c.org_id = org_id
END;
$$

CREATE OR REPLACE FUNCTION msre.GetCustomerBySegment(segments int, remainder int)
  RETURNS SET OF customer
  LANGUAGE PLPGSQL
  AS
$$
BEGIN
    SELECT *
    FROM customer c
    WHERE MOD(c.id,segments) = remainder
END;
$$

CREATE OR REPLACE FUNCTION msre.GetCustomerLog(low timestamp, high timestamp)
  RETURNS SET OF slo_entry
  LANGUAGE PLPGSQL
  AS
$$
BEGIN
    SELECT *
    FROM customer_log l
    WHERE l.change_ts > low AND l.change_ts <= high
END;
$$