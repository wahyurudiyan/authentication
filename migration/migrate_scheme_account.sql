CREATE TABLE account.users (
	id INT UNSIGNED auto_increment NOT NULL,
	account_unique_id varchar(100) NOT NULL,
	firstname varchar(100) NOT NULL,
	surename varchar(100) NULL,
	email varchar(100) NOT NULL,
	phone varchar(100) NOT NULL,
	created_at TIMESTAMP NULL,
	updated_at TIMESTAMP NULL,
	deleted_at TIMESTAMP NULL,
	CONSTRAINT users_PK PRIMARY KEY (id),
	CONSTRAINT users_UN UNIQUE KEY (account_uid)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8
COLLATE=utf8_general_ci;
