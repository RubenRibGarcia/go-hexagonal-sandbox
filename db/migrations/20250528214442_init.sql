-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS bank_accounts (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    balance DECIMAL(10, 2) NOT NULL DEFAULT 0.00
);

CREATE TABLE IF NOT EXISTS transactions (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    bank_account_id UUID NOT NULL REFERENCES bank_accounts(id),
    amount DECIMAL(10, 2) NOT NULL,
    kind VARCHAR(10) NOT NULL CHECK (kind IN ('credit', 'debit'))
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS bank_accounts;
-- +goose StatementEnd
