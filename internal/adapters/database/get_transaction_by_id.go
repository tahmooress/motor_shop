package database

import (
	"context"
	"fmt"

	"github.com/tahmooress/motor-shop/internal/entities/models"
)

func (m *Mysql) GetTransactionByID(ctx context.Context, transactionID models.ID) (*models.Transaction, error) {
	stmt, err := m.db.PrepareContext(ctx, "SELECT id, shop_id, description,factor_number, "+
		"subject, type, amount, created_at, updated_at FROM transactions WHERE id = ?")
	if err != nil {
		return nil, fmt.Errorf("mysql >> GetTransactionByID >> PrepareContext() >> %w", err)
	}

	defer stmt.Close()

	var transaction models.Transaction

	err = stmt.QueryRowContext(ctx, transactionID).Scan(&transaction.ID, &transaction.ShopID, &transaction.Description, &transaction.FactorNumber,
		&transaction.Subject, &transaction.Type, &transaction.Amount, &transaction.CreatedAt, &transaction.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("mysql >> GetTransactionByID >> QueryRowContext() >> %w", err)
	}

	return &transaction, nil
}
