package repository

import (
	"database/sql"
	"digital-library-dashboard/internal/book/model"
)

type Repository interface {
	GetAll() ([]model.Book, error)
	GetByID(id int) (*model.Book, error)
	Create(b model.Book) error
	Update(b model.Book) error
	Delete(id int) error
}

type repo struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repo{db}
}

func (r *repo) GetAll() ([]model.Book, error) {
	rows, err := r.db.Query("SELECT id, title, author, isbn, quantity, category_id, created_by FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []model.Book
	for rows.Next() {
		var b model.Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.ISBN, &b.Quantity, &b.CategoryID, &b.CreatedBy); err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}

func (r *repo) GetByID(id int) (*model.Book, error) {
	row := r.db.QueryRow("SELECT id, title, author, isbn, quantity, category_id, created_by FROM books WHERE id = $1", id)
	var b model.Book
	if err := row.Scan(&b.ID, &b.Title, &b.Author, &b.ISBN, &b.Quantity, &b.CategoryID, &b.CreatedBy); err != nil {
		return nil, err
	}
	return &b, nil
}

func (r *repo) Create(b model.Book) error {
	_, err := r.db.Exec("INSERT INTO books (title, author, isbn, quantity, category_id, created_by) VALUES ($1,$2,$3,$4,$5,$6)",
		b.Title, b.Author, b.ISBN, b.Quantity, b.CategoryID, b.CreatedBy)
	return err
}

func (r *repo) Update(b model.Book) error {
	_, err := r.db.Exec("UPDATE books SET title=$1, author=$2, isbn=$3, quantity=$4, category_id=$5 WHERE id=$6",
		b.Title, b.Author, b.ISBN, b.Quantity, b.CategoryID, b.ID)
	return err
}

func (r *repo) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM books WHERE id=$1", id)
	return err
}
