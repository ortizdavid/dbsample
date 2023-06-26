package samples

func (db *DatabaseSample) GetSalesSample(rdb string) string {
	sql := ""
	switch rdb {
	case "mysql":
		sql = db.createDatabase("db_sales", "mysql") + db.createSalesTablesMySQL() + db.createViewsSales() + db.createSalesProceduresMySQL() + db.insertsSales()
	case "postgres":
		sql = db.createDatabase("db_sales", "postgres") + db.createSalesTablesPostgreSQL() + db.createViewsSales() + db.createSalesProceduresPostgreSQL() + db.insertsSales()
	}
	return sql
}


func (db *DatabaseSample) createSalesTablesMySQL() string {
return `

DROP TABLE IF EXISTS enterprises;
CREATE TABLE enterprises (
    enterprise_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    enterprise_name VARCHAR(30),
    web_site VARCHAR(100),
    acronym VARCHAR(20) NOT NULL,
    identification_number VARCHAR(30),
    logo VARCHAR(100),
    unique_id BINARY(32) DEFAULT (UUID()),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS enterprise_contacts;
CREATE TABLE enterprise_contacts (
    contact_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    enterprise_id INT NOT NULL,
    phone INT,
    email VARCHAR(100)
);

DROP TABLE IF EXISTS stores;
CREATE TABLE stores (
    store_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    enterprise_id INT NOT NULL,
    store_name VARCHAR(100) UNIQUE NOT NULL,
    unique_id BINARY(32) DEFAULT (UUID()),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_enterprise_store FOREIGN KEY(enterprise_id) REFERENCES enterprises(enterprise_id)
);

DROP TABLE IF EXISTS store_address;
CREATE TABLE store_address (
    address_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    store_id INT NOT NULL,
    state VARCHAR(100),
    city VARCHAR(100),
    district VARCHAR(100),
    CONSTRAINT fk_store FOREIGN KEY(store_id) REFERENCES stores(store_id)
);

DROP TABLE IF EXISTS banks;
CREATE TABLE banks (
    bank_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    bank_name VARCHAR(150) NOT NULL,
    acronym VARCHAR(20) NOT NULL
);

DROP TABLE IF EXISTS bank_accounts;
CREATE TABLE bank_accounts (
    account_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    enterprise_id INT NOT NULL,
    account_number VARCHAR(30),
    iban VARCHAR(30) UNIQUE,
    CONSTRAINT fk_enterprise_account FOREIGN KEY(enterprise_id) REFERENCES enterprises(enterprise_id)
);

DROP TABLE IF EXISTS taxes;
CREATE TABLE taxes (
    tax_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    tax_name VARCHAR(100) NOT NULL,
    tax_value FLOAT
);

DROP TABLE IF EXISTS invoice_type;
CREATE TABLE invoice_type (
    type_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    type_name VARCHAR(100) NOT NULL,
    code VARCHAR(20) UNIQUE
);

DROP TABLE IF EXISTS users;
CREATE TABLE users (
    user_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_name VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(200) NOT NULL,
    active ENUM('Yes', 'No') NOT NULL DEFAULT 'Yes',
    unique_id BINARY(32) DEFAULT (UUID()),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS suppliers;
CREATE TABLE suppliers (
    supplier_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    supplier_name VARCHAR(150),
    supplier_acronym VARCHAR(30),
    identification_number VARCHAR(30) UNIQUE,
    image VARCHAR(100),
    unique_id BINARY(32) DEFAULT (UUID()),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS supplier_contacts;
CREATE TABLE supplier_contacts (
    contact_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    supplier_id INT NOT NULL,
    phone INT NOT NULL UNIQUE,
    email VARCHAR(150) UNIQUE NOT NULL,
    CONSTRAINT fk_supplier_contact FOREIGN KEY(supplier_id) REFERENCES suppliers(supplier_id)
);

DROP TABLE IF EXISTS supplier_address;
CREATE TABLE supplier_address (
    addres_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    supplier_id INT NOT NULL,
    state VARCHAR(100),
    city VARCHAR(100),
    district VARCHAR(100),
    CONSTRAINT fk_supplier_address FOREIGN KEY(supplier_id) REFERENCES suppliers(supplier_id)
);

DROP TABLE IF EXISTS client_type;
CREATE TABLE client_type (
    type_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    type_name VARCHAR(100) NOT NULL,
    code VARCHAR(20) UNIQUE
);

DROP TABLE IF EXISTS clients;
CREATE TABLE clients (
    client_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    client_type_id INT,
    client_name VARCHAR(150),
    identification_number VARCHAR(30) UNIQUE,
    unique_id BINARY(32) DEFAULT (UUID()),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_type_client FOREIGN KEY(client_type_id) REFERENCES client_type(type_id)
);

DROP TABLE IF EXISTS client_contacts;
CREATE TABLE client_contacts (
    contact_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    client_id INT NOT NULL,
    phone INT NOT NULL UNIQUE,
    email VARCHAR(150) UNIQUE NOT NULL,
    CONSTRAINT fk_client_ontact FOREIGN KEY(client_id) REFERENCES clients(client_id)
);

DROP TABLE IF EXISTS category_service;
CREATE TABLE category_service (
    category_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    category_name VARCHAR(100) UNIQUE,
    description VARCHAR(100) UNIQUE,
    unique_id BINARY(32) DEFAULT (UUID()),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS category_product;
CREATE TABLE category_product (
    category_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    category_name VARCHAR(100) UNIQUE,
    description VARCHAR(100) UNIQUE,
    unique_id BINARY(32) DEFAULT (UUID()),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS products;
CREATE TABLE products (
    product_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    supplier_id INT NOT NULL,
    category_id INT NOT NULL,
    code VARCHAR(30) UNIQUE NOT NULL,
    product_name VARCHAR(100) NOT NULL,
    unit_price FLOAT NOT NULL DEFAULT 0,
    quantity_per_unit INT DEFAULT 1,
    discount FLOAT DEFAULT 0,
    tax_id INT,
    production_date DATE,
    expiration_date DATE,
    image VARCHAR(150),
    description VARCHAR(200),
    discontinued ENUM('Yes', 'No') DEFAULT 'No',
    unique_id BINARY(32) DEFAULT (UUID()),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_category_product FOREIGN KEY(category_id) REFERENCES category_product(category_id),
    CONSTRAINT fk_tax_product FOREIGN KEY(tax_id) REFERENCES taxes(tax_id),
    CONSTRAINT fk_supplier_product FOREIGN KEY(supplier_id) REFERENCES suppliers(supplier_id)
);

DROP TABLE IF EXISTS services;
CREATE TABLE services (
    service_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    supplier_id INT NOT NULL,
    category_id INT NOT NULL,
    code VARCHAR(30) UNIQUE NOT NULL,
    service_name VARCHAR(100) NOT NULL,
    unit_price FLOAT NOT NULL DEFAULT 0,
    discount FLOAT DEFAULT 0,
    tax_id INT,
    image VARCHAR(150),
    description VARCHAR(200),
    discontinued ENUM('Yes', 'No') DEFAULT 'No',
    unique_id BINARY(32) DEFAULT (UUID()),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_category_service FOREIGN KEY(category_id) REFERENCES category_service(category_id),
    CONSTRAINT fk_tax_service FOREIGN KEY(tax_id) REFERENCES taxes(tax_id),
    CONSTRAINT fk_supplier_service FOREIGN KEY(supplier_id) REFERENCES suppliers(supplier_id)
);

DROP TABLE IF EXISTS product_stock;
CREATE TABLE product_stock (
    stock_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    product_id INT NOT NULL,
    store_id INT,
    quantity INT NOT NULL,
    unique_id BINARY(32) DEFAULT (UUID()),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_product_stoke FOREIGN KEY(product_id) REFERENCES products(product_id),
    CONSTRAINT fk_store_stock FOREIGN KEY(store_id) REFERENCES stores(store_id)
);

DROP TABLE IF EXISTS payment_methods;
CREATE TABLE payment_methods (
    payment_method_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    method_name VARCHAR(100) NOT NULL,
    code VARCHAR(20) UNIQUE
);

DROP TABLE IF EXISTS sale_status;
CREATE TABLE sale_status (
    status_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    status_name VARCHAR(100) NOT NULL,
    code VARCHAR(20) UNIQUE
);

CREATE TABLE sale_products (
    sale_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    client_id INT NOT NULL,
    payment_method_id INT NOT NULL,
    code VARCHAR(30) UNIQUE,
    sale_date DATE,
    amount FLOAT NOT NULL,
    total FLOAT NOT NULL,
    sale_change FLOAT NOT NULL,
    user_id INT NOT NULL,
    store_id INT,
    status_id INT,
    unique_id BINARY(32) DEFAULT (UUID()),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_client_sale FOREIGN KEY(client_id) REFERENCES clients(client_id),
    CONSTRAINT fk_store_sale FOREIGN KEY(status_id) REFERENCES stores(store_id),
    CONSTRAINT fk_status_sale FOREIGN KEY(store_id) REFERENCES sale_status(status_id),
    CONSTRAINT fk_payment_method_sale FOREIGN KEY(payment_method_id) REFERENCES payment_methods(payment_method_id),
    CONSTRAINT fk_user_sale FOREIGN KEY(user_id) REFERENCES users(user_id)
);

CREATE TABLE item_sale_product (
    item_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    sale_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL,
    unique_id BINARY(32) DEFAULT (UUID()),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_sale_item_prod FOREIGN KEY(sale_id) REFERENCES sale_products(sale_id),
    CONSTRAINT fk_product_item FOREIGN KEY(product_id) REFERENCES products(product_id)
);

CREATE TABLE sale_services (
    sale_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    client_id INT NOT NULL,
    payment_method_id INT NOT NULL,
    code VARCHAR(30) UNIQUE,
    sale_date DATE,
    amount FLOAT NOT NULL,
    total FLOAT NOT NULL,
    sale_change FLOAT,
    user_id INT NOT NULL,
    store_id INT,
    status_id INT,
    unique_id BINARY(32) DEFAULT (UUID()),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_client_service FOREIGN KEY(client_id) REFERENCES clients(client_id),
    CONSTRAINT fk_store_service FOREIGN KEY(store_id) REFERENCES stores(store_id),
    CONSTRAINT fk_status_service FOREIGN KEY(store_id) REFERENCES sale_status(status_id),
    CONSTRAINT fk_payment_method_service FOREIGN KEY(payment_method_id) REFERENCES payment_methods(payment_method_id),
    CONSTRAINT fk_user_service FOREIGN KEY(user_id) REFERENCES users(user_id)
);

CREATE TABLE item_sale_service (
    item_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    sale_id INT NOT NULL,
    service_id INT NOT NULL,
    quantity INT NOT NULL,
    unique_id BINARY(32) DEFAULT (UUID()),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_sale_item_serv FOREIGN KEY(sale_id) REFERENCES sale_services(sale_id),
    CONSTRAINT fk_service_item FOREIGN KEY(service_id) REFERENCES services(service_id)
);



`
}




