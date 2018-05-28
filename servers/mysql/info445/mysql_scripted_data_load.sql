/*
alnw_info445_project_4.sql
Created: Sunday May 6, 2018
Modified:
Authors: J. Benjamin Leeds, Caitlin Schaefer, Julian Boss, Sabrina Niklaus
License: None

This script creates the tables and schema for the Airlift Northwest Elevate application database
and populates the database with 100,000 rows of data for performance testing purposes. Population
is accomplished via synthetic transaction. 

Note: Some lookup tables existed previously from the group extra credit assignment 
    and are not created and populated in this script (e.g. tblAGENCY)

Environment: 
    RDBMS: Google Cloud SQL MySQL Database 
    Location: us-west1-b zone
    CPU(s): 1 vCPU (multi-tenant CPU)
    Memory: 614 MB
    Storage: 10GB SSD
    Failover Replica: Disabled

*/

-- Part 1: Drop existing tables to begin with clean database

DROP TABLE IF EXISTS tblAGENCY_TYPE;
DROP TABLE IF EXISTS tblAIRCRAFT;
DROP TABLE IF EXISTS tblAIRCRAFT_STATUS;
DROP TABLE IF EXISTS tblASSIGNED_STATUS;
DROP TABLE IF EXISTS tblMISSION_STATUS;
DROP TABLE IF EXISTS tblASSIGNED_MISSION_STATUS;
DROP TABLE IF EXISTS tblWAYPOINT_TYPE;
DROP TABLE IF EXISTS tblWAYPOINT;
DROP TABLE IF EXISTS tblMISSION_WAYPOINT;
DROP TABLE IF EXISTS tblHOSPITAL;
DROP TABLE IF EXISTS tblAIRPORT;
DROP TABLE IF EXISTS tblPERSONNEL;
DROP TABLE IF EXISTS tblMISSION_PERSONNEL;
DROP TABLE IF EXISTS tblROLES;
DROP TABLE IF EXISTS tblASSIGNED_PERSONNEL_ROLES;
DROP TABLE IF EXISTS tblSEX;
DROP TABLE IF EXISTS tblPATIENT;
DROP TABLE IF EXISTS tblMISSION;

-- Part 2: Create tables and relationships (schema)

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
-- VALUES
--     ("Augusta", "A109E", "Rotorcraft"),
--     ("Lear", "31A", "Fixed-wing"), 
--     ("Airbus", "H-135", "Rotorcraft"),
--     ("Pilatus","PC-12","Fixed-wing");

CREATE TABLE tblAIRCRAFT(
    ac_id INTEGER AUTO_INCREMENT PRIMARY KEY,
    ac_callsign NVARCHAR(100),
    ac_n_number NVARCHAR(10),
    ac_type_id INTEGER,
    ac_lat DECIMAL(9,6),
    ac_long DECIMAL(9,6),
    ac_loc_display_name NVARCHAR(50),
    ac_cell_phone INTEGER,
    ac_sat_phone INTEGER,
    FOREIGN KEY (ac_type_id) REFERENCES tblAIRCRAFT_TYPE(aircraft_type_id)
);

