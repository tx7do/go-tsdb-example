// https://jaxenter.com/apache-cassandra-iot-174970.html

CREATE KEYSPACE IF NOT EXISTS devices
    WITH replication = {
        'class' : 'SimpleStrategy',
        'replication_factor' : 1
        };

-- All events by device
CREATE TABLE IF NOT EXISTS devices.events_by_device
(
    device_id UUID,
    timestamp TIMESTAMP,
    state     TEXT,
    value     TEXT,
    PRIMARY KEY ((device_id), timestamp)
) WITH CLUSTERING ORDER BY (timestamp DESC);

-- Latest known events by device
CREATE TABLE IF NOT EXISTS devices.latest_events_by_device
(
    device_id UUID,
    timestamp TIMESTAMP,
    state     TEXT,
    value     TEXT,
    PRIMARY KEY ((device_id))
);

-- Latest events by state
CREATE INDEX IF NOT EXISTS latest_events_by_state_2i
    ON devices.latest_events_by_device (state);

-- Latest events by state
CREATE TABLE IF NOT EXISTS devices.latest_events_by_state
(
    state     TEXT,
    device_id UUID,
    timestamp TIMESTAMP,
    value     TEXT,
    PRIMARY KEY ((state), device_id)
);

-- Latest 'on' events by device
CREATE TABLE IF NOT EXISTS devices.latest_on_events_by_device
(
    device_id UUID,
    timestamp TIMESTAMP,
    value     TEXT,
    PRIMARY KEY ((device_id))
);

-- Latest 'off' events by device
CREATE TABLE IF NOT EXISTS devices.latest_off_events_by_device
(
    device_id UUID,
    timestamp TIMESTAMP,
    value     TEXT,
    PRIMARY KEY ((device_id))
);

-- Custom hash function
CREATE FUNCTION IF NOT EXISTS devices.hash(id UUID)
    RETURNS NULL ON NULL INPUT
    RETURNS INT
    LANGUAGE Java AS
    'return Integer.parseInt(id.toString().substring(0,3),16) % 3;';

-- Latest 'on' events by device
CREATE TABLE IF NOT EXISTS devices.latest_on_events_by_bucket
(
    bucket    INT,
    device_id UUID,
    timestamp TIMESTAMP,
    value     TEXT,
    PRIMARY KEY ((bucket), device_id)
);

-- Latest 'off' events by device
CREATE TABLE IF NOT EXISTS devices.latest_off_events_by_bucket
(
    bucket    INT,
    device_id UUID,
    timestamp TIMESTAMP,
    value     TEXT,
    PRIMARY KEY ((bucket), device_id)
);
