#!/bin/bash

#set -u

#项目相关的
ProjectName=${ProjectName:-"github.com/xiaomeng79/istio-micro"}
Version=${Version:-"unknow"}
TARGET=${TARGET:-'main'}

#执行环境
GOPROXY=${GOPROXY:-"https://goproxy.io"}
GO111MODULE=${GO111MODULE:-"auto"}
GOPATH=${GOPATH:-${HOME}/"go_path"}
soft_dir=${soft_dir:-${HOME}}
go_version=${go_version:-"1.12.9"}
protoc_version=${protoc_version:-"3.6.1"}
protoc_include_path=${protoc_include_path:-"${soft_dir}/protoc-${protoc_version}-linux-x86_64/include"}
cloc_version=${cloc_version:-"1.76"}
cmd_path=${cmd_path:-"${GOPATH}/bin"}

#将环境变量存入本地环境配置
echo "GOPROXY=${GOPROXY}" >>${HOME}/.profile
echo "protoc_include_path=${protoc_include_path}" >>${HOME}/.profile
echo "GO111MODULE=${GO111MODULE}" >>${HOME}/.profile
echo "GOPATH=${GOPATH}" >>${HOME}/.profile
echo "PATH=${soft_dir}/go/bin:${GOPATH}/bin:${PATH}" >>${HOME}/.profile

source ${HOME}/.profile