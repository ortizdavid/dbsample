--- DATABASE: db_user_roles
--- RDBMS MySQL---

CREATE DATABASE IF NOT EXISTS db_recruitment
USE db_recruitment;


-- CREATE TABLES
DROP TABLE IF EXISTS roles;
CREATE TABLE roles (
    role_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    role_name VARCHAR(50) UNIQUE NOT NULL,
    description VARCHAR(100)
);

DROP TABLE IF EXISTS users;
CREATE TABLE users (
    user_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    role_id INT NOT NULL,
    user_name VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(200) NOT NULL,
    active ENUM('Yes', 'No') NOT NULL DEFAULT 'Yes',
    image VARCHAR(100),
    unique_id BINARY(32) DEFAULT (UUID()),
    created_at DATETIME DEFAULT NOW(),
    updated_at DATETIME DEFAULT NOW(),
    CONSTRAINT fk_role_users FOREIGN KEY(role_id) REFERENCES roles(role_id)
);

DROP TABLE IF EXISTS continents;
CREATE TABLE IF NOT EXISTS continents (
	continent_id INT NOT NULL PRIMARY KEY,
	continent_name VARCHAR(100) NOT NULL UNIQUE
);

DROP TABLE IF EXISTS countries;
CREATE TABLE IF NOT EXISTS countries (
	country_id INT NOT NULL PRIMARY KEY,
	continent_id INT NOT NULL,
	country_name VARCHAR(150) NOT NULL UNIQUE,
    CONSTRAINT fk_continent FOREIGN KEY(country_id) REFERENCES continents(continent_id)
);

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

DROP TABLE IF EXISTS study_areas;
CREATE TABLE IF NOT EXISTS study_areas (
	area_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	area_name VARCHAR(100) NOT NULL UNIQUE,
	description VARCHAR(200) UNIQUE
);

DROP TABLE IF EXISTS degrees;
CREATE TABLE IF NOT EXISTS degrees (
	degree_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	degree_name VARCHAR(100) NOT NULL UNIQUE,
	description VARCHAR(200) UNIQUE
);

DROP TABLE IF EXISTS application_statuses;
CREATE TABLE IF NOT EXISTS application_statuses (
	status_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	status_name VARCHAR(100) NOT NULL UNIQUE,
	code VARCHAR(20) UNIQUE
);

DROP TABLE IF EXISTS candidates;
CREATE TABLE candidates (
    candidate_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    country_id INT,
    identification_type_id INT,
    marital_status_id INT,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    identification_number VARCHAR(30),
    gender ENUM('Masculino', 'Feminino'),
    unique_id BINARY(32) DEFAULT (UUID()),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user_candidate FOREIGN KEY(user_id) REFERENCES users(user_id),
    CONSTRAINT fk_country_candidate FOREIGN KEY(country_id) REFERENCES countries(country_id),
    CONSTRAINT fk_identification_candidate FOREIGN KEY(identification_type_id) REFERENCES identification_types(type_id),
    CONSTRAINT fk_marital_candidate FOREIGN KEY(marital_status_id) REFERENCES marital_statuses(status_id)
);

DROP TABLE IF EXISTS documents;
CREATE TABLE IF NOT EXISTS documents (
	document_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    candidate_id INT NOT NULL,
    document_type_id INT NOT NULL,
    document_name VARCHAR(30),
	file_name VARCHAR(150),
    unique_id BINARY(32) DEFAULT (UUID()),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT fk_candidate_document FOREIGN KEY(candidate_id) REFERENCES candidate(candidate_id),
    CONSTRAINT fk_document_type FOREIGN KEY(document_type_id) REFERENCES document_types(type_id)
);

DROP TABLE IF EXISTS contacts;
CREATE TABLE contacts (
    contact_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    candidate_id INT NOT NULL,
    phone INT NOT NULL,
    email VARCHAR(100) NOT NULL,
    CONSTRAINT fk_candidate FOREIGN KEY(candidate_id) REFERENCES candidates(candidate_id)
);

