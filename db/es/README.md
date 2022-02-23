# ElasticSearch

## Docker部署开发环境

```shell
docker pull bitnami/elasticsearch:latest

docker run -d \
    --name elasticsearch-test \
    -p 9200:9200 \
    -p 9300:9300 \
    -e ELASTICSEARCH_USERNAME=elastic \
    -e ELASTICSEARCH_PASSWORD=elastic \
    -e xpack.security.enabled=true \
    -e discovery.type=single-node \
    -e http.cors.enabled=true \
    -e http.cors.allow-origin=http://localhost:13580,http://127.0.0.1:13580 \
    -e http.cors.allow-headers=X-Requested-With,X-Auth-Token,Content-Type,Content-Length,Authorization \
    -e http.cors.allow-credentials=true \
    bitnami/elasticsearch:latest
```

## 参考网站

- [官方网站](https://www.elastic.co/)
- [Github代码库](https://github.com/elastic/elasticsearch)
