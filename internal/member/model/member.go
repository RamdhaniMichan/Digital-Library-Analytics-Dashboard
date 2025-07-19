package model

import "time"

type Member struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	Status     string    `json:"status"`
	JoinedDate time.Time `json:"joined_date"`
}

type MemberFilter struct {
	Name  string `query:"name"`
	Email string `query:"email"`
	Phone string `query:"phone"`
}
