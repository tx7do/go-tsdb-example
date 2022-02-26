# ClickHouse

ClickHouse 是一款由俄罗斯 Yandex 公司开发的 C++ 开源高性能 OLAP 组件，并于2016年6月宣布开源。在 Yandex 内部, ClickHouse 主要用于在线流量分析产品 Yandex Metrica，类似于 Google Analytics 或者百度统计。

## Docker部署开发环境

```shell
docker pull yandex/clickhouse-server:latest

# 8123为http接口 9000为tcp接口 9004为mysql接口
# 推荐使用DBeaver作为客户端
docker run -d \
    --name clickhouse-server-test \
    -p 8123:8123 \
    -p 9000:9000 \
    -p 9004:9004 \
    --ulimit \
    nofile=262144:262144 \
    yandex/clickhouse-server:latest
```

## 参考网站

- [官方网站](https://clickhouse.com/)
- [官方文档](https://clickhouse.com/docs/zh/)
- [Github代码库](https://github.com/ClickHouse/ClickHouse)
- [基于Clickhouse实现时序数据OLAP分析](http://events.jianshu.io/p/27ba5246df88)
- [如何用ClickHouse存储时序数据](https://www.yisu.com/zixun/528760.html)
- [还在用 ES 查日志吗，快看看石墨文档 Clickhouse 日志架构玩法](https://www.infoq.cn/article/u3z3dQubLIgxTKsgFCQC)
- [What is ClickHouse, how does it compare to PostgreSQL and TimescaleDB](https://www.timescale.com/blog/what-is-clickhouse-how-does-it-compare-to-postgresql-and-timescaledb-and-how-does-it-perform-for-time-series-data/)
- [如何使用ClickHouse实现时序数据管理和挖掘？](https://segmentfault.com/a/1190000038850846)
- [京东ClickHouse的实践之路](https://www.modb.pro/db/223781)
- [ClickHouse 在有赞的实践之路](https://tech.youzan.com/clickhouse-zai-you-zan-de-shi-jian-zhi-lu/)
- [ClickHouse物化视图在微信的实战经验](https://www.modb.pro/db/70716)
- [clickhouse如何玩转时序数据](https://community.qingcloud.com/assets/uploads/files/1553759734256-clickhouse%E5%A6%82%E4%BD%95%E7%8E%A9%E8%BD%AC%E6%97%B6%E5%BA%8F%E6%95%B0%E6%8D%AE-%E5%BC%A0%E5%81%A5.pdf)
- [ClickHouse 在京东能源管理平台的应用实践](https://www.infoq.cn/article/xzmwwdrvgdviy0qqsbtu)
- [ClickHouse表引擎到底怎么选](https://developer.aliyun.com/article/762461)
- [EMQX + ClickHouse 实现物联网数据接入与分析](https://www.emqx.com/zh/blog/emqx-and-clickhouse-for-iot-data-access-and-analysis)
- [ClickHouse 布道者郭炜：讨论ClickHouse的人需要了解它的设计理念](https://www.163.com/dy/article/GQPVTP5K0511D3QS.html)
- [解读clickhouse存算分离在华为云实践](https://juejin.cn/post/7018065345474723853)
- [基于Flink+ClickHouse打造轻量级点击流实时数仓](https://juejin.cn/post/6883745613255540749)
- [每天十亿级数据更新，秒出查询结果，ClickHouse在携程酒店的应用](https://juejin.cn/post/6844903875309207566)
