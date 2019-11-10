/*
Navicat MariaDB Data Transfer

Source Server         : 127.0.0.1
Source Server Version : 100314
Source Host           : 127.0.0.1:3306
Source Database       : test

Target Server Type    : MariaDB
Target Server Version : 100314
File Encoding         : 65001

Date: 2019-11-10 16:39:25
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for info
-- ----------------------------
DROP TABLE IF EXISTS `info`;
CREATE TABLE `info` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `path` varchar(255) DEFAULT NULL,
  `note` varchar(255) DEFAULT NULL,
  `unix` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of info
-- ----------------------------
INSERT INTO `info` VALUES ('2', '0.png', '/static/2019-11-10160029.png', '测试', '1573372829');
INSERT INTO `info` VALUES ('4', 'PS-key_Blue.jpg', '/static/2019-11-10162752.jpg', 'logo', '1573374472');
