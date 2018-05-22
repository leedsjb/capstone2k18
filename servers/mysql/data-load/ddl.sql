/*
    ddl.sql
    Created: Thursday May 17, 2018
    Modified: Monday May 21, 2018
    Authors: J. Benjamin Leeds
    License: None

    This script creates the tables and schema for the
    Airlift Northwest Elevate application database.
*/

DROP TABLE IF EXISTS `tblHOSPITAL`;
DROP TABLE IF EXISTS `tblAIRPORT`;
DROP TABLE IF EXISTS `tblMISSION_WAYPOINT`;
DROP TABLE IF EXISTS `tblWAYPOINT`;
DROP TABLE IF EXISTS `tblWAYPOINT_TYPE`;
CREATE TABLE tblWAYPOINT_TYPE (
    waypointtype_id INTEGER,
    waypointtype_name NVARCHAR(50),
    PRIMARY KEY (waypointtype_id)
);

CREATE TABLE `tblWAYPOINT` (
    `waypoint_id` INTEGER,
    `waypoint_title` NVARCHAR(50),
    `latitude` DECIMAL(9,6),
    `longitude` DECIMAL(9,6),
    `waypoint_type` INTEGER,
    PRIMARY KEY (`waypoint_id`),
    FOREIGN KEY (`waypoint_type`) REFERENCES tblWAYPOINT_TYPE(waypointtype_id)
);

CREATE TABLE `tblHOSPITAL` (
    `waypoint_id` INTEGER,
    `hospital_notes` NVARCHAR(500),
    `pad_time` NVARCHAR(4),
    `frequencies` NVARCHAR(3),
    PRIMARY KEY (`waypoint_id`),
    FOREIGN KEY (waypoint_id) REFERENCES tblWAYPOINT(waypoint_id)
);

CREATE TABLE `tblAIRPORT` (
    `waypoint_id` INTEGER,
    `airport_notes` NVARCHAR(500),
    `airport_identifier` NVARCHAR(5),
    `airport_NOTAM` NVARCHAR(500),
    PRIMARY KEY(waypoint_id),
    FOREIGN KEY(waypoint_id) REFERENCES tblWAYPOINT(waypoint_id)
);

DROP TABLE IF EXISTS `tblASSIGNED_STATUS`;
DROP TABLE IF EXISTS `tblAIRCRAFT_STATUS`;
DROP TABLE IF EXISTS `tblAIRCRAFT_SCHED_SERVICE`;
DROP TABLE IF EXISTS `tblPATIENT`;
DROP TABLE IF EXISTS `tblMISSION_PERSONNEL`;
DROP TABLE IF EXISTS `tblASSIGNED_MISSION_STATUS`;
DROP TABLE IF EXISTS `tblMISSION`;
DROP TABLE IF EXISTS `tblMISSION_TYPE`;
DROP TABLE IF EXISTS `tblAIRCRAFT`;
DROP TABLE IF EXISTS `tblAGENCY`;
DROP TABLE IF EXISTS `tblADDRESS`;
CREATE TABLE `tblADDRESS` (
    `address_id` INTEGER,
    `address_street` NVARCHAR(100),
    `address_street_1` NVARCHAR(100),
    `address_city` NVARCHAR(50),
    `address_state` NVARCHAR(20),
    `address_zip` INTEGER,
    `address_zip_plus4` INTEGER,
    PRIMARY KEY (`address_id`)
);

DROP TABLE IF EXISTS `tblAGENCY_TYPE`;
CREATE TABLE `tblAGENCY_TYPE` (
    `agencytype_id` INTEGER,
    `agencytype_name` NVARCHAR(50),
    `agencytype_desc` NVARCHAR(200),
    PRIMARY KEY (`agencytype_id`)
);

CREATE TABLE `tblAGENCY` (
    `agency_id` INTEGER,
    `agency_name` NVARCHAR(50),
    `agency_area_code` INTEGER,
    `agency_phone` INTEGER,
    `address_id` INTEGER,
    `agencytype_id` INTEGER,
    PRIMARY KEY (`agency_id`),
    FOREIGN KEY(address_id) REFERENCES tblADDRESS(address_id),
    FOREIGN KEY(agencytype_id) REFERENCES tblAGENCY_TYPE(agencytype_id)
);

CREATE TABLE `tblAIRCRAFT_STATUS` (
    `status_id` INTEGER,
    `status_title` NVARCHAR(50),
    `status_long_desc` NVARCHAR(300),
    `status_short_desc` NVARCHAR(50),
    PRIMARY KEY (`status_id`)
);

