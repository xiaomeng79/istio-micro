## istio-micro
[![Build Status](https://travis-ci.org/xiaomeng79/istio-micro.svg?branch=master)](https://travis-ci.org/xiaomeng79/istio-micro)

### 介绍

通过一个前后台都可以操作的用户接口,对用户服务进行操作
这是一个使用服务网格(istio)构建微服务的使用示例

### 技术栈

|技术|描述|
|---|---|
|grpc+protobuf|服务层之间的通讯|
|echo|应用层接口暴露|
|mysql|存储层|
|redis|缓存层|
|kafka|服务之间异步通讯|
|jaeger|链路跟踪|
|EFK|日志收集存储查询(没涉及,只把日志打到文件)[go-log](https://github.com/xiaomeng79/go-log)|
|statik|静态文件打包|
|docker-compose|容器部署|
|istio|流量控制,服务降级,跟踪,服务发现,分流等|

### 模块

- api_backend 后台操作用户数据的RESTful接口
- api_frontend 前台查询用户的接口
- srv_user 用户服务
- srv_socket 推送服务

### 快速演示

### 安装流程

- 依赖安装

go >=1.11
docker-compose


```go
		go get -u github.com/golang/protobuf/proto
		go get -u github.com/golang/protobuf/protoc-gen-go
		go get -u github.com/rakyll/statik
```

- 下载代码

```go
https://github.com/xiaomeng79/istio-micro.git

```

- 编译代码

```go
make allbuild
```

- 运行代码

```go
make compose up
```


#### 测试

1. 浏览器打开消息推送窗口`http://127.0.0.1:5002/public/`

2. 打开命令行插入mysql一条数据

```go
curl -X POST \
  http://127.0.0.1:8888/backend/v1/user \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/json' \
  -d '{"user_name":"meng","iphone":"18201420251","sex":1,"password":"123456"}'
```

3. 查看消息推送窗口是否有变化




