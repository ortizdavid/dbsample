package samples

func (db *DatabaseSample) GetPeopleSample(rdb string) string {
	sql := ""
	switch rdb {
	case "mysql":
		sql = db.createDatabase("db_people", "mysql") + db.createPeopleTablesMySQL() + db.insertsPeople()

	case "postgres":
		sql = db.createDatabase("db_people", "postgres") + db.createPeopleTablesPostgreSQL() + db.insertsPeople()
	}
	return sql
}


func (db *DatabaseSample) createPeopleTablesMySQL() string {
return `

DROP TABLE IF EXISTS marital_statuses;
CREATE TABLE IF NOT EXISTS marital_statuses (
	status_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	status_name VARCHAR(100) NOT NULL UNIQUE,
	code VARCHAR(20) UNIQUE
);

DROP TABLE IF EXISTS identification_types;
CREATE TABLE IF NOT EXISTS identification_types (
	type_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	type_name VARCHAR(100) NOT NULL UNIQUE,
	code VARCHAR(20) UNIQUE
);

DROP TABLE IF EXISTS document_types;
CREATE TABLE IF NOT EXISTS document_types (
	type_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	type_name VARCHAR(100) NOT NULL UNIQUE,
	code VARCHAR(20) UNIQUE
);

DROP TABLE IF EXISTS contact_types;
CREATE TABLE IF NOT EXISTS contact_types (
	type_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    type_name VARCHAR(30),
    code VARCHAR(20) UNIQUE
);

DROP TABLE IF EXISTS continents;
CREATE TABLE continents (
	continent_id INT NOT NULL PRIMARY KEY,
	continent_name VARCHAR(100) NOT NULL UNIQUE
);

DROP TABLE IF EXISTS countries;
CREATE TABLE countries (
	country_id INT NOT NULL PRIMARY KEY,
	continent_id INT NOT NULL,
	country_name VARCHAR(150) NOT NULL UNIQUE,
    CONSTRAINT fk_continent FOREIGN KEY(continent_id) REFERENCES continents(continent_id)
);

DROP TABLE IF EXISTS person;
CREATE TABLE IF NOT EXISTS person (
	person_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	identification_type_id INT,
	marital_status_id INT,
	first_name VARCHAR(100) NOT NULL,
	last_name VARCHAR(100) NOT NULL,
	birth_date DATE,
	gender ENUM('Masculino', 'Feminino'),
	identification_number VARCHAR(30),
	unique_id BINARY(32) DEFAULT (UUID()),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT fk_identification_person FOREIGN KEY(identification_type_id) REFERENCES identification_types(type_id),
	CONSTRAINT fk_marital_statuses FOREIGN KEY(marital_status_id) REFERENCES marital_statuses(status_id)
);

DROP TABLE IF EXISTS contacts;
CREATE TABLE IF NOT EXISTS contacts (
	contact_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    person_id INT NOT NULL,
    contact_type_id INT NOT NULL,
	email VARCHAR(150) UNIQUE,
    phone INT UNIQUE,
    unique_id BINARY(32) DEFAULT (UUID()),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT fk_person FOREIGN KEY(person_id) REFERENCES person(person_id),
    CONSTRAINT fk_contact_type FOREIGN KEY(contact_type_id) REFERENCES contact_types(type_id)
);

DROP TABLE IF EXISTS addresses;
CREATE TABLE IF NOT EXISTS adresses (
	address_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    person_id INT NOT NULL,
    country_id INT NOT NULL,
    state VARCHAR(150),
	city VARCHAR(150),
    district VARCHAR(150),
    postal_code VARCHAR(20),
    unique_id BINARY(32) DEFAULT (UUID()),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT fk_person FOREIGN KEY(person_id) REFERENCES person(person_id),
    CONSTRAINT fk_country FOREIGN KEY(country_id) REFERENCES countries(country_id)
);

DROP TABLE IF EXISTS documents;
CREATE TABLE IF NOT EXISTS documents (
	document_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    person_id INT NOT NULL,
    document_type_id INT NOT NULL,
	file_name VARCHAR(150),
    unique_id BINARY(32) DEFAULT (UUID()),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT fk_person FOREIGN KEY(person_id) REFERENCES person(person_id),
    CONSTRAINT fk_document_type FOREIGN KEY(document_type_id) REFERENCES document_types(type_id)
);

`
}


