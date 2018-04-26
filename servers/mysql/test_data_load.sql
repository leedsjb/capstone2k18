


DROP TABLE IF EXISTS tblAIRCRAFT_STATUS;
CREATE TABLE tblAIRCRAFT_STATUS(
    aircraftStatusID INTEGER AUTO_INCREMENT PRIMARY KEY,
    aircraftStatusName NVARCHAR(100) NOT NULL,
    aircraftStatusDescription NVARCHAR(100)
);

INSERT INTO tblAIRCRAFT_STATUS(aircraftStatusName, aircraftStatusDescription)
VALUES
    ("VFR ONLY - Aircraft Issue", "Aircraft can only fly in Visual Meterological Conditions"),
    ("Delayed Maintenance", "Maintenance not completed on time"),
    ("Delayed Other", "Delayed due to other reasons"),
    ("Hangared", "Aircraft presently in hangar"),
    ("Heavy on Fuel", "Check aircraft fuel weight for weight and balance, payload may be limited"),
    ("Last Out", "*************** What is this?"),
    ("No Riders", "Essential personnel only"),
    ("USA Only", "Cross-border flights prohibited"),
    ("2 Peds", "2 Pediatric flight nurses assigned"),
    ("2 Adult", "2 Adult flight nurses assigned"),
    ("Isolette OOS", "Isolette equipment out of service"),
    ("VFR Only - Pilot Issue", "Flight crew not presently qualified for Instrument Flight Rules"),
    ("MEL - Unable to fly into known icing", "Aircraft Minimum Equipment List prevents Flight Into
    Known Icing (FIKI)");
