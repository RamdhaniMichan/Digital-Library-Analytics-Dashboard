package repository

import (
	"database/sql"
	"digital-library-dashboard/internal/user/model"
)

type Repository interface {
	Create(user *model.User) error
	GetByEmail(email string) (*model.User, error)
	GetByID(id int) (*model.User, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(user *model.User) error {
	return r.db.QueryRow(`INSERT INTO users (name, email, password, role) VALUES ($1, $2, $3, $4) RETURNING id`,
		user.Name, user.Email, user.Password, user.Role).Scan(&user.ID)
}

func (r *repository) GetByEmail(email string) (*model.User, error) {
	user := &model.User{}
	err := r.db.QueryRow(`SELECT id, name, email, password, role FROM users WHERE email=$1`, email).
		Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *repository) GetByID(id int) (*model.User, error) {
	user := &model.User{}
	err := r.db.QueryRow(`SELECT id, name, email, role FROM users WHERE id=$1`, id).
		Scan(&user.ID, &user.Name, &user.Email, &user.Role)
	if err != nil {
		return nil, err
	}
	return user, nil
}
