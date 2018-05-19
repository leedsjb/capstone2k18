/*
    mysql_load_csv_data.sql
    Created: Saturday May 12, 2018
    Modified: Saturday May 19, 2018
    Authors: J. Benjamin Leeds
    License: None

    This script populates the Airlift Northwest Elevate application database 
    with static data.

    Environment: 
        RDBMS: Google Cloud SQL MySQL Database 
        Location: us-west1-b zone
        CPU(s): 1 vCPU (multi-tenant CPU)
        Memory: 614 MB
        Storage: 10GB SSD
        Failover Replica: Disabled

    Google Cloud SQL Proxy Connection Command:
    ./cloud_sql_proxy -instances=airliftnw-uw:us-west1:alnw-elevate-test=tcp:3306
*/

INSERT INTO tblCREW_TYPE (crew_type_id, crew_type_name, crew_type_role)
VALUES
    (1,	"Pilot PIC",	        "Pilot in Command"), 
    (2,	"Pilot SIC",	        "Second in Command"),
    (3,	"Pediatric RN",	        "Medical"),
    (4,	"Adult RN",	            "Medical"),
    (5,	"MD",	                "Medical"),
    (6,	"Observer",	            "Other"),
    (7,	"ECMO",	                "Medical"),
    (8,	"RT",	                "Other"),
    (9,	"Other",	             "Other"),
    (10,"Family Member Rider",	"None"),
    (11,"Ambulance Driver",	    "Driver");

INSERT INTO tblGENDER(gender_name)
VALUES('Male', 'Female', 'Other')

INSERT INTO tblPERSONNEL(personnel_id, personnel_f_name, personnel_l_name, personnel_title, personnel_sms_num, personnel_email)
VALUES
(1,"Tiffany","Chen","Developer","5555555555","tzc@uw.edu"),
(2,"Benjamin","Leeds","Developer","7777777777","jbl@uw.edu")

INSERT INTO tblPERSONNEL_GROUP(personnelGroupID,personnel_id,group_id)
VALUES
(1,1,20),
(2,1,18),
(3,2,12),
(4,2,18)


INSERT INTO tblAIRCRAFT_TYPE(
    aircraft_type_id, aircraft_type_title, aircraft_type_desc,
    aircraft_type_category, aircraft_type_manufacturer
)
INSERT INTO tblAIRCRAFT_TYPE(aircraft_type_id)
VALUES (5)
DESCRIBE tblAIRCRAFT_TYPE
VALUES
(1, "31-A", "Fastest aircraft in-fleet", "Fixed-wing", "Learjet"),
(2, "PC-12", "More economic aircraft", "Fixed-wing", "Pilatus"),
(3, "H-135", "Small helicopter", "Rotorcraft", "Airbus"),
(4, "A109E", "Large helicopter", "Rotorcraft", "Augusta")

INSERT INTO tblAIRCRAFT(
    ac_id, ac_callsign, ac_n_number, ac_type_id, ac_lat, ac_long, ac_loc_display_name,
    ac_cell_phone, ac_sat_phone
)
VALUES
(1, "AL2", "N123AL", "1", "47.4441", "-122.3249", "Bremerton Base", "5555555555","4444444444"),
(2, "AL3", "N124AL", "1", "45.5653", "-122.6448", "Seattle Children's", "5555555555","4444444444"),
(3, "AL5", "N125AL", "2", "47.1042", "-122.87", "Harborview Medical Center", "5555555555","4444444444"),
(4, "AL6", "N126AL", "2", "47.6849", "-122.2968", "UW Medical Center", "5555555555","4444444444"),
(5, "AL7", "N127AL", "3", "45.5653", "-122.6448", "Yakima Base", "5555555555","4444444444"),
(6, "AL8", "N128AL", "3", "45.5287", "-122.6363", "Boeing Field Base", "5555555555","4444444444"),
(7, "AL9", "N129AL", "4", "47.4441", "-122.3249", "Bremerton Base", "5555555555","4444444444")

INSERT INTO tblAIRCRAFT_STATUS(status_id, status_title, status_long_desc, status_short_desc)
VALUES
(1, "Ready", "Aircraft ready for missions", "RFM"),
(2, "Out Of Service", "Aircraft out of service", "OOS"),
(3, "On a Mission", "Aircraft on a mission", "OAM")

INSERT INTO tblASSIGNED_STATUS(aircraftstatus_id,status_id, aircraft_id)
VALUES
(1,1,1),
(2,2,2),
(3,3,3),
(4,3,4),
(5,3,5),
(6,1,6),
(7,2,7)