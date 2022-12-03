
DROP DATABASE RideSharingPlatform 


CREATE DATABASE RideSharingPlatform;

USE RideSharingPlatform;


DROP TABLE `User`;

DROP TABLE `Ride`;


/**=========== Create Tables =================**/

/**====== Table: User ======**/ 
CREATE TABLE `User`
(
 `user_id`       varchar(5) NOT NULL ,
 `user_type`     char(10) NOT NULL,
 `first_name`    varchar(50) NOT NULL ,
 `last_name`     varchar(50) NOT NULL ,
 `mobile_number` varchar(25) NOT NULL ,
 `email_address` varchar(100) NOT NULL ,
 `id_number`      varchar(20) NOT NULL ,
 `car_lic_number` varchar(20) NOT NULL ,

PRIMARY KEY (`user_id`)
);

/**====== Table: Ride ======**/ 
CREATE TABLE `Ride`
(
 `ride_id`       varchar(5) NOT NULL ,
 `user_id`       varchar(5) NOT NULL ,
 `driver_id`     varchar(5) NOT NULL ,
 `pick_up_code`  varchar(6) NOT NULL ,
 `drop_off_code` varchar(6) NOT NULL ,
 `ride_status`   varchar(20) NOT NULL ,

PRIMARY KEY (`ride_id`),
KEY `FK_2` (`user_id`),
CONSTRAINT `FK_2` FOREIGN KEY `FK_2` (`user_id`) REFERENCES `User` (`user_id`),
KEY `FK_3` (`driver_id`),
CONSTRAINT `FK_3` FOREIGN KEY `FK_3` (`driver_id`) REFERENCES `User` (`user_id`)
);


/** INSERTING DATA **/
INSERT INTO User VALUES ('P0001', 'Passenger', 'Chen','Han','98294455','chenhan@gmail.com', 'NULL', 'NULL');
INSERT INTO User VALUES ('P0002', 'Driver', 'Daryl','Yee','9982332','dyhy@gmail.com', 'S1242345', 'B353234');
INSERT INTO Ride VALUES ('R0001', 'P0001', 'P0002', '120324', '143234', "On-going" );


SELECT * FROM User;
SELECT * FROM Ride;




 




