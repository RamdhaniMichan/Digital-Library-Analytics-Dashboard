package model

type Book struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Author     string `json:"author"`
	ISBN       string `json:"isbn"`
	Quantity   int    `json:"quantity"`
	CategoryID int    `json:"category_id"`
	CreatedBy  int    `json:"created_by"`
}

type BookWithCategory struct {
	Title     string   `json:"title"`
	Author    string   `json:"author"`
	ISBN      string   `json:"isbn"`
	Quantity  int      `json:"quantity"`
	Category  Category `json:"category"`
	CreatedBy int      `json:"created_by"`
}

type BookFilter struct {
	Title    string `query:"title"`
	Author   string `query:"author"`
	Category int    `query:"category_id"`
}
