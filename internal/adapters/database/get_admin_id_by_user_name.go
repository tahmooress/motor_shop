package database

import (
	"context"
	"fmt"
	"github.com/tahmooress/motor-shop/internal/entities/models"
)

func (m *Mysql) GetAdminIDByUserName(ctx context.Context, userName string) (*models.ID, error) {
	stmt, err := m.db.PrepareContext(ctx, "SELECT id FROM admin_users WHERE user_name = ?")
	if err != nil {
		return nil, fmt.Errorf("mysql >> GetAdmin() >> PrepareContext() %w", err)
	}

	defer stmt.Close()

	var id models.ID

	err = stmt.QueryRowContext(ctx, userName).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("mysql >> GetAdmin() >> QueryRowContext() %w", err)
	}

	return &id, nil
}
