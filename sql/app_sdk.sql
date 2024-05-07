/*
SQLyog Community v13.1.7 (64 bit)
MySQL - 8.0.26
*********************************************************************
*/
/*!40101 SET NAMES utf8 */;

drop table `app_sdk`

create table `app_sdk` (
    `id` CHAR (36) NOT NULL PRIMARY KEY,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间' ,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    `app` varchar (36),
    `sdk` varchar (36),
    `sdk_key` varchar (32)
);