package repository

import (
	"database/sql"
	"digital-library-dashboard/internal/book/model"
	"fmt"
)

type Repository interface {
	GetAll(offset, limit int, filter model.BookFilter) ([]model.BookWithCategory, int, error)
	GetByID(id int) (*model.BookWithCategory, error)
	Create(b model.Book) error
	Update(b model.Book) error
	UpdateStatusBook(b model.BookStatus) error
	Delete(id int) error
}

type repo struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repo{db}
}

func (r *repo) GetAll(offset, limit int, filter model.BookFilter) ([]model.BookWithCategory, int, error) {
	query := `
        SELECT b.title, b.author, b.isbn, b.quantity, b.created_by,
               c.id as category_id, c.name as category_name
        FROM books b
        LEFT JOIN categories c ON b.category_id = c.id
        WHERE 1=1
    `
	args := []interface{}{}
	argPos := 1

	if filter.Title != "" {
		query += fmt.Sprintf(" AND b.title ILIKE $%d", argPos)
		args = append(args, "%"+filter.Title+"%")
		argPos++
	}

	if filter.Author != "" {
		query += fmt.Sprintf(" AND b.author ILIKE $%d", argPos)
		args = append(args, "%"+filter.Author+"%")
		argPos++
	}

	if filter.Category > 0 {
		query += fmt.Sprintf(" AND b.category_id = $%d", argPos)
		args = append(args, filter.Category)
		argPos++
	}

	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var books []model.BookWithCategory
	for rows.Next() {
		var b model.BookWithCategory
		if err := rows.Scan(
			&b.Title, &b.Author, &b.ISBN, &b.Quantity, &b.CreatedBy,
			&b.Category.ID, &b.Category.Name,
		); err != nil {
			return nil, 0, err
		}
		books = append(books, b)
	}
	return books, len(books), nil
}

func (r *repo) GetByID(id int) (*model.BookWithCategory, error) {
	query := `
		SELECT 
			b.title, b.author, b.isbn, b.quantity, 
			c.id AS category_id, c.name AS category_name, 
			b.created_by
		FROM books b
		JOIN categories c ON b.category_id = c.id
		WHERE b.id = $1
	`

	row := r.db.QueryRow(query, id)

	var b model.BookWithCategory
	if err := row.Scan(
		&b.Title, &b.Author, &b.ISBN, &b.Quantity,
		&b.Category.ID, &b.Category.Name, &b.CreatedBy,
	); err != nil {
		return nil, err
	}

	return &b, nil
}

func (r *repo) Create(b model.Book) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	err = tx.QueryRow(`
		INSERT INTO books (title, author, isbn, quantity, category_id, created_by)
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`,
		b.Title, b.Author, b.ISBN, b.Quantity, b.CategoryID, b.CreatedBy).Scan(&b.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(`
		INSERT INTO book_status (book_id, available_qty, borrowed_qty)
		VALUES ($1, $2, 0)
	`, b.ID, b.Quantity)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *repo) Update(b model.Book) error {
	_, err := r.db.Exec("UPDATE books SET title=$1, author=$2, isbn=$3, quantity=$4, category_id=$5 WHERE id=$6",
		b.Title, b.Author, b.ISBN, b.Quantity, b.CategoryID, b.ID)
	return err
}

func (r *repo) UpdateStatusBook(b model.BookStatus) error {
	_, err := r.db.Exec("UPDATE book_status SET available_qty=$1, borrowed_qty=$2 WHERE book_id=$3",
		b.AvailableQty, b.BorrowedQty, b.BookID)
	return err
}

func (r *repo) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM books WHERE id=$1", id)
	return err
}