DROP TABLE IF EXISTS addresses;
CREATE TABLE addresses (
    address_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    candidate_id INT NOT NULL,
    state VARCHAR(50),
    city VARCHAR(50),
    district VARCHAR(50),
    CONSTRAINT fk_candidate_addres FOREIGN KEY(candidate_id) REFERENCES candidate(candidate_id)
);

DROP TABLE IF EXISTS graduations;
CREATE TABLE IF NOT EXISTS graduations (
	graduation_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    candidate_id INT NOT NULL,
    study_area_id INT NOT NULL,
    institute VARCHAR(150),
	degree_id INT,
    start_year INT(4),
    end_year INT(4),
    unique_id BINARY(32) DEFAULT (UUID()),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT fk_candidate_graduation FOREIGN KEY(candidate_id) REFERENCES candidate(candidate_id),
    CONSTRAINT fk_area_graduation FOREIGN KEY(study_area_id) REFERENCES study_areas(area_id),
    CONSTRAINT fk_degree_graduation FOREIGN KEY(degree_id) REFERENCES degrees(degree_id)
);

DROP TABLE IF EXISTS vacancies;
CREATE TABLE IF NOT EXISTS vacancies (
	vacancy_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    study_area_id INT NOT NULL,
	title VARCHAR(100),
	num_positions INT,
	start_date DATE,
	end_date DATE,
	status ENUM("Aberta", "Fechada") DEFAULT 'Aberta',
	description VARCHAR(300) DEFAULT 'Sem Descrição',
	code VARCHAR(20) NOT NULL UNIQUE,
	unique_id BINARY(32) DEFAULT (UUID()),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT fk_area_vacancy FOREIGN KEY(study_area_id) REFERENCES study_areas(area_id_id)
);

DROP TABLE IF EXISTS job_applications;
CREATE TABLE IF NOT EXISTS job_applications (
	job_application_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	candidate_id INT NOT NULL,
	vacancy_id INT NOT NULL,
	application_status_id INT NOT NULL,
	application_date DATE,
	code VARCHAR(20) NOT NULL UNIQUE,
    reason VARCHAR(200),
	unique_id BINARY(32) DEFAULT (UUID()),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT fk_candidate_application FOREIGN KEY(candidate_id) REFERENCES candidates(candidate_id),
	CONSTRAINT fk_status_application FOREIGN KEY(application_status_id) REFERENCES application_statuses(status_id),
	CONSTRAINT fk_vacancy_application FOREIGN KEY(vacancy_id) REFERENCES vacancies(vacancy_id)
);

DROP TABLE IF EXISTS registrations;
CREATE TABLE registrations (
    registration_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(150),
    gender ENUM('Masculino', 'Feminino'),
    unique_id BINARY(32) DEFAULT (UUID()),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user_registration FOREIGN KEY(user_id) REFERENCES users(user_id),
);
   


-- VIEWS 
-- View view_user_data
DROP VIEW IF exists view_user_data;
CREATE VIEW view_user_data AS 
SELECT us.user_id, us.unique_id,
    us.user_name, us.password, 
    us.active, us.image,
    us.created_at, us.updated_at,
    ro.role_id, ro.role_name,
    ro.description
FROM users us
JOIN roles ro ON(ro.role_id = us.role_id);

-- View view_candidate_data
DROP VIEW IF EXISTS view_candidate_data;
CREATE VIEW view_candidate_data AS
SELECT ca.candidate_id, ca.unique_id,
    ca.first_name, ca.last_name, 
    ca.identification_number, ca.gender,
    ca.created_at, ca.updated_at,
    us.user_id, us.user_name, us.image,
    us.unique_id AS user_unique_id,
    ms.status_id, ms.status_name,
    it.type_id, it.type_name,
    co.country_id, co.country_name
FROM candidates ca
JOIN users us ON(us.user_id = ca.user_id)
JOIN identification_types it ON(it.type_id = ca.identification_type_id)
JOIN marital_statuses ms ON(ms.status_id = ca.marital_status_id)
JOIN countries co ON(co.country_id = ca.country_id);

-- View view_docuent_data
DROP VIEW IF EXISTS view_document_data;
CREATE VIEW view_document_data AS 
SELECT doc.document_id, 
    doc.document_name, doc.file_name,
    dt.type_id, dt.type_name,
    ca.candidate_id
