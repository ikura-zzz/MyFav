rm appsrv 2> /dev/null
export GOOS="linux"
export GOARCH="amd64"
go build
mv appsrv ../../bin/.