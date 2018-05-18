/*
client_stored_procedures.sql
Created: thursday May 17, 2018
Modified:
Author(s): J. Benjamin Leeds
License: None

Use the stored procedures in this file to retrieve data in MySQL to send to clients

*/

DROP PROCEDURE IF EXISTS `uspGetAllGroups`;
CREATE PROCEDURE uspGetAllGroups()
BEGIN
    SELECT tblGROUP.group_id, group_name, personnel_f_name, personnel_l_name
    FROM tblGROUP
    JOIN tblPERSONNEL_GROUP ON tblGROUP.group_id = tblPERSONNEL_GROUP.group_id
    JOIN tblPERSONNEL ON tblPERSONNEL_GROUP.personnel_id = tblPERSONNEL.personnel_id
    ORDER BY tblGROUP.group_name ASC;
END;

DROP PROCEDURE IF EXISTS `uspGetGroupByID`;
CREATE PROCEDURE uspGetGroupByID(
    IN gid INTEGER
)
BEGIN
    SELECT tblGROUP.group_id, group_name, personnel_f_name, personnel_l_name
    FROM tblGROUP
    JOIN tblPERSONNEL_GROUP ON tblGROUP.group_id = tblPERSONNEL_GROUP.group_id
    JOIN tblPERSONNEL ON tblPERSONNEL_GROUP.personnel_id = tblPERSONNEL.personnel_id 
    WHERE tblGROUP.group_id = gid;
END;