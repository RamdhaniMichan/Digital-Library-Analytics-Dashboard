-- +goose Up
-- +goose StatementBegin
INSERT INTO
    book_status (
        book_id,
        available_qty,
        borrowed_qty
    )
VALUES (1, 12, 0),
    (2, 8, 0),
    (3, 5, 0),
    (4, 10, 0),
    (5, 15, 0),
    (6, 7, 0);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd