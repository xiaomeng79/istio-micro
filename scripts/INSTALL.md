## 手动安装

1. 安装主要程序
- [go](https://studygolang.com/dl) >= 1.13.1 go
- [cloc](https://github.com/AlDanial/cloc) >=1.76  代码统计
- [protoc](https://github.com/protocolbuffers/protobuf) >= 3.6.1 proto buffer

**注意:安装完protoc后,需要将protoc下的include目录配置到环境变量,如下**

```bash
echo protoc_include_path=你的protoc地址 >> ~/.profile
```

2. 定义GOPATH并设置go的环境变量
```bash
#新建GOPATH并切换到GOPATH
cd ${GOPATH} 
export GOPATH=你的GOPATH路径 
#如果不能访问外网设置代理
export GOPROXY=https://goproxy.io 
#关闭go mod (这样可以将文件安装到GOPATH)
export GO111MODULE=auto  
#将GOPATH下的bin目录放到PATH环境变量下
export PATH=$GOPATH/bin:$PATH
#将以上环境变量配置到~/.profile并执行
source ~/.profile

```

3. 安装go的依赖包
```bash
go get  github.com/golangci/golangci-lint/cmd/golangci-lint
go get  github.com/golang/protobuf/protoc-gen-go
go get  github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get  github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get  github.com/favadi/protoc-go-inject-tag 
go get  github.com/fzipp/gocyclo 
go get  github.com/rakyll/statik
```
4. 切换到istio-micro目录执行
```bash
#编译全部
make allbuild
#启动
make compose

```
