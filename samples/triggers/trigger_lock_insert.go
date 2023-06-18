package samples


func (trg *TriggerSample) GetTriggerLockInsert(rdb string) string {
	sql := ""
	switch rdb {
	case "mysql":
		sql = trg.createTrigger("trg_lock_insert_settings", "mysql") + trg.createTriggerLockInsertMySQL()
	case "postgres":
		sql = trg.createTrigger("trg_lock_insert_settings", "postgres") + trg.createTriggerLockInsertPostgreSQL() 
	}
	return sql
}


func (trg *TriggerSample) createTriggerLockInsertMySQL() string {
return `
--- TRIGGER: trigger_lock_insert_settings
--- RDBMS MYSQL---

DROP TABLE IF EXISTS app_settings;
CREATE TABLE app_settings (
    app_name VARCHAR(100),
    app_acronym VARCHAR(20),
    max_records_page INT,
    logotype VARCHAR(100),
    favicon VARCHAR(100),
    background_image VARCHAR(100),
    dir_upload_img VARCHAR(100),
    dir_upload_docs VARCHAR(100),
    dir_upload_videos VARCHAR(100),
    created_at DATETIME DEFAULT NOW(),
    updated_at DATETIME DEFAULT NOW()
);

INSERT INTO app_settings (app_name, app_acronym, favicon, logotype, max_records_page,  dir_upload_img)
VALUES ('Sales Management', 'SM', 'ico.png', 'logo.png', 100, 'static/uplodas/images');

-- Trigger to block the insertion of records in the settings table --
-- Must be only one record --


DROP TRIGGER IF EXISTS trg_lock_insert_settings;

DELIMITER $$

CREATE TRIGGER trg_lock_insert_settings 
BEFORE INSERT ON app_settings 
FOR EACH ROW 
BEGIN
    SET @total = (SELECT COUNT(*) FROM app_settings);
    IF (@total > 0) THEN
        SIGNAL SQLSTATE '45000'
            SET MESSAGE_TEXT = 'Table app_settings must have only one(1) record!';
    END IF;
END
$$

DELIMITER ;

`
}

func (trg *TriggerSample) createTriggerLockInsertPostgreSQL() string {
return `
--- TRIGGER: trigger_lock_insert_settings
--- RDBMS MYSQL---

DROP TABLE IF EXISTS app_settings;
CREATE TABLE app_settings (
    app_name VARCHAR(100),
    app_acronym VARCHAR(20),
    max_records_page INT,
    logotype VARCHAR(100),
    favicon VARCHAR(100),
    background_image VARCHAR(100),
    dir_upload_img VARCHAR(100),
    dir_upload_docs VARCHAR(100),
    dir_upload_videos VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO app_settings (app_name, app_acronym, favicon, logotype, max_records_page,  dir_upload_img)
VALUES ('Sales Management', 'SM', 'ico.png', 'logo.png', 100, 'static/uplodas/images');

-- First Create a function fun_lock_settings() --

CREATE OR REPLACE FUNCTION fun_lock_insert_settings()
RETURNS TRIGGER AS $$
BEGIN
    IF (SELECT COUNT(*) FROM app_settings) > 0 THEN
        RAISE EXCEPTION 'Table app_settings must have only one(1) record!';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE 'plpgsql';


-- Trigger to block the insertion of records in the settings table --
-- Must be only one record --

CREATE OR REPLACE TRIGGER tr_lock_insert_settings
    BEFORE INSERT ON app_settings
    FOR EACH ROW
    EXECUTE FUNCTION fun_lock_insert_settings();

`
}

