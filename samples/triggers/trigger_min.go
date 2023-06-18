package samples


func (trg* TriggerSample) GetTriggerMinimalSample(rdb string) string {
	sql := ""
	switch rdb {
	case "mysql":
		sql = trg.createTrigger("trg_min", "mysql") + trg.createTriggerMinimalMySQL()
	case "postgres":
		sql = trg.createTrigger("trg_min", "postgres") + trg.createTriggerMinimalPostgreSQL() 
	}
	return sql
}


func (trg* TriggerSample) createTriggerMinimalMySQL() string {
return `

DROP TABLE IF EXISTS products;
CREATE TABLE products (
    product_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    product_name VARCHAR(100),
    code VARCHAR(20) UNIQUE NOT NULL
);

DROP TABLE IF EXISTS product_history;
CREATE TABLE product_history (
    history_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    product_id INT NOT NULL,
    product_name VARCHAR(100),
    operation VARCHAR(50),
    operation_date DATETIME,
    CONSTRAINT fk_product FOREIGN KEY(product_id) REFERENCES products(product_id)
);

INSERT INTO products (product_name, code) VALUES ('Book', 'PR-001');
INSERT INTO products (product_name, code) VALUES ('Pencil', 'PR-002');
INSERT INTO products (product_name, code) VALUES ('Computer', 'PR-003');
INSERT INTO products (product_name, code) VALUES ('Dictionary', 'PR-004');
INSERT INTO products (product_name, code) VALUES ('Eraser', 'PR-005');


DROP TRIGGER IF EXISTS trg_update_status;
DELIMITER $$

CREATE TRIGGER trg_update_status 
AFTER UPDATE ON products
FOR EACH ROW 
BEGIN
    INSERT INTO product_history (product_name, product_id, operation, operation_date) 
    VALUES (OLD.product_name, OLD.product_id, 'UPDATED', CURRENT_TIMESTAMP);
END $$
DELIMITER ;

`
}

func (trg*TriggerSample) createTriggerMinimalPostgreSQL() string {
return `

DROP TABLE IF EXISTS products;
CREATE TABLE products (
    product_id SERIAL PRIMARY KEY,
    product_name VARCHAR(100),
    code VARCHAR(20) UNIQUE NOT NULL
);

DROP TABLE IF EXISTS product_history;
CREATE TABLE product_history (
    history_id SERIAL PRIMARY KEY,
    product_id INT NOT NULL,
    product_name VARCHAR(100),
    operation VARCHAR(50),
    operation_date TIMESTAMP,
    CONSTRAINT fk_product FOREIGN KEY(product_id) REFERENCES products(product_id)
);

INSERT INTO products (product_name, code) VALUES ('Book', 'PR-001');
INSERT INTO products (product_name, code) VALUES ('Pencil', 'PR-002');
INSERT INTO products (product_name, code) VALUES ('Computer', 'PR-003');
INSERT INTO products (product_name, code) VALUES ('Dictionary', 'PR-004');
INSERT INTO products (product_name, code) VALUES ('Eraser', 'PR-005');

CREATE OR REPLACE FUNCTION fun_update_status()
RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO product_history (product_name, product_id, operation, operation_date) 
    VALUES (OLD.product_name, OLD.product_id, 'UPDATED', CURRENT_TIMESTAMP);
    RETURN NEW;
END;
$$ LANGUAGE 'plpgsql';


CREATE OR REPLACE TRIGGER trg_update_status
    BEFORE UPDATE ON products
    FOR EACH ROW
    EXECUTE FUNCTION fun_update_status();


`
}

