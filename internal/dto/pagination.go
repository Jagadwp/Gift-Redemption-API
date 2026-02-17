package dto

type PaginationQuery struct {
	Page    int    `form:"page"`
	Limit   int    `form:"limit"`
	SortBy  string `form:"sort_by"`
	SortDir string `form:"sort_dir"`
}

func (q *PaginationQuery) Normalize() {
	if q.Page <= 0 {
		q.Page = 1
	}
	if q.Limit <= 0 || q.Limit > 100 {
		q.Limit = 10
	}
	if q.SortBy != "avg_rating" {
		q.SortBy = "created_at"
	}
	if q.SortDir != "asc" {
		q.SortDir = "desc"
	}
}
