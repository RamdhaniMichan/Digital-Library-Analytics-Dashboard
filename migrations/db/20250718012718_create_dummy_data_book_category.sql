-- +goose Up
-- +goose StatementBegin
INSERT INTO
    categories (name)
VALUES ('Fiction'),
    ('Non-Fiction'),
    ('Science'),
    ('History'),
    ('Technology'),
    ('Biography');

INSERT INTO
    books (
        title,
        author,
        isbn,
        quantity,
        category_id,
        created_by,
        created_at,
        updated_at
    )
VALUES (
        'The Silent Patient',
        'Alex Michaelides',
        '9781250301697',
        12,
        1,
        1,
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP
    ),
    (
        'Sapiens: A Brief History of Humankind',
        'Yuval Noah Harari',
        '9780062316097',
        8,
        2,
        2,
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP
    ),
    (
        'Brief Answers to the Big Questions',
        'Stephen Hawking',
        '9781984819192',
        5,
        3,
        3,
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP
    ),
    (
        'A People’s History of the United States',
        'Howard Zinn',
        '9780062397348',
        10,
        4,
        1,
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP
    ),
    (
        'Clean Code',
        'Robert C. Martin',
        '9780132350884',
        15,
        5,
        2,
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP
    ),
    (
        'Becoming',
        'Michelle Obama',
        '9781524763138',
        7,
        6,
        3,
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP
    );

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd