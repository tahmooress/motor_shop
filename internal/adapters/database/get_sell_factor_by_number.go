package database
//
//import (
//	"context"
//	"fmt"
//
//	"github.com/tahmooress/motor-shop/internal/entities/models"
//	"github.com/tahmooress/motor-shop/internal/port/dto/dtosell"
//)
//
//func (m *Mysql) GetSellFactorByNumber(ctx context.Context, factorNumber string) (*dtosell.Response, error) {
//	fcStmt, err := m.db.PrepareContext(ctx, "SELECT f.id, factor_number, total_amount,"+
//		" payed_amount, f.created_at, f.updated_at,"+
//		" s.id, name, last_name,"+
//		" mobile, national_code, company_name, "+
//		"s.created_at, s.updated_at"+
//		" FROM factors f INNER JOIN customers s ON f.customer_id = s.id "+
//		"WHERE factor_number = ?")
//	if err != nil {
//		return nil, fmt.Errorf("mysql >> GetSellFactorByNumber >> PrepareContext() >> %w", err)
//	}
//
//	defer fcStmt.Close()
//
//	var response dtosell.Response
//
//	err = fcStmt.QueryRowContext(ctx, factorNumber).Scan(
//		&response.ID, &response.FactorNumber, &response.TotalAmount,
//		&response.PayedAmount, &response.CreatedAt, &response.UpdatedAt,
//		&response.Customer.ID, &response.Customer.Name, &response.Customer.LastName,
//		&response.Customer.Mobile, &response.Customer.NationalCode, &response.Customer.CompanyName,
//		&response.Customer.CreatedAt, &response.Customer.UpdatedAt)
//	if err != nil {
//		return nil, fmt.Errorf("mysql >> GetFactorByNumber >> Scan() >> %w", err)
//	}
//
//	mStmt, err := m.db.PrepareContext(ctx, "SELECT m.id, model_name, pelak_number, body_number, color, "+
//		"model_year, m.created_at, m.updated_at FROM shop_inventory "+
//		"INNER JOIN motors m ON m.id = shop_inventory.motor_id "+
//		"WHERE shop_inventory.factor_id = ?")
//	if err != nil {
//		return nil, fmt.Errorf("mysql >> GetSellFactorByNumber >> PrepareContext() >> %w", err)
//	}
//
//	defer mStmt.Close()
//
//	rows, err := mStmt.QueryContext(ctx, response.ID)
//	if err != nil {
//		return nil, fmt.Errorf("mysql >> GetFactorByNumber >> QueryContext() >> %w", err)
//	}
//
//	defer rows.Close()
//
//	response.Motors = make([]models.Motor, 0)
//
//	for rows.Next() {
//		var motor models.Motor
//
//		err = rows.Scan(&motor.ID, &motor.ModelName, &motor.PelakNumber, &motor.BodyNumber,
//			&motor.Color, &motor.ModelYear, &motor.CreatedAt, &motor.UpdatedAt)
//		if err != nil {
//			return nil, fmt.Errorf("mysql >> GetFactorByNumber >> rows.Scan() >> %w", err)
//		}
//
//		response.Motors = append(response.Motors, motor)
//	}
//
//	err = rows.Err()
//	if err != nil {
//		return nil, fmt.Errorf("mysql >> GetSellFactorByNumber >> rows.Err() >> %w", err)
//	}
//
//	eStmt, err := m.db.PrepareContext(ctx, "SELECT r.id, amount, status, clear_date, r.created_at, r.updated_at "+
//		"FROM shop_receivable r INNER JOIN factors ON r.factor_id = factors.id "+
//		"WHERE r.factor_id = ?")
//	if err != nil {
//		return nil, fmt.Errorf("mysql >> GetFactorByNumber >> PrepareContext() >> %w", err)
//	}
//
//	defer eStmt.Close()
//
//	r, err := eStmt.QueryContext(ctx, response.ID)
//	if err != nil {
//		return nil, fmt.Errorf("mysql >> GetFactorByNumber >> QueryContext() >> %w", err)
//	}
//
//	defer r.Close()
//
//	response.Equities = make([]dtosell.Equity, 0)
//
//	for r.Next() {
//		var equity dtosell.Equity
//
//		err = r.Scan(&equity.ID, &equity.Amount, &equity.Status, &equity.DueDate, &equity.CreatedAt, &equity.UpdatedAt)
//		if err != nil {
//			return nil, fmt.Errorf("mysql >> GetFactorByNumber >> rows.Scan() >> %w", err)
//		}
//
//		response.Equities = append(response.Equities, equity)
//	}
//
//	err = r.Err()
//	if err != nil {
//		return nil, fmt.Errorf("mysql >> GetFactorByNumber >> rows.Err() >> %w", err)
//	}
//
//	return &response, err
//}
