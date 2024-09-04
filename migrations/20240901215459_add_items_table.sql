-- +goose Up
CREATE TABLE
    items (
        id BIGSERIAL PRIMARY KEY, 
        title TEXT NOT NULL,
        amount REAL NOT NULL,
        quantity INTEGER NOT NULL,
        status TEXT NOT NULL,
        owner_id INTEGER NOT NULL
    );

-- +goose StatementBegin
SELECT
    'up SQL query';

-- +goose StatementEnd
-- +goose Down
DROP TABLE items;

-- +goose StatementBegin
SELECT
    'down SQL query';

-- +goose StatementEnd