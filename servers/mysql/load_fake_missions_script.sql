/*
load_fake_missions_script.sql
Created: Saturday April 21, 2018
Modified: Tuesday April 24, 2018
Authors: J. Benjamin Leeds, Sabrina Niklaus, Caitlin Schaefer, and Julian Boss
License: None

This script loads a MySql v5.7 database hosted in Google Cloud SQL with 5000 rows of fake mission
data for Airlift Northwest Operations. The script uses stored procedures for the bulk of its
operations. Fake data is stored in a CSV in Google Cloud Storage and loaded into a 1NF table
via the Google Cloud Console GUI in Step 2. These data are then copied into an RDBMS schema for
future CRUD operations. 
*/

-- Step 0: Database Setup
CREATE DATABASE TEST
 
/** 
    Note: Re-run mySQL client at this point specifying the newly created TEST database as your
    default for all future queries. Alternatively prepend all your script queries with USE TEST;
**/
 
-- Step 1: create table for .csv import of 1NF data
DROP TABLE IF EXISTS `tblRAW`;
CREATE TABLE tblRAW(
   aircraft_lat DECIMAL(10,0),
   aircraft_long DECIMAL(10,0),
   model_title NVARCHAR(100),
   model_desc NVARCHAR(100),
   aircraft_category NVARCHAR(100),
   aircraft_callsign NVARCHAR(100),
   mission_date DATE,
   agency_name NVARCHAR(100),
   agency_area_code INTEGER,
   agency_phone NVARCHAR(100),
   address_street NVARCHAR(100),
   address_city NVARCHAR(100),
   address_state NVARCHAR(100),
   address_zip INTEGER
);
 
/**
    Step 2: Import 1NF CSV raw data to table following the instructions here:
    https://cloud.google.com/sql/docs/mysql/import-export/importing
**/
 
-- Delete column headers row from imported CSV data
DELETE FROM tblRAW WHERE model_title = 'model-title';

-- *************************************************************************************************

-- Clean up any existing tables and procedures
DROP TABLE IF EXISTS `tblRAW_WITH_PK`;
DROP TABLE IF EXISTS `tblRAW_TO_MISSION`;
DROP TABLE IF EXISTS `tblRAW_TO_AIRCRAFT`;
DROP TABLE IF EXISTS `tblRAW_TO_AGENCIES`;

DROP TABLE IF EXISTS `tblMISSION`;
DROP TABLE IF EXISTS `tblAGENCY`;
DROP TABLE IF EXISTS `tblAIRCRAFT`;

DROP TABLE IF EXISTS `tblADDRESS`;
DROP TABLE IF EXISTS `tblAIRCRAFT_MODELS`;

DROP PROCEDURE IF EXISTS `uspGetAddressID`;
DROP PROCEDURE IF EXISTS `uspGetAircraftModelID`;
DROP PROCEDURE IF EXISTS `uspGetAgencyID`;
DROP PROCEDURE IF EXISTS `uspPopulateAgency`;
DROP PROCEDURE IF EXISTS `uspGetAircraftID`; 
DROP PROCEDURE IF EXISTS `uspPopulateAircraft`;
DROP PROCEDURE IF EXISTS `uspGetMissionID`;
DROP PROCEDURE IF EXISTS `uspPopulateMissionsTable`;
DROP PROCEDURE IF EXISTS `uspPopulateMission`;

-- Step 3: create copy of table with PKs
CREATE TABLE tblRAW_WITH_PK(
   rowID INTEGER NOT NULL AUTO_INCREMENT PRIMARY KEY,
   aircraft_lat DECIMAL(9,6),
   aircraft_long DECIMAL(9,6),
   model_title NVARCHAR(100),
   model_desc NVARCHAR(100),
   aircraft_category NVARCHAR(100),
   aircraft_callsign NVARCHAR(100),
   mission_date DATE,
   agency_name NVARCHAR(100),
   agency_area_code INTEGER,
   agency_phone NVARCHAR(100),
   address_street NVARCHAR(100),
   address_city NVARCHAR(100),
   address_state NVARCHAR(100),
   address_zip INTEGER
);
 
