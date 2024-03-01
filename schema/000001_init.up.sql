CREATE TABLE IF NOT EXISTS projects (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50),
    created_at TIMESTAMP
);


CREATE TABLE IF NOT EXISTS goods (
    id SERIAl PRIMARY KEY,
    project_id INT NOT NULL UNIQUE,
    name VARCHAR(64),
    description VARCHAR(128),
    priority INT,
    removed BOOLEAN,
    created_at TIMESTAMP
);



