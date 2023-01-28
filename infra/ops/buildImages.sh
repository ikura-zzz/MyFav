#!/bin/bash

mv nginx/default_serv.conf nginx/default.conf

chmod 755 goserv/bin/myfav

docker build -t mynginx ../nginx/
docker build -t goserv ../goserv/
docker build -t mysql ../mysql/