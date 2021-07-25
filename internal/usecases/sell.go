package usecases

import (
	"context"
	"fmt"
	"github.com/tahmooress/motor-shop/internal/entities/models"
)

func (u *UseCases) Sell(ctx context.Context, factor models.Factor, shopID models.ID) (*models.Factor, error) {
	respFactor, err := u.IDatabase.CreateSellFactor(ctx, factor, shopID)
	if err != nil {
		return nil, fmt.Errorf("sell >> %w", err)
	}

	return respFactor, nil
}
