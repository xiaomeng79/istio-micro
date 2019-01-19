#!/bin/bash

set -e

#docker-compose 命令
case $1 in
    up) echo "docker-compose up"
    cd deployments && sudo docker-compose up --build
    ;;
    dup) echo "docker-compose dup"
    cd deployments && sudo docker-compose up --build -d
    ;;
    stop) echo "docker-compose stop"
    cd deployments && sudo docker-compose stop
    ;;
    restart) echo "docker-compose restart"
    cd deployments && sudo docker-compose restart
    ;;
    ps) echo "docker-compose ps"
    cd deployments && sudo docker-compose ps
    ;;
    kill) echo "docker-compose kill"
    cd deployments && sudo docker-compose kill
    ;;
    rm) echo "docker-compose rm"
    cd deployments && sudo docker-compose rm
    ;;
    *) echo "docker-compose up"
    cd deployments && sudo docker-compose up --build
    ;;
esac
