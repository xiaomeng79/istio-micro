#!/bin/bash

set -e

#go-bindata -o data/bindata.go -pkg data data/*.json

statik -src=`pwd`/srv/socket/asset -dest=`pwd`/srv/socket

#更新的sql
#statik -src=`pwd`/srv/account/sqlupdate -dest=`pwd`/srv/account

#更新sql数据拷贝
accountsql=`pwd`/deployments/bin/srv_account
mkdir -p ${accountsql} && cp -r `pwd`/srv/account/sqlupdate ${accountsql}