package database

import (
	"context"
	"fmt"
	"strings"

	"github.com/tahmooress/motor-shop/internal/entities/models"
	"github.com/tahmooress/motor-shop/internal/port/dto/dtobuy"
)

func (m *Mysql) CreateBuyFactor(ctx context.Context, request *dtobuy.Request) (string, error) {
	tx, err := m.db.Begin()
	if err != nil {
		return "", fmt.Errorf("mysql >> CreateFactor >> db.Begin() >> %w", err)
	}

	defer tx.Rollback()

	motorStatment := "INSERT INTO motors(id, model_name, pelak_number," +
		" body_number,color, model_year) VALUES "
	motorVals := ""
	motorArgs := make([]interface{}, 0)

	for i, motor := range request.Motors {
		motorVals += "(?,?,?,?,?,?),"
		request.Motors[i].ID = models.NewID()
		motorArgs = append(motorArgs, request.Motors[i].ID, motor.ModelName, motor.PelakNumber,
			motor.BodyNumber, motor.Color, motor.ModelYear)
	}

	motorStatment += strings.TrimRight(motorVals, ",")

	mStmt, err := tx.PrepareContext(ctx, motorStatment)
	if err != nil {
		return "", fmt.Errorf("mysql >> Preparecontext() >> %w", err)
	}

	defer mStmt.Close()

	_, err = mStmt.ExecContext(ctx, motorArgs...)
	if err != nil {
		return "", fmt.Errorf("mysql >> ExecContext() >> %w", err)
	}

	if request.Customer.ID.ID() == 0 {
		id, err := m.CreateCustomerWithTX(ctx, tx, request.Customer)
		if err != nil {
			return "", fmt.Errorf("mysql >> CreateFactor >> %w", err)
		}

		request.Customer.ID = *id
	}

	factorID := models.NewID()

	factorStmt, err := tx.PrepareContext(ctx, "INSERT INTO factors(id, customer_id, type, factor_number,"+
		"total_amount, payed_amount, created_at) VALUES(?,?,?,?,?,?,?)")
	if err != nil {
		return "", fmt.Errorf("mysql >> CreateFactor >> PrepareContext() >> %w", err)
	}

	defer factorStmt.Close()

	_, err = factorStmt.ExecContext(ctx, factorID, request.Customer.ID, models.BUY, request.FactorNumber,
		request.TotalAmount, request.PayedAmount, request.Date)
	if err != nil {
		return "", fmt.Errorf("mysql >> CreateFactor >> ExecContext() >> %w", err)
	}

	inventoryStmt := "INSERT INTO shop_inventory(id, shop_id, motor_id, factor_id) VALUES "
	inventoryVlas := ""
	inventoryArgs := make([]interface{}, 0)

	for _, motor := range request.Motors {
		inventoryVlas += "(?,?,?,?),"
		inventoryArgs = append(inventoryArgs, models.NewID(), request.ShopID, motor.ID, factorID)
	}

	inventoryStmt += strings.TrimRight(inventoryVlas, ",")

	iStmt, err := tx.PrepareContext(ctx, inventoryStmt)
	if err != nil {
		return "", fmt.Errorf("mysql >> CreateFactor >> PrepareContext() >> %w", err)
	}

	defer iStmt.Close()

	_, err = iStmt.ExecContext(ctx, inventoryArgs...)
	if err != nil {
		return "", fmt.Errorf("mysql >> CreateFactor >> ExecContext() >> %w", err)
	}

	paysStmt := "INSERT INTO shop_payable(id, customer_id, factor_id, shop_id, " +
		"amount, status, clear_date) VALUES "
	payVals := ""
	payArgs := make([]interface{}, 0)

	for _, pay := range request.Equities {
		payVals += "(?,?,?,?,?,?,?),"
		payArgs = append(payArgs, models.NewID(), request.Customer.ID, factorID, request.ShopID,
			pay.Amount, models.DEBTOR, pay.DueDate)
	}

	paysStmt += strings.TrimRight(payVals, ",")

	pStmt, err := tx.PrepareContext(ctx, paysStmt)
	if err != nil {
		return "", fmt.Errorf("mysql >> CreateFactor >> PrepareContext() >> %w", err)
	}

	defer pStmt.Close()

	_, err = pStmt.ExecContext(ctx, payArgs...)
	if err != nil {
		return "", fmt.Errorf("mysql >> CreateFactor >> ExecContext() >> %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return "", fmt.Errorf("mysql >> CreateFactor >> tx.Commit() >> %w", err)
	}

	return request.FactorNumber, nil
}
