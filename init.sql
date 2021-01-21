CREATE SCHEMA IF NOT EXISTS `coverage-badge` COLLATE utf8_general_ci;

USE `coverage-badge`;

CREATE TABLE IF NOT EXISTS `project`
(
    `name`      varchar(255) NOT NULL,
    `lang`      varchar(255) NOT NULL,
    `token`     varchar(255) NOT NULL,
    `coverage`  float        NOT NULL,
    `timestamp` timestamp    NULL DEFAULT CURRENT_TIMESTAMP,
    `raw`       text         NOT NULL,
    PRIMARY KEY (`name`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8
  COLLATE = utf8_general_ci;
