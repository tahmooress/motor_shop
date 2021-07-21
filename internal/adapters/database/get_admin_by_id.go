package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/tahmooress/motor-shop/internal/entities/models"
)

func (m *Mysql) GetAdminByID(ctx context.Context, adminID models.ID) (*models.Admin, error) {
	stmt, err := m.db.PrepareContext(ctx, "SELECT id,user_name,password,created_at,updated_at "+
		"FROM admin_users WHERE id = ?")
	if err != nil {
		return nil, fmt.Errorf("mysql >> GetAdminByID >> PrepareContext() >> %w", err)
	}

	defer stmt.Close()

	var admin models.Admin

	err = stmt.QueryRowContext(ctx, adminID).Scan(&admin.ID, &admin.UserName, &admin.Password,
		&admin.CreatedAt, &admin.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrUserNotFound
		}
		return nil, fmt.Errorf("mysql >> GetAdminByID >> QueryRowContext() >> %w", err)
	}

	stm, err := m.db.PrepareContext(ctx, "SELECT shops.id, shops.shop_name, shops.created_at, shops.updated_at FROM "+
		"shops LEFT JOIN accessibility ON shops.id = accessibility.shop_id WHERE accessibility.admin_id = ?")
	if err != nil {
		return nil, fmt.Errorf("mysql >> GetAdminByID >> PrepareContext() >> %w", err)
	}

	defer stm.Close()

	rows, err := stm.QueryContext(ctx, adminID)
	if err != nil {
		return nil, fmt.Errorf("mysql >> GetAdminByID >> QueryContext() >> %w", err)
	}

	defer rows.Close()

	admin.Accessibility = make([]models.ShopIdentity, 0)

	for rows.Next() {
		var access models.ShopIdentity

		err = rows.Scan(&access.ID, &access.ShopName, &access.CreatedAt, &access.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("mysql >> GetAdminByID >> rows.Scan() >> %w", err)
		}

		admin.Accessibility = append(admin.Accessibility, access)
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("mysql >> GetAdminByID >> rows.Err() >> %w", err)
	}

	return &admin, nil
}
