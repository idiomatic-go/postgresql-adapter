DROP TABLE IF EXISTS ping;

CREATE TABLE ping (
    customer_id INT NOT NULL,
    application VARCHAR(40) NOT NULL,

    region VARCHAR(40),
    zone VARCHAR(40),
    sub_zone VARCHAR(40),

    status_code INT NOT NULL,

    start_time TIMESTAMP(6) NOT NULL,
    duration_ms INT NOT NULL,
   /* PRIMARY KEY(customer_id)*/
);
