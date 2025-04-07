-- +goose Up
-- +goose StatementBegin
CREATE TABLE subscriptions
(
    id         SERIAL PRIMARY KEY,
    hash       TEXT NOT NULL,
    order_hash TEXT NOT NULL,
    offer_id   INT REFERENCES offers (id),
    amount     INT  NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),
    deleted_at TIMESTAMPTZ
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE subscriptions;
-- +goose StatementEnd
