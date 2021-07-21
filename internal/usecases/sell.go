package usecases

import (
	"context"
	"fmt"
	"github.com/tahmooress/motor-shop/internal/port/dto/dtosell"
)

func (u *UseCases) Sell(ctx context.Context, request *dtosell.Request) (*dtosell.Response, error) {
	factorNumber, err := u.IDatabase.CreateSellFactor(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("sell >> %w", err)
	}

	response, err := u.IDatabase.GetSellByNumber(ctx, factorNumber)
	if err != nil {
		return nil, fmt.Errorf("sell >> %w", err)
	}

	return response, nil
}
