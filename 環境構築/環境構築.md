# network作成
- mysql-network
``` docker
docker network create mysql-network
```

# imagbuild
- myapache
``` docker
docker build -t myapache centos/
```
- mynginx
``` docker
docker build -t mynginx nginx/
```
- goserv
``` docker
docker build -t goserv goserv/
```
# container起動
- wordpress
``` docker
docker run --network mysql-network -p 80:80 -p 443:443 -d -it --name wordpress wordpress
```
- myapache
``` docker
docker run --privileged --network mysql-network -p 22:22 -p 8080:8080 -d -it --name myapache myapache /sbin/init
```
- mysql
``` docker
docker run --network mysql-network -p 3306:3306 -e MYSQL_ROOT_PASSWORD=mysql -d -it --name mysql mysql
```
- nginx
``` docker
docker run --network mysql-network -p 80:80 -d -it --name mynginx mynginx
```
- goserv
``` docker
docker run --privileged --network mysql-network -p 8080:8080 -d -it --name goserv goserv /sbin/init
```

# データベース初期設定
[リンク](https://www.javadrive.jp/wordpress/install/index1.html#section2)

ただし、ユーザーのホストを%にする必要があった。