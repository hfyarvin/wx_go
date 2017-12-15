/*
Navicat MySQL Data Transfer

Source Server         : mine
Source Server Version : 50720
Source Host           : localhost:3306
Source Database       : maxiiot_test

Target Server Type    : MYSQL
Target Server Version : 50720
File Encoding         : 65001

Date: 2017-12-13 10:48:47
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for maxiiot_accounts
-- ----------------------------
DROP TABLE IF EXISTS `maxiiot_accounts`;
CREATE TABLE `maxiiot_accounts` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `balance` double NOT NULL,
  `version` int(11) DEFAULT '1',
  `created` datetime NOT NULL,
  `updated` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for maxiiot_balances
-- ----------------------------
DROP TABLE IF EXISTS `maxiiot_balances`;
CREATE TABLE `maxiiot_balances` (
  `id` bigint(20) NOT NULL,
  `order_id` varchar(255) NOT NULL,
  `amount` bigint(20) DEFAULT NULL,
  `status` tinyint(4) DEFAULT '0',
  `created` datetime DEFAULT NULL,
  `updated` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for maxiiot_billing_invoices
-- ----------------------------
DROP TABLE IF EXISTS `maxiiot_billing_invoices`;
CREATE TABLE `maxiiot_billing_invoices` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `total_price` bigint(20) NOT NULL,
  `status` tinyint(1) DEFAULT '1',
  `order_id` varchar(255) DEFAULT NULL,
  `expired` datetime NOT NULL,
  `created` datetime NOT NULL,
  `updated` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for maxiiot_support_announcements
-- ----------------------------
DROP TABLE IF EXISTS `maxiiot_support_announcements`;
CREATE TABLE `maxiiot_support_announcements` (
  `id` bigint(20) NOT NULL DEFAULT '1',
  `title` varchar(100) NOT NULL DEFAULT 'ss',
  `content` text NOT NULL,
  `click_times` int(11) DEFAULT '1',
  `created` datetime NOT NULL,
  `updated` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for maxiiot_support_articles
-- ----------------------------
DROP TABLE IF EXISTS `maxiiot_support_articles`;
CREATE TABLE `maxiiot_support_articles` (
  `id` bigint(20) NOT NULL,
  `article_cat_id` bigint(20) DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `content` text,
  `click_times` int(11) DEFAULT '0',
  `created` datetime NOT NULL,
  `updated` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for maxiiot_support_article_cats
-- ----------------------------
DROP TABLE IF EXISTS `maxiiot_support_article_cats`;
CREATE TABLE `maxiiot_support_article_cats` (
  `id` bigint(20) NOT NULL,
  `titlt` varchar(255) DEFAULT NULL,
  `created` datetime NOT NULL,
  `updated` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for maxiiot_support_downloads
-- ----------------------------
DROP TABLE IF EXISTS `maxiiot_support_downloads`;
CREATE TABLE `maxiiot_support_downloads` (
  `id` bigint(20) NOT NULL,
  `url` varchar(255) DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `created` datetime NOT NULL,
  `updated` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for maxiiot_support_tickets
-- ----------------------------
DROP TABLE IF EXISTS `maxiiot_support_tickets`;
CREATE TABLE `maxiiot_support_tickets` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL,
  `department` varchar(100) NOT NULL,
  `subject` varchar(255) DEFAULT NULL,
  `status` tinyint(1) DEFAULT NULL,
  `related_service_id` bigint(20) DEFAULT NULL,
  `priority` tinyint(1) DEFAULT NULL,
  `message` text,
  `attachments` varchar(255) DEFAULT NULL,
  `created` datetime NOT NULL,
  `updated` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for maxiiot_support_ticket_attachments
-- ----------------------------
DROP TABLE IF EXISTS `maxiiot_support_ticket_attachments`;
CREATE TABLE `maxiiot_support_ticket_attachments` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `url` varchar(255) NOT NULL,
  `ticket_id` bigint(20) NOT NULL,
  `created` datetime NOT NULL,
  `updated` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;

-- ----------------------------
-- Table structure for maxiiot_users
-- ----------------------------
DROP TABLE IF EXISTS `maxiiot_users`;
CREATE TABLE `maxiiot_users` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `sex` tinyint(1) DEFAULT NULL,
  `created` datetime NOT NULL,
  `updated` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `first_name` varchar(50) CHARACTER SET utf8 NOT NULL,
  `last_name` varchar(50) CHARACTER SET utf8 NOT NULL,
  `password_hash` varchar(200) NOT NULL,
  `email` varchar(100) NOT NULL,
  `phone` varchar(30) NOT NULL,
  `company` varchar(200) CHARACTER SET utf8 DEFAULT NULL,
  `address` varchar(400) CHARACTER SET utf8 DEFAULT NULL,
  `city` varchar(100) CHARACTER SET utf8 DEFAULT NULL,
  `regin` varchar(100) CHARACTER SET utf8 DEFAULT NULL,
  `zip_code` varchar(30) DEFAULT NULL,
  `country` varchar(100) CHARACTER SET utf8 DEFAULT NULL,
  `identity_card` varchar(50) DEFAULT NULL,
  `is_active` bit(1) DEFAULT NULL,
  `created` datetime NOT NULL,
  `updated` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=latin1;

-- ----------------------------
-- Table structure for users_contact
-- ----------------------------
DROP TABLE IF EXISTS `users_contact`;
CREATE TABLE `users_contact` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL,
  `first_name` varchar(100) CHARACTER SET utf8 NOT NULL,
  `last_name` varchar(100) CHARACTER SET utf8 NOT NULL,
  `company` varchar(200) CHARACTER SET utf8 NOT NULL,
  `email` varchar(50) NOT NULL,
  `phone` varchar(30) DEFAULT NULL,
  `address1` varchar(200) CHARACTER SET utf8 DEFAULT NULL,
  `address2` varchar(200) CHARACTER SET utf8 DEFAULT NULL,
  `city` varchar(50) CHARACTER SET utf8 DEFAULT NULL,
  `region` varchar(100) CHARACTER SET utf8 DEFAULT NULL,
  `zip_code` varchar(30) DEFAULT NULL,
  `country` varchar(100) CHARACTER SET utf8 DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;
