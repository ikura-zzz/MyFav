#!/bin/bash

shlog=/var/log/myfav/sh.log
/home/docker/workspace/myfav/infra/stopContainers.sh >> $shlog
sudo docker run -it --rm --name certbot -p 80:80 -p 443:443 -v "/etc/letsencrypt:/etc/letsencrypt" -v "/var/lib/letsencrypt:/var/lib/letsencrypt" certbot/certbot renew >> $shlog
sudo docker run -it --rm --name certbot -p 80:80 -p 443:443 -v "/etc/letsencrypt:/etc/letsencrypt" -v "/var/lib/letsencrypt:/var/lib/letsencrypt" certbot/certbot certificates >> $shlog
/home/docker/workspace/myfav/infra/startContainers.sh >> $shlog