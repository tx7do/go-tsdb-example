# QuestDB

## Docker部署开发环境

```shell
docker pull questdb/questdb:latest

docker run -d \
    --name questdb-test \
    -p 9000:9000 \
    -p 8812:8812 \
    -p 9009:9009 \
    questdb/questdb:latest
```

## 参考网站

- [官方网站](https://questdb.io/)
- [Github代码库](https://github.com/questdb/questdb)
