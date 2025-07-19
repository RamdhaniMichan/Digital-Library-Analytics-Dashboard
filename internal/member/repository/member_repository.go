package repository

import (
	"database/sql"
	"digital-library-dashboard/internal/member/model"
	"fmt"
)

type Repository interface {
	Create(m *model.Member) error
	GetByID(id int) (*model.Member, error)
	List(page, limit int, filter model.MemberFilter) ([]*model.Member, int, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(m *model.Member) error {
	return r.db.QueryRow(`INSERT INTO members (user_id, name, email, phone, status, joined_date) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id, status, joined_date`,
		m.UserID, m.Name, m.Email, m.Phone, "active", m.JoinedDate).Scan(&m.ID, &m.Status, &m.JoinedDate)
}

func (r *repository) GetByID(id int) (*model.Member, error) {
	m := &model.Member{}
	err := r.db.QueryRow(`SELECT id, user_id, name, email, phone, status, joined_date FROM members WHERE id=$1`, id).
		Scan(&m.ID, &m.UserID, &m.Name, &m.Email, &m.Phone, &m.Status, &m.JoinedDate)
	return m, err
}

func (r *repository) List(page, limit int, filter model.MemberFilter) ([]*model.Member, int, error) {
	query := `
        SELECT id, user_id, name, email, phone, status, joined_date
        FROM members 
        WHERE 1=1`

	args := []interface{}{}
	argPos := 1

	if filter.Name != "" {
		query += fmt.Sprintf(" AND name ILIKE $%d", argPos)
		args = append(args, "%"+filter.Name+"%")
		argPos++
	}

	if filter.Email != "" {
		query += fmt.Sprintf(" AND email ILIKE $%d", argPos)
		args = append(args, "%"+filter.Email+"%")
		argPos++
	}

	if filter.Phone != "" {
		query += fmt.Sprintf(" AND phone LIKE $%d", argPos)
		args = append(args, "%"+filter.Phone+"%")
		argPos++
	}

	query += " ORDER BY id DESC LIMIT $" + fmt.Sprint(argPos) + " OFFSET $" + fmt.Sprint(argPos+1)
	args = append(args, limit, (page-1)*limit)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var list []*model.Member
	for rows.Next() {
		m := &model.Member{}
		err := rows.Scan(&m.ID, &m.UserID, &m.Name, &m.Email, &m.Phone, &m.Status, &m.JoinedDate)
		if err != nil {
			return nil, 0, err
		}
		list = append(list, m)
	}
	return list, len(list), nil
}
