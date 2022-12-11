CREATE DATABASE RideSharingPlatform;


USE RideSharingPlatform;

/*
DROP TABLE `Ride`;
DROP TABLE `Passenger`;
DROP TABLE `Rider`;
DROP TABLE `User`;
*/

/**=========== Create Tables =================**/

/** ====== TableL User ======== **/
CREATE TABLE `User` (
`user_id` int NOT NULL AUTO_INCREMENT,
`user_type` varchar(10) NOT NULL,
`email_address` varchar(100) NOT NULL,
`password` varchar(100) NOT NULL,

PRIMARY KEY (`user_id`)
);


/**====== Table: Passenger ======**/ 
CREATE TABLE `Passenger`
(
 `passenger_id`       int NOT NULL,
 `first_name`    varchar(50) NOT NULL ,
 `last_name`     varchar(50) NOT NULL ,
 `mobile_number` varchar(25) NOT NULL ,
 
PRIMARY KEY (`passenger_id`),
KEY `FK_1` (`passenger_id`),
CONSTRAINT `FK_1` FOREIGN KEY `FK_1` (`passenger_id`) REFERENCES `User` (`user_id`)
);


/**====== Table: Rider ======**/ 
CREATE TABLE `Rider`
(
 `rider_id`      int NOT NULL,
 `first_name`    varchar(50) NOT NULL ,
 `last_name`     varchar(50) NOT NULL ,
 `mobile_number` varchar(25) NOT NULL ,
 `ic_number`     varchar(20) NOT NULL ,
 `car_lic_number` varchar(20) NOT NULL ,

PRIMARY KEY (`rider_id`),
KEY `FK_2` (`rider_id`),
CONSTRAINT `FK_2` FOREIGN KEY `FK_2` (`rider_id`) REFERENCES `User` (`user_id`)
);


/**====== Table: Ride ======**/ 
CREATE TABLE `Ride`
(
 `ride_id`       int NOT NULL AUTO_INCREMENT,
 `passenger_id`  int NOT NULL,
 `passenger_name`varchar(100) NOT NULL ,
 `passenger_phone`varchar(25) NOT NULL ,
 `rider_id`      int NULL ,
 `rider_name`    varchar(100) NULL ,
 `rider_phone`   varchar(25) NULL ,
 `car_lic_number` varchar(20) NULL ,
 `pick_up_code`  varchar(6) NOT NULL ,
 `drop_off_code` varchar(6) NOT NULL ,
 `ride_status`   varchar(20) NOT NULL ,

PRIMARY KEY (`ride_id`),
KEY `FK_3` (`passenger_id`),
CONSTRAINT `FK_3` FOREIGN KEY `FK_3` (`passenger_id`) REFERENCES `Passenger` (`passenger_id`),
KEY `FK_4` (`rider_id`),
CONSTRAINT `FK_4` FOREIGN KEY `FK_4` (`rider_id`) REFERENCES `Rider` (`rider_id`)
);


/** INSERTING DATA **/
INSERT INTO User(user_type, email_address, password)
VALUE ('passenger', 'chenhan@gmail.com', '12345678');
INSERT INTO User(user_type, email_address, password)
VALUES ('rider', 'dyhy@gmail.com', '12345678');

INSERT INTO Passenger
VALUES (1, 'Chen','Han','99887755');
INSERT INTO Rider
VALUES (2, 'Daryl','Yee','99887766','S1232123', 'B353234');

INSERT INTO Ride(passenger_id, passenger_name, passenger_phone, rider_id, rider_name, rider_phone, car_lic_number, pick_up_code, drop_off_code, ride_status)
VALUES (1, 'Chen Han', '99887755', 2, 'Daryl Yee', '99887766', 'B353234', '120324', '143234', "Completed" );