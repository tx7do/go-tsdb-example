-- 创建TimeScaleDB扩展
CREATE EXTENSION IF NOT EXISTS timescaledb CASCADE;
-- 创建PostGIS扩展
CREATE EXTENSION IF NOT EXISTS postgis CASCADE;

-- 创建时序表
CREATE TABLE IF NOT EXISTS ts_kv
(
    entity_id uuid   NOT NULL,
    key       int    NOT NULL,
    ts        bigint NOT NULL,
    bool_v    boolean,
    str_v     varchar(10000000),
    long_v    bigint,
    dbl_v     double precision,
    json_v    json,
    CONSTRAINT ts_kv_pkey PRIMARY KEY (entity_id, key, ts)
);
-- 创建超表
-- SELECT create_hypertable('ts_kv', 'ts');
SELECT create_hypertable('ts_kv', 'ts', chunk_time_interval => 86400000);
-- SELECT create_hypertable('ts_kv', 'ts', chunk_time_interval => INTERVAL '1 day');
-- SELECT add_dimension('ts_kv', 'entity_id', number_partitions => 2);

-- 创建字典表
CREATE TABLE IF NOT EXISTS ts_kv_dictionary
(
    key    varchar(255) NOT NULL,
    key_id serial UNIQUE,
    CONSTRAINT ts_key_id_pkey PRIMARY KEY (key)
);

-- 创建最新数据表
CREATE TABLE IF NOT EXISTS ts_kv_latest
(
    entity_id uuid   NOT NULL,
    key       int    NOT NULL,
    ts        bigint NOT NULL,
    bool_v    boolean,
    str_v     varchar(10000000),
    long_v    bigint,
    dbl_v     double precision,
    json_v    json,
    CONSTRAINT ts_kv_latest_pkey PRIMARY KEY (entity_id, key)
);

CREATE OR REPLACE FUNCTION to_uuid(IN entity_id varchar, OUT uuid_id uuid) AS
$$
BEGIN
    uuid_id := substring(entity_id, 8, 8) || '-' || substring(entity_id, 4, 4) || '-1' || substring(entity_id, 1, 3) ||
               '-' || substring(entity_id, 16, 4) || '-' || substring(entity_id, 20, 12);
END;
$$ LANGUAGE plpgsql;
