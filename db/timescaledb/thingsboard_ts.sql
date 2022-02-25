-- 创建TimeScaleDB扩展
CREATE EXTENSION IF NOT EXISTS timescaledb CASCADE;
-- 创建PostGIS扩展
CREATE EXTENSION IF NOT EXISTS postgis CASCADE;

-- 创建属性键值表
CREATE TABLE IF NOT EXISTS attribute_kv
(
    entity_type    varchar(255),
    entity_id      uuid,
    attribute_type varchar(255),
    attribute_key  varchar(255),
    bool_v         boolean,
    str_v          varchar(10000000),
    long_v         bigint,
    dbl_v          double precision,
    json_v         json,
    last_update_ts bigint,
    CONSTRAINT attribute_kv_pkey PRIMARY KEY (entity_type, entity_id, attribute_type, attribute_key)
);
CREATE INDEX IF NOT EXISTS idx_attribute_kv_by_key_and_last_update_ts ON attribute_kv (entity_id, attribute_key, last_update_ts desc);

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

-- 实体ID -> UUID
CREATE OR REPLACE FUNCTION to_uuid(IN entity_id varchar, OUT uuid_id uuid) AS
$$
BEGIN
    uuid_id := substring(entity_id, 8, 8) || '-' || substring(entity_id, 4, 4) || '-1' || substring(entity_id, 1, 3) ||
               '-' || substring(entity_id, 16, 4) || '-' || substring(entity_id, 20, 12);
END;
$$ LANGUAGE plpgsql;


/**
  * 以下为查询语句
  */

-- 插入一条数据到最新遥测数据表去
INSERT INTO ts_kv_latest (entity_id, key, ts, bool_v, str_v, long_v, dbl_v, json_v)
VALUES ($1, $2, $3, $4, $5, $6, $7, cast($8 AS json))
ON CONFLICT (entity_id, key) DO UPDATE SET ts     = $3,
                                           bool_v = $4,
                                           str_v  = $5,
                                           long_v = $6,
                                           dbl_v  = $7,
                                           json_v = cast($8 AS json);

-- 插入一条数据到遥测数据历史数据表去
INSERT INTO ts_kv (entity_id, key, ts, bool_v, str_v, long_v, dbl_v, json_v)
VALUES ($1, $2, $3, $4, $5, $6, $7, cast($8 AS json))
ON CONFLICT (entity_id, key, ts) DO UPDATE SET bool_v = $4,
                                               str_v  = $5,
                                               long_v = $6,
                                               dbl_v  = $7,
                                               json_v = cast($8 AS json);

-- 插入一条数据到 数据键名-键ID映射表
INSERT INTO ts_kv_dictionary (key)
VALUES ($1);

-- 查询键名的键ID
SELECT key_id
FROM ts_kv_dictionary
WHERE ts_kv_dictionary.key = $1;

-- 查询指定设备的最新遥测数据
SELECT ts_kv_latest.entity_id AS entityId
     , ts_kv_latest.key       AS key
     , ts_kv_dictionary.key   AS strKey
     , ts_kv_latest.str_v     AS strValue
     , ts_kv_latest.bool_v    AS boolValue
     , ts_kv_latest.long_v    AS longValue
     , ts_kv_latest.dbl_v     AS doubleValue
     , ts_kv_latest.json_v    AS jsonValue
     , ts_kv_latest.ts        AS ts
FROM ts_kv_latest
         INNER JOIN ts_kv_dictionary ON ts_kv_latest.key = ts_kv_dictionary.key_id
WHERE ts_kv_latest.entity_id = cast('ad2bfe60-7514-11ec-9a90-af0223be0666' AS uuid);

-- 查询指定设备的所有遥测数据键名
SELECT DISTINCT ts_kv_dictionary.key AS strKey
FROM ts_kv_latest
         INNER JOIN ts_kv_dictionary ON ts_kv_latest.key = ts_kv_dictionary.key_id
WHERE ts_kv_latest.entity_id IN ('ad2bfe60-7514-11ec-9a90-af0223be0666')
ORDER BY ts_kv_dictionary.key;

