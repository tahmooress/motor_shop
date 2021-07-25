package database

import (
	"context"
	"fmt"
	"github.com/tahmooress/motor-shop/internal/entities/models"
	"github.com/tahmooress/motor-shop/internal/pkg/query"
	"github.com/tahmooress/motor-shop/internal/port/dto/dtoshopequity"
)

func (m *Mysql) GetShopReceiveable(ctx context.Context, request *dtoshopequity.Request) (*dtoshopequity.Response, error) {
	q, err := query.New(m.db)
	if err != nil {
		return nil, fmt.Errorf("mysql >> GetShopPayables >> %w", err)
	}

	q.Select = []string{"shop_receivable.id", "shop_receivable.customer_id", "shop_receivable.factor_number",
		"shop_receivable.amount", "shop_receivable.status", "shop_receivable.clear_date",
		"shop_receivable.created_at", "shop_receivable.updated_at"}
	q.Body = "FROM shop_receivable"
	q.Where("shop_receivable.shop_id", []interface{}{request.ShopID}, "=")
	q.QueryFilters = request.Query
	q.Sort("asc", "shop_receivable.clear_date")

	var response dtoshopequity.Response

	response.Equity = make([]models.ShopEquity, 0)

	meta, err := q.Exec(ctx, &response.Equity)
	if err != nil {
		return nil, fmt.Errorf("mysql >> GetShopPayables >> %w", err)
	}

	response.Meta = meta

	return &response, nil
}
