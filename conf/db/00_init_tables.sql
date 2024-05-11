CREATE TABLE countries (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

INSERT INTO countries (id, name) VALUES 
    (1, 'Thailand'), 
    (2, 'France'),
    (999, 'Not Specified');

CREATE TABLE provinces (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

INSERT INTO provinces (id, name) VALUES 
    (1, 'Province 1'), 
    (2, 'Province 2'),
    (999, 'Not Specified');


CREATE TABLE districts (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

INSERT INTO districts (id, name) VALUES 
    (1, 'District 1'), 
    (2, 'District 2'),
    (999, 'Not Specified');

CREATE TABLE subdistricts (
    id SERIAL PRIMARY KEY,
    district_id INTEGER REFERENCES districts(id) NOT NULL,
    name VARCHAR(255) NOT NULL
);

INSERT INTO subdistricts (id, district_id, name) VALUES 
    (1, 1, 'Subdistrict 1-1'), 
    (2, 1, 'Subdistrict 1-2'), 
    (3, 1, 'Subdistrict 1-3'), 
    (4, 2, 'Subdistrict 2-1'),
    (5, 2, 'Subdistrict 2-2'),
    (999, 999, 'Not Specified');

CREATE TABLE members (
    id SERIAL PRIMARY KEY,
    username VARCHAR(64) UNIQUE NOT NULL,
    passwd VARCHAR(64) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    phone VARCHAR(255) UNIQUE NOT NULL,
    address1 VARCHAR(255),
    address2 VARCHAR(255),
    postal_code VARCHAR(10),

    country_id INTEGER REFERENCES countries(id) DEFAULT 999,
    province_id INTEGER REFERENCES provinces(id) DEFAULT 999,
    district_id INTEGER REFERENCES districts(id) DEFAULT 999,
    subdistrict_id INTEGER REFERENCES subdistricts(id) DEFAULT 999,
    organization VARCHAR(255),
    registration_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    confirmation_time TIMESTAMP,
    update_time TIMESTAMP,
    notes TEXT
);
