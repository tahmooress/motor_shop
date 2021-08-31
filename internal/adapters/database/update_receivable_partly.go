package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/tahmooress/motor-shop/internal/entities/models"
	"github.com/tahmooress/motor-shop/internal/port/dto/dtoupdateequities"
)

func (m *Mysql) UpdateReceivablePartly(ctx context.Context, equity models.ShopEquity,
	shopID models.ID, request *dtoupdateequities.Request) (*models.ID, error) {
	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, "Update shop_receivable SET status = ?, amount = ? WHERE id = ?")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, models.CLEAR, request.PayedAmount, equity.ID)
	if err != nil {
		return nil, err
	}

	if n, err := result.RowsAffected(); err != nil || n == 0 {
		return nil, errors.New(fmt.Sprintf("equity not exist >> %s", err))
	}

	newEquity, err := tx.PrepareContext(ctx, "INSERT INTO shop_receivable(id, customer_id, factor_number,"+
		"shop_id, amount, status, clear_date) VALUES(?,?,?,"+
		"?,?,?,?)")
	if err != nil {
		return nil, err
	}

	defer newEquity.Close()

	newEquityID := models.NewID()

	var amount float64 = equity.Amount - request.PayedAmount

	_, err = newEquity.ExecContext(ctx, newEquityID, equity.CustomerID, equity.FactorNumber, shopID, amount,
		models.DEBTOR, request.NextDueDate)
	if err != nil {
		return nil, err
	}

	txStmt, err := tx.PrepareContext(ctx, "INSERT INTO transactions(id, shop_id, type,"+
		" subject, amount, description,factor_number) VALUES(?,?,?,?,?,?,?)")
	if err != nil {
		return nil, fmt.Errorf("mysql >> UpdateShopReceivable >> PrepareContext() >> %w", err)
	}

	defer txStmt.Close()

	_, err = txStmt.ExecContext(ctx, models.NewID(), shopID, models.PAYED, models.EQUITY,
		request.PayedAmount, "", equity.FactorNumber)
	if err != nil {
		return nil, fmt.Errorf("mysql >> UpdateReceivablePartly >> ExecContext() >> %w", err)
	}

	tx.Commit()
	if err != nil {
		return nil, err
	}

	return &newEquityID, nil
}
