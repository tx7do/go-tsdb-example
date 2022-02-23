# TimeScaleDB

## Docker部署开发环境

```shell
docker pull timescale/timescaledb:latest-pg14
docker pull timescale/timescaledb-postgis:latest-pg13
docker pull timescale/pg_prometheus:latest-pg11

docker run -d \
    --name timescale-test \
    -p 5432:5432 \
    -e POSTGRES_PASSWORD=123456 \
    timescale/timescaledb-postgis:latest-pg13
```

## 参考网站

- [官方网站](https://www.timescale.com/)
- [官方文档](https://docs.timescale.com/timescaledb/latest/tutorials/)
- [Github代码库](https://github.com/timescale/timescaledb)
- [Five Data Models for IoT: Managing the Latest IoT Events Based on a State in Apache Cassandra](https://jaxenter.com/apache-cassandra-iot-174970.html)
- [Cassandra for Internet of Things: An Experimental Evaluation](https://www.scitepress.org/papers/2016/58464/58464.pdf)
- [Advanced Time Series Data Modelling](https://www.datastax.com/blog/advanced-time-series-data-modelling)
- [Advanced Time Series with Cassandra](https://www.datastax.com/blog/advanced-time-series-cassandra)
- [Sensor Data Modeling](https://www.datastax.com/learn/data-modeling-by-example/sensor-data-model)
- [Time Series Data Modeling](https://www.datastax.com/learn/data-modeling-by-example/time-series-model)
- [PostgreSQLの時系列データ向け拡張「TimescaleDB」を触ってみた](https://qiita.com/anzai323/items/68d29ea47192bd18cb3a)
