/*
SQLyog Community v13.1.7 (64 bit)
MySQL - 8.0.26
*********************************************************************
*/
/*!40101 SET NAMES utf8 */;

create table `sdk_usage` (
    `id` CHAR (36) NOT NULL PRIMARY KEY,
    `create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间' ,
    `update_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    `user_id` varchar (64),
    `all` bigint (20),
    `used` bigint (20),
    unique index address(`user_id` ASC)
);