-- MySQL dump 10.13  Distrib 5.7.37, for Win64 (x86_64)
--
-- Host: localhost    Database: oneclick
-- ------------------------------------------------------
-- Server version	5.7.24

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
-- Table structure for table `admin_users`
--

DROP TABLE IF EXISTS `admin_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `admin_users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `username` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `password` varchar(80) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `realname` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `avatar_url` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `access` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `admin_users_username_IDX` (`username`),
  KEY `admin_users_deleted_at_IDX` (`deleted_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='管理员';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `admin_users`
--

LOCK TABLES `admin_users` WRITE;
/*!40000 ALTER TABLE `admin_users` DISABLE KEYS */;
/*!40000 ALTER TABLE `admin_users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `article_categories`
--

DROP TABLE IF EXISTS `article_categories`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `article_categories` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '文章分类id',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `parent_id` int(10) unsigned DEFAULT NULL COMMENT '父级ID',
  `title` varchar(128) DEFAULT NULL COMMENT '文章分类标题',
  `image_url` varchar(64) DEFAULT NULL COMMENT '文章分类图片',
  `keywords` varchar(64) DEFAULT NULL COMMENT 'SEO关键词',
  `description` varchar(256) DEFAULT NULL COMMENT 'SEO描述',
  `sort` tinyint(3) unsigned DEFAULT '255' COMMENT '排序',
  PRIMARY KEY (`id`),
  KEY `article_categories_sort_IDX` (`sort`) USING BTREE,
  KEY `article_categories_deleted_at_IDX` (`deleted_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章分类';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `article_categories`
--

LOCK TABLES `article_categories` WRITE;
/*!40000 ALTER TABLE `article_categories` DISABLE KEYS */;
/*!40000 ALTER TABLE `article_categories` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `article_contents`
--

DROP TABLE IF EXISTS `article_contents`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `article_contents` (
  `article_id` int(10) unsigned NOT NULL COMMENT '文章id',
  `content` mediumtext COMMENT '文章内容',
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`article_id`),
  CONSTRAINT `article_contents_FK` FOREIGN KEY (`article_id`) REFERENCES `articles` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章内容';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `article_contents`
--

LOCK TABLES `article_contents` WRITE;
/*!40000 ALTER TABLE `article_contents` DISABLE KEYS */;
/*!40000 ALTER TABLE `article_contents` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `article_statistics`
--

DROP TABLE IF EXISTS `article_statistics`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `article_statistics` (
  `article_id` int(10) unsigned NOT NULL COMMENT '文章id',
  `updated_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `views` int(10) unsigned DEFAULT NULL,
  `favourites` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`article_id`),
  CONSTRAINT `article_statistics_FK` FOREIGN KEY (`article_id`) REFERENCES `articles` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章内容';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `article_statistics`
--

LOCK TABLES `article_statistics` WRITE;
/*!40000 ALTER TABLE `article_statistics` DISABLE KEYS */;
/*!40000 ALTER TABLE `article_statistics` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `article_tag_relates`
--

DROP TABLE IF EXISTS `article_tag_relates`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `article_tag_relates` (
  `article_id` int(10) unsigned NOT NULL,
  `tag_id` int(10) unsigned NOT NULL,
  PRIMARY KEY (`article_id`,`tag_id`),
  CONSTRAINT `article_tag_relates_FK` FOREIGN KEY (`article_id`) REFERENCES `articles` (`id`),
  CONSTRAINT `article_tag_relates_FK_1` FOREIGN KEY (`article_id`) REFERENCES `article_tags` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章标签关联';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `article_tag_relates`
--

LOCK TABLES `article_tag_relates` WRITE;
/*!40000 ALTER TABLE `article_tag_relates` DISABLE KEYS */;
/*!40000 ALTER TABLE `article_tag_relates` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `article_tags`
--

DROP TABLE IF EXISTS `article_tags`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `article_tags` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `name` varchar(32) DEFAULT NULL,
  `slug` varchar(32) DEFAULT NULL,
  `color` varchar(12) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `article_tags_slug_IDX` (`slug`) USING BTREE,
  KEY `article_tags_deleted_at_IDX` (`deleted_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章标签';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `article_tags`
--

LOCK TABLES `article_tags` WRITE;
/*!40000 ALTER TABLE `article_tags` DISABLE KEYS */;
/*!40000 ALTER TABLE `article_tags` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `articles`
--

DROP TABLE IF EXISTS `articles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `articles` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '文章管理ID',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '添加时间',
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `category_id` int(10) unsigned DEFAULT NULL COMMENT '分类id',
  `title` varchar(128) DEFAULT NULL COMMENT '文章标题',
  `image_url` varchar(64) DEFAULT NULL COMMENT '文章图片',
  `origin_url` varchar(64) DEFAULT NULL COMMENT '原文链接',
  `origin_author` varchar(32) DEFAULT NULL COMMENT '原文作者',
  `admin_id` int(10) unsigned DEFAULT NULL COMMENT '管理员id',
  `keywords` varchar(128) DEFAULT NULL COMMENT 'SEO关键词',
  `description` varchar(128) DEFAULT NULL COMMENT 'SEO描述',
  `sort` tinyint(3) unsigned DEFAULT '255' COMMENT '排序',
  PRIMARY KEY (`id`),
  KEY `articles_article_categories_FK` (`category_id`) USING BTREE,
  KEY `articles_title_IDX` (`title`) USING BTREE,
  KEY `articles_FK` (`admin_id`),
  KEY `articles_deleted_at_IDX` (`deleted_at`) USING BTREE,
  KEY `articles_sort_IDX` (`sort`) USING BTREE,
  CONSTRAINT `articles_FK` FOREIGN KEY (`admin_id`) REFERENCES `admin_users` (`id`) ON DELETE SET NULL ON UPDATE CASCADE,
  CONSTRAINT `articles_FK_1` FOREIGN KEY (`category_id`) REFERENCES `article_categories` (`id`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `articles`
--

LOCK TABLES `articles` WRITE;
/*!40000 ALTER TABLE `articles` DISABLE KEYS */;
/*!40000 ALTER TABLE `articles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `web_banners`
--

DROP TABLE IF EXISTS `web_banners`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `web_banners` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `image_url` varchar(128) DEFAULT NULL,
  `redirect_url` varchar(128) DEFAULT NULL,
  `banner_group` varchar(32) DEFAULT NULL,
  `sort` tinyint(3) unsigned DEFAULT '255',
  `remark` varchar(256) DEFAULT NULL,
  `title` varchar(100) DEFAULT NULL,
  `description` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `web_banners_banner_group_IDX` (`banner_group`) USING BTREE,
  KEY `web_banners_deleted_at_IDX` (`deleted_at`) USING BTREE,
  KEY `web_banners_sort_IDX` (`sort`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='轮播图';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `web_banners`
--

LOCK TABLES `web_banners` WRITE;
/*!40000 ALTER TABLE `web_banners` DISABLE KEYS */;
/*!40000 ALTER TABLE `web_banners` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `web_metas`
--

DROP TABLE IF EXISTS `web_metas`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `web_metas` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `meta_key` varchar(64) DEFAULT NULL,
  `meta_value` varchar(512) DEFAULT NULL,
  `meta_group` varchar(32) DEFAULT NULL,
  `meta_name` varchar(128) DEFAULT NULL,
  `remark` varchar(256) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `web_metas_meta_group_IDX` (`meta_group`) USING BTREE,
  KEY `web_metas_deleted_at_IDX` (`deleted_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='元数据';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `web_metas`
--

LOCK TABLES `web_metas` WRITE;
/*!40000 ALTER TABLE `web_metas` DISABLE KEYS */;
/*!40000 ALTER TABLE `web_metas` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'oneclick'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-08-19 21:36:26
