-- Модель данных CRM «Клиентский путь»

CREATE TABLE managers (
    id          SERIAL PRIMARY KEY,
    full_name   TEXT NOT NULL,
    email       TEXT NOT NULL UNIQUE,
    role        TEXT NOT NULL CHECK (role IN ('manager', 'head', 'admin'))
);

CREATE TABLE clients (
    id          SERIAL PRIMARY KEY,
    full_name   TEXT NOT NULL,
    company     TEXT,
    phone       TEXT,
    email       TEXT,
    source      TEXT NOT NULL,            -- сайт, звонок, e-mail, мессенджер
    manager_id  INT REFERENCES managers(id),
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE deals (
    id          SERIAL PRIMARY KEY,
    client_id   INT NOT NULL REFERENCES clients(id),
    stage       TEXT NOT NULL DEFAULT 'first_contact',
    amount      NUMERIC(12, 2),
    lost_reason TEXT,
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE touches (
    id          SERIAL PRIMARY KEY,
    deal_id     INT NOT NULL REFERENCES deals(id),
    channel     TEXT NOT NULL,
    note        TEXT,
    happened_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE contracts (
    id          SERIAL PRIMARY KEY,
    deal_id     INT NOT NULL UNIQUE REFERENCES deals(id),
    number      TEXT NOT NULL UNIQUE,
    signed_at   DATE
);
