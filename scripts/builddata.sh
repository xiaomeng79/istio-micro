#!/bin/bash

set -e

#go-bindata -o data/bindata.go -pkg data data/*.json

statik -src=`pwd`/srv/socket/asset -dest=`pwd`/srv/socket