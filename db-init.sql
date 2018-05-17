CREATE DATABASE IF NOT EXISTS `cheffy` ;
USE `cheffy`;

-- CREATE TABLE "activity" --
CREATE TABLE `activity` ( 
	`id` Char( 36 ) NULL,
	`created_at` DateTime NOT NULL ON UPDATE CURRENT_TIMESTAMP,
	`email` VarChar( 60 ) NULL,
	`environment` VarChar( 60 ) NULL,
	`component` VarChar( 60 ) NULL,
	`message` VarChar( 255 ) NULL,
	`data` JSON NULL )
ENGINE = InnoDB;