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

DROP TABLE IF EXISTS user_words;

CREATE TABLE "user_words" (
  "id" serial PRIMARY KEY,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "deleted_at" timestamp,
  "user_id" integer NOT NULL,
  "rate" numeric NOT NULL,
  "rate_up_at" timestamp NOT NULL,
  "word_id" integer NOT NULL,
  UNIQUE ("user_id" ASC, "word_id" ASC)
);

DROP TABLE IF EXISTS words;

CREATE TABLE "words" (
  "id" serial PRIMARY KEY,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "deleted_at" timestamp,
  "source_text" varchar(256) NOT NULL,
  "target_text" varchar(256) NOT NULL,
  "voice_path" varchar(256),
  UNIQUE ("source_text" ASC)
);