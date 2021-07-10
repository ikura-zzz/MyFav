Remove-Item src\appsrv\appsrv 2> $null
$ENV:GOOS="linux"
$ENV:GOARCH="amd64"
Set-Location src\appsrv
go build 
Set-Location ..\..\
Move-Item -Force src\appsrv\appsrv bin\.