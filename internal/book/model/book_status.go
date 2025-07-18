package model

type BookStatus struct {
	ID           int `json:"id"`
	BookID       int `json:"book_id"`
	AvailableQty int `json:"available_qty"`
	BorrowedQty  int `json:"borrowed_qty"`
}