FROM documents doc
JOIN document_types dt ON(dt.type_id = doc.document_type_id)
JOIN candidates ca ON(ca.candidate_id = doc.candidate_id);

-- View view_graduation_data
DROP VIEW IF EXISTS view_graduation_data;
CREATE VIEW view_graduation_data AS 
SELECT gr.graduation_id, gr.institute,
    gr.start_year, gr.end_year,
    dg.degree_id, dg.degree_name,
    sa.area_id AS study_area_id, 
    sa.area_name AS study_area_name,
    ca.candidate_id
FROM graduations gr
JOIN degrees dg ON(dg.degree_id = gr.degree_id)
JOIN study_areas sa ON(sa.area_id = gr.study_area_id)
JOIN candidates ca ON(ca.candidate_id = gr.candidate_id);

-- View view_vacancy_data
DROP VIEW IF EXISTS view_vacancy_data;
CREATE VIEW view_vacancy_data AS 
SELECT va.vacancy_id, va.unique_id,
    va.title, va.code, 
    va.num_positions,
    va.start_date, va.end_date, 
    va.status, va.description,
    va.created_at, va.updated_at,
    sa.area_id AS study_area_id, 
    sa.area_name AS study_area_name
FROM vacancies va
JOIN study_areas sa ON(sa.area_id = va.study_area_id);

-- VIEW view_job_application_data
DROP VIEW IF EXISTS view_job_application_data;
CREATE VIEW view_job_application_data AS 
SELECT ja.job_application_id, ja.unique_id,
    ja.application_date, ja.code, ja.reason,
	ja.created_at, ja.updated_at, 
	st.status_id, st.status_name AS application_status_name,
	st.code AS application_status_code,
	ca.candidate_id, ca.unique_id AS candidate_unique_id,
	ca.first_name, ca.last_name, 
    us.user_id, us.unique_id AS user_unique_id,
    va.vacancy_id, 
    va.code AS vacancy_code, 
    va.title AS vacancy_title
FROM job_applications ja
JOIN application_statuses st ON(st.status_id = ja.application_status_id)
JOIN candidates ca ON(ca.candidate_id = ja.candidate_id)
JOIN users us ON(us.user_id = ca.user_id)
JOIN vacancies va ON(va.vacancy_id = ja.vacancy_id);



-- INSERTS
INSERT INTO roles (role_id, role_name, description) VALUES
(1, 'administrator', 'Administrador'),
(2, 'secretary', 'Secretaria'),
(3, 'candidate', 'Candidato');

INSERT INTO identification_types (type_id, type_name, code) VALUES
(1, 'Bilhete de Identidade', 'bi'),
(2, 'Passaporte', 'passaporte'),
(3, 'Cartão de Residente', 'residente'),
(4, 'Bilhete de Identidade (Cabo Verde)', 'bi-cverde'),
(5, 'Autorização de Residência', 'autorizacao'),
(6, 'Bilhete de Identidade (militar)', 'bi-militar'),
(7, 'Certificado de Registo de Cidadão UE', 'certificado'),
(8, 'Bilhete de Identidade (estrangeiro)', 'bi-estrangeiro'),
(9, 'Outro', 'outro');

INSERT INTO document_types (type_id, type_name, code) VALUES
(1, 'Diploma', 'diploma'),
(2, 'Passaporte', 'passaporte'),
(3, 'Currículo Vitae', 'curriculo'),
(4, 'Documento de Identificação', 'identificacao'),
(5, 'Certificado de Habilitações', 'certificado'),
(6, 'Cartão de Membro', 'cartao-embro');

INSERT INTO marital_statuses (status_id, status_name, code) VALUES
(1, 'Solteiro', 'solteiro'),
(2, 'Casado (com registo)', 'casado_com_registo'),
(3, 'Casado (sem registo)', 'casado_sem_registo'),
(4, 'Divorciado', 'divorciado'),
(5, 'Separado', 'separado'),
(6, 'Viúvo', 'viuvo'),
(7, 'Outro', 'outro');

INSERT INTO application_statuses (status_id, status_name, code) VALUES
(1, 'Pendente', 'pendente'),
(2, 'Validada', 'validada'),
(3, 'Negada', 'negada'),
(4, 'Cancelada', 'cancelada');

