FROM centos:8.0
LABEL maintainer="ikura-zzz <chukicamata_kbmc@yahoo.co.jp"

RUN yum install -y httpd
RUN yum install -y openssh-server
RUN yum install -y passwd
RUN systemctl enable httpd
RUN systemctl enable openssh-server

RUN useradd web -m -G apache
RUN chgrp -R apache /var/www
RUN -R g+rwx /var/www
