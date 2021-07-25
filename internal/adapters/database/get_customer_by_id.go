package database

import (
	"context"
	"fmt"

	"github.com/tahmooress/motor-shop/internal/entities/models"
)

func (m *Mysql) GetCustomerByID(ctx context.Context, customerID models.ID) (*models.Customer, error) {
	stmt, err := m.db.PrepareContext(ctx, "SELECT id,name, last_name, mobile,"+
		" national_code, company_name, created_at, updated_at FROM customers"+
		" WHERE id = ?")
	if err != nil {
		return nil, fmt.Errorf("mysql >> GetCustomerByID >> PrepareContext() >> %w", err)
	}

	defer stmt.Close()

	var customer models.Customer

	err = stmt.QueryRowContext(ctx, customerID).Scan(&customer.ID, &customer.Name, &customer.LastName, &customer.Mobile,
		&customer.NationalCode, &customer.CompanyName, &customer.CreatedAt, &customer.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("mysql >> GetCustomerByID >> QueryRowContext() >> %w", err)
	}

	return &customer, nil
}
