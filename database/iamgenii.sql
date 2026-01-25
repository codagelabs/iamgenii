CREATE DATABASE  IF NOT EXISTS `u156722531_iamgenii` /*!40100 DEFAULT CHARACTER SET latin1 */;
USE `u156722531_iamgenii`;
-- MySQL dump 10.13  Distrib 5.7.29, for Linux (x86_64)
--
-- Host: localhost    Database: u156722531_iamgenii
-- ------------------------------------------------------
-- Server version	5.7.29-0ubuntu0.18.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `action_recorder`
--

--
-- Table structure for table `address_book`
--

DROP TABLE IF EXISTS `address_book`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `address_book` (
  `address_book_id` int(11) NOT NULL AUTO_INCREMENT,
  `customers_id` int(11) NOT NULL,
  `entry_gender` char(1) COLLATE utf8_unicode_ci DEFAULT NULL,
  `entry_company` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `entry_firstname` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `entry_lastname` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `entry_street_address` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `entry_suburb` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `entry_postcode` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `entry_city` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `entry_state` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `entry_country_id` int(11) NOT NULL DEFAULT '0',
  `entry_zone_id` int(11) NOT NULL DEFAULT '0',
  `entry_locality` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `entry_address_id` int(10) NOT NULL,
  `entry_phone` bigint(10) NOT NULL,
  `entry_delivery` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT  CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL ,
  PRIMARY KEY (`address_book_id`),
  KEY `idx_address_book_customers_id` (`customers_id`),
  FOREIGN KEY(`customers_id`)REFERENCES customers(`customers_id`) 
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `address_book`
--

LOCK TABLES `address_book` WRITE;
/*!40000 ALTER TABLE `address_book` DISABLE KEYS */;
/*!40000 ALTER TABLE `address_book` ENABLE KEYS */;
UNLOCK TABLES;


--
-- Table structure for table `admin_types`
--

DROP TABLE IF EXISTS `admin_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `admin_types` (
  `admin_type_id` int(11) NOT NULL AUTO_INCREMENT,
  `admin_type_name` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT  CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL ,
  `is_active` tinyint(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`admin_type_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `admin_types`
--

LOCK TABLES `admin_types` WRITE;
/*!40000 ALTER TABLE `admin_types` DISABLE KEYS */;
INSERT INTO `admin_types`(`admin_type_id`, `admin_type_name`, `created_at`, `updated_at`,`deleted_at`, `is_active`)
VALUES(1, 'super admin', CURRENT_TIMESTAMP ,CURRENT_TIMESTAMP , NULL , 1);
/*!40000 ALTER TABLE `admin_types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `administrators`
--

DROP TABLE IF EXISTS `administrators`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `administrators` (
  `administrators_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `first_name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `last_name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `email` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `password` varchar(1000) COLLATE utf8_unicode_ci NOT NULL,
  `is_active` tinyint(1) NOT NULL DEFAULT '0',
  `address` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `city` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `state` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `zip` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `country` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `phone` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `image` text COLLATE utf8_unicode_ci NOT NULL,
  `admin_type_id` int(11) NOT NULL,
  `remember_token` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL ,
  `created_at` timestamp NOT NULL DEFAULT  CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL ,
  PRIMARY KEY (`administrators_id`),
  UNIQUE KEY `administrators_user_name_unique` (`username`),
  UNIQUE KEY `administrators_email_unique` (`email`),
  UNIQUE KEY `administrators_phone_unique` (`phone`),
  FOREIGN KEY(`admin_type_id`)REFERENCES admin_types(`admin_type_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `administrators`
--

LOCK TABLES `administrators` WRITE;
/*!40000 ALTER TABLE `administrators` DISABLE KEYS */;
/*!40000 ALTER TABLE `administrators` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `api_calls_list`
--

DROP TABLE IF EXISTS `api_calls_list`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `api_calls_list` (
  `id` int(100) NOT NULL AUTO_INCREMENT,
  `nonce` text NOT NULL,
  `url` varchar(64) NOT NULL,
  `device_id` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `api_calls_list`
--

LOCK TABLES `api_calls_list` WRITE;
/*!40000 ALTER TABLE `api_calls_list` DISABLE KEYS */;
/*!40000 ALTER TABLE `api_calls_list` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `doodle`
--

DROP TABLE IF EXISTS `doodle`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `doodle` (
  `doodle_id` int(11) NOT NULL AUTO_INCREMENT,
  `doodle_title` varchar(64) COLLATE utf8_unicode_ci NOT NULL,
  `doodle_url` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `doodle_image` mediumtext COLLATE utf8_unicode_ci NOT NULL,
  `city` int(10) NOT NULL,
  `doodle_group` varchar(10) COLLATE utf8_unicode_ci NOT NULL,
  `doodle_html_text` mediumtext COLLATE utf8_unicode_ci,
  `expires_impressions` int(7) DEFAULT '0',
  `expires_date` datetime DEFAULT NULL,
  `date_scheduled` datetime DEFAULT NULL,
  `date_added` datetime NOT NULL,
  `date_status_change` datetime DEFAULT NULL,
  `status` int(1) NOT NULL DEFAULT '1',
  `type` varchar(250) COLLATE utf8_unicode_ci NOT NULL,
  `doodle_slug` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT  CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL ,
  PRIMARY KEY (`doodle_id`),
  KEY `idx_doodle_group` (`doodle_group`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `doodle`
--

LOCK TABLES `doodle` WRITE;
/*!40000 ALTER TABLE `doodle` DISABLE KEYS */;
/*!40000 ALTER TABLE `doodle` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `doodle_history`
--

DROP TABLE IF EXISTS `doodle_history`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `doodle_history` (
  `doodle_history_id` int(11) NOT NULL AUTO_INCREMENT,
  `doodle_id` int(11) NOT NULL,
  `doodle_shown` int(5) NOT NULL DEFAULT '0',
  `doodle_clicked` int(5) NOT NULL DEFAULT '0',
  `doodle_history_date` datetime NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT  CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL ,
  PRIMARY KEY (`doodle_history_id`),
  KEY `idx_doodle_history_doodle_id` (`doodle_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `doodle_history`
--

LOCK TABLES `doodle_history` WRITE;
/*!40000 ALTER TABLE `doodle_history` DISABLE KEYS */;
/*!40000 ALTER TABLE `doodle_history` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `categories`
--

DROP TABLE IF EXISTS `categories`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `categories` (
  `categories_id` int(11) NOT NULL AUTO_INCREMENT,
  `categories_name` varchar(32) COLLATE utf8_unicode_ci NOT NULL,
  `categories_image` mediumtext COLLATE utf8_unicode_ci,
  `categories_icon` mediumtext COLLATE utf8_unicode_ci NOT NULL,
  `parent_categories_id` int(11) NOT NULL DEFAULT '0',
  `sort_order` int(3) DEFAULT NULL,
  `categories_slug` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT  CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL ,
  PRIMARY KEY (`categories_id`),
  KEY `idx_categories_parent_id` (`parent_categories_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `categories`
--

LOCK TABLES `categories` WRITE;
/*!40000 ALTER TABLE `categories` DISABLE KEYS */;
/*!40000 ALTER TABLE `categories` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `categories_description`
--

DROP TABLE IF EXISTS `categories_description`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `categories_description` (
  `categories_description_id` int(11) NOT NULL AUTO_INCREMENT,
  `categories_id` int(11) NOT NULL,
  `categories_description` text NOT NULL,
  `language_id` int(11) NOT NULL DEFAULT '1',
  `created_at` timestamp NOT NULL DEFAULT  CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL ,
  PRIMARY KEY (`categories_description_id`),
  FOREIGN KEY (`categories_id`)REFERENCES categories(`categories_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `categories_description`
--

LOCK TABLES `categories_description` WRITE;
/*!40000 ALTER TABLE `categories_description` DISABLE KEYS */;
/*!40000 ALTER TABLE `categories_description` ENABLE KEYS */;
UNLOCK TABLES;



--
-- Table structure for table `cities`
--

DROP TABLE IF EXISTS `cities`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `cities` (
  `cities_id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL,
  `state_id` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT  CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL ,
  PRIMARY KEY (`cities_id`),
  UNIQUE KEY (`name`)

) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `cities`
--

LOCK TABLES `cities` WRITE;
/*!40000 ALTER TABLE `cities` DISABLE KEYS */;
/*!40000 ALTER TABLE `cities` ENABLE KEYS */;
UNLOCK TABLES;


--
-- Table structure for table `customers`
--

DROP TABLE IF EXISTS `customers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `customers` (
  `customers_id` int(11) NOT NULL AUTO_INCREMENT,
  `customers_gender` char(1) COLLATE utf8_unicode_ci DEFAULT '0',
  `customers_firstname` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `customers_lastname` varchar(255) COLLATE utf8_unicode_ci  NOT NULL,
  `customers_dob` date COLLATE utf8_unicode_ci NULL,
  `email` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `user_name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `customers_default_address_id` int(11) DEFAULT NULL,
  `customers_telephone` varchar(255) COLLATE utf8_unicode_ci NULL,
  `customers_phone` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `customers_fax` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `password` varchar(60) COLLATE utf8_unicode_ci NOT NULL,
  `customers_newsletter` char(1) COLLATE utf8_unicode_ci DEFAULT NULL,
  `is_active` tinyint(1) NOT NULL DEFAULT '1',
  `fb_id` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `google_id` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `customers_picture` mediumtext COLLATE utf8_unicode_ci  NULL,
  `is_seen` tinyint(1) NOT NULL DEFAULT '0',
  `remember_token` varchar(255) COLLATE utf8_unicode_ci  NULL,
  `role` varchar(100) COLLATE utf8_unicode_ci  NULL,
  `otp` varchar(100) COLLATE utf8_unicode_ci  NULL,
  `admin_approval` varchar(100) COLLATE utf8_unicode_ci NOT NULL DEFAULT '0',
  `reject_reason` text COLLATE utf8_unicode_ci  NULL,
  `created_at` timestamp NOT NULL DEFAULT  CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL ,
  PRIMARY KEY (`customers_id`),
  KEY `idx_customers_email_address` (`email`),
  UNIQUE KEY `customers_user_name_unique` (`user_name`),
  UNIQUE KEY `customers_email_unique` (`email`),
  UNIQUE KEY `customers_phone_unique` (`customers_phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `customers`
--

LOCK TABLES `customers` WRITE;
/*!40000 ALTER TABLE `customers` DISABLE KEYS */;
/*!40000 ALTER TABLE `customers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `customers_basket`
--

DROP TABLE IF EXISTS `customers_basket`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `customers_basket` (
  `customers_basket_id` int(11) NOT NULL AUTO_INCREMENT,
  `customers_id` int(11) NOT NULL,
  `services_id` int(11) NOT NULL,
  `customers_basket_quantity` int(2) NOT NULL,
  `final_price` decimal(15,2) DEFAULT NULL,
  `customers_basket_date_added` char(10) COLLATE utf8_unicode_ci DEFAULT NULL,
  `is_order` tinyint(1) NOT NULL DEFAULT '0',
  `session_id` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY (`customers_basket_id`),
  KEY `idx_customers_basket_customers_id` (`customers_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `customers_basket`
--

LOCK TABLES `customers_basket` WRITE;
/*!40000 ALTER TABLE `customers_basket` DISABLE KEYS */;
/*!40000 ALTER TABLE `customers_basket` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `customers_basket_attributes`
--

DROP TABLE IF EXISTS `customers_basket_attributes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `customers_basket_attributes` (
  `customers_basket_attributes_id` int(11) NOT NULL AUTO_INCREMENT,
  `customers_basket_id` int(100) NOT NULL,
  `customers_id` int(11) NOT NULL,
  `services_id` text COLLATE utf8_unicode_ci NOT NULL,
  `services_options_id` int(11) NOT NULL,
  `services_options_values_id` int(11) NOT NULL,
  `session_id` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY (`customers_basket_attributes_id`),
  KEY `idx_customers_basket_att_customers_id` (`customers_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `customers_basket_attributes`
--

LOCK TABLES `customers_basket_attributes` WRITE;
/*!40000 ALTER TABLE `customers_basket_attributes` DISABLE KEYS */;
/*!40000 ALTER TABLE `customers_basket_attributes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `customers_info`
--

DROP TABLE IF EXISTS `customers_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `customers_info` (
  `customers_info_id` int(11) NOT NULL,
  `customers_info_date_of_last_logon` datetime DEFAULT NULL,
  `customers_info_number_of_logons` int(5) DEFAULT NULL,
  `customers_info_date_account_created` datetime DEFAULT NULL,
  `customers_info_date_account_last_modified` datetime DEFAULT NULL,
  `global_service_notifications` int(1) DEFAULT '0',
  PRIMARY KEY (`customers_info_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `customers_info`
--

LOCK TABLES `customers_info` WRITE;
/*!40000 ALTER TABLE `customers_info` DISABLE KEYS */;
/*!40000 ALTER TABLE `customers_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `devices`
--

DROP TABLE IF EXISTS `devices`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `devices` (
  `id` int(100) NOT NULL AUTO_INCREMENT,
  `device_id` text COLLATE utf8_unicode_ci NOT NULL,
  `customers_id` int(100) NOT NULL DEFAULT '0',
  `device_type` text COLLATE utf8_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT  CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL ,
  `status` tinyint(1) NOT NULL DEFAULT '0',
  `isDesktop` tinyint(1) NOT NULL DEFAULT '0',
  `onesignal` tinyint(1) NOT NULL DEFAULT '0',
  `isEnableMobile` tinyint(1) NOT NULL DEFAULT '1',
  `isEnableDesktop` tinyint(1) NOT NULL DEFAULT '1',
  `ram` varchar(250) COLLATE utf8_unicode_ci DEFAULT NULL,
  `processor` varchar(250) COLLATE utf8_unicode_ci DEFAULT NULL,
  `device_os` varchar(250) COLLATE utf8_unicode_ci DEFAULT NULL,
  `location` varchar(250) COLLATE utf8_unicode_ci DEFAULT NULL,
  `device_model` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `manufacturer` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `is_notify` tinyint(1) NOT NULL DEFAULT '1',
  `fcm` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `devices`
--

LOCK TABLES `devices` WRITE;
/*!40000 ALTER TABLE `devices` DISABLE KEYS */;
/*!40000 ALTER TABLE `devices` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `employee_details`
--

DROP TABLE IF EXISTS `employee_details`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `employee_details` (
  `empid` int(10) NOT NULL AUTO_INCREMENT,
  `profilepic` varchar(250) COLLATE utf8mb4_unicode_ci NOT NULL,
  `empname` varchar(200) COLLATE utf8mb4_unicode_ci NOT NULL,
  `mobileno` bigint(10) NOT NULL,
  `email` varchar(200) COLLATE utf8mb4_unicode_ci NOT NULL,
  `bike_status` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `licenseno` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `address` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(200) COLLATE utf8mb4_unicode_ci NOT NULL,
  `vendorid` int(10) NOT NULL,
  `approval` varchar(200) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'Pending',
  `fcmtoken` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `rejectreason` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `empstatus` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT  CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL ,
  PRIMARY KEY (`empid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `employee_details`
--

LOCK TABLES `employee_details` WRITE;
/*!40000 ALTER TABLE `employee_details` DISABLE KEYS */;
/*!40000 ALTER TABLE `employee_details` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `languages`
--

DROP TABLE IF EXISTS `languages`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `languages` (
  `languages_id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(32) COLLATE utf8_unicode_ci NOT NULL,
  `code` char(2) COLLATE utf8_unicode_ci NOT NULL,
  `image` mediumtext COLLATE utf8_unicode_ci,
  `directory` varchar(32) COLLATE utf8_unicode_ci DEFAULT NULL,
  `sort_order` int(3) DEFAULT NULL,
  `direction` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `is_default` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`languages_id`),
  KEY `IDX_LANGUAGES_NAME` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `languages`
--

LOCK TABLES `languages` WRITE;
/*!40000 ALTER TABLE `languages` DISABLE KEYS */;
/*!40000 ALTER TABLE `languages` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `mobile_otp_details`
--

DROP TABLE IF EXISTS `mobile_otp_details`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `mobile_otp_details` (
  `varification_id` int(11) NOT NULL AUTO_INCREMENT,
  `mobile_number` bigint(12) NOT NULL,
  `otp` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL,
  `user_type` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `user_id` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT  CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL ,
  PRIMARY KEY (`varification_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `mobile_otp_details`
--

LOCK TABLES `mobile_otp_details` WRITE;
/*!40000 ALTER TABLE `mobile_otp_details` DISABLE KEYS */;
/*!40000 ALTER TABLE `mobile_otp_details` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `notification_details`
--

DROP TABLE IF EXISTS `notification_details`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `notification_details` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `userid` int(10) NOT NULL,
  `moduleid` int(10) NOT NULL,
  `module` varchar(100) NOT NULL,
  `username` varchar(100) NOT NULL,
  `title` varchar(250) NOT NULL,
  `message` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT  CURRENT_TIMESTAMP,
  `IsSeen` tinyint(4) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `notification_details`
--

LOCK TABLES `notification_details` WRITE;
/*!40000 ALTER TABLE `notification_details` DISABLE KEYS */;
/*!40000 ALTER TABLE `notification_details` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `booking`
--

DROP TABLE IF EXISTS `booking`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `booking` (
  `booking_id` int(11) NOT NULL AUTO_INCREMENT,
  `total_tax` decimal(7,2) NOT NULL,
  `customers_id` int(11) NOT NULL,
  `customers_name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `customers_company` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `customers_street_address` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `customers_suburb` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `customers_city` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `customers_postcode` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `customers_state` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `customers_country` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `customers_telephone` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `email` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `customers_address_format_id` int(5) DEFAULT NULL,
  `delivery_name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `delivery_company` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `delivery_street_address` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `delivery_suburb` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `delivery_city` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `delivery_postcode` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `delivery_state` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `delivery_country` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `delivery_address_format_id` int(5) DEFAULT NULL,
  `delivery_locality` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `billing_name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `billing_company` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `billing_street_address` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `billing_suburb` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `billing_city` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `billing_postcode` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `billing_state` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `billing_country` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `billing_address_format_id` int(5) NOT NULL,
  `billing_locality` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `payment_method` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `cc_type` varchar(20) COLLATE utf8_unicode_ci DEFAULT NULL,
  `cc_owner` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `cc_number` varchar(32) COLLATE utf8_unicode_ci DEFAULT NULL,
  `cc_expires` varchar(4) COLLATE utf8_unicode_ci DEFAULT NULL,
  `last_modified` datetime DEFAULT NULL,
  `date_purchased` datetime DEFAULT NULL,
  `booking_date_finished` datetime DEFAULT NULL,
  `currency` char(3) COLLATE utf8_unicode_ci DEFAULT NULL,
  `currency_value` decimal(14,6) DEFAULT NULL,
  `booking_price` decimal(10,2) NOT NULL,
  `additinal_cost` decimal(10,2) NOT NULL,
  `booking_information` mediumtext COLLATE utf8_unicode_ci NOT NULL,
  `is_seen` tinyint(1) NOT NULL DEFAULT '0',
  `coupon_code` text COLLATE utf8_unicode_ci NOT NULL,
  `coupon_amount` int(100) NOT NULL,
  `exclude_service_ids` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `service_categories` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `excluded_service_categories` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `booking_source` tinyint(1) NOT NULL DEFAULT '1' COMMENT '1: Website, 2: App',
  `booking_phone` varchar(30) COLLATE utf8_unicode_ci NOT NULL,
  `billing_phone` varchar(30) COLLATE utf8_unicode_ci NOT NULL,
  `transaction_id` text COLLATE utf8_unicode_ci,
  `payment_status` varchar(100) COLLATE utf8_unicode_ci NOT NULL DEFAULT 'Pending',
  `vendortype` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `vendorid` int(10) NOT NULL,
  `commission_perc` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `empstatus` varchar(100) COLLATE utf8_unicode_ci NOT NULL DEFAULT 'Pending',
  `rejectreason` text COLLATE utf8_unicode_ci NOT NULL,
  `acceptedempid` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `acceptedtime` varchar(200) COLLATE utf8_unicode_ci NOT NULL,
  `completedtime` varchar(200) COLLATE utf8_unicode_ci NOT NULL,
  `rejectedempid` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `image` mediumtext COLLATE utf8_unicode_ci,
  `isPaidByVendor` varchar(10) COLLATE utf8_unicode_ci NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT  CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL ,
  PRIMARY KEY (`booking_id`),
  KEY `idx_booking_customers_id` (`customers_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `booking`
--

LOCK TABLES `booking` WRITE;
/*!40000 ALTER TABLE `booking` DISABLE KEYS */;
/*!40000 ALTER TABLE `booking` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `booking_services`
--

DROP TABLE IF EXISTS `booking_services`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `booking_services` (
  `booking_services_id` int(11) NOT NULL AUTO_INCREMENT,
  `booking_id` int(11) NOT NULL,
  `services_id` int(11) COLLATE utf8_unicode_ci NOT NULL,
  `services_description` varchar(12) COLLATE utf8_unicode_ci DEFAULT NULL,
  `services_name` varchar(64) COLLATE utf8_unicode_ci NOT NULL,
  `services_price` decimal(15,2) NOT NULL,
  `final_price` decimal(15,2) NOT NULL,
  `services_tax` decimal(7,0) NOT NULL,
  `services_quantity` int(2) NOT NULL,
  `admin_comission_perc` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT  CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL ,
  PRIMARY KEY (`booking_services_id`),
  KEY `idx_booking_services_booking_id` (`booking_id`),
  KEY `idx_booking_services_services_id` (`services_id`),
  FOREIGN KEY(`booking_id`) REFERENCES booking(`booking_id`),
  FOREIGN KEY(`services_id`)REFERENCES services(`services_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `booking_services`
--

LOCK TABLES `booking_services` WRITE;
/*!40000 ALTER TABLE `booking_services` DISABLE KEYS */;
/*!40000 ALTER TABLE `booking_services` ENABLE KEYS */;
UNLOCK TABLES;



DROP TABLE IF EXISTS `booking_packages`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `booking_packages` (
  `booking_packages_id` int(11) NOT NULL AUTO_INCREMENT,
  `booking_id` int(11) NOT NULL,
  `packages_id` int(11) NOT NULL,
  `packages_description` varchar(12) COLLATE utf8_unicode_ci DEFAULT NULL,
  `packages_name` varchar(64) COLLATE utf8_unicode_ci NOT NULL,
  `packages_price` decimal(15,2) NOT NULL,
  `final_price` decimal(15,2) NOT NULL,
  `packages_tax` decimal(7,0) NOT NULL,
  `packages_quantity` int(2) NOT NULL,
  `admin_comission_perc` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT  CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL ,
  PRIMARY KEY (`booking_packages_id`),
  KEY `idx_booking_services_booking_id` (`booking_id`),
  KEY `idx_booking_packages_packages_id` (`packages_id`),
  FOREIGN KEY(`booking_id`) REFERENCES booking(`booking_id`),
  FOREIGN KEY(`packages_id`)REFERENCES packages(`packages_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `booking_packages`
--

LOCK TABLES `booking_packages` WRITE;
/*!40000 ALTER TABLE `booking_packages` DISABLE KEYS */;
/*!40000 ALTER TABLE `booking_packages` ENABLE KEYS */;
UNLOCK TABLES;


--
-- Table structure for table `booking_status`
--

DROP TABLE IF EXISTS `booking_status`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `booking_status` (
  `booking_status_id` int(11) NOT NULL DEFAULT '0',
  `language_id` int(11) NOT NULL DEFAULT '1',
  `booking_status_name` varchar(32) COLLATE utf8_unicode_ci NOT NULL,
  `public_flag` int(11) DEFAULT '0',
  `downloads_flag` int(11) DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT  CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL ,
  PRIMARY KEY (`booking_status_id`,`language_id`),
  KEY `idx_booking_status_name` (`booking_status_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `booking_status`
--

LOCK TABLES `booking_status` WRITE;
/*!40000 ALTER TABLE `booking_status` DISABLE KEYS */;
/*!40000 ALTER TABLE `booking_status` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `booking_status_description`
--

DROP TABLE IF EXISTS `booking_status_description`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `booking_status_description` (
  `booking_status_description_id` int(100) NOT NULL AUTO_INCREMENT,
  `booking_status_id` int(11) NOT NULL,
  `booking_status_name` varchar(255) NOT NULL,
  `language_id` int(100) NOT NULL,
  PRIMARY KEY (`booking_status_description_id`),
  FOREIGN KEY(`booking_status_id`)REFERENCES booking_status(`booking_status_id`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `booking_status_description`
--

LOCK TABLES `booking_status_description` WRITE;
/*!40000 ALTER TABLE `booking_status_description` DISABLE KEYS */;
/*!40000 ALTER TABLE `booking_status_description` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `booking_status_history`
--

DROP TABLE IF EXISTS `booking_status_history`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `booking_status_history` (
  `booking_status_history_id` int(11) NOT NULL AUTO_INCREMENT,
  `booking_id` int(11) NOT NULL,
  `booking_status_id` int(5) NOT NULL,
  `date_added` datetime NOT NULL,
  `customer_notified` int(1) DEFAULT '0',
  `comments` mediumtext COLLATE utf8_unicode_ci,
  `created_at` timestamp NOT NULL DEFAULT  CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL ,
  PRIMARY KEY (`booking_status_history_id`),
  KEY `idx_booking_status_history_booking_id` (`booking_id`),
  FOREIGN KEY(`booking_id`)REFERENCES booking(`booking_id`),
  FOREIGN KEY(`booking_status_id`)REFERENCES booking_status(`booking_status_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `booking_status_history`
--

LOCK TABLES `booking_status_history` WRITE;
/*!40000 ALTER TABLE `booking_status_history` DISABLE KEYS */;
/*!40000 ALTER TABLE `booking_status_history` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pages`
--


--
-- Table structure for table `paid_money_by_vendor`
--

DROP TABLE IF EXISTS `paid_money_by_vendor`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `paid_money_by_vendor` (
  `p_id` int(11) NOT NULL AUTO_INCREMENT,
  `vendor_id` int(11) NOT NULL,
  `totalAmount` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `paidAmount` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `PendingAmount` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT  CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL ,
  PRIMARY KEY (`p_id`),
  FOREIGN KEY(`vendor_id`)REFERENCES vendor(`vendor_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `paid_money_by_vendor`
--

LOCK TABLES `paid_money_by_vendor` WRITE;
/*!40000 ALTER TABLE `paid_money_by_vendor` DISABLE KEYS */;
/*!40000 ALTER TABLE `paid_money_by_vendor` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `payment_description`
--

DROP TABLE IF EXISTS `payment_description`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `payment_description` (
  `p_id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `language_id` int(11) NOT NULL,
  `payment_name` varchar(32) NOT NULL,
  `sub_name_1` varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `sub_name_2` varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY (`p_id`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `payment_description`
--

LOCK TABLES `payment_description` WRITE;
/*!40000 ALTER TABLE `payment_description` DISABLE KEYS */;
/*!40000 ALTER TABLE `payment_description` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `payments_setting`
--

DROP TABLE IF EXISTS `payments_setting`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `payments_setting` (
  `payments_id` int(100) NOT NULL AUTO_INCREMENT,
  `braintree_enviroment` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `braintree_name` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `braintree_merchant_id` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `braintree_public_key` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `braintree_private_key` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `brantree_active` tinyint(1) NOT NULL DEFAULT '0',
  `stripe_enviroment` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `stripe_name` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `secret_key` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `publishable_key` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `stripe_active` tinyint(1) NOT NULL DEFAULT '0',
  `cash_on_delivery` tinyint(1) NOT NULL DEFAULT '0',
  `cod_name` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `paypal_name` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `paypal_id` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `paypal_status` tinyint(1) NOT NULL DEFAULT '0',
  `paypal_enviroment` tinyint(1) DEFAULT '0',
  `payment_currency` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `instamojo_enviroment` tinyint(1) NOT NULL,
  `instamojo_name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `instamojo_api_key` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `instamojo_auth_token` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `instamojo_client_id` text COLLATE utf8_unicode_ci NOT NULL,
  `instamojo_client_secret` text COLLATE utf8_unicode_ci NOT NULL,
  `instamojo_active` tinyint(1) NOT NULL DEFAULT '0',
  `cybersource_name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `cybersource_status` tinyint(1) NOT NULL DEFAULT '0',
  `cybersource_enviroment` tinyint(1) NOT NULL DEFAULT '0',
  `hyperpay_enviroment` tinyint(1) NOT NULL DEFAULT '0',
  `hyperpay_name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `hyperpay_active` tinyint(1) NOT NULL DEFAULT '0',
  `hyperpay_userid` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `hyperpay_password` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `hyperpay_entityid` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY (`payments_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `payments_setting`
--

LOCK TABLES `payments_setting` WRITE;
/*!40000 ALTER TABLE `payments_setting` DISABLE KEYS */;
/*!40000 ALTER TABLE `payments_setting` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pincode_area`
--

DROP TABLE IF EXISTS `pincode_area`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `pincode_area` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `pincode` int(6) NOT NULL,
  `cities_id` int(10) NOT NULL,
  `locality` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `delivery` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT  CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL ,
  PRIMARY KEY (`id`),
  FOREIGN KEY(`cities_id`) REFERENCES cities(`cities_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pincode_area`
--

LOCK TABLES `pincode_area` WRITE;
/*!40000 ALTER TABLE `pincode_area` DISABLE KEYS */;
/*!40000 ALTER TABLE `pincode_area` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pincode_master`
--

DROP TABLE IF EXISTS `pincode_master`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `pincode_master` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `first_name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `last_name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `pincode` int(10) NOT NULL,
  `address` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `locality` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `city` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `mobileno` bigint(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT  CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL ,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pincode_master`
--

LOCK TABLES `pincode_master` WRITE;
/*!40000 ALTER TABLE `pincode_master` DISABLE KEYS */;
/*!40000 ALTER TABLE `pincode_master` ENABLE KEYS */;
UNLOCK TABLES;



DROP TABLE IF EXISTS `packages`;

CREATE TABLE `packages` (
  `packages_id` int(11) NOT NULL AUTO_INCREMENT,
  `packages_slug` varchar(255) NOT NULL,
  `packages_name` varchar(64)  NOT NULL,
  `packages_description` text,
  `packages_url` varchar(255) DEFAULT NULL,
  `packages_icon_url` varchar(255) NOT NULL,
  `packages_image_url`  varchar(255) NOT NULL ,
  `packages_price` decimal(15,2) NOT NULL,
  `packages_status` tinyint(1) NOT NULL default '1',
  `packages_tax_class_id` int(11) NOT NULL default '0' ,
  `packages_liked` int(100) NOT NULL DEFAULT '0',
  `is_special`  int(1) NOT NULL DEFAULT '0',
  `packages_viewed` int(5) DEFAULT '0',
  `packages_banner_url` varchar(255) NOT NULL ,
  `added_by` int(11) NULL ,
  `created_at` timestamp NOT NULL DEFAULT  CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL ,
  PRIMARY KEY (`packages_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- Dumping data for table `packages`
--

LOCK TABLES `packages` WRITE;
/*!40000 ALTER TABLE `packages` DISABLE KEYS */;
/*!40000 ALTER TABLE `packages` ENABLE KEYS */;
UNLOCK TABLES;





--
-- Dumping data for table `packages_to_services`
--

DROP TABLE IF EXISTS `packages_to_services`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `packages_to_services` (
  `services_id` int(11) NOT NULL,
  `packages_id` int(11) NOT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL ,
  FOREIGN KEY(`services_id`) REFERENCES services(`services_id`),
  FOREIGN KEY(`packages_id`) REFERENCES packages(`packages_id`),
  CONSTRAINT packages_service_id PRIMARY KEY (packages_id, services_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `packages_to_services`
--

LOCK TABLES `packages_to_services` WRITE;
/*!40000 ALTER TABLE `packages_to_services` DISABLE KEYS */;
/*!40000 ALTER TABLE `packages_to_services` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `services`
--

DROP TABLE IF EXISTS `services`;

CREATE TABLE `services` (
  `services_id` int(11) NOT NULL AUTO_INCREMENT,
  `services_image_url`  varchar(255) NOT NULL ,
  `services_icon_url` varchar(255) NOT NULL,
  `services_price` decimal(15,2) NOT NULL,
  `services_status` tinyint(1) NOT NULL default '1',
  `services_tax_class_id` int(11) NOT NULL default '0' ,
  `services_liked` int(100) NOT NULL DEFAULT '0',
  `is_special`  int(1) NOT NULL DEFAULT '0',
  `services_slug` varchar(255) NOT NULL,
  `added_by` int(11) NULL ,
  `created_at` timestamp NOT NULL DEFAULT  CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL ,
  PRIMARY KEY (`services_id`)
  ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- Dumping data for table `services`
--

LOCK TABLES `services` WRITE;
/*!40000 ALTER TABLE `services` DISABLE KEYS */;
/*!40000 ALTER TABLE `services` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `services_attributes`
--

DROP TABLE IF EXISTS `services_description`;

CREATE TABLE `services_description` (
  `services_id` int(11) NOT NULL AUTO_INCREMENT,
  `services_name` varchar(64)  NOT NULL,
  `services_description` text,
  `services_url` varchar(255) DEFAULT NULL,
  `services_viewed` int(5) DEFAULT '0',
  `services_banner_url` varchar(255) NOT NULL ,
   `deleted_at` timestamp NULL DEFAULT NULL ,
  FOREIGN KEY (`services_id`) REFERENCES services(`services_id`),
  KEY `services_name` (`services_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- Dumping data for table `services_description`
--

LOCK TABLES `services_description` WRITE;
UNLOCK TABLES;

--
-- Table structure for table `services_images`
--

DROP TABLE IF EXISTS `services_images`;

CREATE TABLE `services_images` (
  `image_id` int(11) NOT NULL AUTO_INCREMENT,
  `services_id` int(11) NOT NULL,
  `image_url` varchar(255) COLLATE utf8_unicode_ci,
  `html_content` varchar(1000) COLLATE utf8_unicode_ci,
  `sort_order` int(11) NOT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL ,
  PRIMARY KEY (`image_id`),
  KEY `services_images_id` (`services_id`),
  FOREIGN KEY(`services_id`)REFERENCES services(`services_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- Dumping data for table `services_images`
--

LOCK TABLES `services_images` WRITE;

UNLOCK TABLES;

--
-- Dumping data for table `services_to_categories`
--

DROP TABLE IF EXISTS `services_to_categories`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `services_to_categories` (
  `services_id` int(11) NOT NULL,
  `categories_id` int(11) NOT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL ,
  FOREIGN KEY(`services_id`) REFERENCES services(`services_id`),
  FOREIGN KEY(`categories_id`) REFERENCES categories(`categories_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `services_to_categories`
--

LOCK TABLES `services_to_categories` WRITE;
/*!40000 ALTER TABLE `services_to_categories` DISABLE KEYS */;
/*!40000 ALTER TABLE `services_to_categories` ENABLE KEYS */;
UNLOCK TABLES;


--
-- Table structure for table `reviews`
--

DROP TABLE IF EXISTS `reviews`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `reviews` (
  `reviews_id` int(11) NOT NULL AUTO_INCREMENT,
  `services_id` int(11) NOT NULL,
  `packages_id` int(11) NOT NULL,
  `customers_id` int(11) DEFAULT NULL,
  `vendor_id` int(11) DEFAULT NULL,
  `customers_name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `reviews_rating` int(1) DEFAULT NULL,
  `reviews_status` tinyint(1) NOT NULL DEFAULT '0',
  `reviews_read` int(5) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT  CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL ,
  PRIMARY KEY (`reviews_id`),
  KEY `idx_reviews_services_id` (`services_id`),
  KEY `idx_reviews_customers_id` (`customers_id`),
  FOREIGN KEY(`services_id`)REFERENCES services(`services_id`),
  FOREIGN KEY(`packages_id`)REFERENCES packages(`packages_id`),
  FOREIGN KEY(`vendor_id`)REFERENCES vendor(`vendor_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `reviews`
--

LOCK TABLES `reviews` WRITE;
/*!40000 ALTER TABLE `reviews` DISABLE KEYS */;
/*!40000 ALTER TABLE `reviews` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `reviews_description`
--

DROP TABLE IF EXISTS `reviews_description`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `reviews_description` (
  `reviews_id` int(11) NOT NULL,
  `reviews_text` mediumtext COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY (`reviews_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `reviews_description`
--

LOCK TABLES `reviews_description` WRITE;
/*!40000 ALTER TABLE `reviews_description` DISABLE KEYS */;
/*!40000 ALTER TABLE `reviews_description` ENABLE KEYS */;
UNLOCK TABLES;


--
-- Table structure for table `vendor_bankdetails`
--

DROP TABLE IF EXISTS `vendor_bank_details`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `vendor_bank_details` (
  `vendor_bank_id` int(10) NOT NULL AUTO_INCREMENT,
  `bank_name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `ifsc_code` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `account_number` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `customer_name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT  CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL ,
  `vendor_id` int(10) NOT NULL,
  PRIMARY KEY (`vendor_bank_id`),
  FOREIGN KEY(`vendor_id`) REFERENCES vendor(`vendor_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `vendor_bank_details`
--

LOCK TABLES `vendor_bank_details` WRITE;
/*!40000 ALTER TABLE `vendor_bank_details` DISABLE KEYS */;
/*!40000 ALTER TABLE `vendor_bank_details` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `vendor_collection`
--

DROP TABLE IF EXISTS `vendor_collection`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `vendor_collection` (
  `vendor_collection_id` int(11) NOT NULL AUTO_INCREMENT,
  `vendor_id` int(11)  NOT NULL,
  `total_collection` float(10,2) COLLATE utf8mb4_unicode_ci NOT NULL,
  `total_paid` float(10,2) COLLATE utf8mb4_unicode_ci NOT NULL,
  `total_unpaid` float(10,2) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT  CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`vendor_collection_id`),
  FOREIGN KEY(`vendor_id`) REFERENCES vendor(`vendor_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `vendor_collection`
--

LOCK TABLES `vendor_collection` WRITE;
/*!40000 ALTER TABLE `vendor_collection` DISABLE KEYS */;
/*!40000 ALTER TABLE `vendor_collection` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `vendor_commission`
--

DROP TABLE IF EXISTS `vendor_commission`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `vendor_commission` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `vendor_id` int(10) NOT NULL,
  `category_id` int(10) NOT NULL,
  `commission_perc` float(10,2) NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`vendor_id`) REFERENCES vendor(`vendor_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `vendor_commission`
--

LOCK TABLES `vendor_commission` WRITE;
/*!40000 ALTER TABLE `vendor_commission` DISABLE KEYS */;
/*!40000 ALTER TABLE `vendor_commission` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `vendor`
--

DROP TABLE IF EXISTS `vendors`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `vendors` (
  `vendors_id` int(10) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `business_name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `business_address` varchar(200) COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `city` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `pan_card_id` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `gst_no` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
      `mobile_no` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `vendors_picture` varchar(200) COLLATE utf8mb4_unicode_ci NOT NULL,
  `status` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `rating` int(10) NOT NULL,
  `otp` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `activation_status` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'Pending',
  `activation_date` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `deactivation_reason` varchar(200) COLLATE utf8mb4_unicode_ci NOT NULL,
  `fcm_token` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT  CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL ,
  PRIMARY KEY (`vendors_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `vendors`
--

LOCK TABLES `vendors` WRITE;
/*!40000 ALTER TABLE `vendors` DISABLE KEYS */;
/*!40000 ALTER TABLE `vendors` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `whos_online`
--

DROP TABLE IF EXISTS `whos_online`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `whos_online` (
  `customer_id` int(11) NOT NULL DEFAULT '0',
  `full_name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `session_id` varchar(128) COLLATE utf8_unicode_ci NOT NULL,
  `ip_address` varchar(15) COLLATE utf8_unicode_ci NOT NULL,
  `time_entry` varchar(14) COLLATE utf8_unicode_ci NOT NULL,
  `time_last_click` varchar(14) COLLATE utf8_unicode_ci NOT NULL,
  `last_page_url` mediumtext COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY (`customer_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `whos_online`
--

LOCK TABLES `whos_online` WRITE;
/*!40000 ALTER TABLE `whos_online` DISABLE KEYS */;
/*!40000 ALTER TABLE `whos_online` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-03-08 18:01:00
