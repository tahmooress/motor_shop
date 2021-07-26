package usecases

import (
	"context"
	"fmt"

	"github.com/tahmooress/motor-shop/internal/entities/models"
)

func (u *UseCases) CreateTransaction(ctx context.Context, transaction models.Transaction) (*models.Transaction, error) {
	if transaction.Subject != models.EXPENSES && transaction.Subject != models.EQUITY {
		return nil, fmt.Errorf("CreateTransaction >> %w", models.ErrTxTypeAndSubject)
	}

	if transaction.Type != models.PAYED && transaction.Type != models.RECEIVED {
		return nil, fmt.Errorf("CreateTransaction >> %w", models.ErrTxTypeAndSubject)
	}

	txID, err := u.IDatabase.CreateTransaction(ctx, transaction)
	if err != nil {
		return nil, fmt.Errorf("CreateTransaction >> %w", err)
	}

	resTX, err := u.IDatabase.GetTransactionByID(ctx, *txID)
	if err != nil {
		return nil, models.ErrUnknown
	}

	return resTX, nil
}
