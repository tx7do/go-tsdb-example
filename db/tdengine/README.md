# TDEngine

## Docker部署开发环境

```shell
docker pull tdengine/tdengine:latest

docker run -d \
    --name tdengine-test \
    -p 6030-6041:6030-6041 \
    -p 6030-6041:6030-6041/udp \
    tdengine/tdengine:latest
```

## 参考网站

- [官方网站](https://tdengine.com/)
- [Github代码库](https://github.com/taosdata/TDengine)
- https://www.taosdata.com/downloads/TDengine_Testing_Report_cn.pdf