-- +goose Up
INSERT INTO users (username, password)
VALUES ('admin', '$2a$14$P7N/hEXVuUx2cWokInnK7.3ZE6ZVO0EKnUtbXQKDv/UOmILUk.0VK');
-- +goose StatementBegin
SELECT
    'up SQL query';

-- +goose StatementEnd
-- +goose Down
TRUNCATE TABLE users;

-- +goose StatementBegin
SELECT
    'down SQL query';

-- +goose StatementEnd