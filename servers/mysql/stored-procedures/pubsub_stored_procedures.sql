/*
pubsub_stored_procedures.sql
Created: Monday April 30, 2018
Modified: Saturday May 27, 2018
Author(s): J. Benjamin Leeds
License: None

Use the stored procedures in this file to store data received by Pub/Sub in MySQL

*/

-- uspNewMissionInsert
-- inserts a new mession read from the message queue

DELIMITER $$
DROP PROCEDURE IF EXISTS `uspNewMission`$$
CREATE PROCEDURE uspNewMission(
    IN mission_id_param         INTEGER,
    IN tc_number_param          VARCHAR(10),
    IN aircraft_id_param        INTEGER,
    IN requestor_id_param       INTEGER,
    IN receiver_id_param        INTEGER,
    IN mission_type_id_param    INTEGER,
    
    -- patient details
    IN patient_short_report_param   NVARCHAR(500),
    IN patient_intubated_param      BOOLEAN,
    IN patient_drips_param          TINYINT,
    IN patient_age_param            TINYINT, 
    IN patient_weight_param         SMALLINT,
    IN patient_gender_param         INTEGER, -- may not be able to use an integer
    IN patient_cardiac_param        BOOLEAN,
    IN patient_gi_bleed_param       BOOLEAN,
    IN patient_OB_param             BOOLEAN,

    -- how to handle indeterminate # of crew members
    IN mission_personnel_param NVARCHAR(255)
    -- IN waypoints                -- array
)

BEGIN
		    
	DECLARE asid INTEGER;
	DECLARE msid INTEGER;
    DECLARE pctid INTEGER;
	START TRANSACTION;

    -- Step 1: Create mission
    INSERT INTO tblMISSION(mission_id, aircraft_id, mission_type_id, requestor_id, receiver_id, tc_number) 
    VALUES(mission_id_param, aircraft_id_param, mission_type_id_param, requestor_id_param, receiver_id_param, tc_number_param);
   
	-- determine aircraft_status_id
    SET asid = (SELECT status_id FROM tblAIRCRAFT_STATUS WHERE status_short_desc = "OAM");
    
    -- Step 2: assign aircraft (check if aircraft is RFM and update status to OAM)
    INSERT INTO tblASSIGNED_STATUS(aircraft_status_id, status_id, aircraft_id)
    VALUES(10, @asid, aircraft_id_param);
    
    -- determine mission m_status_id
    SET msid = (SELECT m_status_id FROM tblMISSION_STATUS WHERE m_status_short_desc = "IP");

    -- Step 3: Assign mission status
    INSERT INTO tblASSIGNED_MISSION_STATUS(mission_id, m_status_id)
    VALUES(mission_id_param, @msid);

    -- Step 4: insert patient details
    INSERT INTO tblPATIENT(
        mission_id, patient_gender, patient_short_report, patient_intubated, patient_drips, 
        patient_age, patient_weight, patient_cardiac, patient_gi_bleed, patient_OB
    )
    VALUES(
        mission_id_param, patient_gender_param, patient_short_report_param, patient_intubated_param,
        patient_drips_param, patient_age_param, patient_weight_param, patient_cardiac_param,
        patient_gi_bleed_param, patient_OB_param
    );
    
    -- iterate over crew and add each to mission
    -- NOTE: Logic in progress
    SET i = 1;
    SET mpid = "null";
    
    WHILE i IS NOT null
    DO
		SET mpid = ELT(i, mission_personnel_param); -- retrive 
        SET mptid = ELT(i, mission_personnel_type_param);
    END WHILE;

	-- determine personnel_crew_type_id
    SET pctid = (SELECT personnel_crew_type_id FROM tblPERSONNEL_CREW_TYPE WHERE personnel_id = personnel_id_param AND crew_type_id = crew_type_id_param);
    -- Step 5: Assign mission personnel
    INSERT INTO tblMISSION_PERSONNEL(mission_id, personnel_crew_type_id)
    VALUES(mission_id_param, pctid);

    COMMIT;


END$$
DELIMITER ;

/*

Incoming JSON Object for this Stored Procedure:

 "missionID": "1",                       // Table: Missions.ID
    "TCNum": "18-0013",                     // Table: TC.ID
    "asset": "N123AL",
    "requestorID": "1",                     // Ex. Snoqualmie Pass Ski Area
    "receiverID": "1",                      // Ex. Harborview Medical Center
    "priority": "Emergency",
    "callType": "callTypeID",
    "patient": {
        "shortReport": "head bleed",
        "intubated": "true",
        "drips": "4",
        "age": "42",
        "weight": "50",
        "gender": "M",
        "cardiac": "false",
        "GIBleed": "false",
        "OB": "false",
    },
    "crewMemberID": [
            {
                "crewID":"crewID1",
                "crewRoleID":"crewRoldID1"
            },
            {
                "crewID":"crewID2",
                "crewRoleID":"crewRoldID2"
            },
            {
                "crewID":"crewID3",
                "crewRoleID":"crewRoldID3"
            },
            {
                "crewID":"crewID4",
                "crewRoleID":"crewRoldID4"
            }
    ],
    "waypoints": [
        {
            "ID": "1",
            "ETE": "00:05",             // time to next point
            "ETT": "00:05",             // cumulative mission time
            "active": "true"            // denotes active waypoint
                                        // Table: Missions.CurrentLeg
        },
        {
            "ID": "2",
            "ETE": "00:17",
            "ETT": "00:22",
            "active": "false"           // Table: Missions.CurrentLeg
        },
        {
            "ID": "3",
            "ETE": "00:12",
            "ETT": "00:34",
            "active": "false"           // Table: Missions.CurrentLeg
        }
    ],
}