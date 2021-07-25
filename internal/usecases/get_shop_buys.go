package usecases

import (
	"context"
	"fmt"

	"github.com/tahmooress/motor-shop/internal/port/dto/dtoshoptrades"
)

func (u *UseCases) GetShopBuys(ctx context.Context, request *dtoshoptrades.Request) (*dtoshoptrades.Response, error) {
	response, err := u.IDatabase.GetShopBuys(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("GetShopBuys >> %w", err)
	}

	return response, nil
}
