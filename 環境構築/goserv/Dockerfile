FROM centos:centos8.3.2011
LABEL maintainer="ikura-zzz <chukicamata_kbmc@yahoo.co.jp"

RUN yum install -y openssh-server; yum clean all; systemctl enable sshd.service
RUN yum install -y passwd

COPY sshd_config /etc/ssh/sshd_config

RUN useradd app -m -G apache
RUN chgrp -R apache /var/www
RUN chmod -R g+rwx /var/www
