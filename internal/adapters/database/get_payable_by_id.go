package database

import (
	"context"
	"github.com/tahmooress/motor-shop/internal/entities/models"
)

func (m *Mysql) GetPayableByID(ctx context.Context, id models.ID) (*models.ShopEquity, *models.ID, error){
	stmt, err := m.db.PrepareContext(ctx, "SELECT id, customer_id, factor_number, shop_id,"+
		" amount, status, clear_date FROM shop_payable WHERE id = ?")
	if err != nil {
		return nil, nil, err
	}

	defer stmt.Close()

	var (
		response models.ShopEquity
		shopID   models.ID
	)

	err = stmt.QueryRowContext(ctx, id).Scan(&response.ID, &response.CustomerID, &response.FactorNumber,
		&shopID, &response.Amount, &response.Status, &response.ClearDate)
	if err != nil {
		return nil, nil, err
	}

	return &response, &shopID, nil
}