func (db *DatabaseSample) createSalesTablesPostgreSQL() string {
return `


CREATE TYPE TYPE_STATUS_DISCONTINUED AS ENUM('Yes', 'No');
CREATE TYPE TYPE_STATUS_ACTIVE AS ENUM('Yes', 'No');

DROP TABLE IF EXISTS enterprises;
CREATE TABLE enterprises (
    enterprise_id SERIAL PRIMARY KEY,
    enterprise_name VARCHAR(30),
    web_site VARCHAR(100),
    acronym VARCHAR(20) NOT NULL,
    identification_number VARCHAR(30),
    logo VARCHAR(100),
    unique_id UUID DEFAULT gen_random_uuid(),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS enterprise_contacts;
CREATE TABLE enterprise_contacts (
    contact_id SERIAL PRIMARY KEY,
    enterprise_id INT NOT NULL,
    phone INT,
    email VARCHAR(100)
);

DROP TABLE IF EXISTS stores;
CREATE TABLE stores (
    store_id SERIAL PRIMARY KEY,
    enterprise_id INT NOT NULL,
    store_name VARCHAR(100) UNIQUE NOT NULL,
    unique_id UUID DEFAULT gen_random_uuid(),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_enterprise_store FOREIGN KEY(enterprise_id) REFERENCES enterprises(enterprise_id)
);

DROP TABLE IF EXISTS store_address;
CREATE TABLE store_address (
    address_id SERIAL PRIMARY KEY,
    store_id INT NOT NULL,
    state VARCHAR(100),
    city VARCHAR(100),
    district VARCHAR(100),
    CONSTRAINT fk_store FOREIGN KEY(store_id) REFERENCES stores(store_id)
);

DROP TABLE IF EXISTS banks;
CREATE TABLE banks (
    bank_id SERIAL PRIMARY KEY,
    bank_name VARCHAR(150) NOT NULL,
    acronym VARCHAR(20) NOT NULL
);

DROP TABLE IF EXISTS bank_accounts;
CREATE TABLE bank_accounts (
    account_id SERIAL PRIMARY KEY,
    enterprise_id INT NOT NULL,
    account_number VARCHAR(30),
    iban VARCHAR(30) UNIQUE,
    CONSTRAINT fk_enterprise_account FOREIGN KEY(enterprise_id) REFERENCES enterprises(enterprise_id)
);

DROP TABLE IF EXISTS taxes;
CREATE TABLE taxes (
    tax_id SERIAL PRIMARY KEY,
    tax_name VARCHAR(100) NOT NULL,
    tax_value FLOAT
);

DROP TABLE IF EXISTS invoice_type;
CREATE TABLE invoice_type (
    type_id SERIAL PRIMARY KEY,
    type_name VARCHAR(100) NOT NULL,
    code VARCHAR(20) UNIQUE
);

DROP TABLE IF EXISTS users;
CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    user_name VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(200) NOT NULL,
    active TYPE_STATUS_ACTIVE NOT NULL DEFAULT 'Yes',
    unique_id UUID DEFAULT gen_random_uuid(),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS suppliers;
CREATE TABLE suppliers (
    supplier_id SERIAL PRIMARY KEY,
    supplier_name VARCHAR(150),
    supplier_acronym VARCHAR(30),
    identification_number VARCHAR(30) UNIQUE,
    image VARCHAR(100),
    unique_id UUID DEFAULT gen_random_uuid(),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS supplier_contacts;
CREATE TABLE supplier_contacts (
    contact_id SERIAL PRIMARY KEY,
    supplier_id INT NOT NULL,
    phone INT NOT NULL UNIQUE,
    email VARCHAR(150) UNIQUE NOT NULL,
    CONSTRAINT fk_supplier_contact FOREIGN KEY(supplier_id) REFERENCES suppliers(supplier_id)
);

DROP TABLE IF EXISTS supplier_address;
CREATE TABLE supplier_address (
    addres_id SERIAL PRIMARY KEY,
    supplier_id INT NOT NULL,
    state VARCHAR(100),
    city VARCHAR(100),
    district VARCHAR(100),
    CONSTRAINT fk_supplier_address FOREIGN KEY(supplier_id) REFERENCES suppliers(supplier_id)
);

DROP TABLE IF EXISTS client_type;
CREATE TABLE client_type (
    type_id SERIAL PRIMARY KEY,
    type_name VARCHAR(100) NOT NULL,
    code VARCHAR(20) UNIQUE
);

DROP TABLE IF EXISTS clients;
CREATE TABLE clients (
    client_id SERIAL PRIMARY KEY,
    client_type_id INT,
    client_name VARCHAR(150),
    identification_number VARCHAR(30) UNIQUE,
    unique_id UUID DEFAULT gen_random_uuid(),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_type_client FOREIGN KEY(client_type_id) REFERENCES client_type(type_id)
);

DROP TABLE IF EXISTS client_contacts;
CREATE TABLE client_contacts (
    contact_id SERIAL PRIMARY KEY,
    client_id INT NOT NULL,
    phone INT NOT NULL UNIQUE,
    email VARCHAR(150) UNIQUE NOT NULL,
    CONSTRAINT fk_client_ontact FOREIGN KEY(client_id) REFERENCES clients(client_id)
);

DROP TABLE IF EXISTS category_service;
CREATE TABLE category_service (
    category_id SERIAL PRIMARY KEY,
    category_name VARCHAR(100) UNIQUE,
    description VARCHAR(100) UNIQUE,
    unique_id UUID DEFAULT gen_random_uuid(),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS category_product;
CREATE TABLE category_product (
    category_id SERIAL PRIMARY KEY,
    category_name VARCHAR(100) UNIQUE,
    description VARCHAR(100) UNIQUE,
    unique_id UUID DEFAULT gen_random_uuid(),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS products;
CREATE TABLE products (
    product_id SERIAL PRIMARY KEY,
    supplier_id INT NOT NULL,
    category_id INT NOT NULL,
    code VARCHAR(30) UNIQUE NOT NULL,
    product_name VARCHAR(100) NOT NULL,
    unit_price FLOAT NOT NULL DEFAULT 0,
    quantity_per_unit INT DEFAULT 1,
    discount FLOAT DEFAULT 0,
    tax_id INT,
    production_date DATE,
    expiration_date DATE,
    image VARCHAR(150),
    description VARCHAR(200),
    discontinued TYPE_STATUS_DISCONTINUED DEFAULT 'No',
    unique_id UUID DEFAULT gen_random_uuid(),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_category_product FOREIGN KEY(category_id) REFERENCES category_product(category_id),
    CONSTRAINT fk_tax_product FOREIGN KEY(tax_id) REFERENCES taxes(tax_id),
    CONSTRAINT fk_supplier_product FOREIGN KEY(supplier_id) REFERENCES suppliers(supplier_id)
);

DROP TABLE IF EXISTS services;
CREATE TABLE services (
    service_id SERIAL PRIMARY KEY,
    supplier_id INT NOT NULL,
    category_id INT NOT NULL,
    code VARCHAR(30) UNIQUE NOT NULL,
    service_name VARCHAR(100) NOT NULL,
    unit_price FLOAT NOT NULL DEFAULT 0,
    discount FLOAT DEFAULT 0,
    tax_id INT,
    image VARCHAR(150),
    description VARCHAR(200),
    discontinued TYPE_STATUS_DISCONTINUED DEFAULT 'No',
    unique_id UUID DEFAULT gen_random_uuid(),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_category_service FOREIGN KEY(category_id) REFERENCES category_service(category_id),
    CONSTRAINT fk_tax_service FOREIGN KEY(tax_id) REFERENCES taxes(tax_id),
    CONSTRAINT fk_supplier_service FOREIGN KEY(supplier_id) REFERENCES suppliers(supplier_id)
);

DROP TABLE IF EXISTS product_stock;
CREATE TABLE product_stock (
    stock_id SERIAL PRIMARY KEY,
    product_id INT NOT NULL,
    store_id INT,
    quantity INT NOT NULL,
    unique_id UUID DEFAULT gen_random_uuid(),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_product_stoke FOREIGN KEY(product_id) REFERENCES products(product_id),
    CONSTRAINT fk_store_stock FOREIGN KEY(store_id) REFERENCES stores(store_id)
);

DROP TABLE IF EXISTS payment_methods;
CREATE TABLE payment_methods (
    payment_method_id SERIAL PRIMARY KEY,
    method_name VARCHAR(100) NOT NULL,
    code VARCHAR(20) UNIQUE
);

DROP TABLE IF EXISTS sale_status;
CREATE TABLE sale_status (
    status_id SERIAL PRIMARY KEY,
    status_name VARCHAR(100) NOT NULL,
    code VARCHAR(20) UNIQUE
);

CREATE TABLE sale_products (
    sale_id SERIAL PRIMARY KEY,
    client_id INT NOT NULL,
    payment_method_id INT NOT NULL,
    code VARCHAR(30) UNIQUE,
    sale_date DATE,
    amount FLOAT NOT NULL,
    total FLOAT NOT NULL,
    sale_change FLOAT NOT NULL,
    user_id INT NOT NULL,
    store_id INT,
    status_id INT,
    unique_id UUID DEFAULT gen_random_uuid(),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_client_sale FOREIGN KEY(client_id) REFERENCES clients(client_id),
    CONSTRAINT fk_store_sale FOREIGN KEY(status_id) REFERENCES stores(store_id),
    CONSTRAINT fk_status_sale FOREIGN KEY(store_id) REFERENCES sale_status(status_id),
    CONSTRAINT fk_payment_method_sale FOREIGN KEY(payment_method_id) REFERENCES payment_methods(payment_method_id),
    CONSTRAINT fk_user_sale FOREIGN KEY(user_id) REFERENCES users(user_id)
);

CREATE TABLE item_sale_product (
    item_id SERIAL PRIMARY KEY,
    sale_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL,
    unique_id UUID DEFAULT gen_random_uuid(),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_sale_item_prod FOREIGN KEY(sale_id) REFERENCES sale_products(sale_id),
    CONSTRAINT fk_product_item FOREIGN KEY(product_id) REFERENCES products(product_id)
);

CREATE TABLE sale_services (
    sale_id SERIAL PRIMARY KEY,
    client_id INT NOT NULL,
    payment_method_id INT NOT NULL,
    code VARCHAR(30) UNIQUE,
    sale_date DATE,
    amount FLOAT NOT NULL,
    total FLOAT NOT NULL,
    sale_change FLOAT,
    user_id INT NOT NULL,
    store_id INT,
    status_id INT,
    unique_id UUID DEFAULT gen_random_uuid(),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_client_service FOREIGN KEY(client_id) REFERENCES clients(client_id),
    CONSTRAINT fk_store_service FOREIGN KEY(store_id) REFERENCES stores(store_id),
    CONSTRAINT fk_status_service FOREIGN KEY(store_id) REFERENCES sale_status(status_id),
    CONSTRAINT fk_payment_method_service FOREIGN KEY(payment_method_id) REFERENCES payment_methods(payment_method_id),
    CONSTRAINT fk_user_service FOREIGN KEY(user_id) REFERENCES users(user_id)
);

CREATE TABLE item_sale_service (
    item_id SERIAL PRIMARY KEY,
    sale_id INT NOT NULL,
    service_id INT NOT NULL,
    quantity INT NOT NULL,
    unique_id UUID DEFAULT gen_random_uuid(),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_sale_item_serv FOREIGN KEY(sale_id) REFERENCES sale_services(sale_id),
    CONSTRAINT fk_service_item FOREIGN KEY(service_id) REFERENCES services(service_id)
);


`
}



