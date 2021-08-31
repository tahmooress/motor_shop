package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/tahmooress/motor-shop/internal/entities/models"
)

func (m *Mysql) CreateBuyFactor(ctx context.Context, factor models.Factor, shopID models.ID) (*models.Factor, error) {
	tx, err := m.db.Begin()
	if err != nil {
		return nil, fmt.Errorf("mysql >> CreateFactor >> db.Begin() >> %w", err)
	}

	defer tx.Rollback()

	motorStatment := "INSERT IGNORE INTO motors(id, model_name, pelak_number," +
		" body_number,color, model_year) VALUES "
	motorVals := ""
	motorArgs := make([]interface{}, 0)

	iSmtm := "INSERT INTO items(id,pelak_number, factor_number) VALUES"
	itemVals := ""
	itemsArgs := make([]interface{}, 0)

	for i, motor := range factor.Motors {
		motorVals += "(?,?,?,?,?,?),"
		factor.Motors[i].ID = models.NewID()
		motorArgs = append(motorArgs, factor.Motors[i].ID, motor.ModelName, motor.PelakNumber,
			motor.BodyNumber, motor.Color, motor.ModelYear)

		itemVals += "(?,?,?),"
		itemsArgs = append(itemsArgs, models.NewID(), motor.PelakNumber, factor.FactorNumber)
	}

	motorStatment += strings.TrimRight(motorVals, ",")
	iSmtm += strings.TrimRight(itemVals, ",")

	mStmt, err := tx.PrepareContext(ctx, motorStatment)
	if err != nil {
		return nil, fmt.Errorf("mysql >> Preparecontext() >> %w", err)
	}

	defer mStmt.Close()

	_, err = mStmt.ExecContext(ctx, motorArgs...)
	if err != nil {
		return nil, fmt.Errorf("mysql >> ExecContext() >> %w", err)
	}

	itemsStmt, err := tx.PrepareContext(ctx, iSmtm)
	if err != nil {
		return nil, fmt.Errorf("mysql >> CreateBuyFactor >> PrepareContext >> %w", err)
	}

	defer itemsStmt.Close()

	if factor.Customer.ID.ID() == 0 {
		c, e := m.getCustomerByMobile(ctx, factor.Customer.Mobile)
		if e != nil  {
			if errors.Is(e, sql.ErrNoRows) {
				id, err := m.CreateCustomerWithTX(ctx, tx, factor.Customer)
				if err != nil {
					return nil, fmt.Errorf("mysql >> CreateFactor >> %w", err)
				}

				factor.Customer.ID = *id
			} else {
				return nil, e
			}
		}

		if e == nil {
			factor.Customer.ID = c.ID
		}
	} else if factor.Customer.ID.ID() != 0 {

		customer, err := m.GetCustomerByID(ctx, factor.Customer.ID)
		if err != nil {
			return nil, fmt.Errorf("mysql >> CreateBuyFactor >> %w", err)
		}

		factor.Customer = *customer
	}

	factorID := models.NewID()

	factorStmt, err := tx.PrepareContext(ctx, "INSERT INTO factors(id, customer_id, type, shop_id,factor_number,"+
		"total_amount, payed_amount, created_at) VALUES(?,?,?,?,?,?,?,?)")
	if err != nil {
		return nil, fmt.Errorf("mysql >> CreateFactor >> PrepareContext() >> %w", err)
	}

	defer factorStmt.Close()

	_, err = factorStmt.ExecContext(ctx, factorID, factor.Customer.ID, models.BUY, shopID, factor.FactorNumber,
		factor.TotalAmount, factor.PayedAmount, factor.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("mysql >> CreateFactor >> ExecContext() >> %w", err)
	}

	_, err = itemsStmt.ExecContext(ctx, itemsArgs...)
	if err != nil {
		return nil, fmt.Errorf("mysql >> CreateBuyFactor >> ExecContext() >> %w", err)
	}

	inventoryStmt := "INSERT IGNORE INTO shop_inventory(id, shop_id, motor_id, factor_number) VALUES "
	inventoryVlas := ""
	inventoryArgs := make([]interface{}, 0)

	for _, motor := range factor.Motors {
		inventoryVlas += "(?,?,?,?),"
		inventoryArgs = append(inventoryArgs, models.NewID(), shopID, motor.ID, factor.FactorNumber)
	}

	inventoryStmt += strings.TrimRight(inventoryVlas, ",")

	iStmt, err := tx.PrepareContext(ctx, inventoryStmt)
	if err != nil {
		return nil, fmt.Errorf("mysql >> CreateFactor >> PrepareContext() >> %w", err)
	}

	defer iStmt.Close()

	_, err = iStmt.ExecContext(ctx, inventoryArgs...)
	if err != nil {
		return nil, fmt.Errorf("mysql >> CreateFactor >> ExecContext() >> %w", err)
	}

	paysStmt := "INSERT INTO shop_payable(id, customer_id, factor_number, shop_id, " +
		"amount, status, clear_date) VALUES "
	payVals := ""
	payArgs := make([]interface{}, 0)

	for i, pay := range factor.Equities {
		payVals += "(?,?,?,?,?,?,?),"
		factor.Equities[i].ID = models.NewID()
		factor.Equities[i].Status = models.DEBTOR
		payArgs = append(payArgs, factor.Equities[i].ID, factor.Customer.ID, factor.FactorNumber, shopID,
			pay.Amount, factor.Equities[i].Status, pay.DueDate)
	}

	paysStmt += strings.TrimRight(payVals, ",")

	pStmt, err := tx.PrepareContext(ctx, paysStmt)
	if err != nil {
		return nil, fmt.Errorf("mysql >> CreateFactor >> PrepareContext() >> %w", err)
	}

	defer pStmt.Close()

	_, err = pStmt.ExecContext(ctx, payArgs...)
	if err != nil {
		return nil, fmt.Errorf("mysql >> CreateFactor >> ExecContext() >> %w", err)
	}

	txStmt, err := tx.PrepareContext(ctx, "INSERT INTO transactions(id, shop_id, type, subject,"+
		"amount, description,factor_number,created_at) VALUES(?,?,?,?,"+
		"?,?,?,?)")
	if err != nil {
		return nil, fmt.Errorf("mysql >> CreateBuyFactor >> PrepareContext() >> %w", err)
	}

	defer txStmt.Close()

	_, err = txStmt.ExecContext(ctx, models.NewID(), shopID, models.PAYED, models.EQUITY,
		factor.PayedAmount, "",factor.FactorNumber, factor.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("mysql >> CreateBuyFactor >> ExecContext() >> %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("mysql >> CreateFactor >> tx.Commit() >> %w", err)
	}

	return &models.Factor{
		ID:           factorID,
		FactorNumber: factor.FactorNumber,
		PayedAmount:  factor.PayedAmount,
		TotalAmount:  factor.TotalAmount,
		Motors:       factor.Motors,
		Customer:     factor.Customer,
		Equities:     factor.Equities,
		CreatedAt:    factor.CreatedAt,
		UpdatedAt:    factor.UpdatedAt,
	}, nil
}
