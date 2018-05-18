/*
my_sql_reference.sql
Created: Wednesday May 2, 2018
Modified: Friday May 18, 2018
Authors: J. Benjamin Leeds
License: None

List of common mySQL 5.7 commands for reference.
*/

-- show all tables
SHOW TABLES;

-- show table schema
DESCRIBE TEST.tblMISSION;
DESCRIBE TEST.tblAIRCRAFT;
DESCRIBE TEST.tblAIRCRAFT_TYPE;

-- alter existing table schema
-- https://dev.mysql.com/doc/refman/5.7/en/alter-table.html
ALTER TABLE [tbl_name]
ADD [column_name] INTEGER 

ALTER TABLE tblAIRCRAFT
ADD ac_type_id INTEGER
ADD FOREIGN KEY (ac_type_id) REFERENCES tblAIRCRAFT_TYPE(aircraft_type_id);

-- altering data in existing column
UPDATE [tblNAME]
SET [columnName]=[update]
WHERE [filter clause]

UPDATE tblAIRCRAFT
SET tblAIRCRAFT.ac_type_id = 1
WHERE aircraft_id = 10;

-- show stored procedures
SHOW PROCEDURE STATUS;

-- show current time zone settings for MySQL
SELECT @@global.time_zone, @@session.time_zone, @@system_time_zone
SELECT NOW(); -- select current time

-- show current set flags for MySQL DB:
-- More Details: https://cloud.google.com/sql/docs/mysql/flags
SHOW VARIABLES;