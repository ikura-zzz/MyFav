Remove-Item websrv
$ENV:GOOS="linux"
$ENV:GOARCH="amd64"
go build
Move-Item -Force websrv ..\..\bin\.