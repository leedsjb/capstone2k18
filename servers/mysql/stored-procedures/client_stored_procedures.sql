/*
client_stored_procedures.sql
Created: Thursday May 17, 2018
Modified: Saturday May 26, 2018
Last Change: add uspGetAgencyDetailsByID() stored procedure
update mission_waypoints to add "waypoint_completed"
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
-- CALL uspGetGroupDetailByID(20);
DROP PROCEDURE IF EXISTS `uspGetGroupDetailByID`;
CREATE PROCEDURE uspGetGroupDetailByID(
    IN gid INTEGER
)
BEGIN
    SELECT tblGROUP.group_id, group_name, tblPERSONNEL.personnel_id, personnel_f_name,
    personnel_l_name, personnel_title, personnel_sms_num, personnel_email
    FROM tblGROUP
    JOIN tblPERSONNEL_GROUP ON tblGROUP.group_id = tblPERSONNEL_GROUP.group_id
    JOIN tblPERSONNEL ON tblPERSONNEL_GROUP.personnel_id = tblPERSONNEL.personnel_id 
    WHERE tblGROUP.group_id = gid;
END;

-- endpoint: /v1/personnel
DROP PROCEDURE IF EXISTS `uspGetAllPeople`;
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
-- CALL uspGetPersonDetailByID(5);
DROP PROCEDURE IF EXISTS `uspGetPersonDetailByID`;
CREATE PROCEDURE uspGetPersonDetailByID(
    IN pid INTEGER
)
BEGIN
    SELECT tblPERSONNEL.personnel_id, personnel_f_name, personnel_l_name,
    personnel_title, personnel_sms_num, personnel_email, tblGROUP.group_id, tblGROUP.group_name
    FROM tblPERSONNEL
    LEFT JOIN tblPERSONNEL_GROUP ON tblPERSONNEL.personnel_id=tblPERSONNEL_GROUP.personnel_id
    LEFT JOIN tblGROUP ON tblPERSONNEL_GROUP.group_id = tblGROUP.group_id
    WHERE tblPERSONNEL.personnel_id = pid;
END;


-- endpoint: /v1/aircraft
-- CALL uspGetAllAircraft()
DROP PROCEDURE IF EXISTS `uspGetAllAircraft`;
CREATE PROCEDURE uspGetAllAircraft()
BEGIN
    SELECT ac_id, ac_callsign, ac_n_number, aircraft_type_manufacturer, aircraft_type_title,
    aircraft_type_category, ac_lat, ac_long, ac_loc_display_name, status_short_desc
    FROM tblAIRCRAFT
    JOIN tblAIRCRAFT_TYPE ON tblAIRCRAFT.ac_type_id = tblAIRCRAFT_TYPE.aircraft_type_id
    JOIN tblASSIGNED_STATUS ON tblAIRCRAFT.ac_id = tblASSIGNED_STATUS.aircraft_id
    JOIN tblAIRCRAFT_STATUS ON tblASSIGNED_STATUS.status_id = tblAIRCRAFT_STATUS.status_id;
END;

-- endpoint: /v1/aircraft?status={OAM,RFM,OOS}
-- called when q="", status comes from query status field?
-- CALL uspGetAircraftByStatus("Out of Service");
DROP PROCEDURE IF EXISTS `uspGetAircraftByStatus`;
CREATE PROCEDURE uspGetAircraftByStatus(
    IN statusTitleQuery NVARCHAR(25)
)
BEGIN
    SELECT ac_id, ac_callsign, ac_n_number, aircraft_type_manufacturer, aircraft_type_title,
    aircraft_type_category, ac_lat, ac_long, ac_loc_display_name, status_short_desc
    FROM tblAIRCRAFT
    JOIN tblASSIGNED_STATUS ON tblAIRCRAFT.ac_id = tblASSIGNED_STATUS.aircraft_id
    JOIN tblAIRCRAFT_STATUS ON tblASSIGNED_STATUS.status_id = tblAIRCRAFT_STATUS.status_id
    JOIN tblAIRCRAFT_TYPE ON tblAIRCRAFT.ac_type_id = tblAIRCRAFT_TYPE.aircraft_type_id
    WHERE tblAIRCRAFT_STATUS.status_short_desc = statusTitleQuery;
END;

-- endpoint: /v1/aircraft?category="{Rotorcraft, Fixed-wing}"
-- CALL uspGetAircraftByCategory("Fixed-wing");
DROP PROCEDURE IF EXISTS `uspGetAircraftByCategory`;
CREATE PROCEDURE uspGetAircraftByCategory(
    IN categoryQuery NVARCHAR(25)
)
BEGIN
    SELECT ac_id, ac_callsign, ac_n_number, aircraft_type_manufacturer, aircraft_type_title,
    aircraft_type_category, ac_lat, ac_long, ac_loc_display_name, status_short_desc
    FROM tblAIRCRAFT
    JOIN tblASSIGNED_STATUS ON tblAIRCRAFT.ac_id = tblASSIGNED_STATUS.aircraft_id
    JOIN tblAIRCRAFT_STATUS ON tblASSIGNED_STATUS.status_id = tblAIRCRAFT_STATUS.status_id
    JOIN tblAIRCRAFT_TYPE ON tblAIRCRAFT.ac_type_id = tblAIRCRAFT_TYPE.aircraft_type_id
    WHERE tblAIRCRAFT_TYPE.aircraft_type_category = categoryQuery;
END;

-- endpoint: /v1/aircraft?category="{Rotorcraft, Fixed-wing}"&status={OAM,RFM,OOS}
-- CALL uspGetAircraftByStatusAndCategory("OOS", "Rotorcraft")
DROP PROCEDURE IF EXISTS `uspGetAircraftByStatusAndCategory`;
CREATE PROCEDURE uspGetAircraftByStatusAndCategory(
    IN statusTitleQuery NVARCHAR(25),
    IN categoryQuery NVARCHAR(25)
)
BEGIN
    SELECT ac_id, ac_callsign, ac_n_number, aircraft_type_manufacturer, aircraft_type_title,
    aircraft_type_category, ac_lat, ac_long, ac_loc_display_name, status_short_desc
    FROM tblAIRCRAFT
    JOIN tblASSIGNED_STATUS ON tblAIRCRAFT.ac_id = tblASSIGNED_STATUS.aircraft_id
    JOIN tblAIRCRAFT_STATUS ON tblASSIGNED_STATUS.status_id = tblAIRCRAFT_STATUS.status_id
    JOIN tblAIRCRAFT_TYPE ON tblAIRCRAFT.ac_type_id = tblAIRCRAFT_TYPE.aircraft_type_id
    WHERE tblAIRCRAFT_STATUS.status_short_desc = statusTitleQuery
    AND tblAIRCRAFT_TYPE.aircraft_type_category = categoryQuery;
END;

-- endpoint: /v1/aircraft/{id}
-- CALL uspGetAircraftByID(7);
DROP PROCEDURE IF EXISTS `uspGetAircraftByID`;
CREATE PROCEDURE uspGetAircraftByID(
    IN aid INTEGER
)
BEGIN
    SELECT ac_id, ac_callsign, ac_n_number, aircraft_type_manufacturer, aircraft_type_title,
    aircraft_type_category, ac_lat, ac_long, ac_loc_display_name, status_short_desc
    FROM tblAIRCRAFT
    JOIN tblAIRCRAFT_TYPE ON tblAIRCRAFT.ac_type_id = tblAIRCRAFT_TYPE.aircraft_type_id
    JOIN tblASSIGNED_STATUS ON tblAIRCRAFT.ac_id = tblASSIGNED_STATUS.aircraft_id
    JOIN tblAIRCRAFT_STATUS ON tblASSIGNED_STATUS.status_id = tblAIRCRAFT_STATUS.status_id
    WHERE tblAIRCRAFT.ac_id = aid;
END;

-- endpoint: /v1/aircraft/??
-- Return: Reason and Endtime
-- CALL uspGetOOSByAircraft(7)
DROP PROCEDURE IF EXISTS `uspGetOOSByAircraft`;
CREATE PROCEDURE uspGetOOSByAircraft(
    IN aid INTEGER
)
BEGIN
    SELECT tblAIRCRAFT_SCHED_SERVICE.ac_sched_service_reason, tblAIRCRAFT_SCHED_SERVICE.OOS_end_time
    FROM tblASSIGNED_STATUS
    INNER JOIN tblAIRCRAFT_STATUS ON tblASSIGNED_STATUS.status_id = tblAIRCRAFT_STATUS.status_id
    INNER JOIN tblAIRCRAFT ON tblASSIGNED_STATUS.aircraft_id = tblAIRCRAFT.ac_id
    INNER JOIN tblAIRCRAFT_SCHED_SERVICE ON tblAIRCRAFT.ac_id = tblAIRCRAFT_SCHED_SERVICE.ac_id
    WHERE tblASSIGNED_STATUS.aircraft_id = aid
    AND tblAIRCRAFT_STATUS.status_short_desc = "OOS"; 
END;

-- endpoint: /v1/aircraft/??
-- Return: Reason, Starttime, and Endtime
-- CALL uspGetOOSDetailByAircraft(7)
DROP PROCEDURE IF EXISTS `uspGetOOSDetailByAircraft`;
CREATE PROCEDURE uspGetOOSDetailByAircraft(
    IN aid INTEGER
)
BEGIN
    SELECT tblAIRCRAFT_SCHED_SERVICE.ac_sched_service_reason, 
    tblAIRCRAFT_SCHED_SERVICE.OOS_start_time, tblAIRCRAFT_SCHED_SERVICE.OOS_end_time
    FROM tblASSIGNED_STATUS
    INNER JOIN tblAIRCRAFT_STATUS ON tblASSIGNED_STATUS.status_id = tblAIRCRAFT_STATUS.status_id
    INNER JOIN tblAIRCRAFT ON tblASSIGNED_STATUS.aircraft_id = tblAIRCRAFT.ac_id
    INNER JOIN tblAIRCRAFT_SCHED_SERVICE ON tblAIRCRAFT.ac_id = tblAIRCRAFT_SCHED_SERVICE.ac_id
    WHERE tblASSIGNED_STATUS.aircraft_id = aid
    AND tblAIRCRAFT_STATUS.status_short_desc = "OOS"; 
END;

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

-- endpoint: nested stored procedure
/*
    CALL uspGetMissionIDByAircraft(5, @out);
    SELECT @out;
*/
-- Present issue: returns most recent mission for a given mission. Need to ensure this is the 
-- correct logic. TODO
DROP PROCEDURE IF EXISTS `uspGetMissionIDByAircraft`;
CREATE PROCEDURE uspGetMissionIDByAircraft(
    IN aid INTEGER,
    OUT mid_output_param INTEGER
)
BEGIN
    SET mid_output_param = (
        SELECT mission_id
        FROM tblMISSION
        WHERE tblMISSION.aircraft_id = aid
        ORDER BY tblMISSION.mission_date DESC LIMIT 1
    );
