#!/bin/bash

mynginx=`docker ps -a | grep mynginx | awk '{print $1}'`
goserv=`docker ps -a | grep goserv | awk '{print $1}'`
mysql=`docker ps -a | grep mysql | awk '{print $1}'`
echo "stop container:$mynginx"
docker rm -f $mynginx
echo "stop container:$goserv"
docker rm -f $goserv
echo "stop container:$mysql"
docker stop $mysql
#docker rm -f $(docker ps -q -a)