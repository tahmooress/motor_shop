package database

import (
	"context"
	"fmt"

	"github.com/tahmooress/motor-shop/internal/entities/models"
)

func (m *Mysql) GetFactorByNumber(ctx context.Context, factorNumber string, shopID models.ID) (*models.Factor, error) {
	fcStmt, err := m.db.PrepareContext(ctx, "SELECT id,customer_id, type, total_amount,"+
		"payed_amount, created_at, updated_at FROM factors WHERE factor_number = ? AND shop_id = ?")
	if err != nil {
		return nil, fmt.Errorf("mysql >> GetFactorByNumber >> PrepareContext() >> %w", err)
	}

	defer fcStmt.Close()

	var factor models.Factor

	factor.FactorNumber = factorNumber

	var factorType string

	err = fcStmt.QueryRowContext(ctx, factorNumber, shopID).Scan(
		&factor.ID, &factor.Customer.ID, &factorType, &factor.TotalAmount,
		&factor.PayedAmount, &factor.CreatedAt, &factor.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("mysql >> GetFactorByNumber >> Scan() >> %w", err)
	}

	cStmt, err := m.db.PrepareContext(ctx, "SELECT name,last_name,mobile,national_code,company_name,"+
		"created_at, updated_at FROM customers WHERE id = ?")
	if err != nil {
		return nil, fmt.Errorf("mysql >> GetFactorByNumber >> Preparcontext() >> %w", err)
	}

	defer cStmt.Close()

	err = cStmt.QueryRowContext(ctx, factor.Customer.ID).Scan(&factor.Customer.Name, &factor.Customer.LastName,
		&factor.Customer.Mobile, &factor.Customer.NationalCode, &factor.Customer.CompanyName,
		&factor.Customer.CreatedAt, &factor.Customer.UpdatedAt)

	mStmt, err := m.db.PrepareContext(ctx, "SELECT m.id, m.model_name, m.pelak_number, m.body_number, m.color, "+
		"m.model_year, m.created_at, m.updated_at FROM items "+
		"INNER JOIN motors m ON m.pelak_number = items.pelak_number "+
		"WHERE items.factor_number = ?")
	if err != nil {
		return nil, fmt.Errorf("mysql >> GetFactorByNumber >> PrepareContext() >> %w", err)
	}

	defer mStmt.Close()

	rows, err := mStmt.QueryContext(ctx, factorNumber)
	if err != nil {
		return nil, fmt.Errorf("mysql >> GetFactorByNumber >> QueryContext() >> %w", err)
	}

	defer rows.Close()

	factor.Motors = make([]models.Motor, 0)

	for rows.Next() {
		var motor models.Motor

		err = rows.Scan(&motor.ID, &motor.ModelName, &motor.PelakNumber, &motor.BodyNumber,
			&motor.Color, &motor.ModelYear, &motor.CreatedAt, &motor.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("mysql >> GetFactorByNumber >> rows.Scan() >> %w", err)
		}

		factor.Motors = append(factor.Motors, motor)
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("mysql >> GetFactorByNumber >> rows.Err() >> %w", err)
	}

	if factorType == "BUY" {
		eStmt, err := m.db.PrepareContext(ctx, "SELECT p.id, amount, status, clear_date, p.created_at, p.updated_at "+
			"FROM shop_payable p INNER JOIN factors ON p.factor_number = factors.factor_number "+
			"WHERE factors.factor_number = ?")
		if err != nil {
			return nil, fmt.Errorf("mysql >> GetFactorByNumber >> PrepareContext() >> %w", err)
		}

		defer eStmt.Close()

		rows, err := eStmt.QueryContext(ctx, factorNumber)
		if err != nil {
			return nil, fmt.Errorf("mysql >> GetFactorByNumber >> QueryContext() >> %w", err)
		}

		factor.Equities = make([]models.Equity, 0)

		for rows.Next() {
			var equity models.Equity

			err = rows.Scan(&equity.ID, &equity.Amount, &equity.Status,
				&equity.DueDate, &equity.CreatedAt, &equity.UpdatedAt)
			if err != nil {
				return nil, fmt.Errorf("mysql >> GetFactorByNumber >> rows.Scan() >> %w", err)
			}

			factor.Equities = append(factor.Equities, equity)
		}

		err = rows.Err()
		if err != nil {
			return nil, fmt.Errorf("mysql >> GetFactorByNumber >> rows.Err() >> %w", err)
		}
	}

	if factorType == "SELL" {
		eStmt, err := m.db.PrepareContext(ctx, "SELECT r.id, amount, status, clear_date, r.created_at, r.updated_at "+
			"FROM shop_receivable r INNER JOIN factors ON r.factor_number = factors.factor_number "+
			"WHERE factors.factor_number = ?")
		if err != nil {
			return nil, fmt.Errorf("mysql >> GetFactorByNumber >> PrepareContext() >> %w", err)
		}

		defer eStmt.Close()

		rows, err := eStmt.QueryContext(ctx, factorNumber)
		if err != nil {
			return nil, fmt.Errorf("mysql >> GetFactorByNumber >> QueryContext() >> %w", err)
		}

		factor.Equities = make([]models.Equity, 0)

		for rows.Next() {
			var equity models.Equity

			err = rows.Scan(&equity.ID, &equity.Amount, &equity.Status,
				&equity.DueDate, &equity.CreatedAt, &equity.UpdatedAt)
			if err != nil {
				return nil, fmt.Errorf("mysql >> GetFactorByNumber >> rows.Scan() >> %w", err)
			}

			factor.Equities = append(factor.Equities, equity)
		}

		err = rows.Err()
		if err != nil {
			return nil, fmt.Errorf("mysql >> GetFactorByNumber >> rows.Err() >> %w", err)
		}
	}

	return &factor, nil
}
