/*
SQLyog Community v13.1.7 (64 bit)
MySQL - 8.0.26
*********************************************************************
*/
/*!40101 SET NAMES utf8 */;

create table `active_code_info` (
    `id` CHAR (36) NOT NULL PRIMARY KEY,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间' ,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    `active_key` varchar (256),
    `user_id` varchar (64),
    `app_id` varchar(64),
    `active_ip` varchar(64),
    `device_sn` varchar(64),
    `device_mac` varchar(256),
    `device_identity` varchar(256),
    `active_date` TIMESTAMP COMMENT '激活时间',
    `active_type` bigint(20),
    `active_file` varchar(256),
    `version` varchar(256),
    `start_date` TIMESTAMP,
    `expire_date` TIMESTAMP,
    `status` bigint(20)
);