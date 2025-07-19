package model

import "time"

type Lending struct {
	ID           int        `json:"id"`
	BookID       int        `json:"book_id"`
	MemberID     int        `json:"member_id"`
	BorrowedDate time.Time  `json:"borrowed_date"`
	DueDate      time.Time  `json:"due_date"`
	ReturnDate   *time.Time `json:"return_date"`
	Status       string     `json:"status"`
	CreatedBy    int        `json:"created_by"`
}

type LendingFilter struct {
	MemberID  int       `query:"member_id"`
	BookID    int       `query:"book_id"`
	Status    string    `query:"status"`
	StartDate time.Time `query:"start_date"`
	EndDate   time.Time `query:"end_date"`
}
