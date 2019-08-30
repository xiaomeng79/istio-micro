#!/bin/bash

set -eu

#项目相关的
ProjectName=${ProjectName:-"github.com/xiaomeng79/istio-micro"}
Version=${Version:-"unknow"}
TARGET=${TARGET:-'main'}

#执行环境
GOPROXY=${GOPROXY:-"https://goproxy.io"}
GO111MODULE=${GO111MODULE:-"auto"}
GOPATH=${GOPATH:-${HOME}/"go_path"}
soft_dir=${HOME:-"/tmp"}
go_version=${go_version:-"1.11"}
protoc_version=${protoc_version:-"3.6.1"}
protoc_include_path=${protoc_include_path:-"protoc-${protoc_version}-osx-x86_64/include"}
cloc_version=${cloc_version:-"1.76"}
cmd_path=${cmd_path:-"/usr/bin"}