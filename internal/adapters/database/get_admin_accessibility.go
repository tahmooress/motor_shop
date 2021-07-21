package database

import (
	"context"
	"fmt"
	"github.com/tahmooress/motor-shop/internal/entities/models"
)

func (m *Mysql) GetAdminAccessibility(ctx context.Context, adminID models.ID) ([]string, error) {
	stmt, err := m.db.PrepareContext(ctx, "SELECT shop_name FROM accessibility"+
		"LEFT JOIN shops ON shops.id = accessibility.shop_id WHERE accessibility.admin_id = ?")
	if err != nil {
		return nil, fmt.Errorf("mysql >> GetAdminAccessibility() >> PrepareContext() >> %w", err)
	}

	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, adminID)
	if err != nil {
		return nil, fmt.Errorf("mysql >> GetAdminAccessibility() >> QueryContext() >> %w", err)
	}

	defer rows.Close()

	accessibility := make([]string, 0)

	for rows.Next() {
		var access string

		err = rows.Scan(&access)
		if err != nil {
			return nil, fmt.Errorf("mysql >> GetAdminAccessibility() >> rows.Scan() >> %w", err)
		}

		accessibility = append(accessibility, access)
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("mysql >> GetAdminAccessibility() >> rows.Err() >> %w", err)
	}

	return accessibility, nil
}
