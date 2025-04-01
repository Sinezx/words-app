

DROP TABLE IF EXISTS words;

CREATE TABLE words (
    id serial PRIMARY KEY,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    deleted_at boolean,
    subject varchar(256) NOT NULL UNIQUE,
    translation varchar(256) NOT NULL UNIQUE,
    rate numeric NOT NULL,
    rate_up_at timestamp NOT NULL NOW()
);