-- MySQL dump 10.13  Distrib 8.0.32, for Win64 (x86_64)
--
-- Host: localhost    Database: fifa
-- ------------------------------------------------------
-- Server version	8.0.28

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `confederacions`
--

DROP TABLE IF EXISTS `confederacions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `confederacions` (
  `id` int NOT NULL AUTO_INCREMENT,
  `nombre` varchar(350) NOT NULL,
  `logo` varchar(350) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `confederacions`
--

LOCK TABLES `confederacions` WRITE;
/*!40000 ALTER TABLE `confederacions` DISABLE KEYS */;
INSERT INTO `confederacions` VALUES (1,'Conmebol','https://upload.wikimedia.org/wikipedia/commons/thumb/4/49/Logo_de_la_Conmebol.svg/888px-Logo_de_la_Conmebol.svg.png'),(2,'EUFA','https://2.bp.blogspot.com/-inJxt3l1Ggc/Welu2Bmat6I/AAAAAAABP7w/fQ3jJUftH7IW68jzT0VzLNUe7ZhkxpmTACLcBGAs/s1600/UEFA.png');
/*!40000 ALTER TABLE `confederacions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `edicion_equipos`
--

DROP TABLE IF EXISTS `edicion_equipos`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `edicion_equipos` (
  `id` int NOT NULL AUTO_INCREMENT,
  `id_edicion_torneo` int NOT NULL,
  `id_equipo` int NOT NULL,
  `nombre` int NOT NULL,
  `escudo` int NOT NULL,

  PRIMARY KEY (`id`),
  KEY `fk_edTorneo_idx` (`id_edicion_torneo`),
  KEY `fk_equipo_idx` (`id_equipo`),
  CONSTRAINT `fk_edTorneo` FOREIGN KEY (`id_edicion_torneo`) REFERENCES `edicion_torneos` (`id`),
  CONSTRAINT `fk_equipo` FOREIGN KEY (`id_equipo`) REFERENCES `equipos` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `edicion_equipos`
--

LOCK TABLES `edicion_equipos` WRITE;
/*!40000 ALTER TABLE `edicion_equipos` DISABLE KEYS */;
/*!40000 ALTER TABLE `edicion_equipos` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `edicion_torneos`
--

DROP TABLE IF EXISTS `edicion_torneos`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `edicion_torneos` (
  `id` int NOT NULL AUTO_INCREMENT,
  `id_torneo` int NOT NULL,
  `anio` int NOT NULL,
  `campeon` int DEFAULT NULL,
  `subcampeon` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_campeon_idx` (`campeon`),
  KEY `fk_subcampeon_idx` (`subcampeon`),
  KEY `fk_torneo_idx` (`id_torneo`),
  CONSTRAINT `fk_campeon` FOREIGN KEY (`campeon`) REFERENCES `equipos` (`id`),
  CONSTRAINT `fk_subcampeon` FOREIGN KEY (`subcampeon`) REFERENCES `equipos` (`id`),
  CONSTRAINT `fk_torneo` FOREIGN KEY (`id_torneo`) REFERENCES `torneos` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `edicion_torneos`
--

LOCK TABLES `edicion_torneos` WRITE;
/*!40000 ALTER TABLE `edicion_torneos` DISABLE KEYS */;
/*!40000 ALTER TABLE `edicion_torneos` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `equipos`
--

DROP TABLE IF EXISTS `equipos`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `equipos` (
  `id` int NOT NULL AUTO_INCREMENT,
  `nombre` varchar(350) NOT NULL,
  `escudo` varchar(350) NOT NULL,
  `id_pais` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_pais_idx` (`id_pais`),
  CONSTRAINT `fk_pais` FOREIGN KEY (`id_pais`) REFERENCES `pais` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=131 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `equipos`
--

LOCK TABLES `equipos` WRITE;
/*!40000 ALTER TABLE `equipos` DISABLE KEYS */;
INSERT INTO `equipos` VALUES (1,'Boca Juniors','https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEjOL0rL1mnaqAIv_9cWnQayBcfquddQxXR82Ho7GG9YdcQjYhK077s3wlCOM1Z1OExJShV-b-f3Yzuoq049fzMYoRvgwR06s99-ezsl4uki7P80_dKIxjtIO4kDkpJNbJTxG-K2jS8xkKzvkTgkVXfy2sosC1ZROpCqNkNCiGJBhKbV9_u86I0Bq4wsNgM/s128/Boca%20Juniors128x.png',1),(2,'River Plate','https://blogger.googleusercontent.com/img/a/AVvXsEgTBK7xhnKu2_r_QmZQ1cTNMJsILgHkg9qmEPHvBzWp1YNXQ3yacqii16J9eEz6cUWXrd74cT0d4YE6vTg4bl9rI8zpYAR0AFadmDoF97IZyeHTCSauM9cqvXIwX2zfxk24E_uso78IpGptZ83ITlmLnK9itXSmq2rwz5t_e8XrzwCU_tgCMN8R1DZQ=s128',1),(3,'Independiente','https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEimsJgFxyJXwpXGV_uPf9mBgwnx_Yp8UDgyOxZGLgUK0HoisEqPNa6rZfFHKEnLKlxII9hFeSRsnQ077Iz9nWR3LMZEvQbEWDl3L_LlpkIpN_X4_GW3vHFtxYNYz0jUmafXGyg1tbvN_0Iokh9SydpqgD72ByRPRCHz01ZgON769GQG6npXUmgIB-heLBw/s128/Club%20Atletico%20Independiente128x.png',1),(4,'Racing Club','https://1.bp.blogspot.com/-sgtqeOWWS_I/WWZl7uA4IGI/AAAAAAABMuU/9tk8cPGhac4goZe4cuthJxatrtQw2UzcgCLcBGAs/s1600/Racing%2BClub128x.png',1),(5,'San Lorenzo','https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEiLatQJvwM5qEnrQVL1m1X3lbo_bUkqohF06hxwJrppnnuQmA2DtI_n4QPFH2_iL_SgC0w9ttAobkbg0YYZNjLLEbcNJe4pv945mTq4RQ36M7o6Hdk6rAdawrOCx7H8ON4vyhpqZQM5vu5t-cjYg_3sp7JwsB3EY8AVEoTtA9AU8xb2w3UlTTEuoaeJSzI/s128/CA%20San%20Lorenzo%20de%20Almagro128x.png',1),(6,'Huracan','https://1.bp.blogspot.com/-kwppBEIz7tg/WWZqMtnyqLI/AAAAAAABMvs/cLMEN_Fn2H4WL-u-Y5PMD6BKvSrEVhnjQCLcBGAs/s1600/Club%2BAtletico%2BHuracan128x.png',1),(7,'Belgrano','https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhkk2WH1z7RF1RE00XXNyolmZZnlf05aYqnCtJBBO19yWDAD-7ATTD5mCVpOSlIzKgUg0lDJWiTOS6txwLrd8N4hlyiZIbajuyTds9yzbuBorkt1ccwr0AwUDSaKsB6xB4v5nHSn19bkihi4uEKO83UFiBr5xbNao89-Xh4T2T1PN3MILM7HUluYgUcnZk/s128/Club%20Atletico%20Belgrano128x.png',1),(8,'Talleres','https://1.bp.blogspot.com/-7vtPiTFz4v4/YHIclt-ibfI/AAAAAAABjLA/9QuESkYG81EBZq8OL32rIMqAhBqj9EjEQCLcBGAsYHQ/s128/Talleres%2Bde%2BCordoba128x.png',1),(9,'Rosario Central','https://4.bp.blogspot.com/-_59zDJfWiBQ/XBhZYTlcXqI/AAAAAAABTlY/mugGpUq8FO0dE9udElzJGyponqkNhq6mQCLcBGAs/s1600/Club%2BAtletico%2BRosario%2BCentra128x.png',1),(10,'Newells Old Boys','https://3.bp.blogspot.com/-u5Qd_cGRnJ8/WWZoSf15x3I/AAAAAAABMu8/Vm0fYXwXdyQAfP69Gv6h2WmK0jq2vUbpACLcBGAs/s1600/Club%2BAtletico%2BNewells%2BOld%2BBoys128x.png',1),(11,'Banfield','https://blogger.googleusercontent.com/img/a/AVvXsEiZEUKPDYSMC8jq9i_h2MW4sW5PAODO1tsTY46lGoIuQMphpA0kLQmmVb_DfE85Fgd6rwct9FF_2LYz7T4qjj4RXqAykFFLLgtwHgDEKyQ7IlSYf8IPh3dCXXITgK_2JMhqaautnTRIoLGaBXB4dxQzpqaciKgnRc51bQBqGetAvooSZiRgBM0_cAH9=s128',1),(12,'Lanus','https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEgY0LCvqZHXem9kWtHYUM9sgTasj7XkYDtm9xGPZcXBWdWzAC9MDJPd9T06FbroMSZjaepTDjao038Vz7grV41lvC0-kY5FCv5YdlVLaDDxRv8R-9-OZ4VZBcloZbV9vRxfyBU_LJF8BYx2FpTOsIw_vKaWAMOYwd1OM3dfw-8VlCF9i2tH-kD-r85UsQ4/s128/CA%20Lanus128x.png',1),(13,'Velez Sarfield','https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEirp6aRrGW9ehZvYes7b_Lsz1nHmke4k3nV_Knh_zcmgVpEbFhqWiN2-tJGGOcNU00R2uudedXBlqZBIX0ldnSOw5W6JNB82FR5Kr8qEjseU6xRTPxRE5gsMRuxjLzX1eN84kwHKugkoZYqLCISFFubWIbCNtPZMIYTXdqU6RgLOh9UcQQj-MOJspTL/s128/Velez%20Sarsfield128x.png',1),(14,'Argentinos Juniors','https://2.bp.blogspot.com/-TvMcUIJ76sE/WWZwOvZeuTI/AAAAAAABMxs/8fXQZXEd2vg_E3KxYHoPTQe0ER8jHeFSACLcBGAs/s1600/Argentinos%2BJuniors128x.png',1),(15,'Tigre','https://blogger.googleusercontent.com/img/a/AVvXsEheE2-jT6jhVGlGIRdZPJwS_A7uR_WMCHIQvTd1mUCqJNOau42BT_qO4tWzf7lRE5DJZqKhwo1JDkyQsn5ayntCNsHeCGVFsylbFmRNe8vhjhmHfgdMvJImq74xDJ6fIuvcxTTVg2YDsDzh37D568XL8A1AkCZZm2BkKqwtrA6ezoAVtEKqxEtFHoYs=s128',1),(16,'Defensa y Justicia','https://1.bp.blogspot.com/-CjjTm8k09Ps/YHkUNQFmXuI/AAAAAAABjNM/IW5h2otbuAIA8tg-MsY4W5nrHY8FDC7dACLcBGAsYHQ/s128/Club%2BSocial%2By%2BDeportivo%2BDefensa%2By%2BJusticia128x.png',1),(17,'Instituto','https://3.bp.blogspot.com/-E98lHCoVdLE/WWcRpFeHSjI/AAAAAAABM0U/pg7wy5V9M2QmGZtCn7Ivr42mWkVnoDVbwCLcBGAs/s1600/Instituto%2BAtletico%2BCentral%2BCordoba128x.png',1),(18,'Estudiantes','https://1.bp.blogspot.com/-ZPdWhQlzMP4/Xb446uVZzDI/AAAAAAABWhM/i7xVLJ7qgpUprEy_EyvWtr2KuFhz7A9UwCLcBGAsYHQ/s1600/Club%2BEstudiantes%2Bde%2BLa%2BPlata128x.png',1),(19,'Gimnasia y Esgrima','https://3.bp.blogspot.com/-RvG0kTI9lfY/WWZrIViUlnI/AAAAAAABMwA/sIbTTXBUp9k8WrDnitoMkQWXXKLDQ4_3wCLcBGAs/s1600/Club%2Bde%2BGimnasia%2By%2BEsgrima%2BLa%2BPlata128x.png',1),(20,'Godoy Gruz','https://1.bp.blogspot.com/-DjwWX9izwgk/YUUtL0o92zI/AAAAAAABlaw/BmXxHPpUsPMrXBw63f5J0bfIdyxYO64zACLcBGAsYHQ/s128/Godoy%2BCruz%2BAntonio%2BTomba128x.png',1),(21,'Platense','https://blogger.googleusercontent.com/img/a/AVvXsEhWlS9WXXyR4Zd_AvlO9jODpFkgDw7w2z4GfjzQQ1h_na8Zma7je-7IBJtzjDmgDEL9qKkQyQ-Il96dBy4WsvtZci-DZvZ-_r7SmTf-erteqfMzU6tZc9fpnOKjWyWPFsHaNQt9mF8vJOP5jm3yY9zoBJxf72BclAiUZqr0CxGnHM79CWNtt5m-Nai6=s128',1),(22,'Colon','https://1.bp.blogspot.com/-K92873qufCM/YL0Ubj3z41I/AAAAAAABjrM/4_jrILJmnC4gC9iM8s7JMM_3yxZpxZaqQCLcBGAsYHQ/s128/Club%2BAtletico%2BColon128x.png',1),(23,'Union','https://blogger.googleusercontent.com/img/a/AVvXsEglXo8V_6pxTEHZ-c81Q37Y42rQ80ittTUvm265m9VL8D_vv8_D14wdxqjyDUHW9UtsTBIliG2_Im-BdrFyvmf7ZfIe_rso9ejGquWbD7W8jxDz31BO40uw3kKwJPnGxdMasorbQSPGkIFHgMFm-ImXrbZ14Q-5-rV5hpqhhWlJpDJ9NqHcmFwbrjYC=s128',1),(24,'Arsenal','https://3.bp.blogspot.com/-bKyAVKxc2l4/WWZvybJWzJI/AAAAAAABMxk/QY1zjQQNA8IZhFxjWXNN20pppSUhRQxMgCLcBGAs/s1600/Arsenal%2BFutbol%2BClub128x.png',1),(25,'Atletico Tucuman','https://1.bp.blogspot.com/-_HgCg5zRnZY/Wj005GlFZfI/AAAAAAABQmc/rsAjfBACiec64qQu13qTC-8zFppDQqymwCLcBGAs/s1600/Club%2BAtletico%2BTucuman128x.png',1),(26,'Almirante Brown','https://4.bp.blogspot.com/-CsDyYgpUvAA/W-uuWEcJRLI/AAAAAAABSr0/jgFj3w9BMS0HAAw8EreYUk29aRoTLqeTwCLcBGAs/s1600/Club%2BAlmirante%2BBrown128x.png',1),(27,'Atletico Rafaela','https://4.bp.blogspot.com/-KAIeDknX4lU/WXCfSJdfztI/AAAAAAABNe0/0VuWGyFA8w4JpMHObzYDM6-zdmxFpSIFgCLcBGAs/s1600/Atletico%2Bde%2BRafaela128x.png',1),(28,'Quilmes','https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEiU4tDo5BSy4w7raZp9zycKpa3bvG-IZf0uJWqgzNeNOx398I4FwgOzyyrGi6k1wPmLE6QdJrl20-evWySYw7ONy3xNPR2lpPvv4empkc7UOhNlyRYlI4rs5WW0Vnwkmjf5y5BVXXIbsOOYKZlAD55DmTvhm0sz5-MWoj30mnSwmGz_OZlN2-VvLIgJ/s128/Quilmes%20Atletico%20Club128x.png',1),(29,'Chacarita Juniors','https://2.bp.blogspot.com/-VuAXrAI_-LU/WWcPHhJUaaI/AAAAAAABMz8/eaGiW2S0nc8Mm-C_0dr4cKAK6bpK18qmQCLcBGAs/s1600/Club%2BAtletico%2BChacarita%2BJuniors128x.png',1),(30,'Aldosivi','https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEjYLTS9PrLZVE7UP-Uik6iz9jCN52PMfWsJKdWvPYuKzSXG-_Ti_vZlFQEE_8usH9wtg_y3By5Sc-Dq6h6KPDudFl-176QyNxk75GeOu6snFYmbHnKNGhHiQaBbJ1GqihS-WXANCOoadhwT1KnimCOLaoGMbmaKsBXSCdjar0-Y_fP3vg8A4_6jIvF5/s128/Club%20Atletico%20Aldosivi128x.png',1),(31,'Patronato','https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhHSVELQxaixKKf-HBJqJ2R5HecTYm_fXbRLrJ-w9Vtj1QVvwJUpYJ_iefWcvgVWSavRtp9LK-bGwGmqYBZXsBdPMz0vgohD1UmmvdRiWsKzvSYQuJGGkfdSc7w2YD8vKfESTyN-fwdgQlWp430iFugtNR5r9h3pmSiU4ivZnfOzNAi8JyHm7vyFN-l/s128/Club%20Atletico%20Patronato128x.png',1),(32,'Independiente Rivadavia','https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEj8F07AX4y-MKnEJckcHO1LrsGEnrZ6bQibq0Wn4b2i1ABb56vxpEXOOWTIlbZ9pPyLm762JKkcOPcSgprQLa20W-5fziBNoPUV4eOpIa3sv1777l1NfMuZHEfp7iBy5xHe2neOFVoOWiCV-5uPF31huyZuteDC17P9AHcGU_uK391hSZ_Y5ROufCue/s128/Independiente%20Rivadavia128x.png',1),(33,'San Martin','https://4.bp.blogspot.com/-RN6WHjk7fdE/WWZjjEcDCjI/AAAAAAABMtg/m2X_f-_MT2AjzL4eMPeJYUgFVQVJs3pMgCLcBGAs/s1600/Club%2BAtletico%2BSan%2BMartin128x.png',1),(34,'Ferro Carril Oeste','https://3.bp.blogspot.com/-3XcpH_DYDP8/WWcTJrqB-2I/AAAAAAABM0w/fUt9104cUPsy0n3pb1oEO3yKKFK0ASLwwCLcBGAs/s1600/Club%2BFerro%2BCarril%2BOeste128x.png',1),(35,'Vasco da Gama','https://1.bp.blogspot.com/-uoAjFVWgIoQ/YS7T6mzj_vI/AAAAAAABlNo/6B-GGDJH-EEmf6H5D1mEEnO7bP4oz5h9ACLcBGAsYHQ/s128/Vasco%2Bda%2BGama128x.png',2),(36,'Cruzeiro','https://1.bp.blogspot.com/-SifK8z74BnQ/YA4x_Y7WwYI/AAAAAAABiak/U86V0hhzyn4LgAEL3mvZmVs0atOa3ARugCLcBGAsYHQ/s128/Cruzeiro%2BEsporte%2BClube128x.png',2),(37,'Goias','https://1.bp.blogspot.com/-RzOkis3iZ_M/YB1Ug_RRqII/AAAAAAABifs/hb9kd7987p86dqLrIftBegsO2ed0MNevACLcBGAsYHQ/s128/Goias%2BEsporte%2BClube128x.png',2),(38,'Atletico Paranaense','https://1.bp.blogspot.com/-2vGwpP2TAv8/X59nAE4plbI/AAAAAAABhPA/fooXKkg-GAkUbmO6ed26jG0S9Pmw0U3ogCLcBGAsYHQ/s128/CAP128x.png',2),(39,'Fortaleza','https://1.bp.blogspot.com/-lzxiYCUpk00/XOsfJyWcg1I/AAAAAAABVII/PZY1rLTtHCIMwtxWXu1zRT_HpGjAu1AJQCLcBGAs/s1600/Fortaleza%2BEsporte%2BClube128x.png',2),(40,'Red Bull Bragantino','https://1.bp.blogspot.com/-j02zJGZFtIU/Xg_x78Ri3zI/AAAAAAABXAY/jRxxsQP8e1oxpIR8yah8TZX3oXNJHhItwCLcBGAsYHQ/s1600/Red%2BBull%2BBragantino128x.png',2),(41,'Internacional','https://3.bp.blogspot.com/-lt2iPeVo3lY/WWPxzSv9hOI/AAAAAAABMjk/jNywsNNBSCkG-ZZrHnz9VYYHNi-fveB0ACLcBGAs/s1600/Sport%2BClub%2BInternacional128x.png',2),(42,'Corinthians','https://3.bp.blogspot.com/-f5LBlsWhFmo/WWPQ-O-hLgI/AAAAAAABMfs/JZTBhNVSJk4-tmSsZsM3-TnO-UJiWpAlACLcBGAs/s1600/SC%2BCorinthians%2BPaulista128x.png',2),(43,'Flamengo','https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhus4SZGFTO0T8qOFRtBmhdp-CFdK_CO9JNlfOIsEQG_-Z2mqj4gRTGwGRqei1fq-YdZZOV_AzSz3RcdkwcWjPUc4pn2zkQAr1lytasg7vQdcWstg8fWQEhTeVuWM4evElEIy-BpufwK_oSxeu3x-W9zdpDqRTAO4mfYZ_RoKOpGCSI68Lt65CcEMPadV4/s128/Clube%20de%20Regatas%20do%20Flamengo128x.png',2),(44,'Gremio','https://3.bp.blogspot.com/-DVpv5SzK5pA/WWPPFSQUZBI/AAAAAAABMfY/Fl9PLhQRZNYTOpUxuK32Qvx04KdmPvoogCLcBGAs/s1600/Gremio%2BFoot%2BBall%2BPorto%2BAlegrense128x.png',2),(45,'Santos','https://3.bp.blogspot.com/-zaqtVpHz8Cs/WWPOp0ZsqGI/AAAAAAABMfI/-bNCC35NB7gy_9i7ec8t_BDfddUg6CUbACLcBGAs/s1600/Santos%2BFC128x.png',2),(46,'Palmeiras','https://4.bp.blogspot.com/-ULieMfLG1pQ/WWPNveP293I/AAAAAAABMew/UuT1aPBttGE7tdTDYTrKg6IS4tA701GugCLcBGAs/s1600/Sociedade%2BEsportiva%2BPalmeiras128x.png',2),(47,'Fluminense','https://4.bp.blogspot.com/-pti-GKhgQ8o/WWPNDhk1DjI/AAAAAAABMeY/NwyK4SyvIjMXAPQZCjwwiWVyTmTKOCPNQCLcBGAs/s1600/Fluminense%2BFootball%2BClub128x.png',2),(48,'Atletico Mineiro','https://2.bp.blogspot.com/-nIlIoqzP8Iw/WWPMuwkJQ6I/AAAAAAABMeM/9bSQN8vYFn4Yt5Fr1vK3l8nMTEuYAnfgQCLcBGAs/s1600/Clube%2BAtletico%2BMineiro128x.png',2),(49,'Botafogo','https://3.bp.blogspot.com/--wb1CxJiWg4/WWPL1O-G64I/AAAAAAABMd0/ABxiPfdXcxQDPL-c0gR4YnGsKXg0BdxfQCLcBGAs/s1600/Botafago128x.png',2),(50,'Bahia','https://3.bp.blogspot.com/-r_uvCIXOCGA/WWPI6Z2XWII/AAAAAAABMcw/s9w7ZnPaXe8ViJZCZAQmUleS4KIHyPyuACLcBGAs/s1600/Esporte%2BClube%2BBahia128x.png',2),(51,'Sao Paulo','https://1.bp.blogspot.com/-DSQYCUtUpZw/YAYU-c-F1mI/AAAAAAABiUk/tgsRzOmbZG0AYFTWitKXGbljqkaba3z9gCLcBGAsYHQ/s128/Sao%2BPaulo%2BFutebol%2BClube128x.png',2),(52,'Sport Recife','https://2.bp.blogspot.com/-mgb4Hve9STk/WWPKxDafLwI/AAAAAAABMdc/QLSEUZkFTCgCkOj2jNdp_rrzGDA3WgRYQCLcBGAs/s1600/Sport%2BClub%2Bdo%2BRecife128x.png',2),(53,'Chapecoense','https://1.bp.blogspot.com/-wxFZb_wenN8/YM6g91-mkaI/AAAAAAABjyU/JbzyxhudfP0SBDqmExjgfpr21AjloOZ2gCLcBGAsYHQ/s128/Associacao%2BChapecoense%2Bde%2BFutebol128x.png',2),(54,'Clube Vitoria','https://1.bp.blogspot.com/-TyhpYMCN1A4/YBtULVuEuEI/AAAAAAABieE/DBPyHTgjee4NKehSIMxHBMRQe7Yfp76zQCLcBGAsYHQ/s128/Esporte%2BClube%2BVitoria128x.png',2),(55,'Liverpool','https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhjh7anvMvcjhEsePGgye96kF9aWbNUkgOFjqqdiei3eYtBqvueofsnjeyNw5Y46TZRmCH7JXO4QXodQ6fsZJHKmQOp-XaKPlaA6kB5Pqzk75E63_sLwhuXWlEj1WDWNf9yogwaYNhkUeW9MiqFWzBtQ-692ujTNkHDKCNNYr5FPVD8J3i1OCjjYFcwopE/s128/Liverpool%20Futbol%20Club128x.png',3),(56,'Nacional','https://1.bp.blogspot.com/-vBP0-mz6Gnk/X4pdMXhNtQI/AAAAAAABg9o/YPAkNsfjTZ0rMObW7HNqspWDRtKlnRBDgCLcBGAsYHQ/s128/Club%2BNacional%2Bde%2BFootball128x.png',3),(57,'Montevideo City Torque','https://1.bp.blogspot.com/-t0nQsjbhGH8/X3qYbpKhgtI/AAAAAAABgvM/tA_TpGdemU0wtkOdlVTSPzGUeSW0UuyUACLcBGAsYHQ/s128/Montevideo%2BCity%2BTorque128x.png',3),(58,'River Plate','https://3.bp.blogspot.com/-2jq0uAUPamk/WV6m4J9Ao0I/AAAAAAABL8w/79ujYUulOC0R17JZ_5laebwRgIMuaeTQgCLcBGAs/s1600/Club%2BAtletico%2BRiver%2BPlate128x.png',3),(59,'Danubio','https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEgtwqJdxTLjJ_Xrp5CbbP6ZOnBbpC6gMpBOkhQJWG3b57SrWgJ-BQ8PsMKQx8HISRK2OuRcYV_9pX5neVI4rnzhWu3hTM_ci19oAQPbpTG6GrE7uZciSrF_9G6RnvDpPhTm_IOdNRl780tWycUkNkjk-sPkBBDEknNWjKINZTxGNOL8uZzsZOStvY-8gS8/s128/Danubio%20FC128x.png',3),(60,'Cerro','https://2.bp.blogspot.com/-wfhdsrxcBb8/W0wqt-sZ6cI/AAAAAAABRm4/BbEDoitFnocp51hZucS4ri9SNroMdG2vgCLcBGAs/s1600/Club%2BAtletico%2BCerro128x.png',3),(61,'Montevideo Wanderers','https://4.bp.blogspot.com/-jKPOormJnYE/WV6kkZe8l5I/AAAAAAABL7k/pMtP41ZURpMROHqQ_oC36DmKpPZvA2LuACLcBGAs/s1600/Montevideo%2BWanderers%2BFC128x.png',3),(62,'Peniarol','https://1.bp.blogspot.com/-SIbI219OWyo/XofCAlz8JAI/AAAAAAABYsM/RWskzscCyOEIxOS2ZeEydAvwRk9sGV5BgCLcBGAsYHQ/s1600/Club%2BAtletico%2BPenarol128x.png',3),(63,'Defensor Sporting','https://3.bp.blogspot.com/-y9fS0KyxL14/WV6jNVnkZHI/AAAAAAABL7E/jr7YiHnOkII6Fo9oJTxZQcWMb1si7j2yQCLcBGAs/s1600/Defensor%2BSporting%2BClub128x.png',3),(64,'Fenix','https://1.bp.blogspot.com/-76uuJpGKybc/YaLaT45aUII/AAAAAAABlxk/Gaq6Uj695KsbBjl1d7rBb8NxNeLv36eSQCLcBGAsYHQ/s128/Centro%2BAtletico%2BFenix128x.png',3),(65,'Junior','https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEiSgBb1dgB6FCbvK_JceuzzydjieWBPNNcLN6EYZB415vYsHGDn7DXeZAZhitPbFt3Kdds_F5IR48ErhiYN6fxAhYgJkort-f0zJ1AXi53zaDqaQCp9dLMv8mxJCTsly897nZxM__YXdX7xzJ7nRiyal-YalXQAMjJXn2Ts1W9di3pFeQOd8fRvSiHDXn8/s128/Atletico%20Junior128x.png',4),(66,'Deportes Tolima','https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEjycejWZfgHcQDzYUrI6MPJUyXFBBREYtQQvbOkqBPoh_oGpioJHGGvt6Cphyucyi_6fiLiS5LMBXd7TUH5Z8mrRxPDP3_HOCQFTMoS3HdMvby5RfFGmaUU3Ky7BMoef8Dk0vc0mbFghVr-vpTdo99sa4SKTHlIfRXN-VhVclt-XfoAjsAv0yHXNEjRbiE/s128/Deportes%20Tolima128x.png',4),(67,'Deportivo Pereira','https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEiYpNxp8kLRr9AQctRihkvnZ5rP0ABVH3UMzMtB3jjrdefDX4kHik0B7H-nAv7NojeEzhA-0dU-ZIlCj7FYy66rJVOFkOx9V6LI_GsWxZx-JFTGmD8B5Ceu3wgNyKktSPX-tzMbz6i7rdRQvqjErqMMUlxm_KuglXYRk-Dg3UJyhbCOANdFmG8c70cT/s128/Deportivo%20Pereira128x.png',4),(68,'Independiente Santa Fe','https://4.bp.blogspot.com/-VSObqRQK5xA/XEoRCIpgBTI/AAAAAAABT8U/N5goHXmUI3cBUP2jpPO452NAYfyu8iOpgCLcBGAs/s1600/Independiente%2BSanta%2BFe128x.png',4),(69,'Aguilas Doradas','https://3.bp.blogspot.com/-bLhwTfl4_rM/WWHaYfzrdoI/AAAAAAABMMk/tah9-s42O4sjwey49ivgJ22AjMe0d7tGwCLcBGAs/s1600/Aguilas%2BDoradas128x.png',4),(70,'Once Caldas','https://4.bp.blogspot.com/-fB7XvTwvoxI/WWHZ1eVR8MI/AAAAAAABMMY/Xr_k20l1Xsc7Pwykq_0A0RM-qFu6TPXpQCLcBGAs/s1600/Once%2BCaldas128x.png',4),(71,'Atletico Nacional','https://2.bp.blogspot.com/-Oe-HGsFoBWM/Wj0wrCn6Z5I/AAAAAAABQls/HEKSae-8B5Yyea9XYZ7F0hcsJgwVysyyACLcBGAs/s1600/Club%2BAtletico%2BNacional128x.png',4),(72,'America de Cali','https://2.bp.blogspot.com/-aYwj2MWphHA/XDjJjbnmvII/AAAAAAABTtw/Bj1XZudsvGQUD5MSYCUXcfSivXdOltdiACLcBGAs/s1600/America%2Bde%2BCali128x.png',4),(73,'Independiente Medellin','https://2.bp.blogspot.com/-H_ZuUnNXCus/XGBcDVpKtQI/AAAAAAABUdg/YG6ku27kJNgZAtRJBYQ15J48e14Ah3-QgCLcBGAs/s1600/Deportivo%2BIndependiente%2BMedellin128x.png',4),(74,'Millonarios','https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEjVsAJ0zqMVb6ubdqKbbTB29zCjlXiIgujZV-WBl5CS8TZ-gkwf-l_9ibVmaySedh3mYY3NvhMYqwZ3LE6zAtC4zIh2E0sLGTI9x1oqfiTcsK8LqUkpsW1yRMQYxa4cXL6TdyxrITX16xo6ICo5S_1tBPwO-j2ioNwU9s-sMSzCr3LJsGPutAi4fiNSprM/s128/Millonarios%20FC128x.png',4),(75,'Deportivo Cali','https://3.bp.blogspot.com/-xbLh7umZF_0/WWHSQP-n8DI/AAAAAAABMJo/28t1eSE2gVoJqjctbPBUtpByz3ZqwEdHACLcBGAs/s1600/Deportivo%2BCali128x.png',4),(76,'Cucuta','https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEg99LpzQBx2GhTYw6tej1Kcbbon61NKdo_moYN5-MdrDgNEDxpahJG13_BsoMXFFFmlYijZONojgO0Y3OBQDvJYUAYvzgQN73UHMOa7hpbvM9BULterTODxmBNQbDsatz_0y-DauMEnSGEOHsxce1g59k5ejwldfx6gA57AMSxmBOSM_zZP14zr8-Ff/s128/Cucuta%20Deportivo128x.png',4),(77,'Colo-Colo','https://1.bp.blogspot.com/-hTC_Pvp58tk/X-e676BOQ8I/AAAAAAABhjU/z__Ds_JTFgAcU93Ga68bOIoIo6973A6KwCLcBGAsYHQ/s128/Colo%2BColo128x.png',5),(78,'Everton','https://blogger.googleusercontent.com/img/a/AVvXsEgpqwrORzvfTMt4T08tv9r3Sj-sw5oJi0BCt0piLvwzGpPkNc90U7OeuLdliK4X5enaVs0JATOXsB3IVWsUVvWHHO9o4Ivtr0eSzh9nA0TX4siEtzY8NKKThKjmRvdfOEomHZX2dGrcyTlVCoOswf3HkW01xx_poIKRfy8T2O3P-ep3MknQIssfvpt3=s128',5),(79,'O Higgins','https://3.bp.blogspot.com/-jhV3lopuGWM/WWLznX5Zk-I/AAAAAAABMTw/NRByr-GvoGceWWgstMRRy0HUs1AIKWWXwCLcBGAs/s1600/Club%2BDeportivo%2BO%2BHiggins128x.png',5),(80,'Unión Espaniola','https://1.bp.blogspot.com/-Owx7Nv4owyo/WWL2mdxuDuI/AAAAAAABMUo/wlaho5EfX1sMCq50Kr2x2U_FqMOahVWrwCLcBGAs/s1600/Union%2BEspanola128x.png',5),(81,'Deportes Cobresal','https://blogger.googleusercontent.com/img/a/AVvXsEhA9wChrmK4jFTJsbCk05c7U3BEi8ot8BlerCDOhCngg5nqJvwbxvor6jycydow8j83w0G3f__al3HA5QT1dYJFOLE-ppG3VMQY77pKMa36Uhy6DSbsIZnoDo46KnnAV7xvdM3saPC8Jsmj5Tq1tdB0SWBJHFL8DkmrX49MCl5WCO40N0383TxMWO8c=s128',5),(82,'Universidad Catolica','https://3.bp.blogspot.com/-t8FDiVsqsQc/WWL3KwAMbTI/AAAAAAABMU0/LoNMylJ8rkA1WoIZp6EOrkJd-j6z3ZgFwCLcBGAs/s1600/Universidad%2BCatolica128x.png',5),(83,'Union La Calera','https://2.bp.blogspot.com/-wgBnZzQhD9I/XFDDRyHmsdI/AAAAAAABUPE/9ZM5kZgingEwXa9kwtE05YRsB_QlxoLOQCLcBGAs/s1600/Deportes%2BUnion%2BLa%2BCalera128x.png',5),(84,'Universidad de Chile','https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEiV_dv2GFC8428-M3nIF7b25FD7FbCLCxMTjHiO4cnL_U0uuuA5yT8Ae49xEazEhhKdWPFZPRR_Pq9QLJUua45DSQ9XBPaZR_0lNhfMR8gZXiOMG2giGPoVQCHKykKAzXUPfS-PHzPhKUUKZZNpwfJNVGYvkGvCa3cWLGCoyOBqrPpEzHDtpjN9mSkE/s128/UdeChile128x.png',5),(85,'Palestino','https://blogger.googleusercontent.com/img/a/AVvXsEj6uMd0BCWJJJnODgb9lO4aIsf7nYAKAWy8ahYMJX_pYgk2xSo9iArPq8UGcYgPds2QUgH635BeUWYp4ubqk7O2vHRgQjkmOjvFBc_4-uIyjfWtDgPmyVEdIwTEggi1ePkmB46KKE56SoPf0J7djmMaVfIFZLuOQYQkl9jIA5XJWXqLwPMb9yGR_zf8=s128',5),(86,'Audax Italiano','https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhdtbbd3hDYJ5buong3Fe321weeBeuawUag4_EbOSqQopSgWX-mOqpLcEwUAVUPz_jfJu7woys3QZagdzFbv5tuBK9g-FETUmEIFVwXUalHsrDEBeOFhWfMb5WsWeDtTLsXtoRCXEYhmBMoN989pTeb3R92c8jYp1JAydjobZM2EXowp-vbYZhm-P-2/s128/Audax%20Italiano128x.png',5),(87,'Huachipato','https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEiWhlkPOp7ZZfDQ9QJpqSE_rAeiCOonfVUPer4Mis7EpkJn97dohp2t6u7wYWvfdEDNFB8Oevg-NId9ToQabt40JyvrtTB8csmLi-ggHLkCxJ3seAjg53QCNYcyCL_SDYZkQeYtHh8com_wGg0_AZFJHfPe6q_CG9X6g2PQxXOnhC4tPjTrOpEUx-ciePo/s128/CD%20Huachipato128x.png',5),(88,'Universidad de Concepcion','https://1.bp.blogspot.com/-BPyWp3DB88M/XPnGlGsF5oI/AAAAAAABVO8/Iep1drDhj34vD2g1tr5Ud07k38sSKwRgQCLcBGAs/s1600/CD%2BUniversidad%2Bde%2BConcepcion128x.png',5),(89,'Santiago Wanderers','https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEiNcSpD4cmoeTSm4vLYQKkgXjbbRUbMUG__ZzMzlg-aRbYaHDLWdm5l7pWTeCo9nJ8Eyt2y4QAAHP1fRnAGiFppBczrilIQZa1RGKjcjCCyPI_HJkBMl-nu0i6XcN77Mvd6y8HBOYhtY5Ru6KOwJ56J7hhawx-s5COr9T7T5JlQJjDQaYA1gP5Atf8u/s128/CD%20Santiago%20Wanderers128x.png',5),(90,'Deportes Iquique','https://3.bp.blogspot.com/-CD1J3xVLz6o/WWLt405dmeI/AAAAAAABMS0/HOUT7dv3dWITrBAmKVdwmr4-NK9uhA7VwCLcBGAs/s1600/Club%2Bde%2BDeportes%2BIquique128x.png',5),(91,'Deportes Antofagasta','https://1.bp.blogspot.com/-_daFgD-_Tb0/XPW__SxCklI/AAAAAAABVM8/X7cDIbPlVn4NuT0eUZpX_EdWQJiXYwqZwCLcBGAs/s1600/Club%2BDeportes%2BAntofagasta128x.png',5),(92,'Sport Huancayo','https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEj6tg9e-SJ5tMjB_MP4YjZUSHDRmcwm1b1Z5gmxwsCvp167M5gn0G4SQsmNs6WtFRIAjSkOswBAe2Z00PJ0Ea793XfsmMhT7Pxtz91kQjVGJSE8UZZZjOjGO7HRNmsY_Svvud9J7Qsgzo0inlQ1B0XPBEK64E-fBUj-FsyQMVnpJPqfMjsN28bL_xvR/s128/Sport%20Huancayo128x.png',6),(93,'Cusco','https://1.bp.blogspot.com/-kjUU_SvfI78/X3qTeyFLJeI/AAAAAAABgu0/FG5FVVJKUAILOOGUvN68P6tSUlI81wXiACLcBGAsYHQ/s128/Cusco%2BFC128x.png',6),(94,'Sportivo Cienciano','https://3.bp.blogspot.com/-Fnu9K9Hdock/WZMiOndPe2I/AAAAAAABN-k/nOrg8Aov6tcxxn5PfQy1ShRyz-hJ7Rl4ACLcBGAs/s1600/Club%2BSportivo%2BCienciano%2B128x.png',6),(95,'Sport Boys','https://3.bp.blogspot.com/-UZXA8pTTZPg/WV7iiNNjz9I/AAAAAAABMBc/Ez-VvH2FzkED8AJCDfb2nDHuHevFYzMnwCLcBGAs/s1600/Club%2BSport%2BBoys%2BAssociation128x.png',6),(96,'Melgar','https://3.bp.blogspot.com/-mgPMjaEFyf4/WV65plLq8-I/AAAAAAABMAI/1M-1ONxqBDcWsP2u6oi3Md6gwfxkgXnZgCLcBGAs/s1600/FBC%2BMelgar128x.png',6),(97,'Universitario','https://1.bp.blogspot.com/-vrMr_9IIzEM/WV65Se9t6xI/AAAAAAABL_4/yl-Z2hoRH6QGTKryr1GgdsInVdr3sotKwCLcBGAs/s1600/Club%2BUniversitario%2Bde%2BDeportes128x.png',6),(98,'Sporting Cristal','https://4.bp.blogspot.com/-oMUwo_7ShVA/WV64iigzlBI/AAAAAAABL_g/C1tyCLX_oOQ2TRB-wu3kBkjsi6uhb1JGwCLcBGAs/s1600/Club%2BSporting%2BCristal128x.png',6),(99,'Alianza Lima','https://4.bp.blogspot.com/-A0Jfoe8Xe5s/Wj07KMKS_II/AAAAAAABQm4/O9KJ_EAsttcZjp1FeThZDpWKpw9lGNX1ACLcBGAs/s1600/Alianza%2BLima128x.png',6),(100,'Guarani','https://1.bp.blogspot.com/-W-VskmMPkLg/X58hhzOsH0I/AAAAAAABhNc/f6YT1w0ybFIBXSC2AdK1rOjUFVk-WnLKACLcBGAsYHQ/s128/Club%2BGuarani128x.png',7),(101,'Libertad','https://3.bp.blogspot.com/-ZlfhqUIfMkE/WV_4XxueyrI/AAAAAAABMCg/cT_nfizLs7wd4UJpf9RHA5qyH9CqZ6vDwCLcBGAs/s1600/Club%2BLibertad128x.png',7),(102,'Olimpia','https://4.bp.blogspot.com/-H9HkPfB9e18/WV_5X2Bny-I/AAAAAAABMC4/xmliFMoSqx44Zn5XFsxN2X86V638EnP8QCLcBGAs/s1600/Club%2BOlimpia128x.png',7),(103,'Cerro Portenio','https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEjV__vpUrc07EEQBD_wDz4RSzPzP_EtGEgfDdTxTq2MVIEhW8-jbFbxDHv4IxLvcMh7s6IShpFlVIQUPYIabo8uvLHkzl7JJbdfeyx6g9PWrxRPJW7BdzMIli8GKe9iCSGoNwvtWoe9jX63NeeBN2WGFJGOTtGYqZRBLxGQ7PeFTv85GVw_wRzT1vcs/s128/Club%20Cerro%20Porteno128x.png',7),(104,'Nacional','https://1.bp.blogspot.com/-DsMsrMD7Y5o/WV_43aS289I/AAAAAAABMCs/AM-Ot6OPw2UYNjiCvg7kCfclK1u0J72KQCLcBGAs/s1600/Club%2BNacional128x.png',7),(105,'LDU de Quito','https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEiDIsPLRLyl-SLzxRKRR782JpE6_3wapeZVMV8il4PrNb2QuW9fDe-qzU-wrJgheKRFnuGnzUResybcTlGK8QbtINmRT4aBmAmmOl3oChEg4dSNBhdIrlgck8xsXIiFiQ0YJmjUApelf9whlWb6hAZ8ufg7mAPNexpmwSebbcKufJZbhT8MN6ehJGJdTew/s128/LDU%20de%20Quito128x.png',8),(106,'Aucas','https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEh6cEKFG4n03ISn-M8rdgfXLwl2FkNCPlr3iMcqbcmz8pM6Vuji8mp4Gf018LpuRumUSYeMVIbdB8hdDRH86Ipmff3kAlb3-EdjjcrPm_3H3QZV-CxYoCOmlTNiguGKH3SCsW_1cdIkY3dTU8VhhPZbP26AJhGhAcQW06Bx3ktja4uWN19gqZu_nDre/s128/Sociedad%20Deportiva%20Aucas128x.png',8),(107,'Independiente del Valle','https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEjtp9ZnoN9alyN3r-uCMq10bZ94qGJs4KHuPg48QTYKJ2eZzXt_SKBLJs10CGMQMUOpXQOVHxCgUWgzfWZykddWhDtjdNLybpb4G1zAa-1X4obEIwRmBaPVN_fJ3EQL8Wec7h2034llnDd1wX7gyJcghHiYY5WDUg6KfKQKd3NuKFRwtCpfa8Kddvo6/s128/Independiente%20del%20Valle128x.png',8),(108,'Universidad Catolica','https://2.bp.blogspot.com/-fdMZcwOZXhc/WWDiqadzPqI/AAAAAAABMGA/v3exhgTLZIYl1d6gjfPfAgeB2pJ_jJbdACLcBGAs/s1600/Universidad%2BCatolica128x.png',8),(109,'Emelec','https://1.bp.blogspot.com/-vwYiWieyuj4/YOtjF8RWWDI/AAAAAAABkeg/F1JMvw5fb5wPhwVgZS_CYIoS_jKE6GRSgCLcBGAsYHQ/s128/Club%2BSport%2BEmelec128x.png',8),(110,'El Nacional','https://4.bp.blogspot.com/-rX8l-r5VgB0/WWAFSB3ckWI/AAAAAAABMFU/B0piplfMOR4uTvUgCJUDf8JGaxc8KGnugCLcBGAs/s1600/Club%2BDeportivo%2BEl%2BNacional128x.png',8),(111,'Delfin Sporting','https://1.bp.blogspot.com/-VNTnF0rq4hM/YHyWD6irOxI/AAAAAAABjPo/sOM09KGy6hkLFx5e6bnYQT2Ps14lfat7wCLcBGAsYHQ/s128/Delfin%2BSporting%2BClub128x.png',8),(112,'Barcelona','https://1.bp.blogspot.com/-nWrMPZ7Yu-Y/YHx_HdprlII/AAAAAAABjPU/aJVRR1L6cC0Coey0J7Jv5bPORURbKVL5QCLcBGAsYHQ/s128/Barcelona%2BSporting%2BClub128x.png',8),(113,'Royal Pari','https://1.bp.blogspot.com/-mKZixmROTYM/YTA5j6G7BmI/AAAAAAABlOo/bCFHVj0GJ-A28TRkhC7p850TwiZ5Vu7MACLcBGAsYHQ/s128/Club%2BRoyal%2BPari128x.png',9),(114,'Always Ready','https://1.bp.blogspot.com/-PacpYQUBY6Y/X_NZ-8WRP-I/AAAAAAABh0I/gD3A97AWHZYVrxX64a5M-R1bjCiflJN_QCLcBGAsYHQ/s128/Club%2BAlways%2BReady128x.png',9),(115,'Guabira','https://3.bp.blogspot.com/-p4Gg4gkeXTg/WafcZoRM8-I/AAAAAAABOtY/glX3x89gr888z85cKbeCNr-alwFoznYdwCLcBGAs/s1600/Club%2BDeportivo%2BGuabira128x.PNG',9),(116,'Nacional Potosi','https://3.bp.blogspot.com/-k7MTn_5e58k/WaanUP6LMlI/AAAAAAABOrs/Gbm9OIzSIh4tnphBl6hmQRD50rY5UM-4QCLcBGAs/s1600/Club%2BAtletico%2BNacional%2BPotosi128x.png',9),(117,'Oriente Petrolero','https://3.bp.blogspot.com/-G3_8_3p3pnc/WXPQRD4LQII/AAAAAAABNkE/yGc2lz86tBIpTdyK3D1XcBD8uMQE3WbtgCEwYBhgL/s1600/CD%2BOriente%2BPetrolero128x.png',9),(118,'The Strongest','https://4.bp.blogspot.com/-g7WWDjxhJUs/WWEb3X8HYtI/AAAAAAABMIs/rO6mvvwxKeArDho5wlm5AruE_L2gpkMwgCLcBGAs/s1600/The%2BStrongest128x.png',9),(119,'Jorge Wilstermann','https://3.bp.blogspot.com/-6dFdtBkEMYk/WWEa87BgY6I/AAAAAAABMII/GK4_W8g2f5MotxKSClz7LQgSGe6i50yUwCLcBGAs/s1600/Club%2BJorge%2BWilstermann128x.png',9),(120,'Bolivar','https://1.bp.blogspot.com/-IlkXimignV0/WWEapWJ6o2I/AAAAAAABMH8/QZNlyybb4XchbH8NWMAn1S20WftVxQZmQCLcBGAs/s1600/Club%2BBolivar128x.png',9),(121,'Blooming','https://1.bp.blogspot.com/-NpvR8UEHxYM/X3oxcpvOqZI/AAAAAAABgps/x0KCKKmNxtAPHxrgjFTq86gge3JLlES8QCLcBGAsYHQ/s128/Club%2BBlooming128x.png',9),(122,'Independiente Petrolero','https://blogger.googleusercontent.com/img/a/AVvXsEh7vmO7s9cIQgAHOPQhrqJhvAkZEE8ICNDR6NmWF37UU4opThQ4jUDL3hEoSfXfneZypCeTEV7gMhVTVot4VToQN6Oqlsgcgd3F5D5RyhiQsidoSh4DwrcZhlPpGSzaFDD4dkTJKLaOqWpzgQidSWKeXLG2gDHfEPnV0pkLwgZZpujfRoNIwCVmmS7i=s128',9),(123,'Caracas','https://3.bp.blogspot.com/-95vvX2IawMw/WV6RFqm3kVI/AAAAAAABL3A/nzPoBQcN5u0xz9lUFYaS8XxdP5Iako7CQCLcBGAs/s1600/Caracas%2BFC128x.png',10),(124,'Zamora','https://1.bp.blogspot.com/-4pKMBaYqDAI/XBhbdNcYJHI/AAAAAAABTls/3S_553cNG74SyY6yBbcZxFTnYT5EGyNZgCLcBGAs/s1600/Zamora%2BFC128x.png',10),(125,'Deportivo Tachira','https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhG12LXNTKt0VK4aXzV93vsbKvgSuVUJRkYGpwvtk3OxmCKOWvsopF4APLL6OXCVDFT2A2dBSSEquGpbCGfCgYdhZ-oQpdfRSiaDfXMhPAwWREnLOksNqJalXpodFvL48-G0xeaKaG6qiIQyF5aYeI0sYhVGW0FzMs40JxlXoShXCBAcabW7mrUEkoND64/s128/Deportivo%20Tachira%20FC128x.png',10),(126,'Monagas','https://4.bp.blogspot.com/-dljSHyVRnxM/XHXNZl94iXI/AAAAAAABUiY/5wdHMARnl1oyH0MMh_Y8-oktADy00hdNQCLcBGAs/s1600/Monagas%2BSport%2BClub128x.png',10),(127,'Estudiantes de Merida','https://3.bp.blogspot.com/-sJ1Sj3RRMEA/WV6WjjWmoQI/AAAAAAABL4Q/Zv25cp9sbl83JEggxMHNQwJiM1SbZETwwCLcBGAs/s1600/Estudiantes%2Bde%2BMerida%2BFC128x.png',10),(128,'Mineros','https://1.bp.blogspot.com/-YBCfNT4kxHA/XZDpM8a6_9I/AAAAAAABWH0/UfhQbD-Ap2w4JPYnNvQ2s2n6hRAJW4a4ACLcBGAsYHQ/s1600/AC%2BMineros%2Bde%2BGuayana128x.png',10),(129,'Deportivo La Guaira','https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEg1vvjCas8DULJ79nzFF_GJJuX2XLrWEa0Izmzva4-1mVzfyLl2FC9sTDei1J9OJteI3o9P-P7Jer17TQPm-Dq6c95upPLwQHDFfCn8rjaMoEqzCYFIGewVGHidi8qWyD4Xp39x_g9zq995kBT0wqSLetSonw2lC1F_A-6OiZs2wB4L-rKGJQ7hwxKf/s128/Deportivo%20La%20Guaira128x.png',10),(130,'Carabobo','https://1.bp.blogspot.com/-JnX461IV_mM/XvP0OrTYJGI/AAAAAAABdu4/4pUl7qjIqYQa57ilvZILzUS5mxKMF0TFgCK4BGAsYHg/s128/Carabobo%2BFC128x.png',10);
/*!40000 ALTER TABLE `equipos` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `fases`
--

DROP TABLE IF EXISTS `fases`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `fases` (
  `id` int NOT NULL AUTO_INCREMENT,
  `nombre` varchar(350) NOT NULL,
  `cantidad_equipos` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `fases`
--

LOCK TABLES `fases` WRITE;
/*!40000 ALTER TABLE `fases` DISABLE KEYS */;
INSERT INTO `fases` VALUES (1,'Final',2),(2,'Semifinal',4),(3,'Cuartos de Final',8),(4,'Octavos de Final',16);
/*!40000 ALTER TABLE `fases` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pais`
--

DROP TABLE IF EXISTS `pais`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `pais` (
  `id` int NOT NULL AUTO_INCREMENT,
  `nombre` varchar(350) NOT NULL,
  `bandera` varchar(350) NOT NULL,
  `id_confederacion` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_confederacion_idx` (`id_confederacion`),
  CONSTRAINT `fk_confederacion` FOREIGN KEY (`id_confederacion`) REFERENCES `confederacions` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pais`
--

LOCK TABLES `pais` WRITE;
/*!40000 ALTER TABLE `pais` DISABLE KEYS */;
INSERT INTO `pais` VALUES (1,'Argentina','https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEi5d6sroo7YdfGbudsfcETZfyFpGlwgdLJN1Ah2RotgloDZm5pKyPFk1zBdVezHibLErTo568nz1iOmVHSBaPsoubNbxQJE0z8_mk1ggV9EqQ2LvM8C3z0RoXp6WVfu3N-esDEV6ynYuApYWMadrblKMEqwqkcARyb2m1tfMOrhhPZfPq9ywCav40CJ/s128/Argentina128x.png',1),(2,'Brasil','https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEgPG0fZx3bAFTlovrog0fKZmXM2pGfxDXOFt0mypa2TpzOGIoV-oM2lLEXmqV6EwVwwqVGF_F61BZVmEzNKFQzqbuFtcPjb0eY7zqFr20y90_pQ5gq_Cr-Yx7k2OjgFkbVnbtXY7N-dCprSiIQ04WO2EKxt7hr2RD_fMfnCnoUve1S-gPG6uJUzb8ZB/s128/BRA128x.png',1),(3,'Uruguay','https://1.bp.blogspot.com/-hjj34VmK_fQ/X4pYPI08NyI/AAAAAAABg9Q/FNp3KUuj0uEP9CBDI5LsTfhQ-O2DY1QvwCLcBGAsYHQ/s128/Uruguay128x.png',1),(4,'Colombia','https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEjQEaVh5Ibg4rNsc9a47-VOt8U4hXPIs19dIVVIe0z0qvfW91Y-AHHAqt5Goo4yJZfQEQqzEb6E-pdmJryUZ0J16vmh0L-pp1_8uBVrOTDwdI5QytoIL5KyxDGyWs9HPlsKpZq8Til2eLdfbYWC8Nk5kJpK6-aujVmCOwyo8pBjoezFg0LJ5aTgrxrLORQ/s128/COL128x.png',1),(5,'Chile','https://1.bp.blogspot.com/-8ejrB7stl8k/YM6b5wjHjmI/AAAAAAABjx4/085Dmjd0C7Atip7VKP6Mn7QIu0_7g2XXACLcBGAsYHQ/s128/Chile128x.png',1),(6,'Peru','https://3.bp.blogspot.com/-jJHfDp9eU2o/WeZIXacAFAI/AAAAAAABPxc/yYrba2cUVV4FYm8HuSjTX-jrXPmH_3qqwCLcBGAs/s1600/Peru128x.png',1),(7,'Paraguay','https://2.bp.blogspot.com/-cuGrw42vWi8/WWwYpxogCoI/AAAAAAABM-E/A8iBeh5_RP0-BnYpt161p1zhsY3f2QZqQCLcBGAs/s1600/Paraguay128x.png',1),(8,'Ecuador','https://1.bp.blogspot.com/-AYh01-a9UOA/Xh39IM5B-DI/AAAAAAABXDM/fwYQAaIrBy4dluKMIcV8lQBDWJ7odq1hQCLcBGAsYHQ/s1600/Ecuador128x.png',1),(9,'Bolivia','https://1.bp.blogspot.com/-K48V-qDTuxs/YMKf5KFw2eI/AAAAAAABjtg/Ymxp30vT93IRazg9eIxhk2DWQf07cwYsACLcBGAsYHQ/s128/Bolivia128x.png',1),(10,'Venezuela','https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEjcBCiSxNoPg_tlhShpCipG2nDbRLM0djg9YQEJgbz9gUiCh9rD8WLfUwpNXwfMn3IE2QeRfzPM_0XnEBAlqtL81avGb5xWEsNEMvfKsua7jZ_U8bett1qUQ2JEAGpThubalNM8CsprW3Y-e4H8W84o9Q2iADHaCIJGeW7lGUhssG2tIVNXE-fIMupvDBs/s128/Venezuela128x.png',1);
/*!40000 ALTER TABLE `pais` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `partidos`
--

DROP TABLE IF EXISTS `partidos`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `partidos` (
  `id` int NOT NULL AUTO_INCREMENT,
  `id_edicion_torneo` int NOT NULL,
  `id_fase` int NOT NULL,
  `id_equipo_local` int NOT NULL,
  `id_equipo_visitante` int NOT NULL,
  `goles_local` int DEFAULT NULL,
  `goles_visitante` int DEFAULT NULL,
  `ganador` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_fase_idx` (`id_fase`),
  KEY `fk_equipo_idx` (`id_equipo_local`,`id_equipo_visitante`),
  KEY `fk_Ediciontorneo_idx` (`id_edicion_torneo`),
  KEY `fk_equipo2` (`id_equipo_visitante`),
  CONSTRAINT `fk_Ediciontorneo` FOREIGN KEY (`id_edicion_torneo`) REFERENCES `edicion_torneos` (`id`),
  CONSTRAINT `fk_equipo1` FOREIGN KEY (`id_equipo_local`) REFERENCES `equipos` (`id`),
  CONSTRAINT `fk_equipo2` FOREIGN KEY (`id_equipo_visitante`) REFERENCES `equipos` (`id`),
  CONSTRAINT `fk_fases` FOREIGN KEY (`id_fase`) REFERENCES `fases` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `partidos`
--

LOCK TABLES `partidos` WRITE;
/*!40000 ALTER TABLE `partidos` DISABLE KEYS */;
/*!40000 ALTER TABLE `partidos` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `torneos`
--

DROP TABLE IF EXISTS `torneos`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `torneos` (
  `id` int NOT NULL AUTO_INCREMENT,
  `nombre` varchar(350) NOT NULL,
  `logo` varchar(350) NOT NULL,
  `id_confederacion` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_confederaciones_idx` (`id_confederacion`),
  CONSTRAINT `fk_confederaciones` FOREIGN KEY (`id_confederacion`) REFERENCES `confederacions` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `torneos`
--

LOCK TABLES `torneos` WRITE;
/*!40000 ALTER TABLE `torneos` DISABLE KEYS */;
INSERT INTO `torneos` VALUES (1,'Copa Libertadores','https://upload.wikimedia.org/wikipedia/pt/4/4b/Conmebol_Libertadores_Bridgestone_logo.png',1),(2,'Copa Sudamericana','https://upload.wikimedia.org/wikipedia/pt/e/e4/Conmebol_Sudamericana_logo.png',1),(3,'Champions League','https://upload.wikimedia.org/wikipedia/pt/9/9b/116px-UEFA_Champions_League_logo_2_svg.png',2),(4,'Europa League','https://upload.wikimedia.org/wikipedia/commons/thumb/9/9d/2015_UEFA_Europa_League_logo_%282%29.svg/731px-2015_UEFA_Europa_League_logo_%282%29.svg.png',2);
/*!40000 ALTER TABLE `torneos` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-12-21 13:31:13
