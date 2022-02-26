-- 创建时序表
CREATE TABLE IF NOT EXISTS ts_kv
(
    dt        Date CODEC (ZSTD),
    ts        DateTime CODEC (LZ4HC),
    entity_id Int64,
    key       Int32,
    bool_v    Nullable(UInt8),
    str_v     Nullable(String),
    long_v    Nullable(Int64),
    dbl_v     Nullable(Float32),
    json_v    Nullable(String)
) ENGINE = Log;

-- 创建字典表
CREATE DICTIONARY IF NOT EXISTS ts_kv_dictionary
    (
    key String,
    key_id Int64
    ) PRIMARY KEY key_id;

-- 创建最新数据表
CREATE TABLE IF NOT EXISTS ts_kv_latest
(
    dt        Date CODEC (ZSTD),
    ts        DateTime CODEC (LZ4HC),
    entity_id Int64,
    key       Int32,
    bool_v    Nullable(UInt8),
    str_v     Nullable(String),
    long_v    Nullable(Int64),
    dbl_v     Nullable(Float32),
    json_v    Nullable(String)
) ENGINE = Log;
