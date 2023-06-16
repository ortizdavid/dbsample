package samples


func (db *DatabaseSample) GetUserRolesSample(rdb string) string {
	sql := ""
	switch rdb {
	case "mysql":
		sql = db.createDatabase("db_user_roles", "mysql") + db.createUserRolesTablesMySQL() + db.insertsUserRoles()

	case "postgres":
		sql = db.createDatabase("db_user_roles", "postgres") + db.createUserRolesTablesPostgreSQL() + db.insertsUserRoles()
	}
	return sql
}

func (db *DatabaseSample) createUserRolesTablesPostgreSQL() string {
return `
--- CREATE TABLES ---------------------------------
-----------------------------------------------------------------
CREATE TYPE TYPE_USER_STATUS AS ENUM('Yes', 'No');

DROP TABLE IF EXISTS roles;
CREATE TABLE roles (
    role_id SERIAL PRIMARY KEY,
    role_name VARCHAR(50) UNIQUE NOT NULL,
    description VARCHAR(100),
    unique_id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

DROP TABLE IF EXISTS users;
CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    role_id INT NOT NULL,
    user_name VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(200) NOT NULL,
    active TYPE_USER_STATUS NOT NULL DEFAULT 'Yes',
    unique_id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT fk_role_user FOREIGN KEY(role_id) REFERENCES roles(role_id)
);

DROP TABLE IF EXISTS user_roles;
CREATE TABLE user_roles (
    user_role_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    role_id INT NOT NULL,
    unique_id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(user_id),
    CONSTRAINT fk_role FOREIGN KEY(role_id) REFERENCES roles(role_id)
);

`
}


func (db *DatabaseSample) createUserRolesTablesMySQL() string {
return `

--- CREATE TABLES  ---------------------------
--------------------------------------------------------------

DROP TABLE IF EXISTS roles;
CREATE TABLE roles (
    role_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    role_name VARCHAR(50) UNIQUE NOT NULL,
    description VARCHAR(100),
    unique_id BINARY(32) DEFAULT (UUID()),
    created_at DATETIME DEFAULT NOW(),
    updated_at DATETIME DEFAULT NOW()
);

DROP TABLE IF EXISTS users;
CREATE TABLE users (
    user_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    role_id INT NOT NULL,
    user_name VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(200) NOT NULL,
    active ENUM('Yes', 'No') NOT NULL DEFAULT 'Yes',
    unique_id BINARY(32) DEFAULT (UUID()),
    created_at DATETIME DEFAULT NOW(),
    updated_at DATETIME DEFAULT NOW(),
    CONSTRAINT fk_role FOREIGN KEY(role_id) REFERENCES roles(role_id)
);

DROP TABLE IF EXISTS user_roles;
CREATE TABLE user_roles (
    user_role_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    role_id INT NOT NULL,
    unique_id BINARY(32) DEFAULT (UUID()),
    created_at DATETIME DEFAULT NOW(),
    updated_at DATETIME DEFAULT NOW(),
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(user_id),
    CONSTRAINT fk_role FOREIGN KEY(role_id) REFERENCES roles(role_id)
);
	
`
}


func (db *DatabaseSample) insertsUserRoles() string {
return `
-------- INSERT VALUES ---------------------------------------------------------------------
INSERT INTO roles (role_name, description) VALUES ('super_admin', 'Super Administrator');
INSERT INTO roles (role_name, description) VALUES ('administrator', 'Administrator');
INSERT INTO roles (role_name, description) VALUES ('employee', 'Employee');
INSERT INTO roles (role_name, description) VALUES ('customer', 'Customer');

`
}