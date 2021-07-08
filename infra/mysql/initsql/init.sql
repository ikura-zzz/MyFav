SET @old_autocommit = @@autocommit;
CREATE DATABASE `myfav` DEFAULT CHARACTER SET utf8mb4;
USE `myfav`;
DROP TABLE IF EXISTS `appusers`;
CREATE TABLE `appusers` (
    `id` int NOT NULL AUTO_INCREMENT,
    `username` char(50) NOT NULL DEFAULT '',
    `passhash` binary(32) NOT NULL DEFAULT '',
    `upddate` datetime NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
SET autocommit = @old_autocommit;