func (db *DatabaseSample) createPeopleTablesPostgreSQL() string {
return `

DROP TABLE IF EXISTS marital_statuses;
CREATE TABLE IF NOT EXISTS marital_statuses (
	status_id SERIAL NOT NULL PRIMARY KEY,
	status_name VARCHAR(100) NOT NULL UNIQUE,
	code VARCHAR(20) UNIQUE
);

DROP TABLE IF EXISTS identification_types;
CREATE TABLE IF NOT EXISTS identification_types (
	type_id SERIAL NOT NULL PRIMARY KEY,
	type_name VARCHAR(100) NOT NULL UNIQUE,
	code VARCHAR(20) UNIQUE
);

DROP TABLE IF EXISTS document_types;
CREATE TABLE IF NOT EXISTS document_types (
	type_id SERIAL NOT NULL PRIMARY KEY,
	type_name VARCHAR(100) NOT NULL UNIQUE,
	code VARCHAR(20) UNIQUE
);

DROP TABLE IF EXISTS contact_types;
CREATE TABLE IF NOT EXISTS contact_types (
	type_id SERIAL NOT NULL PRIMARY KEY,
    type_name VARCHAR(30),
    code VARCHAR(20) UNIQUE
);

DROP TABLE IF EXISTS continents;
CREATE TABLE continents (
	continent_id INT NOT NULL PRIMARY KEY,
	continent_name VARCHAR(100) NOT NULL UNIQUE
);

DROP TABLE IF EXISTS countries;
CREATE TABLE countries (
	country_id INT NOT NULL PRIMARY KEY,
	continent_id INT NOT NULL,
	country_name VARCHAR(150) NOT NULL UNIQUE,
    CONSTRAINT fk_continent FOREIGN KEY(continent_id) REFERENCES continents(continent_id)
);

DROP TYPE IF EXISTS TYPE_GENDER;
CREATE TYPE TYPE_GENDER AS ENUM('Masculino', 'Feminino');

DROP TABLE IF EXISTS person;
CREATE TABLE IF NOT EXISTS person (
	person_id SERIAL NOT NULL PRIMARY KEY,
	identification_type_id INT,
	marital_status_id INT,
	first_name VARCHAR(100) NOT NULL,
	last_name VARCHAR(100) NOT NULL,
	birth_date DATE DEFAULT CURRENT_DATE,
	gender TYPE_GENDER,
	identification_number VARCHAR(30),
	unique_id uuid DEFAULT gen_random_uuid() NOT NULL UNIQUE,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT fk_identification_person FOREIGN KEY(identification_type_id) REFERENCES identification_types(type_id),
	CONSTRAINT fk_marital_statuses FOREIGN KEY(marital_status_id) REFERENCES marital_statuses(status_id)
);

DROP TABLE IF EXISTS contacts;
CREATE TABLE IF NOT EXISTS contacts (
	contact_id SERIAL NOT NULL PRIMARY KEY,
    person_id INT NOT NULL,
    contact_type_id INT NOT NULL,
	email VARCHAR(150) UNIQUE,
    phone INT UNIQUE,
    unique_id uuid DEFAULT gen_random_uuid() NOT NULL UNIQUE,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT fk_person FOREIGN KEY(person_id) REFERENCES person(person_id),
    CONSTRAINT fk_contact_type FOREIGN KEY(contact_type_id) REFERENCES contact_types(type_id)
);

DROP TABLE IF EXISTS addresses;
CREATE TABLE IF NOT EXISTS adresses (
	address_id SERIAL NOT NULL PRIMARY KEY,
    person_id INT NOT NULL,
    country_id INT NOT NULL,
    state VARCHAR(150),
	city VARCHAR(150),
    district VARCHAR(150),
    postal_code VARCHAR(20),
    unique_id uuid DEFAULT gen_random_uuid() NOT NULL UNIQUE,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT fk_person FOREIGN KEY(person_id) REFERENCES person(person_id),
    CONSTRAINT fk_country FOREIGN KEY(country_id) REFERENCES countries(country_id)
);

DROP TABLE IF EXISTS documents;
CREATE TABLE IF NOT EXISTS documents (
	document_id SERIAL NOT NULL PRIMARY KEY,
    person_id INT NOT NULL,
    document_type_id INT NOT NULL,
	file_name VARCHAR(150),
    unique_id uuid DEFAULT gen_random_uuid() NOT NULL UNIQUE,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT fk_person FOREIGN KEY(person_id) REFERENCES person(person_id),
    CONSTRAINT fk_document_type FOREIGN KEY(document_type_id) REFERENCES document_types(type_id)
);

`
}


