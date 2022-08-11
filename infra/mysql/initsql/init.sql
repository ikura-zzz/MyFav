SET @old_autocommit = @@autocommit;

CREATE DATABASE `myfav` DEFAULT CHARACTER SET utf8mb4;
CREATE USER 'myfav'@'172.18.0.2' identified by 'wYemyWoRcXti';
GRANT all ON myfav.* TO 'myfav'@'172.18.0.2';
CREATE USER 'myfav'@'172.18.0.3' identified by 'wYemyWoRcXti';
GRANT all ON myfav.* TO 'myfav'@'172.18.0.3';
CREATE USER 'myfav'@'172.18.0.4' identified by 'wYemyWoRcXti';
GRANT all ON myfav.* TO 'myfav'@'172.18.0.4';
USE `myfav`;
DROP TABLE IF EXISTS `appusers`;
CREATE TABLE `appusers` (
    `userid` int NOT NULL AUTO_INCREMENT,
    `username` char(50) NOT NULL UNIQUE,
    `passhash` binary(32) NOT NULL DEFAULT '',
    `upddate` datetime NOT NULL,
    `lastlogin` datetime,
    PRIMARY KEY (`userid`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

DROP TABLE IF EXISTS `favs`;
CREATE TABLE `favs` (
    `favid` int NOT NULL AUTO_INCREMENT,
    `userid` int NOT NULL,
    `title` char(100) NOT NULL,
    `category` char(50) DEFAULT '',
    `publisher` char(50) DEFAULT '',
    `overview` TEXT,
    `impression` MEDIUMTEXT,
    `timing` char(1) NOT NULL,
    `stars` char(1) NOT NULL,
    `openclose` char(1) NOT NULL,
    `upddate` datetime NOT NULL,
    PRIMARY KEY (`favid`),
    FOREIGN KEY (userid) REFERENCES appusers(userid)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

DROP TABLE IF EXISTS `images`;
CREATE TABLE `images`(
    `imageid` int NOT NULL AUTO_INCREMENT,
    `favid` int NOT NULL,
    `image` MEDIUMTEXT NOT NULL,
    PRIMARY KEY (`imageid`),
    FOREIGN KEY (favid) REFERENCES favs(favid)
)ENGINE = InnoDB;

SET autocommit = @old_autocommit;

INSERT INTO `appusers` (username,passhash,upddate) VALUES ('ikura', 0x124ED3717B923C61E52653D77AB0FBE2AFFCD5E00DBE6CD821184C7864691B5B,'2021-07-08 00:00:00');
INSERT INTO `appusers` (username,passhash,upddate) VALUES ('ebima', 0x124ED3717B923C61E52653D77AB0FBE2AFFCD5E00DBE6CD821184C7864691B5B,'2021-07-08 00:00:00');

