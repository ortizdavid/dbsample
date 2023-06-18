package samples


func (view *ViewSample) GetViewProductDataSample(rdb string) string {
	sql := ""
	switch rdb {
	case "mysql":
		sql = view.createView("view_product_data", "mysql") + view.createViewProductDataMySQL()

	case "postgres":
		sql = view.createView("view_product_data", "postgres") + view.createViewProductDataPostgreSQL() 
	}
	return sql
}


func (view *ViewSample) createViewProductDataMySQL() string {
return `
DROP TABLE IF EXISTS suppliers;
CREATE TABLE suppliers (
    supplier_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    supplier_name VARCHAR(150),
    identification_number VARCHAR(30) UNIQUE,
    unique_id BINARY(32) DEFAULT (UUID()),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS categories;
CREATE TABLE categories (
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
    unit_price DOUBLE NOT NULL DEFAULT 0,
    quantity_per_unit INT DEFAULT 1,
    image VARCHAR(150),
    description VARCHAR(200),
    discontinued ENUM('Yes', 'No') DEFAULT 'No',
    unique_id BINARY(32) DEFAULT (UUID()),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_category FOREIGN KEY(category_id) REFERENCES categories(category_id),
    CONSTRAINT fk_supplier FOREIGN KEY(supplier_id) REFERENCES suppliers(supplier_id)
);

DROP TABLE IF EXISTS product_stock;
CREATE TABLE product_stock (
    stock_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    product_id INT NOT NULL,
    quantity INT NOT NULL,
    CONSTRAINT fk_product FOREIGN KEY(product_id) REFERENCES products(product_id)
);


INSERT INTO categories (category_name, description) VALUES ('Office Material', 'Related to Office');
INSERT INTO categories (category_name, description) VALUES ('Electronic', 'Related to Electronic Materials');
INSERT INTO categories (category_name, description) VALUES ('Acessories', 'Related to Acessories');
INSERT INTO categories (category_name, description) VALUES ('Household appliances', 'Related to Household');
INSERT INTO categories (category_name, description) VALUES ('Other', 'Other Products');

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


-- CREATE VIEW
DROP VIEW IF EXISTS view_product_data;

CREATE VIEW view_product_data AS 
SELECT pr.product_id, pr.unique_id,
    pr.product_name, pr.code,
    pr.unit_price, pr.quantity_per_unit,
    pr.created_at, pr.updated_at,
    ps.stock_id, ps.quantity,
    ca.category_id, ca.category_name,
    su.supplier_id, su.supplier_name,
    su.identification_number
FROM products pr
JOIN product_stock ps ON(ps.product_id = pr.product_id)
JOIN categories ca ON(ca.category_id = pr.category_id)
JOIN suppliers su ON(su.supplier_id = pr.supplier_id)
ORDER BY pr.created_at DESC;

-- CALL VIEW 
SELECT * FROM view_product_data;


`
}


func (view *ViewSample) createViewProductDataPostgreSQL() string {
return `
DROP TABLE IF EXISTS suppliers;
CREATE TABLE suppliers (
    supplier_id SERIAL PRIMARY KEY,
    supplier_name VARCHAR(150),
    identification_number VARCHAR(30) UNIQUE,
    unique_id uuid DEFAULT gen_random_uuid(),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS categories;
CREATE TABLE categories (
    category_id SERIAL PRIMARY KEY,
    category_name VARCHAR(100) UNIQUE,
    description VARCHAR(200) UNIQUE,
    unique_id uuid DEFAULT gen_random_uuid(),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

DROP TYPE IF EXISTS TYPE_PRODUCT_STATUS;
CREATE TYPE TYPE_PRODUCT_STATUS AS ENUM('Yes', 'No');

DROP TABLE IF EXISTS products;
CREATE TABLE products (
    product_id SERIAL PRIMARY KEY,
    supplier_id INT NOT NULL,
    category_id INT NOT NULL,
    code VARCHAR(30) UNIQUE NOT NULL,
    product_name VARCHAR(100) NOT NULL,
    unit_price FLOAT NOT NULL DEFAULT 0,
    quantity_per_unit INT DEFAULT 1,
    image VARCHAR(150),
    description VARCHAR(200),
    discontinued TYPE_PRODUCT_STATUS DEFAULT 'No',
    unique_id uuid DEFAULT gen_random_uuid(),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_category FOREIGN KEY(category_id) REFERENCES categories(category_id),
    CONSTRAINT fk_supplier FOREIGN KEY(supplier_id) REFERENCES suppliers(supplier_id)
);

DROP TABLE IF EXISTS product_stock;
CREATE TABLE product_stock (
    stock_id SERIAL PRIMARY KEY,
    product_id INT NOT NULL,
    quantity INT NOT NULL,
    CONSTRAINT fk_product FOREIGN KEY(product_id) REFERENCES products(product_id)
);


INSERT INTO categories (category_name, description) VALUES ('Office Material', 'Related to Office');
INSERT INTO categories (category_name, description) VALUES ('Electronic', 'Related to Electronic Materials');
INSERT INTO categories (category_name, description) VALUES ('Acessories', 'Related to Acessories');
INSERT INTO categories (category_name, description) VALUES ('Household appliances', 'Related to Household');
INSERT INTO categories (category_name, description) VALUES ('Other', 'Other Products');

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


-- CREATE VIEW
DROP VIEW IF EXISTS view_product_data;

CREATE VIEW view_product_data AS 
SELECT pr.product_id, pr.unique_id,
    pr.product_name, pr.code,
    pr.unit_price, pr.quantity_per_unit,
    pr.created_at, pr.updated_at,
    ps.stock_id, ps.quantity,
    ca.category_id, ca.category_name,
    su.supplier_id, su.supplier_name,
    su.identification_number
FROM products pr
JOIN product_stock ps ON(ps.product_id = pr.product_id)
JOIN categories ca ON(ca.category_id = pr.category_id)
JOIN suppliers su ON(su.supplier_id = pr.supplier_id)
ORDER BY pr.created_at DESC;

-- CALL VIEW 
SELECT * FROM view_product_data;

`
}

