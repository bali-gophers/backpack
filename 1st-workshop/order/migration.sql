CREATE DATABASE `workshop_order`;

USE `workshop_order`;

CREATE TABLE `order` (
	`order_id` INT AUTO_INCREMENT,
	`order_number` VARCHAR(15) NOT NULL UNIQUE,
	`full_name` VARCHAR(100) NOT NULL,
	`email` VARCHAR(50) NOT NULL,
	`total` INT NOT NULL,
	`created_at` TIMESTAMP NOT NULL,
	PRIMARY KEY (`order_id`)
) ENGINE=InnoDB DEFAULT CHARACTER SET=utf8;

CREATE TABLE `item` (
	`item_id` INT AUTO_INCREMENT,
	`order_id` INT NOT NULL,
	`title` VARCHAR(30) NOT NULL,
	`count` INT NOT NULL,
	`price` BIGINT NOT NULL,
	PRIMARY KEY (`item_id`),
	INDEX `item_order_id` (`order_id`)
) ENGINE=InnoDB DEFAULT CHARACTER SET=utf8;