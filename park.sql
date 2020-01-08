-- MySQL dump 10.13  Distrib 8.0.17, for Linux (x86_64)
--
-- Host: localhost    Database: parkings
-- ------------------------------------------------------
-- Server version	8.0.17

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `credentials`
--

DROP TABLE IF EXISTS `credentials`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `credentials` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` text,
  `password` text,
  `email` text,
  `points` int(11) DEFAULT NULL,
  `role` text,
  `token` text,
  `expired` text,
  `created_at` text,
  `modified_at` text,
  `deleted_at` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=90 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `credentials`
--

LOCK TABLES `credentials` WRITE;
/*!40000 ALTER TABLE `credentials` DISABLE KEYS */;
INSERT INTO `credentials` VALUES (73,'huydv71271w111','a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3','dohuy171@gmail.com',0,'customer','','','2020-01-04T11:30:57+07:00','',''),(74,'huydv71271w1111','a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3','dohuy171@gmail.com',0,'customer','','','2020-01-04T11:31:17+07:00','',''),(75,'userna','5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8','testemail@gmail.com',0,'customer','','','2020-01-04T11:35:48+07:00','',''),(76,'huydv','a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3','dohuy172@gmail.com',0,'admin','','','2020-01-04T11:43:24+07:00','',''),(77,'hdhshshs','bfa66c40b459b383de1d3abec61ef412399ef5b94d47493a27f323c162d591ff','hddjjs@gmail.com',0,'customer','','','2020-01-04T11:49:31+07:00','',''),(78,'usernssddddsa','5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8','testemagggil@gmail.com',0,'customer','','','2020-01-04T12:08:32+07:00','',''),(79,'test','2c2e309f08dada7a3063bdf03bb9739d0375759035bc4debbb570358f09edcbc','gsd3fg@gmail.com',0,'customer','','','2020-01-04T12:11:36+07:00','',''),(80,'testfff','2c2e309f08dada7a3063bdf03bb9739d0375759035bc4debbb570358f09edcbc','gsd34fg@gmail.com',0,'customer','','','2020-01-04T12:11:45+07:00','',''),(81,'thinhnp5','0e02d3e840eab7d2f8acf30918ed7d3bdaeeb816a87d8659eb3f5b668d702ad7','npthinh124@gmail.com',0,'customer','','','2020-01-04T15:42:57+07:00','',''),(82,'test123','28cb017dfc99073aa1b47c1b30f413e3ce774c4991eb4158de50f9dbb36d8043','gasd123g@gmail.com',0,'customer','','','2020-01-04T17:28:46+07:00','',''),(83,'thinhnp55','0e02d3e840eab7d2f8acf30918ed7d3bdaeeb816a87d8659eb3f5b668d702ad7','dohuy1721@gmail.com',0,'customer','','','2020-01-04T17:30:10+07:00','',''),(84,'test1234','937e8d5fbb48bd4949536cd65b8d35c426b80d2f830c5c308e2cdec422ae2244','sfasdgd2g@gmail.com',0,'customer','','','2020-01-04T22:59:02+07:00','',''),(85,'huydo','7651d3a67ab074886f7a072906ad8ae7cbbdaac69cbe67c89bf4d1109d22aa26','dohuy173@gmail.com',0,'customer','','','2020-01-05T14:22:06+07:00','',''),(86,'user1','932f3c1b56257ce8539ac269d7aab42550dacf8818d075f0bdf1990562aae3ef','dfaf1dfasf@gmail.com',0,'customer','','','2020-01-05T16:28:22+07:00','',''),(87,'user2','932f3c1b56257ce8539ac269d7aab42550dacf8818d075f0bdf1990562aae3ef','dfaf2f3@gmail.com',0,'customer','','','2020-01-05T17:35:20+07:00','',''),(88,'huydvaaa','a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3','dohuy177@gmail.com',0,'customer','','','2020-01-08T20:14:25+07:00','',''),(89,'huydv212aaa','a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3','dohuy277@gmail.com',0,'customer','','','2020-01-08T20:17:58+07:00','','');
/*!40000 ALTER TABLE `credentials` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `owners`
--

DROP TABLE IF EXISTS `owners`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `owners` (
  `credentialId` int(11) NOT NULL,
  `fullName` text,
  `cmndImage` text,
  `address` text,
  `phoneNumber` text,
  `status` text,
  `created_at` text,
  `modified_at` text,
  PRIMARY KEY (`credentialId`),
  CONSTRAINT `fk_credential1` FOREIGN KEY (`credentialId`) REFERENCES `credentials` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `owners`
--

LOCK TABLES `owners` WRITE;
/*!40000 ALTER TABLE `owners` DISABLE KEYS */;
INSERT INTO `owners` VALUES (81,'Nguyễn Phu Thinh','http://localhost:8085/(2020-01-07T20:18:43+07:00)','412 Ngoc Thụy','0123435235','ENABLE','2020-01-07T20:18:43+07:00','');
/*!40000 ALTER TABLE `owners` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `parkings`
--

DROP TABLE IF EXISTS `parkings`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `parkings` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `parkingName` text,
  `properties` text,
  `address` text,
  `kindOf` text,
  `payment` text,
  `longitude` text,
  `latitude` text,
  `capacity` text,
  `ownerId` int(11) DEFAULT NULL,
  `blockAmount` int(11) DEFAULT NULL,
  `created_at` text,
  `modified_at` text,
  `deleted_at` text,
  `describe` text CHARACTER SET utf8mb4,
  `parkingImages` text,
  `certificateOfland` text,
  `status` text,
  PRIMARY KEY (`id`),
  KEY `fk_owner` (`ownerId`),
  CONSTRAINT `fk_owner` FOREIGN KEY (`ownerId`) REFERENCES `owners` (`credentialId`)
) ENGINE=InnoDB AUTO_INCREMENT=43 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `parkings`
--

LOCK TABLES `parkings` WRITE;
/*!40000 ALTER TABLE `parkings` DISABLE KEYS */;
INSERT INTO `parkings` VALUES (39,'458 Phố Minh Khai',',Gửi qua đêm,Có mái che','458 Phố Minh Khai','Điểm đỗ ô tô, xe máy','Tiền mặt','105.8671399','20.997256','50/50',81,50000,'2020-01-07T20:20:02+07:00','','','mo ta','http://localhost:8085/(2020-01-07T20:20:01+07:00)','','APPROVED'),(40,'456 Phố Minh Khai',',Gửi qua đêm,Có mái che','456 Phố Minh Khai','Điểm đỗ xe máy','Tiền mặt','105.868188','20.997611','20/20',81,10000,'2020-01-07T20:58:18+07:00','','','a','http://192.168.137.221:8085/(2020-01-07T20:58:18+07:00)','','APPROVED');
/*!40000 ALTER TABLE `parkings` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ratings`
--

