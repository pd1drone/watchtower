CREATE DATABASE watchtower;

USE watchtower;

CREATE TABLE `Users` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `Username` varchar(255),
  `Password` varchar(255),
  `IsAdmin` bool
);

CREATE TABLE `HistoryLogs` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `WaterLevel` double,
  `Timestamp` integer
);

-- username= admin , password = admin123
INSERT INTO Users (Username, Password, IsAdmin) VALUES ('admin','0192023a7bbd73250516f069df18b500',true);