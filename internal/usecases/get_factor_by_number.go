package usecases

import (
	"context"
	"fmt"
	"github.com/tahmooress/motor-shop/internal/entities/models"
)

func (u *UseCases) GetFactorByNumber(ctx context.Context, factorNumber string,
	shopID models.ID) (*models.Factor, error) {
	factor, err := u.IDatabase.GetFactorByNumber(ctx, factorNumber, shopID)
	if err != nil {
		return nil, fmt.Errorf("getFactorByNumber >> %w", err)
	}

	return factor, nil
}
