FROM centos:centos8.3.2011
LABEL maintainer="ikura-zzz <chukicamata_kbmc@yahoo.co.jp"

RUN mkdir /var/goserv /var/goserv/bin /var/goserv/bin/html /var/goserv/bin/static
COPY bin/websrv /var/goserv/bin/.
COPY bin/html/* /var/goserv/bin/html/.
COPY bin/static/* /var/goserv/bin/static/.
