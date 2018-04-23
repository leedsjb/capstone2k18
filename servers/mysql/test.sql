-- sandbox to interact with Google Cloud SQL MySQL database

CREATE DATABASE TEST

CREATE TABLE tblTEST(
    testField1 INT,
    testField2 NVARCHAR(100)
);


CREATE TABLE tblTEST_IMPORT(
    aircraft NVARCHAR(100),
    aircraft_model NVARCHAR(100),
    mission DATE,
    agency NVARCHAR(100),
    agency_address NVARCHAR(100)
);

INSERT INTO tblTEST_IMPORT(aircraft_model) 
VALUES ('learjet 31A')

SELECT DISTINCT(model_title) FROM tblRAW

CREATE TABLE tblRAW(

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
)


SELECT DISTINCT aircraft_callsign, model_title, model_desc, aircraft_category FROM tblRAW

SELECT mission_date FROM tblRAW LIMIT 100

DELETE FROM tblRAW LIMIT 1

CREATE TABLE tblTEST_DATA_TYPES(
    integer_test INTEGER,
    int_test INT,
    nvarchar_test NVARCHAR(100),
    date_test DATE,
    decimal_test DECIMAL(9,6)
)

SELECT DISTINCT aircraft_callsign, model_title, model_desc, aircraft_category FROM tblRAW LIMIT 50

SELECT aircraft_long FROM tblRAW WHERE aircraft_long=0

DELETE FROM tblRAW LIMIT 1;

SELECT COUNT(
    DISTINCT
        agency_name, agency_phone
) FROM tblRAW

SELECT COUNT(DISTINCT address_street) FROM tblRAW

SELECT agency_name, agency_phone FROM tblRAW LIMIT 100

SELECT COUNT(
    DISTINCT
        address_street, address_city
    
) FROM tblRAW

SHOW TABLES

SELECT IF(
    (
        SELECT COUNT(*)
        FROM information_schema.tables 
        WHERE table_schema = 'TEST'
        AND table_name = 'tblTEST' = 1
    ),
    (DROP TABLE tblTEST),
    ("DOESN'T EXIST")
)


DROP TABLE IF EXISTS `tblTEST`;
CREATE TABLE tblTEST(
    ID INTEGER AUTO_INCREMENT PRIMARY KEY
);

SELECT model_title FROM tblRAW LIMIT 100

SELECT * FROM tblRAW_WITH_PK LIMIT 100

SELECT * FROM tblMISSION LIMIT 50

SELECT * FROM tblRAW LIMIT 50
SELECT * FROM tblAGENCY LIMIT 50
SELECT * FROM tblAIRCRAFT LIMIT 50
SELECT * FROM tblMISSION


SELECT * FROM tblMISSION
JOIN (tblAIRCRAFT, tblAGENCY)
ON (
    tblMISSION.aircraft_id = tblAIRCRAFT.aircraft_id
    AND tblMISSION.agency_id = tblAGENCY.agency_id
)
LIMIT 25 

SELECT * FROM tblMISSION