END;

-- endpoing: internal to apiserver
-- CALL uspGetAircraftIDByCallsign("AL5");
DROP PROCEDURE IF EXISTS `uspGetAircraftIDByCallsign`;
CREATE PROCEDURE uspGetAircraftIDByCallsign(
    IN callsignQuery NVARCHAR(50)
)
BEGIN
    SELECT tblAIRCRAFT.ac_id
    FROM tblAIRCRAFT
    WHERE tblAIRCRAFT.ac_callsign=callsignQuery;
END;

DROP PROCEDURE IF EXISTS `uspGetAircraftCallsignByID`;
CREATE PROCEDURE `uspGetAircraftCallsignByID`( 
    IN idQuery NVARCHAR(50) 
) 
BEGIN 
    SELECT tblAIRCRAFT.ac_callsign
    FROM tblAIRCRAFT 
    WHERE tblAIRCRAFT.ac_id=idQuery; 
END

-- CALL uspGetAircraftIDByMission("1");
DROP PROCEDURE IF EXISTS `uspGetAircraftIDByMission`;
CREATE PROCEDURE `uspGetAircraftIDByMission`( 
    IN idQuery NVARCHAR(50),
    OUT aid_output_param INTEGER
) 
BEGIN
    SET aid_output_param = (
        SELECT tblMISSION.aircraft_id
        FROM tblMISSION 
        WHERE tblMISSION.mission_id=idQuery 
    );
