# InfluxDB

## Docker部署开发环境

```shell
docker pull bitnami/influxdb:latest

docker run -d \
    --name influxdb-test \
    -p 8083:8083 \
    -p 8086:8086 \
    -e INFLUXDB_HTTP_AUTH_ENABLED=true \
    -e INFLUXDB_ADMIN_USER=admin \
    -e INFLUXDB_ADMIN_USER_PASSWORD=123456789 \
    -e INFLUXDB_ADMIN_USER_TOKEN=admintoken123 \
    -e INFLUXDB_DB=my_database \
    bitnami/influxdb:latest

 # 管理后台: http://localhost:8086/
```

## 参考网站

- [官方网站](https://www.influxdata.com/)
- [Github代码库](https://github.com/influxdata/influxdb)
- https://www.influxdata.com/blog/influxdb-clustering/