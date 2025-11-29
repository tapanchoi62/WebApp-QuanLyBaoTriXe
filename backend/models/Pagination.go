package models

type Pagination struct {
	Page       int    `json:"page"`
	PageSize   int    `json:"pageSize"`
	Total      int64  `json:"total"`
	TotalPages int64  `json:"totalPages"`
	Search     string `json:"search,omitempty"`
}