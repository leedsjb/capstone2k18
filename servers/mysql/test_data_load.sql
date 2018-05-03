/*
test_data_load.sql
Created: Wednesday Apil 25, 2018
Modified:
Authors: J. Benjamin Leeds
License: None

This script loads the necessary data into Google Cloud MySQL to begin receiving data additions
and updates from Flight Vector on-premise via Google Cloud Pub/Sub.

*/


DROP TABLE IF EXISTS tblAIRCRAFT_STATUS;
CREATE TABLE tblAIRCRAFT_STATUS(
    aircraftStatusID INTEGER AUTO_INCREMENT PRIMARY KEY,
    aircraftStatusName NVARCHAR(100) NOT NULL,
    aircraftStatusDescription NVARCHAR(100)
);
INSERT INTO tblAIRCRAFT_STATUS(aircraftStatusName, aircraftStatusDescription)
VALUES
    ("VFR ONLY - Aircraft Issue", "Aircraft can only fly in Visual Meterological Conditions"),
    ("Delayed Maintenance", "Maintenance not completed on time"),
    ("Delayed Other", "Delayed due to other reasons"),
    ("Hangared", "Aircraft presently in hangar"),
    ("Heavy on Fuel", "Check aircraft fuel weight for weight and balance, payload may be limited"),
    ("Last Out", "*************** What is this?"),
    ("No Riders", "Essential personnel only"),
    ("USA Only", "Cross-border flights prohibited"),
    ("2 Peds", "2 Pediatric flight nurses assigned"),
    ("2 Adult", "2 Adult flight nurses assigned"),
    ("Isolette OOS", "Isolette equipment out of service"),
    ("VFR Only - Pilot Issue", "Flight crew not presently qualified for Instrument Flight Rules"),
    ("MEL - Unable to fly into known icing", "Aircraft Minimum Equipment List prevents Flight Into
    Known Icing (FIKI)");

DROP TABLE IF EXISTS tblAIRCRAFT_TYPE;
CREATE TABLE tblAIRCRAFT_TYPE(
    aircraft_type_id INTEGER AUTO_INCREMENT PRIMARY KEY, 
    aircraft_type_title NVARCHAR(50) NOT NULL,
    aircraft_type_desc NVARCHAR(250),
    aircraft_type_category NVARCHAR(25) NOT NULL,
    aircraft_type_manufacturer NVARCHAR(50) NOT NULL
);
INSERT INTO tblAIRCRAFT_TYPE(
    aircraft_type_manufacturer, aircraft_type_title, aircraft_type_category
)
VALUES
    ("Diamond", "DA42-NG", "Fixed-wing"), ("Diamond", "DA40", "Fixed-wing"), 
    ("Bell", "206", "Rotorcraft");

DROP TABLE IF EXISTS tblAC_LEVEL_OF_CARE;
CREATE TABLE tblAC_LEVEL_OF_CARE(
    levelOfCareID INTEGER AUTO_INCREMENT PRIMARY KEY,
    levelOfCareName NVARCHAR(100) NOT NULL
)
INSERT INTO tblAC_LEVEL_OF_CARE(levelOfCareName)
VALUES 
    ("Neonatal"), ("Pediatric RN Onboard"), ("Alaska Pediatric");

DROP TABLE IF EXISTS tblROLE;
CREATE TABLE tblROLE(
    roleID INTEGER AUTO_INCREMENT PRIMARY KEY,
    roleName NVARCHAR(100) NOT NULL,
    roleDescription NVARCHAR(100)
)
INSERT INTO tblROLE(roleName)
VALUES 
    ("Pilot PIC"), ("Adult RN"), ("Pediatric RN");

DROP TABLE IF EXISTS tblAIRCRAFT;
CREATE TABLE tblAIRCRAFT(
    acID INTEGER AUTO_INCREMENT PRIMARY KEY, 
    acCallSign NVARCHAR(100) NOT NULL, 
    acNNumber NVARCHAR(100) NOT NULL,
    acType INTEGER FOREIGN KEY REFERENCES tblAIRCRAFT_TYPE(aircraftTypeID),

)


-- Create Personnel Tables

-- Personnel Group: Groups of personnel used primarily for messaging / organizational purposes
DROP TABLE IF EXISTS tblPERSONNEL_GROUP;
CREATE TABLE tblPERSONNEL_GROUP(
    personnelGroupID INTEGER AUTO_INCREMENT PRIMARY KEY, 
    personnelGroupName NVARCHAR(100) NOT NULL, 
    personnelGroupDesc NVARCHAR(100)
)
INSERT INTO tblPERSONNEL_GROUP(personnelGroupName)
VALUES(
    ("Mobile ECMO Team"), ("PAIP"), ("SEA Turbo"), ("SEA Turbo Pilots"), ("SEA Turbo RNs")
);