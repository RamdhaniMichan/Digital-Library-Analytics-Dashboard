package repository

import (
	"database/sql"
	"digital-library-dashboard/internal/analytics/model"
)

type AnalyticsRepository interface {
	GetAnalytics() (model.AnalyticsResponse, error)
}

type analyticsRepo struct {
	db *sql.DB
}

func NewAnalyticsRepository(db *sql.DB) AnalyticsRepository {
	return &analyticsRepo{db: db}
}

func (r *analyticsRepo) GetAnalytics() (model.AnalyticsResponse, error) {
	var res model.AnalyticsResponse

	err := r.db.QueryRow(`SELECT COUNT(*) FROM books`).Scan(&res.TotalBooks)
	if err != nil {
		return res, err
	}

	err = r.db.QueryRow(`
		SELECT 
			COALESCE(SUM(available_qty), 0),
			COALESCE(SUM(borrowed_qty), 0)
		FROM book_status`).Scan(&res.BooksAvailable, &res.BooksBorrowed)
	if err != nil {
		return res, err
	}

	err = r.db.QueryRow(`SELECT COUNT(*) FROM lendings`).Scan(&res.TotalTransactions)
	if err != nil {
		return res, err
	}

	err = r.db.QueryRow(`SELECT COUNT(*) FROM members WHERE status = 'active'`).Scan(&res.TotalMembers)
	if err != nil {
		return res, err
	}

	return res, nil
}
