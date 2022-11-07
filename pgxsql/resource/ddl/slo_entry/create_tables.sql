DROP TABLE IF EXISTS slo_entry;

CREATE TABLE slo_entry(
    id INT,
    customer_id INT NOT NULL,
    name VARCHAR(40) NOT NULL,
    application VARCHAR(40) NOT NULL,
    route_name VARCHAR(40) NOT NULL,
    category INT NOT NULL,
    traffic_type INT NOT NULL,
    traffic_protocol INT,
    processing_interval INT,
    window_interval INT,
    watch_percent INT,
    threshold_percent INT,
    threshold_value INT,
    threshold_minimum INT,
    locality_scope INT,
    rps_low_comparison INT,
    rps_high_comparison INT,
    filter_status_codes VARCHAR(40),
    status_codes VARCHAR(40),
    metric_name VARCHAR(40),
    metric_name_secondary VARCHAR(40),
    function_name VARCHAR(40),
    disable_processing BOOLEAN,
    disable_triage BOOLEAN
    PRIMARY KEY(id)
);

DROP TABLE IF EXISTS slo_entry_log;

CREATE TABLE slo_entry_log (
    id INT,
    slo_entry_id INT NOT NULL,
    name VARCHAR(40) NOT NULL,
    operation TEXT NOT NULL,
    changed_on TIMESTAMP(6) NOT NULL,
    PRIMARY KEY(id)
);