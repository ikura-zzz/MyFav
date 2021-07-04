FROM nginx:1.20

RUN mkdir /usr/share/nginx/html/css
COPY html/*html /usr/share/nginx/html/.
COPY html/css/* /usr/share/nginx/html/css/.
COPY default.conf /etc/nginx/conf.d/.
