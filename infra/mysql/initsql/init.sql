SET @old_autocommit = @@autocommit;
CREATE DATABASE `myfav` DEFAULT CHARACTER SET utf8mb4;
USE `myfav`;
DROP TABLE IF EXISTS `appusers`;
CREATE TABLE `appusers` (
    `userid` int NOT NULL AUTO_INCREMENT,
    `username` char(50) NOT NULL UNIQUE,
    `passhash` binary(32) NOT NULL DEFAULT '',
    `upddate` datetime NOT NULL,
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

INSERT INTO `appusers` (username,passhash,upddate) VALUES ('shigeji', 0x124ED3717B923C61E52653D77AB0FBE2AFFCD5E00DBE6CD821184C7864691B5B,'2021-07-08 00:00:00');
INSERT INTO `appusers` (username,passhash,upddate) VALUES ('maho', 0x124ED3717B923C61E52653D77AB0FBE2AFFCD5E00DBE6CD821184C7864691B5B,'2021-07-08 00:00:00');

INSERT INTO `favs` (userid,title,category,publisher,overview,impression,timing,stars,openclose,upddate) VALUES(1,'ニーア　レプリカント','ゲーム','スクウェア・エニックス','ニーアシリーズの１作目、PS4移植版','にーあがなんやかんやで苦労してとてもつらかった。','1','4','1','2021-07-08 00:00:00');
INSERT INTO `images` (favid,image) VALUE(1,'data:image/jpeg;base64,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');