func (db *DatabaseSample) createSalesProceduresMySQL() string {
return `

-- CREATE PROCEDURES
-- Procedure sp_decrease_product_stock
DROP PROCEDURE IF EXISTS sp_decrease_product_stock;
DELIMITER $$
CREATE PROCEDURE sp_decrease_product_stock(IN param_qty INT, IN param_id INT)
BEGIN
    START TRANSACTION;
    IF param_qty >= 0 THEN  -- CHECK IF quantity less than 0
        UPDATE product_stock 
        SET quantity = quantity - param_qty
        WHERE product_id = param_id;
        COMMIT;
    ELSE
        ROLLBACK; -- RollBack Transaction
        SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = "Quantity must be greeather or equal than 0!";  -- Raise Exception
    END IF;
END $$
DELIMITER ;

-- Procedure sp_increase_product_stock
DROP PROCEDURE IF EXISTS sp_increase_product_stock;
DELIMITER $$
CREATE PROCEDURE sp_increase_product_stock(IN param_qty INT, IN param_id INT)
BEGIN
    START TRANSACTION;
    IF param_qty >= 0 THEN  -- CHECK IF quantity less than 0
        UPDATE product_stock 
        SET quantity = quantity + param_qty
        WHERE product_id = param_id;
        COMMIT;
    ELSE
        ROLLBACK; -- RollBack Transaction
        SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = "Quantity must be greeather or equal than 0!";  -- Raise Exception
    END IF;
END $$
DELIMITER ;


`
}



