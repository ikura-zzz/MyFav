#!/bin/bash

docker rmi $(docker images goserv -q)
docker rmi $(docker images mynginx -q)
#docker rmi $(docker images mysql -q)