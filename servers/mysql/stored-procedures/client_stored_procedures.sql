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

-- endpoint: /v1/personnel
DROP PROCEDURE IF EXISTS `uspGetAllPeople()`;
CREATE PROCEDURE uspGetAllPeople()
BEGIN
    SELECT personnel_id, personnel_f_name, personnel_l_name, personnel_title
    FROM tblPERSONNEL;
END;

-- endpoint: /v1/people?q=searchquery
DROP PROCEDURE IF EXISTS `uspGetPersonByID`;
CREATE PROCEDURE uspGetPersonByID(
    IN pid INTEGER
)
BEGIN
    SELECT personnel_id, personnel_f_name, personnel_l_name, personnel_title
    FROM tblPERSONNEL
    WHERE personnel_id = pid;
END;

-- endpoint: /v1/people/{id}
DROP PROCEDURE IF EXISTS `uspGetPersonDetailByID`;
CREATE PROCEDURE uspGetPersonDetailByID(
    IN pid INTEGER
)
BEGIN
    SELECT personnel_id, personnel_f_name, personnel_l_name,
    personnel_title, personnel_sms_num, personnel_email
    FROM tblPERSONNEL
    WHERE personnel_id = pid;
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

-- *************************************************

-- endpoint: /v1/aircraft/???
-- called when q="", status comes from query status field?
-- CALL uspGetAircraftByStatus("Out of Service");
DROP PROCEDURE IF EXISTS `uspGetAircraftByStatus`;
CREATE PROCEDURE uspGetAircraftByStatus(
    IN statusTitleQuery NVARCHAR(25)
)
BEGIN
    SELECT ac_id, ac_callsign, ac_n_number, aircraft_type_manufacturer, aircraft_type_title,
    aircraft_type_category, ac_lat, ac_long, ac_loc_display_name, status_title
    FROM tblAIRCRAFT
    JOIN tblASSIGNED_STATUS ON tblAIRCRAFT.ac_id = tblASSIGNED_STATUS.aircraft_id
    JOIN tblAIRCRAFT_STATUS ON tblASSIGNED_STATUS.status_id = tblAIRCRAFT_STATUS.status_id
    JOIN tblAIRCRAFT_TYPE ON tblAIRCRAFT.ac_type_id = tblAIRCRAFT_TYPE.aircraft_type_id
    WHERE tblAIRCRAFT_STATUS.status_title = statusTitleQuery;
END;

-- TODO: account for multiple status entries per aircraft:
-- get aircraft by status
-- NOT COMMITTED
DROP PROCEDURE IF EXISTS `uspGetAircraftByStatus`;
CREATE PROCEDURE uspGetAircraftByStatus(
    IN status_query NVARCHAR(25)
)
BEGIN
    SELECT ac_id, ac_callsign, ac_n_number, status_title,
    status_long_desc, status_short_desc, assignedstatus_date
    FROM tblAIRCRAFT test
    INNER JOIN tblASSIGNED_STATUS ON tblAIRCRAFT.ac_id = tblASSIGNED_STATUS.aircraft_id
    INNER JOIN tblAIRCRAFT_STATUS ON tblASSIGNED_STATUS.status_id = tblAIRCRAFT_STATUS.status_id
    WHERE test.assignedstatus_date = (
        SELECT MAX(test2.assignedstatus_date)
        FROM tblAIRCRAFT test2
        INNER JOIN tblASSIGNED_STATUS ON tblAIRCRAFT.ac_id = tblASSIGNED_STATUS.aircraft_id
        INNER JOIN tblAIRCRAFT_STATUS ON tblASSIGNED_STATUS.status_id = tblAIRCRAFT_STATUS.status_id
        WHERE test2.aircraft_id = test.aircraft_id
        );
    
    SELECT MAX(test.assignedstatus_date)
    FROM tblASSIGNED_STATUS test;
END;

-- *************************************************

-- endpoint: /v1/aircraft/{id}
-- CALL uspGetAircraftByID(7);
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

DROP PROCEDURE IF EXISTS `uspGetMissionByAircraft`;
CREATE PROCEDURE uspGetMissionByAircraft(
    IN aid INTEGER
)
BEGIN

    SELECT mission_type_id, tc_number
    FROM tblMISSION
    WHERE tblMISSION.aircraft_id = 5
    ORDER BY tblMISSION.mission_date ASC

    SELECT *
    FROM tblAIRCRAFT
    WHERE ac_id=2;

    SELECT * FROM 

END;

-- endpoint: /v1/aircraft/?? NOT COMMITTED
DROP PROCEDURE IF EXISTS `uspGetOOSDetailByAircraft`;
CREATE PROCEDURE uspGetOOSDetailByAircraft(
    IN aid INTEGER
)
BEGIN

END;

SELECT * FROM tblAIRCRAFT
CALL uspGetA
SELECT * FROM tblAIRCRAFT_SCHED_SERVICE;

-- endpoint: /v1/resources
DROP PROCEDURE IF EXISTS `uspGetResources`;
CREATE PROCEDURE uspGetResources()
BEGIN
    SELECT resource_link_id, resource_short_name, resource_long_name, resource_url,
    resource_thumbnail_photo_url FROM tblRESOURCE_LINKS;
END;

-- active missions:
-- return missions with "Pending" or "In-progess" status
DROP PROCEDURE IF EXISTS `uspGetActiveMissions`;
CREATE PROCEDURE uspGetActiveMissions()
BEGIN
    SELECT tblMISSION.mission_id FROM tblMISSION
    JOIN tblASSIGNED_MISSION_STATUS ON tblMISSION.mission_id = tblASSIGNED_MISSION_STATUS.mission_id
    WHERE tblASSIGNED_MISSION_STATUS.m_status_id IN (
        SELECT m_status_id
        FROM tblMISSION_STATUS
        WHERE m_status_title IN ("Pending", 'In-progress')
    );
END;

-- endpoint: /v1/????
-- returns: mission_type, flight rules, TC, req, rec
-- CALL uspGetMissionDetailByAircraft(7);
DROP PROCEDURE IF EXISTS `uspGetMissionDetailByAircraft`;
CREATE PROCEDURE uspGetMissionDetailByAircraft(
    IN aid INTEGER
)
BEGIN
    SELECT mission_id, mission_type_id, tc_number, requestor_id, receiver_id
    FROM tblMISSION
    WHERE tblMISSION.aircraft_id = aid;
    -- WARNING: TODO: return only most recent mission per aircraft
END;

-- get patient by aircraft
-- Questions: can mission be provided or can only aircraft be provided?
DROP PROCEDURE IF EXISTS `uspGetPatientByAircraft`;
CREATE PROCEDURE uspGetPatientByAircraft(
    IN aircraft_id NVARCHAR(25)
)
BEGIN
    -- how to save list as local var for where clause?
    SET @active_missions_list = (CALL uspGetActiveMissions());
    
    SELECT *
    FROM tblPATIENT
    WHERE mission_id IN (CALL uspGetActiveMissions());

END;


