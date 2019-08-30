#!/bin/bash
set -ux

source scripts/.variables.sh

#定义变量
GOPROXY=${GOPROXY:-"https://goproxy.io"}
soft_dir=${HOME:-"/tmp"}
go_version=${go_version:-"1.11"}
protoc_version=${protoc_version:-"3.6.1"}
protoc_include_path=${protoc_include_path:-"${soft_dir}/protoc-${protoc_version}-osx-x86_64/include"}
cloc_version=${cloc_version:-"1.76"}
GOPATH=${GOPATH:-${HOME}"/go_path"}
cmd_path=${cmd_path:-"${GOPATH}/bin"}


#go
go_install(){
        hash go 2>/dev/null || `
		echo "安装golang环境 go"${go_version} && \
		mkdir -p ${soft_dir} && cd ${soft_dir} && \
		wget -c https://dl.google.com/go/go${go_version}.linux-amd64.tar.gz && \
		tar -xzvf go${go_version}.linux-amd64.tar.gz && \
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
		cp ${soft_dir}/cloc-${cloc_version}/cloc ${cmd_path}/cloc || { echo "cloc文件已经存在"; } && \
		echo "cloc 的版本是:" && cloc --version `

}

protoc_install(){
    hash protoc 2>/dev/null || `
    	echo "安装protobuf工具 " && \
		mkdir -p ${soft_dir} && cd  ${soft_dir} && \
		wget -c https://github.com/protocolbuffers/protobuf/releases/download/v${protoc_version}/protoc-${protoc_version}-linux-x86_64.zip && \
		unzip protoc-${protoc_version}-linux-x86_64.zip -d ./protoc-${protoc_version}-linux-x86_64 && \
		mv ${soft_dir}/protoc-${protoc_version}-linux-x86_64/bin/protoc ${cmd_path} && \
		echo "路径为:"${cmd_path} && \
		echo "protoc 的版本是:" && ${cmd_path}/protoc --version`
}

go_plug(){
        cd ${GOPATH}
		echo "安装 protobuf golang插件 protoc-gen-go protoc-gen-grpc-gateway protoc-gen-swagger protoc-go-inject-tag"
		echo "大概耗时30分钟"
		go get  github.com/golang/protobuf/proto
		go get  -u github.com/golang/protobuf/protoc-gen-go
		go get  -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
		go get  -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
		go get  github.com/favadi/protoc-go-inject-tag
		echo "安装gocyclo圈复杂度计算工具"
		go get  github.com/fzipp/gocyclo
		echo "安装打包静态文件工具"
		go get  github.com/rakyll/statik
		echo "安装go-torch"
		go get github.com/uber/go-torch
		cd ${GOPATH}/src/github.com/uber/go-torch
		`git clone https://github.com/brendangregg/FlameGraph.git` || { echo "FlameGraph已经存在"; }
}


go_install
cloc_install
protoc_install
go_plug
