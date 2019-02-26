#!/bin/bash

set -e

pidfile=server.pid
#build
pprofon() {
    f=deployments/bin/$1_$2/$pidfile
    p=`cat $f`
    echo "进程号:" $p
    kill -10 $p
}

pprofoff() {
    f=deployments/bin/$1_$2/$pidfile
    p=`cat $f`
    echo "进程号:" $p
    kill -12 $p
}

#判断如何build
case $1 in
    pprofon) echo "pprof on:"$2,$3
    if [ -z $2 -o -z $3 ];then
    echo "参数错误"
    exit 2
    fi
    pprofon $2 $3
    ;;
    pprofoff) echo "pprof off:"$2,$3
    if [ -z $2 -o -z $3 ];then
    echo "参数错误"
    exit 2
    fi
    pprofoff $2 $3
    ;;
    *)
    echo "pprof error"
    exit 2
    ;;
esac
