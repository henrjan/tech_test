SET time_zone = "+07:00";

CREATE DATABASE IF NOT EXISTS search_movie DEFAULT CHARACTER SET latin1 COLLATE latin1_swedish_ci;
USE search_movie;

CREATE TABLE IF NOT EXISTS `dth_access_log` (
  `id` varchar(36) NOT NULL,
  `url` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL,
  `method` varchar(10) COLLATE utf8_unicode_ci NOT NULL,
  `request_body` TEXT COLLATE utf8_unicode_ci NOT NULL,
  `response_body` TEXT COLLATE utf8_unicode_ci NOT NULL,
  `created_at` varchar(20) COLLATE utf8_unicode_ci NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
);