END;

-- endpoint: /v1/??/{id}
-- CALL uspGetMissionByAircraft(5);
-- Given an aircraft return follwing details about current or most recent mission:
--      mission_id, mission_type_short_name, tc_number, mission_date, m_status_short_desc
-- Present issue: returns most recent mission for a given mission. Need to ensure this is the 
-- correct logic. TODO
-- Note: edited to return mission type name instead of ID
DROP PROCEDURE IF EXISTS `uspGetMissionByAircraft`;
CREATE PROCEDURE uspGetMissionByAircraft(
    IN aid INTEGER
)
BEGIN

    DECLARE active_mission_id INTEGER; -- declare resets active_mission_id to null w/ each sproc call
    CALL uspGetMissionIDByAircraft(aid ,active_mission_id);

    SELECT mission_type_short_name, tc_number, mission_date, tblMISSION_STATUS.m_status_short_desc
    FROM tblMISSION
    INNER JOIN tblMISSION_TYPE ON tblMISSION.mission_type_id = tblMISSION_TYPE.mission_type_id
    INNER JOIN tblASSIGNED_MISSION_STATUS ON tblMISSION.mission_id = tblASSIGNED_MISSION_STATUS.mission_id
    INNER JOIN tblMISSION_STATUS ON tblASSIGNED_MISSION_STATUS.m_status_id = tblMISSION_STATUS.m_status_id
    WHERE tblMISSION.mission_id = active_mission_id
    ORDER BY tblASSIGNED_MISSION_STATUS.mission_status_date DESC LIMIT 1;
