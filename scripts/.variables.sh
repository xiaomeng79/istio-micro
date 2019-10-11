#!/bin/bash

#set -u

#获取用户自定义变量(myvariables.sh)
if [ -f "scripts/myvariables.sh" ];then
    source scripts/myvariables.sh
fi
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

mkdir -p ${GOPATH}/bin
mkdir -p ${GOPATH}/src

#将环境变量存入本地环境配置
echo "GOPROXY=${GOPROXY}" >>${HOME}/.profile
echo "protoc_include_path=${protoc_include_path}" >>${HOME}/.profile
echo "GO111MODULE=${GO111MODULE}" >>${HOME}/.profile
echo "GOPATH=${GOPATH}" >>${HOME}/.profile
echo "PATH=${soft_dir}/go/bin:${GOPATH}/bin:${PATH}" >>${HOME}/.profile

#手动执行
#source ~/.profile