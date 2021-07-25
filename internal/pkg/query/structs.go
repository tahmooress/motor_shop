package query

import (
	"database/sql"
	"github.com/tahmooress/motor-shop/internal/pkg/server"
)

type Where struct {
	Field     string
	Operation string
	//Values    []string
	Values []interface{}
}

type sort struct {
	sortType string
	field    string
}

type Query struct {
	Select       []string
	Body         string
	Query        string
	metaQuery    string
	GroupBy      string
	Operations   map[string]string
	wheres       []Where
	sorts        []sort
	QueryFilters server.Query
	meta         Meta
	db           *sql.DB
	args         []interface{}
	maxPageSize  int
}

type Meta struct {
	PageNumber int `json:"current_page"`
	PageSize   int `json:"per_page"`
	LastPage   int `json:"last_page"`
	Total      int `json:"total"`
}
