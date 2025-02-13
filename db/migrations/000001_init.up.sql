CREATE TABLE wallets (
    id SERIAL PRIMARY KEY NOT NULL,
    address CHAR(64) UNIQUE NOT NULL,
    balance NUMERIC(15, 2) NOT NULL 
);

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY NOT NULL,
    from_wallet CHAR(64) REFERENCES wallets(address) ON DELETE SET NULL,
    to_wallet CHAR(64) REFERENCES wallets(address) ON DELETE SET NULL,
    amount NUMERIC(15, 2) NOT NULL,
    time TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_transaction_time ON transactions (time);
