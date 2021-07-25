package database

import (
	"context"
	"fmt"

	"github.com/tahmooress/motor-shop/internal/entities/models"
	"github.com/tahmooress/motor-shop/internal/pkg/query"
	"github.com/tahmooress/motor-shop/internal/port/dto/dtocustomers"
)

func (m *Mysql) GetCustomers(ctx context.Context, request *dtocustomers.Request) (*dtocustomers.Response, error) {
	q, err := query.New(m.db)
	if err != nil {
		return nil, fmt.Errorf("mysql >> GetCustomers >> %w", err)
	}

	q.Select = []string{"id", "name", "last_name", "mobile", "national_code",
		"company_name", "created_at", "updated_at"}
	q.Body = "FROM customers"
	q.QueryFilters = request.Query

	var response dtocustomers.Response

	response.Data = make([]models.Customer, 0)

	meta, err := q.Exec(ctx, &response.Data)
	if err != nil {
		return nil, fmt.Errorf("mysql >> GetCustomers >> %w", err)
	}

	response.Meta = meta

	return &response, nil
}
