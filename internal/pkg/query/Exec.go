package query

import (
	"context"
	"fmt"
	"math"

	"github.com/georgysavva/scany/sqlscan"
)

// Exec execute select Query on Query field of Query struct and scan the rows in result
// its parameter result and return a Meta struct that contain data about pagination of Query.
func (q *Query) Exec(ctx context.Context, result interface{}) (Meta, error) {
	err := q.prepareQuery(result)
	if err != nil {
		return Meta{}, err
	}

	q.parseArgs(q.args)

	err = sqlscan.Select(ctx, q.db, result, q.Query, q.args...)
	if err != nil {
		return Meta{}, fmt.Errorf("sqlscan.Select() >> %w", err)
	}

	err = q.db.QueryRow(q.metaQuery, q.args...).Scan(&q.meta.Total)
	if err != nil {
		return Meta{}, fmt.Errorf("exec >> q.db.QueryRow() >> %w", err)
	}

	q.meta.LastPage = int(math.Ceil(float64(q.meta.Total) / float64(q.meta.PageSize)))
	if q.meta.LastPage == 0 {
		q.meta.LastPage = 1
	}

	return q.meta, nil
}
