FROM centos:centos8.3.2011
LABEL maintainer="ikura-zzz <chukicamata_kbmc@yahoo.co.jp"

ENV TZ='Asia/Tokyo'

RUN mkdir /var/goserv /var/goserv/bin /var/goserv/bin/html /var/goserv/bin/static 
#RUN touch /var/log/myfav/myfav.log
#RUN chmod 666 /var/log/myfav/myfav.log
COPY bin/myfav /var/goserv/bin/
COPY bin/html/* /var/goserv/bin/html/
COPY conf/myfav.service /etc/systemd/system/

RUN chmod 755 /var/goserv/bin/myfav
RUN systemctl enable myfav.service
