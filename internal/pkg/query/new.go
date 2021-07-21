package query

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	ascending  = "asc"
	descending = "desc"
)

func New(db *sql.DB) (Query, error) {
	config := os.Getenv("DATABASE_MAX_PAGE_SIZE")

	if config == "" {
		config = "100"
	}

	maxSize, err := strconv.Atoi(config)
	if err != nil {
		return Query{}, fmt.Errorf("new os.getEnv() >> strconv.Atoi() >> %w", err)
	}

	return Query{
		db: db,
		// args is a slice of arguments to pass for corresponding
		// 'where' and 'where in' clause in Query.
		args: make([]interface{}, 0),
		// Operations is map for different operation in Query.
		Operations: map[string]string{
			"eq": "=",
			"gt": ">",
			"lt": "<",
			"ge": ">=",
			"le": "<=",
		},
		maxPageSize: maxSize,
	}, nil
}

// Where is an interface for Query struct to add additional where statement.
func (q *Query) Where(field string, values []string, operation string) {
	w := Where{
		Field:     field,
		Values:    values,
		Operation: operation,
	}

	q.wheres = append(q.wheres, w)
}

// Sort takes sortType 'asc' or 'desc' with default value of 'asc'
// and a field to create 'order by' part of query by them.
func (q *Query) Sort(sortType string, field string) {
	if strings.ToLower(sortType) != ascending && strings.ToLower(sortType) != descending {
		return
	}

	q.sorts = append(q.sorts, sort{strings.ToLower(sortType), field})
}
