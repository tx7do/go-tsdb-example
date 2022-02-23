# go-tsdb-example

本项目主要是用来测试各种不同时序数据库(TSDB)的.

## 时序数据简介

时序数据全称是时间序列（TimeSeries）数据，是按照时间顺序索引的一系列数据点。最常见的是在连续的等时间间隔时间点上获取的序列，因此，它是一系列离散数据[1]。

时序数据几乎无处不在，在目前单向的时间流中，人的脉搏、空气的湿度、股票的价格等都随着时间的流逝不断变化。时序数据是数据的一种，因为它显著而有价值的特点，成为我们特别分析的对象。

将时序数据可以建模为如下部分组成：

- **Metric**：度量的数据集，类似于关系型数据库中的 table，是固定属性，一般不随时间而变化
- **Timestamp**：时间戳，表征采集到数据的时间点
- **Tags**：维度列，用于描述Metric，代表数据的归属、属性，表明是哪个设备/模块产生的，一般不随着时间变化
- **Field/Value**：指标列，代表数据的测量值，可以是单值也可以是多值

## 时序数据库
- [x] [InfluxDB](db/influxdb/README.md)
- [x] [ClickHouse](db/clickhouse/README.md)
- [x] [TimescaleDB](db/timescaledb/README.md)
- [x] [ElasticSearch](db/es/README.md)
- [x] [MongoDB](db/mongodb/README.md)
- [x] [OpenTSDB](db/opentsdb/README.md)
- [x] [QuestDB](db/questdb/README.md)
- [x] [TDEngine](db/tdengine/README.md)
- [x] [Cassandra](db/cassandra/README.md)
- [ ] [CrateDB](db/cratedb/README.md)
- [ ] [SiriDB](db/siridb/README.md)
- [ ] [Timestream](db/timestream/README.md)
- [ ] [VictoriaMetrics](db/victoriametrics/README.md)

## 参考资料

- [Time series - Wikipedia](https://en.wikipedia.org/wiki/Time_series)
- [DB-Engines Ranking of Time Series DBMS](https://db-engines.com/en/ranking/time+series+dbms)
- [Time Series Benchmark Suite (TSBS)](https://github.com/timescale/tsbs)
