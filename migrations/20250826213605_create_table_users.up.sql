CREATE TABLE users
(
    id uuid     NOT NULL,
    email       VARCHAR(255) UNIQUE,
    password_hash VARCHAR(72) NOT NULL,
    first_name  VARCHAR(50) NOT NULL,
    second_name VARCHAR(50) NOT NULL,
    birthdate   TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    biography   TEXT,
    city        VARCHAR(30) NOT NULL,
    created_at  TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    PRIMARY KEY (id)
);