SET @old_autocommit = @@autocommit;
CREATE DATABASE `myfav` DEFAULT CHARACTER SET utf8mb4;
USE `myfav`;
DROP TABLE IF EXISTS `appusers`;
CREATE TABLE `appusers` (
    `id` int NOT NULL AUTO_INCREMENT,
    `username` char(50) NOT NULL UNIQUE,
    `passhash` binary(32) NOT NULL DEFAULT '',
    `upddate` datetime NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

SET autocommit = @old_autocommit;

INSERT INTO `appusers` (username,passhash,upddate) VALUES ('shigeji', 0x124ED3717B923C61E52653D77AB0FBE2AFFCD5E00DBE6CD821184C7864691B5B,'2021-07-08 00:00:00')
