package database

import (
	"context"
	"fmt"

	"github.com/tahmooress/motor-shop/internal/entities/models"
	"github.com/tahmooress/motor-shop/internal/pkg/query"
	"github.com/tahmooress/motor-shop/internal/port/dto/dtotransactions"
)

func (m *Mysql) GetShopTransactions(ctx context.Context,
	request *dtotransactions.Request) (*dtotransactions.Response, error) {
	q, err := query.New(m.db)
	if err != nil {
		return nil, fmt.Errorf("mysql >> GetShopTransactions >> %w", err)
	}

	q.Select = []string{"transactions.id", "transactions.factor_number", "transactions.description",
		"transactions.subject", "transactions.type", "transactions.amount", "transactions.created_at", "transactions.updated_at"}
	q.Body = "FROM transactions"
	q.QueryFilters = request.Query
	q.Where("transactions.shop_id", []interface{}{request.ShopID}, "=")
	q.Sort("desc", "transactions.created_at")

	var response dtotransactions.Response

	response.Data = make([]models.Transaction, 0)

	meta, err := q.Exec(ctx, &response.Data)
	if err != nil {
		return nil, fmt.Errorf("mysql >> GetShopTransactions >> %w", err)
	}

	response.Meta = meta

	return &response, nil
}