DROP TABLE IF EXISTS `ratings`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ratings` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `stars` int(11) DEFAULT NULL,
  `credentialId` int(11) DEFAULT NULL,
  `parkingId` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ratings`
--

LOCK TABLES `ratings` WRITE;
/*!40000 ALTER TABLE `ratings` DISABLE KEYS */;
INSERT INTO `ratings` VALUES (1,3,85,35),(2,4,85,36),(3,5,85,38);
/*!40000 ALTER TABLE `ratings` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transactions`
--

DROP TABLE IF EXISTS `transactions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `transactions` (
  `credentialId` int(11) NOT NULL,
  `parkingId` int(11) NOT NULL,
  `liencePlate` text,
  `session` text,
  `startTime` text,
  `endTime` text,
  `amount` int(11) DEFAULT NULL,
  `status` text,
  `reasonMsg` text,
  `created_at` text,
  `modified_at` text,
  PRIMARY KEY (`credentialId`,`parkingId`),
  KEY `fk_parking` (`parkingId`),
  CONSTRAINT `fk_credential3` FOREIGN KEY (`credentialId`) REFERENCES `credentials` (`id`),
  CONSTRAINT `fk_parking` FOREIGN KEY (`parkingId`) REFERENCES `parkings` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transactions`
--

LOCK TABLES `transactions` WRITE;
/*!40000 ALTER TABLE `transactions` DISABLE KEYS */;
/*!40000 ALTER TABLE `transactions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `userDevices`
--

DROP TABLE IF EXISTS `userDevices`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `userDevices` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `deviceToken` text,
  `credentialId` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_credential2` (`credentialId`),
  CONSTRAINT `fk_credential2` FOREIGN KEY (`credentialId`) REFERENCES `credentials` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `userDevices`
--

LOCK TABLES `userDevices` WRITE;
/*!40000 ALTER TABLE `userDevices` DISABLE KEYS */;
/*!40000 ALTER TABLE `userDevices` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-01-08 13:34:44
