DROP TABLE IF EXISTS counter;

CREATE TABLE counter (
    id INT NOT NULL,
    org_id VARCHAR(40) NOT NULL,
    track_activity BOOLEAN NOT NULL,
    created_ts TIMESTAMP(6) NOT NULL,
    changed_ts TIMESTAMP(6) ,
    PRIMARY KEY(id,org_id)
);

