CREATE TABLE `account.users_account` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `account_unique_id` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `firstname` varchar(100) NOT NULL,
  `surename` varchar(100) DEFAULT NULL,
  `username` varchar(100) NOT NULL,
  `password` varchar(50) NOT NULL,
  `email` varchar(100) NOT NULL,
  `phone` varchar(100) NOT NULL,
  `role` varchar(100) NOT NULL,
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NOT NULL,
  `deleted_at` timestamp NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `users_UN` (`account_unique_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8