func (db *DatabaseSample) createSalesProceduresPostgreSQL() string {
return `
-- CREATE PROCEDURES
-- Procedure sp_decrease_product_stock
DROP PROCEDURE IF EXISTS sp_decrease_product_stock;
CREATE PROCEDURE sp_decrease_product_stock(IN param_qty INT, IN param_id INT)
LANGUAGE plpgsql 
AS $$
BEGIN
    IF param_qty >= 0 THEN  -- CHECK IF quantity less than 0
        UPDATE product_stock 
        SET quantity = quantity - param_qty
        WHERE product_id = param_id;
        COMMIT;
    ELSE
        ROLLBACK; -- RollBack Transaction
        RAISE EXCEPTION 'Quantity must be greeather or equal than 0!';  -- Raise Exception
    END IF;
END;
$$;

-- Procedure sp_increase_product_stock
DROP PROCEDURE IF EXISTS sp_increase_product_stock;
CREATE PROCEDURE sp_increase_product_stock(IN param_qty INT, IN param_id INT)
LANGUAGE plpgsql 
AS $$
BEGIN
    IF param_qty >= 0 THEN  -- CHECK IF quantity less than 0
        UPDATE product_stock 
        SET quantity = quantity - param_qty
        WHERE product_id = param_id;
        COMMIT;
    ELSE
        ROLLBACK; -- RollBack Transaction
        RAISE EXCEPTION 'Quantity must be greeather or equal than 0!';  -- Raise Exception
    END IF;
END;
$$;

`
}


