package database

import (
	"context"
	"fmt"
	"strings"

	"github.com/tahmooress/motor-shop/internal/entities/models"
	"github.com/tahmooress/motor-shop/internal/port/dto/dtosell"
)

func (m *Mysql) CreateSellFactor(ctx context.Context, request *dtosell.Request) (string, error) {
	tx, err := m.db.Begin()
	if err != nil {
		return "", fmt.Errorf("mysql >> CreateSellFactor >> tx.Begin() >> %w", err)
	}

	defer tx.Rollback()

	if request.Customer.ID.ID() == 0 {
		id, err := m.CreateCustomerWithTX(ctx, tx, request.Customer)
		if err != nil {
			return "", fmt.Errorf("mysql >> CreateFactor >> %w", err)
		}

		request.Customer.ID = *id
	}

	factorStmt, err := tx.PrepareContext(ctx, "INSERT INTO factors(id, customer_id, type, factor_number,"+
		"total_amount, payed_amount, created_at) VALUES(?,?,?,?,?,?,?)")
	if err != nil {
		return "", fmt.Errorf("mysql >> CreateFactor >> PrepareContext() >> %w", err)
	}

	defer factorStmt.Close()

	factorID := models.NewID()

	_, err = factorStmt.ExecContext(ctx, factorID, request.Customer.ID, models.SELL, request.FactorNumber,
		request.TotalAmount, request.PayedAmount, request.Date)
	if err != nil {
		return "", fmt.Errorf("mysql >> CreateFactor >> ExecContext() >> %w", err)
	}

	deleteStmt := "DELETE FROM shop_inventory WHERE shop_id = ? AND motor_id IN ("
	dltArgs := make([]interface{}, 0)

	dltArgs = append(dltArgs, request.ShopID)

	for _, motor := range request.Motors {
		deleteStmt += "?,"
		dltArgs = append(dltArgs, motor.ID)
	}

	deleteStmt = strings.TrimRight(deleteStmt, ",") + ")"

	dStmt, err := tx.PrepareContext(ctx, deleteStmt)
	if err != nil {
		return "", fmt.Errorf("mysql >> PrepareContext() >> %w", err)
	}

	defer dStmt.Close()

	result, err := dStmt.ExecContext(ctx, dltArgs...)
	if err != nil {
		return "", fmt.Errorf("mysql >> ExecContext() >> %w", err)
	}

	n, err := result.RowsAffected()
	if err != nil {
		return "", fmt.Errorf("mysql >> RowsAffected >> %w", err)
	}

	if n == 0 {
		return "", models.ErrMotorIsNotExist
	}

	recString := "INSERT INTO shop_receivable(id, customer_id, factor_id, shop_id, " +
		"amount, status, clear_date) VALUES "
	recVals := ""
	recArgs := make([]interface{}, 0)

	for _, receive := range request.Equities {
		recVals += "(?,?,?,?,?,?,?),"
		recArgs = append(recArgs, models.NewID(), request.Customer.ID, factorID, request.ShopID,
			receive.Amount, models.DEBTOR, receive.DueDate)
	}

	recString += strings.TrimRight(recVals, ",")

	recStmt, err := tx.PrepareContext(ctx, recString)
	if err != nil {
		return "", fmt.Errorf("mysql >> CreateSellFactor >> PrepareContext() >> %w", err)
	}

	defer recStmt.Close()

	_, err = recStmt.ExecContext(ctx, recArgs...)
	if err != nil {
		return "", fmt.Errorf("mysql >> CreateSellFactor >> ExecContext() >> %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return "", fmt.Errorf("mysql >> CreateSellFactor >> tx.Commit() >> %w", err)
	}

	return request.FactorNumber, nil
}