DROP TABLE IF EXISTS `tblAIRCRAFT_TYPE`;
CREATE TABLE `tblAIRCRAFT_TYPE` (
    `aircraft_type_id` INTEGER,
    `aircraft_type_title` NVARCHAR(50),
    `aircraft_type_desc` NVARCHAR(250),
    `aircraft_type_category` NVARCHAR(25),
    `aircraft_type_manufacturer` NVARCHAR(50),
    PRIMARY KEY (`aircraft_type_id`)
);

CREATE TABLE `tblAIRCRAFT` (
  `ac_id` INTEGER,
  `ac_callsign` NVARCHAR(50),
  `ac_n_number` NVARCHAR(10),
  `ac_type_id` INTEGER,
  `ac_lat` DECIMAL(9,6),
  `ac_long` DECIMAL(9,6),
  `ac_loc_display_name` NVARCHAR(50),
  `ac_cell_phone` BIGINT,
  `ac_sat_phone` BIGINT,
  PRIMARY KEY (`ac_id`),
  FOREIGN KEY (ac_type_id) REFERENCES tblAIRCRAFT_TYPE(aircraft_type_id)
);

CREATE TABLE `tblASSIGNED_STATUS` (
    `aircraftstatus_id` int,
    `status_id` int,
    `aircraft_id` int,
    `assignedstatus_date` TIMESTAMP,
    PRIMARY KEY (`aircraftstatus_id`),
    FOREIGN KEY(status_id) REFERENCES tblAIRCRAFT_STATUS(status_id),
    FOREIGN KEY(aircraft_id) REFERENCES tblAIRCRAFT(ac_id)
);

CREATE TABLE `tblAIRCRAFT_SCHED_SERVICE` (
  `acSchedServiceID` INTEGER,
  `ac_id` INTEGER,
  `acSchedServiceReason` NVARCHAR(100),
  `OOSStartTime` DATETIME,
  `OOSEndTime` DATETIME,
  PRIMARY KEY (`acSchedServiceID`),
  FOREIGN KEY(ac_id) REFERENCES tblAIRCRAFT(ac_id)
);

CREATE TABLE `tblMISSION_TYPE`(
    `mission_type_id` INTEGER NOT NULL PRIMARY KEY,
    `mission_type_short_name` NVARCHAR(50)
);

CREATE TABLE `tblMISSION`(
    `mission_id` INTEGER,
    `aircraft_id` INTEGER,
    `mission_type_id` INTEGER,
    `requestor_id` INTEGER,
    `receiver_id` INTEGER,
    `mission_date` TIMESTAMP,
    `tc_number` VARCHAR(10),
    PRIMARY KEY (`mission_id`),
    FOREIGN KEY(aircraft_id) REFERENCES tblAIRCRAFT(ac_id),
    FOREIGN KEY(mission_type_id) REFERENCES tblMISSION_TYPE(mission_type_id),
    FOREIGN KEY(requestor_id) REFERENCES tblAGENCY(agency_id),
    FOREIGN KEY(receiver_id) REFERENCES tblAGENCY(agency_id)
);

DROP TABLE IF EXISTS `tblMISSION_STATUS`;
CREATE TABLE `tblMISSION_STATUS` (
    `m_status_id` INTEGER,
    `m_status_title` NVARCHAR(50),
    `m_status_long_desc` NVARCHAR(300),
    `m_status_short_desc` NVARCHAR(50),
    PRIMARY KEY (`m_status_id`)
);

CREATE TABLE `tblASSIGNED_MISSION_STATUS` (
    `missionstatus_id` int,
    `mission_id` int,
    `m_status_id` int,
    `missionstatus_date` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`missionstatus_id`),
    FOREIGN KEY(mission_id) REFERENCES tblMISSION(mission_id),
    FOREIGN KEY(m_status_id) REFERENCES tblMISSION_STATUS(m_status_id)
);

CREATE TABLE `tblMISSION_WAYPOINT` (
    `missionwaypoint_id` INTEGER,
    `mission_id` INTEGER,
    `waypoint_id` INTEGER,
    -- `mission_ETE` TIMESTAMP, -- calculated value: ETA - NOW()
    `mission_ETA` TIMESTAMP NULL, -- do not want auto initialization here
    PRIMARY KEY (`missionwaypoint_id`),
    FOREIGN KEY(mission_id) REFERENCES tblMISSION(mission_id),
    FOREIGN KEY(waypoint_id) REFERENCES tblWAYPOINT(waypoint_id)
);

