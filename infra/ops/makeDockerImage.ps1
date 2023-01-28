Set-Location .\goserv\
.\gomake.ps1
Set-Location ..\

#docker rm -f $(docker ps -q -a)
$nginx=((docker ps -a | Select-String mynginx).Line).Split(" ")[0]
$goserv=((docker ps -a | Select-String goserv).Line).Split(" ")[0]
$mysql=((docker ps -a | Select-String mysql).Line).Split(" ")[0]
docker rm -f $nginx
docker rm -f $goserv
#docker rm -f $mysql

docker rmi $(docker images goserv -q)
docker rmi $(docker images mynginx -q)
#docker rmi $(docker images mysql -q)

docker build -t mynginx ../nginx/
docker build -t goserv ../goserv/
#docker build -t mysql ../mysql/

docker run --privileged --network mysql-network -v "/c/var/log/myfav:/var/log/myfav" -p 8080:8080 -d -it --name goserv goserv /sbin/init
docker run --network mysql-network -p 80:80 -d -it --name mynginx mynginx
#docker run --network mysql-network -p 3306:3306 -d -it --name mysql mysql
If ( $mysql = "\r\n" ) {
    docker start mysql
}
