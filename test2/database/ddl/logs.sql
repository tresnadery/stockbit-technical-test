CREATE TABLE `logs` (
  `id` int NOT NULL AUTO_INCREMENT,
  `request` text,
  `response` text,
  `created_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci