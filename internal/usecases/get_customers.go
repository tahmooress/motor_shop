package usecases

import (
	"context"
	"fmt"
	"github.com/tahmooress/motor-shop/internal/port/dto/dtocustomers"
)

func (u *UseCases) GetCustomers(ctx context.Context, request *dtocustomers.Request) (*dtocustomers.Response, error) {
	response, err := u.IDatabase.GetCustomers(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("GetCustomers >> %w", err)
	}

	return response, nil
}
