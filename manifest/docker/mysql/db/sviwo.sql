/*
MySQL Backup
Database: sviwo
Backup Time: 2024-07-18 19:16:47
*/

SET FOREIGN_KEY_CHECKS=0;
DROP TABLE IF EXISTS `sviwo`.`sw_alarm_log`;
DROP TABLE IF EXISTS `sviwo`.`sw_app_media`;
DROP TABLE IF EXISTS `sviwo`.`sw_app_param`;
DROP TABLE IF EXISTS `sviwo`.`sw_device`;
DROP TABLE IF EXISTS `sviwo`.`sw_product`;
DROP TABLE IF EXISTS `sviwo`.`sw_travel_record`;
DROP TABLE IF EXISTS `sviwo`.`sw_user`;
DROP TABLE IF EXISTS `sviwo`.`sw_user_auth`;
DROP TABLE IF EXISTS `sviwo`.`sw_user_device`;
DROP TABLE IF EXISTS `sviwo`.`sw_version`;
CREATE TABLE `sw_alarm_log` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `type` int NOT NULL DEFAULT '0' COMMENT '告警类型：1=规则告警，2=设备自主告警',
  `data` varchar(3000) DEFAULT NULL COMMENT '触发告警的数据',
  `product_key` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT '0' COMMENT '产品标识',
  `device_key` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '设备标识',
  `status` int NOT NULL DEFAULT '0' COMMENT '告警状态：0=未处理，1=已处理',
  `content` varchar(400) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '处理意见',
  `updated_by` bigint DEFAULT NULL COMMENT '告警处理人员',
  `created_time` datetime NOT NULL COMMENT '告警时间',
  `updated_time` datetime DEFAULT NULL COMMENT '处理时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb3 COMMENT='app操作日志表';