INSERT INTO tblRAW_WITH_PK(aircraft_lat, aircraft_long, model_title, model_desc, aircraft_category,
aircraft_callsign, mission_date, agency_name, agency_area_code, agency_phone, address_street,
address_city, address_state, address_zip)
SELECT tblRAW.aircraft_lat, tblRAW.aircraft_long, tblRAW.model_title, tblRAW.model_desc,
tblRAW.aircraft_category, tblRAW.aircraft_callsign, tblRAW.mission_date, tblRAW.agency_name,
tblRAW.agency_area_code, tblRAW.agency_phone, tblRAW.address_street, tblRAW.address_city,
tblRAW.address_state, tblRAW.address_zip
FROM tblRAW;
 
CREATE TABLE tblADDRESS(
    address_id INTEGER AUTO_INCREMENT,
    address_street NVARCHAR(100),
    address_city NVARCHAR(100),
    address_state NVARCHAR(100),
    address_zip INTEGER,
    PRIMARY KEY (address_id)
);
 
CREATE PROCEDURE uspGetAddressID(
    IN street NVARCHAR(100),
    IN city NVARCHAR(100),
    IN in_state NVARCHAR(100),
    IN zip INT,
    OUT A_ID INT
)
BEGIN
	SET A_ID = (
        SELECT address_id FROM tblADDRESS 
        WHERE address_street = street
        AND address_city = city
        AND address_state = in_state
        AND address_zip = zip
    );
END;

CREATE TABLE tblAIRCRAFT_MODELS(
    model_id INTEGER AUTO_INCREMENT PRIMARY KEY,
    model_title NVARCHAR(100),
    model_desc NVARCHAR(100),
    aircraft_category NVARCHAR(16)
);
 
CREATE PROCEDURE uspGetAircraftModelID(
    IN title NVARCHAR(100),
    IN am_desc NVARCHAR(100),
    IN cat NVARCHAR(16),
    OUT AM_ID INT
)
BEGIN
    SET AM_ID = (
        SELECT model_id FROM tblAIRCRAFT_MODELS
        WHERE title = model_title
        AND am_desc = model_desc
        AND cat = aircraft_category
    );
END;
 

/*
Populate first layer lookup tables
Address and Aircraft_Type
*/
 
INSERT INTO tblADDRESS(address_street, address_city, address_state, address_zip)
SELECT DISTINCT tblRAW_WITH_PK.address_street, tblRAW_WITH_PK.address_city,
tblRAW_WITH_PK.address_state, tblRAW_WITH_PK.address_zip
FROM tblRAW_WITH_PK;
 
INSERT INTO tblAIRCRAFT_MODELS(model_title, model_desc, aircraft_category)
SELECT DISTINCT tblRAW_WITH_PK.model_title, tblRAW_WITH_PK.model_desc,
tblRAW_WITH_PK.aircraft_category
FROM tblRAW_WITH_PK;
 
-- tblAGENCY, tblAIRCRAFT, tblMISSION have foreign keys, require some specialized insert statements,
-- will process with loop over pk temp table

-- attributes for tblAIRCRAFT_MODELS and tblADDRESS not needed here because they are already stored
-- in tblAIRCRAFT and tblAGENCY as foreign keys
 
CREATE TABLE tblAGENCY (
    agency_id INTEGER NOT NULL AUTO_INCREMENT PRIMARY KEY,
    agency_name NVARCHAR(100),
    agency_area_code INTEGER,
    agency_phone NVARCHAR(100),
    address_id INTEGER,
    FOREIGN KEY (address_id) REFERENCES tblADDRESS(address_id)
);
 
CREATE PROCEDURE uspGetAgencyID(
    IN a_name NVARCHAR(100),
    IN area_code INT,
    IN phone NVARCHAR(100),
    OUT A_ID INT
)
BEGIN
    SET A_ID = (
        SELECT agency_id FROM tblAGENCY
        WHERE a_name = agency_name
        AND area_code = agency_area_code
        AND phone = agency_phone
    );
END;

CREATE TABLE tblRAW_TO_AGENCIES(
 	agency_ID INTEGER NOT NULL AUTO_INCREMENT PRIMARY KEY,
    agency_name NVARCHAR(100),
    agency_area_code INTEGER,
    agency_phone NVARCHAR(100),
    address_street NVARCHAR(100),
    address_city NVARCHAR(100),
    address_state NVARCHAR(100),
    address_zip INTEGER
);
 
INSERT INTO tblRAW_TO_AGENCIES(
    agency_name, agency_area_code, agency_phone, address_street, 
    address_city, address_state, address_zip)
