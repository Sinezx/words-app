DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id serial PRIMARY KEY,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    deleted_at timestamp,
    account varchar(256) NOT NULL UNIQUE,
    password varchar(256) NOT NULL
);

CREATE INDEX users_account_idx ON users(account);

DROP TABLE IF EXISTS words;

CREATE TABLE words (
    id serial PRIMARY KEY,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    deleted_at timestamp,
    user_id integer NOT NULL,
    source_text varchar(256) NOT NULL,
    target_text varchar(256) NOT NULL,
    rate numeric NOT NULL,
    rate_up_at timestamp NOT NULL,
    UNIQUE (user_id, source_text)
);