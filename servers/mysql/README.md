Ensure Google Cloud SDK is installed and you are signed in to the correct account. Use the command:
`gcloud auth list` to view your current logged in account. More details:

https://cloud.google.com/sdk/downloads

Install the VSCode MySql Extension:
    MySQL by Jun Han

Option 1: Using the Cloud SQL Proxy

Full Instructions: https://cloud.google.com/sql/docs/mysql/connect-admin-proxy

gcloud auth application-default login

Start Google Cloud SQL Proxy Instance
Terminal session must be in directory where SQL Proxy is installed
./cloud_sql_proxy -instances=airliftnw-uw:us-west1:alnw-dev=tcp:3306

Option 2: Direct Connection

In the mySQL VSCODE extension (bottom-left of Explorer Pane) click the "+" symbol to add a
connection.

    host: Database IP Address
    user: your mySQL username
    password: your mySQL user password
    port: 3306
    certificate file path: only used with service accounts

To upload CSV to Google Cloud Storage for import into Google Cloud SQL:

https://cloud.google.com/storage/docs/uploading-objects

gsutil cp RAW_aircraft_mission_agency_data.csv gs://info445-import-test/

MySQL Style and Configuration Guidelines
More details: http://www.sqlstyle.guide/
---------------------------------------------------

Local vs. User Variables:
    Use local variables wherever possible. They are restricted to the block where they are declared.
    Otherwise, @variable variables (local variables) can be used which are avialable anywhere in a 
    client session. 

Naming Conventions

Attribute Names:
* 30 characters maximum
* all lowercase
* underscore word separators

Datatype Standards:

Use NVARCHAR for all textual data unless data type will never require non-ASCII characters.
Step sizes:
    10 - short categorical values
    25 - 
    50 - longer strings such as address lines
    100 - 
    250 - 
    500 - freeform notes fields

Store all dates in ISO-8601 compliant format. https://en.wikipedia.org/wiki/ISO_8601

True/False Values:
    Use BOOLEAN which is a synonym for its underlying datatype: TINYINT(1)


Google Cloud MySQL Time Zone:

The local time zone can be set for the MySQL instance using the `default_time_zone`
database flag key. Setting the value of the key to -07:00 sets the database to 
Pacific Daylight Time. Note, this does not account for the 1-hour daylight savings time change to -08:00. 

References:
# system_time_zone: inherited from machine server is running on
https://dev.mysql.com/doc/refman/5.7/en/server-system-variables.html#sysvar_system_time_zone

# default-time-zone: overrides system_time_zone to DBA set time zone
https://dev.mysql.com/doc/refman/5.7/en/server-options.html#option_mysqld_default-time-zone

# time_zone: current time zone for each client that connects
https://dev.mysql.com/doc/refman/5.7/en/server-system-variables.html#sysvar_time_zone
