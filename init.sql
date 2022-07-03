CREATE DATABASE coinstream;

CREATE TABLE IF NOT EXISTS expense(
    ID          SERIAL PRIMARY KEY NOT NULL,
    NAME        TEXT            NOT NULL,
    AMOUNT      FLOAT8          NOT NULL,
    CATEGORY    VARCHAR(32)     NOT NULL,
    DATE        DATE            NOT NULL
);

CREATE TABLE IF NOT EXISTS income(
    ID          SERIAL PRIMARY KEY NOT NULL,
    NAME        TEXT            NOT NULL,
    AMOUNT      FLOAT8          NOT NULL,
    DATE        DATE            NOT NULL
);