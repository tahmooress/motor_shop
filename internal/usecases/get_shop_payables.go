package usecases

import (
	"context"
	"fmt"

	"github.com/tahmooress/motor-shop/internal/port/dto/dtoshopequity"
)

func (u *UseCases) GetShopPayables(ctx context.Context, request *dtoshopequity.Request) (*dtoshopequity.Response, error) {
	response, err := u.IDatabase.GetShopPayables(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("GetShopPayables >> %w", err)
	}

	return response, nil
}
