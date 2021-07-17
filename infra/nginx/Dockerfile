FROM nginx:1.20

RUN mkdir /usr/share/nginx/html/css /usr/share/nginx/html/img /usr/share/nginx/html/js
COPY html/*html /usr/share/nginx/html/
COPY html/*.ico /usr/share/nginx/html/
COPY html/img/* /usr/share/nginx/html/img/
COPY html/css/* /usr/share/nginx/html/css/
COPY html/js/* /usr/share/nginx/html/js/
COPY default.conf /etc/nginx/conf.d/
COPY logrotate.conf /etc/logrotate.d/nginx
RUN rm /var/log/nginx/*
RUN touch /var/log/nginx/access.log /var/log/nginx/error.log
