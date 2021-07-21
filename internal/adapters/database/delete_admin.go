package database

import (
	"context"
	"fmt"
	"github.com/tahmooress/motor-shop/internal/port/dto/dtodeleteadmin"
)

func (m *Mysql) DeleteAdmin(ctx context.Context, request *dtodeleteadmin.Request) (*dtodeleteadmin.Response, error) {
	admin, err := m.GetAdminByID(ctx, request.AdminID)
	if err != nil {
		return nil, fmt.Errorf("mysql >> DeleteAdmin >> %w", err)
	}

	stmt, err := m.db.PrepareContext(ctx, "DELETE FROM admin_users WHERE id = ?")
	if err != nil {
		return nil, fmt.Errorf("mysql >> DeleteAdmin >> PrepareContext() >> %w", err)
	}

	_, err = stmt.ExecContext(ctx, request.AdminID)
	if err != nil {
		return nil, fmt.Errorf("mysql >> DeleteAdmin >> ExecContext() >> %w", err)
	}

	return &dtodeleteadmin.Response{
		Admin: *admin,
	}, nil
}
