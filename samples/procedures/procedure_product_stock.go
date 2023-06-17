package samples


func (proc *ProcedureSample) GetProcedureProductStockSample(rdb string) string {
	sql := ""
	switch rdb {
	case "mysql":
		sql = proc.createProcedure("sp_decrease_product_stock", "mysql") + proc.createProcedureProductStockMySQL()

	case "postgres":
		sql = proc.createProcedure("sp_decrease_product_stock", "postgres") + proc.createProcedureProductStockPostgreSQL() 
	}
	return sql
}


func (proc *ProcedureSample) createProcedureProductStockPostgreSQL() string {
return `
DROP TABLE IF EXISTS products;
CREATE TABLE products (
    product_id SERIAL PRIMARY KEY,
    product_name VARCHAR(100) NOT NULL,
    code VARCHAR(30) UNIQUE NOT NULL,
    unique_id uuid DEFAULT gen_random_uuid() NOT NULL UNIQUE,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS product_stock;
CREATE TABLE product_stock (
    stock_id SERIAL PRIMARY KEY,
    product_id INT NOT NULL,
    quantity INT NOT NULL,
    CONSTRAINT fk_product FOREIGN KEY(product_id) REFERENCES products(product_id)
);

INSERT INTO products (product_name, code) VALUES ('Book', 'P001');
INSERT INTO products (product_name, code) VALUES ('Pencil', 'P002');
INSERT INTO products (product_name, code) VALUES ('Computer', 'P003');
INSERT INTO products (product_name, code) VALUES ('Dictionary', 'P004');
INSERT INTO products (product_name, code) VALUES ('Eraser', 'P005');

INSERT INTO product_stock (product_id, quantity) VALUES (1, 100);
INSERT INTO product_stock (product_id, quantity) VALUES (2, 300);
INSERT INTO product_stock (product_id, quantity) VALUES (3, 25);
INSERT INTO product_stock (product_id, quantity) VALUES (4, 200);
INSERT INTO product_stock (product_id, quantity) VALUES (5, 700);


-- PROCEDURE
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

`
}


func (proc *ProcedureSample) createProcedureProductStockMySQL() string {
return `
DROP TABLE IF EXISTS products;
CREATE TABLE products (
    product_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    product_name VARCHAR(100) NOT NULL,
    code VARCHAR(30) UNIQUE NOT NULL,
    unique_id BINARY(32) DEFAULT (UUID()),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS product_stock;
CREATE TABLE product_stock (
    stock_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    product_id INT NOT NULL,
    quantity INT NOT NULL,
    CONSTRAINT fk_product FOREIGN KEY(product_id) REFERENCES products(product_id)
);

INSERT INTO products (product_name, code) VALUES ('Book', 'P001');
INSERT INTO products (product_name, code) VALUES ('Pencil', 'P002');
INSERT INTO products (product_name, code) VALUES ('Computer', 'P003');
INSERT INTO products (product_name, code) VALUES ('Dictionary', 'P004');
INSERT INTO products (product_name, code) VALUES ('Eraser', 'P005');

INSERT INTO product_stock (product_id, quantity) VALUES (1, 100);
INSERT INTO product_stock (product_id, quantity) VALUES (2, 300);
INSERT INTO product_stock (product_id, quantity) VALUES (3, 25);
INSERT INTO product_stock (product_id, quantity) VALUES (4, 200);
INSERT INTO product_stock (product_id, quantity) VALUES (5, 700);


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

`
}