END;

-- endpoint: /v1/????
-- returns: mission_type, flight rules, TC, req, rec
-- CALL uspGetMissionDetailByAircraft(1);
-- Note: no concept of flight rules by mission, by waypoint instead, how to handle?
DROP PROCEDURE IF EXISTS `uspGetMissionDetailByAircraft`;
CREATE PROCEDURE uspGetMissionDetailByAircraft(
    IN aid INTEGER
)
BEGIN
    DECLARE active_mission_id INTEGER; -- declare resets active_mission_id to null w/ each sproc call
    CALL uspGetMissionIDByAircraft(aid ,active_mission_id);

    SELECT
    mission_type_short_name,
    tc_number,
    req_agency.agency_name AS req_agency_name, 
    req_agency.agency_area_code AS req_agency_area_code,
    req_agency.agency_phone AS req_agency_phone,
    req_agency_type.agency_type_name AS req_agency_type_name, 
    req_address.address_street_1 AS req_agency_address_street_1,
    req_address.address_street_2 AS req_agency_address_street_2,
    req_address.address_city AS req_agency_address_city, 
    req_address.address_state AS req_agency_address_state,
    req_address.address_zip AS req_agency_address_zip,
    rec_agency.agency_name AS rec_agency_name,
    rec_agency.agency_area_code AS rec_agency_area_code,
    rec_agency.agency_phone AS rec_agency_phone,
    rec_agency_type.agency_type_name AS rec_agency_type_name,
    rec_address.address_street_1 AS rec_agency_address_street_1,
    rec_address.address_street_2 AS rec_agency_address_street_2,
    rec_address.address_city AS rec_agency_address_city, 
    rec_address.address_state AS rec_agency_address_state,
    rec_address.address_zip AS rec_agency_address_zip,
    tblMISSION_STATUS.m_status_short_desc
    FROM tblMISSION
    INNER JOIN tblMISSION_TYPE ON tblMISSION.mission_type_id = tblMISSION_TYPE.mission_type_id
    INNER JOIN tblAGENCY req_agency ON tblMISSION.requestor_id = req_agency.agency_id
    INNER JOIN tblAGENCY rec_agency ON tblMISSION.receiver_id = rec_agency.agency_id
    INNER JOIN tblAGENCY_TYPE req_agency_type ON req_agency.agency_type_id = req_agency_type.agency_type_id
    INNER JOIN tblAGENCY_TYPE rec_agency_type ON rec_agency.agency_type_id = rec_agency_type.agency_type_id
    INNER JOIN tblADDRESS req_address ON req_agency.address_id = req_address.address_id
    INNER JOIN tblADDRESS rec_address ON rec_agency.address_id = rec_address.address_id
    INNER JOIN tblASSIGNED_MISSION_STATUS ON tblMISSION.mission_id = tblASSIGNED_MISSION_STATUS.mission_id
    INNER JOIN tblMISSION_STATUS ON tblASSIGNED_MISSION_STATUS.m_status_id = tblMISSION_STATUS.m_status_id
    WHERE tblMISSION.mission_id = active_mission_id
    ORDER BY tblASSIGNED_MISSION_STATUS.mission_status_date DESC LIMIT 1;
END;

-- endpoint: API server internal use
-- CALL uspGetAgencyDetailsByID(1);
-- Returns agency name, areacode, phone, type, address
DROP PROCEDURE IF EXISTS `uspGetAgencyDetailsByID`;
CREATE PROCEDURE uspGetAgencyDetailsByID(
    IN agency_id_query INTEGER
)
BEGIN 
    SELECT agency_name, agency_area_code, agency_phone, agency_type_name, 
    address_street_1, address_street_2, address_city, address_state, address_zip
    FROM tblAGENCY
    JOIN tblAGENCY_TYPE ON tblAGENCY.agency_type_id = tblAGENCY_TYPE.agency_type_id
    JOIN tblADDRESS ON tblAGENCY.address_id = tblADDRESS.address_id
    WHERE tblAGENCY.agency_id = agency_id_query;
END;

