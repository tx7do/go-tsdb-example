-- https://docs.timescale.com/timescaledb/latest/tutorials/nyc-taxi-cab

CREATE TABLE "rides"
(
    vendor_id             TEXT,
    pickup_datetime       TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    dropoff_datetime      TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    passenger_count       NUMERIC,
    trip_distance         NUMERIC,
    pickup_longitude      NUMERIC,
    pickup_latitude       NUMERIC,
    rate_code             INTEGER,
    dropoff_longitude     NUMERIC,
    dropoff_latitude      NUMERIC,
    payment_type          INTEGER,
    fare_amount           NUMERIC,
    extra                 NUMERIC,
    mta_tax               NUMERIC,
    tip_amount            NUMERIC,
    tolls_amount          NUMERIC,
    improvement_surcharge NUMERIC,
    total_amount          NUMERIC
);
SELECT create_hypertable('rides', 'pickup_datetime');

CREATE TABLE IF NOT EXISTS "payment_types"
(
    payment_type INTEGER,
    description  TEXT
);
INSERT INTO payment_types(payment_type, description)
VALUES (1, 'credit card'),
       (2, 'cash'),
       (3, 'no charge'),
       (4, 'dispute'),
       (5, 'unknown'),
       (6, 'voided trip');

CREATE TABLE IF NOT EXISTS "rates"
(
    rate_code   INTEGER,
    description TEXT
);
INSERT INTO rates(rate_code, description)
VALUES (1, 'standard rate'),
       (2, 'JFK'),
       (3, 'Newark'),
       (4, 'Nassau or Westchester'),
       (5, 'negotiated fare'),
       (6, 'group ride');

