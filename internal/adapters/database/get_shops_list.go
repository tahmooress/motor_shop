package database

import (
	"context"
	"fmt"
	"github.com/tahmooress/motor-shop/internal/entities/models"
	"github.com/tahmooress/motor-shop/internal/pkg/query"
	"github.com/tahmooress/motor-shop/internal/port/dto/dtogetshops"
)

func (m *Mysql) GetShopsList(ctx context.Context, request *dtogetshops.Request) (*dtogetshops.Response, error) {
	var response dtogetshops.Response

	q, err := query.New(m.db)
	if err != nil {
		return nil, fmt.Errorf("mysql >> GetActiveShops >> %w", err)
	}

	q.Select = []string{"id", "shop_name", "created_at", "updated_at"}
	q.Body = "FROM shops"
	q.QueryFilters = request.Query

	response.Data = make([]models.Shop, 0)

	meta, err := q.Exec(ctx, &response.Data)
	if err != nil {
		return nil, fmt.Errorf("mysql >> GetAdmins >> %w", err)
	}

	response.Meta = meta

	return &response, nil
}
