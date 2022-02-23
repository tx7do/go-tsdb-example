# OpenTSDB

## Docker部署开发环境

```shell
docker pull petergrace/opentsdb-docker:latest

docker run -d \
    --name opentsdb-test \
    -p 4242:4242 \
    petergrace/opentsdb-docker:latest

# 管理后台 http://localhost:4242
```

## 参考网站

- [官方网站](http://opentsdb.net/)
- [Github代码库](https://github.com/OpenTSDB/opentsdb)
