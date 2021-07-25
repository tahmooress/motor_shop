package database

import (
	"context"
	"fmt"
	"strings"

	"github.com/tahmooress/motor-shop/internal/entities/models"
)

func (m *Mysql) CreateSellFactor(ctx context.Context, factor models.Factor, shopID models.ID) (*models.Factor, error) {
	tx, err := m.db.Begin()
	if err != nil {
		return nil, fmt.Errorf("mysql >> CreateSellFactor >> tx.Begin() >> %w", err)
	}

	defer tx.Rollback()

	if factor.Customer.ID.ID() == 0 {
		id, err := m.CreateCustomerWithTX(ctx, tx, factor.Customer)
		if err != nil {
			return nil, fmt.Errorf("mysql >> CreateFactor >> %w", err)
		}

		factor.Customer.ID = *id

	} else if factor.Customer.ID.ID() != 0 {

		customer, _ := m.GetCustomerByID(ctx, factor.Customer.ID)

		factor.Customer = *customer
	}

	factorStmt, err := tx.PrepareContext(ctx, "INSERT INTO factors(id, customer_id, type, shop_id,factor_number,"+
		"total_amount, payed_amount, created_at) VALUES(?,?,?,?,?,?,?,?)")
	if err != nil {
		return nil, fmt.Errorf("mysql >> CreateFactor >> PrepareContext() >> %w", err)
	}

	defer factorStmt.Close()

	factorID := models.NewID()

	_, err = factorStmt.ExecContext(ctx, factorID, factor.Customer.ID, models.SELL, shopID, factor.FactorNumber,
		factor.TotalAmount, factor.PayedAmount, factor.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("mysql >> CreateFactor >> ExecContext() >> %w", err)
	}

	itemsStmt := "INSERT INTO items(id,pelak_number, factor_number) VALUES"
	itemVals := ""
	itemsArgs := make([]interface{}, 0)

	for _, motor := range factor.Motors {
		itemVals += "(?,?,?),"
		itemsArgs = append(itemsArgs, models.NewID(), motor.PelakNumber, factor.FactorNumber)
	}

	itemsStmt += strings.TrimRight(itemVals, ",")

	iSt, err := tx.PrepareContext(ctx, itemsStmt)
	if err != nil {
		return nil, fmt.Errorf("mysql >> CreateBuyFactor >> PrepareContext >> %w", err)
	}

	defer iSt.Close()

	_, err = iSt.ExecContext(ctx, itemsArgs...)
	if err != nil {
		return nil, fmt.Errorf("mysql >> CreateBuyFactor >> ExecContext() >> %w", err)
	}

	deleteStmt := "DELETE FROM shop_inventory WHERE shop_id = ? AND motor_id IN ("
	dltArgs := make([]interface{}, 0)

	dltArgs = append(dltArgs, shopID)

	for _, motor := range factor.Motors {
		deleteStmt += "?,"
		dltArgs = append(dltArgs, motor.ID)
	}

	deleteStmt = strings.TrimRight(deleteStmt, ",") + ")"

	dStmt, err := tx.PrepareContext(ctx, deleteStmt)
	if err != nil {
		return nil, fmt.Errorf("mysql >> PrepareContext() >> %w", err)
	}

	defer dStmt.Close()

	result, err := dStmt.ExecContext(ctx, dltArgs...)
	if err != nil {
		return nil, fmt.Errorf("mysql >> ExecContext() >> %w", err)
	}

	n, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("mysql >> RowsAffected >> %w", err)
	}

	if n == 0 {
		return nil, models.ErrMotorIsNotExist
	}

	recString := "INSERT INTO shop_receivable(id, customer_id, factor_number, shop_id, " +
		"amount, status, clear_date) VALUES "
	recVals := ""
	recArgs := make([]interface{}, 0)

	for i, receive := range factor.Equities {
		recVals += "(?,?,?,?,?,?,?),"
		factor.Equities[i].ID = models.NewID()
		factor.Equities[i].Status = models.DEBTOR
		recArgs = append(recArgs, factor.Equities[i].ID, factor.Customer.ID, factor.FactorNumber, shopID,
			receive.Amount, factor.Equities[i].Status, receive.DueDate)
	}

	recString += strings.TrimRight(recVals, ",")

	recStmt, err := tx.PrepareContext(ctx, recString)
	if err != nil {
		return nil, fmt.Errorf("mysql >> CreateSellFactor >> PrepareContext() >> %w", err)
	}

	defer recStmt.Close()

	_, err = recStmt.ExecContext(ctx, recArgs...)
	if err != nil {
		return nil, fmt.Errorf("mysql >> CreateSellFactor >> ExecContext() >> %w", err)
	}

	txStmt, err := tx.PrepareContext(ctx, "INSERT INTO transactions(id, shop_id, type, subject,"+
		"amount, factor_number,created_at) VALUES(?,?,?,?,"+
		"?,?,?)")
	if err != nil {
		return nil, fmt.Errorf("mysql >> CreateBuyFactor >> PrepareContext() >> %w", err)
	}

	defer txStmt.Close()

	_, err = txStmt.ExecContext(ctx, models.NewID(), shopID, models.RECEIVED, models.EQUITY,
		factor.PayedAmount, factor.FactorNumber, factor.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("mysql >> CreateBuyFactor >> ExecContext() >> %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("mysql >> CreateSellFactor >> tx.Commit() >> %w", err)
	}

	return &models.Factor{
		ID:           factorID,
		FactorNumber: factor.FactorNumber,
		TotalAmount:  factor.TotalAmount,
		PayedAmount:  factor.PayedAmount,
		Motors:       factor.Motors,
		Equities:     factor.Equities,
		Customer:     factor.Customer,
		CreatedAt:    factor.CreatedAt,
		UpdatedAt:    factor.UpdatedAt,
	}, nil
}
