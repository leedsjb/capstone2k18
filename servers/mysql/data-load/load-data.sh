#   mysql_load_csv_data.sql
#   Created: Thursday May 17, 2018
#   Modified: Friday May 18. 2018
#   Authors: J. Benjamin Leeds
#   License: None
#
#   This script populates the Airlift Northwest Elevate application database 
#   with static data in CSV format.


# Step 1: Retrieve CSV Files from Flight Vector SQL Server

# Step 2: Upload CSV files to Google Cloud Storage
# gsutil cp RAW_aircraft_mission_agency_data.csv gs://info445-import-test/

# Step 3: Import Data to MySQL
# to-do: write REST POST API Call to trigger import without using UI

# Step 3 Alternate Method

# LOAD DATA INFILE 'Crew_Export_test.csv' INTO TABLE tblPERSONNEL_CSV_TEST
# FIELDS TERMINATED BY ','
# ENCLOSED BY '"'
# LINES TERMINATED BY '\r\n';
# IGNORE 1 LINES;