SELECT DISTINCT agency_name, agency_area_code, agency_phone, address_street, address_city,
address_state, address_zip
FROM tblRAW_WITH_PK;
 
CREATE PROCEDURE uspPopulateAgency()
BEGIN
    DECLARE err INT DEFAULT FALSE;
    DECLARE run INT;
    DECLARE CONTINUE HANDLER FOR SQLEXCEPTION
        BEGIN
            SET err = true;
        END;
    SET run = (SELECT COUNT(*) FROM tblRAW_TO_AGENCIES);
    
    WHILE run > 0 DO
        SET @agencyID = (SELECT MIN(agency_ID) FROM tblRAW_TO_AGENCIES);
        SET @agency_name = (SELECT agency_name FROM tblRAW_TO_AGENCIES WHERE agency_ID = @agencyID);
        SET @agency_area_code = (SELECT agency_area_code FROM tblRAW_TO_AGENCIES WHERE agency_ID = @agencyID);
        SET @agency_phone = (SELECT agency_phone FROM tblRAW_TO_AGENCIES WHERE agency_ID = @agencyID);
        SET @address_street = (SELECT address_street FROM tblRAW_TO_AGENCIES WHERE agency_ID = @agencyID);
        SET @address_city = (SELECT address_city FROM tblRAW_TO_AGENCIES WHERE agency_ID = @agencyID);
        SET @address_state = (SELECT address_state FROM tblRAW_TO_AGENCIES WHERE agency_ID = @agencyID);
        SET @address_zip = (SELECT address_zip FROM tblRAW_TO_AGENCIES WHERE agency_ID = @agencyID);
        
        CALL uspGetAddressID(
        @address_street, @address_city, @address_state, @address_zip, @address_ID);
        
        START TRANSACTION;
        INSERT INTO tblAGENCY(agency_name, agency_area_code, agency_phone, address_id)
        VALUES (@agency_name, @agency_area_code, @agency_phone, @address_ID);
        
        IF err THEN
            ROLLBACK;
        ELSE
            COMMIT;
        END IF;
        DELETE FROM tblRAW_TO_AGENCIES WHERE agency_ID = @agencyID;
        SET err = FALSE;
        SET run = run - 1;
    END WHILE;
END;
 
CALL uspPopulateAgency();

-- Create and populate aircraft table
 
CREATE TABLE tblAIRCRAFT(
    aircraft_id INTEGER AUTO_INCREMENT PRIMARY KEY,
    aircraft_callsign NVARCHAR(100),
    model_id INTEGER,
    aircraft_lat DECIMAL(9,6),
    aircraft_long DECIMAL(9,6),
    FOREIGN KEY (model_id) REFERENCES tblAIRCRAFT_MODELS(model_id)
);

-- Stored procedure used when populating missions
CREATE PROCEDURE uspGetAircraftID(
    IN callsign NVARCHAR(100),
    OUT A_ID INT
)
BEGIN
    SET A_ID = (
        SELECT aircraft_id FROM tblAIRCRAFT
        WHERE callsign = aircraft_callsign
    );
END;
 

CREATE TABLE tblRAW_TO_AIRCRAFT(
    aircraft_ID INTEGER NOT NULL AUTO_INCREMENT PRIMARY KEY,
    aircraft_lat DECIMAL(9,6),
    aircraft_long DECIMAL(9,6),
    model_title NVARCHAR(100),
    model_desc NVARCHAR(100),
    aircraft_category NVARCHAR(100),
    aircraft_callsign NVARCHAR(100)
);
 
-- copy data from RAW to RAW_AIRCRAFT
INSERT INTO tblRAW_TO_AIRCRAFT(
    aircraft_lat, aircraft_long, model_title, model_desc,
    aircraft_category, aircraft_callsign
)
SELECT DISTINCT aircraft_lat, aircraft_long, model_title, model_desc, aircraft_category,
aircraft_callsign
FROM tblRAW_WITH_PK;

