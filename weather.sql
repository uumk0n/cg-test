-- Adminer 4.8.1 MySQL 8.3.0 dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

DROP TABLE IF EXISTS `actions`;
CREATE TABLE `actions` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `action_time` bigint NOT NULL,
  `request_result` int DEFAULT NULL,
  `temp_c` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `actions_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `actions` (`id`, `user_id`, `action_time`, `request_result`, `temp_c`) VALUES
(13,	2,	1706710987284,	200,	10),
(14,	2,	1706711407617,	200,	10);

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `login` varchar(255) NOT NULL,
  `hashedPassword` varchar(255) NOT NULL,
  `fio` varchar(255) NOT NULL,
  `apiToken` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `login` (`login`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `user` (`id`, `login`, `hashedPassword`, `fio`, `apiToken`) VALUES
(1,	'uumkon',	'$2b$10$Qhh9FD54Iiyk.fNYRRUMzu0ZzzU8eG9XARrolJFeAeqL7FJL2BMWi',	'ilya',	'5ab658b5-2426-4665-8493-1c43a2cdab63'),
(2,	'ilya',	'$2b$10$RqyFPsdF1SfIFuAsKVJEcePQRcODHmIKKjd/zIaxLh3Sn9YyQGcI2',	'ilya',	'44727b37-4b9b-4038-8720-b08814520add');

-- 2024-01-31 14:31:42
