CREATE TABLE `account.users_login_logs` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(100) NOT NULL,
  `role` varchar(100) NOT NULL,
  `login_at` int unsigned NOT NULL,
  `user_detail_id` int unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `login_history_FK` (`user_detail_id`),
  CONSTRAINT `login_history_FK` FOREIGN KEY (`user_detail_id`) REFERENCES `account.users_account` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8