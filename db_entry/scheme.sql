CREATE TABLE clients (
    id serial PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    balance INT DEFAULT 0 CHECK (balance >= 0), -- Ленивый баланс (подсчитан вплоть до операции last_operation)
    last_operation INT DEFAULT -1 -- Указатель на последнюю операцию, учтенную в балансе
);

CREATE TABLE operations (
    id serial PRIMARY KEY,
    client_id INT REFERENCES clients(id),
    change INT -- Любые изменения баланса, включая не валидные
);