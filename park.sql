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
) ENGINE=InnoDB AUTO_INCREMENT=94 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `credentials`
--

LOCK TABLES `credentials` WRITE;
/*!40000 ALTER TABLE `credentials` DISABLE KEYS */;
INSERT INTO `credentials` VALUES (73,'huydv71271w111','a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3','dohuy171@gmail.com',0,'customer','','','2020-01-04T11:30:57+07:00','',''),(74,'huydv71271w1111','a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3','dohuy171@gmail.com',0,'customer','','','2020-01-04T11:31:17+07:00','',''),(75,'userna','5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8','testemail@gmail.com',0,'customer','','','2020-01-04T11:35:48+07:00','',''),(76,'huydv','a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3','dohuy172@gmail.com',0,'admin','','','2020-01-04T11:43:24+07:00','',''),(77,'hdhshshs','bfa66c40b459b383de1d3abec61ef412399ef5b94d47493a27f323c162d591ff','hddjjs@gmail.com',0,'customer','','','2020-01-04T11:49:31+07:00','',''),(78,'usernssddddsa','5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8','testemagggil@gmail.com',0,'customer','','','2020-01-04T12:08:32+07:00','',''),(79,'test','2c2e309f08dada7a3063bdf03bb9739d0375759035bc4debbb570358f09edcbc','gsd3fg@gmail.com',0,'customer','','','2020-01-04T12:11:36+07:00','',''),(80,'testfff','2c2e309f08dada7a3063bdf03bb9739d0375759035bc4debbb570358f09edcbc','gsd34fg@gmail.com',0,'customer','','','2020-01-04T12:11:45+07:00','',''),(81,'thinhnp5','0e02d3e840eab7d2f8acf30918ed7d3bdaeeb816a87d8659eb3f5b668d702ad7','npthinh124@gmail.com',0,'customer','','','2020-01-04T15:42:57+07:00','',''),(82,'test123','28cb017dfc99073aa1b47c1b30f413e3ce774c4991eb4158de50f9dbb36d8043','gasd123g@gmail.com',0,'customer','','','2020-01-04T17:28:46+07:00','',''),(83,'thinhnp55','0e02d3e840eab7d2f8acf30918ed7d3bdaeeb816a87d8659eb3f5b668d702ad7','dohuy1721@gmail.com',0,'customer','','','2020-01-04T17:30:10+07:00','',''),(84,'test1234','937e8d5fbb48bd4949536cd65b8d35c426b80d2f830c5c308e2cdec422ae2244','sfasdgd2g@gmail.com',0,'customer','','','2020-01-04T22:59:02+07:00','',''),(85,'huydo','7651d3a67ab074886f7a072906ad8ae7cbbdaac69cbe67c89bf4d1109d22aa26','dohuy173@gmail.com',0,'customer','','','2020-01-05T14:22:06+07:00','',''),(86,'user1','932f3c1b56257ce8539ac269d7aab42550dacf8818d075f0bdf1990562aae3ef','dfaf1dfasf@gmail.com',0,'customer','','','2020-01-05T16:28:22+07:00','',''),(87,'user2','932f3c1b56257ce8539ac269d7aab42550dacf8818d075f0bdf1990562aae3ef','dfaf2f3@gmail.com',0,'customer','','','2020-01-05T17:35:20+07:00','',''),(88,'huydvaaa','a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3','dohuy177@gmail.com',0,'customer','','','2020-01-08T20:14:25+07:00','',''),(89,'huydv212aaa','a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3','dohuy277@gmail.com',0,'customer','','','2020-01-08T20:17:58+07:00','',''),(90,'dohuy','ef797c8118f02dfb649607dd5d3f8c7623048c9c063d532cc95c5ed7a898a64f','dohuy1677@gmail.com',0,'customer','','','2020-01-12T17:52:42+07:00','',''),(91,'thinhnp96','932f3c1b56257ce8539ac269d7aab42550dacf8818d075f0bdf1990562aae3ef','dfasgv4t4fef@gmail.com',0,'customer','','','2020-01-12T17:52:53+07:00','',''),(92,'test9999','932f3c1b56257ce8539ac269d7aab42550dacf8818d075f0bdf1990562aae3ef','dsfag4gdwag@gmail.com',0,'customer','','','2020-01-13T00:07:27+07:00','',''),(93,'host9999','932f3c1b56257ce8539ac269d7aab42550dacf8818d075f0bdf1990562aae3ef','dvasvaf4v@gmail.com',0,'customer','','','2020-01-13T00:07:59+07:00','','');
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
INSERT INTO `owners` VALUES (81,'Nguyễn Phu Thinh','http://localhost:8085/(2020-01-07T20:18:43+07:00)','412 Ngoc Thụy','0123435235','DISABLED','2020-01-07T20:18:43+07:00','2020-01-12T17:34:57+07:00'),(86,'nguyen phu thinh','http://localhost:8085/(2020-01-12T13:35:21+07:00)','ha noi','099999999','ENABLE','2020-01-12T13:35:21+07:00',''),(91,'Nguyen Phu Thinh','http://10.124.0.32:8085/(2020-01-12T17:53:42+07:00)','HN','0379395295','ENABLE','2020-01-12T17:53:42+07:00',''),(93,'Nguyễn Phu Thịnh','http://10.124.0.32:8085/(2020-01-13T00:09:03+07:00)','HN','0379395295','ENABLE','2020-01-13T00:09:03+07:00','');
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
) ENGINE=InnoDB AUTO_INCREMENT=59 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `parkings`
--

