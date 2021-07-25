package database

import (
	"context"
	"fmt"

	"github.com/tahmooress/motor-shop/internal/entities/models"
)

func (m *Mysql) UpdateShopReceivable(ctx context.Context, equityID models.ID) (*models.ShopEquity, error) {
	var (
		factorNumber string
		ShopID       string
		amount       float64
	)

	fStmt, err := m.db.PrepareContext(ctx, "SELECT shop_id, factor_number, amount FROM shop_receivable WHERE id = ?")
	if err != nil {
		return nil, fmt.Errorf("mysql >> UpdateShopReceivable >> PrepareContext() >> %w", err)
	}

	defer fStmt.Close()

	err = fStmt.QueryRowContext(ctx, equityID).Scan(&ShopID, &factorNumber, &amount)
	if err != nil {
		return nil, fmt.Errorf("mysql >> UpdateShopReceivable >> QueryRowContext() >> %w", err)
	}

	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("mysql >> UpdateShopReceivable >> tx.Begin() >> %w", err)
	}

	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, "UPDATE shop_receivable SET status = ? WHERE id = ?")
	if err != nil {
		return nil, fmt.Errorf("mysql >> UpdateShopReceivable >> PrepareContext() >> %w", err)
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, models.CLEAR, equityID)
	if err != nil {
		return nil, fmt.Errorf("mysql >> UpdateShopReceivable >> ExecContext() >> %w", err)
	}

	txStmt, err := tx.PrepareContext(ctx, "INSERT INTO transactions(id, shop_id, type,"+
		" subject, amount, factor_number) VALUES(?,?,?,?,?,?)")
	if err != nil {
		return nil, fmt.Errorf("mysql >> UpdateShopReceivable >> PrepareContext() >> %w", err)
	}

	defer txStmt.Close()

	_, err = txStmt.ExecContext(ctx, models.NewID(), ShopID, models.RECEIVED, models.EQUITY,
		amount, factorNumber)
	if err != nil {
		return nil, fmt.Errorf("mysql >> UpdateShopReceivable >> ExecContext() >> %w", err)
	}

	stm, err := tx.PrepareContext(ctx, "UPDATE factors SET payed_amount = payed_amount + ? WHERE factor_number = ?")
	if err != nil {
		return nil, fmt.Errorf("mysql >> UpdateShopReceivable >> PrepareContext() >> %w", err)
	}

	defer stm.Close()

	_, err = stm.ExecContext(ctx, amount, factorNumber)
	if err != nil {
		return nil, fmt.Errorf("mysql >> UpdateShopReceivable >> ExecContext() >> %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("mysql >> UpdateShopReceivable >> tx.Commit() >> %w", err)
	}

	sstmt, err := m.db.PrepareContext(ctx, "SELECT customer_id, factor_number, amount,status, clear_date,created_at, updated_at "+
		" FROM shop_receivable WHERE id = ?")
	if err != nil {
		return nil, fmt.Errorf("mysql >> UpdateShopReceivable >> PrepareContext() >> %w", err)
	}

	defer sstmt.Close()

	var equity models.ShopEquity

	equity.ID = equityID

	err = sstmt.QueryRowContext(ctx, equityID).Scan(&equity.CustomerID, &equity.FactorNumber, &equity.Amount,
		&equity.Status, &equity.ClearDate, &equity.CreatedAt, &equity.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("mysql >> UpdateShopReceivable >> QueryRowContext() >> %w", err)
	}

	return &equity, nil
}
