-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server-Version:               11.3.2-MariaDB - mariadb.org binary distribution
-- Server-Betriebssystem:        Win64
-- HeidiSQL Version:             12.10.0.7000
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- Exportiere Datenbank-Struktur für rcjv_app
CREATE DATABASE IF NOT EXISTS `rcjv_app` /*!40100 DEFAULT CHARACTER SET utf16 COLLATE utf16_general_ci */;
USE `rcjv_app`;

-- Exportiere Struktur von Tabelle rcjv_app.institution
CREATE TABLE IF NOT EXISTS `institution` (
  `institution_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`institution_id`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=utf16 COLLATE=utf16_general_ci;

-- Exportiere Daten aus Tabelle rcjv_app.institution: 0 rows
/*!40000 ALTER TABLE `institution` DISABLE KEYS */;
/*!40000 ALTER TABLE `institution` ENABLE KEYS */;

-- Exportiere Struktur von Tabelle rcjv_app.league
CREATE TABLE IF NOT EXISTS `league` (
  `league_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `league` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`league_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf16 COLLATE=utf16_general_ci;

-- Exportiere Daten aus Tabelle rcjv_app.league: 0 rows
/*!40000 ALTER TABLE `league` DISABLE KEYS */;
/*!40000 ALTER TABLE `league` ENABLE KEYS */;

-- Exportiere Struktur von Tabelle rcjv_app.participants
CREATE TABLE IF NOT EXISTS `participants` (
  `participants_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `first_name` varchar(50) DEFAULT NULL,
  `last_name` varchar(50) DEFAULT NULL,
  `team_id` bigint(20) DEFAULT NULL,
  `institution_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`participants_id`),
  KEY `team_id` (`team_id`),
  KEY `institution_id` (`institution_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf16 COLLATE=utf16_general_ci;

-- Exportiere Daten aus Tabelle rcjv_app.participants: 0 rows
/*!40000 ALTER TABLE `participants` DISABLE KEYS */;
/*!40000 ALTER TABLE `participants` ENABLE KEYS */;

-- Exportiere Struktur von Tabelle rcjv_app.teams
CREATE TABLE IF NOT EXISTS `teams` (
  `team_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL,
  `league_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`team_id`),
  KEY `league_id` (`league_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf16 COLLATE=utf16_general_ci;

-- Exportiere Daten aus Tabelle rcjv_app.teams: 0 rows
/*!40000 ALTER TABLE `teams` DISABLE KEYS */;
/*!40000 ALTER TABLE `teams` ENABLE KEYS */;

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
