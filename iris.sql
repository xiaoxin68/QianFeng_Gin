/*
Navicat MySQL Data Transfer

Source Server         : zhang
Source Server Version : 50717
Source Host           : localhost:3306
Source Database       : iris

Target Server Type    : MYSQL
Target Server Version : 50717
File Encoding         : 65001

Date: 2020-06-19 15:54:07
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for admin
-- ----------------------------
DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `admin_name` varchar(32) DEFAULT NULL,
  `create_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE CURRENT_TIMESTAMP,
  `status` int(11) NOT NULL DEFAULT '0',
  `avatar` varchar(255) DEFAULT NULL,
  `pwd` varchar(255) DEFAULT NULL,
  `city_name` varchar(12) DEFAULT NULL,
  `city_id` int(11) DEFAULT NULL,
  `admin_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `IDX_admin_city_id` (`city_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of admin
-- ----------------------------
INSERT INTO `admin` VALUES ('1', 'davie', '2020-05-03 21:36:10', '0', '', '123', '湖北武汉', '1', null);
INSERT INTO `admin` VALUES ('2', 'lili', '2020-05-21 17:05:36', '2', '2', '123', '湖北随州', '2', null);
INSERT INTO `admin` VALUES ('3', 'tom', '2020-05-03 21:36:16', '0', '2', '123', '湖北随州', '2', null);

-- ----------------------------
-- Table structure for city
-- ----------------------------
DROP TABLE IF EXISTS `city`;
CREATE TABLE `city` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(24) DEFAULT NULL,
  `pin_yin` varchar(32) DEFAULT NULL,
  `longitude` float(32,0) DEFAULT NULL,
  `latitude` float(32,0) DEFAULT NULL,
  `area_code` varchar(6) DEFAULT NULL,
  `abbr` varchar(12) DEFAULT NULL,
  `city_id` bigint(20) DEFAULT NULL,
  `city_name` varchar(12) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of city
-- ----------------------------

-- ----------------------------
-- Table structure for food
-- ----------------------------
DROP TABLE IF EXISTS `food`;
CREATE TABLE `food` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `rating` int(11) DEFAULT NULL,
  `month_sales` int(11) DEFAULT NULL,
  `image_path` varchar(255) DEFAULT NULL,
  `activity` varchar(255) DEFAULT NULL,
  `attributes` varchar(255) DEFAULT NULL,
  `specs` varchar(255) DEFAULT NULL,
  `category_id` int(11) DEFAULT NULL,
  `restaurant_id` int(11) DEFAULT NULL,
  `del_flag` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `IDX_food_category_id` (`category_id`),
  KEY `IDX_food_restaurant_id` (`restaurant_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of food
-- ----------------------------
INSERT INTO `food` VALUES ('1', 'wqe', 'qwe', '1', '1', 'das', 'd', 'df', 'df', '12', '1', '1');
INSERT INTO `food` VALUES ('2', 'wqe', 'qwe', '1', '1', 'das', 'd', 'df', 'df', '12', '1', '0');
INSERT INTO `food` VALUES ('3', 'wqe', 'qwe', '1', '1', 'das', 'd', 'df', 'df', '12', '1', '0');
INSERT INTO `food` VALUES ('4', 'lallalal', '123', '0', '0', 'dfds', '2', '', '', '2', '0', '0');
INSERT INTO `food` VALUES ('5', 'asfsdf', '123', '0', '0', 'dfds', '2', '', '', '2', '0', '0');
INSERT INTO `food` VALUES ('6', '张潇潇', '123', '0', '0', 'dfds', '2', '', '', '2', '0', '0');
INSERT INTO `food` VALUES ('7', '张潇潇', '123', '0', '0', 'dfds', '2', '', '', '2', '0', '0');

-- ----------------------------
-- Table structure for food_category
-- ----------------------------
DROP TABLE IF EXISTS `food_category`;
CREATE TABLE `food_category` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `category_name` varchar(255) DEFAULT NULL,
  `category_desc` varchar(255) DEFAULT NULL,
  `level` int(11) DEFAULT NULL,
  `parent_category_id` int(11) DEFAULT NULL,
  `restaurant_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `IDX_food_category_restaurant_id` (`restaurant_id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of food_category
-- ----------------------------
INSERT INTO `food_category` VALUES ('1', '啦啦啦', '的贺卡上GKAS', '1', null, '1');
INSERT INTO `food_category` VALUES ('2', 'DSASDASD', 'SADSADSA', '2', '1', '2');
INSERT INTO `food_category` VALUES ('3', '324', '432', '43', '2', '2');
INSERT INTO `food_category` VALUES ('4', 'SdSA', 'SAD', '2', '1', '1');
INSERT INTO `food_category` VALUES ('5', '张三', '描述来说老师理论上', '1', '0', '1');
INSERT INTO `food_category` VALUES ('6', '张三', '描述来说老师理论上', '1', '0', '1');
INSERT INTO `food_category` VALUES ('7', '张三', '描述来说老师理论上', '1', '0', '1');
INSERT INTO `food_category` VALUES ('8', 'davie', '123', '1', '0', '0');
INSERT INTO `food_category` VALUES ('9', 'davie', '123', '1', '0', '0');
INSERT INTO `food_category` VALUES ('10', 'lallalal', '123', '1', '0', '0');
INSERT INTO `food_category` VALUES ('11', 'lallalal', '123', '1', '0', '0');
INSERT INTO `food_category` VALUES ('12', 'lallalal', '123', '1', '0', '0');
INSERT INTO `food_category` VALUES ('13', 'lallalal', '123', '1', '0', '0');

-- ----------------------------
-- Table structure for permission
-- ----------------------------
DROP TABLE IF EXISTS `permission`;
CREATE TABLE `permission` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `level` varchar(32) DEFAULT NULL,
  `name` varchar(32) DEFAULT NULL,
  `permission_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of permission
-- ----------------------------

-- ----------------------------
-- Table structure for person
-- ----------------------------
DROP TABLE IF EXISTS `person`;
CREATE TABLE `person` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(12) NOT NULL,
  `age` int(11) DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of person
-- ----------------------------
INSERT INTO `person` VALUES ('1', 'Lily', '15');
INSERT INTO `person` VALUES ('2', 'tom', '18');
INSERT INTO `person` VALUES ('3', '张三', '15');
INSERT INTO `person` VALUES ('4', '张三', '16');
INSERT INTO `person` VALUES ('5', '张三', '117');

-- ----------------------------
-- Table structure for shop
-- ----------------------------
DROP TABLE IF EXISTS `shop`;
CREATE TABLE `shop` (
  `shop_id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(32) DEFAULT NULL,
  `address` varchar(128) DEFAULT NULL,
  `latitude` float(32,0) DEFAULT NULL,
  `longitude` float(32,0) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `phone` varchar(255) DEFAULT NULL,
  `promotion_info` varchar(255) DEFAULT NULL,
  `float_delivery_fee` int(11) DEFAULT NULL,
  `float_minimum_order_amount` int(11) DEFAULT NULL,
  `is_premium` varchar(10) DEFAULT NULL,
  `delivery_mode` varchar(10) DEFAULT NULL,
  `new` varchar(10) DEFAULT NULL,
  `bao` varchar(10) DEFAULT NULL,
  `zhun` varchar(10) DEFAULT NULL,
  `piao` varchar(10) DEFAULT NULL,
  `start_time` varchar(255) DEFAULT NULL,
  `end_time` varchar(255) DEFAULT NULL,
  `image_path` varchar(255) DEFAULT NULL,
  `business_license_image` varchar(255) DEFAULT NULL,
  `catering_service_license_image` varchar(255) DEFAULT NULL,
  `category` varchar(255) DEFAULT NULL,
  `status` int(11) DEFAULT NULL,
  `recent_order_num` int(11) DEFAULT NULL,
  `rating_count` int(11) DEFAULT NULL,
  `rating` int(11) DEFAULT NULL,
  `dele` int(11) DEFAULT NULL,
  PRIMARY KEY (`shop_id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of shop
-- ----------------------------
INSERT INTO `shop` VALUES ('1', 'qweqw', 'qw', '324', '324', 'qwe', 'wqe', 'qwe', '213', '2', 'true', 'false', 'false', 'true', 'true', 'true', 'sdgfds', 'sssssssdg', 'sgd', 'sgd', 'sg', 's2d', '0', '0', '0', '0', '1');
INSERT INTO `shop` VALUES ('2', 'wqeewq', 'qw', '324', '324', 'qwe', 'wqe', 'qwe', '213', '2', 'true', 'false', 'false', 'true', 'true', 'true', 'sdgfds', 'sssssssdg', 'sgd', 'sgd', 'sg', 's2d', '0', '0', '0', '0', '1');
INSERT INTO `shop` VALUES ('3', 'ewqwqeweq', 'qw', '324', '324', 'qwe', 'wqe', 'qwe', '213', '2', 'true', 'false', 'false', 'true', 'true', 'true', 'sdgfds', 'sssssssdg', 'sgd', 'sgd', 'sg', 's2d', '0', '0', '0', '0', '0');
INSERT INTO `shop` VALUES ('4', 'xzzdx', 'qw', '324', '324', 'qwe', 'wqe', 'qwe', '213', '2', 'true', 'false', 'false', 'true', 'true', 'true', 'sdgfds', 'sssssssdg', 'sgd', 'sgd', 'sg', 's2d', '0', '0', '0', '0', '0');
INSERT INTO `shop` VALUES ('5', 'hdf', 'qw', '324', '324', 'qwe', 'wqe', 'qwe', '213', '2', 'true', 'false', 'false', 'true', 'true', 'true', 'sdgfds', 'sssssssdg', 'sgd', 'sgd', 'sg', 's2d', '0', '0', '0', '0', '1');
INSERT INTO `shop` VALUES ('6', 'ewqdfuy', 'qw', '324', '324', 'qwe', 'wqe', 'qwe', '213', '2', 'true', 'false', 'false', 'true', 'true', 'true', 'sdgfds', 'sssssssdg', 'sgd', 'sgd', 'sg', 's2d', '0', '0', '0', '0', '0');

-- ----------------------------
-- Table structure for sms_code
-- ----------------------------
DROP TABLE IF EXISTS `sms_code`;
CREATE TABLE `sms_code` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `phone` varchar(11) DEFAULT NULL,
  `biz_id` varchar(30) DEFAULT NULL,
  `code` varchar(4) DEFAULT NULL,
  `create_time` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sms_code
-- ----------------------------
INSERT INTO `sms_code` VALUES ('1', '123', '21', '213', '213');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(12) DEFAULT NULL,
  `register_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE CURRENT_TIMESTAMP,
  `mobile` varchar(11) DEFAULT NULL,
  `is_active` int(11) NOT NULL DEFAULT '0',
  `balance` int(11) NOT NULL DEFAULT '0',
  `avatar` varchar(255) DEFAULT NULL,
  `pwd` varchar(255) DEFAULT NULL,
  `del_flag` int(11) NOT NULL DEFAULT '0',
  `city_name` varchar(24) DEFAULT NULL,
  `user_name` varchar(12) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('1', '张三', '0000-00-00 00:00:00', '123456', '0', '0', '啦啦啦啦', '123', '0', '武汉', null);
INSERT INTO `user` VALUES ('2', '李四', '0000-00-00 00:00:00', '1232132', '0', '0', '哈哈哈哈', '123', '0', '随州', null);
INSERT INTO `user` VALUES ('3', '王二', '0000-00-00 00:00:00', '1232', '0', '0', '嘻嘻嘻', '1213', '1', '襄阳', null);
INSERT INTO `user` VALUES ('4', '12321', '2020-05-04 22:07:56', '123', '0', '0', '123', '231', '0', '123分、', null);
INSERT INTO `user` VALUES ('5', '241123', '2020-05-04 22:07:58', '3412', '0', '0', '423', '2431', '0', '分', null);
INSERT INTO `user` VALUES ('6', '三大·', '2020-05-04 22:07:59', 'asd', '0', '0', '阿萨德', '23123', '0', '抢答', null);
INSERT INTO `user` VALUES ('7', '撒啊啊啊啊啊', '2020-05-04 22:08:01', 'AX', '0', '0', 'ASD', 'ASF', '0', '凄凄切切群', null);
INSERT INTO `user` VALUES ('8', '士大夫', '2020-05-04 22:08:04', '1232131', '0', '0', '的人无', 'dtw', '0', '王企鹅', null);

-- ----------------------------
-- Table structure for user_infos
-- ----------------------------
DROP TABLE IF EXISTS `user_infos`;
CREATE TABLE `user_infos` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(16) DEFAULT NULL,
  `gender` varchar(32) DEFAULT NULL,
  `hobby` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user_infos
-- ----------------------------
INSERT INTO `user_infos` VALUES ('2', '沙河娜扎', '女', '足球');
