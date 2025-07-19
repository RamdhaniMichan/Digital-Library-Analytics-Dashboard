package utils

type Paginate struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalItems int `json:"total_items"`
	TotalPages int `json:"total_pages"`
}

func NewPaginate(page, limit, totalItems int) *Paginate {
	totalPages := 1
	if limit > 0 {
		totalPages = (totalItems + limit - 1) / limit
	}
	return &Paginate{
		Page:       page,
		Limit:      limit,
		TotalItems: totalItems,
		TotalPages: totalPages,
	}
}

func (p *Paginate) Offset() int {
	if p.Page <= 0 || p.Limit <= 0 {
		return 0
	}
	return (p.Page - 1) * p.Limit
}

func (p *Paginate) HasNext() bool {
	return p.Page < p.TotalPages
}

func (p *Paginate) HasPrevious() bool {
	return p.Page > 1
}

func (p *Paginate) NextPage() int {
	if p.HasNext() {
		return p.Page + 1
	}
	return p.Page
}

func (p *Paginate) PreviousPage() int {
	if p.HasPrevious() {
		return p.Page - 1
	}
	return p.Page
}
