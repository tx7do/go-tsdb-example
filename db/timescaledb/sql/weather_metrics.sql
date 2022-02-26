-- https://docs.timescale.com/timescaledb/latest/getting-started/launch-timescaledb/

----------------------------------------
-- Hypertable to store weather metrics
----------------------------------------
-- Step 1: Define regular table
CREATE TABLE IF NOT EXISTS weather_metrics
(
    time             TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    timezone_shift   int                         NULL,
    city_name        text                        NULL,
    temp_c           double PRECISION            NULL,
    feels_like_c     double PRECISION            NULL,
    temp_min_c       double PRECISION            NULL,
    temp_max_c       double PRECISION            NULL,
    pressure_hpa     double PRECISION            NULL,
    humidity_percent double PRECISION            NULL,
    wind_speed_ms    double PRECISION            NULL,
    wind_deg         int                         NULL,
    rain_1h_mm       double PRECISION            NULL,
    rain_3h_mm       double PRECISION            NULL,
    snow_1h_mm       double PRECISION            NULL,
    snow_3h_mm       double PRECISION            NULL,
    clouds_percent   int                         NULL,
    weather_type_id  int                         NULL
);

-- Step 2: Turn into hypertable
SELECT create_hypertable('weather_metrics', 'time');

-- https://s3.amazonaws.com/assets.timescale.com/docs/downloads/weather_data.zip
-- \copy weather_metrics (time, timezone_shift, city_name, temp_c, feels_like_c, temp_min_c, temp_max_c, pressure_hpa, humidity_percent, wind_speed_ms, wind_deg, rain_1h_mm, rain_3h_mm, snow_1h_mm, snow_3h_mm, clouds_percent, weather_type_id) from './weather_data.csv' CSV HEADER;


--------------------------------
-- Average temperature per city
-- in past 2 years
--------------------------------
SELECT city_name, avg(temp_c)
FROM weather_metrics
WHERE time > now() - INTERVAL '2 years'
GROUP BY city_name;

--------------------------------
-- Total snowfall per city
-- in past 5 years
--------------------------------
SELECT city_name, sum(snow_1h_mm)
FROM weather_metrics
WHERE time > now() - INTERVAL '5 years'
GROUP BY city_name;

-----------------------------------
-- time_bucket
-- Average temp per 15 day period
-- for past 6 months, per city
-----------------------------------
SELECT time_bucket('15 days', time) as "bucket"
     , city_name
     , avg(temp_c)
FROM weather_metrics
WHERE time > now() - (12 * INTERVAL '1 month')
GROUP BY bucket, city_name
ORDER BY bucket DESC;

-- non-gapfill query
SELECT time_bucket('30 days', time) as bucket,
       city_name,
       sum(snow_1h_mm)              as sum
FROM weather_metrics
WHERE time > now() - INTERVAL '1 year'
  AND time < now()
GROUP BY bucket, city_name
ORDER BY bucket DESC;

-----------------------------------------
-- time_bucket_gapfill
-- total snow fall per city
-- in 30-day buckets for past 1 year
-----------------------------------------
SELECT time_bucket_gapfill('30 days', time) as bucket,
       city_name,
       sum(snow_1h_mm)                      as sum
FROM weather_metrics
WHERE time > now() - INTERVAL '1 year'
  AND time < now()
GROUP BY bucket, city_name
ORDER BY bucket DESC;

-- Continuous aggregates
-- Define view
CREATE MATERIALIZED VIEW weather_metrics_daily
    WITH (timescaledb.continuous)
AS
SELECT time_bucket('1 day', time) as bucket,
       city_name,
       avg(temp_c)                as avg_temp,
       avg(feels_like_c)          as feels_like_temp,
       max(temp_c)                as max_temp,
       min(temp_c)                as min_temp,
       avg(pressure_hpa)          as pressure,
       avg(humidity_percent)      as humidity_percent,
       avg(rain_3h_mm)            as rain_3h,
       avg(snow_3h_mm)            as snow_3h,
       avg(wind_speed_ms)         as wind_speed,
       avg(clouds_percent)        as clouds
FROM weather_metrics
GROUP BY bucket, city_name
WITH NO DATA;

-- See info about continuous aggregates
SELECT *
FROM timescaledb_information.continuous_aggregates;

-- manual refresh
-- refresh data between 1 Jan 2010 and 2021
CALL refresh_continuous_aggregate('weather_metrics_daily', '2010-01-01', '2021-01-01');

-- Show that manual refresh worked
SELECT *
FROM weather_metrics_daily
WHERE bucket > '2009-01-01'
ORDER BY bucket ASC;

-- create policy
-- refresh the last 6 months of data every 2 weeks
SELECT add_continuous_aggregate_policy('weather_metrics_daily',
                                       start_offset => INTERVAL '6 months',
                                       end_offset => INTERVAL '1 hour',
                                       schedule_interval => INTERVAL '14 days');

-- Continuous aggregate query example
-- Temperature in New York 2015-2021
SELECT bucket, max_temp, avg_temp, min_temp
FROM weather_metrics_daily
WHERE bucket >= '2015-01-01'
  AND bucket < '2021-01-01'
  AND city_name LIKE 'New York'
ORDER BY bucket ASC;

-- Real-time aggregation
SELECT *
FROM weather_metrics_daily
WHERE bucket > now() - 2 * INTERVAL '1 year'
ORDER BY bucket DESC;

-- Enable compression
ALTER TABLE weather_metrics
    SET (
        timescaledb.compress,
        timescaledb.compress_segmentby = 'city_name'
        );

-- See info about compression
SELECT *
FROM timescaledb_information.compression_settings;

-- Add compression policy
SELECT add_compression_policy('weather_metrics', INTERVAL '10 years');

-- Informational view for policy details
SELECT *
FROM timescaledb_information.jobs;

-- Informational view for stats from run jobs
SELECT *
FROM timescaledb_information.job_stats;

---------------------------------------------------
-- Manual compression
---------------------------------------------------
SELECT compress_chunk(i)
FROM show_chunks('weather_metrics', older_than => INTERVAL ' 10 years') i;

-- See effect of compression
SELECT pg_size_pretty(before_compression_total_bytes) as "before compression",
       pg_size_pretty(after_compression_total_bytes)  as "after compression"
FROM hypertable_compression_stats('weather_metrics');

-- Deep and narrow query on compressed data
SELECT avg(temp_c)
FROM weather_metrics
WHERE city_name LIKE 'New York'
  AND time < '2010-01-01';

-- Data retention policy
-- Drop data older than 25 years
SELECT add_retention_policy('weather_metrics', INTERVAL '25 years');

-- Informational view for policy details
SELECT *
FROM timescaledb_information.jobs;
-- Informational view for stats from run jobs
SELECT *
FROM timescaledb_information.job_stats;

-- Manual data retention
SELECT drop_chunks('weather_metrics', INTERVAL '25 years');

