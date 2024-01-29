-- Delete user if already exists
DROP USER IF EXISTS 'user1'@'localhost';

-- Create user with all privileges
CREATE USER 'user1'@'localhost' IDENTIFIED BY 'secret_password';
GRANT ALL PRIVILEGES ON *.* TO 'user1'@'localhost';
--
CREATE DATABASE  IF NOT EXISTS `my_db` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `my_db`;
-- MySQL dump 10.13  Distrib 8.0.25, for Linux (x86_64)
--
-- Host: localhost    Database: my_db
-- ------------------------------------------------------
-- Server version	8.0.29-0ubuntu0.20.04.3

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
-- Table structure for table `products`
--

DROP TABLE IF EXISTS `products`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `products` (
  `id` int DEFAULT NULL,
  `name` varchar(50) DEFAULT NULL,
  `quantity` int DEFAULT NULL,
  `code_value` varchar(50) DEFAULT NULL,
  `is_published` varchar(50) DEFAULT NULL,
  `expiration` date DEFAULT NULL,
  `price` decimal(5,2) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `products`
--

LOCK TABLES `products` WRITE;
/*!40000 ALTER TABLE `products` DISABLE KEYS */;
INSERT INTO `products` VALUES (1,'Corn Shoots',244,'0009-1111','0','2022-01-08',23.27),(2,'Shrimp - Baby, Cold Water',174,'49288-0877','0','2022-08-04',52.12),(3,'Sprouts - Onion',136,'0268-6518','1','2021-12-27',91.95),(4,'Triple Sec - Mcguinness',107,'13537-457','0','2021-06-23',72.60),(5,'Chervil - Fresh',81,'49430-046','0','2022-05-28',34.46),(6,'Wine - German Riesling',212,'24385-804','0','2022-07-07',5.19),(7,'Oil - Sunflower',169,'59779-590','0','2022-07-03',7.24),(8,'Persimmons',238,'45802-327','0','2021-04-14',60.65),(9,'Beer - Labatt Blue',23,'48951-1215','1','2022-06-23',32.99),(10,'Ranchero - Primerba, Paste',59,'0268-1173','1','2021-04-02',17.02),(11,'Tomatillo',95,'0264-7730','0','2021-10-28',92.21),(12,'Pasta - Rotini, Colour, Dry',61,'68151-3826','1','2021-03-19',10.66),(13,'Sherry - Dry',86,'51079-485','1','2021-12-19',63.91),(14,'Chocolate - Unsweetened',67,'68788-9799','1','2022-02-22',36.01),(15,'Soupfoamcont12oz 112con',224,'11084-050','1','2022-06-27',95.66),(16,'Pasta - Gnocchi, Potato',7,'42507-340','1','2022-09-10',39.75),(17,'Beef - Sushi Flat Iron Steak',88,'42254-206','1','2021-12-14',38.50),(18,'Steamers White',217,'51386-737','0','2022-07-05',57.85),(19,'Pasta - Angel Hair',171,'0378-2017','1','2021-03-18',20.37),(20,'Parsley - Dried',197,'21695-130','0','2022-04-23',42.02),(21,'Maintenance Removal Charge',150,'42549-549','1','2021-10-09',77.80),(22,'Dome Lid Clear P92008h',52,'0023-4964','1','2021-09-15',54.49),(23,'Cookies - Oreo, 4 Pack',37,'54868-6276','0','2021-05-20',26.72),(24,'Nantucket - Kiwi Berry Cktl.',15,'65841-777','0','2021-09-18',11.05),(25,'Cookie - Oatmeal',75,'68151-0314','1','2022-05-09',25.62),(26,'Soho Lychee Liqueur',22,'68258-3016','1','2022-08-05',62.07),(27,'Bread - Bistro Sour',131,'0121-0671','1','2021-08-27',92.56),(28,'Chocolate - Pistoles, Lactee, Milk',171,'76181-002','1','2022-07-01',58.98),(29,'Grapefruit - White',22,'36987-1530','1','2022-01-26',73.36),(30,'Passion Fruit',236,'49288-0781','0','2022-03-12',63.36),(31,'Cookies - Englishbay Oatmeal',197,'42291-169','1','2021-12-16',94.37),(32,'Lettuce - Belgian Endive',179,'0378-0152','0','2022-04-02',48.29),(33,'Vaccum Bag 10x13',39,'0406-9907','0','2021-03-18',20.29),(34,'Chocolate - Dark',75,'54575-463','1','2021-04-04',88.37),(35,'Raspberries - Fresh',27,'31722-545','0','2021-12-23',38.82),(36,'Cattail Hearts',98,'45802-472','1','2022-03-26',15.84),(37,'Salt - Sea',34,'66116-360','0','2022-07-30',50.39),(38,'Cheese - Swiss',12,'43493-0001','0','2021-08-30',38.27),(39,'Pasta - Cheese / Spinach Bauletti',86,'48102-102','0','2021-08-31',64.17),(40,'Wine - Sake',216,'0603-0839','1','2021-09-25',39.18),(41,'Tea - Black Currant',19,'59011-458','1','2021-06-17',72.63),(42,'Cheese - Mozzarella, Shredded',243,'36800-099','0','2021-06-05',82.92),(43,'Gooseberry',112,'68788-9834','0','2021-12-10',17.38),(44,'Glass Clear 8 Oz',52,'69244-1001','0','2021-10-09',35.13),(45,'Bread - White, Unsliced',62,'60512-1005','0','2021-05-31',35.07),(46,'Syrup - Monin, Amaretta',139,'49348-559','1','2022-06-03',90.03),(47,'Temperature Recording Station',5,'0942-9395','0','2022-07-19',4.03),(48,'Cheese - Brie, Triple Creme',145,'10702-040','1','2022-06-11',36.63),(49,'Beer - Maudite',204,'15127-738','0','2022-08-22',8.05),(50,'Sesame Seed Black',217,'49884-835','0','2022-01-30',82.25),(51,'Pomegranates',200,'68026-528','0','2021-12-29',50.04),(52,'Wine - Placido Pinot Grigo',18,'52959-991','1','2021-10-02',14.48),(53,'Muffin Mix - Oatmeal',161,'49349-139','0','2021-04-08',91.21),(54,'Beans - Black Bean, Preserved',21,'11410-564','0','2021-05-04',53.26),(55,'Orange - Canned, Mandarin',162,'49738-078','1','2021-10-02',9.13),(56,'Towel - Roll White',3,'52125-232','1','2022-09-10',84.07),(57,'Pail With Metal Handle 16l White',99,'64578-0087','1','2021-07-05',34.05),(58,'Wine - Black Tower Qr',228,'0187-0798','0','2021-12-20',52.98),(59,'Duck - Whole',192,'0409-1755','1','2022-07-27',7.81),(60,'Bag Stand',131,'24470-913','1','2021-11-30',74.23),(61,'Cardamon Seed / Pod',203,'67877-220','1','2022-05-13',94.99),(62,'Vermacelli - Sprinkles, Assorted',110,'43547-254','0','2021-03-20',87.61),(63,'Rolled Oats',124,'49825-128','0','2022-04-21',53.18),(64,'Salad Dressing',243,'36987-2644','0','2021-10-23',50.24),(65,'Crab Meat Claw Pasteurise',101,'55910-402','0','2022-03-19',95.80),(66,'Soup - French Can Pea',96,'10019-955','1','2022-09-01',35.85),(67,'Trout - Rainbow, Frozen',23,'0591-3560','1','2021-05-30',55.11),(68,'Swordfish Loin Portions',40,'55670-122','0','2022-06-12',47.39),(69,'Wine - Red, Wolf Blass, Yellow',87,'43269-648','0','2022-05-13',34.79),(70,'Bread Base - Toscano',64,'36987-3086','0','2021-11-02',88.05),(71,'Cloves - Ground',213,'55319-140','0','2021-07-14',65.20),(72,'Egg - Salad Premix',98,'63777-165','0','2022-02-27',86.82),(73,'Sage Derby',114,'0338-1055','1','2022-09-10',88.12),(74,'Plasticknivesblack',98,'51655-501','0','2022-07-30',28.27),(75,'Kaffir Lime Leaves',27,'36987-2370','1','2021-09-16',56.42),(76,'Breakfast Quesadillas',194,'49348-405','0','2022-01-24',1.22),(77,'Chips - Potato Jalapeno',41,'60232-2582','0','2021-07-28',70.56),(78,'Soap - Pine Sol Floor Cleaner',171,'36987-1854','0','2021-05-17',49.12),(79,'Wine - Casillero Deldiablo',200,'57664-441','1','2021-05-14',10.82),(80,'Lentils - Red, Dry',156,'55154-6970','1','2022-06-05',8.64),(81,'Beer - True North Strong Ale',61,'0206-2405','1','2022-06-12',29.21),(82,'Gingerale - Schweppes, 355 Ml',83,'65862-526','1','2022-09-05',31.88),(83,'Capers - Ox Eye Daisy',154,'11822-3300','0','2022-07-08',14.98),(84,'Tomato - Tricolor Cherry',147,'33342-057','1','2021-07-19',54.13),(85,'Jam - Raspberry,jar',158,'49288-0249','1','2021-06-19',12.98),(86,'Chevril',207,'36987-2164','0','2022-08-27',31.15),(87,'Pastry - Cheese Baked Scones',168,'51630-004','1','2022-06-03',65.45),(88,'Butter - Unsalted',82,'59779-180','0','2022-05-16',35.47),(89,'Cookie Dough - Peanut Butter',129,'68599-6110','0','2021-04-24',89.36),(90,'Fudge - Chocolate Fudge',188,'11523-0259','1','2022-09-02',71.88),(91,'Truffle Shells - White Chocolate',110,'52125-304','1','2022-02-26',2.16),(92,'Dish Towel',214,'0603-2483','1','2021-11-03',72.51),(93,'Molasses - Fancy',68,'65044-1216','1','2021-12-03',11.16),(94,'Peppercorns - Pink',98,'11410-007','1','2022-02-09',90.84),(95,'Cake Circle, Foil, Scallop',200,'46122-027','0','2021-03-27',55.34),(96,'Bread - Mini Hamburger Bun',28,'68026-105','0','2021-06-16',12.91),(97,'Beer - Tetleys',37,'54868-4379','1','2022-09-08',10.52),(98,'Iced Tea - Lemon, 460 Ml',35,'48951-1199','0','2022-06-25',56.69),(99,'Beans - Yellow',36,'60681-0102','1','2022-06-12',45.65),(100,'Peppercorns - Green',34,'64117-115','1','2021-07-24',92.31),(101,'Steam Pan Full Lid',70,'43538-191','0','2022-09-05',50.55),(102,'Nantucket - Carrot Orange',187,'63148-164','0','2021-12-09',13.67),(103,'Flour - Teff',132,'55154-4378','1','2021-08-08',36.80),(104,'Venison - Striploin',176,'68084-692','0','2022-01-12',6.51),(105,'Lamb - Sausage Casings',72,'59667-0103','1','2021-04-28',7.98),(106,'Dates',79,'59779-974','1','2021-09-16',91.00),(107,'Oil - Safflower',63,'66129-101','1','2022-05-11',28.01),(108,'Clams - Canned',19,'68180-236','0','2021-11-01',93.11),(109,'Pastry - Choclate Baked',170,'0093-1006','1','2021-12-20',92.91),(110,'Poppy Seed',97,'0054-8084','0','2021-12-13',32.03),(111,'Longos - Greek Salad',111,'60760-911','0','2021-05-08',69.21),(112,'Bag Stand',13,'42023-136','0','2022-03-26',39.73),(113,'Veal - Provimi Inside',50,'63629-2573','1','2021-06-10',82.79),(114,'Wine - White, Lindemans Bin 95',152,'57955-5080','0','2022-06-30',65.05),(115,'Sprouts - Corn',88,'16714-041','0','2021-05-05',21.41),(116,'Snapple - Mango Maddness',126,'68016-125','1','2022-08-15',85.35),(117,'Table Cloth 54x54 Colour',65,'0615-7521','1','2022-02-19',75.52),(118,'Juice - Orange 1.89l',25,'76329-8261','0','2022-02-15',65.93),(119,'Skirt - 24 Foot',70,'24236-995','0','2021-03-23',79.10),(120,'Lemonade - Mandarin, 591 Ml',172,'64117-714','0','2021-12-09',20.93),(121,'Cod - Black Whole Fillet',244,'58668-4101','1','2021-12-20',79.45),(122,'Sugar Thermometer',29,'0527-1301','0','2021-10-29',39.83),(123,'Compound - Raspberry',152,'68382-179','1','2022-01-16',91.95),(124,'Cookie Trail Mix',36,'51346-257','0','2022-08-22',75.66),(125,'Beef - Top Sirloin',91,'64376-132','0','2021-06-15',79.27),(126,'Bread - Ciabatta Buns',84,'43598-225','1','2021-06-24',59.98),(127,'Soup Campbells - Tomato Bisque',26,'21695-969','0','2021-05-03',96.24),(128,'Radish - Pickled',208,'52959-398','1','2021-08-03',87.16),(129,'Butter Sweet',185,'67510-0085','1','2022-01-01',9.41),(130,'Jam - Raspberry',227,'68400-706','0','2021-06-16',50.71),(131,'Pork - Backfat',92,'39822-3015','1','2021-08-20',32.58),(132,'Yoplait Drink',140,'68788-9165','0','2021-07-27',14.39),(133,'Pastry - Cheese Baked Scones',236,'50181-0016','0','2021-09-29',84.77),(134,'Bread - Pumpernickle, Rounds',61,'63629-2949','0','2021-12-20',35.97),(135,'Veal - Chops, Split, Frenched',223,'68180-655','1','2021-03-16',55.88),(136,'Wine - Red, Marechal Foch',94,'51285-595','0','2022-02-27',91.04),(137,'Beets - Candy Cane, Organic',37,'67296-0538','0','2022-06-03',23.90),(138,'Juice - Clam, 46 Oz',66,'53329-938','0','2022-06-25',57.04),(139,'Chocolate Bar - Smarties',51,'42957-002','0','2022-08-03',59.82),(140,'Mushroom - King Eryingii',156,'0268-1094','0','2022-06-25',22.50),(141,'Gingerale - Schweppes, 355 Ml',132,'54868-5841','1','2022-05-27',8.60),(142,'Wine - Hardys Bankside Shiraz',219,'49349-626','1','2021-10-13',91.57),(143,'Kaffir Lime Leaves',249,'49288-0146','1','2022-07-12',17.02),(144,'Kellogs Special K Cereal',12,'33261-591','1','2021-05-31',5.22),(145,'Cup - 3.5oz, Foam',223,'54868-1173','0','2021-09-28',35.76),(146,'Beef Cheek Fresh',30,'53942-311','1','2021-05-11',32.12),(147,'Beef - Tenderloin',1,'43068-106','0','2022-03-10',39.06),(148,'Paper Cocktail Umberlla 80 - 180',147,'57520-0324','1','2021-05-15',82.87),(149,'Pineapple - Canned, Rings',81,'35000-608','1','2022-04-12',34.78),(150,'Veal Inside - Provimi',124,'49643-460','0','2022-01-08',64.26),(151,'Goulash Seasoning',110,'55910-199','0','2021-08-10',59.24),(152,'Juice - Cranberry, 341 Ml',159,'57955-0065','1','2022-07-27',91.05),(153,'Pastry - Chocolate Marble Tea',6,'68428-037','1','2021-07-24',69.94),(154,'Quinoa',39,'59667-0096','1','2021-03-13',65.15),(155,'Island Oasis - Ice Cream Mix',79,'62011-0006','0','2021-05-30',20.23),(156,'Basil - Dry, Rubbed',225,'51607-001','0','2021-07-25',13.53),(157,'Beef Cheek Fresh',43,'51389-204','0','2021-08-19',79.96),(158,'Scrubbie - Scotchbrite Hand Pad',175,'54569-6100','1','2021-05-06',72.55),(159,'Cookie - Oreo 100x2',211,'41442-150','0','2022-01-23',13.43),(160,'Appetizer - Shrimp Puff',177,'50563-155','0','2022-08-23',29.83),(161,'Island Oasis - Pina Colada',77,'50241-141','0','2021-03-15',83.32),(162,'Graham Cracker Mix',29,'55154-0536','0','2022-09-02',7.92),(163,'Poppy Seed',7,'0407-0690','1','2022-05-21',27.27),(164,'Anchovy Paste - 56 G Tube',118,'55301-007','1','2022-04-21',2.69),(165,'Bread - Sticks, Thin, Plain',213,'62175-129','1','2021-12-04',60.29),(166,'Coffee - Flavoured',1,'0409-4857','0','2022-07-21',31.88),(167,'Macaroons - Homestyle Two Bit',202,'63629-1494','0','2022-05-09',9.67),(168,'Energy - Boo - Koo',89,'50436-0922','1','2021-08-20',13.54),(169,'Sage - Fresh',85,'10424-161','0','2021-11-27',89.57),(170,'Nantucket Pine Orangebanana',237,'0268-1505','0','2022-08-30',51.65),(171,'Sauce - Hp',228,'60793-851','0','2021-10-17',74.07),(172,'Pork - Tenderloin, Frozen',144,'63629-4492','1','2022-08-14',89.39),(173,'Glucose',201,'21695-125','1','2021-10-24',90.99),(174,'Raisin - Golden',86,'57520-0642','1','2021-11-03',12.11),(175,'Brandy Apricot',146,'36800-422','0','2022-05-11',40.27),(176,'Capers - Ox Eye Daisy',90,'54868-5662','1','2022-03-24',42.33),(177,'Puff Pastry - Slab',16,'21695-365','0','2021-05-16',48.79),(178,'Dome Lid Clear P92008h',29,'52685-324','0','2022-01-28',15.69),(179,'Salt And Pepper Mix - Black',219,'0440-1771','1','2021-03-21',79.26),(180,'Artichoke - Bottom, Canned',196,'0093-7658','1','2021-08-30',23.28),(181,'Pasta - Angel Hair',228,'0093-3010','0','2022-06-19',17.91),(182,'Muffin Batt - Choc Chk',141,'13537-447','0','2021-07-08',31.89),(183,'Beef Flat Iron Steak',14,'0054-0003','0','2022-03-03',94.49),(184,'Puree - Mango',147,'60589-005','1','2021-11-10',81.89),(185,'Beer - Camerons Cream Ale',206,'0268-6676','1','2021-11-04',66.32),(186,'Spinach - Frozen',121,'41250-105','0','2021-08-21',79.36),(187,'Bagel - Everything Presliced',153,'49672-100','1','2022-01-11',20.44),(188,'Absolut Citron',121,'55154-4623','0','2021-03-11',65.81),(189,'Honey - Liquid',176,'41520-300','0','2021-06-02',55.05),(190,'Pork - Suckling Pig',187,'61957-1018','0','2021-03-15',19.54),(191,'Beef Striploin Aaa',245,'53645-1021','0','2021-07-26',73.87),(192,'Pepper - Jalapeno',137,'0007-3260','1','2022-07-01',77.85),(193,'Glass - Juice Clear 5oz 55005',126,'53499-5571','1','2021-10-16',49.82),(194,'Devonshire Cream',150,'0363-6230','1','2022-04-03',8.51),(195,'Lobster - Baby, Boiled',21,'0074-6624','1','2021-04-14',37.33),(196,'Soupcontfoam16oz 116con',45,'11673-599','1','2022-05-12',65.19),(197,'French Pastry - Mini Chocolate',240,'0615-1556','0','2022-01-25',7.19),(198,'Sobe - Berry Energy',205,'0069-0122','1','2022-02-12',91.99),(199,'Tea - Jasmin Green',238,'43269-720','0','2022-02-03',34.63),(200,'Scallops - 20/30',229,'68788-9100','1','2022-03-06',30.75);
/*!40000 ALTER TABLE `products` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-05-12 11:30:53
