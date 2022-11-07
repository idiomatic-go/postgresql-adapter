DROP TABLE IF EXISTS slo_entry;

CREATE TABLE slo_entry(
    id INT,
    name VARCHAR(40) NOT NULL,
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