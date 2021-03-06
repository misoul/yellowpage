CREATE DATABASE IF NOT EXISTS testdb1;

USE testdb1;

CREATE TABLE `companies` (
	`id` BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `name` VARCHAR(50),
    `industries` VARCHAR(50),
    `website` VARCHAR(50),
    `found_date` DATETIME,
    `stock_code` VARCHAR(50),
    `desc` VARCHAR(100)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 ROW_FORMAT=FIXED COMMENT='companies';

CREATE TABLE `comments` (
	`id` BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `author` VARCHAR(50),
    `text` VARCHAR(200)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 ROW_FORMAT=FIXED COMMENT='comments';
