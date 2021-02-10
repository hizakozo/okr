/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
SET NAMES utf8mb4;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
SET
character_set_client = utf8mb4 ;
CREATE TABLE `user`
(
    `user_id`   int(11) NOT NULL AUTO_INCREMENT,
    `user_name` varchar(225) NOT NULL,
    `is_delete` tinyint(1) NOT NULL DEFAULT '1',
    `create_at` timestamp default current_timestamp,
    `update_at` timestamp default current_timestamp on update current_timestamp,
    PRIMARY KEY (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

DROP TABLE IF EXISTS `auth`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
SET
character_set_client = utf8mb4 ;
CREATE TABLE `auth`
(
    `auth_id`      int(11) NOT NULL AUTO_INCREMENT,
    `user_id`      int(11) NOT NULL,
    `login_id`     varchar(225) NOT NULL UNIQUE,
    `password`     varchar(223) NOT NULL,
    `mail_address` varchar(223) NOT NULL,
    `create_at`    timestamp default current_timestamp,
    `update_at`    timestamp default current_timestamp on update current_timestamp,
    FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`) ON DELETE CASCADE,
    PRIMARY KEY (`auth_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

DROP TABLE IF EXISTS `company`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
SET
character_set_client = utf8mb4 ;
CREATE TABLE `company`
(
    `company_id` int(11) NOT NULL AUTO_INCREMENT,
    `company_name`   varchar(225) NOT NULL,
    `create_at`   timestamp default current_timestamp,
    `update_at`   timestamp default current_timestamp on update current_timestamp,
    PRIMARY KEY (`company_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

DROP TABLE IF EXISTS `department`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
SET
character_set_client = utf8mb4 ;
CREATE TABLE `department`
(
    `department_id` int(11) NOT NULL AUTO_INCREMENT,
    `company_id` int(11) NOT NULL,
    `department_name`   varchar(225) NOT NULL,
    `create_at`   timestamp default current_timestamp,
    `update_at`   timestamp default current_timestamp on update current_timestamp,
    FOREIGN KEY (`company_id`) REFERENCES `company` (`company_id`) ON DELETE CASCADE,
    PRIMARY KEY (`department_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

DROP TABLE IF EXISTS `user_department`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
SET
character_set_client = utf8mb4 ;
CREATE TABLE `user_department`
(
    `user_department_id` int(11) NOT NULL AUTO_INCREMENT,
    `user_id` int(11) NOT NULL,
    `department_id` int(11) NOT NULL,
    `department_name`   varchar(225) NOT NULL,
    `create_at`   timestamp default current_timestamp,
    `update_at`   timestamp default current_timestamp on update current_timestamp,
    FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`) ON DELETE CASCADE,
    FOREIGN KEY (`department_id`) REFERENCES `department` (`department_id`) ON DELETE CASCADE,
    PRIMARY KEY (`user_department_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

DROP TABLE IF EXISTS `kind`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
SET
character_set_client = utf8mb4 ;
CREATE TABLE `kind`
(
    `kind_id` int(11) NOT NULL AUTO_INCREMENT,
    `kind` int(11) NOT NULL,
    `relation_id`   varchar(225) NOT NULL,
    `create_at`   timestamp default current_timestamp,
    `update_at`   timestamp default current_timestamp on update current_timestamp,
    PRIMARY KEY (`kind_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

DROP TABLE IF EXISTS `objective`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
SET
character_set_client = utf8mb4 ;
CREATE TABLE `objective`
(
    `objective_id` int(11) NOT NULL AUTO_INCREMENT,
    `key_result_id` int(11),
    `kind_id` int(11) NOT NULL,
    `content`   varchar(225) NOT NULL,
    `create_at`   timestamp default current_timestamp,
    `update_at`   timestamp default current_timestamp on update current_timestamp,
    FOREIGN KEY (`kind_id`) REFERENCES `kind` (`kind_id`) ON DELETE CASCADE,
    PRIMARY KEY (`objective_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

DROP TABLE IF EXISTS `key_result`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
SET
character_set_client = utf8mb4 ;
CREATE TABLE `key_result`
(
    `key_result_id` int(11) NOT NULL AUTO_INCREMENT,
    `objective_id` int(11) NOT NULL,
    `content`   varchar(225) NOT NULL,
    `create_at`   timestamp default current_timestamp,
    `update_at`   timestamp default current_timestamp on update current_timestamp,
    FOREIGN KEY (`objective_id`) REFERENCES `objective` (`objective_id`) ON DELETE CASCADE,
    PRIMARY KEY (`key_result_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;