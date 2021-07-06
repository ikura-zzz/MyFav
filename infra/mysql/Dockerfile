FROM mysql:8

EXPOSE 3306

ENV MYSQL_ROOT_PASSWORD=mysql
COPY my.cnf /etc/mysql/conf.d/my.cnf
COPY initsql/init.sql /docker-entrypoint-initdb.d/init.sql