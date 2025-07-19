package repository

import (
	"database/sql"
	"digital-library-dashboard/internal/lending/model"
	"errors"
	"fmt"
)

type LendingRepository interface {
	Create(l model.Lending) error
	GetAll(page, limit int, filter model.LendingFilter) ([]model.Lending, int, error)
	GetByID(id int) (model.Lending, error)
	Update(l model.Lending) error
	Delete(id int) error
}

type lendingRepository struct {
	db *sql.DB
}

func NewLendingRepository(db *sql.DB) LendingRepository {
	return &lendingRepository{db: db}
}

func (r *lendingRepository) Create(l model.Lending) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	query := `
		INSERT INTO lendings (book_id, member_id, borrowed_date, due_date, return_date, status, created_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err = tx.Exec(query,
		l.BookID, l.MemberID, l.BorrowedDate, l.DueDate, l.ReturnDate, l.Status, l.CreatedBy)
	if err != nil {
		return err
	}

	var availableQty int
	err = tx.QueryRow(`
		UPDATE book_status 
		SET available_qty = available_qty - 1, borrowed_qty = borrowed_qty + 1
		WHERE book_id = $1 AND available_qty > 0
		RETURNING available_qty
	`, l.BookID).Scan(&availableQty)

	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("stok buku kosong atau book_id %d tidak ditemukan", l.BookID)
	}
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		UPDATE books 
		SET quantity = $1
		WHERE id = $2
	`, availableQty, l.BookID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *lendingRepository) GetAll(page, limit int, filter model.LendingFilter) ([]model.Lending, int, error) {
	query := `
        SELECT id, book_id, member_id, borrowed_date, due_date, return_date, status, created_by 
        FROM lendings
        WHERE 1=1`

	args := []interface{}{}
	argPos := 1

	if filter.MemberID > 0 {
		query += fmt.Sprintf(" AND member_id = $%d", argPos)
		args = append(args, filter.MemberID)
		argPos++
	}

	if filter.BookID > 0 {
		query += fmt.Sprintf(" AND book_id = $%d", argPos)
		args = append(args, filter.BookID)
		argPos++
	}

	if filter.Status != "" {
		query += fmt.Sprintf(" AND status = $%d", argPos)
		args = append(args, filter.Status)
		argPos++
	}

	if !filter.StartDate.IsZero() {
		query += fmt.Sprintf(" AND borrowed_date >= $%d", argPos)
		args = append(args, filter.StartDate)
		argPos++
	}

	if !filter.EndDate.IsZero() {
		query += fmt.Sprintf(" AND borrowed_date <= $%d", argPos)
		args = append(args, filter.EndDate)
		argPos++
	}

	query += " ORDER BY id DESC LIMIT $" + fmt.Sprint(argPos) + " OFFSET $" + fmt.Sprint(argPos+1)
	args = append(args, limit, (page-1)*limit)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var lendings []model.Lending
	for rows.Next() {
		var l model.Lending
		err := rows.Scan(&l.ID, &l.BookID, &l.MemberID, &l.BorrowedDate, &l.DueDate, &l.ReturnDate, &l.Status, &l.CreatedBy)
		if err != nil {
			return nil, 0, err
		}
		lendings = append(lendings, l)
	}
	return lendings, len(lendings), nil
}

func (r *lendingRepository) GetByID(id int) (model.Lending, error) {
	var l model.Lending
	err := r.db.QueryRow(
		`SELECT id, book_id, member_id, borrowed_date, due_date, return_date, status, created_by FROM lendings WHERE id = $1`, id).
		Scan(&l.ID, &l.BookID, &l.MemberID, &l.BorrowedDate, &l.DueDate, &l.ReturnDate, &l.Status, &l.CreatedBy)
	return l, err
}

func (r *lendingRepository) Update(l model.Lending) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	var oldBookID int
	var oldStatus string
	err = tx.QueryRow(`SELECT book_id, status FROM lendings WHERE id = $1`, l.ID).Scan(&oldBookID, &oldStatus)
	if err != nil {
		return fmt.Errorf("data peminjaman tidak ditemukan: %w", err)
	}

	_, err = tx.Exec(`
		UPDATE lendings 
		SET book_id = $1, member_id = $2, borrowed_date = $3, due_date = $4, return_date = $5, status = $6
		WHERE id = $7`,
		l.BookID, l.MemberID, l.BorrowedDate, l.DueDate, l.ReturnDate, l.Status, l.ID)
	if err != nil {
		return err
	}

	if oldStatus != "returned" && (l.Status == "returned" || l.Status == "late") {
		var availableQty int
		err = tx.QueryRow(`
			UPDATE book_status 
			SET available_qty = available_qty + 1, borrowed_qty = borrowed_qty - 1
			WHERE book_id = $1 AND borrowed_qty > 0
			RETURNING available_qty
		`, l.BookID).Scan(&availableQty)
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("book_id %d tidak valid atau tidak ada buku yang sedang dipinjam", l.BookID)
		}
		if err != nil {
			return err
		}

		_, err = tx.Exec(`
			UPDATE books 
			SET quantity = $1
			WHERE id = $2
		`, availableQty, l.BookID)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *lendingRepository) Delete(id int) error {
	_, err := r.db.Exec(`DELETE FROM lendings WHERE id = $1`, id)
	return err
}
