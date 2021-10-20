/*
    mysql_load_csv_data.sql
    Created: Saturday May 12, 2018
    Modified: Wednesday October 20, 2021
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

/*
Uploading .csv to Google Cloud Storage from ALNWMGMT machine: 
gsutil cp [filename.ext] gs://elevate-test-flight-vector-import
*/

/*
 Populate tblMISSION_TYPE
 FV Source Table(s): [Call Types]
 FV Source Column(s): [Call Types].ID, [Call Types].CallType
 FV T-SQL Query: SELECT TOP(50) ID, CallType FROM [Call Types]
 */

-- Commands to load data from 3 .csv files
 
LOAD DATA 
    LOCAL
    INFILE '/Users/benjaminleeds/Desktop/ALNW/tblGROUPS.csv'
    INTO TABLE tblGROUP
    FIELDS
        TERMINATED BY ','
        ENCLOSED BY '"'
    LINES TERMINATED BY '\r\n';

SELECT * FROM tblMISSION_TYPE;
LOAD DATA
    LOCAL
    INFILE '/Users/benjaminleeds/Desktop/ALNW/tblMISSION_TYPE.csv'
    INTO TABLE tblMISSION_TYPE
    FIELDS
        TERMINATED BY ','
        ENCLOSED BY '"'
    LINES TERMINATED BY '\r\n';

LOAD DATA
    LOCAL
    INFILE '/Users/benjaminleeds/Desktop/ALNW/tblWAYPOINT.csv'
    INTO TABLE tblWAYPOINT
    FIELDS
        TERMINATED BY ','
        ENCLOSED BY '"'
    LINES TERMINATED BY '\r\n';