-- INSERT INTO tblAIRCRAFT(ac_n_number, ac_callsign, ac_type_id, ac_lat, ac_long, ac_cell_phone, ac_sat_phone)
-- VALUES
-- ("N139AM", "AL6", (SELECT aircraft_type_id FROM tblAIRCRAFT_TYPE WHERE aircraft_type_title = "H-135"), 47.658114, 122.298400, 5551234567, 4441234567),
-- ("N951AL", "AL8", (SELECT aircraft_type_id FROM tblAIRCRAFT_TYPE WHERE aircraft_type_title = "A109E"), 47.658114, 122.298400, 5551234568, 4441234568),
-- ("N952AL", "AL5", (SELECT aircraft_type_id FROM tblAIRCRAFT_TYPE WHERE aircraft_type_title = "H-135"), 47.658114, 122.298400, 5551234569, 4441234569),
-- ("N954AL", "AL2", (SELECT aircraft_type_id FROM tblAIRCRAFT_TYPE WHERE aircraft_type_title = "H-135"), 47.658114, 122.298400, 5551234570, 4441234570),
-- ("N235UW", "AL7", (SELECT aircraft_type_id FROM tblAIRCRAFT_TYPE WHERE aircraft_type_title = "H-135"), 47.658114, 122.298400, 5551234571, 4441234571),
-- ("N212AL", "TURBOSPARE", (SELECT aircraft_type_id FROM tblAIRCRAFT_TYPE WHERE aircraft_type_title = "PC-12"), 47.658114, 122.298400, 5551234572, 4441234572),
-- ("N164AL", "LEARSPARE", (SELECT aircraft_type_id FROM tblAIRCRAFT_TYPE WHERE aircraft_type_title = "31A"), 47.658114, 122.298400, 5551234573, 4441234573),
-- ("N165AL", "JNU F", (SELECT aircraft_type_id FROM tblAIRCRAFT_TYPE WHERE aircraft_type_title = "31A"), 47.658114, 122.298400, 5551234574, 4441234574);

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

CREATE TABLE tblAIRCRAFT_STATUS(
    status_id INTEGER AUTO_INCREMENT PRIMARY KEY,
    status_title NVARCHAR(50) NOT NULL,
    status_short_desc NVARCHAR(50),
    status_long_desc NVARCHAR(300)
);

