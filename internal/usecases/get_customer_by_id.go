package usecases

import (
	"context"
	"fmt"
	"github.com/tahmooress/motor-shop/internal/entities/models"
)

func (u *UseCases) GetCustomerByID(ctx context.Context, customerID models.ID) (*models.Customer, error) {
	customer, err := u.IDatabase.GetCustomerByID(ctx, customerID)
	if err != nil {
		return nil, fmt.Errorf("GetCustomerByID >> %w", err)
	}

	return customer, nil
}
