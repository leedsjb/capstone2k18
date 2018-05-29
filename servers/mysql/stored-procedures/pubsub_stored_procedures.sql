/*
pubsub_stored_procedures.sql
Created: Monday April 30, 2018
Modified: Sunday May 27, 2018
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
    
    -- json input
    IN mission_personnel_param NVARCHAR(255), -- a json string
    IN mission_waypoints_param NVARCHAR(255) -- a json string           
)

BEGIN
		    
	DECLARE asid INTEGER;
	DECLARE msid INTEGER;
	-- for crew while loop
    DECLARE iterator INTEGER DEFAULT 0; 
    DECLARE num_crew INTEGER DEFAULT 0;
    DECLARE selector NVARCHAR(10);
    DECLARE crew_member NVARCHAR(50);
    DECLARE crew_id NVARCHAR(10);
    DECLARE crew_role_id NVARCHAR(10);
    DECLARE pctid INTEGER;
    -- for waypoints while loop
    DECLARE iterator_2 INTEGER DEFAULT 0;
    DECLARE num_waypoints INTEGER DEFAULT 0;
    DECLARE selector_2 NVARCHAR(10);
    DECLARE waypoint NVARCHAR(255);
    DECLARE waypoint_id_param NVARCHAR(10); 
    DECLARE waypoint_ETA_param TIMESTAMP; -- time format from T-SQL -> need to convert
    DECLARE waypoint_active_param NVARCHAR(10); -- 0 or 1
    DECLARE waypoint_completed_param NVARCHAR(10); -- 0 or 1
    
    
	START TRANSACTION;

		-- Step 1: Create mission
		INSERT INTO tblMISSION(mission_id, aircraft_id, mission_type_id, requestor_id, receiver_id, tc_number) 
		VALUES(mission_id_param, aircraft_id_param, mission_type_id_param, requestor_id_param, receiver_id_param, tc_number_param);
	   
		-- determine aircraft_status_id
		SET asid = (SELECT status_id FROM tblAIRCRAFT_STATUS WHERE status_short_desc = "OAM");
		
		-- Step 2: assign aircraft (check if aircraft is RFM and update status to OAM)
		INSERT INTO tblASSIGNED_STATUS(status_id, aircraft_id)
		VALUES(asid, aircraft_id_param);
		
		-- determine mission m_status_id
		SET msid = (SELECT m_status_id FROM tblMISSION_STATUS WHERE m_status_short_desc = "IP");

		-- Step 3: Assign mission status
		INSERT INTO tblASSIGNED_MISSION_STATUS(mission_id, m_status_id)
		VALUES(mission_id_param, msid);
        
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
		
        -- Step 5: iterate over crew and add each crewmember to mission in the specified role
		SET num_crew =  JSON_LENGTH(mission_personnel_param); -- 2

		WHILE iterator < num_crew

		DO
			SET selector = CONCAT('$[', iterator,  ']'); -- JSON query selector for crew member in JSON array
			SET crew_member = JSON_EXTRACT(mission_personnel_param, selector); -- retrieve crew member object from JSON array
			SET crew_id = JSON_EXTRACT(crew_member, '$.crewID');
			SET crew_role_id = JSON_EXTRACT(crew_member, '$.crewRoleID');
			
			-- lookup personnel_crew_type_id from lookup tables
			SET pctid = (
				SELECT personnel_crew_type_id
				FROM tblPERSONNEL_CREW_TYPE
				WHERE personnel_id = crew_id AND crew_type_id = crew_role_id
			);

			-- add personnel to mission
			INSERT INTO tblMISSION_PERSONNEL(mission_id, personnel_crew_type_id)
			VALUES (mission_id_param, pctid);
			
			SET iterator = iterator + 1;
            
		END WHILE;
        
        -- Step 6: iterate over waypoints and add each waypoint to mission        
		SET num_waypoints =  JSON_LENGTH(mission_waypoints_param); -- 2

		WHILE iterator_2 < num_waypoints

		DO
			SET selector_2 = CONCAT('$[', iterator_2,  ']'); -- JSON query selector for waypoint in JSON array
			SET waypoint = JSON_EXTRACT(mission_waypoints_param, selector_2); -- retrieve crew member object from JSON array
			
            SET waypoint_id_param = JSON_EXTRACT(waypoint, '$.ID');
			SET waypoint_ETA_param = JSON_EXTRACT(waypoint, '$.ETA');
			SET waypoint_active_param = JSON_EXTRACT(waypoint, '$.active');
			SET waypoint_completed_param = JSON_EXTRACT(waypoint, '$.completed');
	
			-- lookup personnel_crew_type_id from lookup tables
			SET pctid = (
				SELECT personnel_crew_type_id
				FROM tblPERSONNEL_CREW_TYPE
				WHERE personnel_id = crew_id AND crew_type_id = crew_role_id
			);

			-- add mission waypoint
            -- Note: does not handle case where temporary waypoints not in the database are sent TODO
			INSERT INTO tblMISSION_WAYPOINT(mission_id, waypoint_id, mission_ETA, waypoint_active, waypoint_completed)
			VALUES (mission_id_param, waypoint_id_param, waypoint_ETA_param, waypoint_active_param, waypoint_completed_param);
			
			SET iterator = iterator + 1;
            
		END WHILE;
        
    COMMIT;
END$$
DELIMITER ;

/*

Sample new mission:

CALL uspNewMission(
	"5",
    "18-0080",
    "2", -- AL3, N124AL, Lear 31A
    "2", -- King County Sherifft
    "1", -- Harborview Medical Center
    "19", -- FW-Jet
    "Obstructed airway. Choked on a pinecone. Obstruction cleared. Lacerated larynx.",
    "1", -- intubated
    "1", -- drips
    "15", -- age
    "75", -- weight
    "1", -- gender
    "0", -- cardiac
    "1", -- gi bleed
    "0", -- OB
    '[{"crewID":"1","crewRoleID":"1"},{"crewID":"2","crewRoleID":"2"},{"crewID":"3","crewRoleID":"3"},{"crewID":"4","crewRoleID":"4"}]', -- personnel
    '[{"ID":"22810","ETA":"2018-05-28 05:15:27","active":"0","completed":"1"},{"ID":"20001","ETA":"2018-05-28 05:30:27","active":"1","completed":"0"},{"ID":"20070","ETA":"2018-05-28 07:07:27","active":"0","completed":"0"}]' -- waypoints
);

*/
CALL uspUpdateACLocation("1", "47.4441", "-121.3249");

DROP PROCEDURE IF EXISTS `uspUpdateACLocation`;
CREATE PROCEDURE uspUpdateACLocation(
	IN aid INTEGER, 
	IN lat_param DECIMAL(9,6),
	IN long_param DECIMAL(9,6)
)
BEGIN
	UPDATE tblAIRCRAFT
	SET ac_lat = lat_param, ac_long = long_param
	WHERE tblAIRCRAFT.ac_id = aid;
END;

SELECT * FROM tblAIRCRAFT;