package usecases

import (
	"context"
	"fmt"
	"github.com/tahmooress/motor-shop/internal/port/dto/dtogetshopinventory"
)

func (u *UseCases) GetShopInventory(ctx context.Context,
	request *dtogetshopinventory.Request) (*dtogetshopinventory.Response, error) {
	response, err := u.IDatabase.GetShopInventory(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("useCases >> GetShopInventory >> %w", err)
	}

	return response, nil
}
