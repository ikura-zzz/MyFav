Remove-Item appsrv 2> $null
$ENV:GOOS="linux"
$ENV:GOARCH="amd64"
go build
Move-Item -Force appsrv ..\..\bin\.