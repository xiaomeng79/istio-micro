#!/bin/bash

set -eu

source scripts/version.sh
source scripts/.variables.sh

ProjectName=${ProjectName:-"github.com/xiaomeng79/istio-micro"}
Version=${Version:-"unknown-version"}
GoVersion=${GoVersion:-$(go version)}
#GoVersion=${GoVersion:-$(go version | awk '{print $3}')}
GitCommit=${GitCommit:-$(git rev-parse --short HEAD 2> /dev/null || true)}
BuiltTime=${BuiltTime:-$(date --utc --rfc-3339 ns 2> /dev/null | sed -e 's/ /T/')}
#TARGET=${TARGET:-'main'}
#GOOS=$(go env GOHOSTOS)
#
#if [ "${GOOS}" = "windows" ]; then
#	TARGET="${TARGET}.exe"
#fi


#build
build() {
    #判断bin是否存在
    if [ ! -d deployments/bin ];then
    mkdir -p deployments/bin
    fi
    #build

    dirname=./cmd/$1
    if [ -d $dirname ];then
		for f in $dirname/$2.go; do \
		    if [ -f $f ];then \

#		        CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -ldflags '-w' -i -o deployments/bin/$1_$2/$1_$2 -tags $1_$2 ./cmd/$1/
                buildapp $1 $2
                echo build over: $1_$2; \
            fi \
		done \
	fi
}

#build app
buildapp() {

    go build -mod=vendor -a -installsuffix cgo -ldflags                           \
    "                                           \
    -w                                           \
    -X '${ProjectName}/version.Version=${Version}'     \
    -X '${ProjectName}/version.GoVersion=${GoVersion}'       \
    -X '${ProjectName}/version.GitCommit=${GitCommit}'       \
    -X '${ProjectName}/version.BuiltTime=${BuiltTime}'       \
    "                                           \
    -o deployments/bin/${1}_${2}/${1}_${2} -tags ${1}_${2} ./cmd/${1}/
}

#全部build
allbuild() {
    build srv user
    build srv socket
    build api backend
    build api frontend
}
#判断如何build
case $1 in
    allbuild) echo "全部build"
    allbuild
    ;;
    build) echo "build:"$2,$3
    if [ -z $2 -o -z $3 ];then
    echo "参数错误"
    exit 2
    fi
    build $2 $3
    ;;
    *)
    echo "build error"
    exit 2
    ;;
esac
