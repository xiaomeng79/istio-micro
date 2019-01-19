#!/bin/bash

#定义变量
GOPROXY=https://goproxy.io
soft_dir=/tmp
go_version=1.11
protoc_version=3.6.1
cloc_version=1.76
cmd_path=/usr/bin

set -e

#go
go_install(){
        hash go 2>/dev/null || `
		echo "安装golang环境 go"${go_version} && \
		mkdir -p ${soft_dir} && cd ${soft_dir} && \
		wget -c https://dl.google.com/go/go${go_version}.linux-amd64.tar.gz && \
		tar -xzvf go${go_version}.linux-amd64.tar.gz && \
		ln -s ${soft_dir}/go/bin/go ${cmd_path}/go && \
		ln -s ${soft_dir}/go/bin/gofmt ${cmd_path}/gofmt && \
		ln -s ${soft_dir}/go/bin/godoc ${cmd_path}/godoc && \
		go version `
}

#圈复杂分析
cloc_install(){
#安装cloc
	hash cloc 2>/dev/null || `
		echo "安装代码统计工具 cloc" && \
		mkdir -p ${soft_dir} && cd  ${soft_dir} && \
		wget -c https://github.com/AlDanial/cloc/archive/v${cloc_version}.zip && \
		unzip v${cloc_version}.zip && \
		ln -s ${soft_dir}/cloc-${cloc_version}/cloc ${cmd_path}/cloc && \
		echo "cloc 的版本是:" && cloc --version `

}

go_plug(){
		echo "安装 protobuf golang插件 protoc-gen-go"
		go get -u github.com/golang/protobuf/proto
		go get -u github.com/golang/protobuf/protoc-gen-go
		echo "安装gocyclo圈复杂度计算工具"
		go get -u github.com/fzipp/gocyclo
		echo "安装go-torch"
		go get github.com/uber/go-torch
		cd ${GOPATH}/src/github.com/uber/go-torch
		git clone https://github.com/brendangregg/FlameGraph.git
		echo "安装打包静态文件工具"
		go get -u github.com/rakyll/statik
}



















#protoc
#protoc_install(){
#	hash protoc2 2>/dev/null || `
#		echo "安装protobuf 代码生成工具 protoc" && \
#		mkdir -p {soft_dir} && cd  {soft_dir} && \
#		wget -c https://github.com/google/protobuf/releases/download/v${protoc_version}/protobuf-cpp-${protoc_version}.tar.gz && \
#		tar -xzvf protobuf-cpp-${protoc_version}.tar.gz && \
#		cd protobuf-${protoc_version} && \
#		./configure --prefix=${soft_dir}/protobuf && \
#		make -j8 && \
#		make install && \
#		ln -s ${soft_dir}/protobuf/bin/protoc ${cmd_path}/protoc && \
#		protoc --version`
#
#	hash protoc-gen-go 2>/dev/null || `
#		echo "安装 protobuf golang 插件 protoc-gen-go" && \
#		go get -u github.com/golang/protobuf/proto&& \
#		go get -u github.com/golang/protobuf/protoc-gen-go `
#
#}