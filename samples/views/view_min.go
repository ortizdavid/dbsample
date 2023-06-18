package samples


func (view *ViewSample) GetViewMinimalSample(rdb string) string {
	sql := ""
	switch rdb {
	case "mysql":
		sql = view.createView("view_minimal", "mysql") + view.createViewMinimalMySQL()

	case "postgres":
		sql = view.createView("view_minimal", "postgres") + view.createViewMinimalPostgreSQL() 
	}
	return sql
}


func (view *ViewSample) createViewMinimalMySQL() string {
return `
DROP TABLE IF EXISTS clients;
CREATE TABLE clients (
    client_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    identification_number VARCHAR(30) UNIQUE,
    unique_id BINARY(32) DEFAULT (UUID()),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS client_address;
CREATE TABLE client_address (
    addres_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    client_id INT NOT NULL,
    state VARCHAR(100),
    city VARCHAR(100),
    district VARCHAR(100),
    unique_id BINARY(32) DEFAULT (UUID()),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_client_address FOREIGN KEY(client_id) REFERENCES clients(client_id)
);

-- INSERTS
INSERT INTO clients (first_name, last_name, identification_number) VALUES ('John', 'Doe', 'IDN-001');
INSERT INTO clients (first_name, last_name, identification_number) VALUES ('John', 'Wick', 'IDN-002');
INSERT INTO clients (first_name, last_name, identification_number) VALUES ('Nikola', 'Tesla', 'IDN-003');
INSERT INTO clients (first_name, last_name, identification_number) VALUES ('Noah', 'Trevor', 'IDN-004');
INSERT INTO clients (first_name, last_name, identification_number) VALUES ('Anna', 'Mary', 'IDN-005');

INSERT INTO client_address (client_id, state, city, district) VALUES (1, 'Luanda', 'Talatona', 'District 1');
INSERT INTO client_address (client_id, state, city, district) VALUES (2, 'Benguela', 'Sumbe', 'District 2');
INSERT INTO client_address (client_id, state, city, district) VALUES (3, 'Luanda', 'Viana', 'District 3');
INSERT INTO client_address (client_id, state, city, district) VALUES (4, 'São Paulo', 'Cidade 002', 'District 4');
INSERT INTO client_address (client_id, state, city, district) VALUES (5, 'New York', 'Harlem', 'District 5');

-- VIEW view_client_data
DROP VIEW IF EXISTS view_client_data;
CREATE VIEW view_client_data AS 
SELECT cl.client_id, cl.unique_id,
    cl.first_name, cl.last_name,
    cl.identification_number,
    cl.created_at, cl.updated_at,
    ad.addres_id, ad.state, 
    ad.city, ad.district
FROM clients cl
JOIN client_address ad ON(ad.client_id = cl.client_id)
ORDER BY cl.created_at DESC;

-- CALL VIEW 
SELECT * FROM view_client_data;


`
}


func (view *ViewSample) createViewMinimalPostgreSQL() string {
return `
DROP TABLE IF EXISTS clients;
CREATE TABLE clients (
    client_id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    identification_number VARCHAR(30) UNIQUE,
    unique_id uuid DEFAULT gen_random_uuid(),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS client_address;
CREATE TABLE client_address (
    addres_id SERIAL PRIMARY KEY,
    client_id INT NOT NULL,
    state VARCHAR(100),
    city VARCHAR(100),
    district VARCHAR(100),
    unique_id uuid DEFAULT gen_random_uuid(),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_client_address FOREIGN KEY(client_id) REFERENCES clients(client_id)
);

-- INSERTS

-- INSERTS
INSERT INTO clients (first_name, last_name, identification_number) VALUES ('John', 'Doe', 'IDN-001');
INSERT INTO clients (first_name, last_name, identification_number) VALUES ('John', 'Wick', 'IDN-002');
INSERT INTO clients (first_name, last_name, identification_number) VALUES ('Nikola', 'Tesla', 'IDN-003');
INSERT INTO clients (first_name, last_name, identification_number) VALUES ('Noah', 'Trevor', 'IDN-004');
INSERT INTO clients (first_name, last_name, identification_number) VALUES ('Anna', 'Mary', 'IDN-005');

INSERT INTO client_address (client_id, state, city, district) VALUES (1, 'Luanda', 'Talatona', 'District 1');
INSERT INTO client_address (client_id, state, city, district) VALUES (2, 'Benguela', 'Sumbe', 'District 2');
INSERT INTO client_address (client_id, state, city, district) VALUES (3, 'Luanda', 'Viana', 'District 3');
INSERT INTO client_address (client_id, state, city, district) VALUES (4, 'São Paulo', 'Cidade 002', 'District 4');
INSERT INTO client_address (client_id, state, city, district) VALUES (5, 'New York', 'Harlem', 'District 5');

-- VIEW view_client_data
DROP VIEW IF EXISTS view_client_data;
CREATE VIEW view_client_data AS 
SELECT cl.client_id, cl.unique_id,
    cl.first_name, cl.last_name,
    cl.identification_number,
    cl.created_at, cl.updated_at,
    ad.addres_id, ad.state, 
    ad.city, ad.district
FROM clients cl
JOIN client_address ad ON(ad.client_id = cl.client_id)
ORDER BY cl.created_at DESC;

-- CALL VIEW 
SELECT * FROM view_client_data;


`
}

