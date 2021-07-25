package database

import (
	"context"
	"fmt"

	"github.com/tahmooress/motor-shop/internal/entities/models"
	"github.com/tahmooress/motor-shop/internal/pkg/query"
	"github.com/tahmooress/motor-shop/internal/port/dto/dtoshoptrades"
)

func (m *Mysql) GetShopSells(ctx context.Context, request *dtoshoptrades.Request) (*dtoshoptrades.Response, error) {
	q, err := query.New(m.db)
	if err != nil {
		return nil, fmt.Errorf("mysql >> GetShopBuys >> %w", err)
	}

	q.Select = []string{"factors.customer_id", "factor_number", "factors.total_amount",
		"factors.payed_amount", "factors.created_at"}
	q.Body = "FROM factors "
	q.Where("factors.shop_id", []interface{}{request.ShopID}, "=")
	q.Where("factors.type", []interface{}{"SELL"}, "=")
	q.Sort("desc", "factors.created_at")

	var response dtoshoptrades.Response

	response.Data = make([]models.ShopTrades, 0)

	meta, err := q.Exec(ctx, &response.Data)
	if err != nil {
		return nil, fmt.Errorf("mysql >> GetShopBuys >> %w", err)
	}

	response.Meta = meta

	return &response, nil
}