INSERT INTO tblAIRCRAFT_STATUS(status_title, status_short_desc)
VALUES
    ("VFR ONLY - Aircraft Issue", "Aircraft can only fly in Visual Meterological Conditions"),
    ("Delayed Maintenance", "Maintenance not completed on time"),
    ("Delayed Other", "Delayed due to other reasons"),
    ("Hangared", "Aircraft presently in hangar"),
    ("Heavy on Fuel", "Check aircraft fuel weight for weight and balance, payload may be limited"),
    ("Last Out", "***"),
    ("No Riders", "Essential personnel only"),
    ("USA Only", "Cross-border flights prohibited"),
    ("2 Peds", "2 Pediatric flight nurses assigned"),
    ("2 Adult", "2 Adult flight nurses assigned"),
    ("Isolette OOS", "Isolette equipment out of service"),
    ("VFR Only - Pilot Issue", "Flight crew not presently qualified for Instrument Flight Rules"),
    ("MEL - Unable to fly into known icing", "Aircraft Minimum Equipment List prevents Flight Into
    Known Icing (FIKI)");

CREATE TABLE tblASSIGNED_STATUS(
    aircraftstatus_id INTEGER AUTO_INCREMENT PRIMARY KEY,
    status_id INTEGER,
    aircraft_id INTEGER,
    assignedstatus_date TIMESTAMP,
    FOREIGN KEY(status_id) REFERENCES tblAIRCRAFT_STATUS(status_id),
    FOREIGN KEY(aircraft_id) REFERENCES tblAIRCRAFT(ac_id)
);

CREATE TABLE tblMISSION_STATUS(
    m_status_id INTEGER AUTO_INCREMENT PRIMARY KEY,
    m_status_title NVARCHAR(50) NOT NULL,
    m_status_short_desc NVARCHAR(50),
    m_status_long_desc NVARCHAR(300)
);

INSERT INTO tblMISSION_STATUS(m_status_title, m_status_short_desc)
VALUES
("Pending","Mission in planning stage"),
("Active","Mission underway"),
("Completed", "Mission completed");

CREATE TABLE tblASSIGNED_MISSION_STATUS(
    missionstatus_id INTEGER AUTO_INCREMENT PRIMARY KEY,
    mission_id INTEGER,
    m_status_id INTEGER,
    missionstatus_date TIMESTAMP,
    FOREIGN KEY(mission_id) REFERENCES tblMISSION(mission_id),
    FOREIGN KEY(m_status_id) REFERENCES tblMISSION_STATUS(m_status_id)
);

CREATE TABLE tblWAYPOINT_TYPE(
    waypointtype_id INTEGER AUTO_INCREMENT PRIMARY KEY,
    waypointtype_name NVARCHAR(50)
);

CREATE TABLE tblWAYPOINT(
    waypoint_id INTEGER AUTO_INCREMENT PRIMARY KEY,
    waypoint_title NVARCHAR(50),
    latitude DECIMAL(9,6),
    longitude DECIMAL(9,6),
    waypoint_type INTEGER,
    FOREIGN KEY (waypoint_type) REFERENCES tblWAYPOINT_TYPE(waypointtype_id)
);

CREATE TABLE tblMISSION_WAYPOINT(
    missionwaypoint_id INTEGER AUTO_INCREMENT PRIMARY KEY,
    mission_id INTEGER,
    waypoint_id INTEGER,
    mission_ETE TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    mission_ETA TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(mission_id) REFERENCES tblMISSION(mission_id),
    FOREIGN KEY(waypoint_id) REFERENCES tblWAYPOINT(waypoint_id)
);

CREATE TABLE tblHOSPITAL(
    waypoint_id INTEGER PRIMARY KEY,
    hospital_notes NVARCHAR(500),
    pad_time NVARCHAR(4),
    frequencies NVARCHAR(3),
    FOREIGN KEY (waypoint_id) REFERENCES tblWAYPOINT(waypoint_id)
);

CREATE TABLE tblAIRPORT(
    waypoint_id INTEGER PRIMARY KEY,
    airport_notes NVARCHAR(500),
    airport_identifier NVARCHAR(4),
    airport_NOTAM NVARCHAR(3),
    FOREIGN KEY (waypoint_id) REFERENCES tblWAYPOINT(waypoint_id)
);

CREATE TABLE tblPERSONNEL(
    personnel_id INTEGER AUTO_INCREMENT PRIMARY KEY,
    personnel_F_name NVARCHAR(50),
    personnel_L_name NVARCHAR(50)
);

CREATE TABLE tblMISSION_PERSONNEL(
    missionpersonnel_id INTEGER AUTO_INCREMENT PRIMARY KEY,
    mission_id INTEGER,
    personnel_id INTEGER,
    FOREIGN KEY (mission_id) REFERENCES tblMISSION(mission_id),
    FOREIGN KEY (personnel_id) REFERENCES tblPERSONNEL(personnel_id)
);

CREATE TABLE tblGROUP(
    group_id INTEGER AUTO_INCREMENT PRIMARY KEY
    group_name NVARCHAR(100) NOT NULL
)

CREATE TABLE tblPERSONNEL_GROUP(
    personnel_group_id INTEGER AUTO_INCREMENT PRIMARY KEY, 
    personnel_id INTEGER FOREIGN KEY REFERENCES tblPERSONNEL(personnel_id),
    group_id INTEGER FOREIGN KEY REFERENCES tblGROUP(group_id)
)

CREATE TABLE tblROLE(
    role_id INTEGER AUTO_INCREMENT PRIMARY KEY,
    role_title NVARCHAR(50),
    role_desc NVARCHAR(200)
);

INSERT INTO tblROLE(role_title)
VALUES 
    ("Pilot PIC"), ("Adult RN"), ("Pediatric RN");

CREATE TABLE tblASSIGNED_PERSONNEL_ROLES(
    assigned_personnel_role_id INTEGER AUTO_INCREMENT PRIMARY KEY,
    missionpersonnel_id INTEGER,
    role_id INTEGER,
    FOREIGN KEY (missionpersonnel_id) REFERENCES tblMISSION_PERSONNEL(missionpersonnel_id),
    FOREIGN KEY (role_id) REFERENCES tblROLES(role_id)
);

CREATE TABLE tblSEX(
    sex_id INTEGER AUTO_INCREMENT PRIMARY KEY,
    sex_name NVARCHAR(50)
);

CREATE TABLE tblPATIENT(
    mission_id INTEGER PRIMARY KEY,
    patient_sex INTEGER,
    patient_short_report NVARCHAR(500),
    patient_intubated BOOLEAN,
    patient_drips TINYINT,
    patient_age TINYINT,
    patient_weight SMALLINT,
    patient_cardiac BOOLEAN,
    patient_gi_bleed BOOLEAN,
    patient_OB BOOLEAN,
    FOREIGN KEY (mission_id) REFERENCES tblMISSION(mission_id),
    FOREIGN KEY (patient_sex) REFERENCES tblSEX(sex_id)
);

CREATE TABLE tblAGENCY_TYPE(
    agencytype_id INTEGER AUTO_INCREMENT PRIMARY KEY,
    agencytype_name NVARCHAR(50),
    agencytype_desc NVARCHAR(200)
);





-- Part 3: Populate dummy data

CREATE PROCEDURE uspPopulateMissionWrapper(
    IN run INTEGER -- number of rows to populate
)
BEGIN 
SET @num_aircraft = (SELECT COUNT(*) FROM tblAIRCRAFT);
SET @num_agency = (SELECT COUNT(*) FROM tblAGENCY);

    WHILE run > 0 DO 
        DECLARE err INT DEFAULT FALSE;
        SET @aircraft = (SELECT FLOOR(1 + (RAND() * @num_aircraft))); -- choose rand aircraft
        SET @req = (SELECT FLOOR(1 + (RAND() * @num_agency))); -- choose rand requestor
        SET @rec = (SELECT FLOOR(1 + (RAND() * @num_agency))); -- choose rand received
        SET @mdate = (SELECT FROM_UNIXTIME(RAND() * (select UNIX_TIMESTAMP()))); -- choose rand date
        /* should be between 18-000000 and 18-100000 8*/ 
        SET @tcnum = select concat('18-', (LPAD(run, 6, '0'))); -- choose rand mission id

        START TRANSACTION;
            INSERT INTO tblMISSION(aircraft_id, requestor, receiver, mission_date, tc_number)
            VALUES (@aircraft, @req, @rec, @mdate, @tcnum);
            IF err THEN
                ROLLBACK;
            ELSE
                COMMIT;
            END IF;
            SET err = FALSE;
            SET run = run - 1
    END WHILE;
END;

CREATE PROCEDURE uspPopulatePatientWrapper(
    IN run INTEGER
)

BEGIN
SET @num_patient_gender = (SELECT COUNT(*) FROM tblGENDER);

WHILE run > 0 DO
	SET @mission_id = run;
	SET @patient_sex = SELECT FLOOR(1+(RAND()* @num_patient_gender)); -- choose rand sex
	SET @patient_short_report = ‘Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis feugiat arcu sapien, a euismod elit rhoncus vel. Duis semper tincidunt ultricies. Fusce consequat.’;
	SET @patient_intubated = SELECT ROUND(RAND());
	SET @patient_drips = SELECT FLOOR(RAND()*6); -- choose rand number of patient IVs
	SET @patient_age = SELECT FLOOR(RAND()*(90-2+1))+2;
	SET @patient_weight = SELECT FLOOR(RAND()*(255-30+1))+30;
	SET @patient_cardiac = SELECT ROUND(RAND()); -- choose whether patient is a cardiac patient
	SET @patient_gi_bleed = SELECT ROUND(RAND()); -- choose whether patient has gastrointestinal bleed
	SET @patient_ob = SELECT ROUND(RAND()); -- choose if patient is an obstetrics patient
	
	START TRANSACTION;
        INSERT INTO tblPatient(mission_id, patient_sex, patient_short_report, patient_intubated, patient_drips, patient_age, patient_weight, patient_cardiac, patient_gi_bleed, patient_OB)
        VALUES (@mission_id, @patient_sex, @patient_short_report, @patient_intubated, @patient_drips, @patient_age, @patient_weight, @patient_cardiac, @patient_gi_bleed, @patient_ob); 
        IF err THEN
            ROLLBACK;
        ELSE
            COMMIT;
        END IF;
        SET err = FALSE;
        SET run = run - 1
END WHILE;
END;