-- 查询指定租户的所有设备的遥测数据键名
SELECT DISTINCT ts_kv_dictionary.key AS strKey
FROM ts_kv_latest
         INNER JOIN ts_kv_dictionary ON ts_kv_latest.key = ts_kv_dictionary.key_id
WHERE ts_kv_latest.entity_id IN (SELECT id FROM device WHERE tenant_id = 1 limit 100)
ORDER BY ts_kv_dictionary.key;

-- 查询指定租户的指定设备类型的遥测数据键名
SELECT DISTINCT ts_kv_dictionary.key AS strKey
FROM ts_kv_latest
         INNER JOIN ts_kv_dictionary ON ts_kv_latest.key = ts_kv_dictionary.key_id
WHERE ts_kv_latest.entity_id IN
      (SELECT id FROM device WHERE device_profile_id = :device_profile_id AND tenant_id = :tenant_id limit 100)
ORDER BY ts_kv_dictionary.key;

-- 有条件的查询实体的遥测历史数据
SELECT tskv
FROM ts_kv tskv
WHERE tskv.entity_id = :entityId
  AND tskv.key = :entityKey
  AND tskv.ts >= :startTs
  AND tskv.ts < :endTs;

-- 有条件的删除实体的遥测历史数据
DELETE
FROM ts_kv tskv
WHERE tskv.entity_id = :entityId
  AND tskv.key = :entityKey
  AND tskv.ts >= :startTs
  AND tskv.ts < :endTs;

-- TimeScaleDB 聚合查询时序数据 - 平均值
SELECT time_bucket(:timeBucket, tskv.ts, :startTs)          AS tsBucket,
       :timeBucket                                          AS interval,
       SUM(COALESCE(tskv.long_v, 0))                        AS longValue,
       SUM(COALESCE(tskv.dbl_v, 0.0))                       AS doubleValue,
       SUM(CASE WHEN tskv.long_v IS NULL THEN 0 ELSE 1 END) AS longCountValue,
       SUM(CASE WHEN tskv.dbl_v IS NULL THEN 0 ELSE 1 END)  AS doubleCountValue,
       NULL                                                 AS strValue,
       'AVG'                                                AS aggType
FROM ts_kv tskv
WHERE tskv.entity_id = CAST(:entityId AS uuid)
  AND tskv.key = CAST(:entityKey AS int)
  AND tskv.ts > :startTs
  AND tskv.ts <= :endTs
GROUP BY tskv.entity_id, tskv.key, tsBucket
ORDER BY tskv.entity_id, tskv.key, tsBucket;

-- TimeScaleDB 聚合查询时序数据 - 最大值
SELECT time_bucket(:timeBucket, tskv.ts, :startTs)          AS tsBucket,
       :timeBucket                                          AS interval,
       MAX(COALESCE(tskv.long_v, -9223372036854775807))     AS longValue,
       MAX(COALESCE(tskv.dbl_v, -1.79769E+308))             AS doubleValue,
       SUM(CASE WHEN tskv.long_v IS NULL THEN 0 ELSE 1 END) AS longCountValue,
       SUM(CASE WHEN tskv.dbl_v IS NULL THEN 0 ELSE 1 END)  AS doubleCountValue,
       MAX(tskv.str_v)                                      AS strValue,
       'MAX'                                                AS aggType
FROM ts_kv tskv
WHERE tskv.entity_id = CAST(:entityId AS uuid)
  AND tskv.key = CAST(:entityKey AS int)
  AND tskv.ts > :startTs
  AND tskv.ts <= :endTs
GROUP BY tskv.entity_id, tskv.key, tsBucket
ORDER BY tskv.entity_id, tskv.key, tsBucket;

