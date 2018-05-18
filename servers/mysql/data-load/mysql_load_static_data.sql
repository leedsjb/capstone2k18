/*
    mysql_load_csv_data.sql
    Created: Saturday May 12, 2018
    Modified: Friday Mau 18, 2018
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

SELECT tblGROUP.group_id, group_name, personnel_f_name, personnel_l_name FROM tblGROUP
JOIN tblPERSONNEL_GROUP ON tblGROUP.group_id = tblPERSONNEL_GROUP.group_id
JOIN tblPERSONNEL ON tblPERSONNEL_GROUP.personnel_id = tblPERSONNEL.personnel_id
ORDER BY tblGROUP.group_name ASC

DESCRIBE tblGROUP;


DESCRIBE tblPERSONNEL_GROUP;

DESCRIBE tblPERSONNEL;
SELECT * FROM tblPERSONNEL;