DROP TABLE IF EXISTS tblGENDER;
CREATE TABLE `tblGENDER` (
    `GENDER_id` INTEGER,
    `GENDER_name` NVARCHAR(50),
    PRIMARY KEY (`GENDER_id`)
);

CREATE TABLE `tblPATIENT` (
    `mission_id` INTEGER,
    `patient_GENDER` INTEGER,
    `patient_short_report` NVARCHAR(500),
    `patient_intubated` BOOLEAN,
    `patient_drips` TINYINT,
    `patient_age` TINYINT,
    `patient_weight` SMALLINT,
    `patient_cardiac` BOOLEAN,
    `patient_gi_bleed` BOOLEAN,
    `patient_OB` BOOLEAN,
    PRIMARY KEY(mission_id),
    FOREIGN KEY(mission_id) REFERENCES tblMISSION(mission_id),
    FOREIGN KEY(patient_GENDER) REFERENCES tblGENDER(GENDER_id)
);

DROP TABLE IF EXISTS tblPERSONNEL_CREW_TYPE;
DROP TABLE IF EXISTS tblCREW_TYPE;
DROP TABLE IF EXISTS tblPERSONNEL_GROUP;
DROP TABLE IF EXISTS tblPERSONNEL;
CREATE TABLE `tblPERSONNEL` (
    `personnel_id` INTEGER,
    `personnel_f_name` NVARCHAR(50),
    `personnel_l_name` NVARCHAR(50),
    `personnel_title` NVARCHAR(50),
    `personnel_sms_num` NVARCHAR(50),
    `personnel_email` NVARCHAR(50),
    PRIMARY KEY (`personnel_id`)
);

/*
    Related Flight Vector Table: CrewType
    T-SQL Query: SELECT ID, Name, Role FROM CrewType
*/
CREATE TABLE `tblCREW_TYPE` (
    `crew_type_id` INTEGER,
    `crew_type_name` NVARCHAR(25),
    `crew_type_role` NVARCHAR(25),
    PRIMARY KEY (`crew_type_id`)
);

DROP TABLE IF EXISTS tblGROUP;
CREATE TABLE `tblGROUP` (
    `group_id` INTEGER,
    `group_name` NVARCHAR(100) NOT NULL,
    PRIMARY KEY (`group_id`)
);

CREATE TABLE `tblPERSONNEL_CREW_TYPE` (
    `personnel_crew_type_id` INTEGER,
    `personnel_id` INTEGER,
    `crew_type_id` INTEGER,
    PRIMARY KEY (`personnel_crew_type_id`),
    FOREIGN KEY(personnel_id) REFERENCES tblPERSONNEL(personnel_id),
    FOREIGN KEY(crew_type_id) REFERENCES tblCREW_TYPE(crew_type_id)
);

CREATE TABLE `tblPERSONNEL_GROUP` (
    `personnelGroupID` INTEGER,
    `personnel_id` INTEGER,
    `group_id` INTEGER,
    PRIMARY KEY (`personnelGroupID`),
    FOREIGN KEY(personnel_id) REFERENCES tblPERSONNEL(personnel_id),
    FOREIGN KEY(group_id) REFERENCES tblGROUP(group_id)
);

CREATE TABLE `tblMISSION_PERSONNEL` (
    `missionpersonnel_id` INTEGER,
    `mission_id` INTEGER,
    `personnel_crew_type_id` INTEGER,
    PRIMARY KEY (`missionpersonnel_id`),
    FOREIGN KEY(mission_id) REFERENCES tblMISSION(mission_id),
    FOREIGN KEY(personnel_crew_type_id) REFERENCES tblPERSONNEL_CREW_TYPE(personnel_crew_type_id)
);

DROP TABLE IF EXISTS `tblRESOURCE_LINKS`;
CREATE TABLE `tblRESOURCE_LINKS` (
    `resource_link_id` INTEGER AUTO_INCREMENT PRIMARY KEY,
    `resource_short_name` NVARCHAR(25) NOT NULL,
    `resource_long_name` NVARCHAR(100) NOT NULL,
    `resource_url` NVARCHAR(100) NOT NULL, 
    `resource_thumbnail_photo_url` NVARCHAR(200) NOT NULL
);

SHOW TABLES;
