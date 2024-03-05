/*
MySQL Backup
Database: sviwo
Backup Time: 2024-02-29 19:14:25
*/

SET FOREIGN_KEY_CHECKS=0;
DROP TABLE IF EXISTS `sviwo`.`sw_app_log`;
DROP TABLE IF EXISTS `sviwo`.`sw_app_param`;
DROP TABLE IF EXISTS `sviwo`.`sw_app_text`;
DROP TABLE IF EXISTS `sviwo`.`sw_app_videos`;
DROP TABLE IF EXISTS `sviwo`.`sw_battery`;
DROP TABLE IF EXISTS `sviwo`.`sw_car`;
DROP TABLE IF EXISTS `sviwo`.`sw_travel_record`;
DROP TABLE IF EXISTS `sviwo`.`sw_user`;
DROP TABLE IF EXISTS `sviwo`.`sw_user_auth`;
DROP TABLE IF EXISTS `sviwo`.`sw_user_car`;
DROP TABLE IF EXISTS `sviwo`.`sw_version`;
CREATE TABLE `sw_app_log` (
  `app_log_id` bigint(20) NOT NULL,
  `user_id` bigint(20) NOT NULL,
  `business_id` bigint(20) DEFAULT NULL COMMENT '业务id，与类型字段组合查询',
  `log_type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '操作类型：0=其他，1=车辆，2=账户',
  `log_content` varchar(255) DEFAULT NULL COMMENT '操作内容',
  `create_time` datetime NOT NULL,
  `update_time` datetime DEFAULT NULL,
  `is_delete` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否删除：true=已删除，false=正常\n',
  PRIMARY KEY (`app_log_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='app操作日志表';
CREATE TABLE `sw_app_param` (
  `param_id` bigint(20) NOT NULL,
  `parent_id` bigint(20) NOT NULL,
  `param_name` varchar(50) DEFAULT NULL COMMENT '参数名称',
  `param_value` varchar(2000) DEFAULT NULL COMMENT '参数值',
  `create_time` datetime NOT NULL,
  `update_time` datetime DEFAULT NULL,
  `is_delete` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否删除：true=已删除，false=正常\n',
  PRIMARY KEY (`param_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='App参数表';
CREATE TABLE `sw_app_text` (
  `text_id` bigint(20) NOT NULL,
  `text_name` varchar(100) NOT NULL COMMENT '文本名称',
  `text_type` int(1) NOT NULL DEFAULT '0' COMMENT '文本类型：0=账户，1=车辆功能和充电，2=售后服务和维修保养，3=其他',
  `text_content` longtext COMMENT '文本内容',
  `create_time` datetime NOT NULL,
  `update_time` datetime DEFAULT NULL,
  `is_delete` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否删除：true=已删除，false=正常\n',
  PRIMARY KEY (`text_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='App文本表';
CREATE TABLE `sw_app_videos` (
  `videos_id` bigint(20) NOT NULL,
  `videos_name` varchar(100) DEFAULT NULL COMMENT '视频名称',
  `videos_desc` varchar(400) DEFAULT NULL COMMENT '视频简介',
  `videos_type` int(1) NOT NULL DEFAULT '0' COMMENT '视频类型：0=充电，1=车辆驾驶，2=控制和设置，3=锁定和解锁，4=账户',
  `videos_url` varchar(50) DEFAULT NULL COMMENT '视频链接',
  `create_time` datetime NOT NULL,
  `update_time` datetime DEFAULT NULL,
  `is_delete` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否删除：true=已删除，false=正常\n',
  PRIMARY KEY (`videos_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='App视频表';
CREATE TABLE `sw_battery` (
  `battery_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `car_id` bigint(20) NOT NULL COMMENT '车辆id',
  `battery_code` varchar(50) NOT NULL DEFAULT '0' COMMENT '电池编号',
  `battery_temp` int(3) unsigned zerofill NOT NULL DEFAULT '000' COMMENT '电池温度',
  `create_time` datetime NOT NULL,
  `update_time` datetime DEFAULT NULL,
  `is_delete` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否删除：true=已删除，false=正常\n',
  PRIMARY KEY (`battery_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8 COMMENT='电池表';
CREATE TABLE `sw_car` (
  `car_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `car_nickname` varchar(50) DEFAULT NULL COMMENT '车辆昵称（用户自定义）',
  `car_final_position` varchar(50) DEFAULT NULL COMMENT '车辆最后定位',
  `travel_km` int(4) NOT NULL DEFAULT '0' COMMENT '行驶公里数',
  `car_frame_code` varchar(50) DEFAULT NULL COMMENT '车架编号',
  `car_motor_code` varchar(50) DEFAULT NULL COMMENT '车辆电机编号',
  `residue_electricity` int(3) NOT NULL DEFAULT '0' COMMENT '剩余电量',
  `residue_km` int(4) NOT NULL DEFAULT '0' COMMENT '剩余公里数',
  `after_sales_time` date NOT NULL COMMENT '保修日期',
  `activation_time` datetime DEFAULT NULL COMMENT '激活时间',
  `create_time` datetime NOT NULL,
  `update_time` datetime DEFAULT NULL,
  `is_delete` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否删除：true=已删除，false=正常\n',
  PRIMARY KEY (`car_id`)
) ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8 COMMENT='车辆表';
CREATE TABLE `sw_travel_record` (
  `travel_record_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) DEFAULT NULL COMMENT '用户ID',
  `car_id` bigint(20) DEFAULT NULL COMMENT '车辆ID',
  `start_point` varchar(5) DEFAULT NULL COMMENT '起点',
  `end_point` varchar(20) DEFAULT NULL COMMENT '终点',
  `mileage_driven` int(10) DEFAULT '0' COMMENT '行驶里程，单位（m）',
  `start_time` datetime DEFAULT NULL COMMENT '行程开始时间',
  `end_time` datetime DEFAULT NULL COMMENT '行程结束时间',
  `avg_speed` varchar(6) DEFAULT NULL COMMENT '平均时速，单位（m）',
  `consumption` int(4) DEFAULT NULL COMMENT '使用电量',
  `create_time` datetime NOT NULL,
  `update_time` datetime DEFAULT NULL,
  `is_delete` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否删除：true=已删除，false=正常\n',
  PRIMARY KEY (`travel_record_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8 COMMENT='行程记录表';
CREATE TABLE `sw_user` (
  `user_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL,
  `password` varchar(100) NOT NULL,
  `pwd_salt` varchar(255) NOT NULL COMMENT '密码盐值',
  `pwd_encry_num` int(11) NOT NULL DEFAULT '0' COMMENT '密码加密次数',
  `first_name` varchar(100) DEFAULT NULL,
  `last_name` varchar(100) DEFAULT NULL,
  `enable` bit(1) NOT NULL DEFAULT b'1' COMMENT '账号是否可用：true=正常，false=停用\n',
  `head_img` varchar(100) DEFAULT NULL,
  `mobile_phone` varchar(20) DEFAULT NULL COMMENT '手机号',
  `user_address` varchar(255) DEFAULT NULL COMMENT '用户地址',
  `create_time` datetime NOT NULL,
  `update_time` datetime DEFAULT NULL,
  `is_delete` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否删除：true=已删除，false=正常\n',
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `username_un_index` (`username`) USING BTREE COMMENT '用户名唯一键'
) ENGINE=InnoDB AUTO_INCREMENT=10001 DEFAULT CHARSET=utf8 COMMENT='App用户表';
CREATE TABLE `sw_user_auth` (
  `auth_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) DEFAULT NULL,
  `auth_first_name` varchar(100) NOT NULL,
  `auth_last_name` varchar(100) NOT NULL,
  `certificate_front_img` varchar(100) NOT NULL COMMENT '证件正面照片',
  `certificate_back_img` varchar(100) NOT NULL COMMENT '证件背面照片',
  `auth_status` int(1) unsigned zerofill NOT NULL DEFAULT '0' COMMENT '认证状态：0=未认证，1=认证中，2=认证成功，3=认证失败',
  `auth_fail_reason` varchar(255) DEFAULT NULL COMMENT '认证失败原因',
  `auth_time` datetime DEFAULT NULL COMMENT '认证时间',
  `verify_time` datetime DEFAULT NULL COMMENT '审核时间',
  `create_time` datetime NOT NULL,
  `update_time` datetime DEFAULT NULL,
  PRIMARY KEY (`auth_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1763157909782401025 DEFAULT CHARSET=utf8 COMMENT='实名认证表';
CREATE TABLE `sw_user_car` (
  `user_car_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `car_id` bigint(20) NOT NULL,
  `user_id` bigint(20) NOT NULL,
  `mobile_key` bit(1) NOT NULL DEFAULT b'0' COMMENT '手机钥匙开关：0=关，1=开',
  `speed_limit` bit(1) NOT NULL DEFAULT b'0' COMMENT '速度限制开关：0=关，1=开',
  `driving_mode_type` int(1) NOT NULL DEFAULT '0' COMMENT '驾驶模式：0=ECO模式，1=运动模式，2=狂暴模式，3=脱困模式',
  `create_time` datetime NOT NULL,
  PRIMARY KEY (`user_car_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8 COMMENT='用户车辆表';
CREATE TABLE `sw_version` (
  `version_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `version_number` int(5) NOT NULL COMMENT '版本号，用于app版本比较判断',
  `version_code` varchar(20) NOT NULL COMMENT '版本编码：app显示当前版本号使用，例如：V1.1.1',
  `version_type` int(1) NOT NULL DEFAULT '0' COMMENT '版本类型：0=APP更新，1=固件升级',
  `version_update_type` int(1) NOT NULL DEFAULT '0' COMMENT '版本更新类型：0=弱更新，1=强更新',
  `version_status` int(1) NOT NULL DEFAULT '0' COMMENT '版本发布状态：0=待发布，1=已发布，2=已过期',
  `version_url` varchar(200) NOT NULL COMMENT '版本链接',
  `version_desc` varchar(1000) DEFAULT NULL COMMENT '版本描述，用于app显示的新版本信息',
  `create_time` datetime NOT NULL,
  `update_time` datetime DEFAULT NULL,
  `is_delete` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否删除：true=已删除，false=正常\n',
  PRIMARY KEY (`version_id`)
) ENGINE=InnoDB AUTO_INCREMENT=654654321365 DEFAULT CHARSET=utf8 COMMENT='App版本管理表';
BEGIN;
LOCK TABLES `sviwo`.`sw_app_log` WRITE;
DELETE FROM `sviwo`.`sw_app_log`;
UNLOCK TABLES;
COMMIT;
BEGIN;
LOCK TABLES `sviwo`.`sw_app_param` WRITE;
DELETE FROM `sviwo`.`sw_app_param`;
UNLOCK TABLES;
COMMIT;
BEGIN;
LOCK TABLES `sviwo`.`sw_app_text` WRITE;
DELETE FROM `sviwo`.`sw_app_text`;
UNLOCK TABLES;
COMMIT;
BEGIN;
LOCK TABLES `sviwo`.`sw_app_videos` WRITE;
DELETE FROM `sviwo`.`sw_app_videos`;
UNLOCK TABLES;
COMMIT;
BEGIN;
LOCK TABLES `sviwo`.`sw_battery` WRITE;
DELETE FROM `sviwo`.`sw_battery`;
UNLOCK TABLES;
COMMIT;
BEGIN;
LOCK TABLES `sviwo`.`sw_car` WRITE;
DELETE FROM `sviwo`.`sw_car`;
UNLOCK TABLES;
COMMIT;
BEGIN;
LOCK TABLES `sviwo`.`sw_travel_record` WRITE;
DELETE FROM `sviwo`.`sw_travel_record`;
INSERT INTO `sviwo`.`sw_travel_record` (`travel_record_id`,`user_id`,`car_id`,`start_point`,`end_point`,`mileage_driven`,`start_time`,`end_time`,`avg_speed`,`consumption`,`create_time`,`update_time`,`is_delete`) VALUES (1, 1, 1, 'qqq', '11q', 0, '2024-02-27 12:51:46', '2024-02-27 12:51:50', 'qq', 222, '2024-02-27 12:51:56', '2024-02-27 12:51:59', b'1'),(2, 1, 1, 'qqq1', '11q', 0, '2024-02-27 12:55:30', '2024-02-27 12:55:30', 'qq', 222, '2024-02-27 12:55:30', '2024-02-27 12:55:30', b'1'),(3, 1, 1, 'qqq1', '11q', 0, '2024-02-27 12:55:44', '2024-02-27 12:55:44', 'qq', 222, '2024-02-27 12:55:44', '2024-02-27 12:55:44', b'1'),(4, 1, 1, 'qqq1', '11q', 0, '2024-02-27 12:55:47', '2024-02-27 12:55:47', 'qq', 222, '2024-02-27 12:55:47', '2024-02-27 12:55:47', b'1'),(5, 1, 1, 'qqq1', '11q', 0, '2024-02-27 12:55:51', '2024-02-27 12:55:51', 'qq', 222, '2024-02-27 12:55:51', '2024-02-27 12:55:51', b'1'),(6, 1, 1, 'qqq1', '11q', 0, '2024-02-27 12:55:54', '2024-02-27 12:55:54', 'qq', 222, '2024-02-27 12:55:54', '2024-02-27 12:55:54', b'1'),(7, 1, 1, 'qqq1', '11q', 0, '2024-02-27 12:55:57', '2024-02-27 12:55:57', 'qq', 222, '2024-02-27 12:55:57', '2024-02-27 12:55:57', b'1'),(8, 1, 1, 'qqq1', '11q', 0, '2024-02-27 12:56:00', '2024-02-27 12:56:00', 'qq', 222, '2024-02-27 12:56:00', '2024-02-27 12:56:00', b'1'),(9, 1, 1, 'qqq1', '11q', 0, '2024-02-27 12:56:04', '2024-02-27 12:56:04', 'qq', 222, '2024-02-27 12:56:04', '2024-02-27 12:56:04', b'1'),(10, 1, 1, 'qqq1', '11q', 0, '2024-02-27 12:56:18', '2024-02-27 12:56:18', 'qq', 222, '2024-02-27 12:56:18', '2024-02-27 12:56:18', b'1'),(11, 1, 1, 'qqq1', '11q', 0, '2024-02-27 12:56:11', '2024-02-27 12:56:11', 'qq', 222, '2024-02-27 12:56:11', '2024-02-27 12:56:11', b'0')
;
UNLOCK TABLES;
COMMIT;
BEGIN;
LOCK TABLES `sviwo`.`sw_user` WRITE;
DELETE FROM `sviwo`.`sw_user`;
INSERT INTO `sviwo`.`sw_user` (`user_id`,`username`,`password`,`pwd_salt`,`pwd_encry_num`,`first_name`,`last_name`,`enable`,`head_img`,`mobile_phone`,`user_address`,`create_time`,`update_time`,`is_delete`) VALUES (10000, '821317143@qq.com', '7d1ca1fdbecbeb53db69d7e04d58df96', 'NFRWBJdOLf', 20, 'Gou', 'YiChuan', b'1', '', '15928736853', '成都市都江堰市', '2024-02-29 19:02:58', NULL, b'0')
;
UNLOCK TABLES;
COMMIT;
BEGIN;
LOCK TABLES `sviwo`.`sw_user_auth` WRITE;
DELETE FROM `sviwo`.`sw_user_auth`;
INSERT INTO `sviwo`.`sw_user_auth` (`auth_id`,`user_id`,`auth_first_name`,`auth_last_name`,`certificate_front_img`,`certificate_back_img`,`auth_status`,`auth_fail_reason`,`auth_time`,`verify_time`,`create_time`,`update_time`) VALUES (1763157909782401024, 10000, '', '', '', '', 0, '', NULL, NULL, '2024-02-29 19:02:58', NULL)
;
UNLOCK TABLES;
COMMIT;
BEGIN;
LOCK TABLES `sviwo`.`sw_user_car` WRITE;
DELETE FROM `sviwo`.`sw_user_car`;
UNLOCK TABLES;
COMMIT;
BEGIN;
LOCK TABLES `sviwo`.`sw_version` WRITE;
DELETE FROM `sviwo`.`sw_version`;
INSERT INTO `sviwo`.`sw_version` (`version_id`,`version_number`,`version_code`,`version_type`,`version_update_type`,`version_status`,`version_url`,`version_desc`,`create_time`,`update_time`,`is_delete`) VALUES (654654321361, 1, 'V1.0.0', 1, 0, 1, 'asdasdasd', 'sdasdas', '2024-02-22 16:21:02', NULL, b'1'),(654654321362, 2, 'V1.0.1', 1, 0, 1, 'asdasdasd', 'sdasdas', '2024-02-22 16:21:02', NULL, b'1'),(654654321363, 3, 'V1.0.1', 0, 0, 1, 'asdasdasd', 'sdasdas', '2024-02-22 16:21:02', NULL, b'1'),(654654321364, 1, 'V1.0.0', 0, 0, 1, 'asdasdasd', 'sdasdas', '2024-02-22 16:21:02', NULL, b'1')
;
UNLOCK TABLES;
COMMIT;
