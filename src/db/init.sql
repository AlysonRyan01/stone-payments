CREATE TABLE IF NOT EXISTS transaction (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    card_number TEXT        NOT NULL,
    amount      BIGINT      NOT NULL CHECK (amount > 0),
    currency    TEXT        NOT NULL,
    merchant    TEXT        NOT NULL,
    timestamp   TIMESTAMPTZ NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_transaction_card_timestamp
    ON transaction (card_number, timestamp DESC);