-- TimeScaleDB 聚合查询时序数据 - 最小值
SELECT time_bucket(:timeBucket, tskv.ts, :startTs)          AS tsBucket,
       :timeBucket                                          AS interval,
       MIN(COALESCE(tskv.long_v, 9223372036854775807))      AS longValue,
       MIN(COALESCE(tskv.dbl_v, 1.79769E+308))              as doubleValue,
       SUM(CASE WHEN tskv.long_v IS NULL THEN 0 ELSE 1 END) AS longCountValue,
       SUM(CASE WHEN tskv.dbl_v IS NULL THEN 0 ELSE 1 END)  AS doubleCountValue,
       MIN(tskv.str_v)                                      AS strValue,
       'MIN'                                                AS aggType
FROM ts_kv tskv
WHERE tskv.entity_id = CAST(:entityId AS uuid)
  AND tskv.key = CAST(:entityKey AS int)
  AND tskv.ts > :startTs
  AND tskv.ts <= :endTs
GROUP BY tskv.entity_id, tskv.key, tsBucket
ORDER BY tskv.entity_id, tskv.key, tsBucket;

-- TimeScaleDB 聚合查询时序数据 - 总和
SELECT time_bucket(:timeBucket, tskv.ts, :startTs)          AS tsBucket,
       :timeBucket                                          AS interval,
       SUM(COALESCE(tskv.long_v, 0))                        AS longValue,
       SUM(COALESCE(tskv.dbl_v, 0.0))                       AS doubleValue,
       SUM(CASE WHEN tskv.long_v IS NULL THEN 0 ELSE 1 END) AS longCountValue,
       SUM(CASE WHEN tskv.dbl_v IS NULL THEN 0 ELSE 1 END)  AS doubleCountValue,
       NULL                                                 AS strValue,
       NULL                                                 AS jsonValue,
       'SUM'                                                AS aggType
FROM ts_kv tskv
WHERE tskv.entity_id = CAST(:entityId AS uuid)
  AND tskv.key = CAST(:entityKey AS int)
  AND tskv.ts > :startTs
  AND tskv.ts <= :endTs
GROUP BY tskv.entity_id, tskv.key, tsBucket
ORDER BY tskv.entity_id, tskv.key, tsBucket;

-- TimeScaleDB 聚合查询时序数据 - 数量
SELECT time_bucket(:timeBucket, tskv.ts, :startTs)          AS tsBucket,
       :timeBucket                                          AS interval,
       SUM(CASE WHEN tskv.bool_v IS NULL THEN 0 ELSE 1 END) AS booleanValueCount,
       SUM(CASE WHEN tskv.str_v IS NULL THEN 0 ELSE 1 END)  AS strValueCount,
       SUM(CASE WHEN tskv.long_v IS NULL THEN 0 ELSE 1 END) AS longValueCount,
       SUM(CASE WHEN tskv.dbl_v IS NULL THEN 0 ELSE 1 END)  AS doubleValueCount,
       SUM(CASE WHEN tskv.json_v IS NULL THEN 0 ELSE 1 END) AS jsonValueCount
FROM ts_kv tskv
WHERE tskv.entity_id = CAST(:entityId AS uuid)
  AND tskv.key = CAST(:entityKey AS int)
  AND tskv.ts > :startTs
  AND tskv.ts <= :endTs
GROUP BY tskv.entity_id, tskv.key, tsBucket
ORDER BY tskv.entity_id, tskv.key, tsBucket;

-- PostgreDB 聚合查询时序数据 - 字符串最大值
SELECT MAX(tskv.str_v) AS strValue
FROM ts_kv tskv
WHERE tskv.str_v IS NOT NULL
  AND tskv.entity_id = CAST(:entityId AS uuid)
  AND tskv.key = CAST(:entityKey AS int)
  AND tskv.ts >= :startTs
  AND tskv.ts < :endTs;

-- PostgreDB 聚合查询时序数据 - 字符串最小值
SELECT MIN(tskv.str_v) AS strValue
FROM ts_kv tskv
WHERE tskv.str_v IS NOT NULL
  AND tskv.entity_id = CAST(:entityId AS uuid)
  AND tskv.key = CAST(:entityKey AS int)
  AND tskv.ts >= :startTs
  AND tskv.ts < :endTs;

