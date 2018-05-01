/*
pubsub_stored_procedures.sql
Created: Monday April 30, 2018
Modified: Tuesday May 1, 2018
Author(s): J. Benjamin Leeds
License: None

Use the stored procedures in this file to store data received by Pub/Sub in MySQL

*/

-- uspNewMissionInsert
-- inserts a new mession read from the message queue

CREATE PROCEDURE uspNewMissionInsert(
    IN mission_id               INTEGER,
    IN tc_number                VARCHAR(10),
    IN aircraft_id              INTEGER,
    IN requestor_id             INTEGER,
    IN receiver_id              INTEGER,
    IN priority                 NVARCHAR(25),
    IN call_type_id             INTEGER,
    
    -- patient details
    IN patient_short_report     NVARCHAR(500),
    IN patient_intubated        BOOLEAN,
    IN patient_drips            TINYINT,
    IN patient_age              TINYINT, 
    IN patient_weight           SMALLINT,
    IN patient_gender           INTEGER,
    IN patient_cardiac          BOOLEAN,
    IN patient_gi_bleed         BOOLEAN,
    IN patient_OB               BOOLEAN,

    -- how to handle indeterminate # of crew members
    IN assignedCrew             -- array
    IN waypoints                -- array
)

BEGIN

    -- cannot insert into mission table until we know all foreign keys:
        -- tblAIRCRAFT(aircraft_id)
        -- requestor
        -- receiver
        -- priority? use lookup table or naw?
        -- call_type_id
        -- tblPATIENT_DETAILS

    -- insert mission into tblMISSION
    START TRANSACTION;

    INSERT INTO tblMISSION(
        mission_id, tc_number, aircraft_id, requestor_id, receiver_id, priority,
        call_type_id --...
    )
    VALUES mission_id, tc_number, aircraft_id, requestor_id, receiver_id, priority, call_type_id;

    INSERT INTO tblPATIENT(
        mission_id, patient_gender, patient_short_report, patient_intubated, patient_drips, 
        patient_age, patient_weight, patient_cardiac, patient_gi_bleed, patient_OB
    )
    VALUES

    IF err THEN -- research where err comes from
        ROLLBACK;
    ELSE
        COMMIT;
    END IF;
END;

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
            "personID1",
            "personID2",
            "personID3",
            "personID4"
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