INSERT INTO degrees (degree_id, degree_name, description) VALUES
(1, 'Ensino Médio', NULL),
(2, 'Curso Profissional', NULL),
(3, 'Bacharelato', NULL),
(4, 'Pós Graduação', NULL),
(5, 'Licenciatura', NULL),
(6, 'Mestrado', NULL),
(7, 'Doutoramento', NULL);

INSERT INTO study_areas (area_id, area_name, description) VALUES
(1, 'Medicina Geral', NULL),
(2, 'Fisioterapia', NULL),
(3, 'Ortopedia', NULL),
(4, 'Estomatologia', NULL),
(5, 'Laboratório', NULL),
(6, 'Clínica Geral', NULL),
(7, 'Enfermagem', NULL),
(8, 'Análises Clínicas', NULL),
(9, 'Nutrição', NULL),
(10, 'Optometria', NULL),
(11, 'Oftalmologia', NULL),
(12, 'Cardiologia', NULL),
(13, 'Radiologia', NULL),
(14, 'Pediatria', NULL),
(15, 'Obstetrícia', NULL),
(16, 'Imagiologia', NULL),
(17, 'Engenharia Informática', NULL),
(18, 'Ciência da Computação', NULL),
(19, 'Química', NULL),
(20, 'Gestão de Recursos Humanos', NULL),
(21, 'Gestão e Administração', NULL);


INSERT INTO continents (continent_id, continent_name) VALUES
(1, 'Europa'),
(2, 'África'),
(3, 'Ásia'),
(4, 'América'),
(5, 'Oceania');

