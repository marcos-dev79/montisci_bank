CREATE TABLE users (
    id   BIGSERIAL PRIMARY KEY,
    name text      NOT NULL,
    bio  text,
    birthdate date NOT NULL,
    email text NOT NULL,
    last_login TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL, 
    updated_at TIMESTAMP NOT NULL,
    profession_id BIGINT NOT NULL
);

CREATE TABLE PROFESSIONS (
    id   BIGSERIAL PRIMARY KEY,
    name text      NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE ACCOUNTS (
    id   BIGSERIAL PRIMARY KEY,
    name text      NOT NULL,
    balance decimal NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE ACCOUNT_USERS (
    id   BIGSERIAL PRIMARY KEY,
    account_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE ACCOUNT_TRANSACTIONS (
    id   BIGSERIAL PRIMARY KEY,
    account_id BIGINT NOT NULL,
    amount decimal NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);