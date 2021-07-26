#!/bin/bash

mysql=`docker ps | grep mysql | awk '{print $1}'`

docker run --privileged --network mysql-network -p 127.0.0.1:8080:8080 -d -it --name goserv goserv /sbin/init
docker run --network mysql-network -p 443:443 -v "/etc/letsencrypt:/etc/letsencrypt" -d -it --name mynginx mynginx 
#docker run --network mysql-network -p 3306:3306 -d -it --name mysql mysql
if [[ $mysql = "" ]]; then
    docker start mysql
fi