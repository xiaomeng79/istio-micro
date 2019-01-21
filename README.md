## istio-micro
[![Build Status](https://travis-ci.org/xiaomeng79/istio-micro.svg?branch=master)](https://travis-ci.org/xiaomeng79/istio-micro) [![codecov](https://codecov.io/gh/xiaomeng79/istio-micro/branch/master/graph/badge.svg)](https://codecov.io/gh/xiaomeng79/istio-micro)


#### 使用go-micro构建微服务示例请到一下仓库

[go-example](https://github.com/xiaomeng79/go-example)


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

### 快速演示(docker-compose)

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

#### 目录介绍


|技术|描述|
|---|---|
|api|api接口|
|cinit|配置和初始化文件|
|cmd|程序入口|
|deployments|部署文件(docker,k8s,istio)|
|internal|内部公共文件|
|scripts|脚本文件|
|srv|服务|

#### 自动化

Makefile

```go
//格式化代码
make fmt 

//vendor
make vendor

//代码测试,代码检查
make test

//编译单个服务
make build type=srv project=user

//编译全部服务
make allbuild

//protobuf
make proto

//生成单个dockerfile
make dockerfile type=srv project=user

//生成全部dockerfile
make alldockerfile

//docker-compose部署
make compose up

//打包静态文件
make builddata

//提交代码到远程仓库
make push msg="提交信息"
```

#### k8s部署

- 本地安装测试k8s [minikube](https://github.com/kubernetes/minikube)
- k8s安装kafka [kubernetes-kafka](https://github.com/Yolean/kubernetes-kafka)
- k8s安装redis,mysql [k8s-install-scripts](https://github.com/zhuchuangang/k8s-install-scripts)
```go
kubectl apply -f deployments/k8s/api_backend/dev.yaml
kubectl apply -f deployments/k8s/api_frontend/dev.yaml
kubectl apply -f deployments/k8s/srv_user/dev.yaml
kubectl apply -f deployments/k8s/srv_socket/dev.yaml

```

#### istio部署

**待完善**
在k8s部署的基础上,执行deployments/k8s目录下各个network文件和网关文件

#### TODO

- 完善istio配置文件
- 支持swagger接口文档生成


