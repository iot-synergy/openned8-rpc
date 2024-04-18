/*
SQLyog Community v13.1.7 (64 bit)
MySQL - 8.0.26
*********************************************************************
*/
/*!40101 SET NAMES utf8 */;

create table `app_info` (
    `id` CHAR (36) NOT NULL PRIMARY KEY,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间' ,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    `user_id` varchar (64),
    `app_name` varchar (256),
    `summary` varchar (1024),
    `app_category` bigint (20),
    `use_industry` bigint (20),
    `app_category_name` varchar(256),
    `use_industry_name` varchar(256),
    `app_key` varchar(256),
    `app_secret` varchar(256)
);