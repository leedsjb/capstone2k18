
/*
alnw_info445_lab4_syntxn_wrapper.sql
Created: Wednesday May 16, 2018
Modified:
Authors: J. Benjamin Leeds
License: None

INFO 445 Lab 4
This stored procedure inserts a specified number of mission waypoints into the AirliftNW Elevate
database. It uses a wrapper and synthetic transaction stored procedure that generates random values
to pass as parameters to an insert mission stored procedure.
*/

DROP PROCEDURE IF EXISTS `uspInsertMissionWaypointNTimesWrapper`;
CREATE PROCEDURE uspInsertMissionWaypointNTimesWrapper(
    IN num_inserts INTEGER -- number of waypoints to create
)
BEGIN
    SET @num_waypoints = (SELECT COUNT(*) FROM tblWAYPOINT);
    SET @num_missions = (SELECT COUNT(*) FROM tblMISSION);
    WHILE num_inserts > 0 DO
        SET @mid =   (SELECT FLOOR(1 + (RAND() * @num_missions)));
        SET @wpid =  (SELECT FLOOR(1 + (RAND() * @num_waypoints)));
        SET @time_now = (SELECT NOW());
        SET @rand_time_delta = (SELECT FLOOR(1 + (RAND() * 7200)));
        SET @eta = (
            SELECT ADDTIME(@time_now, @rand_time_delta)
        ); -- 0 to 2 hrs
        CALL uspInsertMissionWaypoint(@mid, @wpid, @eta);
    END WHILE;
END;

DROP PROCEDURE IF EXISTS `uspInsertMissionWaypoint`;
CREATE PROCEDURE uspInsertMissionWaypoint(
    IN mid INTEGER,
    IN wpid INTEGER,
    IN eta TIMESTAMP
)
BEGIN
    DECLARE err INT DEFAULT FALSE;
    START TRANSACTION;
        INSERT INTO tblMISSION_WAYPOINT(mission_id, waypoint_id, mission_ETA)
        VALUES(@mid, @wpid, @eta);
        IF err THEN
            ROLLBACK;
        ELSE
            COMMIT;
        END IF;
        SET err = FALSE;
END;


SHOW MASTER STATUS;
SHOW SLAVE STATUS;