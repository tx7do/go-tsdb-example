CREATE TABLE sensor_data
(
    temperature Float32,
    humidity    Float32,
    volume      Float32,
    PM10        Float32,
    pm25        Float32,
    SO2         Float32,
    NO2         Float32,
    CO          Float32,
    sensor_id   String,
    area        Int16,
    coll_time   DateTime,
    coll_date   Date
) engine = Log;