CREATE TABLE `sw_app_media` (
  `id` bigint NOT NULL,
  `parent_id` bigint NOT NULL DEFAULT '0',
  `enable` bit(1) NOT NULL DEFAULT b'1' COMMENT '显示或屏蔽：true=显示，false=屏蔽',
  `page_type` int NOT NULL COMMENT '页面类型：0=帮助页，1=注册用户页，2=服务页',
  `display_type` int NOT NULL DEFAULT '0' COMMENT '显示类型：0=直接显示，1=跳转外链',
  `title` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '标题',
  `media_desc` varchar(1000) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '简介',
  `small_img` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '缩略图',
  `icon` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `content` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci COMMENT '内容',
  `orders` int NOT NULL DEFAULT '999' COMMENT '排序',
  `create_time` datetime NOT NULL,
  `update_time` datetime DEFAULT NULL,
  `is_delete` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除：true=已删除，false=正常\n',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='App视频表';
CREATE TABLE `sw_app_param` (
  `param_id` bigint NOT NULL,
  `parent_id` bigint NOT NULL,
  `param_name` varchar(50) DEFAULT NULL COMMENT '参数名称',
  `param_value` varchar(2000) DEFAULT NULL COMMENT '参数值',
  `param_const` varchar(50) DEFAULT NULL COMMENT '参数常量',
  `create_time` datetime NOT NULL,
  `update_time` datetime DEFAULT NULL,
  `is_delete` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除：true=已删除，false=正常\n',
  PRIMARY KEY (`param_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='App参数表';
CREATE TABLE `sw_device` (
  `device_id` bigint NOT NULL AUTO_INCREMENT,
  `product_id` bigint NOT NULL COMMENT '所属产品',
  `sim_id` varchar(30) NOT NULL,
  `product_key` varchar(100) NOT NULL COMMENT '对应物联网平台产品的ProductKey',
  `device_name` varchar(255) NOT NULL COMMENT '对应物联网平台颁发的设备证书的DeviceName',
  `device_secret` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '对应物联网平台颁发的设备证书的DeviceSecret',
  `device_model` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '设备型号',
  `nickname` varchar(100) DEFAULT NULL COMMENT '产品昵称（目前只有ATV，则等同于车辆昵称）',
  `status` int NOT NULL DEFAULT '0' COMMENT '状态：0=未启用，1=离线，2=在线',
  `online_timeout` int NOT NULL DEFAULT '10' COMMENT '设备在线超时设置，单位：秒',
  `activate_time` datetime DEFAULT NULL COMMENT '激活时间',
  `registry_time` datetime DEFAULT NULL COMMENT '注册时间',
  `version` varchar(20) DEFAULT NULL COMMENT '固件版本号',
  `last_online_time` datetime DEFAULT NULL COMMENT '最后上线时间',
  `metadata_table` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否生成物模型表：0=否，1=是',
  `bluetooth_address` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '蓝牙地址',
  `bluetooth_secret_key` varchar(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '蓝牙握手密钥',
  `create_time` datetime NOT NULL,
  `update_time` datetime DEFAULT NULL,
  `is_delete` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除：true=已删除，false=正常\n',
  PRIMARY KEY (`device_id`) USING BTREE,
  UNIQUE KEY `index_ device_name_unioue` (`device_name`) USING BTREE COMMENT '设备名称唯一键'
) ENGINE=InnoDB AUTO_INCREMENT=10004 DEFAULT CHARSET=utf8mb3 COMMENT='车辆表';
CREATE TABLE `sw_product` (
  `product_id` bigint NOT NULL AUTO_INCREMENT,
  `product_name` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '产品名称',
  `product_key` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '产品key',
  `product_model` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '产品型号',
  `status` int NOT NULL DEFAULT '0' COMMENT '发布状态：0=未发布，1=已发布',
  `metadata` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci COMMENT '物模型',
  `metadata_table` int NOT NULL DEFAULT '0' COMMENT '是否生成物模型表：0=否，1=是',
  `create_time` datetime NOT NULL,
  `update_time` datetime DEFAULT NULL,
  `is_delete` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除：true=已删除，false=正常\n',
  PRIMARY KEY (`product_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10001 DEFAULT CHARSET=utf8mb3 COMMENT='电池表';
CREATE TABLE `sw_travel_record` (
  `travel_record_id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL COMMENT '用户ID',
  `device_id` bigint NOT NULL COMMENT '设备ID',
  `start_point` varchar(5) DEFAULT NULL COMMENT '起点',
  `end_point` varchar(20) DEFAULT NULL COMMENT '终点',
  `mileage_driven` int DEFAULT '0' COMMENT '行驶里程，单位（m）',
  `start_time` datetime DEFAULT NULL COMMENT '行程开始时间',
  `end_time` datetime DEFAULT NULL COMMENT '行程结束时间',
  `avg_speed` varchar(6) DEFAULT NULL COMMENT '平均时速，单位（m）',
  `consumption` int DEFAULT NULL COMMENT '使用电量',
  `create_time` datetime NOT NULL,
  `update_time` datetime DEFAULT NULL,
  `is_delete` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除：true=已删除，false=正常\n',
  PRIMARY KEY (`travel_record_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb3 COMMENT='行程记录表';
CREATE TABLE `sw_user` (
  `user_id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL,
  `password` varchar(100) NOT NULL,
  `pwd_salt` varchar(255) NOT NULL COMMENT '密码盐值',
  `pwd_encry_num` int NOT NULL DEFAULT '0' COMMENT '密码加密次数',
  `first_name` varchar(100) DEFAULT NULL,
  `last_name` varchar(100) DEFAULT NULL,
  `enable` bit(1) NOT NULL DEFAULT b'1' COMMENT '账号是否可用：true=正常，false=停用\n',
  `head_img` varchar(100) DEFAULT NULL,
  `mobile_phone` varchar(20) DEFAULT NULL COMMENT '手机号',
  `user_address` varchar(255) DEFAULT NULL COMMENT '用户地址',
  `create_time` datetime NOT NULL,
  `update_time` datetime DEFAULT NULL,
  `is_delete` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除：true=已删除，false=正常\n',
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `username_un_index` (`username`) USING BTREE COMMENT '用户名唯一键'
) ENGINE=InnoDB AUTO_INCREMENT=10003 DEFAULT CHARSET=utf8mb3 COMMENT='App用户表';
CREATE TABLE `sw_user_auth` (
  `auth_id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint DEFAULT NULL,
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
) ENGINE=InnoDB AUTO_INCREMENT=1813849066451243009 DEFAULT CHARSET=utf8mb3 COMMENT='实名认证表';
CREATE TABLE `sw_user_device` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `device_id` bigint NOT NULL,
  `user_id` bigint NOT NULL,
  `is_select` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否选定：false=未选定，true=已选定',
  `user_device_type` bit(1) NOT NULL DEFAULT b'0' COMMENT '设备用户类型：0=主用户，1=从用户',
  `mobile_key` bit(1) NOT NULL DEFAULT b'0' COMMENT '手机钥匙开关：false=关，true=开',
  `speed_limit` bit(1) NOT NULL DEFAULT b'0' COMMENT '速度限制开关：false=关，true=开',
  `driving_mode_type` int NOT NULL DEFAULT '0' COMMENT '驾驶模式：0=ECO模式，1=运动模式，2=狂暴模式，3=脱困模式',
  `energy_recovery` int NOT NULL DEFAULT '0' COMMENT '动能回收类型：0=无，1=中，2=强',
  `create_time` datetime NOT NULL,
  `update_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1808048639906418689 DEFAULT CHARSET=utf8mb3 COMMENT='用户车辆表';
CREATE TABLE `sw_version` (
  `version_id` bigint NOT NULL AUTO_INCREMENT,
  `version_code` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '版本编码',
  `version_type` int NOT NULL DEFAULT '0' COMMENT '版本类型：0=APP更新，1=固件升级',
  `version_update_type` int NOT NULL DEFAULT '0' COMMENT '版本更新类型：0=弱更新，1=强更新',
  `version_status` int NOT NULL DEFAULT '0' COMMENT '版本发布状态：0=待发布，1=已发布，2=已过期',
  `version_url` varchar(200) NOT NULL COMMENT '版本链接',
  `version_desc` varchar(1000) DEFAULT NULL COMMENT '版本描述，用于app显示的新版本信息',
  `create_time` datetime NOT NULL,
  `update_time` datetime DEFAULT NULL,
  `is_delete` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除：true=已删除，false=正常\n',
  PRIMARY KEY (`version_id`)
) ENGINE=InnoDB AUTO_INCREMENT=654654321365 DEFAULT CHARSET=utf8mb3 COMMENT='App版本管理表';
BEGIN;
LOCK TABLES `sviwo`.`sw_alarm_log` WRITE;
DELETE FROM `sviwo`.`sw_alarm_log`;
INSERT INTO `sviwo`.`sw_alarm_log` (`id`,`type`,`data`,`product_key`,`device_key`,`status`,`content`,`updated_by`,`created_time`,`updated_time`) VALUES (1, 1, '{\"Status\":\"offline\",\"CreateTime\":1719997312}', 'k0ugjmf1ois', 'sviwo-asidh342sjahdk', 0, '', NULL, '2024-07-03 17:03:00', NULL),(2, 1, '{\"Status\":\"offline\",\"CreateTime\":1719997312}', 'k0ugjmf1ois', 'sviwo-asidh342sjahdk', 0, '', NULL, '2024-07-03 17:03:00', NULL),(3, 1, '{\"Status\":\"offline\",\"CreateTime\":1719997312}', 'k0ugjmf1ois', 'sviwo-asidh342sjahdk', 0, '', NULL, '2024-07-03 17:03:00', NULL),(4, 1, '{\"Status\":\"online\",\"CreateTime\":1719997300}', 'k0ugjmf1ois', 'sviwo-asidh342sjahdk', 0, '', NULL, '2024-07-03 17:03:05', NULL),(5, 1, '{\"Status\":\"online\",\"CreateTime\":1719997300}', 'k0ugjmf1ois', 'sviwo-asidh342sjahdk', 0, '', NULL, '2024-07-03 17:03:06', NULL),(6, 1, '{\"Status\":\"online\",\"CreateTime\":1719997300}', 'k0ugjmf1ois', 'sviwo-asidh342sjahdk', 0, '', NULL, '2024-07-03 17:03:06', NULL),(7, 1, '{\"Status\":\"online\",\"CreateTime\":1719997392}', 'k0ugjmf1ois', 'sviwo-asidh342sjahdk', 0, '', NULL, '2024-07-03 17:03:13', NULL),(8, 1, '{\"Status\":\"offline\",\"CreateTime\":1719997401}', 'k0ugjmf1ois', 'sviwo-asidh342sjahdk', 0, '', NULL, '2024-07-03 17:03:22', NULL),(9, 1, '{\"Status\":\"online\",\"CreateTime\":1719997843}', 'k0ugjmf1ois', 'sviwo-asidh342sjahdk', 0, '', NULL, '2024-07-03 17:10:44', NULL),(10, 1, '{\"Status\":\"online\",\"CreateTime\":1719997843}', 'k0ugjmf1ois', 'sviwo-asidh342sjahdk', 0, '', NULL, '2024-07-03 17:10:44', NULL),(11, 1, '{\"Status\":\"online\",\"CreateTime\":1719997844}', 'k0ugjmf1ois', 'sviwo-asidh342sjahdk', 0, '', NULL, '2024-07-03 17:10:44', NULL),(12, 1, '{\"Status\":\"online\",\"CreateTime\":1719997844}', 'k0ugjmf1ois', 'sviwo-asidh342sjahdk', 0, '', NULL, '2024-07-03 17:10:45', NULL),(13, 1, '{\"Status\":\"online\",\"CreateTime\":1720001503}', 'k0ugjmf1ois', 'sviwo-asidh342sjahdk', 0, '', NULL, '2024-07-03 18:11:44', NULL),(14, 1, '{\"Status\":\"online\",\"CreateTime\":1720001503}', 'k0ugjmf1ois', 'sviwo-asidh342sjahdk', 0, '', NULL, '2024-07-03 18:11:44', NULL)
;
UNLOCK TABLES;
COMMIT;
BEGIN;
LOCK TABLES `sviwo`.`sw_app_media` WRITE;
DELETE FROM `sviwo`.`sw_app_media`;
INSERT INTO `sviwo`.`sw_app_media` (`id`,`parent_id`,`enable`,`page_type`,`display_type`,`title`,`media_desc`,`small_img`,`icon`,`content`,`orders`,`create_time`,`update_time`,`is_delete`) VALUES (679874621321, -1, b'1', 1, 0, '注册用户页', NULL, NULL, NULL, NULL, 1, '2024-07-18 18:53:56', NULL, b'0'),(679874621322, 679874621321, b'1', 1, 0, '用户协议', NULL, NULL, NULL, '用户协议用户协议用户协议用户协议用户协议用户协议用户协议用户协议', 1, '2024-07-18 18:53:56', NULL, b'0'),(679874621323, 679874621321, b'1', 1, 0, '隐私政策', NULL, NULL, NULL, '隐私政策隐私政策隐私政策隐私政策隐私政策隐私政策隐私政策隐私政策', 2, '2024-07-18 18:53:56', NULL, b'0'),(679874621324, -1, b'1', 0, 0, '帮助页', NULL, NULL, NULL, '', 2, '2024-07-18 18:53:56', NULL, b'0'),(679874621325, 679874621324, b'1', 0, 1, '车辆无法启动！', NULL, NULL, NULL, 'https://www.zhihu.com', 1, '2024-07-18 18:53:56', NULL, b'0'),(679874621326, 679874621324, b'1', 0, 0, '如何联系我们？', NULL, NULL, NULL, '<p style=\"text-align: center;\">\n    <span style=\"font-size: 60px;\"><strong>15881155743</strong></span>\n</p>', 2, '2024-07-18 18:53:56', NULL, b'0'),(679874621327, 679874621324, b'1', 0, 1, '车辆如何保养？', NULL, NULL, NULL, 'https://www.apple.com.cn', 3, '2024-07-18 18:53:56', NULL, b'0'),(679874621328, -1, b'1', 2, 0, '服务页', NULL, NULL, NULL, '', 3, '2024-07-18 18:53:56', NULL, b'0'),(679874621329, 679874621328, b'1', 2, 1, '视频教程', NULL, NULL, NULL, 'https://www.apple.com.cn', 1, '2024-07-18 18:53:56', NULL, b'0'),(6798746213210, 679874621328, b'1', 2, 0, '支持', NULL, NULL, NULL, '<p style=\"text-align: center;\">\n    <span style=\"font-size: 60px;\"><strong>15881155743</strong></span>\n</p>', 2, '2024-07-18 18:53:56', NULL, b'0')
;
UNLOCK TABLES;
COMMIT;
BEGIN;
LOCK TABLES `sviwo`.`sw_app_param` WRITE;
DELETE FROM `sviwo`.`sw_app_param`;
INSERT INTO `sviwo`.`sw_app_param` (`param_id`,`parent_id`,`param_name`,`param_value`,`param_const`,`create_time`,`update_time`,`is_delete`) VALUES (897621312654654546, -1, '服务电话', '15881155743', 'SERVICE_PHONE', '2024-07-02 14:51:33', NULL, b'0')
;
UNLOCK TABLES;
COMMIT;
BEGIN;
LOCK TABLES `sviwo`.`sw_device` WRITE;
DELETE FROM `sviwo`.`sw_device`;
INSERT INTO `sviwo`.`sw_device` (`device_id`,`product_id`,`sim_id`,`product_key`,`device_name`,`device_secret`,`device_model`,`nickname`,`status`,`online_timeout`,`activate_time`,`registry_time`,`version`,`last_online_time`,`metadata_table`,`bluetooth_address`,`bluetooth_secret_key`,`create_time`,`update_time`,`is_delete`) VALUES (10000, 10000, '111', 'k0ugjmf1ois', 'sviwo_atv', '91154cf19bea3f0c5b76f3ce312f948c', '111', 'SVIWO', 1, 10, '2024-02-28 14:44:19', '2024-02-28 14:44:19', NULL, NULL, b'1', 'kjhkh', 'jhjkh', '2024-02-28 14:44:19', NULL, b'0'),(10001, 10000, '111', 'k0ugjmf1ois', 'sviwo-asdas546a4s6d5', '92cbf83b2c083554f202b6d419f1f509', '111', 'SVIWO1', 1, 10, '2024-03-28 14:46:53', '2024-03-28 14:46:53', NULL, NULL, b'1', '', '', '2024-03-28 14:46:53', NULL, b'0'),(10002, 10000, '111', 'k0ugjmf1ois', 'sviwo-23kj4h2k3b4kk2', '63ab88874e683d02ceccff98015e0aff', '111', 'SVIWO2', 1, 10, '2024-05-28 15:06:46', '2024-05-28 15:06:46', NULL, NULL, b'1', '', '', '2024-05-28 15:06:46', NULL, b'0'),(10003, 10000, '89314404001001776573', 'k0ugjmf1ois', 'sviwo-asidh342sjahdk', '738f2cfe3b5220c69edf3575502bd479', '111', 'SVIWO3', 2, 10, '2024-06-04 14:58:17', '2024-06-04 14:58:17', NULL, '2024-07-03 18:11:44', b'1', '', '', '2024-06-04 14:58:17', NULL, b'0')
;
UNLOCK TABLES;
COMMIT;
BEGIN;
LOCK TABLES `sviwo`.`sw_product` WRITE;
DELETE FROM `sviwo`.`sw_product`;
INSERT INTO `sviwo`.`sw_product` (`product_id`,`product_name`,`product_key`,`product_model`,`status`,`metadata`,`metadata_table`,`create_time`,`update_time`,`is_delete`) VALUES (10000, 'Sviwo-ATV', 'k0ugjmf1ois', '111', 1, '{\"schema\":\"https://iotx-tsl.oss-ap-southeast-1.aliyuncs.com/schema.json\",\"profile\":{\"version\":\"1.0\",\"productKey\":\"k0ugjmf1ois\"},\"properties\":[{\"identifier\":\"VehSpeed\",\"name\":\"速度\",\"accessMode\":\"rw\",\"required\":false,\"dataType\":{\"type\":\"double\",\"specs\":{\"min\":\"0\",\"max\":\"300\",\"unit\":\"km/h\",\"unitName\":\"千米每小时\",\"step\":\"0.1\"}}},{\"identifier\":\"Mileage\",\"name\":\"行驶里程\",\"accessMode\":\"rw\",\"required\":false,\"dataType\":{\"type\":\"double\",\"specs\":{\"min\":\"0\",\"max\":\"1000000000\",\"unit\":\"km\",\"unitName\":\"千米\",\"step\":\"0.1\"}}},{\"identifier\":\"RotateSpeed\",\"name\":\"实际转速\",\"accessMode\":\"rw\",\"required\":false,\"dataType\":{\"type\":\"int\",\"specs\":{\"min\":\"0\",\"max\":\"100000\",\"unit\":\"r/min(rpm)\",\"step\":\"1\"}}},{\"identifier\":\"GeoLocation\",\"name\":\"地理位置\",\"accessMode\":\"rw\",\"desc\":\"{\\\"longitude\\\":104.066541, \\\"latitude\\\":30.572269}\",\"required\":false,\"dataType\":{\"type\":\"text\",\"specs\":{\"length\":\"10240\"}}},{\"identifier\":\"SerialNumber\",\"name\":\"产品编号\",\"accessMode\":\"rw\",\"required\":false,\"dataType\":{\"type\":\"text\",\"specs\":{\"length\":\"20\"}}},{\"identifier\":\"Product\",\"name\":\"产品名称\",\"accessMode\":\"rw\",\"required\":false,\"dataType\":{\"type\":\"text\",\"specs\":{\"length\":\"20\"}}},{\"identifier\":\"BatteryStatus\",\"name\":\"电池状态\",\"accessMode\":\"rw\",\"required\":false,\"dataType\":{\"type\":\"bool\",\"specs\":{\"0\":\"放电\",\"1\":\"充电\"}}},{\"identifier\":\"LockedStatus\",\"name\":\"锁车状态\",\"accessMode\":\"rw\",\"required\":false,\"dataType\":{\"type\":\"bool\",\"specs\":{\"0\":\"关机\",\"1\":\"开机\"}}},{\"identifier\":\"Electricity\",\"name\":\"电池电量\",\"accessMode\":\"rw\",\"required\":false,\"dataType\":{\"type\":\"int\",\"specs\":{\"min\":\"0\",\"max\":\"100\",\"unit\":\"%\",\"unitName\":\"百分比\",\"step\":\"1\"}}},{\"identifier\":\"RemainMile\",\"name\":\"剩余里程\",\"accessMode\":\"rw\",\"required\":false,\"dataType\":{\"type\":\"double\",\"specs\":{\"min\":\"0\",\"max\":\"100000\",\"unit\":\"km\",\"unitName\":\"千米\",\"step\":\"0.1\"}}},{\"identifier\":\"Fortification\",\"name\":\"设防\",\"accessMode\":\"rw\",\"required\":false,\"dataType\":{\"type\":\"bool\",\"specs\":{\"0\":\"关闭\",\"1\":\"开启\"}}},{\"identifier\":\"ProductModel\",\"name\":\"产品型号\",\"accessMode\":\"rw\",\"required\":false,\"dataType\":{\"type\":\"text\",\"specs\":{\"length\":\"20\"}}},{\"identifier\":\"Limit\",\"name\":\"速度阀值设置\",\"accessMode\":\"rw\",\"required\":false,\"dataType\":{\"type\":\"float\",\"specs\":{\"min\":\"0\",\"max\":\"60\",\"unit\":\"km/h\",\"unitName\":\"千米每小时\",\"step\":\"0.1\"}}},{\"identifier\":\"Light\",\"name\":\"灯状态开关\",\"accessMode\":\"rw\",\"desc\":\"0-关闭，1-闪烁， 2-常亮\",\"required\":false,\"dataType\":{\"type\":\"int\",\"specs\":{\"min\":\"0\",\"max\":\"2\",\"step\":\"1\"}}},{\"identifier\":\"PowerModel\",\"name\":\"动能模式切换\",\"accessMode\":\"rw\",\"desc\":\"1:ECO 2:运动 3:狂暴 \",\"required\":false,\"dataType\":{\"type\":\"int\",\"specs\":{\"min\":\"1\",\"max\":\"3\",\"step\":\"1\"}}},{\"identifier\":\"Energy\",\"name\":\"动能回收\",\"accessMode\":\"rw\",\"desc\":\"0（无），1 中，2 强\",\"required\":false,\"dataType\":{\"type\":\"int\",\"specs\":{\"min\":\"0\",\"max\":\"2\",\"step\":\"1\"}}},{\"identifier\":\"speaker\",\"name\":\"喇叭开关\",\"accessMode\":\"rw\",\"desc\":\"喇叭开关，接收指令后间歇性鸣笛10秒后自动取消。\",\"required\":false,\"dataType\":{\"type\":\"bool\",\"specs\":{\"0\":\"关\",\"1\":\"开\"}}}],\"events\":[{\"identifier\":\"post\",\"name\":\"post\",\"type\":\"info\",\"required\":true,\"desc\":\"属性上报\",\"method\":\"thing.event.property.post\",\"outputData\":[{\"identifier\":\"VehSpeed\",\"name\":\"速度\",\"dataType\":{\"type\":\"double\",\"specs\":{\"min\":\"0\",\"max\":\"300\",\"unit\":\"km/h\",\"unitName\":\"千米每小时\",\"step\":\"0.1\"}}},{\"identifier\":\"Mileage\",\"name\":\"行驶里程\",\"dataType\":{\"type\":\"double\",\"specs\":{\"min\":\"0\",\"max\":\"1000000000\",\"unit\":\"km\",\"unitName\":\"千米\",\"step\":\"0.1\"}}},{\"identifier\":\"RotateSpeed\",\"name\":\"实际转速\",\"dataType\":{\"type\":\"int\",\"specs\":{\"min\":\"0\",\"max\":\"100000\",\"unit\":\"r/min(rpm)\",\"step\":\"1\"}}},{\"identifier\":\"GeoLocation\",\"name\":\"地理位置\",\"dataType\":{\"type\":\"text\",\"specs\":{\"length\":\"10240\"}}},{\"identifier\":\"SerialNumber\",\"name\":\"产品编号\",\"dataType\":{\"type\":\"text\",\"specs\":{\"length\":\"20\"}}},{\"identifier\":\"Product\",\"name\":\"产品名称\",\"dataType\":{\"type\":\"text\",\"specs\":{\"length\":\"20\"}}},{\"identifier\":\"BatteryStatus\",\"name\":\"电池状态\",\"dataType\":{\"type\":\"bool\",\"specs\":{\"0\":\"放电\",\"1\":\"充电\"}}},{\"identifier\":\"LockedStatus\",\"name\":\"锁车状态\",\"dataType\":{\"type\":\"bool\",\"specs\":{\"0\":\"关机\",\"1\":\"开机\"}}},{\"identifier\":\"Electricity\",\"name\":\"电池电量\",\"dataType\":{\"type\":\"int\",\"specs\":{\"min\":\"0\",\"max\":\"100\",\"unit\":\"%\",\"unitName\":\"百分比\",\"step\":\"1\"}}},{\"identifier\":\"RemainMile\",\"name\":\"剩余里程\",\"dataType\":{\"type\":\"double\",\"specs\":{\"min\":\"0\",\"max\":\"100000\",\"unit\":\"km\",\"unitName\":\"千米\",\"step\":\"0.1\"}}},{\"identifier\":\"Fortification\",\"name\":\"设防\",\"dataType\":{\"type\":\"bool\",\"specs\":{\"0\":\"关闭\",\"1\":\"开启\"}}},{\"identifier\":\"ProductModel\",\"name\":\"产品型号\",\"dataType\":{\"type\":\"text\",\"specs\":{\"length\":\"20\"}}},{\"identifier\":\"Limit\",\"name\":\"速度阀值设置\",\"dataType\":{\"type\":\"float\",\"specs\":{\"min\":\"0\",\"max\":\"60\",\"unit\":\"km/h\",\"unitName\":\"千米每小时\",\"step\":\"0.1\"}}},{\"identifier\":\"Light\",\"name\":\"灯状态开关\",\"dataType\":{\"type\":\"int\",\"specs\":{\"min\":\"0\",\"max\":\"2\",\"step\":\"1\"}}},{\"identifier\":\"PowerModel\",\"name\":\"动能模式切换\",\"dataType\":{\"type\":\"int\",\"specs\":{\"min\":\"1\",\"max\":\"3\",\"step\":\"1\"}}},{\"identifier\":\"Energy\",\"name\":\"动能回收\",\"dataType\":{\"type\":\"int\",\"specs\":{\"min\":\"0\",\"max\":\"2\",\"step\":\"1\"}}},{\"identifier\":\"speaker\",\"name\":\"喇叭开关\",\"dataType\":{\"type\":\"bool\",\"specs\":{\"0\":\"关\",\"1\":\"开\"}}}]},{\"identifier\":\"lowbattey\",\"name\":\"电量低告警\",\"type\":\"alert\",\"required\":false,\"desc\":\"电量低报警\",\"method\":\"thing.event.lowbattey.post\",\"outputData\":[{\"identifier\":\"BatteryLevel\",\"name\":\"电量水平\",\"dataType\":{\"type\":\"int\",\"specs\":{\"min\":\"0\",\"max\":\"100\",\"step\":\"1\"}}}]},{\"identifier\":\"maxtemp\",\"name\":\"车机温度过高\",\"type\":\"alert\",\"required\":false,\"method\":\"thing.event.maxtemp.post\",\"outputData\":[{\"identifier\":\"temperature\",\"name\":\"温度\",\"dataType\":{\"type\":\"float\",\"specs\":{\"min\":\"0\",\"max\":\"100\",\"unit\":\"°C\",\"unitName\":\"摄氏度\",\"step\":\"0.1\"}}}]},{\"identifier\":\"concharger\",\"name\":\"充电枪连接\",\"type\":\"info\",\"required\":false,\"method\":\"thing.event.concharger.post\",\"outputData\":[{\"identifier\":\"status\",\"name\":\"连接状态\",\"dataType\":{\"type\":\"enum\",\"specs\":{\"0\":\"连接断开\",\"1\":\"连接成功\"}}}]},{\"identifier\":\"WheelAlarm\",\"name\":\"轮动告警\",\"type\":\"alert\",\"required\":false,\"method\":\"thing.event.WheelAlarm.post\",\"outputData\":[]},{\"identifier\":\"VibrationAlarm\",\"name\":\"震动告警\",\"type\":\"alert\",\"required\":false,\"method\":\"thing.event.VibrationAlarm.post\",\"outputData\":[]},{\"identifier\":\"Error\",\"name\":\"设备异常告警\",\"type\":\"alert\",\"required\":false,\"method\":\"thing.event.Error.post\",\"outputData\":[{\"identifier\":\"errorCode\",\"name\":\"异常类型\",\"dataType\":{\"type\":\"text\",\"specs\":{\"length\":\"255\"}}},{\"identifier\":\"error\",\"name\":\"异常详情\",\"dataType\":{\"type\":\"text\",\"specs\":{\"length\":\"255\"}}}]}],\"services\":[{\"identifier\":\"set\",\"name\":\"set\",\"required\":true,\"callType\":\"async\",\"desc\":\"属性设置\",\"method\":\"thing.service.property.set\",\"inputData\":[{\"identifier\":\"VehSpeed\",\"name\":\"速度\",\"dataType\":{\"type\":\"double\",\"specs\":{\"min\":\"0\",\"max\":\"300\",\"unit\":\"km/h\",\"unitName\":\"千米每小时\",\"step\":\"0.1\"}}},{\"identifier\":\"Mileage\",\"name\":\"行驶里程\",\"dataType\":{\"type\":\"double\",\"specs\":{\"min\":\"0\",\"max\":\"1000000000\",\"unit\":\"km\",\"unitName\":\"千米\",\"step\":\"0.1\"}}},{\"identifier\":\"RotateSpeed\",\"name\":\"实际转速\",\"dataType\":{\"type\":\"int\",\"specs\":{\"min\":\"0\",\"max\":\"100000\",\"unit\":\"r/min(rpm)\",\"step\":\"1\"}}},{\"identifier\":\"GeoLocation\",\"name\":\"地理位置\",\"dataType\":{\"type\":\"text\",\"specs\":{\"length\":\"10240\"}}},{\"identifier\":\"SerialNumber\",\"name\":\"产品编号\",\"dataType\":{\"type\":\"text\",\"specs\":{\"length\":\"20\"}}},{\"identifier\":\"Product\",\"name\":\"产品名称\",\"dataType\":{\"type\":\"text\",\"specs\":{\"length\":\"20\"}}},{\"identifier\":\"BatteryStatus\",\"name\":\"电池状态\",\"dataType\":{\"type\":\"bool\",\"specs\":{\"0\":\"放电\",\"1\":\"充电\"}}},{\"identifier\":\"LockedStatus\",\"name\":\"锁车状态\",\"dataType\":{\"type\":\"bool\",\"specs\":{\"0\":\"关机\",\"1\":\"开机\"}}},{\"identifier\":\"Electricity\",\"name\":\"电池电量\",\"dataType\":{\"type\":\"int\",\"specs\":{\"min\":\"0\",\"max\":\"100\",\"unit\":\"%\",\"unitName\":\"百分比\",\"step\":\"1\"}}},{\"identifier\":\"RemainMile\",\"name\":\"剩余里程\",\"dataType\":{\"type\":\"double\",\"specs\":{\"min\":\"0\",\"max\":\"100000\",\"unit\":\"km\",\"unitName\":\"千米\",\"step\":\"0.1\"}}},{\"identifier\":\"Fortification\",\"name\":\"设防\",\"dataType\":{\"type\":\"bool\",\"specs\":{\"0\":\"关闭\",\"1\":\"开启\"}}},{\"identifier\":\"ProductModel\",\"name\":\"产品型号\",\"dataType\":{\"type\":\"text\",\"specs\":{\"length\":\"20\"}}},{\"identifier\":\"Limit\",\"name\":\"速度阀值设置\",\"dataType\":{\"type\":\"float\",\"specs\":{\"min\":\"0\",\"max\":\"60\",\"unit\":\"km/h\",\"unitName\":\"千米每小时\",\"step\":\"0.1\"}}},{\"identifier\":\"Light\",\"name\":\"灯状态开关\",\"dataType\":{\"type\":\"int\",\"specs\":{\"min\":\"0\",\"max\":\"2\",\"step\":\"1\"}}},{\"identifier\":\"PowerModel\",\"name\":\"动能模式切换\",\"dataType\":{\"type\":\"int\",\"specs\":{\"min\":\"1\",\"max\":\"3\",\"step\":\"1\"}}},{\"identifier\":\"Energy\",\"name\":\"动能回收\",\"dataType\":{\"type\":\"int\",\"specs\":{\"min\":\"0\",\"max\":\"2\",\"step\":\"1\"}}},{\"identifier\":\"speaker\",\"name\":\"喇叭开关\",\"dataType\":{\"type\":\"bool\",\"specs\":{\"0\":\"关\",\"1\":\"开\"}}}],\"outputData\":[]},{\"identifier\":\"get\",\"name\":\"get\",\"required\":true,\"callType\":\"async\",\"desc\":\"属性获取\",\"method\":\"thing.service.property.get\",\"inputData\":[\"VehSpeed\",\"Mileage\",\"RotateSpeed\",\"GeoLocation\",\"SerialNumber\",\"Product\",\"BatteryStatus\",\"LockedStatus\",\"Electricity\",\"RemainMile\",\"Fortification\",\"ProductModel\",\"Limit\",\"Light\",\"PowerModel\",\"Energy\",\"speaker\"],\"outputData\":[{\"identifier\":\"VehSpeed\",\"name\":\"速度\",\"dataType\":{\"type\":\"double\",\"specs\":{\"min\":\"0\",\"max\":\"300\",\"unit\":\"km/h\",\"unitName\":\"千米每小时\",\"step\":\"0.1\"}}},{\"identifier\":\"Mileage\",\"name\":\"行驶里程\",\"dataType\":{\"type\":\"double\",\"specs\":{\"min\":\"0\",\"max\":\"1000000000\",\"unit\":\"km\",\"unitName\":\"千米\",\"step\":\"0.1\"}}},{\"identifier\":\"RotateSpeed\",\"name\":\"实际转速\",\"dataType\":{\"type\":\"int\",\"specs\":{\"min\":\"0\",\"max\":\"100000\",\"unit\":\"r/min(rpm)\",\"step\":\"1\"}}},{\"identifier\":\"GeoLocation\",\"name\":\"地理位置\",\"dataType\":{\"type\":\"text\",\"specs\":{\"length\":\"10240\"}}},{\"identifier\":\"SerialNumber\",\"name\":\"产品编号\",\"dataType\":{\"type\":\"text\",\"specs\":{\"length\":\"20\"}}},{\"identifier\":\"Product\",\"name\":\"产品名称\",\"dataType\":{\"type\":\"text\",\"specs\":{\"length\":\"20\"}}},{\"identifier\":\"BatteryStatus\",\"name\":\"电池状态\",\"dataType\":{\"type\":\"bool\",\"specs\":{\"0\":\"放电\",\"1\":\"充电\"}}},{\"identifier\":\"LockedStatus\",\"name\":\"锁车状态\",\"dataType\":{\"type\":\"bool\",\"specs\":{\"0\":\"关机\",\"1\":\"开机\"}}},{\"identifier\":\"Electricity\",\"name\":\"电池电量\",\"dataType\":{\"type\":\"int\",\"specs\":{\"min\":\"0\",\"max\":\"100\",\"unit\":\"%\",\"unitName\":\"百分比\",\"step\":\"1\"}}},{\"identifier\":\"RemainMile\",\"name\":\"剩余里程\",\"dataType\":{\"type\":\"double\",\"specs\":{\"min\":\"0\",\"max\":\"100000\",\"unit\":\"km\",\"unitName\":\"千米\",\"step\":\"0.1\"}}},{\"identifier\":\"Fortification\",\"name\":\"设防\",\"dataType\":{\"type\":\"bool\",\"specs\":{\"0\":\"关闭\",\"1\":\"开启\"}}},{\"identifier\":\"ProductModel\",\"name\":\"产品型号\",\"dataType\":{\"type\":\"text\",\"specs\":{\"length\":\"20\"}}},{\"identifier\":\"Limit\",\"name\":\"速度阀值设置\",\"dataType\":{\"type\":\"float\",\"specs\":{\"min\":\"0\",\"max\":\"60\",\"unit\":\"km/h\",\"unitName\":\"千米每小时\",\"step\":\"0.1\"}}},{\"identifier\":\"Light\",\"name\":\"灯状态开关\",\"dataType\":{\"type\":\"int\",\"specs\":{\"min\":\"0\",\"max\":\"2\",\"step\":\"1\"}}},{\"identifier\":\"PowerModel\",\"name\":\"动能模式切换\",\"dataType\":{\"type\":\"int\",\"specs\":{\"min\":\"1\",\"max\":\"3\",\"step\":\"1\"}}},{\"identifier\":\"Energy\",\"name\":\"动能回收\",\"dataType\":{\"type\":\"int\",\"specs\":{\"min\":\"0\",\"max\":\"2\",\"step\":\"1\"}}},{\"identifier\":\"speaker\",\"name\":\"喇叭开关\",\"dataType\":{\"type\":\"bool\",\"specs\":{\"0\":\"关\",\"1\":\"开\"}}}]}]}', 1, '2024-07-02 15:03:38', NULL, b'0')
;
UNLOCK TABLES;
COMMIT;
BEGIN;
LOCK TABLES `sviwo`.`sw_travel_record` WRITE;
DELETE FROM `sviwo`.`sw_travel_record`;
UNLOCK TABLES;
COMMIT;
BEGIN;
LOCK TABLES `sviwo`.`sw_user` WRITE;
DELETE FROM `sviwo`.`sw_user`;
INSERT INTO `sviwo`.`sw_user` (`user_id`,`username`,`password`,`pwd_salt`,`pwd_encry_num`,`first_name`,`last_name`,`enable`,`head_img`,`mobile_phone`,`user_address`,`create_time`,`update_time`,`is_delete`) VALUES (10000, '821317143@qq.com', '7d1ca1fdbecbeb53db69d7e04d58df96', 'NFRWBJdOLf', 20, 'Gou', 'YiChuan', b'1', '20240718/e78iSrTYruyfg14ZjgYBXLVEM4XKwite.jpg', '15928736853', '成都市都江堰市', '2024-02-29 19:02:58', '2024-07-18 18:44:00', b'0'),(10001, 'gavinegaowen@gmail.com', '005e38f059ea9bbd291bab5037304481', '6FBX2IrV2w', 27, 'hCb', 'nkW', b'1', '', '', '', '2024-07-18 14:30:58', NULL, b'0'),(10002, '332669440@qq.com', '2eee23f5c80b35fe93073a48f34da39f', 'bbp78OIdhU', 26, 'yIx', 'bEO', b'1', '', '', '', '2024-07-18 16:11:32', NULL, b'0')
;
UNLOCK TABLES;
COMMIT;
BEGIN;
LOCK TABLES `sviwo`.`sw_user_auth` WRITE;
DELETE FROM `sviwo`.`sw_user_auth`;
INSERT INTO `sviwo`.`sw_user_auth` (`auth_id`,`user_id`,`auth_first_name`,`auth_last_name`,`certificate_front_img`,`certificate_back_img`,`auth_status`,`auth_fail_reason`,`auth_time`,`verify_time`,`create_time`,`update_time`) VALUES (1763157909782401024, 10000, '', '', '', '', 2, '', NULL, NULL, '2024-02-29 19:02:58', NULL),(1813823758553583616, 10001, '', '', '', '', 2, '', NULL, NULL, '2024-07-18 14:30:58', NULL),(1813849066451243008, 10002, '', '', '', '', 2, '', NULL, NULL, '2024-07-18 16:11:32', NULL)
;
UNLOCK TABLES;
COMMIT;
BEGIN;
LOCK TABLES `sviwo`.`sw_user_device` WRITE;
DELETE FROM `sviwo`.`sw_user_device`;
INSERT INTO `sviwo`.`sw_user_device` (`id`,`device_id`,`user_id`,`is_select`,`user_device_type`,`mobile_key`,`speed_limit`,`driving_mode_type`,`energy_recovery`,`create_time`,`update_time`) VALUES (1, 10003, 10000, b'0', b'0', b'0', b'0', 0, 0, '2024-07-18 16:14:26', '2024-07-18 16:14:29')
;
UNLOCK TABLES;
COMMIT;
BEGIN;
LOCK TABLES `sviwo`.`sw_version` WRITE;
DELETE FROM `sviwo`.`sw_version`;
INSERT INTO `sviwo`.`sw_version` (`version_id`,`version_code`,`version_type`,`version_update_type`,`version_status`,`version_url`,`version_desc`,`create_time`,`update_time`,`is_delete`) VALUES (654654321361, 'V1.0.0', 1, 0, 1, 'asdasdasd', 'sdasdas', '2024-02-22 16:21:02', NULL, b'1'),(654654321362, 'V1.0.1', 1, 0, 1, 'asdasdasd', 'sdasdas', '2024-02-22 16:21:02', NULL, b'1'),(654654321363, 'V1.0.1', 0, 0, 1, 'asdasdasd', 'sdasdas', '2024-02-22 16:21:02', NULL, b'1'),(654654321364, 'V1.0.0', 0, 0, 1, 'asdasdasd', 'sdasdas', '2024-02-22 16:21:02', NULL, b'1')
;
UNLOCK TABLES;
COMMIT;