func (db *DatabaseSample) createViewsSales() string {
return `
-- VIEWS 
-- view_product_data
DROP VIEW IF EXISTS view_product_data;
CREATE VIEW view_product_data AS 
SELECT pr.product_id, pr.unique_id,
    pr.product_name, pr.code,
    pr.unit_price, pr.quantity_per_unit,
    pr.discount, pr.production_date,
    pr.expiration_date, pr.image,
    pr.description, pr.discontinued,
    pr.created_at, pr.updated_at,
    ps.stock_id, ps.quantity,
    ca.category_id, ca.category_name,
    su.supplier_id, su.supplier_name,
    su.identification_number,
    ta.tax_id, ta.tax_name,
    ta.tax_value
FROM products pr
JOIN product_stock ps ON(ps.product_id = pr.product_id)
JOIN category_product ca ON(ca.category_id = pr.category_id)
JOIN suppliers su ON(su.supplier_id = pr.supplier_id)
JOIN taxes ta ON(ta.tax_id = pr.tax_id)
ORDER BY pr.created_at DESC;

-- view_service_data
DROP VIEW IF EXISTS view_service_data;
CREATE VIEW view_service_data AS 
SELECT ser.service_id, ser.unique_id,
    ser.service_name, ser.code,
    ser.unit_price, ser.discount, 
    ser.created_at, ser.updated_at,
    ca.category_id, ca.category_name,
    su.supplier_id, su.supplier_name,
    su.identification_number,
    ta.tax_id, ta.tax_name,
    ta.tax_value
FROM services ser
JOIN category_service ca ON(ca.category_id = ser.category_id)
JOIN suppliers su ON(su.supplier_id = ser.supplier_id)
JOIN taxes ta ON(ta.tax_id = ser.tax_id)
ORDER BY ser.created_at DESC;

-- view_sale_product_data
DROP VIEW IF EXISTS view_sale_product_data;
CREATE VIEW view_sale_product_data AS
SELECT sa.sale_id, sa.unique_id,
    sa.code, sa.sale_date,
    sa.amount, sa.total,
    sa.sale_change,
    sta.status_id, sta.status_name,
    pm.payment_method_id, pm.method_name,
    sto.store_id, sto.store_name,
    us.user_id, us.user_name
FROM sale_products sa
JOIN sale_status sta ON(sta.status_id = sa.status_id)
JOIN payment_methods pm ON(pm.payment_method_id = sa.payment_method_id)
JOIN stores sto ON(sto.store_id = sa.store_id)
JOIN users us ON(us.user_id = sa.user_id)
ORDER BY sa.created_at DESC;

-- view_item_sale_product_data
DROP VIEW IF EXISTS view_item_sale_product_data;
CREATE VIEW view_item_sale_product_data AS 
SELECT it.item_id, it.quantity,
    it.created_at,
    pr.product_id, pr.product_name,
    pr.code, pr.unit_price,
    pr.quantity_per_unit,
    sa.sale_id, sa.total,
    sa.sale_date
FROM item_sale_product it
JOIN sale_products sa ON(sa.sale_id = it.sale_id)
JOIN products pr ON(pr.product_id = it.product_id)
ORDER BY it.created_at DESC;

-- view_sale_service_data
DROP VIEW IF EXISTS view_sale_service_data;
CREATE VIEW view_sale_service_data AS
SELECT sa.sale_id, sa.unique_id,
    sa.code, sa.sale_date,
    sa.amount, sa.total,
    sa.sale_change,
    sta.status_id, sta.status_name,
    pm.payment_method_id, pm.method_name,
    cl.client_id, cl.client_name,
    sto.store_id, sto.store_name,
    us.user_id, us.user_name
FROM sale_services sa
JOIN sale_status sta ON(sta.status_id = sa.status_id)
JOIN clients cl ON(cl.client_id = sa.client_id)
JOIN payment_methods pm ON(pm.payment_method_id = sa.payment_method_id)
JOIN stores sto ON(sto.store_id = sa.store_id)
JOIN users us ON(us.user_id = sa.user_id)
ORDER BY sa.created_at DESC;

-- view_item_sale_service_data
DROP VIEW IF EXISTS view_item_sale_service_data;
CREATE VIEW view_item_sale_service_data AS 
SELECT it.item_id, it.quantity,
    it.created_at,
    ser.service_id, ser.service_name,
    ser.code, ser.unit_price,
    sa.sale_id, sa.total,
    sa.sale_date
FROM item_sale_service it
JOIN sale_services sa ON(sa.sale_id = it.sale_id)
JOIN services ser ON(ser.service_id = it.service_id)
ORDER BY it.created_at DESC;


`
}




