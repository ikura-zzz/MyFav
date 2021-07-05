cd .\goserv\src\appsrv\
.\gomake.ps1
cd ..\..\..\

docker rm -f $(docker ps -q -a)
docker rmi $(docker  images -q)

docker build -t mynginx nginx/
docker build -t goserv goserv/
docker build -t mysql mysql/

docker run --privileged --network mysql-network -p 8080:8080 -d -it --name goserv goserv /sbin/init
docker run --network mysql-network -p 80:80 -d -it --name mynginx mynginx
docker run --network mysql-network -p 3306:3306 -d -it --name mysql mysql
#docker run --network mysql-network -p 3306:3306 -e MYSQL_ROOT_PASSWORD=mysql -d -it --name mysql mysql