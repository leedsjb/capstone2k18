/*
    mysql_load_csv_data.sql
    Created: Saturday May 12, 2018
    Modified: Monday May 21, 2018
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

/* Important: before executing import necessary data from CSV in Cloud Console
Current CSV Import Tables:
    tblGROUP;
    SELECT * FROM tblGROUP;
    UPDATE tblGROUP -- ISSUE TO BE RESOLVED WHERE 1st IMPORTED ROW HAS id=0
    SET group_id=1
    WHERE group_id=0

    tblMISSION_TYPE;
    SELECT * FROM tblMISSION_TYPE;
    UPDATE tblMISSION_TYPE
    SET mission_type_id=1
    WHERE mission_type_id=0
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

INSERT INTO tblGENDER(gender_id,gender_name)
VALUES(1,'Male'), (2,'Female'), (3,'Other');

INSERT INTO tblPERSONNEL(personnel_id, personnel_f_name, personnel_l_name, personnel_title, personnel_sms_num, personnel_email)
VALUES
(1,"Tiffany","Chen","Developer","5555555555","tzc@uw.edu"),
(2,"Benjamin","Leeds","Developer","7777777777","jbl@uw.edu");

INSERT INTO tblPERSONNEL_GROUP(personnelGroupID,personnel_id,group_id)
VALUES
(1,1,20),
(2,1,18),
(3,2,12),
(4,2,18);

INSERT INTO tblAIRCRAFT_TYPE(
    aircraft_type_id, aircraft_type_title, aircraft_type_desc,
    aircraft_type_category, aircraft_type_manufacturer
)
VALUES
(1, "31-A", "Fastest aircraft in-fleet", "Fixed-wing", "Learjet"),
(2, "PC-12", "More economic aircraft", "Fixed-wing", "Pilatus"),
(3, "H-135", "Small helicopter", "Rotorcraft", "Airbus"),
(4, "A109E", "Large helicopter", "Rotorcraft", "Augusta");

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
(7, "AL9", "N129AL", "4", "47.4441", "-122.3249", "Bremerton Base", "5555555555","4444444444");

INSERT INTO tblAIRCRAFT_STATUS(status_id, status_title, status_long_desc, status_short_desc)
VALUES
(1, "Ready", "Aircraft ready for missions", "RFM"),
(2, "Out Of Service", "Aircraft out of service", "OOS"),
(3, "On a Mission", "Aircraft on a mission", "OAM");

INSERT INTO tblASSIGNED_STATUS(aircraftstatus_id,status_id, aircraft_id)
VALUES
(1,1,1),
(2,2,2),
(3,3,3),
(4,3,4),
(5,3,5),
(6,1,6),
(7,2,7);

INSERT INTO tblPERSONNEL_CREW_TYPE(personnel_crew_type_id, personnel_id, crew_type_id)
VALUES
(1,1,1),
(2,2,2);

INSERT INTO tblAGENCY_TYPE(agencytype_id, agencytype_name)
VALUES
(1,"Hospital"),
(2,"EMS Agency"),
(3,"Fire"),
(4,"Police");

INSERT INTO tblADDRESS(
    address_id, address_street, address_city, address_state, address_zip, address_zip_plus4
)
VALUES
(1, "325 9th Ave", "Seattle", "WA", 98104, NULL), -- HMC
(2, "516 3rd Ave", "Seattle", "WA", 98104, NULL), -- King County Sheriff
(3, "301 2nd Ave S", "Seattle", "WA", 98104, NULL); -- Seattle Fire Department

INSERT INTO tblAGENCY(
    agency_id, agency_name, agency_area_code, agency_phone, address_id, agencytype_id
)
VALUES
(1, "Harborview Medical Center", 206, 7443000, 1, 1),
(2, "King County Sheriff", 206, 2963311, 2, 4),
(3, "Seattle Fire Department", 206, 3861400, 3, 3);

DESCRIBE tblMISSION
SELECT * FROM tblMISSION
SELECT * FROM tblMISSION_TYPE
INSERT INTO tblMISSION(
    mission_id, aircraft_id, mission_type_id, requestor_id, receiver_id, tc_number
)
VALUES
(1, 5, 16, 3, 1, "18-0045"),
(2, 1, 6, 2, 1, "18-0036"),
(3, 7, 6, 3, 1, "18-0028");

INSERT INTO tblMISSION_STATUS(m_status_id, m_status_title, m_status_long_desc, m_status_short_desc)
VALUES
(1, "Pending", "Crew not yet departed", "PEND"),
(2, "In-progress", "Mission underway", "IP"),
(3, "Complete", "Mission complete", "COMPLETE");

INSERT INTO tblASSIGNED_MISSION_STATUS(
    missionstatus_id, mission_id, m_status_id
)
VALUES
(1, 1, 1),
(2, 2, 2),
(3, 3, 3);

INSERT INTO tblRESOURCE_LINKS(
    resource_short_name, resource_long_name, resource_url, resource_thumbnail_photo_url
)
VALUES
("PIAP","Post Incident Action Plan", "https://occam.uw.edu", "https://www.computershare.com/PublishingImages/checklist-icon.png"),
("9B","Ninth Brain LMS", "https://suite.ninthbrain.com/Logon.aspx", "https://suite.ninthbrain.com/img/company_logo/logo_standard.png"),
("EMSC","EMS Charts", "https://www.emscharts.com/pub/default.cfm", "https://www.emscharts.com/pub/images/ems_logo.jpg"),
("SP","ALNW SharePoint", "https://portal.airliftnw.org/Pages/Home.aspx", "https://pbs.twimg.com/profile_images/920356547015749632/2in54ehS_400x400.jpg"),
("ALNW","Airlift Northwest Website", "https://airliftnw.org", "https://pbs.twimg.com/profile_images/948624469043511296/pD2bNKBA_400x400.jpg"),
("OCCAM","UW Medicine OCCAM", "https://occam.uw.edu", "https://is4-ssl.mzstatic.com/image/thumb/Purple62/v4/db/ed/93/dbed9341-a3c6-597c-9755-8170e871fca7/mzl.pomnvbtr.jpg/1200x630bb.jpg");



/* 

-- SELECT STATEMENTS
SELECT * FROM tblMISSION;
SELECT * FROM tblAIRCRAFT;
SELECT * FROM tblAGENCY
SELECT * FROM tblAGENCY_TYPE
SELECT * FROM tblADDRESS
SELECT * FROM tblPERSONNEL
SELECT * FROM tblCREW_TYPE
SELECT * FROM tblPERSONNEL_CREW_TYPE

-- mission status
SELECT * FROM tblMISSION_STATUS;
SELECT * FROM tblASSIGNED_MISSION_STATUS;

-- resource links
DESCRIBE tblRESOURCE_LINKS;
SELECT * FROM tblRESOURCE_LINKS;