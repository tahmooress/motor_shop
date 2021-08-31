package database

import (
	"context"

	"github.com/tahmooress/motor-shop/internal/entities/models"
)

func (m *Mysql) getCustomerByMobile(ctx context.Context, mobile string) (*models.Customer, error){
	stmt, err := m.db.PrepareContext(ctx, "SELECT id,name, last_name, mobile, national_code, " +
		" company_name FROM customers WHERE mobile = ?")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var customer models.Customer

	err = stmt.QueryRowContext(ctx, mobile).Scan(&customer.ID, &customer.Name, &customer.LastName, &customer.Mobile,
		&customer.NationalCode, &customer.CompanyName)
	if err != nil {
		return nil,err
	}

	return &customer, nil
}
