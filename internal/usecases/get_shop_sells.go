package usecases

import (
	"context"
	"fmt"

	"github.com/tahmooress/motor-shop/internal/port/dto/dtoshoptrades"
)

func (u *UseCases) GetShopSells(ctx context.Context, request *dtoshoptrades.Request) (*dtoshoptrades.Response, error) {
	response, err := u.IDatabase.GetShopSells(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("GetShopSells >> %w", err)
	}

	return response, nil
}
