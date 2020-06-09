CREATE TABLE auth.login_history (
	id INT UNSIGNED auto_increment NOT NULL,
	account_unique_id varchar(100) NOT NULL,
	username varchar(100) NOT NULL,
	`role` varchar(100) NOT NULL,
	login_at INT UNSIGNED NOT NULL,
	CONSTRAINT login_history_PK PRIMARY KEY (id),
	CONSTRAINT login_history_UN UNIQUE KEY (account_unique_id)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci;
