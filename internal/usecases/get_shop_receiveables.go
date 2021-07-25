package usecases

import (
	"context"
	"fmt"

	"github.com/tahmooress/motor-shop/internal/port/dto/dtoshopequity"
)

func (u *UseCases) GetShopReceiveable(ctx context.Context, request *dtoshopequity.Request) (*dtoshopequity.Response, error) {
	response, err := u.IDatabase.GetShopReceiveable(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("GetShopReceiveable >> %w", err)
	}

	return response, nil
}
