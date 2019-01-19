#!/bin/bash


proto() {
    dirname=./srv/$1/proto
    if [ -d $dirname ];then
		for f in $dirname/*.proto; do \
		    if [ -f $f ];then \
                protoc  --go_out=plugins=grpc:. $f; \
                echo compiled protoc: $f; \
            fi \
		done \
	fi
}

proto user
