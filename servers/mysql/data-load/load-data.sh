#   mysql_load_csv_data.sql
#   Created: Thursday May 17, 2018
#   Modified: Wednesday October 20, 2021
#   Authors: J. Benjamin Leeds
#   License: None
#
#   This script populates the Airlift Northwest Elevate application database 
#   with static data in CSV format.

# Create MySQL Instance

# Step 1: Create Google Cloud MySQL Instance
gcloud services enable sqladmin.googleapis.com
gcloud sql instances create elevate --tier=db-f1-micro --region=us-west1

# Step 2: Connect to Instance Locally using Cloud SQL Proxy

sh start-proxy.sh

# Step 3: Create Database, Load Static Data

    # 3a) Create Schema w/ ddl.sql
    # 3b) Load tblWAYPOINT_TYPE data
    # 3c) Load CSV FIles with mysql_load_csv_data.sql
    # 3d) Load static data with mysql_load_static_data.sql

# Step 4: Load User Stored Procedures

    # client_stored_procedures.sql

# OLD: 

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


