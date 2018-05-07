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
    ("Augusta", "A109E", "Rotorcraft"),
    ("Lear", "31A", "Fixed-wing"), 
    ("Airbus", "H-135", "Rotorcraft"),
    ("Pilatus","PC-12","Fixed-wing");

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

CREATE TABLE tblAIRCRAFT(
    ac_id INTEGER AUTO_INCREMENT PRIMARY KEY,           -- integer
    ac_callsign NVARCHAR(100),                          -- AL1, AL2, etc.
    ac_n_number NVARCHAR(10),                            -- N123AS
    ac_type_id INTEGER,                                 
    ac_lat DECIMAL(9,6),
    ac_long DECIMAL(9,6),
    ac_loc_display_name NVARCHAR(50),
    ac_cell_phone BIGINT,
    ac_sat_phone BIGINT,
    FOREIGN KEY (ac_type_id) REFERENCES tblAIRCRAFT_TYPE(aircraft_type_id)
);

SELECT * FROM tblASSIGNED_STATUS
SELECT * FROM tblAIRCRAFT_TYPE

INSERT INTO tblAIRCRAFT(ac_n_number, ac_callsign, ac_type_id, ac_lat, ac_long, ac_cell_phone, ac_sat_phone)
VALUES
("N139AM", "AL6", (SELECT aircraft_type_id FROM tblAIRCRAFT_TYPE WHERE aircraft_type_title = "H-135"), 47.658114, 122.298400, 5551234567, 4441234567),
("N951AL", "AL8", (SELECT aircraft_type_id FROM tblAIRCRAFT_TYPE WHERE aircraft_type_title = "A109E"), 47.658114, 122.298400, 5551234568, 4441234568),
("N952AL", "AL5", (SELECT aircraft_type_id FROM tblAIRCRAFT_TYPE WHERE aircraft_type_title = "H-135"), 47.658114, 122.298400, 5551234569, 4441234569),
("N954AL", "AL2", (SELECT aircraft_type_id FROM tblAIRCRAFT_TYPE WHERE aircraft_type_title = "H-135"), 47.658114, 122.298400, 5551234570, 4441234570),
("N235UW", "AL7", (SELECT aircraft_type_id FROM tblAIRCRAFT_TYPE WHERE aircraft_type_title = "H-135"), 47.658114, 122.298400, 5551234571, 4441234571),
("N212AL", "TURBOSPARE", (SELECT aircraft_type_id FROM tblAIRCRAFT_TYPE WHERE aircraft_type_title = "PC-12"), 47.658114, 122.298400, 5551234572, 4441234572),
("N164AL", "LEARSPARE", (SELECT aircraft_type_id FROM tblAIRCRAFT_TYPE WHERE aircraft_type_title = "31A"), 47.658114, 122.298400, 5551234573, 4441234573),
("N165AL", "JNU F", (SELECT aircraft_type_id FROM tblAIRCRAFT_TYPE WHERE aircraft_type_title = "31A"), 47.658114, 122.298400, 5551234574, 4441234574)  

DROP TABLE IF EXISTS tblMISSION;
CREATE TABLE tblMISSION(
    mission_id INTEGER AUTO_INCREMENT PRIMARY KEY,
    aircraft_id INTEGER,
    requestor INTEGER,
    receiver INTEGER,
    mission_date DATE,
    tc_number VARCHAR(10),
    FOREIGN KEY(aircraft_id) REFERENCES tblAIRCRAFT(ac_id),
    FOREIGN KEY(requestor) REFERENCES tblAGENCY(agency_id),
    FOREIGN KEY(receiver) REFERENCES tblAGENCY(agency_id)
);


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