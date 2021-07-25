package database

import (
	"context"
	"fmt"

	"github.com/tahmooress/motor-shop/internal/entities/models"
)

func (m *Mysql) CreateShop(ctx context.Context, shopName string) (*models.ID, error) {
	stmt, err := m.db.PrepareContext(ctx, "INSERT IGNORE INTO shops(id, shop_name) VALUES(?,?)")
	if err != nil {
		return nil, fmt.Errorf("mysql >> CreateShop >> PrepareContext() >> %w", err)
	}

	defer stmt.Close()

	id := models.NewID()

	result, err := stmt.ExecContext(ctx, id, shopName)
	if err != nil {
		return nil, fmt.Errorf("mysql >> CreateShop >> ExecContext() >> %w", err)
	}

	if n, err := result.RowsAffected(); err != nil || n == 0 {
		return nil, models.ErrShopAlreadyExist
	}

	return &id, nil
}
