package usecases

import (
	"context"
	"fmt"

	"github.com/tahmooress/motor-shop/internal/port/dto/dtotransactions"
)

func (u *UseCases) GetShopTransactions(ctx context.Context, request *dtotransactions.Request) (*dtotransactions.Response, error) {
	response, err := u.IDatabase.GetShopTransactions(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("GetShopTransactions >> %w", err)
	}

	return response, nil
}
