FROM ubuntu:20.04
LABEL maintainer="ikura-zzz <chukicamata_kbmc@yahoo.co.jp"

RUN mkdir /var/goserv /var/goserv/bin /var/goserv/bin/htmls /var/goserv/bin/static
COPY bin/websrv /var/goserv/bin/.
COPY bin/htmls/* /var/goserv/bin/htmls/.
COPY bin/static/* /var/goserv/bin/static/.
