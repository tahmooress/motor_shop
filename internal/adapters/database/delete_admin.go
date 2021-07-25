package database

import (
	"context"
	"fmt"
	"github.com/tahmooress/motor-shop/internal/entities/models"
)

func (m *Mysql) DeleteAdmin(ctx context.Context, admin models.Admin) (*models.Admin, error) {
	fetchedAdmin, err := m.GetAdminByID(ctx, admin.ID)
	if err != nil {
		return nil, fmt.Errorf("mysql >> DeleteAdmin >> %w", err)
	}

	stmt, err := m.db.PrepareContext(ctx, "DELETE FROM admin_users WHERE id = ?")
	if err != nil {
		return nil, fmt.Errorf("mysql >> DeleteAdmin >> PrepareContext() >> %w", err)
	}

	_, err = stmt.ExecContext(ctx, admin.ID)
	if err != nil {
		return nil, fmt.Errorf("mysql >> DeleteAdmin >> ExecContext() >> %w", err)
	}

	return fetchedAdmin, nil
}
