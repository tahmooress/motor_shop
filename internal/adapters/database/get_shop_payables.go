package database

import (
	"context"
	"fmt"
	"github.com/tahmooress/motor-shop/internal/entities/models"

	"github.com/tahmooress/motor-shop/internal/pkg/query"
	"github.com/tahmooress/motor-shop/internal/port/dto/dtoshopequity"
)

func (m *Mysql) GetShopPayables(ctx context.Context, request *dtoshopequity.Request) (*dtoshopequity.Response, error) {
	q, err := query.New(m.db)
	if err != nil {
		return nil, fmt.Errorf("mysql >> GetShopPayables >> %w", err)
	}

	q.Select = []string{"shop_payable.id", "shop_payable.customer_id", "shop_payable.factor_number",
		"shop_payable.amount", "shop_payable.status", "shop_payable.clear_date",
		"shop_payable.created_at", "shop_payable.updated_at"}
	q.Body = "FROM shop_payable"
	q.Where("shop_payable.shop_id", []interface{}{request.ShopID}, "eq")
	q.QueryFilters = request.Query
	q.Sort("asc", "shop_payable.clear_date")

	var response dtoshopequity.Response

	response.Equity = make([]models.ShopEquity, 0)

	meta, err := q.Exec(ctx, &response.Equity)
	if err != nil {
		return nil, fmt.Errorf("mysql >> GetShopPayables >> %w", err)
	}

	response.Meta = meta

	return &response, nil
}
