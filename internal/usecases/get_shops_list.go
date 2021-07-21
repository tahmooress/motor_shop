package usecases

import (
	"context"
	"fmt"

	"github.com/tahmooress/motor-shop/internal/port/dto/dtogetshops"
)

func (u *UseCases) GetShopsList(ctx context.Context, request *dtogetshops.Request) (*dtogetshops.Response, error) {
	response, err := u.IDatabase.GetShopsList(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("getShopsList >> %w", err)
	}

	return response, nil
}
