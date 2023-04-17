CREATE TABLE if not exists users (
    id SERIAL PRIMARY KEY ,
    user_name VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    user_status CHAR(1) NOT NULL,
    department VARCHAR(255)
);
