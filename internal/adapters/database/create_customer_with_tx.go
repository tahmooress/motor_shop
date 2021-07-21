package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/tahmooress/motor-shop/internal/entities/models"
)

func (m *Mysql) CreateCustomerWithTX(ctx context.Context, tx *sql.Tx, customer models.Customer) (*models.ID, error) {
	Stmt, err := tx.PrepareContext(ctx, "INSERT INTO customers(id, name, last_name, "+
		"mobile, national_code, company_name) VALUES(?,?,?,"+
		"?,?,?)")
	if err != nil {
		return nil, fmt.Errorf("mysql >> PrepareContext() >> %w", err)
	}

	defer Stmt.Close()

	customerID := models.NewID()

	_, err = Stmt.ExecContext(ctx, customerID, customer.Name, customer.LastName,
		customer.Mobile, customer.NationalCode, customer.CompanyName)
	if err != nil {
		return nil, fmt.Errorf("mysql >> ExecContext() >> %w", err)
	}

	return &customerID, err
}
