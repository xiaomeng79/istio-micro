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
|metric|监控报警(influxdb+grafana)|
|docker-compose|容器部署|
|istio|流量控制,服务降级,跟踪,服务发现,分流等|

### 模块

- api_backend 后台操作用户数据的RESTful接口
- api_frontend 前台查询用户的接口
- srv_user 用户服务
- srv_socket 推送服务
- srv_account 账户服务

### 快速演示(docker-compose)

### 安装流程

1. 安装依赖

- 系统依赖安装
    1. git >= 2.17
    2. wget
    3. make
    4. unzip
    5. tar
    
```bash
apt-get install git wget make unzip tar -y
```
    
- 可选部署安装(**任何一种都可以,也可直接部署二进制文件**)
    1. docker >= 1.13.1
    2. docker-compose >=1.19
    3. k8s >=1.12
    4. istio >=1.1

2. 克隆项目
```bash
git clone https://github.com/xiaomeng79/istio-micro.git
```

3. 
3. 安装运行环境
```bash
cd istio-micro 
 make ver
 source ~/.profile 
 make install
```

4. 编译代码

```bash
sudo make allbuild
```

5. 运行代码

```bash
sudo make compose up
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
也可使用grpc-gateway(网关端口:9998)发送信息
```bash
curl -X POST \
  http://127.0.0.1:9998/user \
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

//编译单个服务,同时添加版本信息
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

//开启代码性能分析(如type为api,project为frontend)
make pprofon type=api project=frontend

//关闭代码性能分析(如type为api,project为frontend)
make pprofoff type=api project=frontend

// 清空编译
make clean
```

#### 命令行

```bash
# 每个执行程序，可以查看版本和提交信息,如：srv_user
./srv_user version

```
#### 增量更新sql

原理:通过在程序上,加上版本号如:`v1.0.1`,在数据库记录上一个更新程序程序版本号如:`v0.1.1`,程序启动会判断更新记录,并将中间的版本号的sql,按照版本号从小
到大排序后,依次执行,执行完成后并更新数据库版本号为最新的版本号

[增量更新sql文件说明](./srv/account/README.md)

#### 监控报警

influxdb提供采集数据存储,grafana提供数据展示,报警

http://127.0.0.1:3000 账号密码:admin

```bash
# 新建数据源(influxdb) 地址:http://influxdb:8086
# 导入度量的信息(deployments/config/metrics/gc.json) 可以查看gc和内存信息

```

#### 代码性能分析(可以线上临时开启分析)

[pprof封装库](./pkg/pprof)

程序运行的时候会生成进程id(默认在运行目录下server.pid),通过kill命令发送一个信号(默认10)到程序**开启**性能分析,kill命令发送一个信号(默认12)到程序**关闭**性能分析


```shell
//比如进程id:3125
kill -10 3125 //开启代码性能分析

go tool pprof http://127.0.0.1:38888/debug/pprof/goroutine //goroutine
go tool pprof http://127.0.0.1:38888/debug/pprof/heap //heap
go tool pprof http://127.0.0.1:38888/debug/pprof/profile //profile

kill -12 3125 //关闭代码性能分析
```

#### 生成文档(swagger)
```bash
# 生成网关和文档
make proto
# 本地文档地址(istio-micro/deployments/config/swagger/srv/user/proto/user.swagger.json)
# 在线文档地址(http://127.0.0.1:9998/swagger/user.swagger.json)
# 可以使用swagger-ui(http://editor.swagger.io/)查看

```

#### k8s部署

- 本地安装测试k8s [minikube](https://github.com/kubernetes/minikube)
- k8s安装kafka [kubernetes-kafka](https://github.com/Yolean/kubernetes-kafka)
- k8s安装redis,mysql [k8s-install-scripts](https://github.com/zhuchuangang/k8s-install-scripts)
- k8s安装jaeger [jaeger-kubernetes](https://github.com/jaegertracing/jaeger-kubernetes) `kubectl create -f https://raw.githubusercontent.com/jaegertracing/jaeger-kubernetes/master/all-in-one/jaeger-all-in-one-template.yml`
```go
kubectl apply -f deployments/k8s/api_backend/dev.yaml
kubectl apply -f deployments/k8s/api_frontend/dev.yaml
kubectl apply -f deployments/k8s/srv_user/dev.yaml
kubectl apply -f deployments/k8s/srv_socket/dev.yaml

```

#### istio流量控制

执行deployments/k8s目录下各个network文件和网关文件

#### TODO

- 完善istio其他配置文件


