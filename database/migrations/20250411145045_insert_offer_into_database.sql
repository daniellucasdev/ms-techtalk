-- +goose Up
-- +goose StatementBegin
INSERT INTO offers (hash, name, product_type)
VALUES ('off123', 'IPTV Premium', 'DIGITAL');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
