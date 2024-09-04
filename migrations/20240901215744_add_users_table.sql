-- +goose Up
CREATE TABLE
    users (
        id BIGSERIAL PRIMARY KEY,
        username varchar(50) NOT NULL UNIQUE,
        password varchar(100) NOT NULL
    );
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE users;

-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
