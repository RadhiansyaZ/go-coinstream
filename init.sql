CREATE DATABASE coinstream;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

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

CREATE TABLE IF NOT EXISTS user_coinstream(
    ID              UUID DEFAULT uuid_generate_v4() PRIMARY KEY NOT NULL,
    EMAIL           TEXT                UNIQUE NOT NULL,
    USERNAME        VARCHAR             UNIQUE NOT NULL,
    HASHED_PASSWORD VARCHAR                    NOT NULL,
    NAME            TEXT                       NOT NULL,
    CREATED_AT      TIMESTAMP DEFAULT NOW()    NOT NULL
);