CREATE PROCEDURE uspPopulateAircraft()
BEGIN
    DECLARE err INT DEFAULT FALSE;
    DECLARE run INT;
    DECLARE CONTINUE HANDLER FOR SQLEXCEPTION
    BEGIN
        SET err = true;
    END;
    
    SET run = (SELECT COUNT(*) FROM tblRAW_TO_AIRCRAFT);
    WHILE run > 0 DO -- iterate over each aircraft
    
        SET @aircraft_ID = (SELECT MIN(aircraft_ID) FROM tblRAW_TO_AIRCRAFT);
        SET @aircraft_lat = (SELECT aircraft_lat FROM tblRAW_TO_AIRCRAFT WHERE aircraft_ID = @aircraft_ID);
        SET @aircraft_long = (SELECT aircraft_long FROM tblRAW_TO_AIRCRAFT WHERE aircraft_ID = @aircraft_ID);
        SET @model_title = (SELECT model_title FROM tblRAW_TO_AIRCRAFT WHERE aircraft_ID = @aircraft_ID);
        SET @model_desc = (SELECT model_desc FROM tblRAW_TO_AIRCRAFT WHERE aircraft_ID = @aircraft_ID);
        SET @aircraft_category = (SELECT aircraft_category FROM tblRAW_TO_AIRCRAFT WHERE aircraft_ID = @aircraft_ID);
        SET @aircraft_callsign = (SELECT aircraft_callsign FROM tblRAW_TO_AIRCRAFT WHERE aircraft_ID = @aircraft_ID);
        
        CALL uspGetAircraftModelID(@model_title, @model_desc, @aircraft_category, @model_id);
        
        START TRANSACTION;
        INSERT INTO tblAIRCRAFT (aircraft_callsign, model_id, aircraft_lat, aircraft_long)
        VALUES (@aircraft_callsign, @model_ID, @aircraft_lat, @aircraft_long);
        
        IF err THEN
            ROLLBACK;
        ELSE
            COMMIT;
        END IF;
        
        DELETE FROM tblRAW_TO_AIRCRAFT WHERE aircraft_ID = @aircraft_ID;
        SET run = run - 1;
        SET err = FALSE;
    END WHILE;
END;
 
CALL uspPopulateAircraft();
 
-- copy data from RAW to RAW_TO_MISSION

CREATE TABLE tblMISSION(
    mission_id INTEGER AUTO_INCREMENT PRIMARY KEY,
    aircraft_id INTEGER,
    agency_id INTEGER,
    mission_date DATE,
    FOREIGN KEY(aircraft_id) REFERENCES tblAIRCRAFT(aircraft_id),
    FOREIGN KEY(agency_id) REFERENCES tblAGENCY(agency_id)
);

CREATE PROCEDURE uspGetMissionID(
    IN aircraft_id INT,
    IN agency_id INT,
    IN in_date DATE,
    OUT M_ID INT
)
BEGIN
    SET M_ID = (
        SELECT mission_id FROM tblMISSION
        WHERE aircraft_id = aircraft_id
        AND agency_id = agency_id
        AND in_date = mission_date
    );
END;
 
CREATE TABLE tblRAW_TO_MISSION(
    mission_ID INTEGER NOT NULL AUTO_INCREMENT PRIMARY KEY,
    aircraft_lat DECIMAL(9,6),
    aircraft_long DECIMAL(9,6),
    aircraft_callsign NVARCHAR(100),
    mission_date DATE,
    agency_name NVARCHAR(100),
    agency_area_code INTEGER,
    agency_phone NVARCHAR(100)
);
 
INSERT INTO tblRAW_TO_MISSION(
    aircraft_lat, aircraft_long, aircraft_callsign, mission_date,
    agency_name, agency_area_code, agency_phone
)
SELECT aircraft_lat, aircraft_long, aircraft_callsign, mission_date, agency_name, agency_area_code,
agency_phone
FROM tblRAW_WITH_PK;

-- SELECT * FROM tblRAW LIMIT 25
-- SELECT * FROM tblRAW_TO_MISSION LIMIT 25
CREATE PROCEDURE uspPopulateMissionsTable(
   IN aircraft_lat DECIMAL(9,6),
   IN aircraft_long DECIMAL(9,6),
   IN aircraft_callsign NVARCHAR(100),
   IN agency_name NVARCHAR(100),
   IN agency_area_code INTEGER,
   IN agency_phone NVARCHAR(100),
   IN mission_date DATE
)
BEGIN
    DECLARE err INT DEFAULT FALSE;
    DECLARE CONTINUE HANDLER FOR SQLEXCEPTION
    BEGIN
        SET err = true;
    END;
    -- Execute nested stored procedures
    CALL uspGetAgencyID(agency_name, agency_area_code, agency_phone, @agency_ID);
    CALL uspGetAircraftID(aircraft_callsign, @aircraft_ID);
    START TRANSACTION;
    INSERT INTO tblMISSION(agency_ID, aircraft_ID, mission_date)
    VALUES (@agency_ID, @aircraft_ID, @mission_date);
    IF err THEN
        ROLLBACK;
    ELSE
        COMMIT;
    END IF;
