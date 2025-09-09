CREATE TABLE token_keys (
    id SERIAL PRIMARY KEY,
    private_key TEXT NOT NULL,
    public_key TEXT NOT NULL,
    is_active BOOLEAN DEFAULT true,
);
