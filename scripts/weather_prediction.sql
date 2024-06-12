CREATE DATABASE weather_prediction;
CREATE TABLE `weather_prediction`.`predictions` (
  `weather_day` INT NOT NULL,
  `weather_status` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`weather_day`));
