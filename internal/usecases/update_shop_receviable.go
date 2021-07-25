package usecases

import (
	"context"
	"fmt"

	"github.com/tahmooress/motor-shop/internal/entities/models"
)

func (u *UseCases) UpdateShopReceivable(ctx context.Context, equityID models.ID) (*models.ShopEquity, error) {
	equity, err := u.IDatabase.UpdateShopReceivable(ctx, equityID)
	if err != nil {
		return nil, fmt.Errorf("UpdateShopReceivable >> %w", err)
	}

	return equity, nil
}
