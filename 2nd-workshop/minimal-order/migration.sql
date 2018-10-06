CREATE DATABASE `workshop_minimal_order`;

USE `workshop_minimal_order`;

CREATE TABLE `orders` (
	order_id INT NOT NULL AUTO_INCREMENT,
	order_number VARCHAR(15) NOT NULL,
	total BIGINT NOT NULL,
	created_at TIMESTAMP NOT NULL,
	PRIMARY KEY (`order_id`),
	UNIQUE KEY `order_number_unique` (`order_number`)
) ENGINE=InnoDB DEFAULT CHARACTER SET=utf8;