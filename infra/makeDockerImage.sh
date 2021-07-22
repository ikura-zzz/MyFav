#docker rm -f $(docker ps -q -a)
mynginx=`docker ps | grep mynginx | awk '{print $1}'`
goserv=`docker ps | grep goserv | awk '{print $1}'`
mysql=`docker ps | grep mysql | awk '{print $1}'`
docker rm -f $mynginx
docker rm -f $goserv
#docker rm -f $mysql

docker rmi $(docker images goserv -q)
docker rmi $(docker images mynginx -q)
#docker rmi $(docker images mysql -q)

chmod 755 goserv/bin/myfav

docker build -t mynginx nginx/
docker build -t goserv goserv/
#docker build -t mysql mysql/

docker run --privileged --network mysql-network -p 8080:8080 -d -it --name goserv goserv /sbin/init
docker run --network mysql-network -p 80:80 -d -it --name mynginx mynginx
#ocker run --network mysql-network -p 3306:3306 -d -it --name mysql mysql
if [[ $mysql = "" ]]; then
    docker start mysql
fi