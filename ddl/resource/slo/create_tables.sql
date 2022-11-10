DROP TABLE IF EXISTS msre.slo_entry;

CREATE TABLE msre.slo_entry(
    id INT NOT NULL,
    customer_id INT  REFERENCES customer (id),

    category INT NOT NULL,
    traffic_type INT NOT NULL,
    traffic_protocol INT NOT NULL,
    processing_interval INT NOT NULL,
    window_interval INT NOT NULL,
    watch_percent INT NOT NULL,
    threshold_percent INT NOT NULL,
    threshold_value INT NOT NULL,
    threshold_minimum INT NOT NULL,

    rps_low_comparison INT NOT NULL,
    rps_high_comparison INT NOT NULL,
    locality_scope INT NOT NULL,
    disable_processing BOOLEAN NOT NULL,
    disable_triage BOOLEAN NOT NULL,

    filter_status_codes VARCHAR(40),
    status_codes VARCHAR(40),
    name VARCHAR(40) NOT NULL,
    application VARCHAR(40) NOT NULL,
    route_name VARCHAR(40),

    created_ts TIMESTAMP(4) DEFAULT now(),
    changed_ts TIMESTAMP(4),
    PRIMARY KEY(id,customer_id,name)
);

DROP TABLE IF EXISTS msre.slo_entry_log;

CREATE TABLE msre.slo_entry_log (
    id INT NOT NULL,
    slo_entry_id INT NOT NULL,
    name VARCHAR(40) NOT NULL,
    operation VARCHAR(40) NOT NULL,
    changed_ts TIMESTAMP(4) DEFAULT now(),
    PRIMARY KEY(id,slo_entry_id)
);