## 安装环境变量配置

**可在scripts下新建安装环境变量配置文件(myvariables.sh)**
```bash
#项目相关的
ProjectName=${ProjectName:-"github.com/xiaomeng79/istio-micro"}
Version=${Version:-"unknow"}
TARGET=${TARGET:-'main'}

#执行环境
GOPROXY=${GOPROXY:-"https://goproxy.cn"}
#go mod是否开启
GO111MODULE=${GO111MODULE:-"auto"}
#GOPATH的路径
GOPATH=${GOPATH:-${HOME}"/com_go"}
#其他软件的安装目录
soft_dir=${soft_dir:-${HOME}}
#go安装的版本
go_version=${go_version:-"1.13.1"}
#protoc的版本
protoc_version=${protoc_version:-"3.6.1"}
#protoc引用的路径
protoc_include_path=${protoc_include_path:-"${soft_dir}/protoc-${protoc_version}-linux-x86_64/include"}
#cloc版本
cloc_version=${cloc_version:-"1.76"}
#执行文件路径
cmd_path=${cmd_path:-"${GOPATH}/bin"}
```