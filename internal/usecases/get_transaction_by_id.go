package usecases

import (
	"context"
	"fmt"

	"github.com/tahmooress/motor-shop/internal/entities/models"
)

func (u *UseCases) GetTransactionByID(ctx context.Context, transactionID models.ID) (*models.Transaction, error) {
	tx, err := u.IDatabase.GetTransactionByID(ctx, transactionID)
	if err != nil {
		return nil, fmt.Errorf("GetTransactionByID >> %w", err)
	}

	return tx, nil
}