func (db *DatabaseSample) insertsSales() string {
return `


-- INSERTS
INSERT INTO payment_methods (method_name, code) VALUES ('Numerário', 'numerario');
INSERT INTO payment_methods (method_name, code) VALUES ('TPA', 'tpa');
INSERT INTO payment_methods (method_name, code) VALUES ('Transferência', 'transferencia');

INSERT INTO category_product (category_name, description) VALUES ('Materias de Escritório', 'Relacionado a Escritório');
INSERT INTO category_product (category_name, description) VALUES ('Electronicos', 'Relacionado a Electrónicos');
INSERT INTO category_product (category_name, description) VALUES ('Acessórios', 'Relacionado a Acessórios');
INSERT INTO category_product (category_name, description) VALUES ('Electro-domésticos', 'Relacionado a Electrodomésticos');
INSERT INTO category_product (category_name, description) VALUES ('Alimentação', 'Relacionado a Alimentação');
INSERT INTO category_product (category_name, description) VALUES ('Outro', 'Outros Produtos');

INSERT INTO category_service (category_name, description) VALUES ('Software', 'Relacionado a Softwares');
INSERT INTO category_service (category_name, description) VALUES ('Hospedagem', 'Relacionado aHospedagen');
INSERT INTO category_service (category_name, description) VALUES ('Acessoria', 'Relacionado a Acessoria');
INSERT INTO category_service (category_name, description) VALUES ('Consultoria', 'Relacionado a Consultoria');
INSERT INTO category_service (category_name, description) VALUES ('Formação', 'Relacionado a Formação');
INSERT INTO category_service (category_name, description) VALUES ('Outros', 'Outros Produtos');

INSERT INTO suppliers (supplier_name, identification_number) VALUES ('Supplier 1', 'SUP-001');
INSERT INTO suppliers (supplier_name, identification_number) VALUES ('Supplier 2', 'SUP-002');
INSERT INTO suppliers (supplier_name, identification_number) VALUES ('Supplier 3', 'SUP-003');
INSERT INTO suppliers (supplier_name, identification_number) VALUES ('Supplier 4', 'SUP-004');
INSERT INTO suppliers (supplier_name, identification_number) VALUES ('Supplier 5', 'SUP-005');

INSERT INTO products (category_id, supplier_id, product_name, code, unit_price) VALUES (1, 1, 'Book', 'P001', 100);
INSERT INTO products (category_id, supplier_id, product_name, code, unit_price) VALUES (1, 2, 'Pencil', 'P002', 20);
INSERT INTO products (category_id, supplier_id, product_name, code, unit_price) VALUES (2, 2, 'Computer', 'P003', 6000);
INSERT INTO products (category_id, supplier_id, product_name, code, unit_price) VALUES (2, 3, 'Dictionary', 'P004', 50);
INSERT INTO products (category_id, supplier_id, product_name, code, unit_price) VALUES (2, 3, 'Eraser', 'P005', 5);
INSERT INTO products (category_id, supplier_id, product_name, code, unit_price) VALUES (3, 2, 'Phone', 'P006', 4000);

INSERT INTO product_stock (product_id, quantity) VALUES (1, 100);
INSERT INTO product_stock (product_id, quantity) VALUES (2, 300);
INSERT INTO product_stock (product_id, quantity) VALUES (3, 25);
INSERT INTO product_stock (product_id, quantity) VALUES (4, 200);
INSERT INTO product_stock (product_id, quantity) VALUES (5, 700);
INSERT INTO product_stock (product_id, quantity) VALUES (6, 400);

INSERT INTO taxes (tax_name, tax_value) VALUES ('IVA - Exclusão', 0.00);
INSERT INTO taxes (tax_name, tax_value) VALUES ('IVA - Inclusão', 14.00);
INSERT INTO taxes (tax_name, tax_value) VALUES ('IVA - Simplificado', 7.00);
INSERT INTO taxes (tax_name, tax_value) VALUES ('Produtos constantes da Tabela Anexa à Lei nº 32/21', 7.00);
INSERT INTO taxes (tax_name, tax_value) VALUES ('Produtos da Cesta básica e insumos agrícolas', 5.00);
INSERT INTO taxes (tax_name, tax_value) VALUES ('Província de Cabinda - Operações internas e importação', 2.00);
INSERT INTO taxes (tax_name, tax_value) VALUES ('Retenção Na Fonte', 6.50);
INSERT INTO taxes (tax_name, tax_value) VALUES ('Sector de Hotelaria e Restauração', 7.00);

INSERT INTO banks (bank_name, acronym) VALUES ('Banco Angolano de Investimentos', 'BAI');
INSERT INTO banks (bank_name, acronym) VALUES ('Banco de Poupança e Crédito', 'BPC');
INSERT INTO banks (bank_name, acronym) VALUES ('Banco de Fomento Angola', 'BFA');
INSERT INTO banks (bank_name, acronym) VALUES ('Banco Económico', 'BE');
INSERT INTO banks (bank_name, acronym) VALUES ('Banco de Comércio e Indústria', 'BCI');
INSERT INTO banks (bank_name, acronym) VALUES ('Banco de Indústria Crédito', 'BIC');
INSERT INTO banks (bank_name, acronym) VALUES ('Banco Millenium Atlântico', 'BMA');

INSERT INTO sale_status (status_name, code) VALUES ('Pendente', 'pendente');
INSERT INTO sale_status (status_name, code) VALUES ('Paga', 'paga');
INSERT INTO sale_status (status_name, code) VALUES ('Reembolsada', 'reembolsada');
INSERT INTO sale_status (status_name, code) VALUES ('Incompleta', 'incompleta');
INSERT INTO sale_status (status_name, code) VALUES ('Cancelada', 'cancelada');

INSERT INTO invoice_type (type_name, code) VALUES ('Factura Crédito', 'factura-credito');
INSERT INTO invoice_type (type_name, code) VALUES ('Factura Recibo', 'factura-recibo');
INSERT INTO invoice_type (type_name, code) VALUES ('Nota de Crédito', 'nota-credito');
INSERT INTO invoice_type (type_name, code) VALUES ('Nota de Débito', 'nota-debito');
INSERT INTO invoice_type (type_name, code) VALUES ('Nota de Cobrança', 'nota-cobranca');
INSERT INTO invoice_type (type_name, code) VALUES ('Proforma', 'proforma');
INSERT INTO invoice_type (type_name, code) VALUES ('Recibo', 'recibo');

INSERT INTO client_type (type_name, code) VALUES ('Singular', 'singular');
INSERT INTO client_type (type_name, code) VALUES ('Empresa', 'Empresa');
INSERT INTO client_type (type_name, code) VALUES ('Instituição Publica', 'inst-publica');
INSERT INTO client_type (type_name, code) VALUES ('Instituição Privada', 'inst-privada');


`
}