INSERT INTO countries (country_id, continent_id, country_name) VALUES
(100, 1, 'Guiné'),
(101, 1, 'Guiné Equatorial'),
(102, 1, 'Guiné-Bissau'),
(103, 1, 'Haiti'),
(104, 1, 'Honduras'),
(105, 1, 'Hong Kong'),
(106, 1, 'Hungria'),
(107, 1, 'Iémen'),
(108, 1, 'Ilhas Bouvet'),
(109, 1, 'Ilhas Caimão'),
(10, 1, 'Bélgica'),
(110, 1, 'Ilhas Christmas'),
(111, 1, 'Ilhas Cocos (Keeling)'),
(112, 1, 'Ilhas Cook'),
(113, 1, 'Ilhas Falkland (Malvinas)'),
(114, 1, 'Ilhas Faroé'),
(115, 1, 'Ilhas Fiji'),
(116, 1, 'Ilhas Heard e Ilhas McDonald'),
(117, 1, 'Ilhas Marianas do Norte'),
(118, 1, 'Ilhas Marshall'),
(119, 1, 'Ilhas menores distantes dos Estados Unidos'),
(11, 1, 'África do Sul'),
(120, 1, 'Ilhas Norfolk'),
(121, 1, 'Ilhas Salomão'),
(122, 1, 'Ilhas Virgens (britânicas)'),
(123, 1, 'Ilhas Virgens (Estados Unidos)'),
(124, 1, 'Índia'),
(125, 1, 'Indonésia'),
(126, 1, 'Irão (República Islâmica)'),
(127, 1, 'Iraque'),
(128, 1, 'Islândia'),
(129, 1, 'Israel'),
(12, 1, 'Espanha'),
(130, 1, 'Jamaica'),
(131, 1, 'Japão'),
(132, 1, 'Jibuti'),
(133, 1, 'Jordânia'),
(134, 1, 'Jugoslávia'),
(135, 1, 'Kenya'),
(136, 1, 'Kiribati'),
(137, 1, 'Kuwait'),
(138, 1, 'Laos (República Popular Democrática do)'),
(139, 1, 'Lesoto'),
(13, 1, 'Venezuela'),
(140, 1, 'Letónia'),
(141, 1, 'Líbano'),
(142, 1, 'Libéria'),
(143, 1, 'Líbia (Jamahiriya Árabe da)'),
(144, 1, 'Liechtenstein'),
(145, 1, 'Lituânia'),
(146, 1, 'Luxemburgo'),
(147, 1, 'Macau'),
(148, 1, 'Macedónia (antiga república jugoslava da)'),
(149, 1, 'Madagáscar'),
(150, 1, 'Malásia'),
(151, 1, 'Malawi'),
(152, 1, 'Maldivas'),
(153, 1, 'Mali'),
(154, 1, 'Malta'),
(155, 1, 'Martinica'),
(156, 1, 'Maurícias'),
(157, 1, 'Mauritânia'),
(158, 1, 'Mayotte'),
(159, 1, 'México'),
(15, 1, 'Grã-Bretanha'),
(160, 1, 'Micronésia (Estados Federados da)'),
(161, 1, 'Moldova (República de)'),
(162, 1, 'Mónaco'),
(163, 1, 'Mongólia'),
(164, 1, 'Monserrate'),
(165, 1, 'Myanmar'),
(166, 1, 'Namíbia'),
(167, 1, 'Nauru'),
(168, 1, 'Nepal'),
(169, 1, 'Nicarágua'),
(16, 1, 'Irlanda'),
(170, 1, 'Niger'),
(171, 1, 'Nigéria'),
(172, 1, 'Niue'),
(173, 1, 'Noruega'),
(174, 1, 'Nova Caledónia'),
(175, 1, 'Nova Zelândia'),
(176, 1, 'Omã'),
(177, 1, 'Países Baixos'),
(178, 1, 'Palau'),
(179, 1, 'Panamá'),
(17, 1, 'Moçambique'),
(180, 1, 'Papuásia-Nova Guiné'),
(181, 1, 'Paquistão'),
(182, 1, 'Paraguai'),
(183, 1, 'Peru'),
(184, 1, 'Pitcairn'),
(185, 1, 'Polinésia Francesa'),
(186, 1, 'Polónia'),
(187, 1, 'Porto Rico'),
(188, 1, 'Portugal'),
(189, 1, 'Quirguizistão'),
(18, 1, 'Áustria'),
(190, 1, 'Reino Unido'),
(191, 1, 'República Checa'),
(192, 1, 'República Dominicana'),
(193, 1, 'Reunião'),
(194, 1, 'Roménia'),
(195, 1, 'Ruanda'),
(196, 1, 'Rússia (Federação da)'),
(197, 1, 'Samoa'),
(198, 1, 'Samoa Americana'),
(199, 1, 'Santa Helena'),
(19, 1, 'Costa Rica'),
(200, 1, 'Santa Lúcia'),
(201, 1, 'Santa Sé (Cidade Estado do Vaticano)*'),
(202, 1, 'São Cristóvão e Nevis'),
(203, 1, 'São Marino'),
(204, 1, 'São Pedro e Miquelon'),
(205, 1, 'São Tomé e Príncipe'),
(206, 1, 'São Vicente e Granadinas'),
(207, 1, 'Sara Ocidental'),
(208, 1, 'Senegal'),
(209, 1, 'Serra Leoa'),
(210, 1, 'Seychelles'),
(211, 1, 'Singapura'),
(212, 1, 'Síria (República Árabe da)'),
(213, 1, 'Somália'),
(214, 1, 'Sri Lanka'),
(215, 1, 'Suazilândia'),
(216, 1, 'Sudão'),
(217, 1, 'Suécia'),
(218, 1, 'Suiça'),
(219, 1, 'Suriname'),
(21, 1, 'Marrocos'),
(220, 1, 'Svålbard e a Ilha de Jan Mayen'),
(221, 1, 'Tailândia'),
(222, 1, 'Taiwan (Província da China)'),
(223, 1, 'Tajiquistão'),
(224, 1, 'Tanzânia, República Unida da'),
(225, 1, 'Território Britânico do Oceano Índico'),
(226, 1, 'Território Palestiniano Ocupado'),
(227, 1, 'Territórios Franceses do Sul'),
(228, 1, 'Timor Leste'),
(229, 1, 'Togo'),
(22, 1, 'Afeganistão'),
(230, 1, 'Tokelau'),
(231, 1, 'Tonga'),
(232, 1, 'Trindade e Tobago'),
(233, 1, 'Tunísia'),
(234, 1, 'Turcos e Caicos (Ilhas)'),
(235, 1, 'Turquemenistão'),
(236, 1, 'Turquia'),
(237, 1, 'Tuvalu'),
(238, 1, 'Ucrânia'),
(239, 1, 'Uganda'),
(23, 1, 'Albania'),
(240, 1, 'Uruguai'),
(241, 1, 'Usbequistão'),
(242, 1, 'Vanuatu'),
(243, 1, 'Vietname'),
(244, 1, 'Wallis e Futuna (Ilha)'),
(245, 1, 'Zaire, ver Congo (República Democrática do)'),
(246, 1, 'Zâmbia'),
(247, 1, 'Zimbabwe'),
(24, 1, 'Andorra'),
(25, 1, 'Angola'),
(26, 1, 'Anguila'),
(27, 1, 'Antárctica'),
(28, 1, 'Antígua e Barbuda'),
(29, 1, 'Antilhas holandesas'),
(2, 1, 'Argélia'),
(30, 1, 'Arábia Saudita'),
(31, 1, 'Argentina'),
(32, 1, 'Arménia'),
(33, 1, 'Aruba'),
(34, 1, 'Austrália'),
(35, 1, 'Azerbaijão'),
(36, 1, 'Bahamas'),
(37, 1, 'Bangladesh'),
(38, 1, 'Barbados'),
(39, 1, 'Barém'),
(3, 1, 'Brasil'),
(40, 1, 'Belize'),
(41, 1, 'Benin'),
(42, 1, 'Bermuda'),
(43, 1, 'Bielorrússia'),
(44, 1, 'Bolívia'),
(45, 1, 'Bósnia e Herzegovina'),
(46, 1, 'Botswana'),
(47, 1, 'Brunei Darussalam'),
(48, 1, 'Bulgária'),
(49, 1, 'Burkina Faso'),
(4, 1, 'Alemanha'),
(50, 1, 'Burundi'),
(51, 1, 'Butão'),
(52, 1, 'Cabo Verde'),
(53, 1, 'Camarões'),
(54, 1, 'Camboja'),
(55, 1, 'Catar'),
(56, 1, 'Cazaquistão'),
(57, 1, 'Centro-Africana (República)'),
(58, 1, 'Chade'),
(59, 1, 'Chile'),
(60, 1, 'China'),
(61, 1, 'Chipre'),
(62, 1, 'Cidade do Vaticano ver Santa Sé'),
(63, 1, 'Colômbia'),
(64, 1, 'Comores'),
(65, 1, 'Congo'),
(66, 1, 'Congo (República Democrática do)'),
(67, 1, 'Coreia (República da) '),
(68, 1, 'Coreia (República Popular Democrática da) '),
(69, 1, 'Costa do Marfim'),
(6, 1, 'França'),
(70, 1, 'Croácia'),
(71, 1, 'Cuba'),
(72, 1, 'Dinamarca'),
(73, 1, 'Domínica'),
(74, 1, 'Egipto'),
(75, 1, 'El Salvador'),
(76, 1, 'Emiratos Árabes Unidos'),
(77, 1, 'Equador'),
(78, 1, 'Eritreia'),
(79, 1, 'Eslovaca (República)'),
(7, 1, 'Canadá'),
(80, 1, 'Eslovénia'),
(81, 1, 'Estados Unidos'),
(82, 1, 'Estónia'),
(83, 1, 'Etiópia'),
(84, 1, 'Filipinas'),
(85, 1, 'Finlândia'),
(86, 1, 'Gabão'),
(87, 1, 'Gâmbia'),
(88, 1, 'Gana'),
(89, 1, 'Geórgia'),
(8, 1, 'Itália'),
(90, 1, 'Georgia do Sul e Ilhas Sandwich'),
(91, 1, 'Gibraltar'),
(92, 1, 'Granada'),
(93, 1, 'Grécia'),
(94, 1, 'Gronelândia'),
(95, 1, 'Guadalupe'),
(96, 1, 'Guam'),
(97, 1, 'Guatemala'),
(98, 1, 'Guiana'),
(99, 1, 'Guiana Francesa'),
(9, 1, 'Holanda');


