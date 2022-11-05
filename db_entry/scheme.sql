CREATE TABLE clients (
    id serial PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    balance INT DEFAULT 0 CHECK (balance >= 0),
    last_operation INT DEFAULT -1
);

CREATE TABLE operations (
    id serial PRIMARY KEY,
    client_id INT REFERENCES clients(id),
    change INT
);