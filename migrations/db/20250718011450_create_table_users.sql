-- +goose Up
-- +goose StatementBegin
CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    isbn VARCHAR(20) UNIQUE NOT NULL,
    quantity INTEGER NOT NULL DEFAULT 0,
    category_id INTEGER NOT NULL,
    created_by INTEGER NOT NULL,
    created_at TIMESTAMP
    WITH
        TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP
    WITH
        TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX idx_books_isbn ON books (isbn);

CREATE INDEX idx_books_category_id ON books (category_id);

CREATE INDEX idx_books_created_by ON books (created_by);

CREATE INDEX idx_books_title ON books (LOWER(title));

ALTER TABLE books
ADD CONSTRAINT fk_books_category FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE books;
-- +goose StatementEnd