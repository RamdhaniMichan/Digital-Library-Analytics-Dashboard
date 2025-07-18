package repository

import (
	"context"
	"database/sql"
	"digital-library-dashboard/internal/analytics/model"
)

type AnalyticsRepository interface {
	GetAnalytics(ctx context.Context) (model.AnalyticsResponse, error)
}

type analyticsRepo struct {
	db *sql.DB
}

func NewAnalyticsRepository(db *sql.DB) AnalyticsRepository {
	return &analyticsRepo{db: db}
}

func (r *analyticsRepo) GetAnalytics(ctx context.Context) (model.AnalyticsResponse, error) {
	var res model.AnalyticsResponse

	err := r.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM books`).Scan(&res.TotalBooks)
	if err != nil {
		return res, err
	}

	err = r.db.QueryRowContext(ctx, `
		SELECT 
			COALESCE(SUM(available_qty), 0),
			COALESCE(SUM(borrowed_qty), 0)
		FROM book_status`).Scan(&res.BooksAvailable, &res.BooksBorrowed)
	if err != nil {
		return res, err
	}

	err = r.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM lending`).Scan(&res.TotalTransactions)
	if err != nil {
		return res, err
	}

	err = r.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM members WHERE status = 'active'`).Scan(&res.TotalMembers)
	if err != nil {
		return res, err
	}

	return res, nil
}