func (db *DatabaseSample) insertsPeople() string {
return `

INSERT INTO marital_statuses (code, status_name) VALUES ('solteiro', 'Solteiro');  
INSERT INTO marital_statuses (code, status_name) VALUES ('casado_com_registo', 'Casado (com registo)');  
INSERT INTO marital_statuses (code, status_name) VALUES ('casado_sem_registo', 'Casado (sem registo)');  
INSERT INTO marital_statuses (code, status_name) VALUES ('divorciado', 'Divorciado');  
INSERT INTO marital_statuses (code, status_name) VALUES ('separado', 'Separado');  
INSERT INTO marital_statuses (code, status_name) VALUES ('viuvo', 'Viúvo');  
INSERT INTO marital_statuses (code, status_name) VALUES ('outro', 'Outro');  

INSERT INTO identification_types (code, type_name) VALUES ('bi', 'Bilhete de Identidade');
INSERT INTO identification_types (code, type_name) VALUES ('passaporte','Passaporte');
INSERT INTO identification_types (code, type_name) VALUES ('residente', 'Cartão de Residente');
INSERT INTO identification_types (code, type_name) VALUES ('bi-cverde', 'Bilhete de Identidade (Cabo Verde)');
INSERT INTO identification_types (code, type_name) VALUES ('autorizacao', 'Autorização de Residência');
INSERT INTO identification_types (code, type_name) VALUES ('bi-militar', 'Bilhete de Identidade (militar)');
INSERT INTO identification_types (code, type_name) VALUES ('certificado', 'Certificado de Registo de Cidadão UE');
INSERT INTO identification_types (code, type_name) VALUES ('bi-estrangeiro', 'Bilhete de Identidade (estrangeiro)');
INSERT INTO identification_types (code, type_name) VALUES ('outro', 'Outro');

INSERT INTO document_types (type_name, code) VALUES ('Diploma', 'diploma');  
INSERT INTO document_types (type_name, code) VALUES ('Nº de Identificação', 'nif');  
INSERT INTO document_types (type_name, code) VALUES ('Currículo Vitae', 'curriculo');  
INSERT INTO document_types (type_name, code) VALUES ('Bilhete de Identidade', 'bilhete');  
INSERT INTO document_types (type_name, code) VALUES ('Registo Militar', 'registo-militar');  
INSERT INTO document_types (type_name, code) VALUES ('Documento Bancário', 'doc-bancario'); 
INSERT INTO document_types (type_name, code) VALUES ('Registo Criminal', 'registo-criminal');  
INSERT INTO document_types (type_name, code) VALUES ('Recenseamento Militar', 'recenseamento');  
INSERT INTO document_types (type_name, code) VALUES ('Certificado de Habilitações', 'certificado');  

INSERT INTO contact_types (type_name, code) VALUES ('Casa', 'casa');  
INSERT INTO contact_types (type_name, code) VALUES ('pessoal', 'pessoal');  
INSERT INTO contact_types (type_name, code) VALUES ('Empresa', 'empresa');  
INSERT INTO contact_types (type_name, code) VALUES ('Familiar', 'familiar');  
INSERT INTO contact_types (type_name, code) VALUES ('Outro', 'outro');  

INSERT INTO continents (continent_id, continent_name) VALUES (1, 'Europa');
INSERT INTO continents (continent_id, continent_name) VALUES (2, 'África');
INSERT INTO continents (continent_id, continent_name) VALUES (3, 'Ásia');
INSERT INTO continents (continent_id, continent_name) VALUES (4, 'América');
INSERT INTO continents (continent_id, continent_name) VALUES (5, 'Oceania');

INSERT INTO countries (country_id, country_name, continent_id) VALUES (100, 'Guiné', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (101, 'Guiné Equatorial', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (102, 'Guiné-Bissau', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (103, 'Haiti', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (104, 'Honduras', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (105, 'Hong Kong', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (106, 'Hungria', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (107, 'Iémen', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (108, 'Ilhas Bouvet', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (109, 'Ilhas Caimão', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (10, 'Bélgica', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (110, 'Ilhas Christmas', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (111, 'Ilhas Cocos (Keeling)', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (112, 'Ilhas Cook', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (113, 'Ilhas Falkland (Malvinas)', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (114, 'Ilhas Faroé', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (115, 'Ilhas Fiji', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (116, 'Ilhas Heard e Ilhas McDonald', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (117, 'Ilhas Marianas do Norte', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (118, 'Ilhas Marshall', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (119, 'Ilhas menores distantes dos Estados Unidos', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (11, 'África do Sul', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (120, 'Ilhas Norfolk', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (121, 'Ilhas Salomão', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (122, 'Ilhas Virgens (britânicas)', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (123, 'Ilhas Virgens (Estados Unidos)', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (124, 'Índia', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (125, 'Indonésia', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (126, 'Irão (República Islâmica)', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (127, 'Iraque', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (128, 'Islândia', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (129, 'Israel', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (12, 'Espanha', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (130, 'Jamaica', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (131, 'Japão', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (132, 'Jibuti', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (133, 'Jordânia', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (134, 'Jugoslávia', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (135, 'Kenya', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (136, 'Kiribati', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (137, 'Kuwait', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (138, 'Laos (República Popular Democrática do)', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (139, 'Lesoto', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (13, 'Venezuela', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (140, 'Letónia', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (141, 'Líbano', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (142, 'Libéria', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (143, 'Líbia (Jamahiriya Árabe da)', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (144, 'Liechtenstein', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (145, 'Lituânia', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (146, 'Luxemburgo', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (147, 'Macau', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (148, 'Macedónia (antiga república jugoslava da)', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (149, 'Madagáscar', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (150, 'Malásia', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (151, 'Malawi', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (152, 'Maldivas', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (153, 'Mali', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (154, 'Malta', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (155, 'Martinica', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (156, 'Maurícias', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (157, 'Mauritânia', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (158, 'Mayotte', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (159, 'México', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (15, 'Grã-Bretanha', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (160, 'Micronésia (Estados Federados da)', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (161, 'Moldova (República de)', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (162, 'Mónaco', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (163, 'Mongólia', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (164, 'Monserrate', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (165, 'Myanmar', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (166, 'Namíbia', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (167, 'Nauru', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (168, 'Nepal', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (169, 'Nicarágua', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (16, 'Irlanda', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (170, 'Niger', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (171, 'Nigéria', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (172, 'Niue', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (173, 'Noruega', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (174, 'Nova Caledónia', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (175, 'Nova Zelândia', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (176, 'Omã', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (177, 'Países Baixos', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (178, 'Palau', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (179, 'Panamá', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (17, 'Moçambique', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (180, 'Papuásia-Nova Guiné', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (181, 'Paquistão', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (182, 'Paraguai', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (183, 'Peru', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (184, 'Pitcairn', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (185, 'Polinésia Francesa', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (186, 'Polónia', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (187, 'Porto Rico', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (188, 'Portugal', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (189, 'Quirguizistão', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (18, 'Áustria', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (190, 'Reino Unido', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (191, 'República Checa', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (192, 'República Dominicana', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (193, 'Reunião', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (194, 'Roménia', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (195, 'Ruanda', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (196, 'Rússia (Federação da)', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (197, 'Samoa', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (198, 'Samoa Americana', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (199, 'Santa Helena', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (19, 'Costa Rica', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (200, 'Santa Lúcia', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (201, 'Santa Sé (Cidade Estado do Vaticano)*', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (202, 'São Cristóvão e Nevis', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (203, 'São Marino', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (204, 'São Pedro e Miquelon', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (205, 'São Tomé e Príncipe', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (206, 'São Vicente e Granadinas', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (207, 'Sara Ocidental', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (208, 'Senegal', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (209, 'Serra Leoa', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (210, 'Seychelles', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (211, 'Singapura', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (212, 'Síria (República Árabe da)', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (213, 'Somália', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (214, 'Sri Lanka', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (215, 'Suazilândia', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (216, 'Sudão', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (217, 'Suécia', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (218, 'Suiça', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (219, 'Suriname', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (21, 'Marrocos', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (220, 'Svålbard e a Ilha de Jan Mayen', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (221, 'Tailândia', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (222, 'Taiwan (Província da China)', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (223, 'Tajiquistão', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (224, 'Tanzânia, República Unida da', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (225, 'Território Britânico do Oceano Índico', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (226, 'Território Palestiniano Ocupado', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (227, 'Territórios Franceses do Sul', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (228, 'Timor Leste', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (229, 'Togo', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (22, 'Afeganistão', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (230, 'Tokelau', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (231, 'Tonga', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (232, 'Trindade e Tobago', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (233, 'Tunísia', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (234, 'Turcos e Caicos (Ilhas)', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (235, 'Turquemenistão', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (236, 'Turquia', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (237, 'Tuvalu', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (238, 'Ucrânia', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (239, 'Uganda', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (23, 'Albania', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (240, 'Uruguai', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (241, 'Usbequistão', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (242, 'Vanuatu', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (243, 'Vietname', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (244, 'Wallis e Futuna (Ilha)', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (245, 'Zaire, ver Congo (República Democrática do)', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (246, 'Zâmbia', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (247, 'Zimbabwe', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (24, 'Andorra', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (25, 'Angola', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (26, 'Anguila', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (27, 'Antárctica', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (28, 'Antígua e Barbuda', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (29, 'Antilhas holandesas', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (2, 'Argélia', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (30, 'Arábia Saudita', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (31, 'Argentina', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (32, 'Arménia', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (33, 'Aruba', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (34, 'Austrália', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (35, 'Azerbaijão', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (36, 'Bahamas', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (37, 'Bangladesh', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (38, 'Barbados', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (39, 'Barém', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (3, 'Brasil', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (40, 'Belize', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (41, 'Benin', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (42, 'Bermuda', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (43, 'Bielorrússia', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (44, 'Bolívia', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (45, 'Bósnia e Herzegovina', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (46, 'Botswana', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (47, 'Brunei Darussalam', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (48, 'Bulgária', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (49, 'Burkina Faso', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (4, 'Alemanha', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (50, 'Burundi', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (51, 'Butão', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (52, 'Cabo Verde', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (53, 'Camarões', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (54, 'Camboja', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (55, 'Catar', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (56, 'Cazaquistão', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (57, 'Centro-Africana (República)', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (58, 'Chade', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (59, 'Chile', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (60, 'China', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (61, 'Chipre', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (62, 'Cidade do Vaticano ver Santa Sé', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (63, 'Colômbia', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (64, 'Comores', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (65, 'Congo', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (66, 'Congo (República Democrática do)', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (67, 'Coreia (República da) ', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (68, 'Coreia (República Popular Democrática da) ', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (69, 'Costa do Marfim', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (6, 'França', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (70, 'Croácia', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (71, 'Cuba', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (72, 'Dinamarca', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (73, 'Domínica', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (74, 'Egipto', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (75, 'El Salvador', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (76, 'Emiratos Árabes Unidos', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (77, 'Equador', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (78, 'Eritreia', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (79, 'Eslovaca (República)', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (7, 'Canadá', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (80, 'Eslovénia', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (81, 'Estados Unidos', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (82, 'Estónia', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (83, 'Etiópia', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (84, 'Filipinas', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (85, 'Finlândia', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (86, 'Gabão', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (87, 'Gâmbia', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (88, 'Gana', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (89, 'Geórgia', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (8, 'Itália', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (90, 'Georgia do Sul e Ilhas Sandwich', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (91, 'Gibraltar', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (92, 'Granada', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (93, 'Grécia', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (94, 'Gronelândia', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (95, 'Guadalupe', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (96, 'Guam', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (97, 'Guatemala', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (98, 'Guiana', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (99, 'Guiana Francesa', 1);
INSERT INTO countries (country_id, country_name, continent_id) VALUES (9, 'Holanda', 1);


`
}