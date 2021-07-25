package usecases

import (
	"context"
	"fmt"

	"github.com/tahmooress/motor-shop/internal/entities/models"
)

func (u *UseCases) UpdateShopPayable(ctx context.Context, equityID models.ID) (*models.ShopEquity, error) {
	equity, err := u.IDatabase.UpdateShopPayable(ctx, equityID)
	if err != nil {
		return nil, fmt.Errorf("UpdateShopPayable >> %w", err)
	}

	return equity, nil
}
