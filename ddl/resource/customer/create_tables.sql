DROP TABLE IF EXISTS customer;

CREATE TABLE customer (
    id INT NOT NULL,
    org_id VARCHAR(40) NOT NULL,
    track_activity BOOLEAN NOT NULL,
    created_ts TIMESTAMP(6) NOT NULL,
    changed_ts TIMESTAMP(6) ,
    PRIMARY KEY(id,org_id)
);

DROP TABLE IF EXISTS customer_log;

CREATE TABLE customer_log (
    id INT NOT NULL,
    customer_id INT REFERENCES customer (id),
    org_id VARCHAR(40) NOT NULL,
    operation VARCHAR(40) NOT NULL,
    changed_ts TIMESTAMP(6) NOT NULL,
    PRIMARY KEY(id,customer_id)
);

DROP TABLE IF EXISTS customer_locality;

CREATE TABLE customer_locality (
     id INT NOT NULL,
     customer_id INT REFERENCES customer (id),
     region VARCHAR(40) NOT NULL,
     zone VARCHAR(40) NOT NULL,
     sub_zone VARCHAR(40) NOT NULL,
     created_ts TIMESTAMP(6) NOT NULL,
     PRIMARY KEY(id,customer_id)
);

DROP TABLE IF EXISTS customer_metric;

CREATE TABLE customer_metric (
     id INT NOT NULL,
     customer_id INT REFERENCES customer (id),
     application VARCHAR(40) NOT NULL,
     route_name VARCHAR(40) NOT NULL,
     region VARCHAR(40) NOT NULL,
     zone VARCHAR(40) NOT NULL,
     sub_zone VARCHAR(40) NOT NULL,
     name VARCHAR(40) NOT NULL,
     value INT NOT NULL,
     created_ts TIMESTAMP(6) NOT NULL,
     changed_ts TIMESTAMP(6) ,
     PRIMARY KEY(id,customer_id)
);

