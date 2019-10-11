#!/bin/bash

source scripts/.variables.sh
set -xe

#定义变量
GOPROXY=${GOPROXY:-"https://goproxy.io"}
soft_dir=${HOME:-"/tmp"}
go_version=${go_version:-"1.13.1"}
protoc_version=${protoc_version:-"3.6.1"}
protoc_include_path=${protoc_include_path:-"${soft_dir}/protoc-${protoc_version}-linux-x86_64/include"}
cloc_version=${cloc_version:-"1.76"}
GOPATH=${GOPATH:-${HOME}"/go_path"}
cmd_path=${cmd_path:-"${GOPATH}/bin"}

#go
go_install(){
		echo "安装golang环境 go"${go_version} && \
		mkdir -p ${soft_dir} && cd ${soft_dir} && \
		wget -c https://dl.google.com/go/go${go_version}.linux-amd64.tar.gz && \
		tar -xzvf go${go_version}.linux-amd64.tar.gz && \
		go version && \
		go env -w GOPROXY=${GOPROXY},direct && \
		go env -w GOPATH=${GOPATH} && \
		go env -w GO111MODULE=auto
}

#圈复杂分析
cloc_install(){
#安装cloc
		echo "安装代码统计工具 cloc" && \
		mkdir -p ${soft_dir} && cd  ${soft_dir} && \
		wget -c https://github.com/AlDanial/cloc/archive/v${cloc_version}.zip && \
		unzip v${cloc_version}.zip && \
		mv ${soft_dir}/cloc-${cloc_version}/cloc ${cmd_path} || { echo "cloc文件已经存在"; } && \
		echo "cloc 的版本是:" && cloc --version

}

protoc_install(){
    	echo "安装protobuf工具 " && \
		mkdir -p ${soft_dir} && cd  ${soft_dir} && \
		wget -c https://github.com/protocolbuffers/protobuf/releases/download/v${protoc_version}/protoc-${protoc_version}-linux-x86_64.zip && \
		unzip protoc-${protoc_version}-linux-x86_64.zip -d ./protoc-${protoc_version}-linux-x86_64 && \
		mv ${soft_dir}/protoc-${protoc_version}-linux-x86_64/bin/protoc ${cmd_path} && \
		echo "protoc 的版本是:" && protoc --version
}

golangci-lint(){
    echo "安装golangci-lint工具"
}

go_plug(){
        cd ${GOPATH} && export GOPROXY=https://goproxy.io && export GO111MODULE=off  && export GOPATH=${GOPATH} && \
        echo "GOPATH为:"${GOPATH} && \
		echo "安装 protobuf golang插件 protoc-gen-go protoc-gen-grpc-gateway protoc-gen-swagger protoc-go-inject-tag" && \
		echo "大概耗时30分钟" && \
		mkdir -p ${GOPATH}/src/golang && cd ${GOPATH}/src/golang && git clone https://github.com/golang/protobuf.git --depth 1
		mkdir -p ${GOPATH}/src/google.golang.org && cd ${GOPATH}/src/google.golang.org && \
		git clone https://github.com/google/go-genproto.git --depth 1  && mv  go-genproto/ genproto/ && \
		cd ${GOPATH} && \
		go get   github.com/golang/protobuf/protoc-gen-go && \
		go get   github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway && \
		go get   github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger && \
		go get  github.com/favadi/protoc-go-inject-tag && \
		go get  github.com/golangci/golangci-lint/cmd/golangci-lint && \
		echo "安装gocyclo圈复杂度计算工具" && \
		go get  github.com/fzipp/gocyclo && \
		echo "安装打包静态文件工具" && \
		go get  github.com/rakyll/statik && \
		echo "安装go-torch" && \
		go get github.com/uber/go-torch && \
		cd ${GOPATH}/src/github.com/uber/go-torch && \
		`git clone https://github.com/brendangregg/FlameGraph.git` || { echo "FlameGraph已经存在"; }
}

# 安装一些依赖工具到GOPATH
tool_install(){
    cd ${GOPATH}
    #代码风格审查
    go get  github.com/golangci/golangci-lint/cmd/golangci-lint
    go get  github.com/golang/protobuf/protoc-gen-go
    go get  github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
    go get  github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
    go get  github.com/favadi/protoc-go-inject-tag
    go get  github.com/rakyll/statik
}

# 安装依赖
case $1 in
    tool) echo "依赖工具安装"
        protoc_install
        tool_install
    ;;
    go) echo "安装go程序"
        go_install
    ;;
    other)
    echo "安装其他工具"
        cloc_install
    ;;
    alltool)
    echo "全部工具安装"
        go_plug
    ;;
    *)
    echo "安装go程序和依赖工具"
        go_install
        protoc_install
        tool_install
    ;;
esac
