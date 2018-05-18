/*
    mysql_load_csv_data.sql
    Created: Saturday May 12, 2018
    Modified: Thursday May 17, 2018
    Authors: J. Benjamin Leeds
    License: None

    This script populates the Airlift Northwest Elevate application database 
    with data from Flight Vector's T-SQL database from CSV.

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

CREATE TABLE tblPERSONNEL(
    personnel_id INTEGER PRIMARY KEY NOT NULL,
    personnel_f_name NVARCHAR(50) NOT NULL,
    personnel_l_name NVARCHAR(50) NOT NULL,
    personnel_title NVARCHAR(50), 
    crew_type_id INTEGER,
    personnel_sms_num NVARCHAR(50),
    personnel_email NVARCHAR(50)
    FOREIGN KEY (crew_type_id) REFERENCES tblCREW_TYPE(crew_type_id)
);

CREATE TABLE tblGENDER(
    gender_id INTEGER AUTO_INCREMENT PRIMARY KEY, -- ** AUTO_INC??
    gender_name NVARCHAR(50) NOT NULL
)
INSERT INTO tblGENDER(gender_name)
VALUES('Male', 'Female', 'Other')

-- Possible Method to load data: 
LOAD DATA INFILE 'Crew_Export_test.csv' INTO TABLE tblPERSONNEL_CSV_TEST
FIELDS TERMINATED BY ','
ENCLOSED BY '"'
LINES TERMINATED BY '\r\n';
-- IGNORE 1 LINES;