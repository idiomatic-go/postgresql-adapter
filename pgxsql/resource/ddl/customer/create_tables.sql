DROP TABLE IF EXISTS customer;

CREATE TABLE customer (
    id INT,
    org_id VARCHAR(40) NOT NULL,
    track_slo_activity BOOLEAN,
    created_ts TIMESTAMP(6) NOT NULL,
    updated_ts TIMESTAMP(6) NOT NULL,
    PRIMARY KEY(id)
);

DROP TABLE IF EXISTS customer_log;

CREATE TABLE customer_log (
    id INT,
    customer_id INT NOT NULL,
    org_id VARCHAR(40) NOT NULL,
    operation TEXT NOT NULL,
    changed_ts TIMESTAMP(6) NOT NULL,
    PRIMARY KEY(id)
);

DROP TABLE IF EXISTS customer_locality;

CREATE TABLE customer_locality (
     id INT,
     customer_id INT NOT NULL,
     region VARCHAR(40),
     zone VARCHAR(40),
     sub_zone VARCHAR(40),
     PRIMARY KEY(id)
);
