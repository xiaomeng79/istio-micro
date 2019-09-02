#!/bin/bash

set -eu

source scripts/.variables.sh

#生成dockerfile
gen(){
#程序名称
pname="$1_$2"
#模板
filename=./deployments/bin/"$pname"
swagger=./deployments/config/swagger/$1/$2/proto
#判断bin是否存在
if [ ! -d deployments/bin/"$pname" ];then
mkdir -p deployments/bin/"$pname"/proto
fi

#判断swagger存在复制到这个目录下面
if [ -d $swagger ];then
cp -r $swagger/ $filename
fi

#添加本地时区
#RUN set -xe && apk add --no-cache tzdata && cp -r -f /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

cat>$filename/Dockerfile<<EOF
FROM alpine:3.5
RUN  echo 'http://mirrors.ustc.edu.cn/alpine/v3.5/main' > /etc/apk/repositories \
    && echo 'http://mirrors.ustc.edu.cn/alpine/v3.5/community' >>/etc/apk/repositories \
&& apk update && apk add tzdata \
&& ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
&& echo "Asia/Shanghai" > /etc/timezone

ADD $pname /$pname
ADD proto/ /swagger
ADD sqlupdate/ /sqlupdate
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
    gen srv account
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