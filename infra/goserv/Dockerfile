FROM golang:1.18
LABEL maintainer="ikura-zzz <chukicamata_kbmc@yahoo.co.jp"

WORKDIR /usr/src/app

RUN mkdir /usr/local/bin/html /var/log/myfav/
ADD src/golang/ ./
ADD src/html/ /usr/local/bin/html/
RUN go mod download && go mod verify

WORKDIR /usr/src/app/appsrv
RUN go build -v -o /usr/local/bin/app ./

CMD ["app"]