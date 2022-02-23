# Cassandra

## Docker部署开发环境

```shell
docker pull bitnami/cassandra:latest

docker run -itd \
    --name cassandra-test \
    -p 7000:7000  \
    -p 9042:9042  \
    -e CASSANDRA_USER=cassandra \
    -e CASSANDRA_PASSWORD=cassandra \
    -e CASSANDRA_ENABLE_USER_DEFINED_FUNCTIONS=true \
    bitnami/cassandra:latest
```

## 参考网站

- [官方网站](https://cassandra.apache.org/_/index.html)
- [Github代码库](https://github.com/apache/cassandra)
- [Five Data Models for IoT: Managing the Latest IoT Events Based on a State in Apache Cassandra](https://jaxenter.com/apache-cassandra-iot-174970.html)
