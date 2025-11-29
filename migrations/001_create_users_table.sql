CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(64) UNIQUE NOT NULL,
    password_hash BYTEA NOT NULL,
    refresh_token TEXT,
    test BOOLEAN DEFAULT TRUE 
);
