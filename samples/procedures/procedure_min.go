package samples


func (proc *ProcedureSample) GetProcedureMinimalSample(rdb string) string {
	sql := ""
	switch rdb {
	case "mysql":
		sql = proc.createProcedure("sp_copy_data", "mysql") + proc.createProcedureMinimalMySQL()
	case "postgres":
		sql = proc.createProcedure("sp_copy_data", "postgres") + proc.createProcedureMinimalPostgreSQL() 
	}
	return sql
}


func (proc *ProcedureSample) createProcedureMinimalMySQL() string {
return `
DROP TABLE IF EXISTS clients;
CREATE TABLE clients (
    id INT NOT NULL AUTO_INCREMENT KEY,
    name VARCHAR(100) NOT NULL,
    code VARCHAR(20) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS clients_copy;
CREATE TABLE clients_copy (
    id INT NOT NULL AUTO_INCREMENT KEY,
    name VARCHAR(100) NOT NULL,
    code VARCHAR(20) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO clients (name, code) VALUES ('Maria', 'CL-001');
INSERT INTO clients (name, code) VALUES ('John', 'CL-002');
INSERT INTO clients (name, code) VALUES ('Alfred', 'CL-003');
INSERT INTO clients (name, code) VALUES ('Patrick', 'CL-004');
INSERT INTO clients (name, code) VALUES ('Anna', 'CL-005');

-- PROCEDURE
DROP PROCEDURE IF EXISTS sp_copy_data;

DELIMITER $$

CREATE PROCEDURE sp_copy_data()
BEGIN
    START TRANSACTION;

    INSERT INTO clients_copy(name, code)
    SELECT name, code FROM clients;

    COMMIT;
END $$

DELIMITER ;

`
}


func (proc *ProcedureSample) createProcedureMinimalPostgreSQL() string {
return `
DROP TABLE IF EXISTS clients;
CREATE TABLE clients (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    code VARCHAR(20) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS clients_copy;
CREATE TABLE clients_copy (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    code VARCHAR(20) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO clients (name, code) VALUES ('Maria', 'CL-001');
INSERT INTO clients (name, code) VALUES ('John', 'CL-002');
INSERT INTO clients (name, code) VALUES ('Alfred', 'CL-003');
INSERT INTO clients (name, code) VALUES ('Patrick', 'CL-004');
INSERT INTO clients (name, code) VALUES ('Anna', 'CL-005');

-- PROCEDURE
DROP PROCEDURE IF EXISTS sp_copy_data;

CREATE PROCEDURE sp_copy_data()
LANGUAGE plpgsql 
AS $$
BEGIN

    INSERT INTO clients_copy(name, code)
    SELECT name, code FROM clients;

    COMMIT;
END;
$$;

`
}

