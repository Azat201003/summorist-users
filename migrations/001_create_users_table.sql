CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(64) UNIQUE,
    password_hash BYTEA NOT NULL,
    refresh_token TEXT
);
