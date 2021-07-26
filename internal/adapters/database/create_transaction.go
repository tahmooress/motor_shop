package database

import (
	"context"
	"fmt"

	"github.com/tahmooress/motor-shop/internal/entities/models"
)

func (m *Mysql) CreateTransaction(ctx context.Context, transaction models.Transaction) (*models.ID, error) {
	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("mysql >> CreateTransaction >> tx.BeginTx() >> %w", err)
	}

	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, "INSERT INTO transactions(id, shop_id, type, subject,"+
		"amount, description, factor_number) VALUES(?,?,?,?,?,?,?)")
	if err != nil {
		return nil, fmt.Errorf("mysql >> CreateTransaction >> PrepareContext() >> %w", err)
	}

	defer stmt.Close()

	transactionID := models.NewID()

	_, err = stmt.ExecContext(ctx, transactionID, transaction.ShopID, transaction.Type, transaction.Subject,
		transaction.Amount, transaction.Description, transaction.FactorNumber)
	if err != nil {
		return nil, fmt.Errorf("mysql >> CreateTransaction >> ExecContext() >> %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("mysql >> CreateTransaction >> tx.Commit() >> %w", err)
	}

	return &transactionID, nil
}
