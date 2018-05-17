

DROP PROCEDURE IF EXISTS `uspLastMission`;
CREATE PROCEDURE uspLastMission()(
    IN aircraft_callsign NVARCHAR(100)
)
BEGIN
SELECT mission_date
FROM tblMISSION
ORDER BY mission_date ASC
LIMIT 1

CASE mission_date
WHEN NOW() - mission_date < 1
THEN 'last mission more than 1 hour old'
WHEN NOW() - mission_date < 5
THEN `last mission less than 5 hours ago`
END

CALL uspClosestBaseToAddress('Portland');


CASE expression
WHEN condition1 THEN result1
WHEN condition2 THEN result2
ELSE result -- used when no prior cases match
END CASE

-- list stored procedures
SHOW PROCEDURE STATUS


SELECT * FROM tblAIRCRAFT


