/*
SQLyog Community v13.1.7 (64 bit)
MySQL - 8.0.26
*********************************************************************
*/
/*!40101 SET NAMES utf8 */;

create table `sdk_info` (
    `id` CHAR (36) NOT NULL PRIMARY KEY,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间' ,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    `name` varchar (256),
    `avatar` varchar (256),
    `desc` bigint (20),
    `download_url` varchar (512)
);