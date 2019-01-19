#!/bin/bash

set -e

#生成dockerfile
gen(){
#程序名称
pname="$1_$2"
#模板
filename=./deployments/bin/"$pname"
#判断bin是否存在
if [ ! -d deployments/bin/"$pname" ];then
mkdir -p deployments/bin/"$pname"
fi

cat>$filename/Dockerfile<<EOF
FROM alpine:3.2
RUN set -xe && apk add --no-cache tzdata && cp -r -f /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
ADD $pname /$pname
RUN chmod +x /$pname
ENTRYPOINT [ "/$pname" ]
EOF
echo "生成dockerfile $pname"
}

#全部生成dockerfile
allgen() {
    gen api backend
    gen api frontend
    gen srv user
    gen srv socket
}

#判断如何build
case $1 in
    alldf) echo "全部生成dockerfile"
    allgen
    ;;
    df) echo "生成dockerfile:"$2,$3
    if [ -z $2 -o -z $3 ];then
    echo "参数错误"
    exit 2
    fi
    gen $2 $3
    ;;
    *)
    echo "生成dockerfile error"
    exit 2
    ;;
esac