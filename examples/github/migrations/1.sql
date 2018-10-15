CREATE TABLE `github_profile`(
	id INT NOT NULL,
	name VARCHAR(200) NOT NULL,
	email VARCHAR(200) NOT NULL,
	avatar_url VARCHAR(200) NOT NULL,
	bio VARCHAR(200) NOT NULL,
	public_repos INT NOT NULL,
	created_at TIMESTAMP NOT NULL,
	PRIMARY KEY (`id`),
	UNIQUE KEY `github_profile_email_unique` (`email`)
) ENGINE=InnoDB DEFAULT CHARACTER SET=utf8;