## 安装依赖

#### 配置环境变量

GOPROXY=${GOPROXY:-"https://goproxy.io"}
GO111MODULE=${GO111MODULE:-"auto"}
GOPATH=${GOPATH:-${HOME}/"go_path"}
soft_dir=${soft_dir:-${HOME}}
go_version=${go_version:-"1.12.9"}
protoc_version=${protoc_version:-"3.6.1"}
protoc_include_path=${protoc_include_path:-"${soft_dir}/protoc-${protoc_version}-linux-x86_64/include"}
cloc_version=${cloc_version:-"1.76"}
cmd_path=${cmd_path:-"${GOPATH}/bin"}

```bash
echo "GOPROXY=https://goproxy.io" >>${HOME}/.profile
echo "protoc_include_path=${HOME}/protoc/include" >>${HOME}/.profile
echo "GO111MODULE=auto" >>${HOME}/.profile
echo "GOPATH=${HOME}/go_path" >>${HOME}/.profile
echo "PATH=${soft_dir}/go/bin:${GOPATH}/bin:${PATH}" >>${HOME}/.profile
```

#### 安装系统依赖
```bash
    1. git >= 2.17
    2. wget
    3. make
    4. unzip
    5. tar
```
#### 安装go程序
```bash

```