END;
 
-- iterate over missions in RAW and called uspPopulateMissionsTable stored procedure

CREATE PROCEDURE uspPopulateMission()
BEGIN
    DECLARE run INT;
    SET run = (SELECT COUNT(*) FROM tblRAW_TO_MISSION);
    WHILE run > 0 DO -- iterate over each mission
    
        SET @mission_ID = (SELECT MIN(mission_ID) FROM tblRAW_TO_MISSION);
        SET @aircraft_lat = (SELECT aircraft_lat FROM tblRAW_TO_MISSION WHERE mission_ID = @mission_ID);
        SET @aircraft_long = (SELECT aircraft_long FROM tblRAW_TO_MISSION WHERE mission_ID = @mission_ID);
        SET @aircraft_callsign = (SELECT aircraft_callsign FROM tblRAW_TO_MISSION WHERE mission_ID = @mission_ID);
        SET @agency_name = (SELECT agency_name FROM tblRAW_TO_MISSION WHERE mission_ID = @mission_ID);
        SET @agency_area_code = (SELECT agency_area_code FROM tblRAW_TO_MISSION WHERE mission_ID = @mission_ID);
        SET @agency_phone = (SELECT agency_phone FROM tblRAW_TO_MISSION WHERE mission_ID = @mission_ID);
        SET @mission_date = (SELECT mission_date FROM tblRAW_TO_MISSION WHERE mission_ID = @mission_ID);
        CALL uspPopulateMissionsTable(
            @aircraft_lat,
            @aircraft_long,
            @aircraft_callsign,
            @agency_name,
            @agency_area_code,
            @agency_phone,
            @mission_date
        );
            
        DELETE FROM tblRAW_TO_MISSION WHERE mission_ID = @mission_ID;
        SET run = run - 1;
    END WHILE;
END;
 
CALL uspPopulateMission();

-- Benjamin's Case Statement:
DROP PROCEDURE IF EXISTS `uspLastMission`;
CREATE PROCEDURE uspLastMission(
    IN aircraft_callsign NVARCHAR(100)
)
BEGIN
    DECLARE recent_mission_date DATE;
 
    SET recent_mission_date = (
        SELECT mission_date
        FROM tblMISSION
        JOIN tblAIRCRAFT ON tblMISSION.aircraft_id = tblAIRCRAFT.aircraft_id
        WHERE aircraft_callsign = aircraft_callsign
        ORDER BY mission_date DESC
        LIMIT 1
    );
 
    SET @delta = DATEDIFF(NOW(), recent_mission_date);
   
    CASE 
        WHEN @delta > 360 THEN "last mission was more than 1 year ago"
        WHEN @delta > 90 THEN "last mission was more than 90 days ago"
        WHEN @delta > 30 THEN "last mission was more than 30 days ago"
        WHEN @delta > 10 THEN "last mission was more than 10 days ago"
        WHEN @delta > 3 THEN "last mission was more than 3 days ago"
        ELSE "mission 3 days ago or less"
    END;
END

-- Test uspLastMission case statement with "AL1" aircraft
CALL uspLastMission("AL1");

-- Caitlin's Case Statement
SELECT tblADDRESS.address_state, tblAGENCY.agency_name,
CASE tblADDRESS.address_state
	WHEN 'Washington' THEN 'In State'
	WHEN 'Oregon' OR 'Montana' OR 'Alaska' OR 'Idaho' THEN 'In Range'
	ELSE 'Out of Range'
END AS 'Agency In Region?'

FROM tblADDRESS
INNER JOIN tblAGENCY 
ON tblAGENCY.address_ID = tblADDRESS.address_ID

-- Sabrina's Case Statement
SELECT mission_id,
CASE
    WHEN Month(mission_date) BETWEEN 3 AND 5 THEN 'SPRING'
    WHEN Month(mission_date) BETWEEN 6 AND 8 THEN 'SUMMER'
    WHEN Month(mission_date) BETWEEN 9 AND 11 THEN 'FALL'
    ELSE 'WINTER'
END
AS mission_season
FROM tblMISSION;
