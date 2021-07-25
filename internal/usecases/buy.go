package usecases

import (
	"context"
	"fmt"
	"github.com/tahmooress/motor-shop/internal/entities/models"
)

func (u *UseCases) Buy(ctx context.Context, factor models.Factor, shopID models.ID) (*models.Factor, error) {
	respFactor, err := u.IDatabase.CreateBuyFactor(ctx, factor, shopID)
	if err != nil {
		return nil, fmt.Errorf("buy >> %w", err)
	}

	return respFactor, nil
}
