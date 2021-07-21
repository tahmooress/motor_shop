package query

import (
	"fmt"
	"strings"
)

// prepareQuery create a select Query from filters that comes from http.

// nolint : funlen, nestif
func (q *Query) prepareQuery(result interface{}) error {
	// create select part of query.
	selct := q.createSelect()

	body := fmt.Sprintf("%s ", q.Body)

	// sanitize filters to remove redundant and invalid fields.
	filters, sorts := q.sanitizeFilters(result, q.Select)

	// create 'where' part of the Query.
	where := q.parseFilters(filters)

	// create 'group by' part of the query.
	groupBy := q.groupBy()

	// create 'order by' part of the query.
	orderBy := q.orderBy(sorts)

	// create limit/offset part of the query.
	limitOffset := q.metaCreator(q.QueryFilters.Page.Size, q.QueryFilters.Page.Number)

	// create meta query to supplement meta data.
	q.metaQuery = q.createMetaQuery(selct, body, where, groupBy, orderBy)

	// set Query filed of Query.
	q.Query = selct + body + where + groupBy + orderBy + limitOffset

	return nil
}

// createSelect creates select portion of Query.
func (q *Query) createSelect() string {
	query := "SELECT "

	// concatenate columns to select Query.
	for i := range q.Select {
		if i == 0 {
			query += q.Select[i]

			continue
		}

		query += fmt.Sprintf(", %s", q.Select[i])
	}

	return query + " "
}

// where takes a where struct and create desired 'where' or 'where in ' clause and return it.
func (q *Query) where(where Where) string {
	whereQuery := ""
	// whereLen is a number to consider if we have to use 'where in' or 'where' clause in query.
	const whereLen = 1

	if len(where.Values) == whereLen {
		whereQuery += fmt.Sprintf("%s %s ? ", where.Field, where.Operation)
		q.args = append(q.args, where.Values[0])
	} else {
		n := len(where.Values)

		whereQuery += fmt.Sprintf("%s IN(%s) ", where.Field, questionMarkSequence(n))

		for _, w := range where.Values {
			q.args = append(q.args, w)
		}
	}

	return whereQuery
}

// orderBy takes a sort then create an 'ORDER BY' portion of a database select Query.

// nolint : nestif
func (q *Query) orderBy(sorts []sort) string {
	var (
		orderBy  string
		sortPart string
	)

	if sorts == nil || len(sorts) == 0 {
		return orderBy
	}

	orderBy = "ORDER BY "

	for _, s := range sorts {
		if s.sortType == "asc" {
			sortPart += fmt.Sprintf("%s ASC ,", s.field)
			continue
		}

		sortPart += fmt.Sprintf("%s DESC ,", s.field)
	}

	return orderBy + strings.TrimLeft(strings.TrimRight(sortPart, ","), ",")
}

func (q *Query) groupBy() string {
	var query string
	// concatenate GROUP BY part of Query if exist.
	if q.GroupBy != "" {
		query += fmt.Sprintf(" GROUP BY %s ", q.GroupBy)
	}

	return query
}
