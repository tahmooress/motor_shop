package usecases

import (
	"context"
	"fmt"

	"github.com/tahmooress/motor-shop/internal/entities/models"
	"github.com/tahmooress/motor-shop/internal/port/dto/dtogetshops"
)

func (u *UseCases) CreateShop(ctx context.Context, shopName string) (*models.Shop, error) {
	_, err := u.IDatabase.CreateShop(ctx, shopName)
	if err != nil {
		return nil, fmt.Errorf("createShop >> %w", err)
	}

	var request dtogetshops.Request

	request.Query.Filter = map[string]map[string][]string{"shops.shop_name": {"=": []string{shopName}}}

	response, err := u.IDatabase.GetShopsList(ctx, &request)
	if err != nil {
		return nil, fmt.Errorf("CreateShop >> %w", err)
	}

	if response.Data == nil || len(response.Data) == 0 {
		return nil, models.ErrUnknown
	}

	return &response.Data[0], nil
}