-- PostgreDB 聚合查询时序数据 - 数值最大值
SELECT MAX(COALESCE(tskv.long_v, -9223372036854775807))     AS longValue,
       MAX(COALESCE(tskv.dbl_v, -1.79769E+308))             AS doubleValue,
       SUM(CASE WHEN tskv.long_v IS NULL THEN 0 ELSE 1 END) AS longCountValue,
       SUM(CASE WHEN tskv.dbl_v IS NULL THEN 0 ELSE 1 END)  AS doubleCountValue,
       'MAX'                                                AS aggType
FROM ts_kv tskv
WHERE tskv.entity_id = CAST(:entityId AS uuid)
  AND tskv.key = CAST(:entityKey AS int)
  AND tskv.ts >= :startTs
  AND tskv.ts < :endTs;

-- PostgreDB 聚合查询时序数据 - 数值最小值
SELECT MIN(COALESCE(tskv.long_v, -9223372036854775807))     AS longValue,
       MIN(COALESCE(tskv.dbl_v, -1.79769E+308))             AS doubleValue,
       SUM(CASE WHEN tskv.long_v IS NULL THEN 0 ELSE 1 END) AS longCountValue,
       SUM(CASE WHEN tskv.dbl_v IS NULL THEN 0 ELSE 1 END)  AS doubleCountValue,
       'MIN'                                                AS aggType
FROM ts_kv tskv
WHERE tskv.entity_id = CAST(:entityId AS uuid)
  AND tskv.key = CAST(:entityKey AS int)
  AND tskv.ts >= :startTs
  AND tskv.ts < :endTs;

-- PostgreDB 聚合查询时序数据 - 数量
SELECT SUM(CASE WHEN tskv.bool_v IS NULL THEN 0 ELSE 1 END) AS booleanValueCount,
       SUM(CASE WHEN tskv.str_v IS NULL THEN 0 ELSE 1 END)  AS strValueCount,
       SUM(CASE WHEN tskv.long_v IS NULL THEN 0 ELSE 1 END) AS longValueCount,
       SUM(CASE WHEN tskv.dbl_v IS NULL THEN 0 ELSE 1 END)  AS doubleValueCount,
       SUM(CASE WHEN tskv.json_v IS NULL THEN 0 ELSE 1 END) AS jsonValueCount
FROM ts_kv tskv
WHERE tskv.entity_id = CAST(:entityId AS uuid)
  AND tskv.key = CAST(:entityKey AS int)
  AND tskv.ts >= :startTs
  AND tskv.ts < :endTs;

-- PostgreDB 聚合查询时序数据 - 平均值
SELECT SUM(COALESCE(tskv.long_v, 0))                        AS longValue,
       SUM(COALESCE(tskv.dbl_v, 0.0))                       AS doubleValue,
       SUM(CASE WHEN tskv.long_v IS NULL THEN 0 ELSE 1 END) AS longCountValue,
       SUM(CASE WHEN tskv.dbl_v IS NULL THEN 0 ELSE 1 END)  AS doubleCountValue,
       NULL                                                 AS strValue,
       'AVG'                                                AS aggType
FROM ts_kv tskv
WHERE tskv.entity_id = CAST(:entityId AS uuid)
  AND tskv.key = CAST(:entityKey AS int)
  AND tskv.ts >= :startTs
  AND tskv.ts < :endTs;

-- PostgreDB 聚合查询时序数据 - 求和
SELECT SUM(COALESCE(tskv.long_v, 0))                        AS longValue,
       SUM(COALESCE(tskv.dbl_v, 0.0))                       AS doubleValue,
       SUM(CASE WHEN tskv.long_v IS NULL THEN 0 ELSE 1 END) AS longCountValue,
       SUM(CASE WHEN tskv.dbl_v IS NULL THEN 0 ELSE 1 END)  AS doubleCountValue,
       NULL                                                 AS strValue,
       NULL                                                 AS jsonValue,
       'SUM'                                                AS aggType
FROM ts_kv tskv
WHERE tskv.entity_id = CAST(:entityId AS uuid)
  AND tskv.key = CAST(:entityKey AS int)
  AND tskv.ts >= :startTs
  AND tskv.ts < :endTs;
