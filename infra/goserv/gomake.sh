#!/bin/bash
rm src/appsrv/myfav 2> /dev/null
export GOOS="linux"
export GOARCH="arm"
cd src/golang/appsrv
go build -o myfav
cd ../../../
mv -f src/golang/appsrv/myfav bin/.