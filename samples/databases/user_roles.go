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
-----------------------------------------------------

DROP TABLE IF EXISTS roles;
CREATE TABLE roles (
    role_id SERIAL PRIMARY KEY,
    role_name VARCHAR(50) UNIQUE NOT NULL,
    description VARCHAR(100),
    unique_id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

DROP TYPE IF EXISTS TYPE_USER_STATUS;
CREATE TYPE TYPE_USER_STATUS AS ENUM('Yes', 'No');

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


DROP TABLE IF EXISTS permissions;
CREATE TABLE permissions (
    permission_id SERIAL PRIMARY KEY,
    permission_name VARCHAR(50) NOT NULL,
    description VARCHAR(200),
    unique_id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

DROP TABLE IF EXISTS permission_roles;
CREATE TABLE permission_roles (
    permission_role_id SERIAL PRIMARY KEY,
    role_id INT NOT NULL,
    permission_id INT NOT NULL,
    unique_id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT fk_role FOREIGN KEY(role_id) REFERENCES roles(role_id),
    CONSTRAINT fk_permission FOREIGN KEY(permission_id) REFERENCES permissions(permission_id)
);

`
}


func (db *DatabaseSample) createUserRolesTablesMySQL() string {
return `

--- CREATE TABLES ---------------------------------
-----------------------------------------------------

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

DROP TABLE IF EXISTS permissions;
CREATE TABLE permissions (
    permission_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    permission_name VARCHAR(100) NOT NULL,
    description VARCHAR(200),
    unique_id BINARY(32) DEFAULT (UUID()),
    created_at DATETIME DEFAULT NOW(),
    updated_at DATETIME DEFAULT NOW()
);

DROP TABLE IF EXISTS permission_roles;
CREATE TABLE permission_roles (
    permission_role_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    role_id INT NOT NULL,
    permission_id INT NOT NULL,
    unique_id BINARY(32) DEFAULT (UUID()),
    created_at DATETIME DEFAULT NOW(),
    updated_at DATETIME DEFAULT NOW(),
    CONSTRAINT fk_role FOREIGN KEY(role_id) REFERENCES roles(role_id),
    CONSTRAINT fk_permission FOREIGN KEY(permission_id) REFERENCES permissions(permission_id)
);
	
`
}


func (db *DatabaseSample) insertsUserRoles() string {
return `

INSERT INTO roles (role_name, description) VALUES ('super_admin', 'Super Administrator');
INSERT INTO roles (role_name, description) VALUES ('administrator', 'Administrator');
INSERT INTO roles (role_name, description) VALUES ('employee', 'Employee');
INSERT INTO roles (role_name, description) VALUES ('secretary', 'Secretary');
INSERT INTO roles (role_name, description) VALUES ('customer', 'Customer');
INSERT INTO roles (role_name, description) VALUES ('supplier', 'Supplier');

INSERT INTO permissions (permission_name, description) VALUES ('create-role', 'Create Roles');
INSERT INTO permissions (permission_name, description) VALUES ('update-role', 'Update Roles');
INSERT INTO permissions (permission_name, description) VALUES ('assign-role', 'Assign Roles');
INSERT INTO permissions (permission_name, description) VALUES ('list-role', 'List Roles');
INSERT INTO permissions (permission_name, description) VALUES ('create-user', 'Create User');
INSERT INTO permissions (permission_name, description) VALUES ('delete-user', 'Delete User');
INSERT INTO permissions (permission_name, description) VALUES ('update-user', 'Update User');
INSERT INTO permissions (permission_name, description) VALUES ('list-user', 'List Users');
INSERT INTO permissions (permission_name, description) VALUES ('change-password', 'Change Password');
INSERT INTO permissions (permission_name, description) VALUES ('change-username', 'Change Username');
INSERT INTO permissions (permission_name, description) VALUES ('upload-user-image', 'Upload User Image');
INSERT INTO permissions (permission_name, description) VALUES ('upload-document', 'Upload Document');
INSERT INTO permissions (permission_name, description) VALUES ('change-configuration', 'Change Configuration');
INSERT INTO permissions (permission_name, description) VALUES ('list-configuration', 'List Configurations');
INSERT INTO permissions (permission_name, description) VALUES ('lock-application', 'Lock Application');
INSERT INTO permissions (permission_name, description) VALUES ('create-product', 'Create Product');
INSERT INTO permissions (permission_name, description) VALUES ('update-product', 'Update Product');
INSERT INTO permissions (permission_name, description) VALUES ('view-product', 'View Product');
INSERT INTO permissions (permission_name, description) VALUES ('delete-product', 'Delete Product');
INSERT INTO permissions (permission_name, description) VALUES ('extract-report', 'Extract Reports');
INSERT INTO permissions (permission_name, description) VALUES ('extract-statistics', 'Extract Statistics');


`
}