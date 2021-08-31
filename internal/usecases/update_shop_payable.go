package usecases

import (
	"context"

	"github.com/tahmooress/motor-shop/internal/entities/models"
)

func (u *UseCases) UpdateShopPayable(ctx context.Context, equityID models.ID) (*models.ShopEquity, error) {
	response, err := u.IDatabase.UpdateShopPayable(ctx, equityID)
	if err != nil {
		return nil, err
	}

	return response, nil
}
