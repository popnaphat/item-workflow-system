-- +goose Up
INSERT INTO items (title, amount, quantity, status, owner_id)
VALUES ('สาย LAN', 1000, 10, 'PENDING', 1);
-- +goose StatementBegin
SELECT
    'up SQL query';

-- +goose StatementEnd
-- +goose Down
TRUNCATE TABLE items;

-- +goose StatementBegin
SELECT
    'down SQL query';

-- +goose StatementEnd