FROM nginx:1.20

RUN mkdir /usr/share/nginx/html/css /usr/share/nginx/html/img
COPY html/*html /usr/share/nginx/html/.
COPY html/img/* /usr/share/nginx/html/img/.
#COPY html/css/* /usr/share/nginx/html/css/.
COPY default.conf /etc/nginx/conf.d/.
