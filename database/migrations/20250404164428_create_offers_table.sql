-- +goose Up
-- +goose StatementBegin
CREATE TABLE offers
(
    id           SERIAL PRIMARY KEY,
    hash         TEXT        NOT NULL,
    name         TEXT        NOT NULL,
    product_type TEXT        NOT NULL,
    created_at   TIMESTAMPTZ DEFAULT now(),
    updated_at   TIMESTAMPTZ DEFAULT now(),
    deleted_at   TIMESTAMPTZ NULL
);
CREATE INDEX idx_offers_hash on offers (id);
CREATE INDEX idx_deleted_at on offers (deleted_at);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE offers;
-- +goose StatementEnd