LOCK TABLES `parkings` WRITE;
/*!40000 ALTER TABLE `parkings` DISABLE KEYS */;
INSERT INTO `parkings` VALUES (39,'458 Phố Minh Khai',',Gửi qua đêm,Có mái che','458 Phố Minh Khai','Điểm đỗ ô tô, xe máy','Tiền mặt','105.8671399','20.997256','50/50',81,50000,'2020-01-07T20:20:02+07:00','','','mo ta','http://localhost:8085/(2020-01-07T20:20:01+07:00)','','APPROVED'),(40,'456 Phố Minh Khai',',Gửi qua đêm,Có mái che','456 Phố Minh Khai','Điểm đỗ xe máy','Tiền mặt','105.868188','20.997611','20/20',81,10000,'2020-01-07T20:58:18+07:00','','','a','http://192.168.137.221:8085/(2020-01-07T20:58:18+07:00)','','APPROVED'),(43,'500 Phố Minh Khai',',Đỗ qua đêm,Mái che','500 Phố Minh Khai','Điểm đỗ ô tô, xe máy','Hỗ trợ cả tiền mặt và điểm','105.868888','20.998487','50/50',86,20000,'2020-01-12T13:36:27+07:00','','','doi dien time city','http://localhost:8085/(2020-01-12T13:36:27+07:00)','','APPROVED'),(44,'346 Bến Vân Đồn','','346 Bến Vân Đồn','Điểm đỗ xe máy','Tiền mặt','106.6924247','10.757158','44/44',86,44,'2020-01-12T15:03:43+07:00','','','','http://10.124.0.32:8085/(2020-01-12T15:03:43+07:00)','','APPROVED'),(45,'600 Phố Minh Khai',',Đỗ qua đêm,Mái che','600 Phố Minh Khai','Điểm đỗ xe máy','Hỗ trợ cả tiền mặt và điểm','105.8705111','20.9990473','50/50',81,50000,'2020-01-12T17:02:10+07:00','2020-01-12T17:02:30+07:00','','','http://10.124.0.32:8085/(2020-01-12T17:02:09+07:00)','','APPROVED'),(46,'25 Phố Minh Khai',',Đỗ qua đêm,Mái che','25 Phố Minh Khai','Điểm đỗ ô tô, xe máy','Hỗ trợ cả tiền mặt và điểm','105.851579','20.995777','50/50',91,20000,'2020-01-12T17:55:03+07:00','2020-01-12T17:55:21+07:00','','gan time city','http://10.124.0.32:8085/(2020-01-12T17:55:02+07:00)','','APPROVED'),(47,'412 Ngọc Thuỵ',',Đỗ qua đêm,Mái che','412 Ngọc Thuỵ','Điểm đỗ ô tô, xe máy','Tiền mặt','105.8650209','21.0626644','50/50',86,100000,'2020-01-12T22:19:29+07:00','','','Nha có cửa đỏ','http://10.124.0.32:8085/(2020-01-12T22:19:29+07:00)','','APPROVED'),(48,'22 Minh Khai',',Đỗ qua đêm,Mái che','22 Minh Khai','Điểm đỗ ô tô, xe máy','Hỗ trợ cả tiền mặt và điểm','106.684977','20.860589','50/50',86,20000,'2020-01-12T22:54:06+07:00','','','aaaaa','http://10.124.0.32:8085/(2020-01-12T22:54:06+07:00)','','APPROVED'),(49,'200 Phố Minh Khai',',Đỗ qua đêm,Mái che','200 Phố Minh Khai','Điểm đỗ ô tô, xe máy','Hỗ trợ cả tiền mặt và điểm','105.8544738','20.9952696','50/50',81,50,'2020-01-12T23:03:36+07:00','','','','http://10.124.0.32:8085/(2020-01-12T22:59:40+07:00)','','APPROVED'),(50,'100 Phố Minh Khai','Camera,Đỗ qua đêm,Mái che','100 Phố Minh Khai','Điểm đỗ ô tô, xe máy','Hỗ trợ cả tiền mặt và điểm','105.8520615','20.9955449','50/50',81,20000,'2020-01-12T23:04:29+07:00','','','','http://10.124.0.32:8085/(2020-01-12T23:04:29+07:00)','','APPROVED'),(51,'3 Minh Khai','Camera,Đỗ qua đêm,Mái che','3 Minh Khai','Điểm đỗ ô tô, xe máy','Hỗ trợ cả tiền mặt và điểm','106.685146','20.863504','22/22',81,20000,'2020-01-12T23:20:10+07:00','','','','http://10.124.0.32:8085/(2020-01-12T23:19:38+07:00)','','APPROVED'),(52,'45 Phố Minh Khai','a,b,c','45 Phố Minh Khai','Điểm đỗ xe máy','Tiền mặt','105.851967','20.995729','55/55',81,5555,'2020-01-12T23:26:43+07:00','','','','http://10.124.0.32:8085/(2020-01-12T23:26:43+07:00)','','APPROVED'),(53,'60 Minh Khai','Camera,Đỗ qua đêm,Mái che','60 Minh Khai','Điểm đỗ xe máy','Tiền mặt','106.6851234','20.8589675','50/50',81,7668,'2020-01-12T23:28:58+07:00','','','','http://10.124.0.32:8085/(2020-01-12T23:28:57+07:00)','','APPROVED'),(54,'50 Bạch Mai','Camera,Đỗ qua đêm,Mái che','50 Bạch Mai','Điểm đỗ ô tô, xe máy','Hỗ trợ cả tiền mặt và điểm','105.851546','20.9987229','50/50',81,5,'2020-01-12T23:34:19+07:00','','','','http://10.124.0.32:8085/(2020-01-12T23:34:19+07:00)','','APPROVED'),(55,'100 Phố Minh Khai','','100 Phố Minh Khai','Điểm đỗ xe máy','Tiền mặt','105.8520615','20.9955449','50/50',81,50000,'2020-01-12T23:43:19+07:00','','','','http://10.124.0.32:8085/(2020-01-12T23:43:18+07:00)','','APPROVED'),(56,'255 Bạch Mai','','255 Bạch Mai','Điểm đỗ xe máy','Tiền mặt','105.8512886','21.0033936','50/50',81,5555,'2020-01-12T23:45:36+07:00','','','','http://10.124.0.32:8085/(2020-01-12T23:45:36+07:00)','','APPROVED'),(57,'12 Đại Cồ Việt','Camera,Đỗ qua đêm,Mái che','12 Đại Cồ Việt','Điểm đỗ ô tô, xe máy','Điểm','105.850379','21.008759','50/50',93,50000,'2020-01-13T00:16:31+07:00','','','gan bach mai','http://10.124.0.32:8085/(2020-01-13T00:16:31+07:00)','','APPROVED'),(58,'20 Đại Cồ Việt','','20 Đại Cồ Việt','Điểm đỗ ô tô','Điểm','105.850003','21.008711','40/40',93,15000,'2020-01-13T00:17:58+07:00','','','','http://10.124.0.32:8085/(2020-01-13T00:17:57+07:00)','','APPROVED');
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
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ratings`
--

LOCK TABLES `ratings` WRITE;
/*!40000 ALTER TABLE `ratings` DISABLE KEYS */;
INSERT INTO `ratings` VALUES (4,3,81,43),(5,5,81,43),(6,4,81,43),(7,4,81,0),(8,5,81,0),(9,2,81,0),(10,2,92,26),(11,1,92,58),(12,5,92,58),(13,4,92,58),(14,4,92,58);
/*!40000 ALTER TABLE `ratings` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transactions`
--

