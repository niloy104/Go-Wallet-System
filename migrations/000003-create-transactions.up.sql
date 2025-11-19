-- +migrate Up
CREATE TABLE transactions (
    id            BIGSERIAL PRIMARY KEY,
    from_user_id  BIGINT,
    to_user_id    BIGINT,
    amount        BIGINT NOT NULL CHECK (amount > 0),
    type          VARCHAR(20) NOT NULL, -- 'topup', 'transfer_out', 'transfer_in'
    status        VARCHAR(20) NOT NULL DEFAULT 'success',
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_transactions_from_user ON transactions(from_user_id);
CREATE INDEX idx_transactions_to_user   ON transactions(to_user_id);
CREATE INDEX idx_transactions_created_at ON transactions(created_at DESC);