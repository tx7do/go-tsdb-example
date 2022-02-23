# MongoDB

## Docker部署开发环境

```shell
docker pull bitnami/mongodb:latest

docker run -itd \
    --name mongodb-test \
    -p 27017:27017 \
    -e MONGODB_ROOT_USER=root \
    -e MONGODB_ROOT_PASSWORD=123456 \
    -e MONGODB_USERNAME=test \
    -e MONGODB_PASSWORD=123456 \
    -e MONGODB_DATABASE=test \
    bitnami/mongodb:latest
```

## 参考网站

- [官方网站](https://www.mongodb.com/)
- [Github代码库](https://github.com/mongodb/mongo)
