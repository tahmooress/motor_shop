package usecases

import (
	"context"
	"fmt"

	"github.com/tahmooress/motor-shop/internal/port/dto/dtobuy"
)

func (u *UseCases) Buy(ctx context.Context, request *dtobuy.Request) (*dtobuy.Response, error) {
	factorNumber, err := u.IDatabase.CreateBuyFactor(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("buy >> %w", err)
	}

	response, err := u.IDatabase.GetBuyFactorByNumber(ctx, factorNumber)
	if err != nil {
		return nil, fmt.Errorf("buy >> %w", err)
	}

	return response, nil
}
