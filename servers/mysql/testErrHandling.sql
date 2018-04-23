-- Sandbox to test error handling in stored procedures

CREATE PROCEDURE testErrorHandling()
BEGIN

    DECLARE err INT DEFAULT FALSE;
    DECLARE run INT;

    BEGIN
    DECLARE CONTINUE HANDLER FOR SQLEXCEPTION
        BEGIN
            SET err = true;
        END;

    IF err THEN
        ROLLBACK;
    ELSE
        COMMIT;
    END IF;

    DECLARE err INT DEFAULT FALSE;
    SET err = FALSE;

END

SHOW PROCEDURE STATUS
CREATE PROCEDURE p()
BEGIN
  DECLARE counter INT DEFAULT 0;
  WHILE counter < 10 DO
    -- ... do work ...
    SET counter = counter + 1;
  END WHILE;
END;
