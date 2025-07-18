-- +goose Up
-- +goose StatementBegin
CREATE TABLE lendings (
    id SERIAL PRIMARY KEY,
    book_id INTEGER NOT NULL REFERENCES books (id) ON DELETE CASCADE,
    member_id INTEGER NOT NULL REFERENCES members (id) ON DELETE CASCADE,
    borrowed_date TIMESTAMP NOT NULL,
    due_date TIMESTAMP NOT NULL,
    return_date TIMESTAMP,
    status VARCHAR(50) NOT NULL,
    created_by INTEGER NOT NULL REFERENCES users (id),
    created_at TIMESTAMP
    WITH
        TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP
    WITH
        TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_lendings_member_id ON lendings (member_id);

CREATE INDEX idx_lendings_status ON lendings (status);

CREATE INDEX idx_lendings_borrowed_date ON lendings (borrowed_date);

CREATE INDEX idx_lendings_due_date ON lendings (due_date);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE lendings;
-- +goose StatementEnd