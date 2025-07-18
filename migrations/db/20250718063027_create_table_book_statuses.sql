-- +goose Up
-- +goose StatementBegin
CREATE TABLE book_status (
    id SERIAL PRIMARY KEY,
    book_id INTEGER NOT NULL UNIQUE REFERENCES books (id) ON DELETE CASCADE,
    available_qty INTEGER NOT NULL DEFAULT 0,
    borrowed_qty INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP
    WITH
        TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP
    WITH
        TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE book_status;
-- +goose StatementEnd