DROP TABLE IF EXISTS `transactions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `transactions` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `credentialId` int(11) NOT NULL,
  `parkingId` int(11) NOT NULL,
  `licence` text,
  `session` text,
  `phoneNumber` text,
  `requestedPayment` text,
  `startTime` text,
  `endTime` text,
  `amount` int(11) DEFAULT NULL,
  `status` text,
  `reasonMsg` text,
  `created_at` text,
  `modified_at` text,
  PRIMARY KEY (`id`),
  KEY `fk_credential` (`credentialId`),
  KEY `fk_parking` (`parkingId`),
  CONSTRAINT `fk_credential` FOREIGN KEY (`credentialId`) REFERENCES `credentials` (`id`),
  CONSTRAINT `fk_parking` FOREIGN KEY (`parkingId`) REFERENCES `parkings` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transactions`
--

LOCK TABLES `transactions` WRITE;
/*!40000 ALTER TABLE `transactions` DISABLE KEYS */;
INSERT INTO `transactions` VALUES (1,81,39,'727w7\n30f9','1','079797070','Tiền mặt','2020-01-12T13:45:00.000Z','2020-01-12T14:45:00.000Z',50000,'5','','2020-01-12T13:29:13+07:00',''),(2,81,43,'rdkdkd','8','0329184025','Điểm','2020-01-12T13:45:00.000Z','2020-01-12T21:45:00.000Z',160000,'5','','2020-01-12T13:38:01+07:00',''),(3,86,45,'30f94244','2','0379395295','Điểm','2020-01-12T17:15:00.000Z','2020-01-12T19:15:00.000Z',100000,'5','','2020-01-12T17:05:55+07:00',''),(4,90,46,'30z1 9529','2','0967240005','Tiền mặt','2020-01-12T17:00:00.000Z','2020-01-12T19:00:00.000Z',40000,'5','','2020-01-12T17:56:43+07:00',''),(5,91,46,'30A9999','1','0912412424','Điểm','2020-01-12T17:00:00.000Z','2020-01-12T18:00:00.000Z',20000,'5','','2020-01-12T17:58:29+07:00',''),(6,91,45,'30F94244','2','0379395295','Điểm','2020-01-12T19:45:00.000Z','2020-01-12T21:45:00.000Z',100000,'5','','2020-01-12T19:42:45+07:00',''),(7,91,45,'30a9999','2','0913024433','Điểm','2020-01-12T19:45:00.000Z','2020-01-12T21:45:00.000Z',100000,'4','','2020-01-12T19:46:53+07:00',''),(8,81,43,'30f1111','2','0913053311','Điểm','2020-01-12T20:00:00.000Z','2020-01-12T22:00:00.000Z',40000,'5','','2020-01-12T20:00:37+07:00',''),(9,81,45,'30k943833','2','0976407646','Điểm','2020-01-12T20:00:00.000Z','2020-01-12T22:00:00.000Z',100000,'4','','2020-01-12T20:01:36+07:00',''),(10,81,40,'30jsjsj2','2','0913649944','Tiền mặt','2020-01-12T20:00:00.000Z','2020-01-12T22:00:00.000Z',20000,'4','','2020-01-12T20:02:22+07:00',''),(11,81,40,'30sjsjsj','1','097979797','Tiền mặt','2020-01-12T20:15:00.000Z','2020-01-12T21:15:00.000Z',10000,'4','','2020-01-12T20:07:34+07:00',''),(12,81,40,'snsnsn','2','70707070','Tiền mặt','2020-01-12T20:30:00.000Z','2020-01-12T22:30:00.000Z',20000,'4','','2020-01-12T20:11:52+07:00',''),(13,81,40,'bsbs','1','777007','Tiền mặt','2020-01-12T20:15:00.000Z','2020-01-12T21:45:00.000Z',10000,'3','','2020-01-12T20:16:17+07:00',''),(14,81,40,'bssbsh','0','09090909','Tiền mặt','2020-01-12T22:45:00.000Z','2020-01-12T22:45:00.000Z',0,'4','','2020-01-12T20:27:57+07:00',''),(15,81,47,'30z1 9529','0','0967240005','Tiền mặt','2020-01-12T22:45:00.000Z','2020-01-12T22:45:00.000Z',0,'4','','2020-01-12T22:22:53+07:00',''),(16,81,47,'30f9422','0','0967045577','Tiền mặt','2020-01-14T22:45:00.000Z','2020-01-14T22:45:00.000Z',0,'4','','2020-01-12T22:27:47+07:00',''),(17,81,47,'30f3333','1','0913053311','Tiền mặt','2020-01-12T22:45:00.000Z','2020-01-12T23:45:00.000Z',100000,'4','','2020-01-12T22:38:53+07:00',''),(18,81,47,'30f9333','23','097979797','Tiền mặt','2020-01-12T23:45:00.000Z','2020-01-13T22:45:00.000Z',2300000,'4','','2020-01-12T22:39:36+07:00',''),(19,81,47,'gsga','1','079797','Tiền mặt','2020-01-12T22:45:00.000Z','2020-01-12T23:45:00.000Z',100000,'4','','2020-01-12T22:42:46+07:00',''),(20,81,47,'hwus','23','9797','Tiền mặt','2020-01-12T23:45:00.000Z','2020-01-13T22:45:00.000Z',2300000,'4','','2020-01-12T22:43:56+07:00',''),(21,81,47,'hshw','93','9797','Tiền mặt','2020-01-13T01:00:00.000Z','2020-01-16T22:45:00.000Z',9300000,'5','','2020-01-12T22:44:38+07:00',''),(22,92,57,'30Z1-9529','1','0967240005','Điểm','2020-01-13T01:45:00.000Z','2020-01-13T02:45:00.000Z',50000,'1','','2020-01-13T00:47:46+07:00',''),(23,92,57,'30z1-9529','1','0967240005','Điểm','2020-01-13T01:45:00.000Z','2020-01-13T02:45:00.000Z',50000,'4','','2020-01-13T00:54:54+07:00',''),(24,92,57,'30f94244','1','0912312412','Điểm','2020-01-13T01:45:00.000Z','2020-01-13T02:45:00.000Z',50000,'5','','2020-01-13T01:02:57+07:00',''),(25,92,57,'30F94244','1','0913053311','Điểm','2020-01-13T03:30:00.000Z','2020-01-13T04:30:00.000Z',50000,'5','','2020-01-13T01:11:48+07:00',''),(26,92,58,'30a99999','1','0913053311','Điểm','2020-01-13T13:45:00.000Z','2020-01-13T14:45:00.000Z',15000,'5','','2020-01-13T01:30:00+07:00','');
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

-- Dump completed on 2020-01-13  9:40:37
