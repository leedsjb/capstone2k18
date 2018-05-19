/*
client_stored_procedures.sql
Created: Thursday May 17, 2018
Modified: Saturday May 19, 2018
Author(s): J. Benjamin Leeds
License: None

Use the stored procedures in this file to retrieve data in MySQL to send to clients

*/

-- return list of all groups
-- endpoint: /v1/groups
DROP PROCEDURE IF EXISTS `uspGetAllGroups`;
CREATE PROCEDURE uspGetAllGroups()
BEGIN
    SELECT tblGROUP.group_id, group_name, personnel_f_name, personnel_l_name
    FROM tblGROUP
    JOIN tblPERSONNEL_GROUP ON tblGROUP.group_id = tblPERSONNEL_GROUP.group_id
    JOIN tblPERSONNEL ON tblPERSONNEL_GROUP.personnel_id = tblPERSONNEL.personnel_id
    ORDER BY tblGROUP.group_name ASC;
END;

-- used to populate group search trie
-- endpoint: none
DROP PROCEDURE IF EXISTS `uspGetGroupByID`;
CREATE PROCEDURE uspGetGroupByID(
    IN gid INTEGER
)
BEGIN
    SELECT tblGROUP.group_id, group_name, personnel_f_name, personnel_l_name
    FROM tblGROUP
    JOIN tblPERSONNEL_GROUP ON tblGROUP.group_id = tblPERSONNEL_GROUP.group_id
    JOIN tblPERSONNEL ON tblPERSONNEL_GROUP.personnel_id = tblPERSONNEL.personnel_id 
    WHERE tblGROUP.group_id = gid;
END;

-- endpoint: /v1/groups/{id}
DROP PROCEDURE IF EXISTS `uspGetGroupDetailByID`;
CREATE PROCEDURE uspGetGroupDetailByID(
    IN gid INTEGER
)
BEGIN
    SELECT tblGROUP.group_id, group_name, personnel_f_name,
    personnel_l_name, tblPERSONNEL.personnel_id, personnel_title
    FROM tblGROUP
    JOIN tblPERSONNEL_GROUP ON tblGROUP.group_id = tblPERSONNEL_GROUP.group_id
    JOIN tblPERSONNEL ON tblPERSONNEL_GROUP.personnel_id = tblPERSONNEL.personnel_id 
    WHERE tblGROUP.group_id = gid;
END;

-- endpoint: /v1/aircraft
DROP PROCEDURE IF EXISTS `uspGetAllAircraft`;
CREATE PROCEDURE uspGetAllAircraft()
BEGIN
    SELECT ac_id, ac_callsign, ac_n_number, aircraft_type_manufacturer, aircraft_type_title,
    aircraft_type_category, ac_lat, ac_long, ac_loc_display_name, status_title
    FROM tblAIRCRAFT
    JOIN tblAIRCRAFT_TYPE ON tblAIRCRAFT.ac_type_id = tblAIRCRAFT_TYPE.aircraft_type_id
    JOIN tblASSIGNED_STATUS ON tblAIRCRAFT.ac_id = tblASSIGNED_STATUS.aircraft_id
    JOIN tblAIRCRAFT_STATUS ON tblASSIGNED_STATUS.status_id = tblAIRCRAFT_STATUS.status_id;
END;

-- endpoint: /v1/aircraft/{id}
DROP PROCEDURE IF EXISTS `uspGetAircraftByID`;
CREATE PROCEDURE uspGetAircraftByID(
    IN aid INTEGER
)
BEGIN
 SELECT ac_id, ac_callsign, ac_n_number, aircraft_type_manufacturer, aircraft_type_title,
    aircraft_type_category, ac_lat, ac_long, ac_loc_display_name, status_title
    FROM tblAIRCRAFT
    JOIN tblAIRCRAFT_TYPE ON tblAIRCRAFT.ac_type_id = tblAIRCRAFT_TYPE.aircraft_type_id
    JOIN tblASSIGNED_STATUS ON tblAIRCRAFT.ac_id = tblASSIGNED_STATUS.aircraft_id
    JOIN tblAIRCRAFT_STATUS ON tblASSIGNED_STATUS.status_id = tblAIRCRAFT_STATUS.status_id
    WHERE tblAIRCRAFT.ac_id = aid;
END;