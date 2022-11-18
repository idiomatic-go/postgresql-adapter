DROP TABLE IF EXISTS access_log_ingress;

CREATE TABLE access_log_ingress (
    customer_id INT NOT NULL,
    service VARCHAR(40) NOT NULL,
    route VARCHAR(40),

    region VARCHAR(40),
    zone VARCHAR(40),
    sub_zone VARCHAR(40),

    http_method VARCHAR(40),
    http_protocol BOOLEAN NOT NULL,
    grpc_protocol BOOLEAN NOT NULL,
    envoy_status BOOLEAN NOT NULL,

    status_class INT NOT NULL,
    status_code INT NOT NULL,
    grpc_status INT NOT NULL,

    start_time TIMESTAMP(6) NOT NULL,
    duration_ms INT NOT NULL,
    /*PRIMARY KEY(customer_id)*/
);