-- get patient by aircraft
-- Questions: can mission be provided or can only aircraft be provided?
-- CALL uspGetPatientByAircraft(1);
DROP PROCEDURE IF EXISTS `uspGetPatientByAircraft`;
CREATE PROCEDURE uspGetPatientByAircraft(
    IN aircraft_id INTEGER
)
BEGIN
    DECLARE active_mission_id INTEGER; -- declare resets active_mission_id to null w/ each sproc call
    CALL uspGetMissionIDByAircraft(aircraft_id ,active_mission_id);
    SELECT mission_id, patient_short_report, patient_intubated, patient_drips, patient_age,
    patient_weight, tblGENDER.gender_name, patient_cardiac, patient_gi_bleed, patient_OB
    FROM tblPATIENT
    JOIN tblGENDER ON tblPATIENT.patient_gender_id = tblGENDER.gender_id
    WHERE mission_id = active_mission_id;
END;

-- get mission crew by aircraft
-- endpoint: /v1/???
-- CALL uspGetMissionCrewByAircraft(2);
-- 	PersonnelID, FName, LName, Role (crew_type_name)
DROP PROCEDURE IF EXISTS `uspGetMissionCrewByAircraft`; -- TODO: DROP uspGetCrewByAircraft;
CREATE PROCEDURE uspGetMissionCrewByAircraft(
    IN aircraft_id INTEGER
)
BEGIN
    DECLARE active_mission_id INTEGER;
    CALL uspGetMissionIDByAircraft(aircraft_id, active_mission_id);
    
    -- determine current mission for aircraft
    SELECT tblPERSONNEL.personnel_id, personnel_f_name,
    personnel_l_name, tblCREW_TYPE.crew_type_name
    FROM tblMISSION_PERSONNEL
    JOIN tblPERSONNEL_CREW_TYPE
    ON tblMISSION_PERSONNEL.personnel_crew_type_id = tblPERSONNEL_CREW_TYPE.personnel_crew_type_id
    JOIN tblPERSONNEL ON tblPERSONNEL_CREW_TYPE.personnel_id = tblPERSONNEL.personnel_id
    JOIN tblCREW_TYPE ON tblPERSONNEL_CREW_TYPE.crew_type_id = tblCREW_TYPE.crew_type_id
    WHERE mission_id = active_mission_id;
END;

-- get assigned crew by aircraft
-- CALL uspGetAssignedCrewByAircraft(1);
DROP PROCEDURE IF EXISTS `uspGetAssignedCrewByAircraft`;
CREATE PROCEDURE `uspGetAssignedCrewByAircraft`(
    IN aircraft_id INTEGER
)
BEGIN
    SELECT tblPERSONNEL.personnel_id, personnel_f_name, -- TODO: use SELECT DISTINCT(personnel_id)?
    personnel_l_name, tblCREW_TYPE.crew_type_name
    FROM tblAIRCRAFT_PERSONNEL
    INNER JOIN tblPERSONNEL_CREW_TYPE
    ON tblAIRCRAFT_PERSONNEL.personnel_crew_type_id = tblPERSONNEL_CREW_TYPE.personnel_crew_type_id
    INNER JOIN tblPERSONNEL ON tblPERSONNEL_CREW_TYPE.personnel_id = tblPERSONNEL.personnel_id
    INNER JOIN tblCREW_TYPE ON tblPERSONNEL_CREW_TYPE.crew_type_id = tblCREW_TYPE.crew_type_id
    WHERE tblAIRCRAFT_PERSONNEL.ac_id = aircraft_id
    AND tblAIRCRAFT_PERSONNEL.shift_start < NOW()
    AND tblAIRCRAFT_PERSONNEL.shift_end > NOW(); 
END;

-- endpoint: /v1/waypoints??
-- CALL uspGetWaypointsByAircraft(5);
-- Returns: waypoint_title, waypoint_ETA, waypoint_status, waypoint_flight_rules
DROP PROCEDURE IF EXISTS `uspGetWaypointsByAircraft`;
CREATE PROCEDURE uspGetWaypointsByAircraft(
    IN aid INTEGER
)
BEGIN
    DECLARE active_mission_id INTEGER;
    CALL uspGetMissionIDByAircraft(aid, active_mission_id);
    SELECT mission_waypoint_id, waypoint_title, mission_ETA, waypoint_active, waypoint_completed,
    flight_rules,latitude, longitude
    FROM tblMISSION_WAYPOINT
    INNER JOIN tblWAYPOINT ON tblMISSION_WAYPOINT.waypoint_id = tblWAYPOINT.waypoint_id
    WHERE tblMISSION_WAYPOINT.mission_id = active_mission_id;
END;
