Remove-Item src\appsrv\myfav 2> $null
$ENV:GOOS="linux"
$ENV:GOARCH="amd64"
Set-Location src\golang\appsrv
go build -o myfav
Set-Location ..\..\
Move-Item -Force src\golang\appsrv\myfav bin\.