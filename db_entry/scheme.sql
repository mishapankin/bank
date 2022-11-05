CREATE TABLE clients (
    id serial PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    balance INT DEFAULT 0 CHECK (balance > 0)
);