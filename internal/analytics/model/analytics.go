package model

type AnalyticsResponse struct {
	TotalBooks        int `json:"total_books"`
	BooksAvailable    int `json:"books_available"`
	BooksBorrowed     int `json:"books_borrowed"`
	TotalTransactions int `json:"total_transactions"`
	TotalMembers      int `json:"total_members"`
}
