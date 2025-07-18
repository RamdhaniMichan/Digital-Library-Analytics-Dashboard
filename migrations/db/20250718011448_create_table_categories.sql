-- +goose Up
-- +goose StatementBegin
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE
);

CREATE UNIQUE INDEX idx_categories_name_ci ON categories (LOWER(name));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE categories;
-- +goose StatementEnd